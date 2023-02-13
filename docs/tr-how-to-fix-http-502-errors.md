# A Solution to HTTP 502 Errors with AWS ALB

# 使用 AWS ALB 解决 HTTP 502 错误

https://www.tessian.com/blog/how-to-fix-http-502-errors/

https://www.tessian.com/blog/how-to-fix-http-502-errors/

by Samson DanzigerFriday, October 1st, 2021

作者：Samson Danziger 2021 年 10 月 1 日，星期五

At Tessian, we have many applications that interact with each other using [REST APIs](https://www.redhat.com/en/topics/api/what-is-a-rest-api). We noticed in the logs that at random times, uncorrelated with traffic, and seemingly unrelated to any code we had actually written, we were getting a lot of HTTP 502 “Bad Gateway” errors.

在 Tessian，我们有许多使用 [REST API](https://www.redhat.com/en/topics/api/what-is-a-rest-api) 相互交互的应用程序。我们在日志中注意到，在随机时间，与流量无关，并且看似与我们实际编写的任何代码无关，我们会收到大量 HTTP 502“Bad Gateway”错误。

Now that the issue is fixed, I wanted to explain what this error means, how you get it and how to solve it. My hope is that if you’re having to solve this same issue, this article will explain why and what to do.

既然问题已经解决，我想解释一下这个错误的含义、你是如何得到它的以及如何解决它。我希望，如果您必须解决同样的问题，本文将解释原因和解决方法。

## First, let’s talk about load balancing

## 首先来说说负载均衡

![](https://www.tessian.com/wp-content/uploads/2020/11/image1.png)

An example of a development system where clients communicate directly with the server.

客户端直接与服务器通信的开发系统示例。

In a development system, you usually run one instance of a server and you communicate directly with it. You send HTTP requests to it, it returns responses, everything is golden.

在开发系统中，您通常运行一个服务器实例并直接与其通信。您向它发送 HTTP 请求，它返回响应，一切都很好。

For a production system running at any non-trivial scale, this doesn’t work. Why? Because the amount of traffic going to the server is much greater, and you need it to not fall over even if there are tens of thousands of users.

对于以任何非平凡规模运行的生产系统，这是行不通的。为什么？因为到服务器的流量要大得多，你需要它即使有几万用户也不会倒下。

Typically, servers have a maximum number of connections they can support. If it goes over this number, new people can’t connect, and you have to wait until a new connection is freed up. In the old days, the solution might have been to have a bigger machine, with more resources, and more available connections.

通常，服务器有它们可以支持的最大连接数。如果超过这个数字，新人将无法连接，您必须等待新连接被释放。在过去，解决方案可能是拥有更大的机器、更多资源和更多可用连接。

Now we use a load balancer to manage connections from the client to multiple instances of the server. The load balancer sits in the middle and routes client requests to any available server that can handle them in a pool.

现在我们使用负载均衡器来管理从客户端到服务器的多个实例的连接。负载均衡器位于中间并将客户端请求路由到池中可以处理它们的任何可用服务器。

If one server goes down, traffic is automatically routed to one of the others in the pool. If a new server is added, traffic is automatically routed to that, too. This all happens to reduce load on the others.

如果一台服务器出现故障，流量会自动路由到池中的另一台服务器。如果添加了新服务器，流量也会自动路由到该服务器。这一切恰好是为了减轻其他人的负担。

![](https://www.tessian.com/wp-content/uploads/2020/11/image2.png)

An example of a production system that uses a load balancer to manage connections from the client to multiple instances of the server.

使用负载平衡器管理从客户端到服务器的多个实例的连接的生产系统示例。

## What are 502 errors?

## 什么是 502 错误？

On the web, there are a variety of HTTP status codes that are sent in response to requests to let the user know what happened. Some might be pretty familiar:

在 web 上，有各种各样的 HTTP 状态代码是为了响应请求而发送的，让用户知道发生了什么。有些人可能很熟悉：

- **200 OK** – Everything is fine.
- **301 Moved Permanently** – I don’t have what you’re looking for, try here instead.
- **403 Forbidden** – I understand what you’re looking for, but you’re not allowed here.
- **404 Not Found** – I can’t find whatever you’re looking for.
- **503 Service Unavailable** – I can’t handle the request right now, probably too busy.

- **200 OK** – 一切都很好。
- **301 永久移动** – 我没有您要找的东西，请在此处尝试。
- **403 Forbidden** – 我知道你在找什么，但你不被允许在这里。
- **404 Not Found** – 我找不到你要找的东西。
- **503 服务不可用** – 我现在无法处理请求，可能是太忙了。

4xx and 5xx both deal with errors. 4xx are for client errors, where the user has done something wrong. 5xx, on the other hand, are server errors, where something is wrong on the server and it’s not your fault.

4xx 和 5xx 都处理错误。 4xx 用于客户端错误，即用户做错了什么。另一方面，5xx 是服务器错误，服务器出现问题，这不是您的错。

All of these are specified by a standard called [RFC7231](https://tools.ietf.org/html/rfc7231#section-6.6.3). For 502 it says:

所有这些都由名为 [RFC7231](https://tools.ietf.org/html/rfc7231#section-6.6.3) 的标准指定。对于 502，它说：

_The 502 (Bad Gateway) status code indicates that the server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while attempting to fulfill the request._

_502（错误网关）状态代码表示服务器在充当网关或代理时，在尝试完成请求时从其访问的入站服务器收到了无效响应。_

The load balancer sits in the middle, between the client and the actual service you want to talk to. Usually it acts as a dutiful messenger passing requests and responses back and forth. But, if the service returns an invalid or malformed response, instead of returning that nonsensical information to the client, it sends back a 502 error instead.

负载均衡器位于中间，介于客户端和您要与之通信的实际服务之间。通常它充当尽职尽责的信使来回传递请求和响应。但是，如果服务返回无效或格式错误的响应，它不会将无意义的信息返回给客户端，而是发回 502 错误。

This lets the client know that the response the load balancer received was invalid.

这让客户端知道负载均衡器收到的响应无效。

![](https://www.tessian.com/wp-content/uploads/2020/11/image6.png)

High level overview of requests between the client, the ALB, and the server Created with Lucidchart.

使用 Lucidchart 创建的客户端、ALB 和服务器之间请求的高级概述。

## The actual issue 

## 实际问题

Adam Crowder has done [a full analysis of this problem](https://adamcrowder.net/posts/node-express-api-and-aws-alb-502/) by tracking it _all_ the way down to TCP packet capture to assess what's going wrong. That’s a bit out of scope for this post, but here’s a brief summary of what’s happening:

Adam Crowder 已经完成[对这个问题的全面分析](https://adamcrowder.net/posts/node-express-api-and-aws-alb-502/) 通过跟踪它_all_一直到 TCP 数据包捕获到评估出了什么问题。这有点超出本文的范围，但这里是对正在发生的事情的简要总结：

1.  At Tessian, we have lots of interconnected services. Some of them have Application Load Balancers (ALBs) managing the connections to them.
2.  In order to make an HTTP request, we must open a TCP socket from the client to the server. Opening a socket involves performing a three-way handshake with the server before either side can send any data.
3.  Once we’ve finished sending data, the socket is closed with a 4 step process. These 3 and 4 step processes can be a large overhead when not much actual data is sent.
4.  Instead of opening and then closing one socket per HTTP request, we can keep a socket open for longer and reuse it for multiple HTTP requests. **This is called HTTP Keep-Alive.** Either the client _or_ the server can then initiate a close of the socket with a FIN segment (either for fun or due to timeout).

1. 在 Tessian，我们有很多相互关联的服务。其中一些具有应用程序负载平衡器 (ALB) 来管理与它们的连接。
2. 为了发出 HTTP 请求，我们必须打开从客户端到服务器的 TCP 套接字。打开套接字涉及在任何一方发送任何数据之前与服务器执行三次握手。
3. 一旦我们发送完数据，套接字将通过 4 个步骤关闭。当发送的实际数据不多时，这些 3 步和 4 步过程可能会产生很大的开销。
4. 我们可以将一个套接字保持打开更长时间，并将其重新用于多个 HTTP 请求，而不是针对每个 HTTP 请求打开然后关闭一个套接字。 **这称为 HTTP Keep-Alive。** 客户端 _ 或 _ 服务器然后可以使用 FIN 段启动套接字的关闭（为了好玩或由于超时）。

![](https://www.tessian.com/wp-content/uploads/2020/11/Screenshot-2020-11-03-at-16.14.13-e1604420186830.png)

Source: An Introduction to Computer Networks, edition 1.9.21

资料来源：计算机网络简介，1.9.21 版

The 502 Bad Gateway error is caused when the ALB sends a request to a service at the same time that the service closes the connection by sending the FIN segment to the ALB socket. The ALB socket receives FIN, acknowledges, and starts a new handshake procedure.

502 Bad Gateway 错误是在 ALB 向服务发送请求的同时服务通过将 FIN 段发送到 ALB 套接字来关闭连接而引起的。 ALB 套接字接收 FIN、确认并启动新的握手过程。

Meanwhile, the socket on the service side has just received a data request referencing the previous (now closed) connection. Because it can’t handle it, it sends an RST segment _back_ to the ALB, and then the ALB returns a 502 to the user.

同时，服务端的套接字刚刚收到引用先前（现已关闭）连接的数据请求。因为它无法处理，它向 ALB _back_ 发送一个 RST 段，然后 ALB 返回一个 502 给用户。

The diagram and table below show what happens between sockets of the ALB and the Server.

下图和下表显示了 ALB 的套接字和服务器之间发生的情况。

![](https://www.tessian.com/wp-content/uploads/2020/11/image5.png)

Socket communication segments between the Application Load Balancer and the Server. You can imagine the Client on the left.

Application Load Balancer 和服务器之间的套接字通信段。你可以想象左边的客户。

![](https://www.tessian.com/wp-content/uploads/2020/11/Screenshot-2020-11-03-at-16.11.14.png)

Table of segments sent between the Application Load Balancer and the Server.

在 Application Load Balancer 和服务器之间发送的分段表。

## How to fix 502 errors

## 如何修复 502 错误

It’s fairly simple. Just make sure that the service doesn’t send the FIN segment before the ALB sends a FIN segment to the service. In other words, make sure the service doesn’t close the HTTP Keep-Alive connection before the ALB.

这很简单。只需确保在 ALB 向服务发送 FIN 段之前服务不发送 FIN 段。换句话说，确保服务不会在 ALB 之前关闭 HTTP Keep-Alive 连接。

[The default timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#connection-idle-timeout) for the AWS Application Load Balancer is 60 seconds, so we changed the service timeouts to 65 seconds. Barring two hiccups shortly after deploying, this has totally fixed it.

AWS Application Load Balancer 的[默认超时](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#connection-idle-timeout) 是 60 秒，所以我们将服务超时更改为 65 秒。除非在部署后不久出现两次小问题，否则这已经完全解决了。

### The actual configuration change

### 实际配置更改

I have included the configuration for common Python and Node server frameworks below. If you are using any of those, you can just copy and paste. If not, these should at least point you in the right direction.

我在下面包含了常见 Python 和 Node 服务器框架的配置。如果您正在使用其中任何一个，则只需复制并粘贴即可。如果没有，这些至少应该为您指明正确的方向。

### uWSGI (Python)

### uWSGI（Python）

As a config file:

作为配置文件：

```
# app.ini
[uwsgi]
...
harakiri = 65
add-header = Connection: Keep-Alive
http-keepalive = 1
...
```

Or as command line arguments:

或者作为命令行参数：

```
--add-header "Connection: Keep-Alive"
--http-keepalive --harakiri 65
```

### Gunicorn (Python)

### 独角兽（Python）

As command line arguments:

作为命令行参数：

```
--keep-alive 65
```

### Express (Node)

### 快递（节点）

In Express, specify the time in milliseconds on the server object.

在 Express 中，在服务器对象上指定时间（以毫秒为单位）。

```
const express = require('express');
const app = express();
const server = app.listen(80);

server.keepAliveTimeout = 65000
```

Looking for more tips from engineers and other cybersecurity news? Keep up with [our blog](https://www.tessian.com/blog/) and follow us on [LinkedIn](https://www.linkedin.com/company/tessian/).

寻找来自工程师和其他网络安全新闻的更多提示？关注 [我们的博客](https://www.tessian.com/blog/) 并在 [LinkedIn](https://www.linkedin.com/company/tessian/) 上关注我们。

Learn More About Life at Tessian](https://www.tessian.com/blog/category/team-culture/)

了解更多关于 Tessian 的生活](https://www.tessian.com/blog/category/team-culture/)

Samson Danziger
Python Backend Engineer

萨姆森·丹兹格
Python后端工程师

Subscribe to our blog 

订阅我们的博客

