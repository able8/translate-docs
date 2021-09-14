# My Choice of Programming Languages

July 28, 2021

[Ranting,](http://iximiuz.com/en/categories/?category=Ranting) [Programming](http://iximiuz.com/en/categories/?category=Programming)

_It's a no holy war post! Any choice of languages cannot be 100% objective, and I'm just sharing my experience here._

When I was a kid, I used to spend days tinkering with woodworking tools. I was lucky enough to have a wide set of tools at my disposal. However, there was no one around to give me a hint about what tool to use when. So, I quickly came up with a heuristic: if my fingers and a tool survived an exercise, I've used the right tool; if either the fingers or the tool got damaged, I'd try other tools for the same task until I find the right one. And it worked quite well for me! Since then, I'm an apologist of the idea that every tool is good only for a certain set of tasks.

A programming language is yet another kind of tool. When I became a software developer, I adapted my heuristic to the new reality: if, while solving a task using a certain language, I suffer too much (fingers damage) or I need to hack things more often than not (tool damage), it's a wrong choice of a language.

Since the language is just a tool, my programming toolbox is defined by the tasks I work on the most often. Since 2010, I've worked in many domains, starting from web UI development and ending with writing code for infrastructure components. I find pleasure in being a generalist (jack of all trades), but there is always a pitfall of spreading yourself too thin (master of none). So, for the past few years, I've been trying to limit my sphere of expertise with the **server-side, distributed systems, and infrastructure**. Hence, the following choice of languages.

![Language logos](http://iximiuz.com/my-choice-of-programming-languages/kdpv-2000-opt.png)

## Python

The most appealing thing for me in [Python](http://iximiuz.com/en/categories/?category=Python) is the tremendous pace of development this language can provide. It's easy to write code in Python. There is often one clear way to accomplish a task, and you don't have to think much about how to convince the language to do what you need. Instead, you can focus on the business logic. I use Python for **quick prototyping**, for the performance-tolerant **server-side code**, for **ad hoc scripting**, and of course, for **data analysis**. I started my Python journey in 2014 and, without any prior experience with the language, managed to single-handedly develop a service that quickly got thousands of active users a day and [survived the load without a lot of rewriting](http://iximiuz.com/en/posts/save-the-day-with-gevent/).

## Go

If I had to describe [Go](http://iximiuz.com/en/categories/?category=Go) using just one word, I'd be torn apart between _pragmatic_ and _boring_. But I guess those are synonyms. Go is the language that could have replaced all other languages in my toolbox. Well, maybe except JavaScript. The development in Go is not as fast as in Python, but you still can have decent productivity. And the resulting code is almost as fast as code in C/C++ or Java. I use Go for **performance-sensitive services**, **infrastructure components**, and **command-line tools**. Static linking makes it especially good for tool development because the distribution is simplified a lot. Another reason to pick Go for me is the availability of some packages. Most of the [Cloud Native](http://iximiuz.com/en/posts/making-sense-out-of-cloud-native-buzz/) projects is written in Go, so it often becomes a [default choice](https://github.com/iximiuz/goimagego) for [certain kind of projects](https://github.com/iximiuz/conman). I have Go in my toolbox since 2016.

## Rust

Unlike Python, you may need to fight the compiler to get things done with [Rust](http://iximiuz.com/en/categories/?category=Rust). But if you manage to compile your code, the result is guaranteed to be safe (well, unless you abuse `unsafe` blocks). And most of the time, the result is also quite performant. I personally find this dancing with compiler quite rewarding, but I have to admit that the development pace degrades a lot. That's why I use Rust only when it's really necessary - for **performance-critical tools** or **[systems](https://github.com/iximiuz/shimmy) [programming](https://github.com/iximiuz/reapme)**. I don't consider myself a knowledgeable Rust developer just yet. The language is in my toolbox only since 2020. But I hope with [pq](https://github.com/iximiuz/pq) I'll be spending more time writing Rust this year.

# C

I [don't write code](https://github.com/iximiuz/golife.c) in [C](http://iximiuz.com/en/categories/?category=C). But I find the ability to read C code fluently a must-have in my domain. The Linux Kernel is written in C, and it makes it a native language for the whole family of operating systems. I always strive to **understand how things work one- or two layers of abstraction below my code**. Often it means I need to dig down to _libc_ calls and try to [reproduce snippets from man pages](https://github.com/iximiuz/ptyme) using bare C. Sometimes [reading the libc code itself](https://github.com/iximiuz/popen2) also helps. When I really have to wear a hat of a systems programmer, I [turn to Rust instead](http://iximiuz.com/en/posts/dealing-with-processes-termination-in-Linux/). My acquaintance with C lasts from 2006-2007.

## JavaScript

To some extent, I use JavaScript since 2011. Most of the time, it's about **making UIs** (web, desktop, mobile). For a few years, I've been trying to avoid such kinds of tasks and focus more on the server-side and systems programming. But I find it extremely useful to have such a tool in the toolbox. For instance, when I needed to [**visualize some data**](http://iximiuz.com/en/posts/pq/) recently, I quickly hacked a simple [HTML page](https://github.com/iximiuz/pq/blob/be4751177f014dd60e3144c3170b17f8e8c5e0fc/graph.html) to plot my datasets with [chart.js](https://www.chartjs.org/). I'm not even trying to keep up with all the modern web and UI frameworks, although I've written a few applications in React 🙈

## Other languages

I used to do some freelancing in **C++** and **Delphi**, but it ended a long, long time ago. I still use C++ sporadically to solve [Leetcode and HackerRank problems](https://twitter.com/iximiuz/status/1404545363730743303) but only when my Python solution times out. Rewriting it as-is in C++ rarely helps, though. Leaving the university-time freelancing aside, I started my professional career as a **PHP** developer. But I quickly found the language quite limiting. It's well-tailored for web development, but I don't see it as universal as Python or Go. Hence, no fit for my toolbox. I also spent a few years writing services in **Node.js** with **JavaScript** and **TypeScript**. Node.js experience definitely [broadened my understanding of asynchronous programming](http://iximiuz.com/en/posts/explain-event-loop-in-100-lines-of-code/). And I find TypeScript just lovely. But I try to limit the use of these tools only for UI making. In the past two years, I also had a chance to write some **Perl** and **Java** for high-loaded production services. Perl is a mind-blowing language with a [tremendous impact](https://twitter.com/iximiuz/status/1161980017796161538) on many of the mainstream programming languages. And the modern Java is also quite decent. But again, there is just very little fit for them, given my daily tasks.

## Looking for a mentor?

If you are at the beginning of your programming career or having a hard time understanding some basic concepts in one of my areas of expertise, I'd be happy to help. Just [drop me a message](http://iximiuz.com/en/about/), and we'll see how to go about it.

### Read also

[My 10 Years of Programming Experience](http://iximiuz.com/en/posts/my-10-years-of-programming-experience/)

[programming,](javascript: void 0) [python,](javascript: void 0) [golang,](javascript: void 0) [rust,](javascript: void 0) [c,](javascript: void 0) [javascript](javascript: void 0)

#### Written by Ivan Velichko

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom) [RSS](http://iximiuz.com/feed.rss)

