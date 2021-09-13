# Structs, methods & interfaces

# 结构、方法和接口

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/structs)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/structs)**

Suppose that we need some geometry code to calculate the perimeter of a rectangle given a height and width. We can write a `Perimeter(width float64, height float64)` function, where `float64` is for floating-point numbers like `123.45`.

假设我们需要一些几何代码来计算给定高度和宽度的矩形的周长。我们可以编写一个 `Perimeter(width float64, height float64)` 函数，其中 `float64` 用于浮点数，如 `123.45`。

The TDD cycle should be pretty familiar to you by now.

现在您应该非常熟悉 TDD 循环了。

## Write the test first

## 先写测试

```go
func TestPerimeter(t *testing.T) {
    got := Perimeter(10.0, 10.0)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
```

Notice the new format string? The `f` is for our `float64` and the `.2` means print 2 decimal places.

注意到新的格式字符串了吗？ `f` 用于我们的 `float64`，`.2` 表示打印 2 个小数位。

## Try to run the test

## 尝试运行测试

`./shapes_test.go:6:9: undefined: Perimeter`



## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

```go
func Perimeter(width float64, height float64) float64 {
    return 0
}
```

Results in `shapes_test.go:10: got 0.00 want 40.00`.



## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}
```

So far, so easy. Now let's create a function called `Area(width, height float64)` which returns the area of a rectangle.

到目前为止，很容易。现在让我们创建一个名为“Area(width, height float64)”的函数，它返回矩形的面积。

Try to do it yourself, following the TDD cycle.

尝试按照 TDD 循环自己做。

You should end up with tests like this

你应该以这样的测试结束

```go
func TestPerimeter(t *testing.T) {
    got := Perimeter(10.0, 10.0)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    got := Area(12.0, 6.0)
    want := 72.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
```

And code like this

和这样的代码

```go
func Perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
    return width * height
}
```

## Refactor

## 重构

Our code does the job, but it doesn't contain anything explicit about rectangles. An unwary developer might try to supply the width and height of a triangle to these functions without realising they will return the wrong answer.

我们的代码完成了这项工作，但它不包含任何关于矩形的明确内容。粗心的开发人员可能会尝试为这些函数提供三角形的宽度和高度，而没有意识到它们会返回错误的答案。

We could just give the functions more specific names like `RectangleArea`. A neater solution is to define our own _type_ called `Rectangle` which encapsulates this concept for us.

我们可以给函数提供更具体的名称，例如“RectangleArea”。一个更简洁的解决方案是定义我们自己的 _type_ 称为“矩形”，它为我们封装了这个概念。

We can create a simple type using a **struct**. [A struct](https://golang.org/ref/spec#Struct_types) is just a named collection of fields where you can store data.

我们可以使用 **struct** 创建一个简单的类型。 [结构体](https://golang.org/ref/spec#Struct_types) 只是一个命名的字段集合，您可以在其中存储数据。

Declare a struct like this

声明一个这样的结构

```go
type Rectangle struct {
    Width float64
    Height float64
}
```

Now let's refactor the tests to use `Rectangle` instead of plain `float64`s.

现在让我们重构测试以使用 `Rectangle` 而不是普通的 `float64`s。

```go
func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    rectangle := Rectangle{12.0, 6.0}
    got := Area(rectangle)
    want := 72.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
```

Remember to run your tests before attempting to fix. The tests should show a helpful error like

请记住在尝试修复之前运行测试。测试应该显示一个有用的错误，如

```text
./shapes_test.go:7:18: not enough arguments in call to Perimeter
    have (Rectangle)
    want (float64, float64)
```

You can access the fields of a struct with the syntax of `myStruct.field`.

您可以使用 `myStruct.field` 的语法访问结构体的字段。

Change the two functions to fix the test.

更改两个函数以修复测试。

```go
func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.Width * rectangle.Height
}
```

I hope you'll agree that passing a `Rectangle` to a function conveys our intent more clearly, but there are more benefits of using structs that we will cover later.

我希望你会同意，将 `Rectangle` 传递给函数可以更清楚地传达我们的意图，但使用结构有更多好处，我们将在后面介绍。

Our next requirement is to write an `Area` function for circles.

我们的下一个要求是为圆编写一个 `Area` 函数。

## Write the test first

## 先写测试

```go
func TestArea(t *testing.T) {

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        got := Area(rectangle)
        want := 72.0

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        got := Area(circle)
        want := 314.1592653589793

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })

}
```

As you can see, the `f` has been replaced by `g`, with good reason. 

如您所见，`f` 已被`g` 取代，这是有充分理由的。

Use of `g` will print a more precise decimal number in the error message \([fmt options](https://golang.org/pkg/fmt/)\).
For example, using a radius of 1.5 in a circle area calculation, `f` would show `7.068583` whereas `g` would show `7.0685834705770345`.

使用 `g` 将在错误消息 \([fmt options](https://golang.org/pkg/fmt/)\) 中打印更精确的十进制数。
例如，在圆面积计算中使用半径 1.5，`f` 会显示 `7.068583` 而 `g` 会显示 `7.0685834705770345`。

## Try to run the test

## 尝试运行测试

`./shapes_test.go:28:13: undefined: Circle`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to define our `Circle` type.

我们需要定义我们的 `Circle` 类型。

```go
type Circle struct {
    Radius float64
}
```

Now try to run the tests again

现在尝试再次运行测试

`./shapes_test.go:29:14: cannot use circle (type Circle) as type Rectangle in argument to Area`

`./shapes_test.go:29:14: 不能在 Area 的参数中使用圆（圆类型）作为矩形类型`

Some programming languages allow you to do something like this:

一些编程语言允许你做这样的事情：

```go
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```

But you cannot in Go

但是你不能在 Go 中

`./shapes.go:20:32: Area redeclared in this block`

`./shapes.go:20:32：在这个块中重新声明的区域`

We have two choices:

我们有两个选择：

* You can have functions with the same name declared in different _packages_. So we could create our `Area(Circle)` in a new package, but that feels overkill here.
* We can define [_methods_](https://golang.org/ref/spec#Method_declarations) on our newly defined types instead.

* 您可以在不同的_packages_ 中声明具有相同名称的函数。所以我们可以在一个新的包中创建我们的 `Area(Circle)`，但这在这里感觉有点过分了。
* 我们可以在新定义的类型上定义 [_methods_](https://golang.org/ref/spec#Method_declarations)。

### What are methods?

### 什么是方法？

So far we have only been writing _functions_ but we have been using some methods. When we call `t.Errorf` we are calling the method `Errorf` on the instance of our `t` \(`testing.T`\).

到目前为止，我们只编写了_functions_，但我们一直在使用一些方法。当我们调用 `t.Errorf` 时，我们正在调用我们的 `t` 实例上的方法 `Errorf` \(`testing.T`\)。

A method is a function with a receiver.
A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

方法是带有接收器的函数。
方法声明将标识符、方法名称绑定到方法，并将方法与接收者的基类型相关联。

Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on "things".

方法与函数非常相似，但它们是通过在特定类型的实例上调用它们来调用的。你可以在任何你喜欢的地方调用函数，比如`Area(rectangle)`，你只能在“事物”上调用方法。

An example will help so let's change our tests first to call methods instead and then fix the code.

一个例子会有所帮助，所以让我们先改变我们的测试，改为调用方法，然后修复代码。

```go
func TestArea(t *testing.T) {

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        got := rectangle.Area()
        want := 72.0

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        got := circle.Area()
        want := 314.1592653589793

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })

}
```

If we try to run the tests, we get

如果我们尝试运行测试，我们会得到

```text
./shapes_test.go:19:19: rectangle.Area undefined (type Rectangle has no field or method Area)
./shapes_test.go:29:16: circle.Area undefined (type Circle has no field or method Area)
```

> type Circle has no field or method Area

> 类型 Circle 没有字段或方法 Area

I would like to reiterate how great the compiler is here. It is so important to take the time to slowly read the error messages you get, it will help you in the long run.

我想重申编译器在这里有多棒。花时间慢慢阅读您收到的错误消息非常重要，从长远来看它会对您有所帮助。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Let's add some methods to our types

让我们为我们的类型添加一些方法

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64  {
    return 0
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64  {
    return 0
}
```

The syntax for declaring methods is almost the same as functions and that's because they're so similar. The only difference is the syntax of the method receiver `func (receiverName ReceiverType) MethodName(args)`.

声明方法的语法几乎与函数相同，这是因为它们非常相似。唯一的区别是方法接收器`func (receiverName ReceiverType) MethodName(args)` 的语法。

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. In many other programming languages this is done implicitly and you access the receiver via `this`.

当您的方法在该类型的变量上调用时，您可以通过 `receiverName` 变量获得对其数据的引用。在许多其他编程语言中，这是隐式完成的，您可以通过“this”访问接收器。

It is a convention in Go to have the receiver variable be the first letter of the type.

Go 中的约定是将接收者变量作为类型的第一个字母。

```go
r Rectangle
```

If you try to re-run the tests they should now compile and give you some failing output.

如果您尝试重新运行测试，它们现在应该编译并为您提供一些失败的输出。

## Write enough code to make it pass

## 编写足够的代码使其通过

Now let's make our rectangle tests pass by fixing our new method

现在让我们通过修复我们的新方法来让我们的矩形测试通过

```go
func (r Rectangle) Area() float64  {
    return r.Width * r.Height
}
```

If you re-run the tests the rectangle tests should be passing but circle should still be failing.

如果您重新运行测试，矩形测试应该会通过，但圆应该仍然会失败。

To make circle's `Area` function pass we will borrow the `Pi` constant from the `math` package \(remember to import it\).

为了使 circle 的 `Area` 函数通过，我们将从 `math` 包中借用 `Pi` 常量 \（记得导入它\）。

```go
func (c Circle) Area() float64  {
    return math.Pi * c.Radius * c.Radius
}
```

## Refactor

## 重构

There is some duplication in our tests.

我们的测试中有一些重复。

All we want to do is take a collection of _shapes_, call the `Area()` method on them and then check the result. 

我们要做的就是获取_shapes_的集合，对它们调用`Area()`方法，然后检查结果。

We want to be able to write some kind of `checkArea` function that we can pass both `Rectangle`s and `Circle`s to, but fail to compile if we try to pass in something that isn't a shape.

我们希望能够编写某种`checkArea` 函数，我们可以将`Rectangle`s 和`Circle`s 都传递给它，但是如果我们试图传递一些不是形状的东西，则编译失败。

With Go, we can codify this intent with **interfaces**.

使用 Go，我们可以用 **interfaces** 来编码这个意图。

[Interfaces](https://golang.org/ref/spec#Interface_types) are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

[接口](https://golang.org/ref/spec#Interface_types) 是 Go 等静态类型语言中一个非常强大的概念，因为它们允许您创建可用于不同类型的函数并创建高度解耦的代码，同时仍然保持类型安全。

Let's introduce this by refactoring our tests.

让我们通过重构我们的测试来介绍这一点。

```go
func TestArea(t *testing.T) {

    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    }

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        checkArea(t, rectangle, 72.0)
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        checkArea(t, circle, 314.1592653589793)
    })

}
```

We are creating a helper function like we have in other exercises but this time we are asking for a `Shape` to be passed in. If we try to call this with something that isn't a shape, then it will not compile.

我们正在创建一个辅助函数，就像我们在其他练习中所做的那样，但这次我们要求传入一个“Shape”。如果我们尝试使用不是形状的东西来调用它，那么它将无法编译。

How does something become a shape? We just tell Go what a `Shape` is using an interface declaration

东西是如何变成形状的？我们只是通过接口声明告诉 Go 什么是“Shape”

```go
type Shape interface {
    Area() float64
}
```

We're creating a new `type` just like we did with `Rectangle` and `Circle` but this time it is an `interface` rather than a `struct`.

我们正在创建一个新的 `type`，就像我们对 `Rectangle` 和 `Circle` 所做的一样，但这次它是一个 `interface` 而不是 `struct`。

Once you add this to the code, the tests will pass.

将其添加到代码中后，测试将通过。

### Wait, what?

### 等等，什么？

This is quite different to interfaces in most other programming languages. Normally you have to write code to say `My type Foo implements interface Bar`.

这与大多数其他编程语言中的接口完全不同。通常，您必须编写代码来说明“我的类型 Foo 实现了接口 Bar”。

But in our case

但在我们的情况下

* `Rectangle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface
* `Circle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface
* `string` does not have such a method, so it doesn't satisfy the interface
* etc.

* `Rectangle` 有一个名为 `Area` 的方法，它返回一个 `float64`，因此它满足 `Shape` 接口
* `Circle` 有一个名为 `Area` 的方法，它返回一个 `float64`，因此它满足 `Shape` 接口
* `string` 没有这样的方法，所以不满足接口
* 等等。

In Go **interface resolution is implicit**. If the type you pass in matches what the interface is asking for, it will compile.

在 Go 中**接口解析是隐式的**。如果您传入的类型与接口要求的类型相匹配，它将编译。

### Decoupling

### 解耦

Notice how our helper does not need to concern itself with whether the shape is a `Rectangle` or a `Circle` or a `Triangle`. By declaring an interface, the helper is _decoupled_ from the concrete types and only has the method it needs to do its job.

注意我们的助手不需要关心形状是“矩形”还是“圆形”或“三角形”。通过声明一个接口，帮助器从具体类型中_decoupled_，并且只有完成其工作所需的方法。

This kind of approach of using interfaces to declare **only what you need** is very important in software design and will be covered in more detail in later sections.

这种使用接口声明**只需要的东西**的方法在软件设计中非常重要，将在后面的章节中详细介绍。

## Further refactoring

## 进一步重构

Now that you have some understanding of structs we can introduce "table driven tests".

现在您对结构有了一些了解，我们可以介绍“表驱动测试”。

[Table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) are useful when you want to build a list of test cases that can be tested in the same manner.

[表驱动测试](https://github.com/golang/go/wiki/TableDrivenTests) 当您想构建可以以相同方式进行测试的测试用例列表时非常有用。

```go
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
```

The only new syntax here is creating an "anonymous struct", `areaTests`. We are declaring a slice of structs by using `[]struct` with two fields, the `shape` and the `want`. Then we fill the slice with cases.

这里唯一的新语法是创建一个“匿名结构”，`areaTests`。我们通过使用带有两个字段的 `[]struct` 来声明一个结构体，`shape` 和 `want`。然后我们用案例填充切片。

We then iterate over them just like we do any other slice, using the struct fields to run our tests.

然后我们像对任何其他切片一样迭代它们，使用结构字段来运行我们的测试。

You can see how it would be very easy for a developer to introduce a new shape, implement `Area` and then add it to the test cases. In addition, if a bug is found with `Area` it is very easy to add a new test case to exercise it before fixing it.

您可以看到开发人员引入新形状、实现“区域”然后将其添加到测试用例中是多么容易。此外，如果在 `Area` 中发现了一个错误，在修复它之前很容易添加一个新的测试用例来练习它。

Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests.
They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.

表驱动测试可能是您工具箱中的重要项目，但请确保您需要测试中的额外噪音。
当您希望测试接口的各种实现时，或者如果传递给函数的数据有许多需要测试的不同需求，它们非常适合。

Let's demonstrate all this by adding another shape and testing it; a triangle.

让我们通过添加另一个形状并对其进行测试来演示这一切；一个三角形。

## Write the test first

## 先写测试

Adding a new test for our new shape is very easy. Just add `{Triangle{12, 6}, 36.0},` to our list.

为我们的新形状添加新测试非常简单。只需将 `{Triangle{12, 6}, 36.0},` 添加到我们的列表中。

```go
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
        {Triangle{12, 6}, 36.0},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
```

## Try to run the test

## 尝试运行测试

Remember, keep trying to run the test and let the compiler guide you toward a solution.

请记住，继续尝试运行测试并让编译器引导您找到解决方案。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

`./shapes_test.go:25:4: undefined: Triangle`

We have not defined `Triangle` yet

我们还没有定义`Triangle`

```go
type Triangle struct {
    Base   float64
    Height float64
}
```

Try again

再试一次

```text
./shapes_test.go:25:8: cannot use Triangle literal (type Triangle) as type Shape in field value:
    Triangle does not implement Shape (missing Area method)
```

It's telling us we cannot use a `Triangle` as a shape because it does not have an `Area()` method, so add an empty implementation to get the test working

它告诉我们我们不能使用 `Triangle` 作为形状，因为它没有 `Area()` 方法，所以添加一个空的实现来让测试工作

```go
func (t Triangle) Area() float64 {
    return 0
}
```

Finally the code compiles and we get our error

最后代码编译，我们得到我们的错误

`shapes_test.go:31: got 0.00 want 36.00`

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (t Triangle) Area() float64 {
    return (t.Base * t.Height) * 0.5
}
```

And our tests pass!

我们的测试通过了！

## Refactor

## 重构

Again, the implementation is fine but our tests could do with some improvement.

同样，实现很好，但我们的测试可以做一些改进。

When you scan this

当你扫描这个

```go
{Rectangle{12, 6}, 72.0},
{Circle{10}, 314.1592653589793},
{Triangle{12, 6}, 36.0},
```

It's not immediately clear what all the numbers represent and you should be aiming for your tests to be easily understood.

目前还不清楚所有数字代表什么，您的目标应该是让您的测试易于理解。

So far you've only been shown syntax for creating instances of structs `MyStruct{val1, val2}` but you can optionally name the fields.

到目前为止，您只看到了创建结构体实例`MyStruct{val1, val2}` 的语法，但您可以选择命名字段。

Let's see what it looks like

让我们看看它长什么样

```go
         {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
        {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
```

In [Test-Driven Development by Example](https://g.co/kgs/yCzDLF) Kent Beck refactors some tests to a point and asserts:

在 [Test-Driven Development by Example](https://g.co/kgs/yCzDLF) 中，Kent Beck 重构了一些测试并断言：

> The test speaks to us more clearly, as if it were an assertion of truth, **not a sequence of operations**

> 测试更清楚地告诉我们，好像它是对真理的断言，**不是一系列操作**

Now our tests - rather, the list of test cases - make assertions of truth about shapes and their areas.

现在我们的测试——更确切地说，是测试用例列表——对形状及其区域的真实性做出断言。

## Make sure your test output is helpful

## 确保您的测试输出有用

Remember earlier when we were implementing `Triangle` and we had the failing test? It printed `shapes_test.go:31: got 0.00 want 36.00`.

还记得之前我们在实现 `Triangle` 并且测试失败时吗？它打印了“shapes_test.go:31: got 0.00 want 36.00”。

We knew this was in relation to `Triangle` because we were just working with it.
But what if a bug slipped in to the system in one of 20 cases in the table?
How would a developer know which case failed?
This is not a great experience for the developer, they will have to manually look through the cases to find out which case actually failed.

我们知道这与“三角形”有关，因为我们只是在使用它。
但是，如果在表中的 20 个案例之一中有一个错误滑入系统怎么办？
开发人员如何知道哪个案例失败了？
这对开发人员来说不是很好的体验，他们将不得不手动查看案例以找出实际失败的案例。

We can change our error message into `%#v got %g want %g`. The `%#v` format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.

我们可以将错误消息更改为 `%#v got %g want %g`。 `%#v` 格式字符串将打印出我们的结构及其字段中的值，因此开发人员可以一目了然地看到正在测试的属性。

To increase the readability of our test cases further, we can rename the `want` field into something more descriptive like `hasArea`.

为了进一步提高测试用例的可读性，我们可以将 `want` 字段重命名为更具描述性的内容，例如 `hasArea`。

One final tip with table driven tests is to use `t.Run` and to name the test cases.

表驱动测试的最后一个技巧是使用 `t.Run` 并命名测试用例。

By wrapping each case in a `t.Run` you will have clearer test output on failures as it will print the name of the case

通过将每个案例包装在一个 `t.Run` 中，您将获得更清晰的失败测试输出，因为它将打印案例的名称

```text
--- FAIL: TestArea (0.00s)
    --- FAIL: TestArea/Rectangle (0.00s)
        shapes_test.go:33: main.Rectangle{Width:12, Height:6} got 72.00 want 72.10
```

And you can run specific tests within your table with `go test -run TestArea/Rectangle`.

您可以使用 `go test -run TestArea/Rectangle` 在表中运行特定测试。

Here is our final test code which captures this

这是我们最终的测试代码，它捕获了这个

```go
func TestArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })

    }

}
```

## Wrapping up

##  总结

This was more TDD practice, iterating over our solutions to basic mathematic problems and learning new language features motivated by our tests.

这是更多的 TDD 实践，迭代我们对基本数学问题的解决方案，并学习由我们的测试激发的新语言特征。

* Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer 

* 声明结构以创建您自己的数据类型，这使您可以将相关数据捆绑在一起并使代码的意图更清晰

* Declaring interfaces so you can define functions that can be used by different types \([parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)\)
* Adding methods so you can add functionality to your data types and so you can implement interfaces
* Table driven tests to make your assertions clearer and your test suites easier to extend & maintain

* 声明接口，以便您可以定义可由不同类型使用的函数 \([参数多态性](https://en.wikipedia.org/wiki/Parametric_polymorphism)\)
* 添加方法，以便为数据类型添加功能，从而实现接口
* 表驱动测试使您的断言更清晰，并使您的测试套件更易于扩展和维护

This was an important chapter because we are now starting to define our own types. In statically typed languages like Go, being able to design your own types is essential for building software that is easy to understand, to piece together and to test.

这是重要的一章，因为我们现在开始定义我们自己的类型。在像 Go 这样的静态类型语言中，能够设计自己的类型对于构建易于理解、拼凑和测试的软件至关重要。

Interfaces are a great tool for hiding complexity away from other parts of the system. In our case our test helper _code_ did not need to know the exact shape it was asserting on, only how to "ask" for its area.

接口是一个很好的工具，可以将复杂性从系统的其他部分隐藏起来。在我们的例子中，我们的测试助手 _code_ 不需要知道它断言的确切形状，只需要如何“询问”它的面积。

As you become more familiar with Go you will start to see the real strength of interfaces and the standard library. You'll learn about interfaces defined in the standard library that are used _everywhere_ and by implementing them against your own types, you can very quickly re-use a lot of great functionality. 

随着您对 Go 越来越熟悉，您将开始看到接口和标准库的真正优势。您将了解标准库中定义的接口，这些接口_无处不在_，并且通过针对您自己的类型实现它们，您可以非常快速地重用许多强大的功能。

