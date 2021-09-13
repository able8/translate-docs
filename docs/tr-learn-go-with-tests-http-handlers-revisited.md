# HTTP Handlers Revisited

# 重新审视 HTTP 处理程序

**[You can find all the code here](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/http-handlers-revisited)**

**[你可以在这里找到所有代码](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/http-handlers-revisited)**

This book already has a chapter on [testing a HTTP handler](http-server.md) but this will feature a broader discussion on designing them, so they are simple to test.

本书已有一章关于[测试 HTTP 处理程序](http-server.md)，但这将更广泛地讨论设计它们，因此它们易于测试。

We'll take a look at a real example and how we can improve how it's designed by applying principles such as single responsibility principle and separation of concerns. These principles can be realised by using [interfaces](structs-methods-and-interfaces.md) and [dependency injection](dependency-injection.md). By doing this we'll show how testing handlers is actually quite trivial.

我们将看一个真实的例子，以及我们如何通过应用诸如单一职责原则和关注点分离等原则来改进它的设计方式。这些原理可以通过使用[接口](structs-methods-and-interfaces.md)和[依赖注入](dependency-injection.md)来实现。通过这样做，我们将展示测试处理程序实际上是多么微不足道。

![Common question in Go community illustrated](amazing-art.png)

Testing HTTP handlers seems to be a recurring question in the Go community, and I think it points to a wider problem of people misunderstanding how to design them.

测试 HTTP 处理程序似乎是 Go 社区中反复出现的问题，我认为它指向了人们误解如何设计它们的更广泛问题。

So often people's difficulties with testing stems from the design of their code rather than the actual writing of tests. As I stress so often in this book:

人们在测试方面的困难往往源于他们的代码设计，而不是实际的测试编写。正如我在本书中经常强调的那样：

> If your tests are causing you pain, listen to that signal and think about the design of your code.

> 如果您的测试让您感到痛苦，请倾听该信号并考虑您的代码设计。

## An example

##  一个例子

[Santosh Kumar tweeted me](https://twitter.com/sntshk/status/1255559003339284481)

[Santosh Kumar 给我发了推文](https://twitter.com/sntshk/status/1255559003339284481)

> How do I test a http handler which has mongodb dependency?

> 如何测试具有 mongodb 依赖性的 http 处理程序？

Here is the code

这是代码

```go
func Registration(w http.ResponseWriter, r *http.Request) {
    var res model.ResponseResult
    var user model.User

    w.Header().Set("Content-Type", "application/json")

    jsonDecoder := json.NewDecoder(r.Body)
    jsonDecoder.DisallowUnknownFields()
    defer r.Body.Close()

    // check if there is proper json body or error
    if err := jsonDecoder.Decode(&user);err != nil {
        res.Error = err.Error()
        // return 400 status codes
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(res)
        return
    }

    // Connect to mongodb
    client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err := client.Connect(ctx)
    if err != nil {
        panic(err)
    }
    defer client.Disconnect(ctx)
    // Check if username already exists in users datastore, if so, 400
    // else insert user right away
    collection := client.Database("test").Collection("users")
    filter := bson.D{{"username", user.Username}}
    var foundUser model.User
    err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
    if foundUser.Username == user.Username {
        res.Error = UserExists
        // return 400 status codes
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(res)
        return
    }

    pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        res.Error = err.Error()
        // return 400 status codes
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(res)
        return
    }
    user.Password = string(pass)

    insertResult, err := collection.InsertOne(context.TODO(), user)
    if err != nil {
        res.Error = err.Error()
        // return 400 status codes
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(res)
        return
    }

    // return 200
    w.WriteHeader(http.StatusOK)
    res.Result = fmt.Sprintf("%s: %s", UserCreated, insertResult.InsertedID)
    json.NewEncoder(w).Encode(res)
    return
}
```

Let's just list all the things this one function has to do:

让我们列出这个函数必须做的所有事情：

1. Write HTTP responses, send headers, status codes, etc.
2. Decode the request's body into a `User`
3. Connect to a database (and all the details around that)
4. Query the database and applying some business logic depending on the result
5. Generate a password
6. Insert a record

1. 编写 HTTP 响应、发送标头、状态码等。
2. 将请求的主体解码为`User`
3. 连接到数据库（以及所有相关细节）
4.查询数据库并根据结果应用一些业务逻辑
5. 生成密码
6.插入一条记录

This is too much.

这太多了。

## What is a HTTP Handler and what should it do ?

## 什么是 HTTP 处理程序，它应该做什么？

Forgetting specific Go details for a moment, no matter what language I've worked in what has always served me well is thinking about the [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns) and the [ single responsibility principle](https://en.wikipedia.org/wiki/Single-responsibility_principle).

暂时忘记了特定的 Go 细节，无论我用什么语言工作过，一直对我有用的都是在考虑 [关注点分离](https://en.wikipedia.org/wiki/Separation_of_concerns) 和 [单一责任原则](https://en.wikipedia.org/wiki/Single-responsibility_principle)。

This can be quite tricky to apply depending on the problem you're solving. What exactly _is_ a responsibility?

根据您要解决的问题，这可能很难应用。究竟什么是责任？

The lines can blur depending on how abstractly you're thinking and sometimes your first guess might not be right.

线条可能会模糊，具体取决于您思考的抽象程度，有时您的第一个猜测可能不正确。

Thankfully with HTTP handlers I feel like I have a pretty good idea what they should do, no matter what project I've worked on:

值得庆幸的是，对于 HTTP 处理程序，无论我从事什么项目，我都非常清楚它们应该做什么：

1. Accept a HTTP request, parse and validate it.
2. Call some `ServiceThing` to do `ImportantBusinessLogic` with the data I got from step 1. 

1. 接受一个 HTTP 请求，解析并验证它。
2. 调用一些 `ServiceThing` 来使用我从第 1 步得到的数据做 `ImportantBusinessLogic`。

3. Send an appropriate `HTTP` response depending on what `ServiceThing` returns.

3. 根据 `ServiceThing` 返回的内容发送适当的 `HTTP` 响应。

I'm not saying every HTTP handler _ever_ should have roughly this shape, but 99 times out of 100 that seems to be the case for me.

我并不是说每个 HTTP 处理程序 _ever_ 都应该大致具有这种形状，但是 100 次中有 99 次对我来说似乎是这种情况。

When you separate these concerns:

  当您将这些关注点分开时：

 - Testing handlers becomes a breeze and is focused a small number of concerns.
  - Importantly testing `ImportantBusinessLogic` no longer has to concern itself with `HTTP`, you can test the business logic cleanly.
  - You can use `ImportantBusinessLogic` in other contexts without having to modify it.
  - If `ImportantBusinessLogic` changes what it does, so long as the interface remains the same you don't have to change your handlers.

- 测试处理程序变得轻而易举，并且只关注少数问题。
 - 重要的是测试 `ImportantBusinessLogic` 不再需要关注 `HTTP`，你可以干净地测试业务逻辑。
 - 您可以在其他上下文中使用“ImportantBusinessLogic”而无需修改它。
 - 如果`ImportantBusinessLogic` 改变了它的作用，只要界面保持不变，你就不必改变你的处理程序。

## Go's Handlers

## Go 的处理程序

[`http.HandlerFunc`](https://golang.org/pkg/net/http/#HandlerFunc)

[`http.HandlerFunc`](https://golang.org/pkg/net/http/#HandlerFunc)

> The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.

> HandlerFunc 类型是一个适配器，允许将普通函数用作 HTTP 处理程序。

`type HandlerFunc func(ResponseWriter, *Request)`

`type HandlerFunc func(ResponseWriter, *Request)`

Reader, take a breath and look at the code above. What do you notice?

读者，深呼吸，看看上面的代码。你注意到什么？

**It is a function that takes some arguments**

**这是一个接受一些参数的函数**

There's no framework magic, no annotations, no magic beans, nothing.

没有框架魔法，没有注释，没有魔法豆，什么都没有。

It's just a function, _and we know how to test functions_.

它只是一个函数，_而且我们知道如何测试函数_。

It fits in nicely with the commentary above:

它与上面的评论非常吻合：

- It takes a [`http.Request`](https://golang.org/pkg/net/http/#Request) which is just a bundle of data for us to inspect, parse and validate.
- > [A `http.ResponseWriter` interface is used by an HTTP handler to construct an HTTP response.](https://golang.org/pkg/net/http/#ResponseWriter)

- 它需要一个 [`http.Request`](https://golang.org/pkg/net/http/#Request)，它只是一组供我们检查、解析和验证的数据。
- > [HTTP 处理程序使用 `http.ResponseWriter` 接口来构造 HTTP 响应。](https://golang.org/pkg/net/http/#ResponseWriter)

### Super basic example test

### 超级基础示例测试

```go
func Teapot(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusTeapot)
}

func TestTeapotHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    res := httptest.NewRecorder()

    Teapot(res, req)

    if res.Code != http.StatusTeapot {
        t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
    }
}
```

To test our function, we _call_ it.

为了测试我们的函数，我们_调用_它。

For our test we pass a `httptest.ResponseRecorder` as our `http.ResponseWriter` argument, and our function will use it to write the `HTTP` response. The recorder will record (or _spy_ on) what was sent, and then we can make our assertions.

对于我们的测试，我们传递一个 `httptest.ResponseRecorder` 作为我们的 `http.ResponseWriter` 参数，我们的函数将使用它来编写 `HTTP` 响应。记录器将记录（或 _spy_ on）发送的内容，然后我们可以做出断言。

## Calling a `ServiceThing` in our handler

## 在我们的处理程序中调用一个 `ServiceThing`

A common complaint about TDD tutorials is that they're always "too simple" and not "real world enough". My answer to that is:

关于 TDD 教程的一个常见抱怨是它们总是“太简单”而不是“足够真实的世界”。我的回答是：

> Wouldn't it be nice if all your code was simple to read and test like the examples you mention?

> 如果您的所有代码都像您提到的示例一样易于阅读和测试，这不是很好吗？

This is one of the biggest challenges we face but need to keep striving for. It _is possible_ (although not necessarily easy) to design code, so it can be simple to read and test if we practice and apply good software engineering principles.

这是我们面临的最大挑战之一，但需要继续努力。 _有可能_（尽管不一定容易）设计代码，因此如果我们练习和应用良好的软件工程原理，阅读和测试会很简单。

Recapping what the handler from earlier does:

回顾之前的处理程序所做的事情：

1. Write HTTP responses, send headers, status codes, etc.
2. Decode the request's body into a `User`
3. Connect to a database (and all the details around that)
4. Query the database and applying some business logic depending on the result
5. Generate a password
6. Insert a record

1. 编写 HTTP 响应、发送标头、状态码等。
2. 将请求的主体解码为`User`
3. 连接到数据库（以及所有相关细节）
4.查询数据库并根据结果应用一些业务逻辑
5. 生成密码
6.插入一条记录

Taking the idea of a more ideal separation of concerns I'd want it to be more like:

考虑到更理想的关注点分离，我希望它更像：

1. Decode the request's body into a `User`
2. Call a `UserService.Register(user)` (this is our `ServiceThing`)
3. If there's an error act on it (the example always sends a `400 BadRequest` which I don't think is right, I'll just have a catch-all handler of a `500 Internal Server Error` _for now_. I must stress that returning `500` for all errors makes for a terrible API! Later on we can make the error handling more sophisticated, perhaps with [error types](error-types.md).
4. If there's no error, `201 Created` with the ID as the response body (again for terseness/laziness)

1. 将请求的正文解码为`User`
2. 调用一个`UserService.Register(user)`（这是我们的`ServiceThing`）
3. 如果有一个错误行为（这个例子总是发送一个我认为不正确的`400 BadRequest`，我现在只有一个`500 Internal Server Error`的catch-all处理程序_。我必须强调的是，为所有错误返回 `500` 是一个糟糕的 API！稍后我们可以使错误处理更加复杂，也许使用 [error types](error-types.md)。
4. 如果没有错误，`201 Created` 以 ID 作为响应主体（再次为了简洁/懒惰）

For the sake of brevity I won't go over the usual TDD process, check all the other chapters for examples.

为简洁起见，我不会讨论通常的 TDD 过程，请查看所有其他章节以获取示例。

### New design

###  新设计

```go
type UserService interface {
    Register(user User) (insertedID string, err error)
}

type UserServer struct {
    service UserService
}

func NewUserServer(service UserService) *UserServer {
    return &UserServer{service: service}
}

func (u *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    // request parsing and validation
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)

    if err != nil {
        http.Error(w, fmt.Sprintf("could not decode user payload: %v", err), http.StatusBadRequest)
        return
    }

    // call a service thing to take care of the hard work
    insertedID, err := u.service.Register(newUser)

    // depending on what we get back, respond accordingly
    if err != nil {
        //todo: handle different kinds of errors differently
        http.Error(w, fmt.Sprintf("problem registering new user: %v", err), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, insertedID)
}
```

Our `RegisterUser` method matches the shape of `http.HandlerFunc` so we're good to go. We've attached it as a method on a new type `UserServer` which contains a dependency on a `UserService` which is captured as an interface.

我们的 `RegisterUser` 方法匹配 `http.HandlerFunc` 的形状，所以我们很高兴。我们将它作为一种方法附加到新类型“UserServer”上，该类型包含对作为接口捕获的“UserService”的依赖。

Interfaces are a fantastic way to ensure our `HTTP` concerns are decoupled from any specific implementation; we can just call the method on the dependency, and we don't have to care _how_ a user gets registered.

接口是确保我们的“HTTP”关注点与任何特定实现分离的绝佳方式；我们可以只调用依赖项上的方法，我们不必关心_如何_注册用户。

If you wish to explore this approach in more detail following TDD read the [Dependency Injection](dependency-injection.md) chapter and the [HTTP Server chapter of the "Build an application" section](http-server.md).

如果您想在 TDD 之后更详细地探索这种方法，请阅读 [Dependency Injection](dependency-injection.md) 章节和 [“构建应用程序”部分的 HTTP Server 章节](http-server.md)。

Now that we've decoupled ourselves from any specific implementation detail around registration writing the code for our handler is straightforward and follows the responsibilities described earlier.

既然我们已经将自己与围绕注册的任何特定实现细节分离，为我们的处理程序编写代码就很简单了，并遵循前面描述的职责。

### The tests!

### 测试！

This simplicity is reflected in our tests.

这种简单性反映在我们的测试中。

```go
type MockUserService struct {
    RegisterFunc    func(user User) (string, error)
    UsersRegistered []User
}

func (m *MockUserService) Register(user User) (insertedID string, err error) {
    m.UsersRegistered = append(m.UsersRegistered, user)
    return m.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
    t.Run("can register valid users", func(t *testing.T) {
        user := User{Name: "CJ"}
        expectedInsertedID := "whatever"

        service := &MockUserService{
            RegisterFunc: func(user User) (string, error) {
                return expectedInsertedID, nil
            },
        }
        server := NewUserServer(service)

        req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
        res := httptest.NewRecorder()

        server.RegisterUser(res, req)

        assertStatus(t, res.Code, http.StatusCreated)

        if res.Body.String() != expectedInsertedID {
            t.Errorf("expected body of %q but got %q", res.Body.String(), expectedInsertedID)
        }

        if len(service.UsersRegistered) != 1 {
            t.Fatalf("expected 1 user added but got %d", len(service.UsersRegistered))
        }

        if !reflect.DeepEqual(service.UsersRegistered[0], user) {
            t.Errorf("the user registered %+v was not what was expected %+v", service.UsersRegistered[0], user)
        }
    })

    t.Run("returns 400 bad request if body is not valid user JSON", func(t *testing.T) {
        server := NewUserServer(nil)

        req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("trouble will find me"))
        res := httptest.NewRecorder()

        server.RegisterUser(res, req)

        assertStatus(t, res.Code, http.StatusBadRequest)
    })

    t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {
        user := User{Name: "CJ"}

        service := &MockUserService{
            RegisterFunc: func(user User) (string, error) {
                return "", errors.New("couldn't add new user")
            },
        }
        server := NewUserServer(service)

        req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
        res := httptest.NewRecorder()

        server.RegisterUser(res, req)

        assertStatus(t, res.Code, http.StatusInternalServerError)
    })
}
```

Now our handler isn't coupled to a specific implementation of storage it is trivial for us to write a `MockUserService` to help us write simple, fast unit tests to exercise the specific responsibilities it has.

现在我们的处理程序没有与存储的特定实现耦合，编写一个 `MockUserService` 来帮助我们编写简单、快速的单元测试来执行它所具有的特定职责对我们来说是微不足道的。

### What about the database code? You're cheating!

### 数据库代码呢？你在作弊！

This is all very deliberate. We don't want HTTP handlers concerned with our business logic, databases, connections, etc.

这都是非常刻意的。我们不希望 HTTP 处理程序与我们的业务逻辑、数据库、连接等有关。

By doing this we have liberated the handler from messy details, we've _also_ made it easier to test our persistence layer and business logic as it is also no longer coupled to irrelevant HTTP details.

通过这样做，我们将处理程序从杂乱的细节中解放出来，我们_也_使测试我们的持久层和业务逻辑变得更加容易，因为它也不再与无关的 HTTP 细节耦合。

All we need to do is now implement our `UserService` using whatever database we want to use

我们现在需要做的就是使用我们想要使用的任何数据库来实现我们的“UserService”

```go
type MongoUserService struct {
}

func NewMongoUserService() *MongoUserService {
    //todo: pass in DB URL as argument to this function
    //todo: connect to db, create a connection pool
    return &MongoUserService{}
}

func (m MongoUserService) Register(user User) (insertedID string, err error) {
    // use m.mongoConnection to perform queries
    panic("implement me")
}
```

We can test this separately and once we're happy in `main` we can snap these two units together for our working application.

我们可以单独测试它，一旦我们对 `main` 感到满意，我们就可以将这两个单元组合在一起用于我们的工作应用程序。

```go
func main() {
    mongoService := NewMongoUserService()
    server := NewUserServer(mongoService)
    http.ListenAndServe(":8000", http.HandlerFunc(server.RegisterUser))
}
```

### A more robust and extensible design with little effort

### 更健壮和可扩展的设计，只需很少的努力

These principles not only make our lives easier in the short-term they make the system easier to extend in the future. 

这些原则不仅使我们的生活在短期内更轻松，而且使系统在未来更容易扩展。

It wouldn't be surprising that further iterations of this system we'd want to email the user a confirmation of registration.

我们希望通过电子邮件向用户发送注册确认信息，对该系统的进一步迭代并不奇怪。

With the old design we'd have to change the handler _and_ the surrounding tests. This is often how parts of code become unmaintainable, more and more functionality creeps in because it's already _designed_ that way; for the "HTTP handler" to handle... everything!

使用旧设计，我们必须更改处理程序_和_周围的测试。这通常是代码部分变得不可维护的原因，越来越多的功能潜入其中，因为它已经_设计_了；对于“HTTP 处理程序”来处理......一切！

By separating concerns using an interface we don't have to edit the handler _at all_ because it's not concerned with the business logic around registration.

通过使用接口分离关注点，我们不必编辑处理程序_根本_，因为它不关心注册的业务逻辑。

## Wrapping up

##  总结

Testing Go's HTTP handlers is not challenging, but designing good software can be!

测试 Go 的 HTTP 处理程序并不具有挑战性，但设计好的软件却可以！

People make the mistake of thinking HTTP handlers are special and throw out good software engineering practices when writing them which then makes testing them challenging.

人们错误地认为 HTTP 处理程序是特殊的，并在编写它们时抛弃了良好的软件工程实践，这使得测试它们具有挑战性。

Reiterating again; **Go's http handlers are just functions**. If you write them like you would other functions, with clear responsibilities, and a good separation of concerns you will have no trouble testing them, and your codebase will be healthier for it. 

再次重申； **Go 的 http 处理程序只是函数**。如果您像编写其他函数一样编写它们，具有明确的职责和良好的关注点分离，那么您将可以轻松测试它们，并且您的代码库将因此更健康。

