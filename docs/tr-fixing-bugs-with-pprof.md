## Fixing Service Performance with PProf in Go

## 在 Go 中使用 PProf 修复服务性能

Mar 1, 2019

2019 年 3 月 1 日

I tweeted the other day about how I managed to reduce the CPU  consumption on one of my services by 90% by removing a single line of  code.

前几天我发了一条推文，说我如何通过删除一行代码设法将我的一个服务的 CPU 消耗减少了 90%。

![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/tweet.png)

One of the things I did not do, mainly since it is difficult to make a full explanation in 256 characters is explain how I identified the  problem and the process I took to locate the root cause. This short post sets out to rectify that, and I hope it serves as a useful resource to  save you all from repeating my embarrassing mistake.

我没有做的一件事，主要是因为很难在 256 个字符中做出完整的解释是解释我如何识别问题以及我找到根本原因的过程。这篇短文旨在纠正这一点，我希望它是一种有用的资源，可以避免大家重复我的尴尬错误。

## The Service

##  服务

The service I was working on was a simple API aggregation service; this exposes a public HTTP API consumed from a ReactJS website. It  interacts with two other upstream services, a cache which uses gRPC for  the transport and a face detection service which is using Matt Ryer and  David Hernandez [FaceBox](https://machinebox.io/docs/facebox).

我正在开发的服务是一个简单的 API 聚合服务；这公开了从 ReactJS 网站使用的公共 HTTP API。它与其他两个上游服务交互，一个使用 gRPC 进行传输的缓存和一个使用 Matt Ryer 和 David Hernandez [FaceBox](https://machinebox.io/docs/facebox) 的人脸检测服务。

I make no apologies for the implementation details of this service; it is not the model on an excellent microservice in fact there are  components which really should be delegated out into other services. What I was building was a simple system which would allow me to  demonstrate how to use the Consul Connect service mesh and Envoy’s  reliability and observability features. If you would like to take a look at the source code, you can find the link on GitHub: https://github.com/emojify-app/api.

对于这项服务的实施细节，我不道歉；它不是一个优秀的微服务模型，事实上，有些组件确实应该委托给其他服务。我正在构建的是一个简单的系统，它允许我演示如何使用 Consul Connect 服务网格以及 Envoy 的可靠性和可观察性功能。如果您想查看源代码，可以在 GitHub 上找到链接：https://github.com/emojify-app/api。

## The Problem

##  问题

The service itself was functioning fine, the latency was low, and  there were no errors, but the CPU consumption did feel a little high for the traffic received and the work the service is doing. Running this on my Kubernetes cluster it was easy to miss this, but when running this  in an environment with lower resources it was a problem as the CPU  consumption was starving the other services. This unusual behavior  caused me to start to take a look and investigate.

服务本身运行良好，延迟很低，并且没有错误，但是对于接收到的流量和服务正在执行的工作，CPU 消耗确实有点高。在我的 Kubernetes 集群上运行它很容易错过这一点，但是当在资源较少的环境中运行它时，这是一个问题，因为 CPU 消耗正在耗尽其他服务。这种不寻常的行为让我开始审视和调查。

**API Service CPU Consumption** ![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/before_cpu.png)

The above chart is showing the CPU consumption from the service; this might be nothing normal; looking at something like CPU without context  is not the best way to draw a conclusion. What started to look unusual  was when I combined the CPU chart, an understanding on the actual work  the service was doing at the time which was streaming files over a gRPC  connection from another service, and the limited number of requests.

上图显示了服务的 CPU 消耗；这可能不正常；在没有上下文的情况下查看 CPU 之类的东西并不是得出结论的最佳方式。开始看起来不寻常的是，当我结合 CPU 图表时，对服务当时正在做的实际工作的理解，即通过来自另一个服务的 gRPC 连接流式传输文件，以及有限数量的请求。

**API Service Requests per Second** ![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/before_rps.png)

Something looks fishy here I have a hunch that there is something not quite right as I have built many Go based services and generally they  are terrifically efficient on their CPU and memory consumption. Let’s  take a look at the upstream that the API is calling and see how that is  performing as this service is responsible for sending the files.

这里看起来有些可疑，我有一种预感，因为我已经构建了许多基于 Go 的服务，而且通常它们在 CPU 和内存消耗方面的效率非常高。让我们看看 API 正在调用的上游，看看它是如何执行的，因为该服务负责发送文件。

**Cache Service CPU** ![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/cache_cpu.png)

There is a dramatic difference there; the Cache is using 10% of the  CPU of the API for the same number of requests. This Cache service is  reading from a file and sending the bytes of data as a gRPC message. The API service receives that message and writes it as an HTTP response. There is nothing complicated going on there; there should not be such a  difference in the two services.

那里有一个巨大的差异。对于相同数量的请求，缓存使用 API 的 CPU 的 10%。此缓存服务正在读取文件并将数据字节作为 gRPC 消息发送。 API 服务接收该消息并将其写入为 HTTP 响应。没有什么复杂的事情发生。这两种服务不应该有这样的差异。

## Investigating the Problem 

## 调查问题

I was now pretty confident that there was a bug somewhere in the API  service which was causing it to consume way too much CPU and I needed to investigate. Luckily for me Go has an excellent tool called pprof https://golang.org/pkg/net/http/pprof/ which allows you to inspect the internal working of your application,  you can see incredible detail like timings for memory allocation and the execution time for individual blocks of code. Adding this to your code  is also incredibly easy, so I decided, I would deploy a new build of my  service with the diagnostics included so that I could run a profile.

我现在非常确信 API 服务中的某个地方存在一个错误，导致它消耗了过多的 CPU，我需要进行调查。幸运的是，Go 有一个出色的工具 pprof https://golang.org/pkg/net/http/pprof/ 它允许您检查应用程序的内部工作，您可以看到令人难以置信的细节，例如内存分配的时间和单个代码块的执行时间。将它添加到您的代码中也非常容易，所以我决定，我将部署一个包含诊断的新服务版本，以便我可以运行配置文件。

To enable profiling, I only had to add a couple of lines of code; the first was to import and enable the pprof package.

为了启用分析，我只需要添加几行代码；第一个是导入并启用 pprof 包。

```
import    _ "net/http/pprof"
```

If you already have a web server in your application, then pprof automatically attaches itself to `http.DefaultServeMux` enables the API at the path `/debug/pprof/`

如果您的应用程序中已经有一个 web 服务器，那么 pprof 会自动将自身附加到 `http.DefaultServeMux` 中，在路径 `/debug/pprof/` 启用 API

I was not using the DefaultServeMux in my application as I am using  the Gorilla Mux package for my http handlers. Because I was using  Gorilla, I had to add another line of code to enable the HTTP routing to pprof

我没有在我的应用程序中使用 DefaultServeMux，因为我将 Gorilla Mux 包用于我的 http 处理程序。因为我使用的是 Gorilla，所以我不得不添加另一行代码来启用到 pprof 的 HTTP 路由

```
r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
```

After building and re-deploying my application with the  instrumentation enabled I could then sample the running processes using  the pprof tool. Again this is a straightforward process; the pprof has  excellent documentation on how to collect the different profiles, CPU,  heap, blocking, etc.

在启用检测的情况下构建和重新部署我的应用程序后，我可以使用 pprof 工具对正在运行的进程进行采样。同样，这是一个简单的过程； pprof 有关于如何收集不同配置文件、CPU、堆、阻塞等的优秀文档。

I point the pprof tool at my profile endpoint which is: `https://myservice/debug/pprof/profile?seconds=5` this collects a five-second profile of the CPU. The short profile is  fine for my requirements as I know the requests are completing quickly, I don’t know why they are consuming so much CPU.

我将 pprof 工具指向我的配置文件端点，即：`https://myservice/debug/pprof/profile?seconds=5` 这会收集 CPU 的五秒配置文件。简短的配置文件很适合我的要求，因为我知道请求正在快速完成，我不知道为什么它们消耗这么多 CPU。

```
$ go tool pprof "https://myservice/debug/pprof/profile?seconds=5"
Fetching profile over HTTP from https://myservice/debug/pprof/profile?seconds=5
Saved profile in /home/jacksonnic/pprof/pprof.emojify-api.samples.cpu.006.pb.gz
File: emojify-api
Type: cpu
Time: Mar 1, 2019 at 5:09am (UTC)
Duration: 5s, Total samples = 40ms (  0.8%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

Once I had the pprof tool running the first thing I like to do is to  view a visual call trace of the collected profile, this allows me to  zoom in quickly to the source of the problem. You need to have Graphviz  installed to do this, but by merely executing the command pdf pprof  outputs a graphical overview of the profile.

一旦我让 pprof 工具运行，我喜欢做的第一件事就是查看收集的配置文件的可视调用跟踪，这使我能够快速放大到问题的根源。您需要安装 Graphviz 才能执行此操作，但只需执行命令 pdf pprof 即可输出配置文件的图形概览。

The output looked like this:

输出如下所示：

![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile_before_1.png)

Immediately I can see that `writeString` in the protobuf package is consuming a considerable amount of the CPU. Next step is to trace this back up to the source, in my code to  understand why.

我立即可以看到 protobuf 包中的 `writeString` 正在消耗大量 CPU。下一步是在我的代码中将其追溯到源头以了解原因。

Following the trace, I finally get to code I have written and saw this:

在跟踪之后，我终于看到了我编写的代码并看到了这一点：

![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile_before_2.png)

The source of the problem seems to be a fmt.Println statement, time  to dig into the code and see what that is doing and why it is even  there.

问题的根源似乎是一个 fmt.Println 语句，是时候深入研究代码并看看它在做什么以及为什么它甚至在那里。

```
d, err := c.cache.Get(context.Background(), &wrappers.StringValue{Value: f})
fmt.Println("err", d, err)
```

So it seems that when I was debugging my application, I was writing  the response from the gRPC cache service which is a protobuf message to  StdOut. Like a good developer, I forgot this was there and then deployed the application. Serializing this protobuf message in this way was  incredibly expensive, annoyingly it was also very unnecessary.

因此，似乎在调试应用程序时，我正在将 gRPC 缓存服务的响应写入 StdOut，这是一个 protobuf 消息。就像一个优秀的开发人员，我忘记了它在那里，然后部署了应用程序。以这种方式序列化这个 protobuf 消息非常昂贵，令人讨厌的是，这也是非常不必要的。

## Fixing the Service

## 修复服务

The fix could not be more straightforward, delete the line of code,  rebuild and deploy the service again, with the new service deployed  there was an immediate impact to the CPU consumption.

修复再简单不过了，删除代码行，重新构建并重新部署服务，部署新服务后，CPU 消耗立即受到影响。

**CPU post fix** ![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/after_cpu.png)

Over a 90% reduction in CPU, and what looks like far more normal  operating conditions for the service with this traffic. It is good to  double check; however, so again I ran a profile, this time the profile  showed the hot spot as syscall. The result this time is what I would  expect as the service is mainly reading and writing to a TCP socket.

CPU 减少了 90% 以上，并且对于具有此流量的服务来说，看起来更正常的操作条件。仔细检查一下是很好的；然而，我再次运行了一个配置文件，这次配置文件将热点显示为系统调用。这次的结果是我所期望的，因为该服务主要是读取和写入 TCP 套接字。

**CPU profile after fix** ![img](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile_after.png)

## Summary

##  概括

I always say no experience is bad if you can learn something from it, I certainly learned that I am prone to making stupid mistakes, but it  was also fun to dig into pprof again. The whole process of finding the  problem and fixing it took me approximately 30 minutes, this, of course, could have been so much longer had the issue not been so pronounced. However, it highlights just how amazing the tooling in the Go ecosystem  is.

我总是说，如果你能从中学到一些东西，没有经验是不好的，我当然知道我很容易犯愚蠢的错误，但再次深入研究 pprof 也很有趣。发现问题并修复它的整个过程花了我大约 30 分钟，当然，如果问题不那么明显，这个过程可能会更长。然而，它强调了 Go 生态系统中的工具是多么的神奇。

If you want a takeaway, I have two:

如果你想要外卖，我有两个：

- Code review, a fresh pair of eyes would probably spot the unnecessary `Println` statement
- Profile your services before major deployments, it does not take  long, and a quick eyeball of the results can save embarrassing mistakes

- 代码审查，一双新的眼睛可能会发现不必要的`Println`语句
- 在主要部署之前分析您的服务，不需要很长时间，快速观察结果可以避免令人尴尬的错误

If you would like to see more detail on the before and after traces, you can download a PDF from the following links:

如果您想查看有关前后轨迹的更多详细信息，可以从以下链接下载 PDF：

- Before: [profile001.pdf](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile001.pdf)
- After: [profile002.pdf](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile002.pdf)

- 之前：[profile001.pdf](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile001.pdf)
- 之后：[profile002.pdf](https://nicholasjackson.io/images/posts/fixing-bugs-with-pprof/profile002.pdf)

One final thing, remember to remove the instrumentation from your  service, especially if this is running in production. I may or may not  have just done this.

最后一件事，请记住从您的服务中删除检测，尤其是在生产中运行时。我可能刚刚也可能没有这样做。

Have fun,

玩得开心，

Nic

网卡

Posted with : [Golang](https://nicholasjackson.io/tag/go/)

发表于：[Golang](https://nicholasjackson.io/tag/go/)

------

Nic Jackson's Picture

尼克杰克逊的照片

#### Nic Jackson

#### 尼克杰克逊

Author, Geek, Ultrarunner and in my spare time developer advocate at HashiCorp. Author of [Building Microservices in Go](https://goo.gl/vUKtPT). Maintainer of Minke https://minke.rocks, build and test tool for microservices. 

作者、极客、Ultrarunner 和我业余时间在 HashiCorp 的开发人员倡导者。 [Building Microservices in Go](https://goo.gl/vUKtPT) 的作者。 Minke 的维护者 https://minke.rocks，微服务的构建和测试工具。

