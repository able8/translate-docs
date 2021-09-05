# Gracefully shutting down a Nodejs HTTP server

# 优雅地关闭 Nodejs HTTP 服务器

25 Nov 2015

Looking to gracefully shutdown your Nodejs HTTP server? Well,  it's actually a little more difficult than you think. Without handling  process signals, like "ctrl-c", your application terminates immediately. Which means, there may be requests still processing that you terminated before completion. In a live environment, this is a horrible user  experience. Unfortunately, even if you did catch the process' signals,  Node doesn't actually have a built in mechanism for gracefully shutting  down a running HTTP server.

想要优雅地关闭你的 Nodejs HTTP 服务器？嗯，它实际上比你想象的要困难一些。如果不处理进程信号，例如“ctrl-c”，您的应用程序将立即终止。这意味着，可能有您在完成前终止的请求仍在处理中。在实时环境中，这是一种可怕的用户体验。不幸的是，即使您确实捕获了进程的信号，Node 实际上也没有内置机制来优雅地关闭正在运行的 HTTP 服务器。

## Defining the problem

## 定义问题

Graceful shutdown is the process by which a running HTTP server:

1. Stops accepting new connections
2. Stops fulfilling new requests
3. Waits for in-flight requests to complete

正常关闭是运行 HTTP 服务器的过程：

1. 停止接受新连接
2. 停止满足新的请求
3. 等待飞行中的请求完成

Once these three prerequisites are fufilled the HTTP server is  "shutdown". All the while, it has not interrupted any concurrent request while transitioning to this state.

一旦满足这三个先决条件，HTTP 服务器就会“关闭”。一直以来，它在转换到此状态时都没有中断任何并发请求。

## Node's naive solution

## Node 的天真解决方案

Let's take the classic HTTP server written with Node's native HTTP library:

```javascript
var http = require('http');

var server = http.createServer(function(req, res) {
    res.end('Goodbye!');
});

server.listen(3000);
```


This example, while trivial, demonstrates a working HTTP server. Now, what if I wanted to terminate it, but not disrupt any concurrent  traffic? How about the `close` method?

这个例子虽然微不足道，但演示了一个工作的 HTTP 服务器。现在，如果我想终止它，但又不想中断任何并发流量怎么办？ `close` 方法怎么样？

```javascript
server.close(function() {
    console.log('We closed!');
    process.exit();
});
```


Unfortunately, not really, there's a problem with this. Reading the documentation might make it a little more clear:

> Stops the server from accepting new connections and keeps existing  connections. This function is asynchronous, the server is finally closed when all connections are ended and the server emits a 'close' event. -  From [Nodejs Docs](https://nodejs.org/api/net.html#net_server_close_callback)

不幸的是，并非如此，这有问题。阅读文档可能会使它更清楚一点：

> 停止服务器接受新连接并保持现有连接。此函数是异步的，当所有连接都结束并且服务器发出“关闭”事件时，服务器最终关闭。 - 来自 [Nodejs 文档](https://nodejs.org/api/net.html#net_server_close_callback)

The important part of the description to remember is "keeps existing connections". While `close` does stop the listening socket from accepting new ones, sockets that  are already connected may continue to operate - which is fine if they're mid-request - but this also includes sockets that are connected with a  'keep-alive' connection type.

要记住的描述的重要部分是“保持现有连接”。虽然 `close` 确实会阻止侦听套接字接受新的套接字，但已经连接的套接字可能会继续运行 - 如果它们处于中间请求，这很好 - 但这也包括与“保持活动”连接的套接字连接类型。

> HTTP persistent connection, also called HTTP keep-alive, or HTTP  connection reuse, is the idea of using a single TCP connection to send  and receive multiple HTTP requests/responses, as opposed to opening a  new connection for every single request/response pair - From [Wikipedia](https://en.wikipedia.org/wiki/HTTP_persistent_connection)

> HTTP 持久连接，也称为 HTTP keep-alive，或 HTTP 连接重用，是使用单个 TCP 连接发送和接收多个 HTTP 请求/响应的想法，而不是为每个单个请求/响应对打开一个新连接- 来自 [维基百科](https://en.wikipedia.org/wiki/HTTP_persistent_connection)

This means that sockets that are kept alive will remain alive and  still capable of making additional HTTP requests. This is obviously not  what we want. While it fulfils points one and three of our definition of graceful shutdown, it does not fulfill number two and is thus, not a  complete solution. A graceful shutdown should programatically:

1. Close the listening socket to prevent new connections
2. Close all idle keep-alive sockets to prevent new requests during shutdown
3. Wait for all in-flight requests to finish before closing their sockets.

这意味着保持活动的套接字将保持活动状态并且仍然能够发出额外的 HTTP 请求。这显然不是我们想要的。虽然它满足了我们对优雅关机的定义的第一点和第三点，但它不满足第二点，因此不是一个完整的解决方案。正常关闭应该以编程方式：

1. 关闭监听socket，防止新连接
2. 关闭所有空闲的keep-alive sockets，防止shutdown时有新的请求
3. 在关闭套接字之前等待所有进行中的请求完成。

## The solution

##  解决方案

[http-shutdown](https://github.com/thedillonb/http-shutdown) is a [NPM package](https://www.npmjs.com/package/http-shutdown) that provides graceful shutdown functionality. It does this by  leveraging several mechanisms for keeping track of active vs idle  sockets. The principle is as follows:

1. Listen for socket "connection" and "close" events. When a socket  connects, add it to a list of connected sockets and mark that socket as  idle (it hasn't made a request yet). When it's closed, remove it from  that list. This gives us a mechanism to track all sockets that are  currently connected and allows us to keep track of which of those are  active vs idle.
2. Listen on the HTTP server's "request" and "finish" event. When the  "request" event triggers, mark the underlying socket as active. When the "finish" event occurs, mark the socket as idle. In addition, look to  see if the system is "shutting down". If so, call `destroy` on that socket to close it, making sure no additional traffic flows through it.

[http-shutdown](https://github.com/thedillonb/http-shutdown) 是一个 [NPM 包](https://www.npmjs.com/package/http-shutdown)，提供优雅的关闭功能。它通过利用多种机制来跟踪活动和空闲套接字来做到这一点。原理如下：

1. 监听套接字“连接”和“关闭”事件。当套接字连接时，将其添加到已连接套接字列表中并将该套接字标记为空闲（它还没有发出请求）。当它关闭时，将其从该列表中删除。这为我们提供了一种机制来跟踪当前连接的所有套接字，并允许我们跟踪哪些套接字是活动的还是空闲的。
2. 监听 HTTP 服务器的“请求”和“完成”事件。当“请求”事件触发时，将底层套接字标记为活动的。当“完成”事件发生时，将套接字标记为空闲。此外，查看系统是否正在“关闭”。如果是这样，请在该套接字上调用 `destroy` 以关闭它，确保没有额外的流量通过它。

Given the two mechanisms above, we can easily implement a `shutdown` method that will active a graceful shutdown. It only need to do the following: 

鉴于上述两种机制，我们可以轻松实现一个“关闭”方法，该方法将激活正常关闭。它只需要执行以下操作：

1. Call the native `close` method on the HTTP server to close the listening socket and prevent all future socket connections.
2. Set a flag that indicates the server is shutting down. (Principle  two above uses this flag to close sockets when they're done with their  request)
3. Loop through all currently connected sockets where they're idle flag is true: call `destory` on these sockets to close them out since they're not serving traffic and we want to prevent them from serving any more.

1. 在HTTP服务器上调用原生`close`方法关闭监听socket，阻止以后所有的socket连接。

2. 设置一个标志，指示服务器正在关闭。 （上面的原则二使用此标志在完成请求后关闭套接字）
3. 循环所有当前连接的套接字，它们空闲标志为真：在这些套接字上调用 `destory` 将它们关闭，因为它们不提供流量，我们希望阻止它们再提供服务。

If you're interested in the source code behind the [NPM library](https://www.npmjs.com/package/http-shutdown), check out the [GitHub repository](https://github.com/thedillonb/http-shutdown). Or, more specifically, check out the [single file](https://github.com/thedillonb/http-shutdown/blob/master/index.js) which contains all of which I discussed above. 

如果您对 [NPM 库](https://www.npmjs.com/package/http-shutdown) 背后的源代码感兴趣，请查看 [GitHub 存储库](https://github.com/thedillonb/http-关闭)。或者，更具体地说，查看 [单个文件](https://github.com/thedillonb/http-shutdown/blob/master/index.js)，其中包含我上面讨论的所有内容。

