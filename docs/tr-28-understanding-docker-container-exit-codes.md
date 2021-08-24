# Understanding Docker Container Exit Codes
## The most common exit codes, what they mean, and what causes them

## 了解 Docker 容器退出代码
### 最常见的退出代码，它们的含义，以及导致它们的原因

[Oct 21, 2019](https://betterprogramming.pub/understanding-docker-container-exit-codes-5ee79a1d58f6?source=post_page-----5ee79a1d58f6--------------------------------) · 4 min read

Docker container exit code — how to use them for troubleshooting?
It’s one of the most common question that I come across: “Why is my container not running?”
Can docker container exit codes help troubleshoot this issue?

Docker 容器退出代码——如何使用它们进行故障排除？
这是我遇到的最常见的问题之一：“为什么我的容器没有运行？”
docker 容器退出代码可以帮助解决这个问题吗？

The first step in answering this question is to identify the exit code for  the docker container. The exit code may give a hint as to what happened  to stop the container running. This article lists the most common exit  codes when working with docker containers and aims to answer two  important questions:
- What does this specific exit code mean?
- What action caused this exit code?

回答这个问题的第一步是确定 docker 容器的退出代码。退出代码可能会提示停止容器运行的原因。本文列出了使用 docker 容器时最常见的退出代码，旨在回答两个重要问题：
- 这个特定的退出代码是什么意思？
- 什么操作导致了这个退出代码？

This will ultimately help answer the original question: “Why is my container not running?”

这最终将有助于回答最初的问题：“为什么我的容器没有运行？”

# How to Find Exit Codes

# 如何查找退出代码

## Option 1: List all containers that exited

## 选项 1：列出所有退出的容器

```
docker ps --filter "status=exited"
```


## Option 2: Grep by container name

## 选项 2：按容器名称 Grep

```
docker ps -a grep <container-name>Example: docker ps -a | grep hello-world
```


## Option 3: Inspect by container id

## 选项 3：通过容器 ID 检查

```
docker inspect <container-id> --format='{{.State.ExitCode}}'Example: docker inspect ca6cbb290468 --format='{{.State.ExitCode}}'
```


# Exit Codes

# 退出代码

Common exit codes associated with docker containers are:

- **Exit Code 0**: Absence of an attached foreground process
- **Exit Code 1**: Indicates failure due to application error
- **Exit Code 137**: Indicates failure as container received SIGKILL (Manual intervention or ‘oom-killer’ [OUT-OF-MEMORY])
- **Exit Code 139**: Indicates failure as container received SIGSEGV
- **Exit Code 143**: Indicates failure as container received SIGTERM

与 docker 容器相关的常见退出代码是：
- **退出代码 0**：没有附加的前台进程
- **退出代码 1**：表示由于应用程序错误而失败
- **退出代码 137**：表示容器收到 SIGKILL 失败（手动干预或“oom-killer”[OUT-OF-MEMORY]）
- **退出代码 139**：表示容器收到 SIGSEGV 失败
- **退出代码 143**：表示容器收到 SIGTERM 失败

## Exit Code 0

## 退出代码 0

- Exit code 0 indicates that the specific container does not have a foreground process attached.
- This exit code is the exception to all the other exit codes to follow. It does not necessarily mean something bad happened.
- Developers use this exit code if they want to automatically stop their container once it has completed its job.

- 退出代码 0 表示特定容器没有附加前台进程。
- 此退出代码是要遵循的所有其他退出代码的例外。这并不一定意味着发生了不好的事情。
- 如果开发人员希望在容器完成其工作后自动停止容器，则使用此退出代码。

Here is an example using the public docker container — “hello-world”. If you have docker installed on your system or VM instance, run this:

这是一个使用公共 docker 容器的示例——“hello-world”。如果您的系统或 VM 实例上安装了 docker，请运行以下命令：

```
docker run hello-world
```


You will get a message, “Hello from docker!” but try to find the container using this code:

您将收到一条消息，“来自 docker 的您好！” 但尝试使用此代码查找容器：

```
docker ps -a | grep hello-world
```


You’ll notice that the container exited and the exit code is 0. This is  because the container does not have any foreground process attached,  such as a Java process or a shell process that runs until a SIGTERM  event occurs

你会注意到容器退出了，退出代码为 0。这是因为容器没有附加任何前台进程，例如 Java 进程或运行直到 SIGTERM 事件发生的 shell 进程

![img](https://miro.medium.com/max/60/1*onZujADMgDc3KvsDt8k0BA.png?q=20)

![img](https://miro.medium.com/max/1400/1*onZujADMgDc3KvsDt8k0BA.png)

Exit code 0

退出代码 0

## Exit Code 1

## 退出代码 1

- Indicates that the container stopped due to either an application error or an  incorrect reference in Dockerfile to a file that is not present in the  container.
- An application error can be as simple as “divide by 0” or as complex as  “Reference to a bean name that conflicts with existing, non-compatible  bean definition of same name and class.”
- An incorrect reference in Dockerfile to a file not present in the container can be as simple as a typo (the example below has `sample.ja` instead of `sample.jar`)

- 表示容器由于应用程序错误或 Dockerfile 中对容器中不存在的文件的错误引用而停止。
- 应用程序错误可以像“除以 0”一样简单，也可以像“引用与现有的、不兼容的相同名称和类的 bean 定义冲突的 bean 名称”一样复杂。
- Dockerfile 中对容器中不存在的文件的不正确引用可能就像拼写错误一样简单（下面的示例使用 `sample.ja` 而不是 `sample.jar`）

![img](https://miro.medium.com/max/60/1*0C1noFKfmrj5FVnhpbO53w.png?q=20)

![img](https://miro.medium.com/max/1400/1*0C1noFKfmrj5FVnhpbO53w.png)

## Exit Code 137

## 退出代码 137

- This indicates that container received SIGKILL
- A common event that initiates a SIGKILL is a docker kill. This can be initiated either manually by user or by the docker daemon:

- 这表明容器收到了 SIGKILL
- 启动 SIGKILL 的常见事件是 docker kill。这可以由用户或 docker 守护进程手动启动：

```
docker kill <container-id>
```


- `docker kill` can be initiated manually by the user or by the host machine. If  initiated by host machine, then it is generally due to being out of  memory. To confirm if the container exited due to being out of memory,  verify `docker inspect` against the container id for the section below and check if `OOMKilled` is true (which would indicate it is out of memory):

- `docker kill` 可以由用户或主机手动启动。如果是宿主机发起的，那么一般是因为内存不足。要确认容器是否因内存不足而退出，请根据以下部分的容器 ID 验证 `docker inspect`，并检查 `OOMKilled` 是否为 true（这表示它内存不足）：

```
"State": {
  "Status": "exited",
  "Running": false,
  "Paused": false,
  "Restarting": false,
  "OOMKilled": true,
  "Dead": false,
  "Pid": 0,
  "ExitCode": 137,
  "Error": "",
  "StartedAt": "2019-10-21T01:13:51.7340288Z",
  "FinishedAt": "2019-10-21T01:13:51.7961614Z"
 }
```


## Exit Code 139

## 退出代码 139

- This indicates that container received SIGSEGV 
- 这表明容器收到了 SIGSEGV
- SIGSEGV indicates a segmentation fault. This occurs when a program attempts to access a [memory](https://en.wikipedia.org/wiki/Computer_memory) location that it’s not allowed to access, or attempts to access a memory location in a way that’s not allowed.
- From the Docker container standpoint, this either indicates an issue with  the application code or sometimes an issue with the base images used by  the container.

- SIGSEGV 表示分段错误。当程序尝试访问不允许访问的 [memory](https://en.wikipedia.org/wiki/Computer_memory) 位置，或尝试以不允许的方式访问内存位置时，就会发生这种情况。
- 从 Docker 容器的角度来看，这要么表明应用程序代码存在问题，要么有时表明容器使用的基础镜像存在问题。

## Exit Code 143

## 退出代码 143

- This indicates that container received SIGTERM.
- Common events that initiate a SIGTERM are `docker stop` or `docker-compose stop`. In this case there was a manual termination that forced the container to exit:

- 这表明容器收到了 SIGTERM。
- 启动 SIGTERM 的常见事件是 `docker stop` 或 `docker-compose stop`。在这种情况下，有一个手动终止迫使容器退出：

```
docker stop <container-id>
 OR
docker-compose down <container-id>
```


- **Note:** sometimes `docker stop` can also result in exit code 137. This typically happens if the  application tied to the container doesn’t handle SIGTERM — the docker  daemon waits ten seconds then issues SIGKILL

- **注意：** 有时`docker stop` 也会导致退出代码 137。如果绑定到容器的应用程序不处理 SIGTERM，通常会发生这种情况——docker 守护进程等待 10 秒然后发出 SIGKILL

### **Some uncommon exit codes with Docker containers (typically with shell script usage)**

### **Docker 容器的一些不常见退出代码（通常使用 shell 脚本）**

- **Exit Code 126**: Permission problem or command is not executable
- **Exit Code 127**: Possible typos in shell script with unrecognizable characters

- **退出代码 126**：权限问题或命令不可执行
- **退出代码 127**：shell 脚本中可能存在无法识别字符的拼写错误

[Better Programming](https://betterprogramming.pub/?source=post_sidebar--------------------------post_sidebar-----------)


