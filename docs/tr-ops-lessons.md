# (A few) Ops Lessons We All Learn The Hard Way

#（一些）运维课程我们都以艰难的方式学习

January 24th, 2020 From: https://www.netmeister.org/blog/ops-lessons.html

Nope, not another [Falsehoods](https://www.netmeister.org/blog/cs-falsehoods.html) post, but not entirely unlike one. Only here we have a few lessons in [operations](https://www.netmeister.org/blog/defining-operations.html) that we all (eventually) (have to) learn; often the hard way. Why things are the way they are, or what the lessons *mean* is left to the reader to interpret, agree, or disagree with. It's more fun that way. Enjoy!

不，不是另一个 [Falsehoods](https://www.netmeister.org/blog/cs-falsehoods.html) 帖子，但并非完全不同。只有在这里，我们在 [操作](https://www.netmeister.org/blog/defining-operations.html)中有一些我们（最终）（必须)学习的课程；往往是艰难的方式。为什么事情是这样的，或者课程*意味着*是什么留给读者解释、同意或不同意。这样更有趣。享受！

------

1. Email is the worst monitoring and alerting mechanism except for all the others.
2. Absence of a signal is itself a signal.
3. The severity of an incident is measured by the number of rules broken in resolving it.
4. The mobile hotspot you're paying for so you can leave your house while you're oncall only works at home and in the office.
5. The only other person who knows how this works is also on vacation.

1. 电子邮件是除其他所有机制之外最差的监控和警报机制。
2. 没有信号本身就是一个信号。
3. 事件的严重性是通过在解决它时违反的规则数量来衡量的。
4. 您付费购买的移动热点，以便您可以在待命时离开家，只能在家中和办公室使用。
5. 唯一知道这是怎么回事的人也在休假。

6. If a post-mortem follow-up task is not picked up within a week, it's unlikely to be completed at all.
7. That janky script you put together during the outage -- the one that uses `expect(1)` and '`ssh -t -t`' -- now is the foundation of the entire team's toolchest.
8. NTP being off may not be a root cause, but it sure didn't help.
9. [UTC or GTFO](https://teespring.com/utc-or-gtfo).
10. Your infrastructure uses a lot more self-signed certificates than you think. A *lot* more. In places that make you weep.

6. 如果一个尸检后续任务没有在一周内完成，它根本不可能完成。
7. 你在停电期间编写的那个笨拙的脚本——使用`expect(1)`和``ssh -t -t`'的那个脚本——现在是整个团队工具箱的基础。
8. NTP 关闭可能不是根本原因，但肯定没有帮助。
9. [UTC 或 GTFO](https://teespring.com/utc-or-gtfo)。
10. 您的基础设施使用的自签名证书比您想象的要多得多。多很多。在让你哭泣的地方。

11. Self-signed certificates beget long lived certs, which beget lack of certificate validity monitoring, which begets `curl -k`, which begets a lack of certificate deployment automation, which begets self-signed certificates.
12. For any N applications, at most N/2+1 use the same certificate bundle.
13. The system you're troubleshooting doesn't use the one the tool you're troubleshooting it with does.
14. An [API without a reference](https://twitter.com/jschauma/status/1086305575443542017) implementation and command-line client is called a gray box.
15. Restricted shells are not as restricted as you think.

11. 自签名证书导致长期存在的证书，这导致缺乏证书有效性监控，这导致`curl -k`，这导致缺乏证书部署自动化，这导致自签名证书。
12. 对于任何 N 个应用程序，最多 N/2+1 个使用相同的证书包。
13. 您正在排除故障的系统没有使用您正在排除故障的工具。
14. [没有参考的 API](https://twitter.com/jschauma/status/1086305575443542017) 实现和命令行客户端称为灰盒。
15. 受限制的shell并没有你想象的那么受限制。

16. Very few operations are truly idempotent.
17. "Asserting state" beats "monitoring for compliance" any day.
18. [One in a Million is next Tuesday.](https://docs.microsoft.com/en-us/archive/blogs/larryosterman/one-in-a-million-is-next-tuesday)
19. People give talks at conferences not to convince others that their work is awesome and totally worth the time and effort they put in, but themselves.
20. It's ok to use shell for complex stuff; it often times is easier, faster, and still less of a mess than juggling libraries and dependencies.

16. 很少有操作是真正幂等的。
17. “断言状态”胜过“监控合规性”。
18. [百万分之一是下周二。](https://docs.microsoft.com/en-us/archive/blogs/larryosterman/one-in-a-million-is-next-tuesday)
19. 人们在会议上发表演讲不是为了让别人相信他们的工作很棒，完全值得他们投入的时间和精力，而是他们自己。
20. 复杂的东西可以用shell；与处理库和依赖项相比，它通常更容易、更快，而且还没有那么混乱。

21. There's nothing wrong with Perl.
22. Ok, we all at times keep adding `$`, `{`, `}`, and `@` in random places trying to make things work, but still.
23. Serverless isn't.
24. [Y38K](https://en.wikipedia.org/wiki/Year_2038_problem) is [already here](https://twitter.com/jxxf/status/1219009308438024200), it's just not evenly distributed.
25. If you determine "[human error](https://www.netmeister.org/blog/humanerrno.html)" as the root cause, then you're doing it wrong.

21. Perl 没有任何问题。
22. 好吧，我们有时会在随机的地方不断添加`$`、`{`、`}` 和`@`，试图让事情正常运行，但仍然如此。
23. 无服务器不是。
24. [Y38K](https://en.wikipedia.org/wiki/Year_2038_problem)[已经在这里](https://twitter.com/jxxf/status/1219009308438024200)，只是分布不均匀。
25. 如果您确定“[人为错误](https://www.netmeister.org/blog/humanerrno.html)”是根本原因，那么您就做错了。
26. Your network team has a way into the network that your security team doesn't know about.
27. And don't even as much as *mention* the serial console and IPMI networks, but boy are you glad you have 'em.
28. Blocking TCP port 53 traffic leads to very strange failures. Don't.
29. Somewhere in your infrastructure a service you didn't know uses DNS for endpoint discovery in a very surprising way.
30. Do. Not. Monkey. Around. With. `/etc/hosts`.

26. 您的网络团队有一种方式进入您的安全团队不知道的网络。
27. 甚至不要*提及*串行控制台和IPMI网络，但是男孩，你很高兴你拥有它们。
28. 阻塞 TCP 端口 53 流量会导致非常奇怪的故障。别。
29. 在您的基础设施中，某个您不知道的服务以一种非常令人惊讶的方式使用 DNS 进行端点发现。
30. 做。不是。猴。大约。和。 `/etc/hosts`。

31. If you break it, you own it - for now; if you fix it, you own it - forever.
32. Turning it off and on again is actually quite a reasonable way to fix many things.
33. A `README.md` in git is no substitute for a manual page that's shipped with your tool.
34. A search for a document you know exists will only turn up links to documents referencing *but not actually linking to* the one you're looking for.
35. The document you're looking for was marked as obsolete and not migrated to the new content management solution.

31. 如果你打破它，你就拥有它——暂时；如果你修好它，你就拥有它——永远。
32. 将其关闭再打开实际上是解决许多问题的合理方法。
33. git 中的`README.md` 不能替代工具随附的手册页。
34. 搜索您知道存在的文档只会找到引用*但实际上不会链接到*您正在查找的文档的链接。
35. 您要查找的文档被标记为过时且未迁移到新的内容管理解决方案。

36. Sure, your current content management system sucks, but it's still better than the one you're moving to.
37. Nobody knows how git works; everybody simply `rm -fr && git checkout`'s periodically.
38. There are very few network restrictions creative and determined use of `ssh(1)`port forwarding can't overcome.
39. This is both incredibly useful and concerning. 

36. 当然，您当前的内容管理系统很烂，但它仍然比您要迁移到的更好。
37. 没有人知道 git 是如何工作的；每个人都只是定期`rm -fr && git checkout`。
38. 很少有网络限制创造性地使用`ssh(1)`端口转发无法克服。
39. 这既非常有用又令人担忧。

40. It is tempting to jump right into implementing a solution when the right thing may well be to not do the thing that requires the solution in the first place.

40. 当正确的事情很可能是首先不做需要解决方案的事情时，立即开始实施解决方案是很诱人的。

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

41. 永久关闭是 [令人惊讶的困难](https://twitter.com/jschauma/status/1074684924588974080)。
42. “*古代*”在软件和协议方面是一个非常相对的术语。
43. “*Obsolete*”并不意味着它没有被使用和依赖。
44. 数据中心停电前后在线的系统只有相交。一些旧系统上线会立即导致不同的中断。
45. 你的一些最关键的服务是由少数人的工作描述根本没有提到这些服务的。
46. 在最初的“每个人或我 ermahgehrd Slack 宕机”下降之后，生产力在整个停电期间呈线性增长。
47. 您受 [CAP 定理](https://en.wikipedia.org/wiki/CAP_theorem) 约束的次数比您想象的要多得多。 [停止问题](https://en.wikipedia.org/wiki/Halting_problem) 也是个婊子。
48. 当您正在调试的系统尚未收敛时，最终的一致性无济于事。
49. 您正在查看的源代码不是在生产中运行的代码。
50. `strace(1)`/`ktrace(1)` 不会说谎。

51. Unless somebody's been playing `LD_PRELOAD` games.
52. Schrödinger's Backup -- "The condition of any backup is unknown until a restore is attempted." -- is overly optimistic.
53. There's [an xkcd](https://xkcd.com/305/) for the precise situation you find yourself in. (There's also one for at least half of these.)
54. At some point in your career you will implement [half of kerberos](https://twitter.com/jschauma/status/643444266300239873). Poorly.
55. Any sufficiently successful product launch is indistinguishable from a DDoS; any sufficiently advanced user indistinguishable from an attacker.

51. 除非有人一直在玩“LD_PRELOAD”游戏。
52. 薛定谔的备份——“在尝试恢复之前，任何备份的状态都是未知的。” ——过于乐观了。
53. 有 [an xkcd](https://xkcd.com/305/)用于您发现自己所处的确切情况。（至少还有一半的情况。)
54. 在您职业生涯的某个时刻，您将实施 [一半的 kerberos](https://twitter.com/jschauma/status/643444266300239873)。不良。
55. 任何足够成功的产品发布都与 DDoS 没有区别；任何与攻击者无法区分的足够高级的用户。

56. Debugging any sufficiently complex open source product is indistinguishable from reverse engineering a black box.
57. "We've always done it this way." is not a good reason by itself, but there's bound to be one for why.
58. That reason may or may not be valid any longer, however.
59. A junior engineer asking "why" and pointing out the docs don't reflect reality is at least as valuable as the senior engineer working blindly off tribal knowledge.
60. Your herculean efforts to upgrade the OS across your entire fleet completed just in time for the EOL announcement of the version you upgraded to.

56. 调试任何足够复杂的开源产品与逆向工程黑盒没有区别。
57. “我们一直都是这样做的。”本身不是一个很好的理由，但一定有一个原因。
58. 然而，这个理由可能不再有效，也可能不再有效。
59. 询问“为什么”并指出文档不反映现实的初级工程师至少与盲目依赖部落知识的高级工程师一样有价值。
60. 您在整个机群中升级操作系统的艰巨努力在您升级到的版本的 EOL 公告之前及时完成。

61. This phenomenon was first described in Dante's *Inferno* as the Ninth Circle of Hell, Ring Four, aka `RedHat Canto XXXIV`.
62. Containers create at least as many problems as they solve.
63. The most ninja move the expert you hired for that third party black box product you rely on is to say "Let me ping the support team".
64. Somewhere, somebody ran into this *exact* problem, but they never bothered to post a solution.
65. That completely automated solution you set up requires at least three manual steps you didn't document.

61. 这种现象在但丁的 *Inferno* 中首次被描述为地狱的第九圈，四环，又名“RedHat Canto XXXIV”。
62. 容器产生的问题至少与它们解决的问题一样多。
63. 最让你为你依赖的第三方黑匣子产品聘请的专家感动的就是说“让我联系一下支持团队”。
64. 在某处，有人遇到了这个*确切*的问题，但他们从不费心发布解决方案。
65. 您设置的完全自动化的解决方案至少需要三个您没有记录的手动步骤。

66. CAPEX budget always increases, OPEX budget always decreases.
67. CAPEX costs can be reasonably estimated, OPEX costs can only be ballparked.
68. Doubling your time estimate in the hopes of beating expectations won't work because your manager takes your estimate, has a hardy laugh, and then resets it back to what they already promised upchain.
69. Your quarterly planning means bubkes when the next re-org rolls around.
70. Most of your actual work is not covered by your [OKRs](https://www.netmeister.org/blog/okr-distractions.html).

66. CAPEX 预算总是增加，OPEX 预算总是减少。
67. CAPEX 成本可以合理估算，OPEX 成本只能粗略估计。
68. 将你的时间估计加倍以希望超出预期是行不通的，因为你的经理接受了你的估计，哈哈大笑，然后将它重新设置回他们已经承诺的上行链路。
69. 你的季度计划意味着下一次重组来临时。
70. 你的[OKRs](https://www.netmeister.org/blog/okr-distractions.html)没有涵盖你的大部分实际工作。

71. Recursively applying the [Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle) is a surprisingly accurate way to gauge your low hanging fruit, determine your high impact objectives, and ballpark your required effort.
72. Although, to be honest, it [only works in about 80% of cases](https://twitter.com/jschauma/status/962448995690967041).
73. Management will always happily [spend $$$ on outside consultants](https://www.netmeister.org/blog/crazy-like-a-fox.html) to tell them what you've been saying for years.
74. Management will much rather invest in inventing a new, square wheel than fixing an old round one. 

71. 递归地应用[帕累托原则](https://en.wikipedia.org/wiki/Pareto_principle) 是一种非常准确的方法来衡量你的低悬果，确定你的高影响目标，并估计你需要的努力。
72. 虽然，老实说，它[仅在大约 80% 的情况下有效](https://twitter.com/jschauma/status/962448995690967041)。
73. 管理层总是很乐意[在外部顾问身上花费$$](https://www.netmeister.org/blog/crazy-like-a-fox.html) 告诉他们你多年来一直在说的话。
74. 管理层更愿意投资于发明一个新的方形轮子，而不是修理一个旧的圆形轮子。

75. In any organization practicing [continuous integration](https://twitter.com/jschauma/status/1225455309919006720), half of all commits are to fake out CI tests.

75. 在任何实施 [持续集成](https://twitter.com/jschauma/status/1225455309919006720) 的组织中，一半的提交都是伪造 CI 测试。

76. Good software development practices do not always translate well to ops and friends.
77. Mandatory [code reviews](https://twitter.com/jschauma/status/1019410471999467525) do not automatically improve code quality nor reduce the frequency of incidents.
78. Every new paradigm tends to mostly add layers of abstractions; cutting through them and identifying what basic principles continue to apply is half the battle.
79. Real change can only be implemented [above layer 7](https://en.wikipedia.org/wiki/Layer_8).
80. "Prod" is just another name for "staging".

76. 良好的软件开发实践并不总是能很好地转化为操作人员和朋友。
77. 强制[代码审查](https://twitter.com/jschauma/status/1019410471999467525) 不会自动提高代码质量，也不会降低事故频率。
78. 每一种新范式都倾向于增加抽象层；切开它们并确定继续适用的基本原则是成功的一半。
79. 真正的改变只能实现[第7层之上](https://en.wikipedia.org/wiki/Layer_8)。
80. “Prod”只是“staging”的别称。

81. Your source of truth lies.
82. Also: it's incomplete.
83. `pcap` or it didn't happen.
84. `grep(1)` > Splunk (there, I said it)
85. Multithreading is rarely worth the added complexity.
86. Parallelism is not Concurrency.
87. Simplicity is King.
88. Nobody knows what exactly it is you do.

81. 你的真相之源在于。
82. 另外：它是不完整的。
83. `pcap` 或者它没有发生。
84. `grep(1)` > Splunk（我说过了）
85. 多线程很少值得增加复杂性。
86. 并行不是并发。
87. 简单为王。
88. 没人知道你到底在做什么。

January 24th, 2020 

2020 年 1 月 24 日

