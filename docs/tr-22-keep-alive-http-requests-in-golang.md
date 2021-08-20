# Keep-Alive in http requests in golang

16 Dec 2016 From https://awmanoj.github.io/tech/2016/12/16/keep-alive-http-requests-in-golang/

The [three-way handshake](https://en.wikipedia.org/wiki/Handshaking#TCP_three-way_handshake) in a TCP connection setup is heavy wrt performance and re-using the  already created connections is a considerable optimization. Keep alive  is a method to allow the same tcp connection for HTTP conversation  instead of opening a new one with each new request. Most servers and  clients allow for configurations and options to utilise this  optimisation.

TCP 连接设置中的 [三次握手](https://en.wikipedia.org/wiki/Handshaking#TCP_three-way_handshake) 对性能有很大影响，重用已经创建的连接是一个相当大的优化。保持活动是一种允许 HTTP 会话使用相同 tcp 连接的方法，而不是为每个新请求打开一个新连接。大多数服务器和客户端允许配置和选项来利用这种优化。

I had a strange encounter with the way this is handled in go.

我对 go 中处理这种情况的方式有过一次奇怪的遭遇。

I’ve a go application - an [nsq](http://nsq.io/) (a  distributed message queue) consumer - which reads some data off nsq and  send that data over a HTTP POST to a remote server over SSL. This  involves a [TLS handshake](https://www.ssl.com/article/ssl-tls-handshake-overview/) which is even more heavy than non-ssl handshake.

我有一个 go 应用程序——一个 [nsq](http://nsq.io/)（一个分布式消息队列）消费者——它从 nsq 读取一些数据，并通过 HTTP POST 将这些数据通过 SSL 发送到远程服务器。这涉及到 [TLS 握手](https://www.ssl.com/article/ssl-tls-handshake-overview/)，它比非 ssl 握手更重。

We found that nsq queue depth for the topics was high and increasing, additionally we kept seeing errors like in nsqd logs:

我们发现主题的 nsq 队列深度很高并且不断增加，此外我们不断看到类似 nsqd 日志中的错误：

```
[nsqd] 2016/12/16 17:48:34.563958 ERROR: [19.68.8.7:55239] - E_FIN_FAILED FIN 0b3caf88 failed ID not in flight - ID not in flight
 [nsqd] 2016/12/16 17:48:34.567675 ERROR: [19.68.8.7:55239] - E_FIN_FAILED FIN 0b3ca9f4 failed ID not in flight - ID not in flight
 [nsqd] 2016/12/16 17:48:34.571008 ERROR: [19.68.8.7:55239] - E_FIN_FAILED FIN 0b3c955c failed ID not in flight - ID not in flight
 [nsqd] 2016/12/16 17:48:34.573917 ERROR: [19.68.8.7:55239] - E_FIN_FAILED FIN 0b3c86f8 failed ID not in flight - ID not in flight
```


Based on the [limited](https://github.com/nsqio/nsq/issues/660) [number](https://github.com/nsqio/nsq/issues/729) [of documents](https: //github.com/nsqio/nsq/issues/762) we got over the internet it seemed like this was due to a slow consumer (hints like need for heartbeat from consumer pointed in this  direction). Since consumer was very simple - just doing some http posts  so it was the only suspect.

基于[有限](https://github.com/nsqio/nsq/issues/660) [数量](https://github.com/nsqio/nsq/issues/729) [文件](https: //github.com/nsqio/nsq/issues/762）我们通过互联网似乎这是由于消费者缓慢（暗示消费者需要心跳指向这个方向）。由于消费者非常简单 - 只是做一些 http 帖子，所以它是唯一的嫌疑人。

We did a `strace` on the running process to see the `read,connect` system calls. It indeed was doing some certificate interchange (read) for each connect or POST:

我们对正在运行的进程执行了 `strace` 以查看 `read,connect` 系统调用。它确实为每个连接或 POST 进行了一些证书交换（读取）：

```
[pid 18043] read(90, "g\201\f\1\2\0020|\6\10+\6\1\5\5\7\1\1\4p0n0$\6\10+\6 \1\5\5\0070\1\206\30http://ocsp.digicert.com0F\6\10+\6\1\5\5\0070\2\206:http://cacerts.digicert. com/DigiCertSHA2SecureServerCA.crt0\f\6\3U\35\23\1\1\377\4\0020\0000\r\6\t*\206H\206\367\r\1\1\v\5 \0\3\202\1\1\0\222\327\361\374\211T\330\363\305~8Z\270\210\313<Z\336o\274?\366\34\252\ 206Q\332\0044\257'\3751\205\273sa\200\251\325\224\267\\*\221\1/Ws\246\351bl=\330q?\200\256f\21\257\ 331\246\242w\313\202\245\361\241a?\2\345<a\35\313n\276\220\217k\377C\335}\235\2000+%?… ", 3166) = 1948
```


From the code, we noticed that:
- a http client was being created for each incoming message in the handler. this was overkill clearly.
- the http client created was using default transport. default http client does have a keep-alive setting.

从代码中，我们注意到：
- 处理程序为每个传入消息创建一个 http 客户端。这显然没有必要。
- 创建的 http 客户端使用默认 transport。默认的 http 客户端确实有一个 keep-alive 设置。

Since keep-alive is already there so we thought #1 is the problem  which is causing new connections to be created for every request which  in turn leads to the TLS handshake (and eventually to slow consumer). We fixed this by creating a global client:

由于 keep-alive 已经存在，所以我们认为 #1 是导致为每个请求创建新连接的问题，这反过来导致 TLS 握手（并最终减慢消费者）。我们通过创建一个全局客户端来解决这个问题：

```
  14 var client *http.Client
  15
  16 func init() {
  17     tr := &http.Transport{
  18         MaxIdleConnsPerHost: 1024,
  19         TLSHandshakeTimeout: 0 * time.Second,
  20     }
  21     client = &http.Client{Transport: tr}
  22 }
```


and using this to do the Post:

```
  31     resp, err := client.Post("https://api.some-web.com/v2/events", "application/json", bytes.NewBuffer(eventJson))
  32     if err != nil {
  33         log.Println("err", err)
  34         return err
  35     }
  36
  37     defer resp.Body.Close()
```


But we observed that problem persisted still (slowness as evident from errors in nsqd logs). We confirmed by doing `strace` again:

但是我们观察到问题仍然存在（从 nsqd 日志中的错误中可以看出速度很慢）。我们通过再次执行 `strace` 来确认：

```
$ sudo strace -s 2000 -f -p 18120 -e 'read,connect' 

[pid 18120] read(90, "g\201\f\1\2\0020|\6\10+\6\1\5\5\7\1\1\4p0n0$\6\10+\6 \1\5\5\0070\1\206\30http://ocsp.digicert.com0F\6\10+\6\1\5\5\0070\2\206:http://cacerts.digicert. com/DigiCertSHA2SecureServerCA.crt0\f\6\3U\35\23\1\1\377\4\0020\0000\r\6\t*\206H\206\367\r\1\1\v\5 \0\3\202\1\1\0\222\327\361\374\211T\330\363\305~8Z\270\210\313<Z\336o\274?\366\34\252\ 206Q\332\0044\257'\3751\205\273sa\200\251\325\224\267\\*\221\1/Ws\246\351bl=\330q?\200\256f\21\257\ 331\246\242w\313\202\245\361\241a?\2\345<a\35\313n\276\220\217k\377C\335}\235\2000+%?… ", 3166) = 1948
```

This appeared for each connection request. This meant that golang was somehow not honouring the keep-alive. This is when [this thread on stack overflow](http://stackoverflow.com/questions/17948827/reusing-http-connections-in-golang) helped us -

这出现在每个连接请求中。这意味着 golang 莫名不遵循 keep-alive。
> You should ensure that you read until the response is complete before calling Close().
> 在调用 Close() 之前，您应该确保在响应完成之前阅读。

```
res, _ := client.Do(req)
 io.Copy(ioutil.Discard, res.Body)
 res.Body.Close()
```

To ensure http.Client connection reuse be sure to do two things:
- Read until Response is complete (i.e. ioutil.ReadAll(rep.Body))
- Call Body.Close()

为确保 http.Client 连接重用，请务必做两件事：
- 阅读直到响应完成（即 ioutil.ReadAll(rep.Body)）
- 调用 Body.Close()

So we followed the suggestion:

所以我们遵循了这个建议：

```
resp, err := client.Post("https://api.some-web.com/v2/events", "application/json", bytes.NewBuffer(eventJson))
 if err != nil {
     log.Println("err", err)
     return defaultErrStatus, err
 }

 io.Copy(ioutil.Discard, resp.Body)   // <= NOTE

 defer resp.Body.Close()
```


While, as a I see the golang implementation, it does include [a subtle comment in the code](https://github.com/golang/go/blob/master/src/net/http/client.go#L666) but I don't understand yet why this idiosyncrasy exists. But would be something to be aware of and keep on the back of mind.

虽然，正如我看到的 golang 实现，它确实包含 [代码中的一个微妙的注释](https://github.com/golang/go/blob/master/src/net/http/client.go#L666)但我还不明白为什么这种特质存在。但将是需要注意并牢记的事情。

```
// Post issues a POST to the specified URL.
// Caller should close resp.Body when done reading from it.
// If the provided body is an io.Closer, it is closed after the request.
```


### Take Away

### 要点

- Use a global client as much as possible to avoid new connections.
- Even when you don’t need the response back, read it fully before closing the response body - to take benefit of persistent connections  (keep-alive). 
- 尽可能使用全局客户端以避免新连接。
- 即使您不需要返回响应，也请在关闭响应正文之前完整阅读它 - 以利用持久连接（保持活动）。

