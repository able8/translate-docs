# Debugging Software Deployments with strace

# 使用 strace 调试软件部署

Published 14 November 2019

2019 年 11 月 14 日发布

Tags: [Ops](http://theartofmachinery.com/tags/ops/ "Posts Tagged Ops") , [Tools](http://theartofmachinery.com/tags/tools/ "Posts Tagged Tools") , [Reliability ](http://theartofmachinery.com/tags/reliability/ "Posts Tagged Reliability") , [Servers](http://theartofmachinery.com/tags/servers/ "Posts Tagged Servers")

Most of my paid work involves deploying software systems, which means I spend a lot of time trying to answer the following questions:

- This software works on the original developer’s machine, so why doesn’t it work on mine?
- This software worked on my machine yesterday, so why doesn’t it work today?

我的大部分有偿工作都涉及部署软件系统，这意味着我花了很多时间尝试回答
以下问题：

- 这个软件可以在原开发者的机器上运行，为什么不能在我的机器上运行？
- 这个软件昨天在我的机器上运行了，为什么今天不运行了？

That’s a kind of debugging, but it’s a different kind of debugging from normal software debugging. Normal debugging is usually about the logic of the code, but deployment debugging is usually about the interaction between the code and its environment. Even when the root cause is a logic bug, the fact that the software apparently worked on another machine means that the environment is usually involved somehow.

这是一种调试，但与普通的软件调试不同。正常调试通常是关于代码的逻辑，但部署调试通常是关于代码和它的环境。即使根本原因是逻辑错误，但该软件显然可以在另一个machine 意味着通常以某种方式涉及环境。

So, instead of using normal debugging tools like `gdb`, I have another toolset for debugging deployments. My favourite tool for “Why isn’t this software working on this machine?” is `strace`.

因此，我没有使用像 `gdb` 这样的普通调试工具，而是有另一个用于调试部署的工具集。我最喜欢的工具“为什么这个软件不能用于这个机器？”是`strace`。

## What is `strace`?

## 什么是`strace`？

[`strace`](https://strace.io/) is a tool for “system call tracing”. It’s primarily a Linux tool, but you can do the same kind of debugging tricks with tools for other systems (such as [DTrace](http://dtrace.org/blogs/about/) and [ktrace](https://man.openbsd.org/ktrace)).

[`strace`](https://strace.io/) 是一个工具“系统调用跟踪”。它主要是一个 Linux 工具，但您可以使用以下工具执行相同类型的调试技巧其他系统（例如 [DTrace](http://dtrace.org/blogs/about/) 和 [ktrace](https://man.openbsd.org/ktrace))。

The basic usage is very simple. Just run it against a command and it dumps all the system calls (you’ll probably need to install `strace` first):

基本用法非常简单。只需针对命令运行它，它就会转储所有系统调用（您可能会
需要先安装`strace`）：

```shell
$ strace echo Hello
...Snip lots of stuff...
write(1, "Hello\n", 6)                  = 6
close(1)                                = 0
close(2)                                = 0
exit_group(0)                           = ?
+++ exited with 0 +++

```

What are these system calls? They’re like the API for the operating system kernel. Once upon a time, software used to have direct access to the hardware it ran on. If it needed to display something on the screen, for example, it could twiddle with ports and/or memory-mapped registers for the video hardware. That got chaotic when multitasking computer systems became popular because different applications would “fight” over hardware, and bugs in one application could crash other applications, or even bring down the whole system. So CPUs started supporting different privilege modes (or “protection rings”). They let an operating system kernel run in the most privileged mode with full hardware access, while spawning less-privileged software applications that must ask the kernel to interact with the hardware for them using system calls.

这些系统调用是什么？它们就像操作系统内核的 API。曾几何时，使用的软件直接访问它运行的硬件。例如，如果它需要在屏幕上显示一些东西，它可以摆弄视频硬件的端口和/或内存映射寄存器。多任务处理计算机时变得混乱系统变得流行是因为不同的应用程序会“争夺”硬件，并且一个应用程序中的错误可能使其他应用程序崩溃，甚至使整个系统崩溃。因此 CPU 开始支持不同的特权模式（或“保护环”）。他们让操作系统内核在具有完全硬件访问权限的最高特权模式下运行，同时产生低权限的软件应用程序，这些应用程序必须要求内核为它们与硬件交互使用系统调用。

At the binary level, making a system call is a bit different from making a simple function call, but most programs use wrappers in a standard library. E.g. the POSIX C standard library contains a `write()` function call that contains all the architecture-dependent code for making the `write` system call. 

在二进制层面，进行系统调用与进行简单的函数调用有点不同，但大多数程序在标准库中使用包装器。例如。 POSIX C 标准库包含一个 `write()` 函数调用，其中包含所有与体系结构相关的代码用于进行 `write` 系统调用。

[![How system calls sit between applications and the kernel](http://theartofmachinery.com/images/strace/system_calls.svg)](http://theartofmachinery.com/images/strace/system_calls.svg)

In short, an application’s interaction with its environment (the computer system) is all done through system calls. So when software works on one machine but not another, looking at system call traces is a good way to find what’s wrong. More specifically, here are the typical things you can analyse using a system call trace:

- Console input and output (IO)
- Network IO
- Filesystem access and file IO
- Process/thread lifetime management
- Raw memory management
- Access to special device drivers

简而言之，应用程序与其环境（计算机系统）的交互都是通过系统调用完成的。因此，当软件在一台机器上运行而不在另一台机器上运行时，查看系统调用跟踪是找到什么是错误的。更具体地说，以下是您可以使用系统调用跟踪分析的典型事物：

- 控制台输入和输出 (IO)
- 网络IO
- 文件系统访问和文件 IO
- 进程/线程生命周期管理
- 原始内存管理
- 访问特殊设备驱动程序

## When can `strace` be used?

## 什么时候可以使用`strace`？

In theory, `strace` can be used with any userspace program because all userspace programs have to make system calls. It’s more effective with compiled, lower-level programs, but still works with high-level languages like Python if you can wade through the extra noise from the runtime environment and interpreter.

理论上，`strace` 可以用于任何用户空间程序因为所有用户空间程序都必须进行系统调用。它对编译后的低级程序更有效，但如果您可以克服来自运行时环境的额外噪音，仍然可以使用 Python 等高级语言和口译员。

`strace` shines with debugging software that works fine on one machine, but on another machine fails with a vague error message about files or permissions or failure to run some command or something. Unfortunately, it’s not so great with higher-level problems, like a certificate verification
failure. They usually need a combination of `strace`, sometimes [`ltrace`](https://linux.die.net/man/1/ltrace), and higher-level tooling (like the `openssl` command line tool for certificate debugging).

`strace` 在大放异彩，在调试软件在一台机器，但在另一台机器上失败，并显示有关文件或权限的模糊错误消息或无法运行某些命令什么的。不幸的是，对于更高级别的问题，它并不是那么好，比如证书验证失败。他们通常需要结合`strace`，有时[`ltrace`](https://linux.die.net/man/1/ltrace)，以及更高级别的工具（如用于证书调试）。

The examples in this post are based on a standalone server, but system call tracing can often be done on more complicated deployment platforms, too. Just search for appropriate tooling.

本文中的示例基于独立服务器，但系统调用跟踪通常可以在更多服务器上完成复杂的部署平台。只需搜索合适的工具即可。

## A simple debugging example

## 一个简单的调试例子

Let’s say you’re trying to run an awesome server application called foo, but here’s what happens:

```shell
$ foo
Error opening configuration file: No such file or directory
```



Obviously it’s not finding the configuration file that you’ve written. This can happen because package managers sometimes customise the expected locations of files when compiling an application, so following an installation guide for one distro leads to files in the wrong place on another distro. You could fix the problem in a few seconds if only the error message told you where the configuration file is expected to be, but it doesn’t. How can you find out?

显然它没有找到您编写的配置文件。这可能是因为包管理器有时在编译应用程序时自定义文件的预期位置，因此请遵循安装指南因为一个发行版导致文件在另一个发行版上的错误位置。你可以在几秒钟内解决问题，如果只是错误消息告诉你配置文件应该在哪里，但它没有。你怎么知道？

If you have access to the source code, you could read it and work it out. That’s a good fallback plan, but not the fastest solution. You also could use a stepping debugger like `gdb` to see what the program does, but it’s more efficient to use a tool that’s specifically designed to show the interaction with the environment: `strace`.

如果您可以访问源代码，则可以阅读并解决它。这是一个很好的后备计划，但不是最快的解决方案。你也可以使用像`gdb`这样的步进调试器来查看程序做了什么，但使用工具更有效专门设计用于显示与环境的交互：`strace`。

The output of `strace` can be a bit overwhelming at first, but the good news is that you can ignore most of it. It often helps to use the `-o` switch to save the trace to a separate file:

`strace` 的输出一开始可能有点让人不知所措，
但好消息是您可以忽略其中的大部分内容。使用“-o”开关将跟踪保存到单独的文件通常会有所帮助：

```shell
$ strace -o /tmp/trace foo
Error opening configuration file: No such file or directory
$ cat /tmp/trace
execve("foo", ["foo"], 0x7ffce98dc010 /* 16 vars */) = 0
brk(NULL)                               = 0x56363b3fb000
access("/etc/ld.so.preload", R_OK)      = -1 ENOENT (No such file or directory)
openat(AT_FDCWD, "/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
fstat(3, {st_mode=S_IFREG|0644, st_size=25186, ...}) = 0
mmap(NULL, 25186, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f2f12cf1000
close(3)                                = 0
openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
read(3, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\260A\2 \0\0\0\0\0"..., 832) = 832
fstat(3, {st_mode=S_IFREG|0755, st_size=1824496, ...}) = 0
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f2f12cef000
mmap(NULL, 1837056, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f2f12b2e000
mprotect(0x7f2f12b50000, 1658880, PROT_NONE) = 0
mmap(0x7f2f12b50000, 1343488, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x22000) = 0x7f2f12b50000
mmap(0x7f2f12c98000, 311296, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x16a000) = 0x7f2f12c98000
mmap(0x7f2f12ce5000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1b6000) = 0x7f2f12ce5000
mmap(0x7f2f12ceb000, 14336, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f2f12ceb000
close(3)                                = 0
arch_prctl(ARCH_SET_FS, 0x7f2f12cf0500) = 0
mprotect(0x7f2f12ce5000, 16384, PROT_READ) = 0
mprotect(0x56363b08b000, 4096, PROT_READ) = 0
mprotect(0x7f2f12d1f000, 4096, PROT_READ) = 0
munmap(0x7f2f12cf1000, 25186)           = 0
openat(AT_FDCWD, "/etc/foo/config.json", O_RDONLY) = -1 ENOENT (No such file or directory)
dup(2)                                  = 3
fcntl(3, F_GETFL)                       = 0x2 (flags O_RDWR)
brk(NULL)                               = 0x56363b3fb000
brk(0x56363b41c000)                     = 0x56363b41c000
fstat(3, {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0x8), ...}) = 0
write(3, "Error opening configuration file"..., 60) = 60
close(3)                                = 0
exit_group(1)                           = ?
+++ exited with 1 +++
```

The first page or so of `strace` output is typically low-level process startup. (You can see a lot of `mmap`, `mprotect`, `brk` calls for things like allocating raw memory and mapping dynamic
libraries.) Actually, when debugging an error, `strace` output is best read from the bottom up. You can see the `write` call that outputs the error message at the end. If you work up, the first failing system call is the `openat` call that fails with `ENOENT` (“No such file or directory”) trying to open `/etc/foo/config.json`. And now we know where the configuration file is supposed to be.

`strace` 输出的第一页左右通常是低级进程启动。 （你可以看到很多`mmap`，`mprotect`、`brk` 调用诸如分配原始内存和动态映射之类的东西库。）实际上，在调试错误时，`strace` 输出最好从下往上读。你可以看到 `write` 调用。最后输出错误信息。如果你成功了，第一个失败的系统调用是 `openat` 调用，它以 `ENOENT`（“没有这样的文件或目录”）试图打开 `/etc/foo/config.json` 失败。现在我们知道配置文件在哪里
应该是。

That’s a simple example, but I’d say at least 90% of the time I use `strace`, I’m not doing anything more complicated. Here’s the complete debugging formula step-by-step:

1. Get frustrated by a vague system-y error message from a program
2. Run the program again with`strace`
3. Find the error message in the trace
4. Work upwards to find the first failing system call

这是一个简单的例子，但我会说至少 90% 的时间我使用 `strace`，我没有做任何更复杂的事情。这是完整的一步步调试公式：

1. 对来自程序的模糊系统错误消息感到沮丧
2. 用`strace`再次运行程序
3. 在trace中找到错误信息
4. 向上查找第一个失败的系统调用

There’s a very good chance the system call in step 4 shows you what went wrong.

第 4 步中的系统调用很有可能会告诉您出了什么问题。

## Some tips

##  一些技巧

Before walking through a more complicated example, here are some useful tips for using `strace` effectively:

### `man` is your friend

在浏览一个更复杂的例子之前，这里有一些有效使用 `strace` 的有用提示：

### `man` 是你的朋友

On many \*nix systems, you can get a list of all kernel system calls by running `man syscalls`. You’ll see things like `brk(2)`, which means you can get more information by running `man 2 brk`.

在许多\*nix 系统上，您可以通过运行`man syscalls` 获得所有内核系统调用的列表。你会看到类似 `brk(2)` 的内容，这意味着你可以通过运行 `man 2 brk` 来获取更多信息。

One little gotcha: `man 2 fork` shows me the man page for the `fork()` wrapper in GNU `libc`, which is actually now implemented using the `clone` system call instead. The semantics of `fork` are the same, but if I write a program using `fork()` and `strace` it, I won’t find any `fork` calls in the trace, just `clone` calls. Gotchas like that are only confusing if you’re comparing
source code to `strace` output.

一个小问题：`man 2 fork` 向我展示了GNU `libc` 中的 `fork()` 包装器，现在实际上是使用 `clone` 系统调用来实现的。 `fork` 的语义是相同的，但是如果我使用 `fork()` 和 `strace` 编写一个程序，我不会在跟踪中找到任何 `fork` 调用，只有 `clone` 调用。如果您进行比较，这样的问题只会令人困惑`strace` 输出的源代码。

### Use `-o` to save output to a file

### 使用 `-o` 将输出保存到文件

`strace` can generate a lot of output so it’s often helpful to store the trace in a separate file (as in the example above). It also avoids mixing up program output with `strace` output in the console.

`strace` 可以产生大量输出，所以它通常很有帮助将跟踪存储在单独的文件中（如上例所示）。它还避免将程序输出与`strace` 控制台输出。

### Use `-s` to see more argument data

### 使用`-s`查看更多参数数据

You might have noticed that the second part of the error message doesn’t appear in the example trace above. That’s because `strace` only shows the first 32 bytes of string arguments by default. If you need to capture more, add something like `-s 128` to the `strace` invocation.

您可能已经注意到错误消息的第二部分没有出现在上面的示例跟踪中。那是因为 `strace` 只显示字符串的前 32 个字节默认参数。如果您需要捕获更多内容，请在 `strace` 调用中添加类似 `-s 128` 的内容。

### `-y` makes it easier to track files/sockets/etc

### `-y` 可以更轻松地跟踪文件/套接字/等

“Everything is a file” means \*nix systems do all IO using file descriptors, whether it’s to an actual file or over networks or through interprocess pipes. That’s convenient for programming, but makes it harder to follow what’s really going on when you see generic `read` and `write` in the system call trace.

“一切都是文件”意味着 \*nix 系统使用文件描述符来执行所有 IO，无论是对实际文件还是在其他文件上网络或通过进程间管道。这对编程来说很方便，但更难遵循真正的当您在系统调用跟踪中看到通用的“read”和“write”时，继续进行。

Adding the `-y` switch makes `strace` annotate every file descriptor in the output with a note about what it points to.

添加 `-y` 开关使 `strace` 注释输出中的每个文件描述符，并附上关于它指向什么。

### Attach to an already-running process with `-p`

### 使用 `-p` 附加到一个已经运行的进程

As we’ll see in the example later, sometimes you want to trace a program that’s already running. If you know it’s running as process 1337 (say, by looking at the output of `ps`), you can trace it like this:

正如我们稍后将在示例中看到的，有时您想跟踪一个已经在运行的程序。如果你知道这是
作为进程 1337 运行（例如，通过查看 `ps` 的输出），您可以像这样跟踪它：

```shell
$ strace -p 1337
...system call trace output...

```

You probably need root.

你可能需要root。

### Use `-f` to follow child processes

### 使用 `-f` 来跟踪子进程

By default, `strace` only traces the one process. If that  process spawns a child process, you’ll see the system call for spawning the process (normally `clone` nowadays), but not any of the calls made by the child process.

默认情况下，`strace` 只跟踪一个进程。如果说进程产生一个子进程，你会看到产生进程的系统调用（现在通常是 `clone`），但不会看到子进程进行的任何调用。

If you think the bug is in a child process, you’ll need to use the `-f` switch to enable tracing it. A downside is that the output can be more confusing. When tracing one process and one thread, `strace ` can show you a single stream of call events. When tracing multiple processes, you might see the start of a call cut off with `<unfinished ...>`, then a bunch of calls for other threads of execution, before seeing the end of the original call with `<... foocall resumed>`. Alternatively, you can separate all the traces into different files by using the `-ff` switch as well
(see [the `strace`manual](https://linux.die.net/man/1/strace) for details).

如果您认为错误在子进程中，则需要使用“-f”开关来启用跟踪。一个缺点是输出可能更多令人困惑。跟踪一个进程和一个线程时，strace 可以向您显示单个呼叫事件流。跟踪多个进程时，您可能会看到呼叫开始中断使用 `<unfinished ...>`，然后是一堆其他的调用执行线程，在看到带有 `<... foocall resumed>` 的原始调用结束之前。或者，您可以将所有也可以使用 `-ff` 开关跟踪到不同的文件（见[`strace`\手册](https://linux.die.net/man/1/strace)了解详情)。

### You can filter the trace with `-e`

### 您可以使用`-e`过滤跟踪

As you’ve seen, the default trace output is a firehose of all system calls. You can filter which calls get traced using the `-e` flag (see [the `strace` manual](https://linux.die.net/man/1/strace)).
The main advantage is that it’s faster to run the program under a filtered `strace` than to trace everything and `grep` the results later. Honestly, I don’t bother most of the time.

如您所见，默认跟踪输出是所有系统调用的流水。您可以过滤哪些调用被跟踪使用 `-e` 标志（参见 [`strace` 手册](https://linux.die.net/man/1/strace))。

主要优点是在过滤后的 `strace` 下运行程序比跟踪所有内容然后 `grep` 结果更快。老实说，我大部分时间都不打扰。

### Not all errors are bad

### 不是所有的错误都是坏的

A simple and common example is a program searching for a file in multiple places, like a shell searching for which`bin/` directory has an executable:

一个简单而常见的例子是一个程序在多个地方搜索一个文件，就像一个 shell 搜索哪个
`bin/` 目录有一个可执行文件：

```shell
$ strace sh -c uname
...
stat("/home/user/bin/uname", 0x7ffceb817820) = -1 ENOENT (No such file or directory)
stat("/usr/local/bin/uname", 0x7ffceb817820) = -1 ENOENT (No such file or directory)
stat("/usr/bin/uname", {st_mode=S_IFREG|0755, st_size=39584, ...}) = 0
...

```

The “last failed call before the error message” heuristic is pretty good at finding relevent errors. In any case, working from the bottom up makes sense.

“错误消息之前的最后一次失败调用”启发式非常适合查找相关错误。任何状况之下，
自下而上的工作是有道理的。

### C programming guides are good for understanding system calls

### C 编程指南有助于理解系统调用

Standard C library calls aren’t system calls, but they’re only thin layers on top. So if you understand (even just roughly) how to do something in C, it’s easier to read a system call trace. For example, if you’re having trouble debugging networking system calls, you could try skimming through [Beej’s classic Guide to Network Programming](https://beej.us/guide/bgnet/html/index.html).

标准 C 库调用不是系统调用，但它们只是顶层的薄层。所以如果你明白（即使只是
大致）如何在 C 中做某事，更容易阅读系统调用跟踪。例如，如果您遇到问题
调试网络系统调用，你可以尝试浏览[Beej的经典网络编程指南](https://beej.us/guide/bgnet/html/index.html)。

## A more complicated debugging example

## 一个更复杂的调试例子

As I said, that simple debugging example is representative of most of my `strace` usage. However, sometimes a little more detective work is required, so here’s a slightly more complicated (and real) example.

正如我所说，这个简单的调试示例代表了我的大部分 `strace` 用法。然而，有时更多的侦探工作是需要，所以这里有一个稍微复杂（和真实）的例子。

[`bcron`](https://untroubled.org/bcron/) is a job scheduler that’s yet another implementation of the classic \*nix `cron` daemon. It’s been installed on a server, but here’s what happens when someone tries to edit a job schedule:

[`bcron`](https://untroubled.org/bcron/) 是一份工作调度程序，它是经典 \*nix `cron` 守护进程的另一个实现。它已安装在服务器上，但会发生以下情况当有人试图编辑工作计划时：

```shell
# crontab -e -u logs
bcrontab: Fatal: Could not create temporary file

```

Okay, so bcron tried to write some file, but it couldn’t, and isn’t telling us why. This is a debugging job for `strace`:

好的，所以 bcron 试图写一些文件，但它不能，也没有告诉我们为什么。这是一个调试工作
`strace`：

```shell
# strace -o /tmp/trace crontab -e -u logs
bcrontab: Fatal: Could not create temporary file
# cat /tmp/trace
...
openat(AT_FDCWD, "bcrontab.14779.1573691864.847933", O_RDONLY) = 3
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f82049b4000
read(3, "#Ansible: logsagg\n20 14 * * * lo"..., 8192) = 150
read(3, "", 8192)                       = 0
munmap(0x7f82049b4000, 8192)            = 0
close(3)                                = 0
socket(AF_UNIX, SOCK_STREAM, 0)         = 3
connect(3, {sa_family=AF_UNIX, sun_path="/var/run/bcron-spool"}, 110) = 0
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f82049b4000
write(3, "156:Slogs\0#Ansible: logsagg\n20 1"..., 161) = 161
read(3, "32:ZCould not create temporary f"..., 8192) = 36
munmap(0x7f82049b4000, 8192)            = 0
close(3)                                = 0
write(2, "bcrontab: Fatal: Could not creat"..., 49) = 49
unlink("bcrontab.14779.1573691864.847933") = 0
exit_group(111)                         = ?
+++ exited with 111 +++
```



There’s the error message `write` near the end, but a couple of things are different this time. First, there’s no relevant system call error that happens before it. Second, we see that the error message has just been `read` from somewhere else. It looks like the real problem is happening somewhere else, and `bcrontab` is just replaying the message.

接近尾声时有错误消息“write”，但有几个这次的情况有所不同。首先，在它之前没有发生相关的系统调用错误。其次，我们看到错误消息刚刚从某处“读取”别的。看起来真正的问题发生在其他地方，而 `bcrontab` 只是在重播消息。

If you look at `man 2 read`, you’ll see that the first argument (the 3) is a file descriptor, which is what \*nix uses for all IO handles. How do you know what file descriptor 3 represents? In this specific case, you could run `strace with the `-y` switch (as explained above) and it would tell you
automatically, but it’s useful to know how to read and analyse traces to figure things like this out.

如果你看一下‘man 2 read’，你会看到第一个参数（3）是一个文件描述符，它是 \*nix 用于所有 IO 句柄的。你怎么知道是什么文件描述符3代表？在这种特定情况下，您可以运行 `strace`
使用 `-y` 开关（如上所述），它会告诉你自动，但了解如何读取和分析跟踪以解决此类问题很有用。

A file descriptor can come from one of many system calls (depending on whether it’s a descriptor for the console, a network socket, an actual file, or something else), but in any case we can search for calls returning 3 (i.e., search for “= 3” in the trace). There are two in this trace: the `openat` at the top, and the `socket` in the middle. `openat` opens a file, but the `close(3)` afterwards shows that it gets closed again. (Gotcha: file descriptors can be reused as they’re opened and closed.) The `socket` call is the relevant one (it’s the last one before the `read`), which tells us `bcrontab` is talking to something over a network socket. The next line,`connect` shows file descriptor 3 being configured as a Unix domain socket connection to `/var/run/bcron-spool`.

文件描述符可以来自许多系统调用之一（取决于它是控制台的描述符，还是网络套接字、实际文件或其他东西），但无论如何我们都可以搜索返回 3 的调用（即搜索对于跟踪中的“= 3”）。此跟踪中有两个：顶部的“openat”和中间的“socket”。 `openat` 打开一个文件，但之后的 `close(3)` 显示它再次关闭。 （陷阱：文件描述符可以在打开和关闭时重用。）`socket` 调用是相关的（它是调用之前的最后一个）`read`)，它告诉我们 `bcrontab` 正在通过网络套接字与某些东西交谈。下一行，`connect` 显示文件描述符 3 被配置为 Unix域套接字连接到`/var/run/bcron-spool`。

So now we need to figure out what’s listening on the other side of the Unix socket. There are a couple of neat tricks for that, both useful for debugging server deployments. One is to use `netstat` or the newer `ss` (“socket status”). Both commands describe active network sockets on
the system, and take the `-l` switch for describing listening (server) sockets, and the `-p` switch to get information about what program is using the socket. (There are many more useful options, but those two are enough to get this job done.)

所以现在我们需要弄清楚 Unix 套接字的另一端正在侦听什么。有几个整洁的技巧，都对调试服务器部署很有用。一种是使用 `netstat` 或更新的 `ss`（“套接字状态”）。这两个命令都描述了活动的网络套接字系统，并使用 `-l` 开关来描述监听（服务器）套接字，以及“-p”开关以获取有关，什么程序正在使用套接字。 （还有更多有用的选项，但这两个足以得到这份工作完毕。）

```shell
# ss -pl | grep /var/run/bcron-spool
u_str LISTEN 0   128   /var/run/bcron-spool 1466637   * 0   users:(("unixserver",pid=20629,fd=3))

```



That tells us that the listener is a command `unixserver` running as process ID 20629. (It’s a coincidence that it’s also using file descriptor 3 for the socket.) 

这告诉我们侦听器是一个命令 `unixserver` 作为进程 ID 20629 运行。（巧合的是，它还为套接字使用了文件描述符 3。）

The second really useful tool for finding the same information is `lsof`. It can list all open files (or file descriptors) on the system.
Alternatively, we can get information about a specific file:

查找相同信息的第二个真正有用的工具是 `lsof`。它可以列出系统上所有打开的文件（或文件描述符）。
或者，我们可以获取有关特定文件的信息：

```shell
# lsof /var/run/bcron-spool
COMMAND   PID   USER  FD  TYPE  DEVICE              SIZE/OFF  NODE    NAME
unixserve 20629 cron  3u  unix  0x000000005ac4bd83  0t0       1466637 /var/run/bcron-spool type=STREAM

```

Process 20629 is a long-running server, so we can attach `strace` to it using something like `strace -o /tmp/trace -p 20629`. If we then try to edit the cron schedule
in another terminal, we can capture a trace while the error is happening. Here’s the result:

进程 20629 是一个长时间运行的服务器，因此我们可以使用类似 `strace -o /tmp/trace -p 20629` 的东西将 `strace` 附加到它。如果我们然后尝试编辑 cron 计划
在另一个终端中，我们可以在错误发生时捕获跟踪。结果如下：

```
accept(3, NULL, NULL)                   = 4
clone(child_stack=NULL, flags=CLONE_CHILD_CLEARTID|CLONE_CHILD_SETTID|SIGCHLD, child_tidptr=0x7faa47c44810) = 21181
close(4)                                = 0
accept(3, NULL, NULL)                   = ? ERESTARTSYS (To be restarted if SA_RESTART is set)
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=21181, si_uid=998, si_status=0, si_utime=0, si_stime=0} ---
wait4(0, [{WIFEXITED(s) && WEXITSTATUS(s) == 0}], WNOHANG|WSTOPPED, NULL) = 21181
wait4(0, 0x7ffe6bc36764, WNOHANG|WSTOPPED, NULL) = -1 ECHILD (No child processes)
rt_sigaction(SIGCHLD, {sa_handler=0x55d244bdb690, sa_mask=[CHLD], sa_flags=SA_RESTORER|SA_RESTART, sa_restorer=0x7faa47ab9840}, {sa_handler=0x55d244bdb690, sa_mask=[CHLD], sa_flags=SA_RESTORER|SA_RESTART, sa_restorer=0x7faa47ab9840}, 8) = 0
rt_sigreturn({mask=[]})                 = 43
accept(3, NULL, NULL)                   = 4
clone(child_stack=NULL, flags=CLONE_CHILD_CLEARTID|CLONE_CHILD_SETTID|SIGCHLD, child_tidptr=0x7faa47c44810) = 21200
close(4)                                = 0
accept(3, NULL, NULL)                   = ? ERESTARTSYS (To be restarted if SA_RESTART is set)
--- SIGCHLD {si_signo=SIGCHLD, si_code=CLD_EXITED, si_pid=21200, si_uid=998, si_status=111, si_utime=0, si_stime=0} ---
wait4(0, [{WIFEXITED(s) && WEXITSTATUS(s) == 111}], WNOHANG|WSTOPPED, NULL) = 21200
wait4(0, 0x7ffe6bc36764, WNOHANG|WSTOPPED, NULL) = -1 ECHILD (No child processes)
rt_sigaction(SIGCHLD, {sa_handler=0x55d244bdb690, sa_mask=[CHLD], sa_flags=SA_RESTORER|SA_RESTART, sa_restorer=0x7faa47ab9840}, {sa_handler=0x55d244bdb690, sa_mask=[CHLD], sa_flags=SA_RESTORER|SA_RESTART, sa_restorer=0x7faa47ab9840}, 8) = 0
rt_sigreturn({mask=[]})                 = 43
accept(3, NULL, NULL
```

(The last `accept` doesn’t complete during the trace period.) Unfortunately, once again, this trace doesn’t contain the error we’re after. We don’t see any of the messages that we saw `bcrontab` sending to and receiving from the socket. Instead, we see a lot of process management ( `clone`,
`wait4`, `SIGCHLD`, etc.). This process is spawning a child process, which we can guess is doing the real work. If we want to catch a trace of that, we have to add `-f` to the `strace` invocation. Here’s what we find if we search for the error message after getting a new trace with `strace -f -o /tmp/trace -p 20629`:

```
21470 openat(AT_FDCWD, "tmp/spool.21470.1573692319.854640", O_RDWR|O_CREAT|O_EXCL, 0600) = -1 EACCES (Permission denied)
21470 write(1, "32:ZCould not create temporary f"..., 36) = 36
21470 write(2, "bcron-spool[21470]: Fatal: logs:"..., 84) = 84
21470 unlink("tmp/spool.21470.1573692319.854640") = -1 ENOENT (No such file or directory)
21470 exit_group(111)                   = ?
21470 +++ exited with 111 +++

```



Now we’re getting somewhere. Process ID 21470 is getting a permission denied error trying to create a file at the path `tmp/spool.21470.1573692319.854640` (relative to the current working directory). If we just knew the current working directory, we would know the full path and could figure out why the process can’t create create its temporary file there. Unfortunately, the process has already exited, so we can’t just use `lsof -p 21470` to find out the current 

现在我们正在到达某个地方。进程 ID 21470 在尝试创建文件时出现权限被拒绝错误路径`tmp/spool.21470.1573692319.854640`（相对于当前工作目录）。如果我们只知道当前的工作目录，我们就会知道完整路径并可以计算出找出为什么进程不能在那里创建它的临时文件。不幸的是，进程已经退出，所以我们不能只使用 `lsof -p 21470` 来找出当前

directory, but we can work backwards looking for PID 21470 system calls that change directory. (If there aren’t any, PID 21470 must have inherited it from its parent, and we can `lsof -p` that.) That system call is `chdir` (which is easy to find out using today’s web search engines). Here’s the result of working backwards through the trace, all the way to the server PID 20629:

目录，但我们可以向后查找更改目录的 PID 21470 系统调用。 （如果没有，PID 21470 一定是从它的父级继承的，我们可以`lsof -p`那个。）那个系统调用是`chdir`（这很容易
找出使用今天的网络搜索引擎）。这是通过跟踪向后工作的结果，一直到服务器PID 20629：

```
20629 clone(child_stack=NULL, flags=CLONE_CHILD_CLEARTID|CLONE_CHILD_SETTID|SIGCHLD, child_tidptr=0x7faa47c44810) = 21470
...
21470 execve("/usr/sbin/bcron-spool", ["bcron-spool"], 0x55d2460807e0 /* 27 vars */) = 0
...
21470 chdir("/var/spool/cron")          = 0
...
21470 openat(AT_FDCWD, "tmp/spool.21470.1573692319.854640", O_RDWR|O_CREAT|O_EXCL, 0600) = -1 EACCES (Permission denied)
21470 write(1, "32:ZCould not create temporary f"..., 36) = 36
21470 write(2, "bcron-spool[21470]: Fatal: logs:"..., 84) = 84
21470 unlink("tmp/spool.21470.1573692319.854640") = -1 ENOENT (No such file or directory)
21470 exit_group(111)                   = ?
21470 +++ exited with 111 +++
```

(If you’re getting lost here, you might want to read [my previous post\
about \*nix process management and shells](http://theartofmachinery.com/2018/11/07/writing_a_nix_shell.html).) Okay, so the server PID 20629 doesn’t have permission to create a file at `/var/spool/cron/tmp/spool.21470.1573692319.854640`. The most likely reason would be classic \*nix filesystem permission settings. Let’s check:

（如果你在这里迷路了，你可能想阅读[我以前的帖子关于 \*nix 进程管理和 shell](http://theartofmachinery.com/2018/11/07/writing_a_nix_shell.html)。好的，所以服务器 PID 20629 没有创建文件的权限在`/var/spool/cron/tmp/spool.21470.1573692319.854640`。这最可能的原因是经典的 \*nix 文件系统权限设置。让我们检查：

```shell
# ls -ld /var/spool/cron/tmp/
drwxr-xr-x 2 root root 4096 Nov  6 05:33 /var/spool/cron/tmp/
# ps u -p 20629
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
cron     20629  0.0  0.0   2276   752 ? Ss   Nov14   0:00 unixserver -U /var/run/bcron-spool -- bcron-spool

```

There’s the problem! The server is running as user `cron`, but only `root` has permissions to write to that `/var/spool/cron/tmp/` directory. A simple `chown cron  /var/spool/cron/tmp/` makes `bcron` work properly. (If that weren’t the problem, the next most likely suspect would be a kernel security module like SELinux or AppArmor, so I’d check the kernel logs with `dmesg`.)

问题来了！服务器以用户“cron”的身份运行，但只有 `root` 有权限写入该 `/var/spool/cron/tmp/` 目录。一个简单的 `chown cron /var/spool/cron/tmp/` 使 `bcron` 正常工作。 （如果这不是问题，下一个最有可能的怀疑是像 SELinux 或 AppArmor 这样的内核安全模块，所以我会用 `dmesg` 检查内核日志。）

## Summary

##  概括

System call traces can be overwhelming at first, but I hope I’ve shown that they’re a fast way to debug a whole class of common deployment problems. Imagine trying to debug that multi-process `bcron` problem using a stepping debugger.

系统调用跟踪一开始可能会让人不知所措，但我希望我已经证明它们是一种快速调试整体的方法
一类常见的部署问题。想象一下尝试使用步进调试器调试多进程“bcron”问题。

Working back through a chain of system calls takes practice, but as I said, most of the time I use `strace` I just get a trace and look for errors, working from the bottom up. In any case, `strace` has saved me hours and hours of debugging time. I hope it’s useful for you, too.

通过一系列系统调用返回需要练习，但正如我所说，大多数时候我使用 `strace` 我只是得到一个跟踪并查找错误，从底部开始工作向上。无论如何，`strace` 为我节省了数小时的时间
调试时间。我希望它对你也有用。

