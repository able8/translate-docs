# Panics, stack traces and how to recover [best practice]

# 恐慌、堆栈跟踪以及如何恢复 [最佳实践]

yourbasic.org/golang

## A panic is an exception in Go

## 恐慌是 Go 中的一个例外

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds. To signify events such as end-of-file, Go programs use the built-in `error` type.  See [Error handling best practice](http://yourbasic.org/golang/errors-explained/) and [3 simple ways to create an error](http://yourbasic.org/golang/create-error/) for more on errors.

恐慌类似于 C++ 和 Java 异常，但仅用于运行时错误，例如跟随 nil 指针或尝试索引数组越界。为了表示诸如文件结束之类的事件，Go 程序使用内置的 `error` 类型。请参阅 [错误处理最佳实践](http://yourbasic.org/golang/errors-explained/) 和 [3 种创建错误的简单方法](http://yourbasic.org/golang/create-error/) 了解更多关于错误。

A panic stops the normal execution of a goroutine:

恐慌会停止 goroutine 的正常执行：

- When a program panics, it immediately starts to unwind the call stack.
- This continues until the program crashes and prints a stack trace,
- or until the built-in`recover` function is called.

- 当程序发生恐慌时，它会立即开始展开调用堆栈。
- 这一直持续到程序崩溃并打印堆栈跟踪，
- 或者直到调用内置的`recover` 函数。

A panic is caused either by a runtime error, or an explicit call to the built-in `panic` function.

恐慌是由运行时错误引起的，或对内置 `panic` 函数的显式调用。

## Stack traces

## 堆栈跟踪

A **stack trace** – a report of all active stack frames – is typically printed to the console when a panic occurs.
Stack traces can be very useful for debugging:

**堆栈跟踪** – 所有活动堆栈帧的报告 –通常在发生恐慌时打印到控制台。
堆栈跟踪对于调试非常有用：

- not only do you see **where** the error happened,
- but also **how** the program arrived in this place.

- 您不仅可以看到错误发生的**位置**，
- 还有**如何**程序到达这个地方。

### Interpret a stack trace

### 解释堆栈跟踪

Here’s an example of a stack trace:

这是堆栈跟踪的示例：

```
goroutine 11 [running]:
testing.tRunner.func1(0xc420092690)
    /usr/local/go/src/testing/testing.go:711 +0x2d2
panic(0x53f820, 0x594da0)
    /usr/local/go/src/runtime/panic.go:491 +0x283
github.com/yourbasic/bit.(*Set).Max(0xc42000a940, 0x0)
    ../src/github.com/bit/set_math_bits.go:137 +0x89
github.com/yourbasic/bit.TestMax(0xc420092690)
    ../src/github.com/bit/set_test.go:165 +0x337
testing.tRunner(0xc420092690, 0x57f5e8)
    /usr/local/go/src/testing/testing.go:746 +0xd0
created by testing.(*T).Run
    /usr/local/go/src/testing/testing.go:789 +0x2de

```


It can be read from the bottom up:

可以从下往上读：

- `testing.(*T).Run` has called `testing.tRunner`,
- which has called`bit.TestMax`,
- which has called`bit.(*Set).Max`,
- which has called`panic`,
- which has called`testing.tRunner.func1`.

  


The indented lines show the source file and line number at which the function was called.
The hexadecimal numbers refer to parameter values, including values of pointers and internal data structures.
[Stack Traces in Go](https://www.goinggo.net/2015/01/stack-traces-in-go.html) has more details.

缩进的行显示调用函数的源文件和行号。
十六进制数指的是参数值，包括指针和内部数据结构的值。
[Go 中的堆栈跟踪](https://www.goinggo.net/2015/01/stack-traces-in-go.html) 有更多细节。

### Print and log a stack trace

### 打印并记录堆栈跟踪

To print the stack trace for the current goroutine, use [`debug.PrintStack`](https://golang.org/pkg/runtime/debug/#PrintStack) from package [`runtime/debug`](https://golang.org/pkg/runtime/debug/).

要打印当前 goroutine 的堆栈跟踪，请使用包 [`runtime/debug`](https://golang.org/pkg/runtime/debug/#PrintStack) 中的 [`debug.PrintStack`](https://golang.org/pkg/runtime/debug/#PrintStack)golang.org/pkg/runtime/debug/)。

You can also examine the current stack trace programmatically by calling [`runtime.Stack`](https://golang.org/pkg/runtime/#Stack).

您还可以通过调用 [`runtime.Stack`](https://golang.org/pkg/runtime/#Stack) 以编程方式检查当前堆栈跟踪。

### Level of detail

###  详细程度

The [`GOTRACEBACK`](https://golang.org/pkg/runtime/#hdr-Environment_Variables) variable controls the amount of output generated when a Go program fails.

[`GOTRACEBACK`](https://golang.org/pkg/runtime/#hdr-Environment_Variables) 变量控制 Go 程序失败时生成的输出量。

- `GOTRACEBACK=none` omits the goroutine stack traces entirely.
- `GOTRACEBACK=single` (the default) prints a stack trace for the current goroutine, eliding functions internal to the run-time system. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time.
- `GOTRACEBACK=all` adds stack traces for all user-created goroutines.
- `GOTRACEBACK=system` is like `all` but adds stack frames for run-time functions and shows goroutines created internally by the run-time.



- `GOTRACEBACK=none` 完全省略了 goroutine 堆栈跟踪。
- `GOTRACEBACK=single`（默认）打印当前 goroutine 的堆栈跟踪，省略运行时系统内部的函数。
  如果当前没有 goroutine，则失败会打印所有 goroutine 的堆栈跟踪或者故障是运行时内部的。
- `GOTRACEBACK=all` 为所有用户创建的 goroutine 添加堆栈跟踪。
- `GOTRACEBACK=system` 类似于 `all`，但为运行时函数添加了堆栈帧并显示由运行时内部创建的 goroutine。

## Recover and catch a panic

## 恢复并陷入恐慌

The built-in `recover` function can be used to regain control of a panicking goroutine and resume normal execution.

内置的“recover”功能可用于恢复控制恐慌的 goroutine 并恢复正常执行。

- A call to`recover` stops the unwinding and returns the argument passed to `panic`.
- If the goroutine is not panicking, `recover` returns `nil`. 
   
- 调用`recover` 停止展开并返回传递给 `panic` 的参数。
- 如果 goroutine 没有恐慌，`recover` 返回 `nil`。

Because the only code that runs while unwinding is inside [deferred functions](http://yourbasic.org/golang/defer/), `recover` is only useful inside such functions.

因为在展开时运行的唯一代码是在 [延迟函数](http://yourbasic.org/golang/defer/) 中，`recover` 仅在此类函数中有用。

### Panic handler example

### 恐慌处理程序示例

```
func main() {
    n := foo()
    fmt.Println("main received", n)
}

func foo() int {
    defer func() {
        if err := recover();err != nil {
            fmt.Println(err)
        }
    }()
    m := 1
    panic("foo: fail")
    m = 2
    return m
}
```


```
foo: fail
main received 0
```


Since the panic occurred before `foo` returned a value,`n` still has its initial zero value.

由于恐慌发生在 `foo` 返回值之前，`n` 仍然具有其初始零值。

### Return a value

### 返回一个值

To return a value during a panic, you must use a [named return value](http://yourbasic.org/golang/named-return-values-parameters/).

要在恐慌期间返回值，您必须使用 [命名返回值](http://yourbasic.org/golang/named-return-values-parameters/)。

```
func main() {
    n := foo()
    fmt.Println("main received", n)
}

func foo() (m int) {
    defer func() {
        if err := recover();err != nil {
            fmt.Println(err)
            m = 2
        }
    }()
    m = 1
    panic("foo: fail")
    m = 3
    return m
}
```


```
foo: fail
main received 2
```


## Test a panic (utility function)

## 测试恐慌（实用程序）

In this example, we use reflection to check if a list of interface variables have types corre­sponding to the para­meters of a given function. If so, we call the function with those para­meters to check if there is a panic.

在这个例子中，我们使用反射来检查接口变量列表具有与给定函数的参数对应的类型。如果是这样，我们会调用带有这些参数的函数来检查是否有恐慌。

```
// Panics tells if function f panics with parameters p.
func Panics(f interface{}, p ...interface{}) bool {
    fv := reflect.ValueOf(f)
    ft := reflect.TypeOf(f)
    if ft.NumIn() != len(p) {
        panic("wrong argument count")
    }
    pv := make([]reflect.Value, len(p))
    for i, v := range p {
        if reflect.TypeOf(v) != ft.In(i) {
            panic("wrong argument type")
        }
        pv[i] = reflect.ValueOf(v)
    }
    return call(fv, pv)
}

func call(fv reflect.Value, pv []reflect.Value) (b bool) {
    defer func() {
        if err := recover();err != nil {
            b = true
        }
    }()
    fv.Call(pv)
    return
}
```



