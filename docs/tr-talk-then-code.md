# Talk, then code

# 讨论，然后编码

[February 18, 2019](https://dave.cheney.net/2019/02/18/talk-then-code)

The open source projects that I contribute to follow a philosophy which I describe as _talk, then code_. I think this is generally a good way to develop software and I want to spend a little time talking about the benefits of this methodology.

我参与的开源项目遵循一种哲学，我将其描述为 _talk，然后是 code_。我认为这通常是开发软件的好方法，我想花一点时间谈谈这种方法的好处。

### Avoiding hurt feelings

### 避免受伤的感觉

The most important reason for discussing the change you want to make is it avoids hurt feelings. Often I see a contributor work hard in isolation on a pull request only to find their work is rejected. This can be for a bunch of reasons; the PR is too large, the PR doesn’t follow the local style, the PR fixes an issue which wasn’t important to the project or was recently fixed indirectly, and many more.

讨论您想要做出的改变的最重要原因是避免伤害感情。我经常看到一个贡献者在一个 pull request 上孤立地工作，结果发现他们的工作被拒绝了。这可能有很多原因； PR 太大，PR 不遵循本地风格，PR 修复了对项目不重要或最近间接修复的问题，等等。

The underlying cause of all these issues is a lack of communication. The goal of the _talk, then code_ philosophy is not to impede or frustrate, but to ensure that a feature lands correctly the first time, without incurring significant maintenance debt, and neither the author of the change, or the reviewer, has to carry the emotional burden of dealing with hurt feelings when a change appears out of the blue with an implicit “well, I've done the work, all you have to do is merge it, right?”

所有这些问题的根本原因是缺乏沟通。 _talk，然后是代码_哲学的目标不是阻碍或挫败，而是确保功能第一次正确登陆，而不会产生重大的维护债务，并且更改的作者或审阅者都不必承担当一个变化突然出现并隐含着“好吧，我已经完成了工作，你所要做的就是合并它，对吧？”时，处理伤害感受的情感负担。

### What does discussion look like?

### 讨论是什么样的？

Every new feature or bug fix should be discussed with the maintainer(s) of the project before work commences. It’s fine to experiment privately, but do not send a change without discussing it first.

在工作开始之前，每个新功能或错误修复都应该与项目的维护者进行讨论。私下试验是可以的，但不要在没有事先讨论的情况下发送更改。

The definition of _talk_ for simple changes can be as little as a design sketch in a GitHub issue. If your PR fixes a bug, you should link to the bug it fixes. If there isn’t one, you should raise a bug and wait for the maintainers to acknowledge it before sending a PR. This might seem a little backward–who wouldn’t want a bug fixed–but consider the bug could be a misunderstanding in how the software works or it could be a symptom of a larger problem that needs further investigation.

_talk_ 的简单更改定义可以像 GitHub 问题中的设计草图一样小。如果您的 PR 修复了错误，您应该链接到它修复的错误。如果没有，你应该提出一个错误并等待维护者在发送 PR 之前确认它。这似乎有点落后——谁不想修复错误——但考虑到错误可能是对软件工作方式的误解，或者它可能是需要进一步调查的更大问题的征兆。

For more complicated changes, especially feature requests, I recommend that a design document be circulated and agreed upon before sending code. This doesn’t have to be a full blown document, a sketch in an issue may be sufficient, but the key is to reach agreement using words, before locking it in stone with code.

对于更复杂的更改，尤其是功能请求，我建议在发送代码之前分发一份设计文档并达成一致。这不一定是一个完整的文档，一个问题的草图可能就足够了，但关键是在用代码将其锁定之前用文字达成一致。

In all cases you shouldn’t proceed to send code until there is a positive agreement from the maintainer that the approach is one they are happy with. A pull request is for life, not just for Christmas.

在所有情况下，您都不应该继续发送代码，直到维护者同意该方法是他们满意的方法。拉取请求是终身的，而不仅仅是圣诞节。

### Code review, not design by committee

### 代码审查，而不是由委员会设计

A code review is not the place for arguments about design. This is for two reasons. First, most code review tools are not suitable for long comment threads, GitHub’s PR interface is very bad at this, Gerrit is better, but few have a team of admins to maintain a Gerrit instance. More importantly, disagreements at the code review stage suggests there wasn’t agreement on how the change should be implemented.

代码审查不是讨论设计的地方。这是出于两个原因。首先，大多数代码审查工具不适合长评论线程，GitHub 的 PR 接口在这方面非常糟糕，Gerrit 更好，但很少有管理员团队来维护 Gerrit 实例。更重要的是，代码审查阶段的分歧表明在如何实施更改方面没有达成一致。

* * *

Talk about what you want to code, then code what you talked about. Please don’t do it the other way around.

谈论您想编码的内容，然后编码您所谈论的内容。请不要反过来做。

### Related posts:

###  相关文章：

1. [How to include C code in your Go package](https://dave.cheney.net/2013/09/07/how-to-include-c-code-in-your-go-package "How to include C code in your Go package")
2. [Let’s talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging "Let’s talk about logging")
3. [The value of TDD](https://dave.cheney.net/2016/04/11/the-value-of-tdd "The value of TDD") 

