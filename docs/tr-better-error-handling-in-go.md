# Better Error Handling, in Go

# 更好的错误处理，在 Go 中

Tuesday March 26th 2019

2019 年 3 月 26 日星期二

It’s no secret that we’re big golang fans at bet365. It has allowed us to create exciting new products, like Bet Builder and our custom sports search engine. Within the very near future, the majority of our sports website will be powered by _Go_.

我们是 bet365 的 golang 忠实粉丝，这已经不是什么秘密了。它使我们能够创建令人兴奋的新产品，例如 Bet Builder 和我们的自定义体育搜索引擎。在不久的将来，我们的大部分体育网站都将由 _Go_ 提供支持。

We’ve used this move in technology as an opportunity to review what simplicity means within our vast codebase, and embrace the idiomatic ‘Go way’.

我们利用技术上的这一举措来回顾我们庞大的代码库中的简单性意味着什么，并接受惯用的“Go way”。

Our internal standards document bares signage to “leave your OO baggage at the door”.

我们的内部标准文件标明“将您的 OO 行李留在门口”。

And we did; Our code organization has become more literal, with less obfuscating OO patterns, clearer logic paths and a focus on reducing ‘code golf’ for maintainability. Go promotes this kind of behavior and generally makes it obvious when it’s punishing you for not abiding by its rules.

我们做到了；我们的代码组织变得更加文字化，减少了混淆的 OO 模式，更清晰的逻辑路径，并专注于减少“代码高尔夫”以实现可维护性。 Go 提倡这种行为，并且通常会在它惩罚你不遵守其规则时表现出来。

## I’ll get to the point

## 我进入正题

One great example of this simplicity is the way it handles errors at runtime; Rather than adding a bespoke language construct for dealing with the raising and propagation of errors throughout your call stacks, go chose to simply make use of its ability to return more than one value from a function. This is great because there is no confusion when you’re invoking a function or method as to whether or not it is your responsibility to handle what happens when something goes wrong. If you see an `error` is part of that functions return signature, then it’s on you to deal with it.

这种简单性的一个很好的例子是它在运行时处理错误的方式； go 没有添加定制的语言结构来处理整个调用堆栈中错误的产生和传播，而是选择简单地利用其从函数返回多个值的能力。这很好，因为当您调用函数或方法时，您不会混淆是否有责任处理出现问题时发生的情况。如果您看到 `error` 是该函数返回签名的一部分，那么您需要处理它。

Now, whilst this does make your code declaratively more obvious, one thing we noticed was if you're calling a few functions which return errors it's quite easy for things to get a little _too obvious_ that you're doing a bit of error handling by falling into the pattern I'm about to demonstrate.

现在，虽然这确实使您的代码在声明上更加明显，但我们注意到的一件事是，如果您正在调用一些返回错误的函数，很容易让事情变得有点_太明显了_您正在执行一些错误处理陷入我将要展示的模式。

You may have seen this pattern occurring yourself, you may have solved this problem yourself in a different way, you may be totally fine with this code and not see it as a problem at all, and that’s okay.

您可能已经亲眼见过这种模式发生，您可能自己以不同的方式解决了这个问题，您可能对这段代码完全没问题，根本不认为它是一个问题，这没关系。

## Our problem, an example

## 我们的问题，一个例子

Let’s take a look at a somewhat contrived example of what I’m talking about. We’ll use a typical HTTP handler for some context

让我们看一个有点人为的例子来说明我在说什么。我们将对某些上下文使用典型的 HTTP 处理程序

``` js

``` js

func myHandler(w http.Response, r *http.Request) {

     func myHandler(w http.Response, r *http.Request) {

    err := validateRequest(r)
     if err != nil {
         log.Printf("error validating request to myHandler - err: %v", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     错误:= 验证请求(r)
    如果错误！= nil {
        log.Printf("验证对 myHandler 的请求时出错 - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        返回
    }

    user, err := getUserFromRequest(r)
     if err != nil {
         log.Printf("error getting user from request in myHandler - err: %v", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     用户，错误：= getUserFromRequest(r)
    如果错误！= nil {
        log.Printf("从 myHandler 中的请求获取用户时出错 - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        返回
    }

    dataset, err := db.GetUserData(user)
     if err != nil {
         log.Printf("error retrieving user data in myHandler - err: %v", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     数据集，错误：= db.GetUserData（用户）
    如果错误！= nil {
        log.Printf("在 myHandler 中检索用户数据时出错 - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        返回
    }

    buffer := newBuffer()
     err := serialize.UserData(dataset, &buffer)
     if err != nil {
         log.Printf("error serializing user data in myHandler - err %v", err)
         w.WriteHeader(http.StatusInternalServerError)
         return
     }

     缓冲区 := newBuffer()
    错误 := serialize.UserData(dataset, &buffer)
    如果错误！= nil {
        log.Printf("在 myHandler 中序列化用户数据时出错 - err %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        返回
    }

    err := buffer.WriteTo(w);
     if err != nil {
         log.Printf("error writing buffer to response in myHandler - err %v", err)
         return
     }
}

错误：= buffer.WriteTo(w);
    如果错误！= nil {
        log.Printf("将缓冲区写入 myHandler 中的响应时出错 - err %v", err)
        返回
    }
}

```

``

Okay, so a nice straightforward HTTP handler. Let’s quickly break down what tasks it was performing for us.

好的，这是一个很好的直接 HTTP 处理程序。让我们快速分解一下它为我们执行的任务。

- Validate the request
- Grab a user from the request
- Fetch a user dataset from the database
- Serialize the dataset into a buffer
- Write the buffer to the HTTP response

- 验证请求
- 从请求中获取用户
- 从数据库中获取用户数据集
- 将数据集序列化为缓冲区
- 将缓冲区写入 HTTP 响应

So we've done those five things, fine, but take a look at that code again – it doesn't feel like we just did five basic things there, it feels like a lot more was happening due to the noise created by the error handling after each action we took. 

所以我们已经完成了这五件事，很好，但是再看看那段代码——感觉我们不只是在那里做了五件基本的事情，感觉由于错误产生的噪音，发生了更多的事情在我们采取的每个行动之后处理。

The first thing to consider here is the amount of additional context we’re wrapping into each one of these errors for logging purposes. The messages generally start off with `"there was an error doing thing"` – which seems like an obvious thing to pad out your error message with, but it could actually be redundant when you consider the function it originated from could've (should 've) added enough context about where the error originated from itself. We found we were able to remove a lot of `return fmt.Errorf("there was a db error %v", err)` statements that were simply adding zero additional information.

这里要考虑的第一件事是我们为了记录目的而包装到每个错误中的额外上下文的数量。这些消息通常以“做事时有错误”开头——这似乎是一个显而易见的事情来填补你的错误消息，但当你考虑它源自的函数可能已经（应该）时，它实际上可能是多余的've) 添加了足够的关于错误源自何处的上下文。我们发现我们能够删除很多 `return fmt.Errorf("there was an db error %v", err)` 语句，这些语句只是添加了零附加信息。

Of course, that doesn’t necessarily mean the error was raised in the function we called. The point is wherever we have an error ‘origin site’ within our own codebase, we should take the opportunity to decorate the error with enough context that it doesn’t need wrapping multiple times as we pass it back up the stack. This could be parameter information, generated URLs or other context useful to someone diagnosing an issue later on. We can then simply bare-return errors up the stack up until the point we can handle them – in our case, exit safely and log the error.

当然，这并不一定意味着错误是在我们调用的函数中引发的。关键是在我们自己的代码库中有一个错误的“起源站点”，我们应该抓住机会用足够的上下文来修饰错误，这样当我们将它传回堆栈时，它不需要多次包装。这可能是参数信息、生成的 URL 或其他对稍后诊断问题有用的上下文。然后我们可以简单地在堆栈中直接返回错误，直到我们可以处理它们——在我们的例子中，安全退出并记录错误。

So how do we remove the need to wrap these errors multiple times during the execution path of our handler but still maintain the same level of information? After all, the information we wrapped it with wasn’t useless, it also stated `in myHandler` as the error ‘destination site’, which again could be vital in a diagnosis investigation by an engineer looking at the issue later on.

那么我们如何在处理程序的执行路径中消除多次包装这些错误的需要，但仍然保持相同级别的信息？毕竟，我们用它包装的信息并不是无用的，它还将“在 myHandler”中声明为错误的“目标站点”，这对于稍后查看问题的工程师进行的诊断调查也很重要。

After trying a few different approaches, and there are many ways to do this, we took inspiration from how go 2.0 had proposed solving the problem ( [check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)) and came up with a more concise solution we can use today using what we call the error defer pattern. A halfway house, if you like.

在尝试了几种不同的方法之后，有很多方法可以做到这一点，我们从 go 2.0 提出的解决问题的建议中获得了灵感（[check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)) 并提出了一个更简洁的解决方案，我们今天可以使用我们称之为错误延迟模式的方法。一个中途的房子，如果你喜欢。

## The solution, an example

## 解决方法，一个例子

This is the same code, but we’re going to switch out the error handling to use error defer. Let’s take a look

这是相同的代码，但我们将关闭错误处理以使用错误延迟。让我们来看看

``` js

``` js

func myHandler(w http.Response, r *http.Request) {

     func myHandler(w http.Response, r *http.Request) {

    var err error
     defer func() {
         if err != nil {
             log.Printf("error in myHandler - error: %v", err)
             w.WriteHeader(http.StatusInternalServerErrror)
         }
     }()

     变量错误
    延迟功能（）{
        如果错误！= nil {
            log.Printf("myHandler 中的错误 - 错误：%v", err)
            w.WriteHeader(http.StatusInternalServerErrror)
        }
    }()

    err = validateRequest(r)
     if err != nil { return }

     错误 = 验证请求（r）
    如果错误 != nil { 返回 }

    user, err := getUserFromRequest(r)
     if err != nil { return  }

     用户，错误：= getUserFromRequest(r)
    如果错误 != nil { 返回 }

    dataset, err := db.GetUserData(user)
     if err != nil { return }

     数据集，错误：= db.GetUserData（用户）
    如果错误 != nil { 返回 }

    buffer := newBuffer()
     err = serialize.UserData(dataset, &buffer)
     if err != nil { return }

     缓冲区 := newBuffer()
    err = serialize.UserData(dataset, &buffer)
    如果错误 != nil { 返回 }

    err2 := buffer.WriteTo(w)
     if err2 != nil {
         log.Printf("error writing buffer to response in myHandler - error %v", err2)
         return
     }
}

err2 := buffer.WriteTo(w)
    如果 err2 != nil {
        log.Printf("将缓冲区写入 myHandler 中的响应时出错 - 错误 %v", err2)
        返回
    }
}

```

``

Straight away, that feels much easier and quicker to digest. We can see the application logic hasn’t changed but has become more visible due to the error handling blocks being reduced to a more basic form. We’ve also made more use of the context added at the origin sites within the functions we’ve called.

直接，这感觉更容易和更快地消化。我们可以看到应用程序逻辑没有改变，但由于错误处理块被简化为更基本的形式而变得更加可见。我们还更多地使用了在我们调用的函数中添加到源站点的上下文。

The error we respond with is now tracked by the top level `var err error`. Each function which returns an error writes to that single declaration, then simply checks the value is not `nil` each time it can be set, before it continues to the next statement. If the value isn’t `nil` we just return. When we return, no matter what stage we do, we always drop into the `defer func` we declared at the top of the handler. From there, we check the error, wrap with the destination site information and continue to handle as previously by logging and setting an http status code.

我们响应的错误现在由顶层 `var err error` 跟踪。每个返回错误的函数都会写入该单个声明，然后在每次可以设置时简单地检查该值是否为“nil”，然后再继续执行下一条语句。如果值不是 `nil`，我们就返回。当我们返回时，无论我们做什么阶段，我们总是会进入我们在处理程序顶部声明的 `defer func`。从那里，我们检查错误，用目标站点信息包装并通过记录和设置 http 状态代码继续像以前一样处理。

You'll notice the last block shows an exception to the pattern and how we can deal with scenarios where sharing an error handler is not sufficient – We failed the write, so there's little point setting an http status code at that point, we just need to log the error, so we don't write to `err`. 

你会注意到最后一个块显示了模式的一个异常，以及我们如何处理共享错误处理程序不够的场景——我们写入失败，所以此时设置 http 状态代码没什么意义，我们只需要记录错误，所以我们不写入 `err`。

Without wanting to bang on about errors too much more, there are also other options within this approach, for example we can tag more information into custom `error` structs, then check their types in the `defer func` to give ourselves additional context or configuration whilst still keeping this handling code out of the main execution path, like so

不想过多地讨论错误，这种方法中还有其他选项，例如我们可以将更多信息标记到自定义 `error` 结构中，然后在 `defer func` 中检查它们的类型以提供额外的上下文或配置，同时仍将此处理代码保留在主执行路径之外，就像这样

``` js

``` js

if custom, ok := err.(CustomErrorType); ok {
      /// handle the custom type
}

如果自定义，好的 := err.(CustomErrorType);好的 {
     /// 处理自定义类型
}

```

``

## Conclusion

##  结论

Go is a fantastic language to work with, but like any tool it can easily become unwieldy without proper care for how it is used. Error handling has been identified by the wider community as being somewhat of a bugbear, so this is how we solved that problem in a way that worked for us. We look forward to using [check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md) when it finally arrives in golang 2.0!

Go 是一种非常适合使用的语言，但与任何工具一样，如果不注意如何使用它，它很容易变得笨拙。更广泛的社区认为错误处理有点麻烦，所以这就是我们以适合我们的方式解决该问题的方式。我们期待在 [check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md) 最终到达 golang 2.0 时使用它！

#### Pete G

#### 皮特G

Senior Software Architect 

高级软件架构师

