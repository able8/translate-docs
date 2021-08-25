# The awesomeness of the httptest package in Go

# Go 中 httptest 包的厉害之处

25 Feb 2020 ·   Five minute read From: https://gianarb.it/blog/golang-mockmania-httptest

Go has a nice http package. I am able to say that because I am not aware of any other implementation of it in Go other than the one provided by the standard library. This is for me a good sign.

Go 有一个很好的 http 包。我能这么说是因为除了标准库提供的实现之外，我不知道它在 Go 中的任何其他实现。这对我来说是个好兆头。

```
resp, err := http.Get("http://example.com/")
 if err != nil {
 // handle error
 }
 defer resp.Body.Close()
 body, err := ioutil.ReadAll(resp.Body)
```


This example comes from the [documentation](https://golang.org/pkg/net/http/) itself.

此示例来自 [文档](https://golang.org/pkg/net/http/) 本身。

We are here to read about testing, so who cares about the http package itself! What matters is the [httptest](https://golang.org/pkg/net/http/httptest/) package! Way cooler.

我们在这里阅读有关测试的信息，所以谁在乎 http 包本身！重要的是 [httptest](https://golang.org/pkg/net/http/httptest/) 包！凉快多了。

This article is not the first one for the MockMania series, I wrote about titled [“InfluxDB Client v2”](https://gianarb.it/blog/golang-mockmania-influxdb-v2-client), it uses the httptest service already! But hey it deserves its own blog post.

这篇文章不是 MockMania 系列的第一篇，我写的标题是 [“InfluxDB Client v2”](https://gianarb.it/blog/golang-mockmania-influxdb-v2-client)，它使用了 httptest 服务已经！但是，嘿，它值得拥有自己的博客文章。

## Server Side

## 服务器端

The http package provides a client and a server. The server is made of handlers. The handler takes a request and based on that it returns a response. This is its interface:

http 包提供了一个客户端和一个服务器。服务器由处理程序组成。处理程序接受一个请求，并根据它返回一个响应。这是它的界面：

```
type Handler interface {
     ServeHTTP(ResponseWriter, *Request)
 }
```


As you can see if gets a ResponseWriter to compose a response based on the Request it gets. This process can be as complicated as you like, it can reaches databases, third party services but in the end, it writes a response.

如您所见，是否获取 ResponseWriter 以根据它获取的请求撰写响应。这个过程可以随心所欲地复杂，它可以到达数据库、第三方服务，但最终它会写一个响应。

It means that mocking all the dependencies to get the right scenario we use the ResponseWriter to figure out if the handler made what we want.

这意味着模拟所有依赖项以获得正确的场景，我们使用 ResponseWriter 来确定处理程序是否做了我们想要的。

The httptest package provides a replacement for the ResponseWriter called ResponseRecorder. We can pass it to the handler and check how it looks like after its execution:

httptest 包提供了 ResponseWriter 的替代品，称为 ResponseRecorder。我们可以将它传递给处理程序并检查它执行后的样子：

```
 handler := func(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, "ping")
 }

 req := httptest.NewRequest("GET", "http://example.com/foo", nil)
 w := httptest.NewRecorder()
 handler(w, req)

 resp := w.Result()
 body, _ := ioutil.ReadAll(resp.Body)

 fmt.Println(resp.StatusCode)
 fmt.Println(string(body))
```


This handler is very simple, it just manipulates the response body. If your handler is more complicated and it has dependencies you have to be sure to replace them as well, injecting the appropriate one.

这个处理程序非常简单，它只是操作响应体。如果您的处理程序更复杂并且它有依赖项，您必须确保也替换它们，注入适当的一个。

## Client-Side

## 客户端

Handlers are useful if you can’t use them. The Go http package provides an http client as well that you can use to interact with an http server. An http client by itself is useless, but it is the entry point for all the manipulation and transformation you do on the information you get via HTTP. With the proliferation of microservices, it is a very common situation.

如果您不能使用处理程序，处理程序很有用。 Go http 包还提供了一个 http 客户端，您可以使用它与 http 服务器进行交互。 http 客户端本身是无用的，但它是您对通过 HTTP 获得的信息进行的所有操作和转换的入口点。随着微服务的激增，这是一种非常普遍的情况。

The workflow is well understood, you have an HTTP backend to interact with, you fetch data from there are you manipulate them with your business logic. When testing what you can do is to mock the http backend in order to return what you want, testing that your business logic does what it is supposed to do based on the input you get from the HTTP server.

工作流程很好理解，你有一个 HTTP 后端与之交互，你从那里获取数据，然后用你的业务逻辑来操作它们。在测试时，您可以做的是模拟 http 后端以返回您想要的内容，测试您的业务逻辑是否根据您从 HTTP 服务器获得的输入执行它应该执行的操作。

During our first example, the handler was the subject of our testing, this is not the case anymore, we are testing the consumer this time, so we have to mimic and handler in order to get what we expect to return

在我们的第一个例子中，处理程序是我们测试的主题，现在不是这样了，这次我们正在测试消费者，所以我们必须模仿和处理程序以获得我们期望返回的内容

```
 ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintln(w, "I am a super server")
 }))
 defer ts.Close()
```


As you can see we are creating a new HTTP server via the httptest. It accepts a handler. The goal for this handler returns what we would like to gest our code on. In theory, it should just use the ResponseWriter to compose the response we expect.

如您所见，我们正在通过 httptest 创建一个新的 HTTP 服务器。它接受一个处理程序。这个处理程序的目标是返回我们想要获取代码的内容。理论上，它应该只使用 ResponseWriter 来编写我们期望的响应。

The server has a bunch of methods, the one you are looking for is the URL one. Because we can pass it to an http.Client, the one we will use as a mock for our function

服务器有很多方法，您要查找的方法是 URL 方法。因为我们可以将它传递给一个 http.Client，我们将使用它作为我们函数的模拟

```
 res, err := http.Get(ts.URL)
 if err != nil {
  log.Fatal(err)
 }
 bb, err := ioutil.ReadAll(res.Body)
 res.Body.Close()
```


That’s it, as you can see `ts.URL` points the http.Client to the mock server we created.

就是这样，你可以看到 `ts.URL` 将 http.Client 指向我们创建的模拟服务器。

## Conclusion

## 结论

I use the httptest package a lot even when writing SDKs for services that do not have integration with Go because I can follow their documentation mocking their server and I do not need to reach them until I am confident with the code I wrote.

即使在为没有与 Go 集成的服务编写 SDK 时，我也经常使用 httptest 包，因为我可以按照他们的文档模拟他们的服务器，并且在我对我编写的代码充满信心之前不需要联系他们。

My suggestion is to test your client code for edge cases as well because of the httptest.Server gives you the flexibility to write any response you can think about. You can mimic an authorized response to seeing how your code with handle it, or an empty body or a rate limit. The only limit is our laziness. 

我的建议是测试客户端代码的边缘情况，因为 httptest.Server 使您可以灵活地编写您能想到的任何响应。您可以模仿授权响应来查看您的代码如何处理它，或者一个空的正文或一个速率限制。唯一的限制是我们的懒惰。
