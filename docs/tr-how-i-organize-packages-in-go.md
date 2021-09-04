# How I organize packages in Go

# 我如何在 Go 中组织包

Structuring the source code can be as challenging as writing it. There are many approaches to do so. Bad decisions can be painful and refactoring can be very time-consuming. On the other hand, it’s almost impossible to perfectly design your application at the beginning. What's more, some solutions may work at some application's size and should the application [develop over time](https://www.amazon.com/Building-Evolutionary-Architectures-Support-Constant-ebook/dp/B075RR1XVG/ref=sr_1_1?keywords=evolutionary+architecture&qid=1565498731&s=gateway&sr=8-1). Our software should grow with the problem it’s solving.

构建源代码与编写源代码一样具有挑战性。有很多方法可以做到这一点。错误的决定可能会很痛苦，而重构可能非常耗时。另一方面，在一开始就完美地设计您的应用程序几乎是不可能的。此外，一些解决方案可能适用于某些应用程序的大小，并且应用程序应该[随着时间的推移而开发](https://www.amazon.com/Building-Evolutionary-Architectures-Support-Constant-ebook/dp/B075RR1XVG/ref=sr_1_1?keywords=evolutionary+architecture&qid=1565498731&s=gateway&sr=8-1)。我们的软件应该随着它解决的问题而增长。

I mostly develop microservices and this architecture fits great to my needs. Projects with much more domain in it or more infrastructure applications may require a different approach. Let me know in the comments below what’s your design and where it makes sense the most.

我主要开发微服务，这种架构非常适合我的需求。具有更多领域或更多基础设施应用程序的项目可能需要不同的方法。在下面的评论中让我知道您的设计是什么以及它最有意义的地方。

## Packages and its dependencies

## 包及其依赖项

When it comes to developing domain services, it’s useful to split the service by domain’s components. Every component should be independent and, in theory, be able to be extracted to an external service if needed. What does it mean and how to achieve that?

在开发域服务时，将服务按域的组件拆分是很有用的。每个组件都应该是独立的，理论上，如果需要，可以提取到外部服务。它是什么意思以及如何实现它？

Let's assume that we have a service which handles everything related to placing orders like sending an email confirmation, saving information to a database, connecting with a payment provider etc. Every of the package should have a name which clearly [says what's for](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164) and compatible with [the naming standard](https://blog.golang.org/package-names).

假设我们有一个服务来处理与下订单相关的所有事情，比如发送电子邮件确认、将信息保存到数据库、与支付提供商连接等。每个包裹都应该有一个名称，清楚地[说明用途](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164)并与[命名标准](https://blog.golang.org/package-names)兼容。

![](http://developer20.com/images/organize-go.png)

This is only an example of a project where we have 3 packages: `confemails`, `payproviders` and `warehouse`. Names should be short and self-explaining.

这只是我们有 3 个包的项目示例：`confemails`、`payproviders` 和 `warehouse`。名称应该简短且不言自明。

Every of the package has his own `Setup()` function. The function accepts only bare requirements to be able to work correctly and be able to communicate with the world outside of the package. For example, if the package exposes an HTTP endpoint, the Setup() function accepts an HTTP router like mux Router. When you’re package requires access to the database then Setup() function accepts sql.DB as well. Of course, the package can need another package, too.

每个包都有自己的“Setup()”函数。该函数仅接受能够正常工作并能够与包外的世界进行通信的基本要求。例如，如果包公开了一个 HTTP 端点，则 Setup() 函数接受一个 HTTP 路由器，如 mux 路由器。当包需要访问数据库时，Setup() 函数也接受 sql.DB。当然，该包也可能需要另一个包。

## Inside of the package

## 包内

When we know the external dependencies of our module, we should focus on how to organize the code inside of it. At the very beginning, the package contains the following files:

- setup.go - where the Setup() function leaves
- service.go - it’s a place where the logic has its place
- repository.go - we need to fetch/save the information somewhere

当我们知道我们模块的外部依赖后，我们应该关注如何组织它内部的代码。一开始，该包包含以下文件：

- setup.go - Setup() 函数离开的地方
- service.go - 这是一个逻辑有其位置的地方
- repository.go - 我们需要在某处获取/保存信息

The `Setup()` function is responsible for building every building block of the module that is: services, repositories, registering event handlers or HTTP handlers and so on. This is an example of real production code which uses this approach.

`Setup()` 函数负责构建模块的每个构建块，即：服务、存储库、注册事件处理程序或 HTTP 处理程序等。这是使用这种方法的实际生产代码的示例。

```go
func Setup(router *mux.Router, httpClient httpGetter, auth jwtmiddleware.Authorization, logger logger) {
    h := httpHandler{
        logger:        logger,
        requestClaims: jwtutil.NewHTTPRequestClaims(client),
        service:       service{client: httpClient},
    }
    auth.CreateRoute("/v1/lastAnswerTime", h.proxyRequest, http.MethodGet)
}
```


As you can see, it builds a JWT middleware, a service which handles all the business logic (and where a logger is passed) and registers the HTTP handler. Thanks to that, the module is very independent and (in theory) can be moved out to a separate microservice without much work. And at the end, all of the packages are configured in the main function. 

如您所见，它构建了一个 JWT 中间件，一个处理所有业务逻辑（以及记录器传递的位置）并注册 HTTP 处理程序的服务。多亏了这一点，该模块非常独立，并且（理论上）可以移出到单独的微服务中，而无需做太多工作。最后，所有的包都在主函数中配置。

Sometimes, a few handlers or repositories are needed. For example, some information can be stored in a database and then sent via an event to a different part of your platform. Keeping only one repository with a method like saveToDb() isn’t that handy at all. All of elements like that are split by the functionality: repository\_order.go or service\_user.go. If there are more than 3 types of the object, there are moved to a separate subfolder.

有时，需要一些处理程序或存储库。例如，一些信息可以存储在数据库中，然后通过事件发送到平台的不同部分。使用 saveToDb() 之类的方法只保留一个存储库一点也不方便。所有类似的元素都按功能拆分：repository\_order.go 或 service\_user.go。如果有超过 3 种类型的对象，则将它们移动到单独的子文件夹中。

![](http://developer20.com/images/organizing-go-1.png)

## Testing

## 测试

When it comes to testing, I stick to a few rules. Firstly, use interfaces in the Setup() function. Those interfaces should be as small as possible. In the example above, there’s httpGetter interface. The interface has only `Get()` function in it.

在测试方面，我遵守一些规则。首先，在 Setup() 函数中使用接口。这些接口应该尽可能小。在上面的例子中，有 httpGetter 接口。该接口中只有“Get()”函数。

```go
type httpGetter interface {
 Get(url string) (resp *http.Response, err error)
}
```


Thank’s to that, I only have to mock only 1 method. The interface is always as close to its usage as possible.

多亏了这一点，我只需要模拟 1 个方法。界面总是尽可能接近其用途。

Secondly, try to write fewer tests which will cover more code at the same time. There’s no sense to write a test for every repository or service separately. For every domain decision/operation, one successful and one failed test should be sufficient and cover about 80% of the code. Sometimes, there is some critical part of the application. Then, this part can be covered by separate test cases.

其次，尝试编写更少的测试，同时覆盖更多的代码。单独为每个存储库或服务编写测试是没有意义的。对于每个域决策/操作，一次成功和一次失败的测试就足够了，并且覆盖了大约 80% 的代码。有时，应用程序有一些关键部分。然后，这部分可以由单独的测试用例覆盖。

Finally, write tests in separate package suffixed with `_test` and put it inside of the module. It will help to keep everything in one place.

最后，在以`_test` 为后缀的单独包中编写测试并将其放入模块中。这将有助于将所有内容放在一个地方。

When you want to test the whole application, prepare every dependency in the `setup()` function next to the main function. It’ll give you the same setup for both production and test environments that can save you some bugs. Tests should reuse the setup() function and mock only those dependencies which aren’t easy to mock (like external APIs).

当您要测试整个应用程序时，请在 main 函数旁边的 `setup()` 函数中准备每个依赖项。它将为生产和测试环境提供相同的设置，可以为您节省一些错误。测试应该重用 setup() 函数并只模拟那些不容易模拟的依赖项（如外部 API）。

## Summary

##  概括

All the rest files like .travis.yaml etc are kept in the project root. It gives me a clear view of the whole project. I know where to look for the domain files, where infrastructure-related files are and there aren’t mixed. Otherwise, the main folder of the project would become a mess.

.travis.yaml 等所有其他文件都保存在项目根目录中。它让我对整个项目有一个清晰的认识。我知道在哪里寻找域文件，与基础设施相关的文件在哪里，没有混合。否则，项目的主文件夹会变得一团糟。

As I said in the introduction, I know that all of the projects won’t benefit from this way of organizing project but smaller applications like microservices can find it very useful.


正如我在介绍中所说，我知道所有项目都不会从这种组织项目的方式中受益，但像微服务这样的小型应用程序会发现它非常有用。


### See Also

###  也可以看看

- [Golang Tips & Tricks #6 - the \_test package](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - blank identifier in structs](http://developer20.com/golang-tips-and-trics-v/)
- [GoGoConf 2019 - report](http://developer20.com/gogoconf-2019/)
- [Golang Tips & Tricks #4 - internal folders](http://developer20.com/golang-tips-and-trics-iv/)
- [Golang Tips & Tricks #3 - graceful shutdown](http://developer20.com/golang-tips-and-trics-iii/)

- [Golang Tips & Tricks #6 - \_test 包](http://developer20.com/golang-tips-and-trics-vi/)
- [Golang Tips & Tricks #5 - 结构体中的空白标识符](http://developer20.com/golang-tips-and-trics-v/)
- [GoGoConf 2019 - 报告](http://developer20.com/gogoconf-2019/)
- [Golang Tips & Tricks #4 - 内部文件夹](http://developer20.com/golang-tips-and-trics-iv/)
- [Golang Tips & Tricks #3 - 优雅关机](http://developer20.com/golang-tips-and-trics-iii/)

[←](http://developer20.com/golang-tips-and-trics-vi/)[→](http://developer20.com/golang-tips-and-trics-vii/)Top

[←](http://developer20.com/golang-tips-and-trics-vi/)[→](http://developer20.com/golang-tips-and-trics-vii/)顶部

© 2021 . Made with [Hugo](https://gohugo.io) using the [Tale](https://github.com/EmielH/tale-hugo/) theme. 

© 2021。 [Hugo](https://gohugo.io) 使用 [Tale](https://github.com/EmielH/tale-hugo/) 主题制作。

