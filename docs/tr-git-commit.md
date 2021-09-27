# How to Write a Git Commit Message

# 如何编写 Git 提交消息

Commit messages matter. Here's how to write them well.

提交消息很重要。以下是如何写好它们。

31 Aug 2014• 11 min read

2014 年 8 月 31 日• 11 分钟阅读

_**Contents:** [Introduction](http://chris.beams.io#intro) \| [The Seven Rules](http://chris.beams.io#seven-rules) \| [Tips](http://chris.beams.io#tips)_

_**内容：** [简介](http://chris.beams.io#intro) \|  [七大法则](http://chris.beams.io#seven-rules) \| [提示](http://chris.beams.io#tips)_

## Introduction: Why good commit messages matter

## 介绍：为什么好的提交信息很重要

If you browse the log of any random Git repository, you will probably find its commit messages are more or less a mess. For example, take a look at [these gems](https://github.com/spring-projects/spring-framework/commits/e5f4b49?author=cbeams) from my early days committing to Spring:

如果您浏览任何随机 Git 存储库的日志，您可能会发现它的提交消息或多或少是一团糟。例如，看看我早期致力于 Spring 的 [these gems](https://github.com/spring-projects/spring-framework/commits/e5f4b49?author=cbeams)：

```
$ git log --oneline -5 --author cbeams --before "Fri Mar 26 2009"

e5f4b49 Re-adding ConfigurationPostProcessorTests after its brief removal in r814.@Ignore-ing the testCglibClassesAreLoadedJustInTimeForEnhancement() method as it turns out this was one of the culprits in the recent build breakage.The classloader hacking causes subtle downstream effects, breaking unrelated tests.The test method is still useful, but should only be run on a manual basis to ensure CGLIB is not prematurely classloaded, and should not be run as part of the automated build.
2db0f12 fixed two build-breaking issues: + reverted ClassMetadataReadingVisitor to revision 794 + eliminated ConfigurationPostProcessorTests until further investigation determines why it causes downstream tests to fail (such as the seemingly unrelated ClassPathXmlApplicationContextTests)
147709f Tweaks to package-info.java files
22b25e0 Consolidated Util and MutableAnnotationUtils classes into existing AsmUtils
7f96f57 polishing
```

Yikes. Compare that with these [more recent](https://github.com/spring-projects/spring-framework/commits/5ba3db?author=philwebb) commits from the same repository:

哎呀。将其与来自同一存储库的这些 [最近](https://github.com/spring-projects/spring-framework/commits/5ba3db?author=philwebb) 提交进行比较：

```
$ git log --oneline -5 --author pwebb --before "Sat Aug 30 2014"

5ba3db6 Fix failing CompositePropertySourceTests
84564a0 Rework @PropertySource early parsing logic
e142fd1 Add tests for ImportSelector meta-data
887815f Update docbook dependency and generate epub
ac8326d Polish mockito usage

```

Which would you rather read?

你更愿意读哪个？

The former varies in length and form; the latter is concise and consistent.

前者的长度和形式各不相同；后者简洁一致。

The former is what happens by default; the latter never happens by accident.

前者是默认情况下发生的事情；后者绝不会偶然发生。

While many repositories’ logs look like the former, there are exceptions. The [Linux kernel](https://github.com/torvalds/linux/commits/master) and [Git itself](https://github.com/git/git/commits/master) are great examples. Look at [Spring Boot](https://github.com/spring-projects/spring-boot/commits/master), or any repository managed by [Tim Pope](https://github.com/tpope/vim-pathogen/commits/master).

虽然许多存储库的日志看起来像前者，但也有例外。 [Linux 内核](https://github.com/torvalds/linux/commits/master) 和 [Git 本身](https://github.com/git/git/commits/master) 就是很好的例子。查看 [Spring Boot](https://github.com/spring-projects/spring-boot/commits/master)，或任何由 [Tim Pope] 管理的存储库(https://github.com/tpope/vim-病原体/提交/主人)。

The contributors to these repositories know that a well-crafted Git commit message is the best way to communicate _context_ about a change to fellow developers (and indeed to their future selves). A diff will tell you _what_ changed, but only the commit message can properly tell you _why_. Peter Hutterer [makes this point](http://who-t.blogspot.co.at/2009/12/on-commit-messages.html) well:

这些存储库的贡献者知道，精心设计的 Git 提交消息是与其他开发人员（以及他们未来的自己）交流 _context_ 更改的最佳方式。 diff 会告诉你_what_改变了，但只有提交消息可以正确地告诉你_why_。 Peter Hutterer [说明了这一点](http://who-t.blogspot.co.at/2009/12/on-commit-messages.html) 很好：

> Re-establishing the context of a piece of code is wasteful. We can’t avoid it completely, so our efforts should go to [reducing it](http://www.osnews.com/story/19266/WTFs_m) [as much] as possible. Commit messages can do exactly that and as a result, _a commit message shows whether a developer is a good collaborator_.

> 重新建立一段代码的上下文是一种浪费。我们不能完全避免它，所以我们的努力应该尽可能地[减少它](http://www.osnews.com/story/19266/WTFs_m)。提交消息完全可以做到这一点，因此，_提交消息显示开发人员是否是一个好的合作者_。

If you haven’t given much thought to what makes a great Git commit message, it may be the case that you haven’t spent much time using `git log` and related tools. There is a vicious cycle here: because the commit history is unstructured and inconsistent, one doesn’t spend much time using or taking care of it. And because it doesn’t get used or taken care of, it remains unstructured and inconsistent.

如果你还没有考虑过什么是好的 Git 提交消息，可能是因为你没有花太多时间使用 `git log` 和相关工具。这里有一个恶性循环：因为提交历史是非结构化和不一致的，人们不会花太多时间使用或照顾它。而且因为它没有得到使用或照顾，它仍然是非结构化和不一致的。

But a well-cared for log is a beautiful and useful thing. `git blame`, `revert`, `rebase`, `log`, `shortlog` and other subcommands come to life. Reviewing others’ commits and pull requests becomes something worth doing, and suddenly can be done independently. Understanding why something happened months or years ago becomes not only possible but efficient. 

但是一个精心照料的原木是一件美丽而有用的事情。 `gitblame`、`revert`、`rebase`、`log`、`shortlog` 和其他子命令变得栩栩如生。审查他人的提交和拉取请求成为一件值得做的事情，并且突然可以独立完成。了解为什么几个月或几年前发生的事情不仅可能而且有效。

A project’s long-term success rests (among other things) on its maintainability, and a maintainer has few tools more powerful than his project’s log. It’s worth taking the time to learn how to care for one properly. What may be a hassle at first soon becomes habit, and eventually a source of pride and productivity for all involved.

一个项目的长期成功取决于（除其他外）它的可维护性，维护者没有比他的项目日志更强大的工具了。值得花时间学习如何正确照顾一个人。一开始可能会很麻烦的事情很快就会变成习惯，并最终成为所有相关人员的骄傲和生产力的源泉。

In this post, I am addressing just the most basic element of keeping a healthy commit history: how to write an individual commit message. There are other important practices like commit squashing that I am not addressing here. Perhaps I’ll do that in a subsequent post.

在这篇文章中，我只讨论保持健康提交历史的最基本要素：如何编写单独的提交消息。还有其他重要的实践，比如提交压缩，我没有在这里讨论。也许我会在下一篇文章中做到这一点。

Most programming languages have well-established conventions as to what constitutes idiomatic style, i.e. naming, formatting and so on. There are variations on these conventions, of course, but most developers agree that picking one and sticking to it is far better than the chaos that ensues when everybody does their own thing.

大多数编程语言对于什么构成惯用风格都有完善的约定，即命名、格式等。当然，这些约定有多种变化，但大多数开发人员都同意，选择一个并坚持它比每个人都做自己的事情时随之而来的混乱要好得多。

A team’s approach to its commit log should be no different. In order to create a useful revision history, teams should first agree on a commit message convention that defines at least the following three things:

团队处理提交日志的方法应该没有什么不同。为了创建有用的修订历史，团队应该首先就提交消息约定达成一致，该约定至少定义了以下三件事：

**Style.** Markup syntax, wrap margins, grammar, capitalization, punctuation. Spell these things out, remove the guesswork, and make it all as simple as possible. The end result will be a remarkably consistent log that’s not only a pleasure to read but that actually _does get read_ on a regular basis.

**样式。** 标记语法、换行边距、语法、大小写、标点符号。把这些事情拼出来，消除猜测，让一切尽可能简单。最终结果将是一个非常一致的日志，这不仅是一种阅读的乐趣，而且实际上_确实会定期阅读_。

**Content.** What kind of information should the body of the commit message (if any) contain? What should it _not_ contain?

**内容。** 提交消息的正文（如果有）应包含哪些类型的信息？它应该_不_包含什么？

**Metadata.** How should issue tracking IDs, pull request numbers, etc. be referenced?

**元数据。**应如何引用问题跟踪 ID、拉取请求编号等？

Fortunately, there are well-established conventions as to what makes an idiomatic Git commit message. Indeed, many of them are assumed in the way certain Git commands function. There’s nothing you need to re-invent. Just follow the [seven rules](http://chris.beams.io#seven-rules) below and you’re on your way to committing like a pro.

幸运的是，对于什么是惯用的 Git 提交消息，有完善的约定。事实上，它们中的许多都是以某些 Git 命令的运行方式假设的。您无需重新发明任何东西。只需遵循下面的 [七项规则](http://chris.beams.io#seven-rules)，您就可以像专业人士一样做出承诺。

## The seven rules of a great Git commit message

## 一个伟大的 Git 提交信息的七个规则

> _Keep in mind: [This](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html) [has](https://www.git-scm.com/book/en/v2/Distributed-Git-Contributing-to-a-Project#_commit_guidelines) [all](https://github.com/torvalds/subsurface-for-dirk/blob/master/README.md#contributing) [been](http://who-t.blogspot.co.at/2009/12/on-commit-messages.html) [said](https://github.com/erlang/otp/wiki/writing-good-commit-messages) [before](https://github.com/spring-projects/spring-framework/blob/30bce7/CONTRIBUTING.md#format-commit-messages)._

> _记住：[这个](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)[有](https://www.git-scm.com/book/en/v2/Distributed-Git-Contributing-to-a-Project#_commit_guidelines) [全部](https://github.com/torvalds/subsurface-for-dirk/blob/master/README.md#contributing) [been](http://who-t.blogspot.co.at/2009/12/on-commit-messages.html) [说](https://github.com/erlang/otp/wiki/writing-good-commit-messages) [之前](https://github.com/spring-projects/spring-framework/blob/30bce7/CONTRIBUTING.md#format-commit-messages)._

1. [Separate subject from body with a blank line](http://chris.beams.io#separate)
2. [Limit the subject line to 50 characters](http://chris.beams.io#limit-50)
3. [Capitalize the subject line](http://chris.beams.io#capitalize)
4. [Do not end the subject line with a period](http://chris.beams.io#end)
5. [Use the imperative mood in the subject line](http://chris.beams.io#imperative)
6. [Wrap the body at 72 characters](http://chris.beams.io#wrap-72)
7. [Use the body to explain _what_ and _why_ vs. _how_](http://chris.beams.io#why-not-how)

1. [用空行将主题与正文分开](http://chris.beams.io#separate)
2. [主题行限制在50个字符](http://chris.beams.io#limit-50)
3. [主题行大写](http://chris.beams.io#capitalize)
4. [主题行不要以句号结尾](http://chris.beams.io#end)
5. [在主题行中使用祈使语气](http://chris.beams.io#imperative)
6. [72字包裹正文](http://chris.beams.io#wrap-72)
7. [用身体解释_what_和_why_ vs. _how_](http://chris.beams.io#why-not-how)

For example:

例如：

```
Summarize changes in around 50 characters or less

More detailed explanatory text, if necessary.Wrap it to about 72
characters or so.In some contexts, the first line is treated as the
subject of the commit and the rest of the text as the body.The
blank line separating the summary from the body is critical (unless
you omit the body entirely);various tools like `log`, `shortlog`
and `rebase` can get confused if you run the two together.

Explain the problem that this commit is solving.Focus on why you
are making this change as opposed to how (the code explains that).
Are there side effects or other unintuitive consequences of this
change?Here's the place to explain them.

Further paragraphs come after blank lines.

 - Bullet points are okay, too

 - Typically a hyphen or asterisk is used for the bullet, preceded
by a single space, with blank lines in between, but conventions
vary here

If you use an issue tracker, put references to them at the bottom,
like this:

Resolves: #123
See also: #456, #789

```

### 1\. Separate subject from body with a blank line

### 1. 用空行将主题与正文分开

From the `git commit` [manpage](https://www.kernel.org/pub/software/scm/git/docs/git-commit.html#_discussion): 

从`git commit` [手册页](https://www.kernel.org/pub/software/scm/git/docs/git-commit.html#_discussion)：

> Though not required, it’s a good idea to begin the commit message with a single short (less than 50 character) line summarizing the change, followed by a blank line and then a more thorough description. The text up to the first blank line in a commit message is treated as the commit title, and that title is used throughout Git. For example, Git-format-patch(1) turns a commit into email, and it uses the title on the Subject line and the rest of the commit in the body.

> 虽然不是必需的，但最好以一个简短的（少于 50 个字符）行总结更改开始提交消息，然后是一个空行，然后是更全面的描述。提交消息中直到第一个空行的文本被视为提交标题，并且该标题在整个 Git 中使用。例如，Git-format-patch(1) 将提交转换为电子邮件，它使用主题行上的标题和正文中提交的其余部分。

Firstly, not every commit requires both a subject and a body. Sometimes a single line is fine, especially when the change is so simple that no further context is necessary. For example:

首先，并不是每个提交都需要一个主题和一个主体。有时一行就可以了，尤其是当更改非常简单以至于不需要进一步的上下文时。例如：

```
Fix typo in introduction to user guide

```

Nothing more need be said; if the reader wonders what the typo was, she can simply take a look at the change itself, i.e. use `git show` or `git diff` or `git log -p`.

无需多言；如果读者想知道错别字是什么，她可以简单地看一下更改本身，即使用 `git show` 或 `git diff` 或 `git log -p`。

If you’re committing something like this at the command line, it’s easy to use the `-m` option to `git commit`:

如果你在命令行提交这样的东西，很容易使用 `git commit` 的 `-m` 选项：

```
$ git commit -m"Fix typo in introduction to user guide"

```

However, when a commit merits a bit of explanation and context, you need to write a body. For example:

但是，当提交需要一些解释和上下文时，您需要编写一个正文。例如：

```
Derezz the master control program

MCP turned out to be evil and had become intent on world domination.
This commit throws Tron's disc into MCP (causing its deresolution)
and turns it back into a chess game.

```

Commit messages with bodies are not so easy to write with the `-m` option. You’re better off writing the message in a proper text editor. If you do not already have an editor set up for use with Git at the command line, read [this section of Pro Git](https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration).

使用 `-m` 选项编写带有正文的提交消息并不容易。您最好在适当的文本编辑器中编写消息。如果您还没有在命令行中设置用于 Git 的编辑器，请阅读 [Pro Git 的这一部分](https://git-scm.com/book/en/v2/Customizing-Git-Git-配置)。

In any case, the separation of subject from body pays off when browsing the log. Here’s the full log entry:

无论如何，在浏览日志时，主体与主体的分离是有回报的。这是完整的日志条目：

```
$ git log
commit 42e769bdf4894310333942ffc5a15151222a87be
Author: Kevin Flynn <kevin@flynnsarcade.com>
Date:   Fri Jan 01 00:00:00 1982 -0200

Derezz the master control program

MCP turned out to be evil and had become intent on world domination.
This commit throws Tron's disc into MCP (causing its deresolution)
and turns it back into a chess game.

```

And now `git log --oneline`, which prints out just the subject line:

现在`git log --oneline`，它只打印出主题行：

```
$ git log --oneline
42e769 Derezz the master control program

```

Or, `git shortlog`, which groups commits by user, again showing just the subject line for concision:

或者，`git shortlog`，它按用户对提交进行分组，再次显示简洁的主题行：

```
$ git shortlog
Kevin Flynn (1):
      Derezz the master control program

Alan Bradley (1):
      Introduce security program "Tron"

Ed Dillinger (3):
      Rename chess program to "MCP"
      Modify chess program
      Upgrade chess program

Walter Gibbs (1):
      Introduce protoype chess program

```

There are a number of other contexts in Git where the distinction between subject line and body kicks in—but none of them work properly without the blank line in between.

在 Git 中有许多其他上下文，其中主题行和正文之间的区别开始发挥作用——但如果没有空白行，它们都不能正常工作。

### 2\. Limit the subject line to 50 characters

### 2. 将主题行限制为 50 个字符

50 characters is not a hard limit, just a rule of thumb. Keeping subject lines at this length ensures that they are readable, and forces the author to think for a moment about the most concise way to explain what’s going on.

50 个字符不是硬性限制，只是一个经验法则。保持这个长度的主题行确保它们可读，并迫使作者思考一下最简洁的方式来解释正在发生的事情。

> _Tip: If you’re having a hard time summarizing, you might be committing too many changes at once. Strive for [atomic commits](https://www.freshconsulting.com/atomic-commits/) (a topic for a separate post)._

> _Tip: 如果你很难总结，你可能一次提交了太多的更改。争取[原子提交](https://www.freshconsulting.com/atomic-commits/)（单独帖子的主题)。_

GitHub’s UI is fully aware of these conventions. It will warn you if you go past the 50 character limit:

GitHub 的 UI 完全了解这些约定。如果您超过 50 个字符的限制，它会警告您：

![gh1](https://i.imgur.com/zyBU2l6.png)

And will truncate any subject line longer than 72 characters with an ellipsis:

并且会用省略号截断任何超过 72 个字符的主题行：

![gh2](https://i.imgur.com/27n9O8y.png)

So shoot for 50 characters, but consider 72 the hard limit.

所以拍摄 50 个字符，但考虑 72 个字符的硬限制。

### 3\. Capitalize the subject line

### 3. 大写主题行

This is as simple as it sounds. Begin all subject lines with a capital letter.

这听起来很简单。所有主题行都以大写字母开头。

For example:

例如：

- Accelerate to 88 miles per hour

- 加速到每小时 88 英里

Instead of:

代替：

- accelerate to 88 miles per hour

- 加速到每小时 88 英里

### 4\. Do not end the subject line with a period

### 4. 不要以句点结束主题行

Trailing punctuation is unnecessary in subject lines. Besides, space is precious when you’re trying to keep them to [50 chars or less](https://chris.beams.io/posts/git-commit/#limit-50).

主题行中不需要尾随标点符号。此外，当您试图将它们保持在 [50 个字符或更少](https://chris.beams.io/posts/git-commit/#limit-50) 时，空间是宝贵的。

Example:

例子：

- Open the pod bay doors

- 打开吊舱舱门

Instead of:

代替：

- Open the pod bay doors.

- 打开吊舱舱门。

### 5\. Use the imperative mood in the subject line

### 5. 在主题行中使用祈使语气

_Imperative mood_ just means “spoken or written as if giving a command or instruction”. A few examples:

_祈使语气_仅表示“口语或书面语，好像在发出命令或指示”。几个例子：

- Clean your room
- Close the door
- Take out the trash 

- 收拾你的房间
- 关门
- 把垃圾带出去

Each of the seven rules you’re reading about right now are written in the imperative (“Wrap the body at 72 characters”, etc.).

您现在正在阅读的七个规则中的每一个都是用命令式编写的（“将正文包裹在 72 个字符处”等）。

The imperative can sound a little rude; that’s why we don’t often use it. But it’s perfect for Git commit subject lines. One reason for this is that **Git itself uses the imperative whenever it creates a commit on your behalf**.

命令式听起来有点粗鲁。这就是为什么我们不经常使用它。但它非常适合 Git 提交主题行。原因之一是 **Git 本身在代表您创建提交时使用命令式**。

For example, the default message created when using `git merge` reads:

例如，使用 `git merge` 时创建的默认消息如下：

```
Merge branch 'myfeature'

```

And when using `git revert`:

当使用 `git revert` 时：

```
Revert "Add the thing with the stuff"

This reverts commit cc87791524aedd593cff5a74532befe7ab69ce9d.

```

Or when clicking the “Merge” button on a GitHub pull request:

或者在 GitHub 拉取请求上单击“合并”按钮时：

```
Merge pull request #123 from someuser/somebranch

```

So when you write your commit messages in the imperative, you’re following Git’s own built-in conventions. For example:

因此，当您以命令式编写提交消息时，您将遵循 Git 自己的内置约定。例如：

- Refactor subsystem X for readability
- Update getting started documentation
- Remove deprecated methods
- Release version 1.0.0

- 重构子系统 X 以提高可读性
- 更新入门文档
- 删除不推荐使用的方法
- 发布版本 1.0.0

Writing this way can be a little awkward at first. We’re more used to speaking in the _indicative mood_, which is all about reporting facts. That’s why commit messages often end up reading like this:

这样写一开始可能有点尴尬。我们更习惯于以_指示性情绪_说话，这完全是为了报告事实。这就是为什么提交消息通常最终会读成这样：

- Fixed bug with Y
- Changing behavior of X

- 修复了 Y 的错误
- 改变 X 的行为

And sometimes commit messages get written as a description of their contents:

有时提交消息会被写成对其内容的描述：

- More fixes for broken stuff
- Sweet new API methods

- 对损坏的东西的更多修复
- 甜蜜的新 API 方法

To remove any confusion, here’s a simple rule to get it right every time.

为了消除任何混淆，这里有一个简单的规则，每次都正确。

**A properly formed Git commit subject line should always be able to complete the following sentence**:

**格式正确的 Git 提交主题行应始终能够完成以下句子**：

- If applied, this commit will_your subject line here_

- 如果应用，此提交将_您的主题行在这里_

For example:

例如：

- If applied, this commit will_refactor subsystem X for readability_
- If applied, this commit will_update getting started documentation_
- If applied, this commit will_remove deprecated methods_
- If applied, this commit will_release version 1.0.0_
- If applied, this commit will_merge pull request #123 from user/branch_

- 如果应用，此提交将_重构子系统 X 以提高可读性_
- 如果应用，此提交将_更新入门文档_
- 如果应用，此提交将_删除不推荐使用的方法_
- 如果应用，此提交将_发布版本 1.0.0_
- 如果应用，此提交 will_merge 来自用户/分支的拉取请求 #123_

Notice how this doesn’t work for the other non-imperative forms:

请注意这对其他非命令形式不起作用：

- If applied, this commit will_fixed bug with Y_
- If applied, this commit will_changing behavior of X_
- If applied, this commit will_more fixes for broken stuff_
- If applied, this commit will_sweet new API methods_

- 如果应用，这个提交将_fixed bug with Y_
- 如果应用，此提交将_改变 X_ 的行为
- 如果应用，此提交将_更多修复损坏的东西_
- 如果应用，此提交将_甜蜜的新 API 方法_

> _Remember: Use of the imperative is important only in the subject line. You can relax this restriction when you’re writing the body._

> _记住：命令式的使用仅在主题行中很重要。你在写正文的时候可以放宽这个限制。_

### 6\. Wrap the body at 72 characters

### 6\.将正文包裹在 72 个字符处

Git never wraps text automatically. When you write the body of a commit message, you must mind its right margin, and wrap text manually.

Git 从不自动换行文本。当您编写提交消息的正文时，您必须注意其右边距，并手动换行文本。

The recommendation is to do this at 72 characters, so that Git has plenty of room to indent text while still keeping everything under 80 characters overall.

建议以 72 个字符执行此操作，以便 Git 有足够的空间缩进文本，同时仍将所有内容保持在 80 个字符以内。

A good text editor can help here. It’s easy to configure Vim, for example, to wrap text at 72 characters when you’re writing a Git commit. Traditionally, however, IDEs have been terrible at providing smart support for text wrapping in commit messages (although in recent versions, IntelliJ IDEA has [finally](https://youtrack.jetbrains.com/issue/IDEA-53615) [gotten] (https://youtrack.jetbrains.com/issue/IDEA-53615#comment=27-448299) [better](https://youtrack.jetbrains.com/issue/IDEA-53615#comment=27-446912) about this).

一个好的文本编辑器可以在这里提供帮助。例如，很容易配置 Vim，以便在编写 Git 提交时将文本换行为 72 个字符。然而，传统上，IDE 在为提交消息中的文本换行提供智能支持方面一直很糟糕（尽管在最近的版本中，IntelliJ IDEA [终于](https://youtrack.jetbrains.com/issue/IDEA-53615) [gotten] (https://youtrack.jetbrains.com/issue/IDEA-53615#comment=27-448299) [更好](https://youtrack.jetbrains.com/issue/IDEA-53615#comment=27-446912)关于这个)。

### 7\. Use the body to explain what and why vs. how

### 7. 用身体来解释什么、为什么和如何

This [commit from Bitcoin Core](https://github.com/bitcoin/bitcoin/commit/eb0b56b19017ab5c16c745e6da39c53126924ed6) is a great example of explaining what changed and why:

这个[来自比特币核心的提交](https://github.com/bitcoin/bitcoin/commit/eb0b56b19017ab5c16c745e6da39c53126924ed6)是一个很好的例子，可以解释发生了什么变化以及为什么：

```
commit eb0b56b19017ab5c16c745e6da39c53126924ed6
Author: Pieter Wuille <pieter.wuille@gmail.com>
Date:   Fri Aug 1 22:57:55 2014 +0200

Simplify serialize.h's exception handling

Remove the 'state' and 'exceptmask' from serialize.h's stream
implementations, as well as related methods.

As exceptmask always included 'failbit', and setstate was always
called with bits = failbit, all it did was immediately raise an
exception.Get rid of those variables, and replace the setstate
with direct exception throwing (which also removes some dead
code).

As a result, good() is never reached after a failure (there are
only 2 calls, one of which is in tests), and can just be replaced
by !eof().

fail(), clear(n) and exceptions() are just never called.Delete
them.

```

Take a look at the [full diff](https://github.com/bitcoin/bitcoin/commit/eb0b56b19017ab5c16c745e6da39c53126924ed6) and just think how much time the author is saving fellow and future committers by taking the time to provide this context here and now. If he didn’t, it would probably be lost forever.

看看 [full diff](https://github.com/bitcoin/bitcoin/commit/eb0b56b19017ab5c16c745e6da39c53126924ed6)，想想作者花多少时间在这里提供这个上下文和现在。如果他不这样做，它可能会永远丢失。

In most cases, you can leave out details about how a change has been made. Code is generally self-explanatory in this regard (and if the code is so complex that it needs to be explained in prose, that’s what source comments are for). Just focus on making clear the reasons why you made the change in the first place—the way things worked before the change (and what was wrong with that), the way they work now, and why you decided to solve it the way you did .

在大多数情况下，您可以省略有关如何进行更改的详细信息。代码在这方面通常是不言自明的（如果代码非常复杂以至于需要用散文来解释，这就是源注释的用途）。只需专注于明确您最初进行更改的原因——更改前的工作方式（以及其中有什么问题）、它们现在的工作方式，以及您决定以您的方式解决问题的原因.

The future maintainer that thanks you may be yourself!

感谢您的未来维护者可能就是您自己！

## Tips

##  提示

### Learn to love the command line. Leave the IDE behind.

### 学会爱上命令行。将 IDE 抛在脑后。

For as many reasons as there are Git subcommands, it’s wise to embrace the command line. Git is insanely powerful; IDEs are too, but each in different ways. I use an IDE every day (IntelliJ IDEA) and have used others extensively (Eclipse), but I have never seen IDE integration for Git that could begin to match the ease and power of the command line (once you know it).

出于与 Git 子命令一样多的原因，使用命令行是明智之举。 Git 非常强大； IDE 也是如此，但每个都有不同的方式。我每天都使用 IDE (IntelliJ IDEA) 并广泛使用其他 IDE (Eclipse)，但我从未见过 Git 的 IDE 集成可以开始匹配命令行的易用性和强大功能（一旦您了解它）。

Certain Git-related IDE functions are invaluable, like calling `git rm` when you delete a file, and doing the right stuff with `git` when you rename one. Where everything falls apart is when you start trying to commit, merge, rebase, or do sophisticated history analysis through the IDE.

某些与 Git 相关的 IDE 功能非常宝贵，例如在删除文件时调用 `git rm`，在重命名文件时使用 `git` 做正确的事情。当您开始尝试通过 IDE 提交、合并、变基或进行复杂的历史分析时，一切都会崩溃。

When it comes to wielding the full power of Git, it’s command-line all the way.

当谈到使用 Git 的全部功能时，它一直是命令行。

Remember that whether you use Bash or Zsh or Powershell, there are [tab](https://git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-Bash) [completion](https://git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-Zsh) [scripts](https://git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-PowerShell) that take much of the pain out of remembering the subcommands and switches.

请记住，无论您使用 Bash 还是 Zsh 或 Powershell，都有 [tab](https://git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-Bash) [完成](https://git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-Zsh) [脚本](https:///git-scm.com/book/en/v2/Appendix-A%3A-Git-in-Other-Environments-Git-in-PowerShell)，它可以减轻记住子命令和开关的痛苦。

### Read Pro Git

### 阅读 Pro Git

The [Pro Git](https://git-scm.com/book/en/v2) book is available online for free, and it’s fantastic. Take advantage! 

[Pro Git](https://git-scm.com/book/en/v2) 书可在线免费获得，非常棒。好好利用！

