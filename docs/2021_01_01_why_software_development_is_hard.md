# Why Software Development is Hard

​    Posted on January  1, 2021    

There’s this idea that having better programming languages will make  software development much easier and more productive. That no doubt used to be true, back when assembly or Fortran came along. However,  languages are now good enough that the main difficulties – and thus  opportunities for improvement – are found elsewhere. Programming is  still hard, but for reasons that have nothing to do with the language  used.

## Amdahl’s law

When you have a sequential series of tasks, [Amdahl’s law](https://en.wikipedia.org/wiki/Amdahl's_law) applies. It tells us that there is a hard limit to how much you can  speed up the entire series of tasks by speeding up just one of those  tasks.

Say boiling water takes 10 minutes and then cooking the pasta takes  another 10 minutes. If you work on finding a way to boil water faster,  you’ll never make dinner take less time than the 10 minutes the pasta  needs. An infinitely powerful burner will never give more than a 2x  speedup.

The general formula is that if something takes *p* portion of the total time, you can never get a speedup greater than 1/(1 – *p*). If a portion of the job takes 90% of the time, then *p* = 0.90. Optimizing that part down to zero time would speed up the overall job by 1/(1 – 0.90) = 10x.

The key to Amdahl’s law is that the best possible speedup you can get is limited by the size of the part you are optimizing.

Programming is hard for lots of reasons. As a simplification, we can  think of the things that make it difficult as tasks that must be done  sequentially. After all, humans are not very good at multitasking. At  any given point in time, you are using a build tool, reading  documentation, writing code, or sitting in a meeting. Or maybe writing  code instead of paying attention to the meeting. You deal with one  challenge at once, so Amdahl’s law roughly applies[1](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn1). If you manage to get build times down to zero, your projects will only  get done a little bit faster. Your productivity is still limited by all  the other stuff that goes into getting your project done.

Translating a plan for a program into something a computer can run  used to be incredibly difficult. Long ago, this involved translating the program all the way down to ones and zeros and tediously inputting  them. I do not know how much time this took, but for the sake of  argument (and easy math), let’s say it was 90% of the work of  programming. That would mean that a better way of telling a computer  what to do (such as Python) could give as much as a 10x improvement in  programming productivity.

But now we have better languages. It takes less time to tell a  computer what to do – we got the promised productivity improvement.  Translating a plan for a program into code no longer takes 90% of the  time. It now takes (again, for the sake of argument) only 10% of the  time. This would mean the maximum improvement you could get from making  this part easier is now only about 1.11x. That’s 81 times smaller than  the speedup that used to be available[2](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fn2)!

This is because the other 90% of software development is a lot of  hard tasks that a better programming language won’t (directly) make  easier.

## Why programming is still hard

### How to lose friends

I’m saying that programming is hard in ways that have nothing to do  with the programming language. To see why, let’s start by pretending we  don’t have to worry about computers at all. Instead of teaching a  computer what needs to be done, you’ll tell your friend what to do. You  can’t cheat and tell your friend to rely on common sense – you have to  make all the decisions for them.

You’ll find that you need a long time just to explain the key  background information. Your friend will need to know about the things  in the real world that the program works with (“OK, so we have these  things called bundles, where you buy all the products in the bundle and  get a discount”) and what you’ve decided the program should do (“If the  user returns just one of the items in a bundle…”). Acronyms and terms  will have to be explained, external factors will have to be discussed  (“it’s illegal for a manufacturer to dictate the price we sell the  product for, but not for them to dictate the price we advertise the  product for, so …”).

This friend needs to know about all the different situations that  might come up. There’s an incredible number of little details that must  be handled (“the user cannot enter a negative number of a product in the cart”). When you start to look at all the pairwise combinations of  states of different features, all the possible actions a user could try  to take, and all the possible events (a package gets lost in shipping),  you find that there’s a huge number of edge cases you need to teach your friend about.

Explaining all this to your friend is difficult in a couple of  different ways. First, you have to know about all the real–world details that are relevant to the program (there are products, they can be out  of stock, they can have discounts, etc). Second, you must know all the  decisions that have been made about what the program should do in each  possible situation. Third, you have to communicate all of this in a way  that your friend will understand. This means you need to organize your  thoughts well enough to make them digestible. If you write essays or  blog posts, you know that communicating a lot of information is not an  easy task!

Note that thus far, none of the work has involved a computer in any  way. It certainly hasn’t involved a programming language. Understanding  how the world works, knowing what the program should do, and organizing  the expression of those ideas are all tasks that are just plain hard.

### Description vs Specification

There’s a mental trap here that’s easy to fall into. It’s easy to  miss the distinction between describing something and specifying it.  When you have a description (“a red car”), you can test whether or not a thing meets that description (“yes it’s red, but it isn’t a car”), but  it isn’t enough to tell you how to create a car. That’s what a  specification is for.

Creating something requires making a lot of decisions. If you were to write down the results of every decision made, you would have a  (disorganized) specification. Writing a program requires that you’ve  made those decisions, so a mere description won’t do – you need a  specification. When you have a description (“it should list files”),  it’s easy to think that it’s a specification and thus it should be easy  to tell a computer to do it. But you’ve left out the million tiny  decisions that need to be made (“what order should it list the files in? Should they each go on their own line?”).

When you go to write a program, you are forced to confront any places where your specification is actually only a description. The computer  won’t accept just “draw a rectangle” – it will want to know where on the screen it should appear, how big it should be, and what color to make  it. Translating an idea into code tends to reveal the decisions you  haven’t made yet. Making these decisions is a big effort. It’s easy to  mistake the source of that effort, blaming it on the programming  language instead of acknowledging the simple fact that it’s hard to  create a specification given only a description.

### Back to Computers

There’s more to developing software than understanding what it should do and translating the ideas into code. Computer themselves bring their own problems that the program must solve. Your program has to run fast  enough on real hardware and networks. The program might need to handle  machines failing. The complexities of tools and protocols add even more  problems to the domain. These aren’t difficulties caused by the process  of explaining to the computer what to do – they are just more things  that need to be explained.

You also have to run parts of a program in your head. Sometimes the  flow of logic is as easy to follow as a story, but other times the  sequence of events and the state to keep track of can completely  overwhelm what you are able to fit into your head. Getting the details  of a program right – or fixing them when they are not right – requires  understanding the state of the program itself in various situations.

Writing code crystallizes your current idea of how the program works. But that doesn’t mean the program will stop changing. You’ll find bugs, want new features, and need to change existing behavior. The way the  program was organized may have worked at first, but that doesn’t mean  it’s a going to be the right structure forever. You’ll end up spending  time doing some combination of trying to predict the future and cleaning up the mess left when you inevitably discover that you weren’t  clairvoyant.

### Two is a crowd

If you aren’t just writing the program by yourself, you have to work  with other people. This brings an entirely new set of challenges.

All the people working on the project have to be organized in some  way, each person with their own work to do. You don’t want people to get in each other’s way, so you have to divide up the work. Creating a  sensible division of work requires understanding how the program is  structured. [Conway’s law](https://en.wikipedia.org/wiki/Conway's_law) applies.

If you have multiple teams, things get even harder. Each team has  different goals, and will thus be optimizing for different things. A  decision that’s good for the other team might block you from getting  your work done. Understanding the position of the people across the  table from you and finding a good compromise is hard work, but needs to  be done.

In big projects, there’s no one team – let alone one person – who  understands the whole thing. Yet you still have to figure out how to  design parts of that system and make those parts fit together. This is a lot harder than if you were just creating the entire design yourself.

Even though it’s several steps removed from actually writing code,  dealing with the people side is a very real part of developing software.

## Is there hope?

We can look for ways in which Amdahl’s law might not apply. If the  speeds of the separate tasks aren’t totally independent – if you can  speed up one task by optimizing another – then there’s hope that a  technical solution might help.

A vastly better language and development environment might be the  loophole we need. If it allowed a program to be written by fewer people – say two people instead of a team, or one team instead of a department – you could seriously cut down on the organizational overhead. You  wouldn’t need to have a meeting to decide on an interface if you are  personally writing all the code on both sides of that interface. The  increase in productivity wouldn’t just lower the cost of writing the  code, it would change the shape of the work in a way that lowers the  cost of the other tasks as well. That said, there’s a limit to how far  down this path you can go because one programmer cannot fit everything a business does into their head.

Iteration speed is another lever. In order to write a program, you  need to understand the domain and decisions to be made. To create this  understanding, you might gather all the details into your head, and then arrange it into a mental model. That’s one way to do it, but maybe not  the most efficient way. Another approach is to build a small mental  model based on the obvious details. Then, create a small program from  that model to test the ideas against reality. Iterate based on the  feedback this provides, creating richer and more accurate models each  time around. This seems to work better for how people actually learn.  For this method to be effective, you need to be able to test the ideas  and get the feedback quickly. The ideal state is that the new code  starts running as soon as you are done typing. Changing the development  environment to allow faster iteration cycles would let developers shift  from the first to the second approach for building their understanding  of the problem.

I’m not especially optimistic that a more expressive language will  meaningfully increase productivity. I do still hold some hope that a  significantly better development environment is possible. If we had  better tools to understand existing code, faster development iteration  cycles, and less tedious “toil” work, it might meaningfully change how  software development is done in a way that has compounding – instead of  diminishing – returns.

------

1. This relies on the assumption that a speedup in one area has no impact on how the other areas work. This probably isn’t actually the case. If build times go from one hour to one minute, you will  probably take advantage of this by creating more builds. If  communication has less latency (e.g. Slack vs email), you don’t just  change how much time you spend communicating. You also change how you  communicate.[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref1)
2. A 10x speedup is a 900% improvement. A 1.111…x speedup is an 11.11…% improvement. 900/11.111 = 81.[↩](http://jeremymikkola.com/posts/2021_01_01_why_software_development_is_hard.html#fnref2)
