# Why Software Development is Hard

# 为什么软件开发很难

Posted on January  1, 2021

There’s this idea that having better programming languages will make  software development much easier and more productive. That no doubt used to be true, back when assembly or Fortran came along. However,  languages are now good enough that the main difficulties – and thus opportunities for improvement – are found elsewhere. Programming is  still hard, but for reasons that have nothing to do with the language  used.

有一种想法是，拥有更好的编程语言将使软件开发更容易、更高效。毫无疑问，在汇编或 Fortran 出现时，这曾经是正确的。然而，语言现在已经足够好了，主要的困难——以及改进的机会——都可以在别处找到。编程仍然很难，但原因与所使用的语言无关。

## Amdahl’s law

## 阿姆达尔定律

When you have a sequential series of tasks, [Amdahl’s law](https://en.wikipedia.org/wiki/Amdahl's_law) applies. It tells us that there is a hard limit to how much you can  speed up the entire series of tasks by speeding up just one of those  tasks.

当您有一系列连续的任务时，[阿姆达尔定律](https://en.wikipedia.org/wiki/Amdahl's_law) 适用。它告诉我们，通过加速其中一个任务，您可以在多大程度上加速整个系列任务。

Say boiling water takes 10 minutes and then cooking the pasta takes  another 10 minutes. If you work on finding a way to boil water faster,  you’ll never make dinner take less time than the 10 minutes the pasta  needs. An infinitely powerful burner will never give more than a 2x  speedup.

假设沸水需要 10 分钟，然后煮意大利面需要另外 10 分钟。如果您努力寻找一种更快地煮沸水的方法，那么晚餐所需的时间永远不会少于意大利面所需的 10 分钟。一个无限强大的燃烧器永远不会提供超过 2 倍的加速。

The general formula is that if something takes *p* portion of the total time, you can never get a speedup greater than 1/(1 – *p*). If a portion of the job takes 90% of the time, then *p* = 0.90. Optimizing that part down to zero time would speed up the overall job by 1/(1 – 0.90) = 10x.

一般公式是，如果某件事占用了总时间的 *p* 部分，则永远不会获得大于 1/(1 – *p*) 的加速比。如果工作的一部分占用了 90% 的时间，则 *p* = 0.90。将该部分优化到零时间将使整体工作速度提高 1/(1 – 0.90) = 10 倍。

The key to Amdahl’s law is that the best possible speedup you can get is limited by the size of the part you are optimizing.

Amdahl 定律的关键在于，您可以获得的最佳加速速度受到您正在优化的零件尺寸的限制。

Programming is hard for lots of reasons. As a simplification, we can  think of the things that make it difficult as tasks that must be done  sequentially. After all, humans are not very good at multitasking. At  any given point in time, you are using a build tool, reading  documentation, writing code, or sitting in a meeting. Or maybe writing  code instead of paying attention to the meeting. You deal with one  challenge at once, so Amdahl’s law roughly applies[1](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn1). If you manage to get build times down to zero, your projects will only  get done a little bit faster. Your productivity is still limited by all  the other stuff that goes into getting your project done.

编程之所以困难，原因有很多。作为一种简化，我们可以将困难的事情视为必须按顺序完成的任务。毕竟，人类并不擅长多任务处理。在任何给定的时间点，您都在使用构建工具、阅读文档、编写代码或参加会议。或者也许写代码而不是关注会议。您一次应对一个挑战，因此 Amdahl 定律大致适用[1](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn1)。如果您设法将构建时间减少到零，您的项目只会更快地完成。您的工作效率仍然受到完成项目所需的所有其他因素的限制。

Translating a plan for a program into something a computer can run  used to be incredibly difficult. Long ago, this involved translating the program all the way down to ones and zeros and tediously inputting  them. I do not know how much time this took, but for the sake of  argument (and easy math), let’s say it was 90% of the work of  programming. That would mean that a better way of telling a computer  what to do (such as Python) could give as much as a 10x improvement in  programming productivity.

将程序的计划转换成计算机可以运行的东西过去非常困难。很久以前，这涉及将程序一直翻译成 1 和 0 并繁琐地输入它们。我不知道这花了多少时间，但为了论证（和简单的数学），假设它占了编程工作的 90%。这意味着告诉计算机该做什么的更好方法（例如 Python）可以将编程效率提高 10 倍之多。

But now we have better languages. It takes less time to tell a  computer what to do – we got the promised productivity improvement. Translating a plan for a program into code no longer takes 90% of the  time. It now takes (again, for the sake of argument) only 10% of the  time. This would mean the maximum improvement you could get from making  this part easier is now only about 1.11x. That’s 81 times smaller than  the speedup that used to be available[2](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn2)!

但是现在我们有了更好的语言。告诉计算机该做什么花费的时间更少——我们得到了承诺的生产力改进。将程序的计划翻译成代码不再需要 90% 的时间。现在（再次，为了论证）它只需要 10% 的时间。这意味着您可以通过使这部分更容易获得的最大改进现在只有大约 1.11 倍。这比以前可用的加速小 81 倍 [2](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn2)！

This is because the other 90% of software development is a lot of  hard tasks that a better programming language won’t (directly) make  easier.

这是因为另外 90% 的软件开发是许多艰巨的任务，更好的编程语言不会（直接）使它们变得更容易。

## Why programming is still hard

## 为什么编程仍然很难

### How to lose friends

### 如何失去朋友

I’m saying that programming is hard in ways that have nothing to do  with the programming language. To see why, let’s start by pretending we  don’t have to worry about computers at all. Instead of teaching a  computer what needs to be done, you’ll tell your friend what to do. You  can’t cheat and tell your friend to rely on common sense – you have to  make all the decisions for them. 

我是说编程很难，与编程语言无关。要了解原因，让我们先假装我们根本不必担心计算机。与其教计算机需要做什么，不如告诉朋友该做什么。你不能欺骗并告诉你的朋友依靠常识——你必须为他们做所有的决定。

You’ll find that you need a long time just to explain the key  background information. Your friend will need to know about the things  in the real world that the program works with (“OK, so we have these  things called bundles, where you buy all the products in the bundle and  get a discount”) and what you've decided the program should do (“If the  user returns just one of the items in a bundle…”). Acronyms and terms  will have to be explained, external factors will have to be discussed  (“it's illegal for a manufacturer to dictate the price we sell the  product for, but not for them to dictate the price we advertise the  product for, so …” ).

你会发现仅仅解释关键背景信息就需要很长时间。您的朋友将需要了解该程序在现实世界中使用的东西（“好吧，我们有这些东西称为捆绑包，您可以在其中购买捆绑包中的所有产品并获得折扣”）以及您拥有的东西决定程序应该做什么（“如果用户只返回捆绑包中的一个项目......”）。必须解释首字母缩略词和术语，必须讨论外部因素（“制造商规定我们销售产品的价格是非法的，但他们规定我们宣传产品的价格是非法的，所以……” ）。

This friend needs to know about all the different situations that  might come up. There’s an incredible number of little details that must  be handled (“the user cannot enter a negative number of a product in the cart”). When you start to look at all the pairwise combinations of  states of different features, all the possible actions a user could try  to take, and all the possible events (a package gets lost in shipping),  you find that there's a huge number of edge cases you need to teach your friend about.

这位朋友需要了解可能出现的所有不同情况。必须处理数量惊人的小细节（“用户不能在购物车中输入产品的负数”）。当您开始查看不同特征状态的所有成对组合、用户可能尝试采取的所有可能操作以及所有可能的事件（包裹在运输过程中丢失）时，您会发现存在大量边缘你需要教给你朋友的案例。

Explaining all this to your friend is difficult in a couple of  different ways. First, you have to know about all the real–world details that are relevant to the program (there are products, they can be out  of stock, they can have discounts, etc). Second, you must know all the  decisions that have been made about what the program should do in each  possible situation. Third, you have to communicate all of this in a way  that your friend will understand. This means you need to organize your  thoughts well enough to make them digestible. If you write essays or  blog posts, you know that communicating a lot of information is not an  easy task!

以几种不同的方式向您的朋友解释这一切很困难。首先，您必须了解与该计划相关的所有真实世界的细节（有产品，它们可能缺货，它们可能有折扣等）。其次，您必须了解有关程序在每种可能情况下应该做什么的所有决定。第三，你必须以你的朋友能够理解的方式传达所有这些。这意味着你需要很好地组织你的想法，使它们易于消化。如果您撰写论文或博客文章，您就会知道传达大量信息并非易事！

Note that thus far, none of the work has involved a computer in any  way. It certainly hasn’t involved a programming language. Understanding  how the world works, knowing what the program should do, and organizing  the expression of those ideas are all tasks that are just plain hard.

请注意，到目前为止，没有一项工作以任何方式涉及计算机。它当然不涉及编程语言。了解世界是如何运作的，知道程序应该做什么，以及组织这些想法的表达，都是非常困难的任务。

### Description vs Specification

### 描述与规格

There’s a mental trap here that’s easy to fall into. It’s easy to  miss the distinction between describing something and specifying it. When you have a description (“a red car”), you can test whether or not a thing meets that description (“yes it's red, but it isn't a car”), but  it isn't enough to tell you how to create a car. That’s what a  specification is for.

这里有一个容易陷入的心理陷阱。很容易忽略描述某事和指定某事之间的区别。当你有一个描述（“一辆红色的车”）时，你可以测试一个东西是否符合那个描述（“是的，它是红色的，但它不是一辆汽车”），但这还不足以告诉你如何创造一辆汽车。这就是规范的用途。

Creating something requires making a lot of decisions. If you were to write down the results of every decision made, you would have a  (disorganized) specification. Writing a program requires that you’ve  made those decisions, so a mere description won’t do – you need a  specification. When you have a description (“it should list files”),  it’s easy to think that it’s a specification and thus it should be easy  to tell a computer to do it. But you’ve left out the million tiny  decisions that need to be made (“what order should it list the files in? Should they each go on their own line?”).

创造一些东西需要做出很多决定。如果你要写下每个决定的结果，你就会有一个（杂乱无章的）规范。编写程序需要你做出这些决定，所以仅仅描述是行不通的——你需要一个规范。当你有一个描述（“它应该列出文件”）时，很容易认为它是一个规范，因此告诉计算机这样做应该很容易。但是你忽略了需要做出的数百万个小决定（“它应该以什么顺序列出文件？它们应该各自独立吗？”）。

When you go to write a program, you are forced to confront any places where your specification is actually only a description. The computer  won’t accept just “draw a rectangle” – it will want to know where on the screen it should appear, how big it should be, and what color to make  it. Translating an idea into code tends to reveal the decisions you  haven’t made yet. Making these decisions is a big effort. It’s easy to  mistake the source of that effort, blaming it on the programming  language instead of acknowledging the simple fact that it’s hard to  create a specification given only a description.

当您开始编写程序时，您不得不面对任何您的规范实际上只是描述的地方。计算机不会只接受“画一个矩形”——它会想知道它应该出现在屏幕上的哪个位置，它应该有多大，以及它的颜色。将想法转化为代码往往会揭示您尚未做出的决定。做出这些决定是一项巨大的努力。很容易弄错这种努力的来源，将其归咎于编程语言，而不是承认一个简单的事实，即仅给出描述就很难创建规范。

### Back to Computers 

### 回到电脑

There’s more to developing software than understanding what it should do and translating the ideas into code. Computer themselves bring their own problems that the program must solve. Your program has to run fast  enough on real hardware and networks. The program might need to handle  machines failing. The complexities of tools and protocols add even more  problems to the domain. These aren’t difficulties caused by the process  of explaining to the computer what to do – they are just more things  that need to be explained.

开发软件不仅仅是了解它应该做什么并将想法转化为代码。计算机本身带来了程序必须解决的问题。你的程序必须在真实的硬件和网络上运行得足够快。该程序可能需要处理机器故障。工具和协议的复杂性给该领域增加了更多问题。这些并不是向计算机解释要做什么的过程造成的困难——它们只是更多需要解释的事情。

You also have to run parts of a program in your head. Sometimes the  flow of logic is as easy to follow as a story, but other times the  sequence of events and the state to keep track of can completely  overwhelm what you are able to fit into your head. Getting the details  of a program right – or fixing them when they are not right – requires  understanding the state of the program itself in various situations.

您还必须在头脑中运行程序的一部分。有时，逻辑流程就像故事一样容易理解，但有时，事件的顺序和要跟踪的状态可能会完全压倒您的头脑。正确获取程序的细节——或者在它们不正确时修复它们——需要了解程序本身在各种情况下的状态。

Writing code crystallizes your current idea of how the program works. But that doesn’t mean the program will stop changing. You’ll find bugs, want new features, and need to change existing behavior. The way the  program was organized may have worked at first, but that doesn’t mean  it’s a going to be the right structure forever. You’ll end up spending  time doing some combination of trying to predict the future and cleaning up the mess left when you inevitably discover that you weren’t  clairvoyant.

编写代码将您当前对程序如何工作的想法具体化。但这并不意味着该程序将停止更改。你会发现错误，想要新功能，并需要改变现有的行为。该计划的组织方式一开始可能有效，但这并不意味着它将永远是正确的结构。当你不可避免地发现自己不是千里眼时，你最终会花时间尝试预测未来和清理留下的烂摊子。

### Two is a crowd

### 两个是人群

If you aren’t just writing the program by yourself, you have to work  with other people. This brings an entirely new set of challenges.

如果您不只是自己编写程序，则必须与其他人一起工作。这带来了一系列全新的挑战。

All the people working on the project have to be organized in some  way, each person with their own work to do. You don’t want people to get in each other’s way, so you have to divide up the work. Creating a  sensible division of work requires understanding how the program is  structured. [Conway’s law](https://en.wikipedia.org/wiki/Conway's_law) applies.

所有参与项目的人都必须以某种方式组织起来，每个人都有自己的工作要做。你不希望人们互相妨碍，所以你必须分工。创建合理的工作分工需要了解程序的结构。 [康威定律](https://en.wikipedia.org/wiki/Conway's_law) 适用。

If you have multiple teams, things get even harder. Each team has  different goals, and will thus be optimizing for different things. A  decision that’s good for the other team might block you from getting  your work done. Understanding the position of the people across the  table from you and finding a good compromise is hard work, but needs to  be done.

如果您有多个团队，事情会变得更加困难。每个团队都有不同的目标，因此会针对不同的事情进行优化。对其他团队有利的决定可能会阻止您完成工作。了解你对面的人的立场并找到一个好的折衷方案是一项艰巨的工作，但需要完成。

In big projects, there’s no one team – let alone one person – who  understands the whole thing. Yet you still have to figure out how to  design parts of that system and make those parts fit together. This is a lot harder than if you were just creating the entire design yourself.

在大型项目中，没有一个团队——更不用说一个人——了解整件事。然而，您仍然必须弄清楚如何设计该系统的各个部分并将这些部分组合在一起。这比您自己创建整个设计要困难得多。

Even though it’s several steps removed from actually writing code,  dealing with the people side is a very real part of developing software.

尽管与实际编写代码相距几个步骤，但与人打交道是开发软件的一个非常真实的部分。

## Is there hope?

## 还有希望吗？

We can look for ways in which Amdahl’s law might not apply. If the  speeds of the separate tasks aren’t totally independent – if you can  speed up one task by optimizing another – then there’s hope that a  technical solution might help. 

我们可以寻找可能不适用阿姆达尔定律的方式。如果单独任务的速度不是完全独立的——如果你可以通过优化另一个任务来加速另一个任务——那么技术解决方案可能会有所帮助。

A vastly better language and development environment might be the  loophole we need. If it allowed a program to be written by fewer people – say two people instead of a team, or one team instead of a department – you could seriously cut down on the organizational overhead. You  wouldn’t need to have a meeting to decide on an interface if you are  personally writing all the code on both sides of that interface. The  increase in productivity wouldn’t just lower the cost of writing the  code, it would change the shape of the work in a way that lowers the  cost of the other tasks as well. That said, there’s a limit to how far  down this path you can go because one programmer cannot fit everything a business does into their head.

一个更好的语言和开发环境可能是我们需要的漏洞。如果它允许由更少的人编写程序——比如两个人而不是一个团队，或者一个团队而不是一个部门——你可以大大减少组织开销。如果您亲自编写接口两侧的所有代码，您就不需要开会来决定接口。生产力的提高不仅会降低编写代码的成本，还会以一种降低其他任务成本的方式改变工作的形式。也就是说，在这条道路上你能走多远是有限制的，因为一个程序员不可能把企业所做的一切都放在他们的脑海里。

Iteration speed is another lever. In order to write a program, you  need to understand the domain and decisions to be made. To create this  understanding, you might gather all the details into your head, and then arrange it into a mental model. That’s one way to do it, but maybe not  the most efficient way. Another approach is to build a small mental  model based on the obvious details. Then, create a small program from  that model to test the ideas against reality. Iterate based on the  feedback this provides, creating richer and more accurate models each  time around. This seems to work better for how people actually learn. For this method to be effective, you need to be able to test the ideas  and get the feedback quickly. The ideal state is that the new code  starts running as soon as you are done typing. Changing the development  environment to allow faster iteration cycles would let developers shift  from the first to the second approach for building their understanding  of the problem.

迭代速度是另一个杠杆。为了编写程序，您需要了解要做出的领域和决策。为了建立这种理解，您可以将所有细节收集到您的脑海中，然后将其整理成一个心智模型。这是一种方法，但可能不是最有效的方法。另一种方法是根据明显的细节建立一个小的心理模型。然后，根据该模型创建一个小程序，以根据实际情况测试这些想法。根据所提供的反馈进行迭代，每次都创建更丰富、更准确的模型。这似乎更适合人们的实际学习方式。为了使这种方法有效，您需要能够测试想法并快速获得反馈。理想的状态是新代码在您完成输入后立即开始运行。改变开发环境以允许更快的迭代周期将使开发人员从第一种方法转变为第二种方法来建立他们对问题的理解。

I’m not especially optimistic that a more expressive language will  meaningfully increase productivity. I do still hold some hope that a  significantly better development environment is possible. If we had  better tools to understand existing code, faster development iteration  cycles, and less tedious “toil” work, it might meaningfully change how  software development is done in a way that has compounding – instead of  diminishing – returns.

对于更具表现力的语言会有意义地提高生产力，我并不特别乐观。我仍然抱有一些希望，一个明显更好的开发环境是可能的。如果我们有更好的工具来理解现有代码、更快的开发迭代周期和更少乏味的“辛劳”工作，它可能会有意义地改变软件开发的方式，使回报复合——而不是递减——。

------

1. This relies on the assumption that a speedup in one area has no impact on how the other areas work. This probably isn’t actually the case. If build times go from one hour to one minute, you will  probably take advantage of this by creating more builds. If  communication has less latency (e.g. Slack vs email), you don’t just  change how much time you spend communicating. You also change how you  communicate.[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref1)
2. A 10x speedup is a 900% improvement. A 1.111…x speedup is an 11.11…% improvement. 900/11.111 = 81.[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref2) 

1. 这依赖于这样一个假设：一个区域的加速对其他区域的工作方式没有影响。实际情况可能并非如此。如果构建时间从一小时缩短到一分钟，您可能会通过创建更多构建来利用这一点。如果沟通的延迟更短（例如 Slack 与电子邮件），您不仅会改变沟通所花费的时间。你也会改变你的沟通方式。[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref1)
2. 10 倍加速是 900% 的改进。 1.111…x 的加速是 11.11…% 的改进。 900/11.111 = 81.[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref2)

