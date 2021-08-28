# Structuring and testing HTTP handlers in Go

# 在 Go 中构建和测试 HTTP 处理程序

Written 15th of October 2020 From: https://www.maragu.dk/blog/structuring-and-testing-http-handlers-in-go/

There are many ways to structure your HTTP handlers in your web application code in Go. It would be nice to have a default way to do this that makes it easy to:
- Inject your dependencies, to make the handlers and the rest of your code loosely coupled
- See which route paths go to which handlers, and have them close together in code, for readability
- Unit test the handlers in isolation

在 Go 的 Web 应用程序代码中构建 HTTP 处理程序的方法有很多种。最好有一种默认方式来执行此操作，以便轻松：
- 注入您的依赖项，使处理程序和其余代码松散耦合
- 查看哪些路由路径到达哪些处理程序，并在代码中将它们放在一起，以提高可读性
- 单独对处理程序进行单元测试

After quite a few different designs, I've found a way I like, and in this post, I'll show you.

经过多次不同的设计，我找到了一种我喜欢的方式，在这篇文章中，我将向您展示。

If you want to check out a simple project implementing this, see[github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing).

如果你想查看一个实现这个的简单项目，请参阅[github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing)。

## The handler

## 处理程序

I'll start by showing you the design, and then breaking it down. A handler generally looks like this:

我将首先向您展示设计，然后将其分解。处理程序通常如下所示：

```go
package handlers

import (
    "net/http"

    "github.com/go-chi/chi"
)

type partyStarterRepo interface {
    StartParty(id string) error
}

func PartyStarter(mux chi.Router, s partyStarterRepo) {
    mux.Post("/start/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
        id := chi.URLParam(r, "id")
        if err := s.StartParty(id);err != nil {
            http.Error(w, err.Error(), http.StatusBadGateway)
            return
        }
        w.WriteHeader(http.StatusAccepted)
    })
}
```


### Request multiplexer as parameter

### 请求多路复用器作为参数

The `PartyStarter` function takes the request multiplexer `mux` (in this case [chi](https://github.com/go-chi/chi),but use any you like) as the first parameter. This means that the handler registers itself, including defining the route and its parameters. It's nice to have this very close to the handler code, both for increased readability, but also that it's very clear that they belong together and should be changed together. For example, if the `id` parameter changes in name or content, the code right below should reflect that.

`PartyStarter` 函数将请求多路复用器 `mux`（在本例中为 [chi](https://github.com/go-chi/chi)，但可以使用任何你喜欢的）作为第一个参数。这意味着处理程序会注册自己，包括定义路由及其参数。很高兴将此代码与处理程序代码非常接近，既提高了可读性，又很明显它们属于一起并且应该一起更改。例如，如果 `id` 参数的名称或内容发生变化，下面的代码应该反映这一点。

### Dependency as private interface parameter

### 依赖作为私有接口参数

The business logic dependency is passed as an interface, `partyStarterRepo`, that is defined specifically for this handler. We can do this in Go because interfaces are implicit, meaning that anything that has a method with signature `StartParty(id string) error` can be passed to this function. We will use this in testing.

业务逻辑依赖作为一个接口`partyStarterRepo` 传递，该接口是专门为此处理程序定义的。我们可以在 Go 中做到这一点，因为接口是隐式的，这意味着任何具有签名为“StartParty(id string) error”的方法都可以传递给这个函数。我们将在测试中使用它。

This enables us to define exactly what this handler needs from its dependencies, nothing more, nothing less. So if your dependency has a lot of extra functionality (for example, a`StopParty` function), this handler doesn't know about it.

这使我们能够准确定义该处理程序从其依赖项中需要什么，仅此而已。因此，如果您的依赖项具有许多额外功能（例如，`StopParty` 函数），则该处理程序不知道它。

### Handlers in a separate package

### 处理程序在单独的包中

To isolate the handlers, they are in a separate package called `handlers`. Note that because of the use of private interfaces for dependencies, we don't import our business logic packages in the handlers. This reduces coupling, and makes it easier to swap the underlying dependencies, for example.

为了隔离处理程序，它们位于一个名为“处理程序”的单独包中。请注意，由于对依赖项使用了私有接口，因此我们不会在处理程序中导入我们的业务逻辑包。例如，这减少了耦合，并使交换底层依赖关系变得更容易。

## Testing the handler

## 测试处理程序

To test the handler, we can use the `httptest` package from the standard library, along with a very small mock for the dependency.

为了测试处理程序，我们可以使用标准库中的 `httptest` 包，以及一个非常小的依赖模拟。

```go
package handlers

import (
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/go-chi/chi"
)

type partyStarterRepoMock struct {
    err error
}

func (m *partyStarterRepoMock) StartParty(id string) error {
    return m.err
}

func TestPartyStarter(t *testing.T) {
    t.Run("sends bad gateway on start party error", func(t *testing.T) {
        mux := chi.NewMux()
        PartyStarter(mux, &partyStarterRepoMock{err: errors.New("no snacks")})

        req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
        rec := httptest.NewRecorder()
        mux.ServeHTTP(rec, req)

        res := rec.Result()
        if res.StatusCode != http.StatusBadGateway {
            t.FailNow()
        }
    })

    t.Run("sends accepted on start party success", func(t *testing.T) {
        mux := chi.NewMux()
        PartyStarter(mux, &partyStarterRepoMock{})

        req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
        rec := httptest.NewRecorder()
        mux.ServeHTTP(rec, req)

        res := rec.Result()
        if res.StatusCode != http.StatusAccepted {
            t.FailNow()
        }
    })
}
```


See how the mock is tiny, because we're testing only exactly what this handler needs? No more autogenerating huge mocks with all your business logic functions on it.

看看模拟是多么小，因为我们只测试这个处理程序需要什么？不再自动生成包含所有业务逻辑功能的大型模拟。

Also note that we don't have to start our server to check that our routes work as expected, because the routes are right there in the handler.

另请注意，我们不必启动服务器来检查我们的路由是否按预期工作，因为路由就在处理程序中。

## Conclusion 

##  结论

In this post, I've shown you how to structure your HTTP handlers in Go so they are loosely coupled with their dependencies, using private interfaces, and easy to test, using routes that are defined inside the handler. To see a simple project showing you all of this, check out [github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing).

在这篇文章中，我向您展示了如何在 Go 中构建您的 HTTP 处理程序，以便它们与它们的依赖项松散耦合，使用私有接口，并且易于测试，使用在处理程序中定义的路由。要查看一个向您展示所有这些的简单项目，请查看 [github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing)。

If you have any questions or comments about this, feel free to reach out to me on Twitter. I'm at [@markusrgw](https://twitter.com/markusrgw).

如果您对此有任何问题或意见，请随时在 Twitter 上与我联系。我在 [@markusrgw](https://twitter.com/markusrgw)。

## About me

##  关于我

I'm Markus, a professional software consultant and developer. 🤓✨ You can reach me at [markus@maragu.dk](mailto:Markus from maragu) or [@markusrgw](https://twitter.com/markusrgw).

我是 Markus，一位专业的软件顾问和开发人员。 🤓✨ 你可以通过 [markus@maragu.dk](mailto:Markus from maragu) 或 [@markusrgw](https://twitter.com/markusrgw) 联系我。

I'm currently [building Go courses over at golang.dk](https://www.golang.dk/). 

我目前正在 [在 golang.dk 上构建 Go 课程](https://www.golang.dk/)。

