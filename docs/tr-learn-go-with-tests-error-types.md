# Error types

# 错误类型

**[You can find all the code here](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/error-types)**

**[你可以在这里找到所有代码](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/error-types)**

**Creating your own types for errors can be an elegant way of tidying up your code, making your code easier to use and test.**

**创建自己的错误类型是整理代码的一种优雅方式，使代码更易于使用和测试。**

Pedro on the Gopher Slack asks

Gopher Slack 上的 Pedro 问

> If I’m creating an error like `fmt.Errorf("%s must be foo, got %s", bar, baz)`, is there a way to test equality without comparing the string value?

> 如果我创建了一个类似 `fmt.Errorf("%s must be foo, got %s", bar, baz)` 这样的错误，有没有办法在不比较字符串值的情况下测试相等性？

Let's make up a function to help explore this idea.

让我们编写一个函数来帮助探索这个想法。

```go
// DumbGetter will get the string body of url if it gets a 200
func DumbGetter(url string) (string, error) {
    res, err := http.Get(url)

    if err != nil {
        return "", fmt.Errorf("problem fetching from %s, %v", url, err)
    }

    if res.StatusCode != http.StatusOK {
        return "", fmt.Errorf("did not get 200 from %s, got %d", url, res.StatusCode)
    }

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body) // ignoring err for brevity

    return string(body), nil
}
```

It's not uncommon to write a function that might fail for different reasons and we want to make sure we handle each scenario correctly.

编写一个可能因不同原因而失败的函数并不少见，我们希望确保正确处理每个场景。

As Pedro says, we _could_ write a test for the status error like so.

正如 Pedro 所说，我们_可以_像这样为状态错误编写一个测试。

```go
t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {

    svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        res.WriteHeader(http.StatusTeapot)
    }))
    defer svr.Close()

    _, err := DumbGetter(svr.URL)

    if err == nil {
        t.Fatal("expected an error")
    }

    want := fmt.Sprintf("did not get 200 from %s, got %d", svr.URL, http.StatusTeapot)
    got := err.Error()

    if got != want {
        t.Errorf(`got "%v", want "%v"`, got, want)
    }
})
```

This test creates a server which always returns `StatusTeapot` and then we use its URL as the argument to `DumbGetter` so we can see it handles non `200` responses correctly.

这个测试创建了一个总是返回 `StatusTeapot` 的服务器，然后我们使用它的 URL 作为 `DumbGetter` 的参数，所以我们可以看到它正确处理非 `200` 响应。

## Problems with this way of testing

## 这种测试方式的问题

This book tries to emphasise _listen to your tests_ and this test doesn't _feel_ good:

这本书试图强调_听你的测试_而这个测试_感觉_不好：

- We're constructing the same string as production code does to test it
- It's annoying to read and write
- Is the exact error message string what we're _actually concerned with_ ?

- 我们正在构建与生产代码相同的字符串来测试它
- 阅读和写作很烦人
- 确切的错误消息字符串是我们_实际关心的_吗？

What does this tell us? The ergonomics of our test would be reflected on another bit of code trying to use our code.

这告诉我们什么？我们测试的人体工程学将反映在尝试使用我们代码的另一段代码上。

How does a user of our code react to the specific kind of errors we return? The best they can do is look at the error string which is extremely error prone and horrible to write.

我们代码的用户对我们返回的特定类型的错误有何反应？他们能做的最好的事情就是查看错误字符串，该字符串极易出错且编写起来很糟糕。

## What we should do

## 我们应该做什么

With TDD we have the benefit of getting into the mindset of:

使用 TDD，我们可以进入以下思维模式：

> How would _I_ want to use this code?

> _I_ 要如何使用此代码？

What we could do for `DumbGetter` is provide a way for users to use the type system to understand what kind of error has happened.

我们可以为 `DumbGetter` 做的是为用户提供一种使用类型系统来了解发生了什么样的错误的方法。

What if `DumbGetter` could return us something like

如果`DumbGetter` 可以返回给我们类似的东西怎么办

```go
type BadStatusError struct {
    URL    string
    Status int
}
```

Rather than a magical string, we have actual _data_ to work with.

而不是一个神奇的字符串，我们有实际的 _data_ 可以使用。

Let's change our existing test to reflect this need

让我们改变我们现有的测试来反映这个需求

```go
t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {

    svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        res.WriteHeader(http.StatusTeapot)
    }))
    defer svr.Close()

    _, err := DumbGetter(svr.URL)

    if err == nil {
        t.Fatal("expected an error")
    }

    got, isStatusErr := err.(BadStatusError)

    if !isStatusErr {
        t.Fatalf("was not a BadStatusError, got %T", err)
    }

    want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}

    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
})
```

We'll have to make `BadStatusError` implement the error interface.

我们必须让 `BadStatusError` 实现错误接口。

```go
func (b BadStatusError) Error() string {
    return fmt.Sprintf("did not get 200 from %s, got %d", b.URL, b.Status)
}
```

### What does the test do?

### 测试有什么作用？

Instead of checking the exact string of the error, we are doing a [type assertion](https://tour.golang.org/methods/15) on the error to see if it is a `BadStatusError`. This reflects our desire for the _kind_ of error clearer. Assuming the assertion passes we can then check the properties of the error are correct.

我们没有检查错误的确切字符串，而是对错误进行 [类型断言](https://tour.golang.org/methods/15) 以查看它是否是“BadStatusError”。这反映了我们对 _kind_ 错误更清晰的渴望。假设断言通过，我们就可以检查错误的属性是否正确。

When we run the test, it tells us we didn't return the right kind of error

当我们运行测试时，它告诉我们没有返回正确类型的错误

```
--- FAIL: TestDumbGetter (0.00s)
    --- FAIL: TestDumbGetter/when_you_dont_get_a_200_you_get_a_status_error (0.00s)
        error-types_test.go:56: was not a BadStatusError, got *errors.errorString
```

Let's fix `DumbGetter` by updating our error handling code to use our type

让我们通过更新我们的错误处理代码来使用我们的类型来修复 `DumbGetter`

```go
if res.StatusCode != http.StatusOK {
    return "", BadStatusError{URL: url, Status: res.StatusCode}
}
```

This change has had some _real positive effects_ 

这一变化产生了一些_真正的积极影响_

- Our `DumbGetter` function has become simpler, it's no longer concerned with the intricacies of an error string, it just creates a `BadStatusError`.
- Our tests now reflect (and document) what a user of our code _could_ do if they decided they wanted to do some more sophisticated error handling than just logging. Just do a type assertion and then you get easy access to the properties of the error.
- It is still "just" an `error`, so if they choose to they can pass it up the call stack or log it like any other `error`.

- 我们的 `DumbGetter` 函数变得更简单，它不再关心错误字符串的复杂性，它只是创建一个 `BadStatusError`。
- 我们的测试现在反映（和记录）我们代码的用户_可以_如果他们决定做一些更复杂的错误处理而不只是记录。只需进行类型断言，然后您就可以轻松访问错误的属性。
- 它仍然“只是”一个“错误”，所以如果他们选择这样做，他们可以将它传递到调用堆栈或像任何其他“错误”一样记录它。

## Wrapping up

##  总结

If you find yourself testing for multiple error conditions don't fall in to the trap of comparing the error messages.

如果您发现自己测试了多个错误条件，请不要陷入比较错误消息的陷阱。

This leads to flaky and difficult to read/write tests and it reflects the difficulties the users of your code will have if they also need to start doing things differently depending on the kind of errors that have occurred.

这会导致不稳定且难以读/写的测试，并且它反映了如果您的代码用户还需要根据发生的错误类型开始以不同的方式做事时将会遇到的困难。

Always make sure your tests reflect how _you'd_ like to use your code, so in this respect consider creating error types to encapsulate your kinds of errors. This makes handling different kinds of errors easier for users of your code and also makes writing your error handling code simpler and easier to read.

始终确保您的测试反映了_您希望_如何使用您的代码，因此在这方面考虑创建错误类型来封装您的错误类型。这使得代码用户更容易处理不同类型的错误，也使编写错误处理代码更简单、更容易阅读。

## Addendum

## 附录

As of Go 1.13 there are new ways to work with errors in the standard library which is covered in the [Go Blog](https://blog.golang.org/go1.13-errors)

从 Go 1.13 开始，[Go 博客](https://blog.golang.org/go1.13-errors) 中介绍了处理标准库中错误的新方法

```go
t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {

    svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        res.WriteHeader(http.StatusTeapot)
    }))
    defer svr.Close()

    _, err := DumbGetter(svr.URL)

    if err == nil {
        t.Fatal("expected an error")
    }

    var got BadStatusError
    isBadStatusError := errors.As(err, &got)
    want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}

    if !isBadStatusError {
        t.Fatalf("was not a BadStatusError, got %T", err)
    }

    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
})
```

In this case we are using [`errors.As`](https://golang.org/pkg/errors/#example_As) to try and extract our error into our custom type. It returns a `bool` to denote success and extracts it into `got` for us. 

在这种情况下，我们使用 [`errors.As`](https://golang.org/pkg/errors/#example_As) 尝试将我们的错误提取到我们的自定义类型中。它返回一个 `bool` 来表示成功，并为我们将其提取到 `got` 中。

