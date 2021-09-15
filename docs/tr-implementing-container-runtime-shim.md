# Implementing Container Runtime Shim: runc

# 实现容器运行时 Shim：runc

December 15, 2019 (Updated: August 24, 2021)

## What is a shim?

## 什么是垫片？

[A container runtime shim](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#runtime-shims) is a piece of software that resides in between [a container manager ( _containerd_, _cri-o_, _podman_)](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-management) and [a container runtime ( _runc_ , _crun_)](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes) solving the integration problem of these counterparts.

[一个容器运行时垫片](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#runtime-shims) 是一个位于 [a container manager ( _containerd_, _cri-o_, _podman_)](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-management) 和 [容器运行时 ( _runc_ , _crun_)](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes) 解决这些同行的集成问题。

![Layered Docker architecture: docker (cli) -> dockerd -> containerd -> containerd-shim -> runc](http://iximiuz.com/implementing-container-runtime-shim/docker-containerd-runc-2000-opt.png)



_Layered Docker architecture_

_分层Docker架构_

The easiest way to spot a shim is to inspect the process tree on a Linux host with a running docker container:

发现 shim 的最简单方法是在运行着 docker 容器的 Linux 主机上检查进程树：

![Spotting container runtime shim process](http://iximiuz.com/implementing-container-runtime-shim/ps-shim-example.png)

_`ps auxf` output on a host running `docker run -it ubuntu bash`; notice `containerd-shim` process in between `containerd` and `bash`._

_`ps auxf` 在运行 `docker run -it ubuntu bash` 的主机上输出；注意`containerd`和`bash`之间的`containerd-shim`过程。_

On the one hand, runtimes need shims to be able to survive managers restarts. On the other hand, shims are helping container managers to deal with the quirky behavior of runtimes. As a part of [the container manager implementation series](http://iximiuz.com/en/posts/conman-the-container-manager-inception/), we will try to create our own shim and then integrate it with [ _conman_](https://github.com/iximiuz/conman), an experimental container manager. Hopefully, during the development, we will gain an in-depth understanding of the topic.

一方面，运行时需要垫片才能在管理器重启后幸存下来。另一方面，垫片正在帮助容器管理器处理运行时的古怪行为。作为[容器管理器实现系列](http://iximiuz.com/en/posts/conman-the-container-manager-inception/)的一部分，我们将尝试创建自己的shim，然后将其与[_conman_](https://github.com/iximiuz/conman)，一个实验性的容器管理器。希望在开发过程中，我们能够深入了解该主题。

However, before jumping to the shim development, we need to familiarize ourselves with the container runtime component of the choice. Unsurprisingly, _conman_ uses _runc_ as a container runtime, so I will start the article by covering basic _runc_ use cases alongside its design quirks. Then I'll show the naive way to use _runc_ from code and explain some related pitfalls. The final part of the article will provide an overview of the shim's design.

但是，在跳转到 shim 开发之前，我们需要熟悉所选的容器运行时组件。不出所料，_conman_ 使用 _runc_ 作为容器运行时，所以我将通过介绍基本的 _runc_ 用例及其设计怪癖来开始这篇文章。然后我将展示从代码中使用 _runc_ 的天真方法并解释一些相关的陷阱。文章的最后一部分将概述垫片的设计。

## Playing with runc

## 玩 runc

The detailed explanation of _what is a container runtime_ and _why do we need one_ is out of the scope of this article. If you are lacking this knowledge, read [this section first](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes), then check out [ _runc_](https://github.com/opencontainers/runc) README file, and try out `runc --help`.

_什么是容器运行时_和_为什么需要一个_的详细解释超出了本文的范围。如果您缺乏这方面的知识，请先阅读 [本节](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/#container-runtimes)，然后查看 [ _runc_](https://github.com/opencontainers/runc) README 文件，并尝试使用 `runc --help`。

Long story short, _runc_ is **a command-line tool** for running containerized applications. If this sentence sounds too fancy, you can think of _runc_ as just a tool to spawn a new ordinary Linux process but inside of an isolated environment. The isolation among other things includes a dedicated root file system and a new process tree and is being achieved via Linux _namespaces_ and _cgroups_ facilities. Let's call this new process that _runc_ launches **a container process**. This process becomes the first process (i.e. _PID=1_) inside of the newly started container. We will be referring to this process frequently throughout the rest of the article.

长话短说，_runc_ 是**一个用于运行容器化应用程序的命令行工具**。如果这句话听起来太花哨，您可以将 _runc_ 视为只是在隔离环境中生成新的普通 Linux 进程的工具。除其他外，隔离包括一个专用的根文件系统和一个新的进程树，并且正在通过 Linux _namespaces_ 和 _cgroups_ 设施实现。让我们称这个新进程 _runc_ 启动 **一个容器进程**。该进程成为新启动容器内的第一个进程（即 _PID=1_）。我们将在本文的其余部分经常提到这个过程。

Even though _runc_ is most often used behind _Docker_ (through _containerd_) or _cri-o_ daemons while both oftentimes reside behind a _kubelet_, it's a standalone executable file, i.e. _runc_ is by no means a library. We will see further on how this design makes the usage of _runc_ from another software somewhat complicated, but now let's try to just play with this tool in our favorite terminal emulator:

尽管 _runc_ 最常在 _Docker_（通过 _containerd_）或 _cri-o_ 守护进程后面使用，而两者通常都驻留在 _kubelet_ 后面，但它是一个独立的可执行文件，即 _runc_ 绝不是一个库。我们将进一步了解这种设计如何使来自另一个软件的 _runc_ 的使用有些复杂，但现在让我们尝试在我们最喜欢的终端模拟器中使用这个工具：

_**NB:** runc supports running containers controlled by a [Linux pseudoterminal](http://iximiuz.com/en/posts/linux-pty-what-powers-docker-attach-functionality/), but today we will not use this functionality and leave its consideration to a separate article._

_**注意：** runc 支持运行由 [Linux 伪终端](http://iximiuz.com/en/posts/linux-pty-what-powers-docker-attach-functionality/) 控制的容器，但今天我们不会使用此功能并将其考虑留给单独的文章。_

```bash
# Prepare some directories.
$ mkdir -p container1/rootfs
$ cd container1

# Use docker to create a root
# filesystem for our new container.
$ sudo bash -c 'docker export $(docker create busybox) |tar -C rootfs -xvf -'

# Create a default bundle (i.e. container) config.
$ runc spec

# Change the default command (sh) to
# sh -c 'echo Hi, my PID is $$;sleep 10;echo Bye Bye'
$ sed -i 's/"sh"/"sh", "-c", "echo Hi, my PID is $$; sleep 10; echo Bye Bye"/' config.json

# Do not use pseudoterminal (PTY) to control the container process.
# PTY use cases lie out of the scope of this article.
$ sed -i 's/"terminal": true/"terminal": false/' config.json

# Review the bundle config (optional).
$ less config.json

# Run the container with ID cont1.
$ sudo runc run cont1

```

If we check the corresponding process tree using `ps axfo pid,ppid,command` in the separate terminal session we will see something similar to this:

如果我们在单独的终端会话中使用 `ps axfo pid,ppid,command` 检查相应的进程树，我们将看到类似的内容：

![Process tree with running container](http://iximiuz.com/implementing-container-runtime-shim/runc-foreground-ps-tree.png)

The process hierarchy seems absolutely normal. Our login _bash_ session ( _PID 9503_) [fork-execed](https://en.wikipedia.org/wiki/Fork%E2%80%93exec) _runc_ process which in turn forked itself (probably due to [PID namespace implementation related reasons](http://man7.org/linux/man-pages/man7/pid_namespaces.7.html)) and the second fork finally launched the _sh_ shell ( _PID 22437_) inside of the containerized environment.

进程层次结构看起来绝对正常。我们的登录 _bash_ 会话 (_PID 9503_) [fork-execed](https://en.wikipedia.org/wiki/Fork%E2%80%93exec) _runc_ 进程又分叉了自己（可能是由于 [PID 命名空间实现相关原因](http://man7.org/linux/man-pages/man7/pid_namespaces.7.html)) 和第二个 fork 终于在容器化环境中启动了 _sh_ shell (_PID 22437_)。

Now, let's take a look at the produced output:

现在，让我们看看生成的输出：

```
Hi, my PID is 1
Bye Bye

```

Notice, that even though from the host system we see the PID of the _sh_ process as _22437_, our tiny script printed out `Hi, my PID is 1`. That's perfect proof that it got its own process ID namespace. We also should set our eyes on the fact that the _stdout_ content of the container was just printed out to our terminal, meaning that _stdio streams_ of our login bash shell have been [passed through](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#-pass-through) to the container and then set as its own _stdio streams_. **That’s an important observation we need to keep in mind.** We can try to represent the process structure as follows:

请注意，即使从主机系统我们看到 _sh_ 进程的 PID 为 _22437_，我们的小脚本打印出“嗨，我的 PID 是 1”。这是它拥有自己的进程 ID 命名空间的完美证明。我们还应该关注容器的 _stdout_ 内容刚刚打印到我们的终端这一事实，这意味着我们登录 bash shell 的 _stdio 流_ 已经[通过](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#-pass-through) 到容器，然后设置为它自己的 _stdio 流_。 **这是我们需要牢记的一个重要观察。** 我们可以尝试将流程结构表示如下：

![runc reparenting container process](http://iximiuz.com/implementing-container-runtime-shim/runc-foreground-stdio.png)

The approach to run containers we just explored is called [**foreground**](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#foreground) in _runc_ terminology. It means that the _runc_ process always stays in between the container process and the launching process (bash shell in our example). There is another mode supported by _runc_, called [**detached**](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#detached). Let's try to use it:

我们刚刚探索的运行容器的方法在 _runc_ 术语中称为 [**foreground**](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#foreground)。这意味着 _runc_ 进程始终位于容器进程和启动进程（在我们的示例中为 bash shell)之间。 _runc_ 支持另一种模式，称为 [**detached**](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#detached)。让我们尝试使用它：

```bash
# from the container1 folder
sudo runc run --detach cont1-detached

```

Notice, how _runc_ releases the execution back to our login shell almost instantly. If we check the corresponding process tree, we will see the following picture:

请注意，_runc_ 如何几乎立即将执行释放回我们的登录 shell。如果我们查看对应的进程树，会看到如下图：

![runc in detached mode](http://iximiuz.com/implementing-container-runtime-shim/runc-detached-ps-tree.png)

Seems like _runc_ exited completely after it has spawned the container process. And the container process has been reparented to the host's _PID 1_ process. There is no connection between the launching process (our login shell) and the container process... except for the passed-through _stdio streams_! The output of the container process once again has been printed out to our terminal, but this time the lines produced by the container are interleaved with the login shell prompt appearance:

似乎 _runc_ 在产生容器进程后完全退出。并且容器进程已重新分配给主机的 _PID 1_ 进程。启动进程（我们的登录shell）和容器进程之间没有任何联系……除了传递的_stdio流_！容器进程的输出再次打印到我们的终端，但这次容器产生的行与登录 shell 提示出现交错：

```
$ sudo runc run --detach cont1-detached
Hi, my PID is 1
$ Bye Bye

```

We can try to depict the corresponding structure of the involved processes and their _stdio streams_ as follows:

我们可以尝试将所涉及的进程及其_stdio流_的相应结构描述如下：

![container and container runtime - process tree](http://iximiuz.com/implementing-container-runtime-shim/runc-detached-stdio.png)

That's how the detached mode is described by _runc_ [documentation](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#detached):

这就是 _runc_ [文档](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#detached) 描述分离模式的方式：

> In contrast to foreground mode, in detached mode there is no long-running foreground runc process once the container has started. In fact, there is no long-running `runc` process at all. However, this means that it is up to the caller to handle the stdio after `runc` has set it up for you. In a shell this means that the `runc` command will exit and control will return to the shell, after the container has been set up.

> 与前台模式相反，在分离模式下，一旦容器启动，就没有长时间运行的前台 runc 进程。事实上，根本没有长时间运行的 `runc` 进程。然而，这意味着在 `runc` 为你设置好 stdio 之后，由调用者来处理它。在 shell 中，这意味着在设置容器后，`runc` 命令将退出并且控制权将返回到 shell。

> You can run `runc` in detached mode in one of the following ways:

> 您可以通过以下方式之一在分离模式下运行 `runc`：

> - `runc run -d` ... which operates similar to `runc run` but is detached.
> - `runc create` followed by `runc start` which is the standard container lifecycle defined by the OCI runtime specification ( `runc create` sets up the container completely, waiting for `runc start` to begin execution of user code). 

> - `runc run -d` ... 其操作类似于 `runc run`，但是是分离的。
> - `runc create` 后跟 `runc start`，这是 OCI 运行时规范定义的标准容器生命周期（`runc create` 完全设置容器，等待 `runc start` 开始执行用户代码）。

> The main use-case of detached mode is for higher-level tools that want to be wrappers around `runc`. By running `runc` in detached mode, those tools have far more control over the container's `stdio` without `runc` getting in the way (most wrappers around `runc` like `cri-o` or `containerd` use detached mode for this reason).

> 分离模式的主要用例是用于希望成为 `runc` 包装器的高级工具。通过在分离模式下运行 `runc`，这些工具可以更好地控制容器的 `stdio`，而不会受到 `runc` 的影响（大多数围绕 `runc` 的包装器，如 `cri-o` 或 `containerd` 使用分离模式来这个原因）。

> Unfortunately using detached mode is a bit more complicated and requires more care than the foreground mode -- mainly because it is now up to the caller to handle the `stdio` of the container.

> 不幸的是，使用分离模式比前台模式要复杂一些，需要更多的关注——主要是因为现在由调用者来处理容器的`stdio`。

We need to admit, that the foreground mode has some significant drawbacks and the detached mode has been designed in a way to eliminate them. From [the same document](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#foreground):

我们需要承认，前台模式有一些明显的缺点，而分离模式的设计旨在消除它们。来自 [同一文档](https://github.com/opencontainers/runc/blob/201b06374548b64212f4ceb1529688d435e42899/docs/terminals.md#foreground)：

> The main drawback of the foreground mode of operation is that it requires a long-running foreground runc process. If you kill the foreground runc process then you will no longer have access to the stdio of the container (and in most cases this will result in the container dying abnormally due to SIGPIPE or some other error). By extension this means that any bug in the long-running foreground runc process (such as a memory leak) or a stray OOM-kill sweep could result in your container being killed through no fault of the user. In addition, there is no way in foreground mode of passing a file descriptor directly to the container process as its stdio (like --preserve-fds does).

> 前台操作模式的主要缺点是它需要一个长时间运行的前台 runc 进程。如果您终止前台 runc 进程，那么您将无法再访问容器的 stdio（并且在大多数情况下，这将导致容器由于 SIGPIPE 或其他一些错误而异常死亡）。通过扩展，这意味着长时间运行的前台 runc 进程中的任何错误（例如内存泄漏）或杂散的 OOM-kill 扫描都可能导致您的容器因用户的过错而被终止。此外，在前台模式下无法像其 stdio 那样将文件描述符直接传递给容器进程（就像 --preserve-fds 那样）。

## A naive attempt to use runc from code

## 从代码中使用 runc 的幼稚尝试

By this time we already know how quirky _runc_ is when it comes to handling its _stdio_ streams. We also know the desired mode to use if we want to run _runc_ from container management software. Let's try to make a naive program launching a container with the following _sh_ program:

到目前为止，我们已经知道 _runc_ 在处理其 _stdio_ 流时是多么古怪。如果我们想从容器管理软件运行 _runc_，我们还知道要使用的所需模式。让我们尝试使用以下 _sh_ 程序创建一个简单的程序来启动容器：

```bash
#!/bin/sh

function report_sigpipe() {
    echo "got SIGPIPE, exiting..." > /var/trap.log;
    exit 1;
}

trap report_sigpipe SIGPIPE;

while sleep 1
do
    echo "time is $(date)";
done

```

The idea of the script is fairly simple. Every second it writes the current time to its _stdout_. Additionally, it catches the _SIGPIPE_ signal (using [`trap`](http://man7.org/linux/man-pages/man1/trap.1p.html)) and reports it by writing `/var/trap.log ` file, then exits.

脚本的想法相当简单。它每秒将当前时间写入其 _stdout_。此外，它捕获 _SIGPIPE_ 信号（使用 [`trap`](http://man7.org/linux/man-pages/man1/trap.1p.html))并通过写入 `/var/trap.log ` 文件，然后退出。

Now, let's prepare a _runc_ bundle:

现在，让我们准备一个 _runc_ 包：

```bash
$ mkdir -p container2/rootfs
$ cd container2

$ sudo bash -c 'docker export $(docker create busybox) |tar -C rootfs -xvf -'

$ cat > rootfs/entrypoint.sh <<EOF
#!/bin/sh

function report_sigpipe() {
    echo "got SIGPIPE, exiting..." > /var/trap.log;
    exit 1;
}

trap report_sigpipe SIGPIPE;

while sleep 1
do
    echo "time is $(date)";
done
EOF

$ runc spec

$ sed -i 's/"sh"/"sh", "entrypoint.sh"/' config.json
$ sed -i 's/"terminal": true/"terminal": false/' config.json

# Make the rootfs writable to be able to save /var/trap.log.
$ sed -i 's/"readonly": true/"readonly": false/' config.json

```

Finally, let's quickly make a tiny _Go_ container management program. As any container manager, it should be interested in the container's _stdout_ content. It will be reading the data written by the container to its _stdout_ and reporting it back to us. And as we learned earlier, we need to use _runc_ in the detached mode:

最后，让我们快速制作一个小巧的_Go_容器管理程序。作为任何容器管理器，它应该对容器的 _stdout_ 内容感兴趣。它将读取容器写入其 _stdout_ 的数据并将其报告给我们。正如我们之前了解到的，我们需要在分离模式下使用 _runc_：

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    // We want to have a full control on the container's
    // stdout, so we are creating a pipe to redirect it.
    rd, wr, err := os.Pipe()
    if err != nil {
        panic(err)
    }
    defer rd.Close()
    defer wr.Close()

    // Start runc in detached mode.
    fmt.Println("Launching runc")

    cmd := exec.Command("runc", "run", "--detach", os.Args[1])
    cmd.Stdin = nil  // i.e. /dev/null
    cmd.Stderr = nil // i.e. /dev/null
    cmd.Stdout = wr
    if err := cmd.Run();err != nil {
        panic(err)
    }

    // Read some data from the container's stdout.
    buf := make([]byte, 1024)
    for i := 0;i < 10;i++ {
        n, err := rd.Read(buf)
        if err != nil {
            panic(err)
        }
        output := strings.TrimSuffix(string(buf[:n]), "\n")
        fmt.Printf("Container produced: [%s]\n", output)

    }

    // Get bored quickly, give up and exit.
    fmt.Println("We are done, exiting...")
}

```

Now let's try to use our manager to run a container with the aforementioned script:

现在让我们尝试使用我们的管理器来运行一个包含上述脚本的容器：

```bash
$ sudo `which go` run main.go cont2
Launching runc
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
Container produced: [time is Sun Dec 15 16:34:25 UTC 2019]
We are done, exiting...

```

It seems like a successful attempt. However, if we check the status of the container after the manager's exit we will notice that the container has been stopped as well:

这似乎是一次成功的尝试。但是，如果我们在管理器退出后检查容器的状态，我们会注意到容器也已停止：

```bash
$ sudo runc state cont2
{
"ociVersion": "1.0.1-dev",
"id": "cont2",
"pid": 0,
"status": "stopped",
"bundle": "/home/vagrant/container2",
"rootfs": "/home/vagrant/container2/rootfs",
"created": "2019-12-15T17:03:24.266072187Z",
"owner": ""
}

```

This is obviously not the desired behavior. Even though the container manager exited, we want to have the container running indefinitely. We even used the detached mode, so we know that the container process was reparented to the host's _init_ process right after the `runc run --detach` execution. So, what happened to the container?

这显然不是预期的行为。即使容器管理器退出，我们也希望容器无限期地运行。我们甚至使用了分离模式，所以我们知道容器进程在 `runc run --detach` 执行后立即被重新分配给主机的 _init_ 进程。那么，容器发生了什么？

It's not hard to guess, that the actual container termination reason should have something to do with the ability to write to its _stdout_. If we check the `/var/trap.log` file we will prove that theory quickly:

不难猜测，实际容器终止原因应该与写入其 _stdout_ 的能力有关。如果我们检查 `/var/trap.log` 文件，我们将很快证明这个理论：

```bash
$ cat rootfs/var/trap.log
got SIGPIPE, exiting...

```

**This exercise shows that if we want to keep control of the container's _stdio_ streams, the container process cannot be independent of the launching process.** And since we know, that the container manager can be restarted due to a crash, update, or some other reasons, that makes it impossible to launch _runc_ directly from the container manager process. Thus, we need to have a helper process living as long as the underlying container process and serving it. And such a process called **container runtime shim**.

**这个练习表明，如果我们想保持对容器的 _stdio_ 流的控制，容器进程不能独立于启动进程。** 而且我们知道，容器管理器可以因崩溃、更新、或其他一些原因，这使得无法直接从容器管理器进程启动 _runc_。因此，我们需要一个辅助进程与底层容器进程一样存活并为其提供服务。而这样的过程叫做**container runtime shim**。

## The full scope of the shim

## 垫片的完整范围

A container runtime shim is a lightweight daemon launching _runc_ and controlling the container process. The shim's process is tightly bound to the container's process but is completely detached from the manager's process. All the communications between the container and the manager happen through the shim. Examples of the production-ready shims out there is [conmon](https://github.com/containers/conmon) and containerd [_runtime shim_](https://github.com/containerd/containerd/blob/master/runtime/v2/shim.go). The shim usually is responsible for the following (and probably some other) things:

容器运行时 shim 是一个轻量级守护进程，它启动 _runc_ 并控制容器进程。 shim 的进程与容器的进程紧密绑定，但与管理器的进程完全分离。容器和管理器之间的所有通信都通过 shim 进行。生产就绪垫片的示例有 [conmon](https://github.com/containers/conmon) 和 containerd [_runtime shim_](https://github.com/containerd/containerd/blob/master/runtime/v2/shim.go）。垫片通常负责以下（可能还有其他一些)事情：

- Serves container's _stdout_ and _stderr_ streams even during the manager restart. This allows container managers to have containers _stdout_ and _stderr_ forwarded to log files at any given moment. This is the feature powering `docker logs <container>` and `kubectl logs <pod> -c <container>` commands. When a container manager is being requested to provide some container logs, it just can read them directly from the predefined location on disk. For instance, `kubectl logs` triggers the following delegation chain `kubectl <-- network --> Kubernetes Core API <-- network --> kubelet <-- (CRI gRPC API) --> CRI Runtime Service (cri-o , containerd, docker) <-- read() --> logs on node's disk`.

- 即使在管理器重启期间也提供容器的 _stdout_ 和 _stderr_ 流。这允许容器管理器在任何给定时刻将容器 _stdout_ 和 _stderr_ 转发到日志文件。这是支持 `docker logs <container>` 和 `kubectl logs <pod> -c <container>` 命令的功能。当容器管理器被要求提供一些容器日志时，它可以直接从磁盘上的预定义位置读取它们。例如，`kubectl logs` 触发以下委托链`kubectl <-- network --> Kubernetes Core API <-- network --> kubelet <-- (CRI gRPC API) --> CRI Runtime Service (cri-o , containerd, docker) <-- read() --> 登录节点的磁盘`。

- Attaches to a running container. Container managers usually provide ways to stream some data in and out from the container, including PTY-controlled scenarios. For that, the shim needs to keep the _stdin_ of the container also open. The shim can establish a socket server accepting connections and performing the streaming to and from the container's _stdio_ to the attached clients. This powers `kubectl run -i`, `podman run -i --tty`, as well as the famous PTY-controlled interactive Docker use case (eg. `docker run -it ubuntu:latest bash`). 

- 附加到正在运行的容器。容器管理器通常提供一些方法来将一些数据传入和传出容器，包括 PTY 控制的场景。为此，垫片需要保持容器的 _stdin_ 也打开。 shim 可以建立一个套接字服务器，接受连接并执行从容器的 _stdio_ 到附加客户端的流传输。这为 `kubectl run -i`、`podman run -i --tty` 以及著名的 PTY 控制的交互式 Docker 用例（例如`docker run -it ubuntu:latest bash`）提供支持。

- Keeps track of container exit code. In the detached mode, _runc_ deliberately daemonizes the container process by forking and then exiting the foreground process. The container process then gets reparented, by default to the host's _init_ process. Having containers detached leads to an absence of container status update. One way to address this problem is to make the shim process a [_subreaper_](http://iximiuz.com/en/posts/dealing-with-processes-termination-in-Linux/#awaiting-a-grandchild-process-termination). This way the container process will be reparented to the shim process. Then the shim can wait for the container's process termination and report its exit code to a predefined destination (eg. a file on the disk). The corresponding container manager can pick it up later.

- 跟踪容器退出代码。在分离模式下，_runc_故意通过fork然后退出前台进程来守护容器进程。然后容器进程被重新设置为父进程，默认情况下是主机的 _init_ 进程。分离容器会导致缺少容器状态更新。解决此问题的一种方法是使 shim 进程成为 [_subreaper_](http://iximiuz.com/en/posts/dealing-with-processes-termination-in-Linux/#awaiting-a-grandchild-process-终止）。这样，容器进程将被重新分配给 shim 进程。然后 shim 可以等待容器的进程终止并将其退出代码报告给预定义的目标（例如磁盘上的文件)。对应的容器管理器可以稍后取用。

- Synchronizes container manager with the container creation status. Since _runc_ daemonizes the container creation process we need a side-channel (eg. a Unix socket) to communicate the actual start (or failure) of the container back to the container manager. `runc` reports the container creation error via its _stderr_. Since exactly the same file descriptor, later on, can become the container process' _stderr_ the shim needs to carefully consume all the data in it until the runtime process termination occurs and report it to the container manager immediately. Thus, if we make a typo in our container command, the actual error will be reported back to us during the container creation phase: `docker run -it ubuntu bahs docker: Error response from daemon: OCI runtime create failed: container_linux.go: 345: starting container process caused "exec: \"bahs\": executable file not found in $PATH": unknown.`

- 将容器管理器与容器创建状态同步。由于 _runc_ 守护容器创建过程，我们需要一个侧通道（例如 Unix 套接字）将容器的实际启动（或失败）传达回容器管理器。 `runc` 通过其 _stderr_ 报告容器创建错误。由于完全相同的文件描述符稍后可以成为容器进程的 _stderr_，因此 shim 需要小心地使用其中的所有数据，直到运行时进程终止发生并立即将其报告给容器管理器。因此，如果我们在容器命令中打错字，实际的错误将在容器创建阶段报告给我们：`docker run -it ubuntu bahs docker: Error response from daemon: OCI runtime create failed: container_linux.go: 345：启动容器进程导致“exec：\”bahs\”：在$PATH 中找不到可执行文件”：未知。`

## Shim implementation overview

## Shim 实现概述

We already know that the container runtime shim has to be a long-lived daemon tightly bound to the container process. Its structure can be expressed as the following diagram:

我们已经知道容器运行时 shim 必须是一个与容器进程紧密绑定的长期存在的守护进程。其结构可用下图表示：

![Container runtime shim lifecycle](http://iximiuz.com/implementing-container-runtime-shim/shim-implementation.png)

The main process of the shim is short-lived and serves the purpose of the daemonization of the shim. It forks the actual shim daemon process, writes its PID on disk and exits immediately, leaving the shim detached from the launching process (i.e. a container manager). The long-lived shim daemon process starts from creating a new session and detaching its _stdio_ streams from the parent (by redirecting them to `/dev/null`). This is somewhat common steps for any daemon-alike software. Then it forks one more process, the predecessor of the container process. This process `exec` s `runc create` with the provided parameters (bundle dir, config.json, etc). The shim daemon process waits for the termination of the container predecessor and then reports the status of this operation back to the container manager. At this point, we have only a single shim daemon process and a detached container process. However, shim process is a subreaper of the container process. At last, the shim daemon process can start serving the container's _stdio_ streams as well as awaiting the container termination. Once the container termination status is known, the shim process writes it to a predefined location on disk and exits.

shim 的主要进程是短暂的，用于 shim 的守护进程。它派生实际的 shim 守护进程，将其 PID 写入磁盘并立即退出，使 shim 与启动进程（即容器管理器）分离。长期存在的 shim 守护进程从创建一个新会话并将其 _stdio_ 流与父级分离（通过将它们重定向到`/dev/null`）开始。对于任何类似守护进程的软件来说，这都是一些常见的步骤。然后它又派生出一个进程，即容器进程的前身。此过程使用提供的参数（包目录、config.json 等）执行 `exec` 和 `runc create`。 shim 守护进程等待容器前驱的终止，然后将此操作的状态报告回容器管理器。此时，我们只有一个 shim 守护进程和一个分离的容器进程。但是，shim 进程是容器进程的子收割者。最后，shim 守护进程可以开始为容器的 _stdio_ 流提供服务并等待容器终止。一旦知道容器终止状态，shim 进程将其写入磁盘上的预定义位置并退出。

## Stay tuned

##  敬请关注

In the following articles, we will see:

在以下文章中，我们将看到：

- [How to implement a basic container runtime shim in Rust](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/) and integrate it with the [experimental container manager](https://github.com/iximiuz/conman).
- [How to implement _attach_ functionality, bringing interactive containers support](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/) and revealing the magic behind `docker run -i` and ` kubectl run --stdin`.
- How to unleash the power of PTY-controlled containers and shed some light on technologies behind`docker run -it` and `kubectl run --stdin --tty`.

- [如何在 Rust 中实现基本的容器运行时 shim](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/) 并将其与 [实验性容器管理器](https://github.com/iximiuz/conman)。
- [如何实现 _attach_ 功能，带来交互式容器支持](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/) 并揭示 `docker run -i` 和 `docker run -i` 背后的魔力kubectl 运行 --stdin`。
- 如何释放 PTY 控制的容器的力量并阐明`docker run -it` 和`kubectl run --stdin --tty` 背后的技术。

Make code, not war!

编写代码，而不是战争！

### Related articles

###  相关文章

- [Implementing Container Runtime Shim: First Code](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/)
- [Implementing Container Runtime Shim: Interactive Containers](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/)
- [conman - [the] container manager: inception](http://iximiuz.com/en/posts/conman-the-container-manager-inception/) 

- [实现容器运行时 Shim：第一个代码](http://iximiuz.com/en/posts/implementing-container-runtime-shim-2/)
- [实现容器运行时 Shim：交互式容器](http://iximiuz.com/en/posts/implementing-container-runtime-shim-3/)
- [conman - [the] 容器管理器：inception](http://iximiuz.com/en/posts/conman-the-container-manager-inception/)

- [A journey from containerization to orchestration and beyond](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

- [从容器化到编排及其他的旅程](http://iximiuz.com/en/posts/journey-from-containerization-to-orchestration-and-beyond/)

### More Container insights from this blog

### 来自此博客的更多容器见解

- [Not every container has an operating system inside](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [You don't need an image to run a container](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [You need containers to build images](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [Container Networking Is Simple!](http://iximiuz.com/en/posts/container-networking-is-simple/)

- [并非每个容器内部都有操作系统](http://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
- [你不需要图像来运行容器](http://iximiuz.com/en/posts/you-dont-need-an-image-to-run-a-container/)
- [你需要容器来构建镜像](http://iximiuz.com/en/posts/you-need-containers-to-build-an-image/)
- [容器网络很简单！](http://iximiuz.com/en/posts/container-networking-is-simple/)

[docker,](javascript: void 0) [runc,](javascript: void 0) [OCI-runtime,](javascript: void 0) [containerd-shim,](javascript: void 0) [container](javascript: void 0)

[docker,](javascript: void 0) [runc,](javascript: void 0) [OCI-runtime,](javascript: void 0) [containerd-shim,](javascript: void 0) [container](javascript:无效 0)

#### Written by Ivan Velichko

#### 由伊万·维利奇科 (Ivan Velichko) 撰写

_Follow me on twitter [@iximiuz](https://twitter.com/iximiuz)_

_在推特上关注我 [@iximiuz](https://twitter.com/iximiuz)_

Liked this article? Let it be the beginning of a great friendship. Leave your email so I could notify you about new articles or any other interesting happenings around the topics of this blog. No spam whatsoever, I promise!

喜欢这篇文章吗？让它成为一段伟大友谊的开始。留下您的电子邮件，以便我可以通知您有关此博客主题的新文章或任何其他有趣的事件。没有任何垃圾邮件，我保证！

Copyright Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss) 

版权所有 Ivan Velichko © 2021 Feed [Atom](http://iximiuz.com/feed.atom)[RSS](http://iximiuz.com/feed.rss)

