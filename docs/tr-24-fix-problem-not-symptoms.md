# Fix the Problem, Not the Symptoms

# 解决问题，而不是症状

From: https://product.hubspot.com/blog/fix-problem-not-symptoms

Here’s something I think everyone can relate to: you fix a problem, only to  discover later on that it’s still happening. Why is that? Sometimes  we’re in a rush, or are a little careless. I believe it’s often because  we did not understand the problem well enough to actually address it. By being a little more deliberate we can avoid this, and reliably fix  problems once and for all. And it’s not a particularly technical skill. We need to start by asking “why” more, and looking for the answers to  understand the problem well enough to confidently solve it.

这是我认为每个人都可以理解的事情：你解决了一个问题，后来发现它仍然在发生。为什么？有时我们很匆忙，或者有点粗心。我相信这通常是因为我们对问题的理解不够深入，无法真正解决它。通过更加谨慎，我们可以避免这种情况，并可靠地一劳永逸地解决问题。而且这不是一项特别的技术技能。我们需要从多问“为什么”开始，寻找答案以充分理解问题，从而自信地解决它。

## The case of the OutOfMemoryErrors

## OutOfMemoryErrors 的情况

At HubSpot one of the things I’m  responsible for is a web crawler. This system has had a multitude of  issues but one class of them, OutOfMemoryErrors, have been perennial. At times it seemed there was another every other week. They no longer  plague me, and I attribute that to continually questioning why the error was occurring and searching for the answers before fixing it. An example from earlier this year started like most others: with a heap dump.

在 HubSpot，我负责的一项工作是网络爬虫。这个系统有很多问题，但其中一类，OutOfMemoryErrors，一直存在。有时似乎每隔一周就有一次。他们不再困扰我，我将其归因于不断质疑错误发生的原因并在修复之前寻找答案。今年早些时候的一个示例与大多数其他示例一样：使用堆转储。

-  **Why did it fill the heap?** Turns out while parsing the HTML a base 64 encoded PNG was taking up a considerable amount of memory.

- **为什么它会填满堆？** 结果在解析 HTML 时，base 64 编码的 PNG 占用了大量内存。

-  **Why did it have an image at all?** The service is  supposed to collect only the HTML. The code had set a flag to not  download images for the headless browser client so that naturally made  me question:

- **为什么它有一个图像？** 该服务应该只收集 HTML。代码设置了一个标志，不为无头浏览器客户端下载图像，所以这自然让我产生了疑问：

- **Why then did the image appear in the heap?** Looking  further into the client I could see that the flag was being passed down, but I was able to repeatedly capture images with that client, with the  flag set.

- **为什么图像会出现在堆中？** 进一步查看客户端，我可以看到标志正在传递，但我能够在设置标志的情况下使用该客户端重复捕获图像。

- **So why were the images present?** A quick search  revealed that the name of that flag was wrong. I was able to submit a  pull request to the client, with screenshot evidence of an example  webpage before and after my change to demonstrate that it was in fact  not properly omitting images before.

- **那么为什么会出现这些图像？** 快速搜索发现该标志的名称是错误的。我能够向客户端提交拉取请求，并附上更改前后示例网页的截图证据，以证明它实际上没有正确省略之前的图像。

Had I just added more memory to the  service, this issue would have been buried for longer. It would have  consumed more resources, which cost us. Other users of the headless  browser client would be inadvertently downloading images as well. Best  of all, I don’t get OutOfMemoryErrors any longer.

如果我刚刚为服务添加了更多内存，这个问题就会被埋葬更长时间。它会消耗更多的资源，这会让我们付出代价。无头浏览器客户端的其他用户也会无意中下载图像。最重要的是，我不再收到 OutOfMemoryErrors 了。

## A habit of asking “why”

## 问“为什么”的习惯

At each step of the way I was able to find some concrete evidence that led me to the next step. Often the code  itself is not enough to correctly discover an issue so these artifacts  are incredibly important. Things like heap dumps, thread dumps, and log  files can give us a picture of the system as it runs. They also  typically have timestamps included so it is possible to collate the  evidence to construct a more complete picture of what happened. By cultivating a habit of asking “why” and searching for evidence, we can solve problems rather than treating  symptoms. I say “habit” because it’s a practice that we add to our work  rather than a specific skill. You may find that it takes a while to  become successful at it, but like many things, the more repetition, the  easier it will become, until you find it’s happening automatically for  you. 

在每一步，我都能找到一些具体的证据，引导我进入下一步。通常，代码本身不足以正确发现问题，因此这些工件非常重要。堆转储、线程转储和日志文件之类的内容可以为我们提供系统运行时的图片。它们通常还包含时间戳，因此可以整理证据以构建更完整的事件图景。通过培养问“为什么”和寻找证据的习惯，我们可以解决问题而不是治标不治本。我说“习惯”是因为它是我们添加到工作中的一种实践，而不是一种特定的技能。你可能会发现要在这件事上取得成功需要一段时间，但就像许多事情一样，重复得越多，它就会变得越容易，直到你发现它自动发生在你身上。

