# Iteration

# 迭代

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/for)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/for)**

To do stuff repeatedly in Go, you'll need `for`. In Go there are no `while`, `do`, `until` keywords, you can only use `for`. Which is a good thing!

要在 Go 中重复做一些事情，你需要 `for`。在 Go 中没有 `while`、`do`、`until` 关键字，你只能使用 `for`。这是一件好事！

Let's write a test for a function that repeats a character 5 times.

让我们为一个重复字符 5 次的函数编写一个测试。

There's nothing new so far, so try and write it yourself for practice.

到目前为止没有什么新东西，所以尝试自己编写以进行练习。

## Write the test first

## 先写测试

```go
package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}
```

## Try and run the test

## 尝试并运行测试

`./repeat_test.go:6:14: undefined: Repeat`



## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

_Keep the discipline!_ You don't need to know anything new right now to make the test fail properly.

_保持纪律！_ 你现在不需要知道任何新的东西来使测试正确地失败。

All you need to do right now is enough to make it compile so you can check your test is written well.

你现在需要做的就是让它编译，这样你就可以检查你的测试是否写得很好。

```go
package iteration

func Repeat(character string) string {
    return ""
}
```

Isn't it nice to know you already know enough Go to write tests for some basic problems? This means you can now play with the production code as much as you like and know it's behaving as you'd hope.

知道你已经足够了解 Go 来为一些基本问题编写测试，这不是很好吗？这意味着您现在可以随心所欲地使用生产代码，并知道它的行为与您希望的一样。

`repeat_test.go:10: expected 'aaaaa' but got ''`

## Write enough code to make it pass

## 编写足够的代码使其通过

The `for` syntax is very unremarkable and follows most C-like languages.

`for` 语法非常不起眼，并且遵循大多数类 C 语言。

```go
func Repeat(character string) string {
    var repeated string
    for i := 0;i < 5;i++ {
        repeated = repeated + character
    }
    return repeated
}
```

Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces `{ }` are always required. You might wonder what is happening in the row

与 C、Java 或 JavaScript 等其他语言不同，for 语句的三个组件周围没有括号，并且始终需要大括号“{}”。你可能想知道这一行发生了什么

```go
     var repeated string
```

as we've been using `:=` so far to declare and initializing variables. However, `:=` is simply [short hand for both steps](https://gobyexample.com/variables). Here we are declaring a `string` variable only. Hence, the explicit version. We can also use `var` to declare functions, as we'll see later on.

因为到目前为止我们一直在使用 `:=` 来声明和初始化变量。但是，`:=` 只是 [两个步骤的简写](https://gobyexample.com/variables)。这里我们只声明了一个 `string` 变量。因此，显式版本。我们也可以使用 `var` 来声明函数，稍后我们会看到。

Run the test and it should pass.

运行测试，它应该通过。

Additional variants of the for loop are described [here](https://gobyexample.com/for).

[此处](https://gobyexample.com/for) 描述了 for 循环的其他变体。

## Refactor

## 重构

Now it's time to refactor and introduce another construct `+=` assignment operator.

现在是重构并引入另一个构造`+=`赋值运算符的时候了。

```go
const repeatCount = 5

func Repeat(character string) string {
    var repeated string
    for i := 0;i < repeatCount;i++ {
        repeated += character
    }
    return repeated
}
```

`+=` called _"the Add AND assignment operator"_, adds the right operand to the left operand and assigns the result to left operand. It works with other types like integers.

`+=` 称为_“加与赋值运算符”_，将右操作数与左操作数相加，并将结果分配给左操作数。它适用于其他类型，如整数。

### Benchmarking

### 基准测试

Writing [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks) in Go is another first-class feature of the language and it is very similar to writing tests.

在 Go 中编写 [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks) 是该语言的另一个一流特性，它与编写测试非常相似。

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0;i < b.N;i++ {
        Repeat("a")
    }
}
```

You'll see the code is very similar to a test.

您会看到代码与测试非常相似。

The `testing.B` gives you access to the cryptically named `b.N`.

`testing.B` 使您可以访问隐秘命名的 `b.N`。

When the benchmark code is executed, it runs `b.N` times and measures how long it takes.

执行基准代码时，它会运行 `b.N` 次并测量需要多长时间。

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

代码运行的次数对您来说无关紧要，框架将确定什么是“好”值，以便让您获得一些不错的结果。

To run the benchmarks do `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`)

要运行基准测试，请执行`go test -bench=.`（或者如果您在 Windows Powershell 中使用 `go test -bench="."`）

```text
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
```

What `136 ns/op` means is our function takes on average 136 nanoseconds to run \(on my computer\). Which is pretty ok! To test this it ran it 10000000 times.

`136 ns/op` 的意思是我们的函数运行\（在我的计算机上\）平均需要 136 纳秒。这很好！为了测试它，它运行了 10000000 次。

_NOTE_ by default Benchmarks are run sequentially.

_注意_ 默认情况下，基准测试按顺序运行。

## Practice exercises

## 练习练习

* Change the test so a caller can specify how many times the character is repeated and then fix the code
* Write `ExampleRepeat` to document your function
* Have a look through the [strings](https://golang.org/pkg/strings) package. Find functions you think could be useful and experiment with them by writing tests like we have here. Investing time learning the standard library will really pay off over time.

* 更改测试，以便调用者可以指定字符重复的次数，然后修复代码
* 编写 `ExampleRepeat` 来记录你的函数
* 查看 [strings](https://golang.org/pkg/strings) 包。找到您认为可能有用的函数，并通过编写我们在这里的测试来试验它们。花时间学习标准库会随着时间的推移真正得到回报。

## Wrapping up

##  总结

* More TDD practice
* Learned `for`
* Learned how to write benchmarks 

* 更多 TDD 实践
* 学习`for`
* 学习了如何编写基准

