# Programming with errors

# 编程出错

**2019 09 11**

**2019 09 11**

Go 1.13 introduces an enhanced [package errors](https://golang.org/pkg/errors) (née [xerrors](https://godoc.org/golang.org/x/xerrors)) which roughly standardizes programming with errors. Personally, I find the API confusing. This is a quick reference for how to use it effectively.

Go 1.13 引入了增强的 [包错误](https://golang.org/pkg/errors) (née [xerrors](https://godoc.org/golang.org/x/xerrors))，它大致标准化了编程错误。就个人而言，我觉得 API 令人困惑。这是有关如何有效使用它的快速参考。

### Creating errors

### 创建错误

Sentinel errors work the same as before. Name them as ErrXxx, and create them with errors.New.

哨兵错误的工作方式与以前相同。将它们命名为 ErrXxx，并使用 errors.New 创建它们。

```go
var ErrFoo = errors.New("foo error")
```

Error types basically work the same as before. Name them as XxxError, and make sure they have an Error method, to satisfy the error interface.

错误类型的工作方式与以前基本相同。将它们命名为 XxxError，并确保它们具有 Error 方法，以满足错误接口。

```go
type BarError struct {
    Reason string
}

func (e BarError) Error() string {
    return fmt.Sprintf("bar error: %s", e.Reason)
}
```

If your error type wraps another error, provide an Unwrap method.

如果您的错误类型包含另一个错误，请提供一个 Unwrap 方法。

```go
type BazError struct {
    Reason string
    Inner  error
}

func (e BazError) Error() string {
    if e.Inner != nil {
        return fmt.Sprintf("baz error: %s: %v", e.Reason, e.Inner)
    }
    return fmt.Sprintf("baz error: %s", e.Reason)
}

func (e BazError) Unwrap() error {
    return e.Inner
}
```

### Wrapping and returning errors

### 包装和返回错误

By default, when you encounter an error in a function and need to return it to the caller, wrap it with some context about what went wrong, using [fmt.Errorf](https://golang.org/pkg/fmt#Errorf) and the new `%w` verb.

默认情况下，当您在函数中遇到错误并需要将其返回给调用者时，使用 [fmt.Errorf](https://golang.org/pkg/fmt#Errorf) 和新的 `%w` 动词。

```go
func process(j Job) error {
    result, err := preprocess(j)
    if err != nil {
        return fmt.Errorf("error preprocessing job: %w", err)
    }
```

This process is called error annotation. Avoid returning un-annotated errors, because that can make it difficult for callers to understand what went wrong.

这个过程称为错误注释。避免返回未注释的错误，因为这会使调用者难以理解哪里出了问题。

Also, consider wrapping errors via a custom error type (like BazError, above) for more sophisticated use cases.

此外，对于更复杂的用例，请考虑通过自定义错误类型（如上面的 BazError）包装错误。

```go
     p := getPriority()
    widget, err := manufacture(p, result)
    if err != nil {
        return ManufacturingError{Priority: p, Error: err}
    }
```

### Checking errors

### 检查错误

Most of the time, when you receive an error, you don’t care about the details. Whatever you were trying to do failed, so you either need to report the error (e.g. log it) and carry on; or, if it’s not possible to continue, annotate the error with context, and return it to the caller.

大多数情况下，当您收到错误消息时，您并不关心细节。无论你尝试做什么都失败了，所以你要么需要报告错误（例如记录它）并继续；或者，如果无法继续，请使用上下文注释错误，并将其返回给调用者。

If you care about which error you received, you can check for sentinel errors with [errors.Is](https://golang.org/pkg/errors#Is), and for error values with [errors.As](https://golang.org/pkg/errors#As).

如果您关心收到的错误，可以使用 [errors.Is](https://golang.org/pkg/errors#Is) 检查标记错误，并使用 [errors.As](https://golang.org/pkg/errors#As)。

```go
err := f()
if errors.Is(err, ErrFoo) {
    // you know you got an ErrFoo
    // respond appropriately
}

var bar *BarError
if errors.As(err, &bar) {
    // you know you got a BarError
    // bar's fields are populated
    // respond appropriately
}
```

errors.Is and errors.As will both try to unwrap the error, recursively, in order to find a match. **[This code](https://play.golang.org/p/GorSR6HTWzf) demonstrates basic error wrapping and checking techniques.** Look at the order of the checks in `func a()`, and then try changing the error that's returned by `func c()`, to get an intuition about how everything works.

errors.Is 和 errors.As 都将尝试递归地解包错误，以便找到匹配项。 **[此代码](https://play.golang.org/p/GorSR6HTWzf) 演示了基本的错误包装和检查技术。** 查看 `func a()` 中的检查顺序，然后尝试更改`func c()` 返回的错误，以获得关于一切如何工作的直觉。

As the [package errors docs](https://golang.org/pkg/errors/) state, prefer using errors.Is over checking plain equality, e.g. `if err == ErrFoo`; and prefer using errors.As over plain type assertions, e.g. `if e, ok := err.(MyError)`, because the plain versions don’t perform unwrapping. If you explicitly don’t want to allow callers to unwrap errors, provide a different formatting verb to `fmt.Errorf`, like `%v`; or don’t provide an `Unwrap` method on your error type. But these cases should be rare. 

正如 [包错误文档](https://golang.org/pkg/errors/) 所述，更喜欢使用错误。比检查简单的相等性，例如`如果错误 == ErrFoo`;并且更喜欢使用errors.As而不是普通类型的断言，例如`if e, ok := err.(MyError)`，因为普通版本不执行解包。如果您明确不想让调用者解开错误，请为 `fmt.Errorf` 提供不同的格式动词，例如 `%v`；或者不要在你的错误类型上提供 `Unwrap` 方法。但这些情况应该很少见。

