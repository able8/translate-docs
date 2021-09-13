# Intro to generics

# 泛型介绍

(At the time of writing) Go does not have support for user-defined generics, but [the proposal](https://blog.golang.org/generics-proposal) [has been accepted](https://github.com/golang/go/issues/43651#issuecomment-776944155) and will be included in version 1.18.

（在撰写本文时）Go 不支持用户定义的泛型，但 [提案](https://blog.golang.org/generics-proposal)[已被接受](https://github.com)。 com/golang/go/issues/43651#issuecomment-776944155) 并将包含在 1.18 版中。

However, there are ways to experiment with the upcoming implementation using the [go2go playground](https://go2goplay.golang.org/) _today_. So to work through this chapter you'll have to leave your precious editor of choice and instead do the work within the playground.

但是，有一些方法可以使用 [go2go playground](https://go2goplay.golang.org/) _today_ 来试验即将到来的实现。因此，要完成本章，您将不得不离开您选择的宝贵编辑器，而是在操场内完成工作。

This chapter will give you an introduction to generics, dispel reservations you may have about them and, give you an idea how to simplify some of your code in the future. After reading this you'll know how to write:

本章将向您介绍泛型，消除您可能对它们的保留意见，并让您了解如何在未来简化一些代码。阅读本文后，您将知道如何编写：

- A function that takes generic arguments
- A generic data-structure

- 一个接受泛型参数的函数
- 通用数据结构

## Setting up the playground

## 设置游乐场

In the _go2go playground_ we can't run `go test`. How are we going to write tests to explore generic code?

在 _go2go playground_ 中，我们无法运行 `go test`。我们将如何编写测试来探索通用代码？

The playground _does_ let us execute code, and because we're programmers that means we can work around the lack of a test runner by **making one of our own**.

操场_确实_让我们执行代码，因为我们是程序员，这意味着我们可以通过**制作我们自己的**来解决缺少测试运行程序的问题。

## Our own test helpers (`AssertEqual`, `AssertNotEqual`)

## 我们自己的测试助手 (`AssertEqual`, `AssertNotEqual`)

To explore generics we'll write some test helpers that'll kill the program and print something useful if a test fails.

为了探索泛型，我们将编写一些测试助手，它们会在测试失败时终止程序并打印一些有用的东西。

### Assert on integers

### 对整数断言

Let's start with something basic and iterate toward our goal

让我们从一些基本的东西开始，朝着我们的目标迭代

```go
package main

import (
    "log"
)

func main() {
    AssertEqual(1, 1)
    AssertNotEqual(1, 2)

    AssertEqual(50, 100) // this should fail

    AssertNotEqual(2, 2) // so you wont see this print
}

func AssertEqual(got, want int) {
    if got != want {
        log.Fatalf("FAILED: got %d, want %d", got, want)
    } else {
        log.Printf("PASSED: %d did equal %d\n", got, want)
    }
}

func AssertNotEqual(got, want int) {
    if got == want {
        log.Fatalf("FAILED: got %d, want %d", got, want)
    } else {
        log.Printf("PASSED: %d did not equal %d\n", got, want)
    }
}
```

[This program prints](https://go2goplay.golang.org/p/WywgJnAp34v):

 [本程序打印](https://go2goplay.golang.org/p/WywgJnAp34v)：

```
2009/11/10 23:00:00 PASSED: 1 did equal 1
2009/11/10 23:00:00 PASSED: 1 did not equal 2
2009/11/10 23:00:00 FAILED: got 50, want 100
```

### Assert on strings

### 对字符串进行断言

Being able to assert on the equality of integers is great but what if we want to assert on `string` ?

能够断言整数的相等性很好，但是如果我们想断言 `string` 呢？

```go
func main() {
    AssertEqual("CJ", "CJ")
}
```

You'll get an error

你会得到一个错误

```
type checking failed for main
prog.go2:8:14: cannot use "CJ" (untyped string constant) as int value in argument to AssertEqual
```

If you take your time to read the error, you'll see the compiler is complaining that we're trying to pass a `string` to a function that expects an `integer`.

如果你花时间阅读错误，你会看到编译器抱怨我们试图将一个 `string` 传递给一个需要 `integer` 的函数。

#### Recap on type-safety

#### 回顾类型安全

If you've read the previous chapters of this book, or have experience with statically typed languages, this should not surprise you. The Go compiler expects you to write your functions, structs etc by describing what types you wish to work with.

如果您已经阅读了本书的前几章，或者有使用静态类型语言的经验，这应该不会让您感到惊讶。 Go 编译器希望您通过描述您希望使用的类型来编写函数、结构等。

You can't pass a `string` to a function that expects an `integer`.

您不能将“字符串”传递给需要“整数”的函数。

Whilst this can feel like ceremony, it can be extremely helpful. By describing these constraints you,

虽然这感觉像是仪式，但它可能非常有帮助。通过描述这些限制，你，

- Make function implementation simpler. By describing to the compiler what types you work with, you **constrain the number of possible valid implementations**. You can't "add" a `Person` and a `BankAccount`. You can't capitalise an `integer`. In software, constraints are often extremely helpful.
- Are prevented from accidentally passing data to a function you didn't mean to.

- 使功能实现更简单。通过向编译器描述您使用的类型，您**限制了可能的有效实现的数量**。你不能“添加”一个 `Person` 和一个 `BankAccount`。您不能将“整数”大写。在软件中，约束通常非常有用。
- 防止意外将数据传递给您不希望传递的函数。

Go currently offers you a way to be more abstract with your types with interfaces, so that you can design functions that do not take concrete types but instead, types that offer the behaviour you need. This gives you some flexibility whilst maintaining type-safety.

Go 目前为您提供了一种使您的类型与接口更加抽象的方法，以便您可以设计不采用具体类型的函数，而是提供您需要的行为的类型。这为您提供了一些灵活性，同时保持了类型安全。

### A function that takes a string or an integer? (or indeed, other things)

### 一个接受字符串或整数的函数？ （或者实际上，其他事情）

The other option that Go _currently_ gives is declaring the type of your argument as `interface{}` which means "anything".

Go _currently_ 提供的另一个选项是将参数的类型声明为`interface{}`，意思是“任何东西”。

Try changing the signatures to use this type instead.

尝试更改签名以使用此类型。

```go
func AssertEqual(got, want interface{}) {

func AssertNotEqual(got, want interface{}) {

```

The tests should now compile and pass. The output will be a bit ropey because we're using the integer `%d` format string to print our messages, so change them to the general `%+v` format for a better output of any kind of value.

测试现在应该编译并通过。输出会有点乱，因为我们使用整数`%d` 格式字符串来打印我们的消息，所以将它们更改为通用的`%+v` 格式，以便更好地输出任何类型的值。

### Tradeoffs made without generics 

### 没有泛型的权衡

Our `AssertX` functions are quite naive but conceptually aren't too different to how other [popular libraries offer this functionality](https://github.com/matryer/is/blob/master/is.go#L150)

我们的 `AssertX` 函数非常幼稚，但在概念上与其他 [流行库提供此功能](https://github.com/matryer/is/blob/master/is.go#L150) 的方式没有太大不同

```go
func (is *I) Equal(a, b interface{}) {
```

So what's the problem?

所以有什么问题？

By using `interface{}` the compiler can't help us when writing our code, because we're not telling it anything useful about the types of things passed to the function. Go back to the _go2go playground_ and try comparing two different types,

通过使用`interface{}`，编译器在编写代码时无法帮助我们，因为我们没有告诉它任何关于传递给函数的事物类型的有用信息。回到 _go2go playground_ 并尝试比较两种不同的类型，

```go
AssertNotEqual(1, "1")
```

In this case, we get away with it; the test compiles, and it fails as we'd hope, although the error message `got 1, want 1` is unclear; but do we want to be able to compare strings with integers? What about comparing a `Person` with an `Airport`?

在这种情况下，我们侥幸逃脱；测试编译通过，但正如我们所希望的那样失败，尽管错误消息“得到 1，想要 1”不清楚；但是我们是否希望能够将字符串与整数进行比较？将“人”与“机场”进行比较怎么样？

Writing functions that take `interface{}` can be extremely challenging and bug-prone because we've _lost_ our constraints, and we have no information at compile time as to what kinds of data we're dealing with.

编写使用`interface{}` 的函数可能极具挑战性并且容易出错，因为我们已经_失去_了我们的约束，并且在编译时我们没有关于我们正在处理的数据类型的信息。

Often developers have to use reflection to implement these *ahem* generic functions, which is usually painful and can hurt the performance of your program.

通常，开发人员必须使用反射来实现这些*ahem* 泛型函数，这通常很痛苦并且会损害程序的性能。

## Our own test helpers with generics

## 我们自己的泛型测试助手

Ideally, we don't want to have to make specific `AssertX` functions for every type we ever deal with. We'd like to be able to have _one_ `AssertEqual` function that works with _any_ type but does not let you compare [apples and oranges](https://en.wikipedia.org/wiki/Apples_and_oranges).

理想情况下，我们不想为我们曾经处理过的每种类型都创建特定的 `AssertX` 函数。我们希望能够拥有 _one_ `AssertEqual` 函数，该函数适用于 _any_ 类型，但不允许您比较 [apples and oranges](https://en.wikipedia.org/wiki/Apples_and_oranges)。

Generics offer us a new way to make abstractions (like interfaces) by letting us **describe our constraints** in ways we cannot currently do.

泛型通过让我们以目前无法做到的方式**描述我们的约束**，为我们提供了一种创建抽象（如接口）的新方法。

```go
package main

import (
    "log"
)

func main() {
    AssertEqual(1, 1)
    AssertEqual("1", "1")
    AssertNotEqual(1, 2)
    //AssertEqual(1, "1") - uncomment me to see compilation error
}

func AssertEqual[T comparable](got, want T) {
    if got != want {
        log.Fatalf("FAILED: got %+v, want %+v", got, want)
    } else {
        log.Printf("PASSED: %+v did equal %+v\n", got, want)
    }
}

func AssertNotEqual[T comparable](got, want T) {
    if got == want {
        log.Fatalf("FAILED: got %+v, want %+v", got, want)
    } else {
        log.Printf("PASSED: %+v did not equal %+v\n", got, want)
    }
}
```

[go2go playground link](https://go2goplay.golang.org/p/a-6MzWrjeAx)

[go2go 游乐场链接](https://go2goplay.golang.org/p/a-6MzWrjeAx)

To write generic functions in Go, you need to provide "type parameters" which is just a fancy way of saying "describe your generic type and give it a label".

要在 Go 中编写泛型函数，您需要提供“类型参数”，这只是“描述您的泛型类型并给它一个标签”的一种奇特方式。

In our case the type of our type parameter is [`comparable`](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#comparable-types-in-constraints) and we've given it the label of `T`. This label then lets us describe the types for the arguments to our function (`got, want T`).

在我们的例子中，我们的类型参数的类型是 [`comparable`](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#comparable-types- in-constraints），我们给它贴上了“T”的标签。然后这个标签让我们描述函数参数的类型（`got, want T`)。

We're using `comparable` because we want to describe to the compiler that we wish to use the `==` and `!=` operators on things of type `T` in our function, we want to compare! If you try changing the type to `any`,

我们使用 `comparable` 是因为我们想向编译器描述我们希望在我们的函数中对类型为 `T` 的事物使用 `==` 和 `!=` 运算符，我们想要比较！如果您尝试将类型更改为 `any`，

```go
func AssertNotEqual[T any](got, want T) {
```

You'll get the following error:

你会得到以下错误：

```
prog.go2:15:5: cannot compare got != want (operator != not defined for T)
```

Which makes a lot of sense, because you can't use those operators on every (or `any`) type.

这很有意义，因为您不能在每种（或“任何”）类型上使用这些运算符。

### Is [`any`](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#the-constraint) the same as `interface{ }` ?

### [`any`](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#the-constraint) 与 `interface{ 相同}`？

Consider two functions

考虑两个函数

```go
func GenericFoo[T any](x, y T)
```

```go
func InterfaceyFoo(x, y interface{})
```

What's the point of generics here? Doesn't `any` describe... anything?

这里泛型的重点是什么？没有`any` 描述......什么吗？

In terms of constraints, `any` does mean "anything" and so does `interface{}`. The difference with the generic version is _you're still describing a specific type_ and what that means is we've still constrained this function to only work with _one_ type.

就约束而言，`any` 确实意味着“任何事情”，`interface{}` 也是如此。与通用版本的不同之处在于_您仍在描述特定类型_，这意味着我们仍然限制此函数仅适用于 _one_ 类型。

What this means is you can call `InterfaceyFoo` with any combination of types (e.g `InterfaceyFoo(apple, orange)`). However `GenericFoo` still offers some constraints because we've said that it only works with _one_ type, `T`.

这意味着您可以使用任意类型的组合调用“InterfaceyFoo”（例如“InterfaceyFoo(apple, orange)”）。然而 `GenericFoo` 仍然提供了一些限制，因为我们已经说过它只适用于 _one_ 类型，`T`。

Valid:

有效的：

- `GenericFoo(apple1, apple2)`
- `GenericFoo(orange1, orange2)`
- `GenericFoo(1, 2)`
- `GenericFoo("one", "two")`

-`GenericFoo(apple1, apple2)`
-`GenericFoo(orange1, orange2)`
-`GenericFoo(1, 2)`
- `GenericFoo("one", "two")`

Not valid (fails compilation):

无效（编译失败）：

- `GenericFoo(apple1, orange1)`
- `GenericFoo("1", 1)` 

-`GenericFoo(apple1, orange1)`
-`GenericFoo("1", 1)`

`any` is especially useful when making data types where you want it to work with various types, but you don't actually _use_ the type in your own data structure (typically you're just storing it). Things like, `Set` and `LinkedList`, are all good candidates for using `any`.

`any` 在创建数据类型时特别有用，您希望它可以与各种类型一起使用，但您实际上并没有_使用_您自己的数据结构中的类型（通常您只是存储它）。诸如“Set”和“LinkedList”之类的东西都是使用“any”的好选择。

## Next: Generic data types

## 下一个：通用数据类型

We're going to create a [stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)) data type. Stacks should be fairly straightforward to understand from a requirements point of view. They're a collection of items where you can `Push` items to the "top" and to get items back again you `Pop` items from the top (LIFO - last in, first out).

我们将创建一个 [stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)) 数据类型。从需求的角度来看，堆栈应该相当容易理解。它们是一组项目，您可以在其中“将”项目“推”到“顶部”，然后从顶部“弹出”项目（LIFO - 后进先出)再次取回项目。

For the sake of brevity I've omitted the TDD process that arrived me at the [following code](https://go2goplay.golang.org/p/HghXymv1OKm) for a stack of `int`s, and a stack of ` string`s.

为简洁起见，我省略了到达 [以下代码](https://go2goplay.golang.org/p/HghXymv1OKm) 的 TDD 过程，用于一堆 `int` 和一堆 `字符串的。

```go
package main

import (
    "log"
)

type StackOfInts struct {
    values []int
}

func (s *StackOfInts) Push(value int) {
    s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
    if s.IsEmpty() {
        return 0, false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}

type StackOfStrings struct {
    values []string
}

func (s *StackOfStrings) Push(value string) {
    s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
    if s.IsEmpty() {
        return "", false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}

func main() {
    // INT STACK

    myStackOfInts := new(StackOfInts)

    // check stack is empty
    AssertTrue(myStackOfInts.IsEmpty())

    // add a thing, then check it's not empty
    myStackOfInts.Push(123)
    AssertFalse(myStackOfInts.IsEmpty())

    // add another thing, pop it back again
    myStackOfInts.Push(456)
    value, _ := myStackOfInts.Pop()
    AssertEqual(value, 456)
    value, _ = myStackOfInts.Pop()
    AssertEqual(value, 123)
    AssertTrue(myStackOfInts.IsEmpty())

    // STRING STACK

    myStackOfStrings := new(StackOfStrings)

    // check stack is empty
    AssertTrue(myStackOfStrings.IsEmpty())

    // add a thing, then check it's not empty
    myStackOfStrings.Push("one two three")
    AssertFalse(myStackOfStrings.IsEmpty())

    // add another thing, pop it back again
    myStackOfStrings.Push("four five six")
    strValue, _ := myStackOfStrings.Pop()
    AssertEqual(strValue, "four five six")
    strValue, _ = myStackOfStrings.Pop()
    AssertEqual(strValue, "one two three")
    AssertTrue(myStackOfStrings.IsEmpty())
}

func AssertTrue(thing bool) {
    if thing {
        log.Printf("PASSED: Expected thing to be true and it was\n")
    } else {
        log.Fatalf("FAILED: expected true but got false")
    }
}

func AssertFalse(thing bool) {
    if !thing {
        log.Printf("PASSED: Expected thing to be false and it was\n")
    } else {
        log.Fatalf("FAILED: expected false but got true")
    }
}

func AssertEqual[T comparable](got, want T) {
    if got != want {
        log.Fatalf("FAILED: got %+v, want %+v", got, want)
    } else {
        log.Printf("PASSED: %+v did equal %+v\n", got, want)
    }
}

func AssertNotEqual[T comparable](got, want T) {
    if got == want {
        log.Fatalf("FAILED: got %+v, want %+v", got, want)
    } else {
        log.Printf("PASSED: %+v did not equal %+v\n", got, want)
    }
}
```

### Problems

###  问题

- The code for both `StackOfStrings` and `StackOfInts` is almost identical. Whilst duplication isn't always the end of the world, this doesn't feel great and does add an increased maintenance cost.
- As we're duplicating the logic across two types, we've had to duplicate the tests too.

- `StackOfStrings` 和 `StackOfInts` 的代码几乎相同。虽然重复并不总是世界末日，但这感觉并不好，并且确实增加了维护成本。
- 由于我们要在两种类型之间复制逻辑，因此我们也必须复制测试。

We really want to capture the _idea_ of a stack in one type, and have one set of tests for them. We should be wearing our refactoring hat right now which means we should not be changing the tests because we want to maintain the same behaviour.

我们真的想以一种类型捕获堆栈的 _idea_，并为它们进行一组测试。我们现在应该戴上重构的帽子，这意味着我们不应该更改测试，因为我们希望保持相同的行为。

Pre-generics, this is what we _could_ do

前泛型，这是我们_可以_做的

```go
type StackOfInts = Stack
type StackOfStrings = Stack

type Stack struct {
    values []interface{}
}

func (s *Stack) Push(value interface{}) {
    s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
    return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
    if s.IsEmpty() {
        var zero interface{}
        return zero, false
    }

    index := len(s.values) - 1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}
```

- We're aliasing our previous implementations of `StackOfInts` and `StackOfStrings` to a new unified type `Stack`
- We've removed the type safety from the `Stack` by making it so `values` is a [slice](https://github.com/quii/learn-go-with-tests/blob/main/arrays-and-slices.md) of `interface{}`

- 我们将之前的 `StackOfInts` 和 `StackOfStrings` 实现别名为一个新的统一类型 `Stack`
- 我们已经从 `Stack` 中移除了类型安全，让 `values` 成为一个 [slice](https://github.com/quii/learn-go-with-tests/blob/main/arrays-和-slices.md) 的`interface{}`

... And our tests still pass. Who needs generics?

...我们的测试仍然通过。谁需要泛型？

### The problem with throwing out type safety

### 抛出类型安全的问题

The first problem is the same as we saw with our `AssertEquals` - we've lost type safety. I can now `Push` apples onto a stack of oranges.

第一个问题与我们在 `AssertEquals` 中看到的一样——我们已经失去了类型安全性。我现在可以将苹果“推”到一堆橙子上。

Even if we have the discipline not to do this, the code is still unpleasant to work with because when methods **return `interface{}` they are horrible to work with**.

即使我们有纪律不这样做，代码仍然令人不愉快，因为当方法**返回`interface{}`时，它们很难使用**。

Add the following test,

添加以下测试，

```go
myStackOfInts.Push(1)
myStackOfInts.Push(2)
firstNum, _ := myStackOfInts.Pop()
secondNum, _ := myStackOfInts.Pop()
AssertEqual(firstNum+secondNum, 3)
```

You get a compiler error, showing the weakness of losing type-safety:

你得到一个编译器错误，显示了失去类型安全的弱点：

```go
prog.go2:59:14: invalid operation: operator + not defined for firstNum (variable of type interface{})
```

When `Pop` returns `interface{}` it means the compiler has no information about what the data is and therefore severely limits what we can do. It can't know that it should be an integer, so it does not let us use the `+` operator.

当 `Pop` 返回 `interface{}` 时，这意味着编译器没有关于数据是什么的信息，因此严重限制了我们可以做的事情。它不能知道它应该是一个整数，所以它不让我们使用 `+` 运算符。

To get around this, the caller has to do a [type assertion](https://golang.org/ref/spec#Type_assertions) for each value.

为了解决这个问题，调用者必须为每个值做一个 [类型断言](https://golang.org/ref/spec#Type_assertions)。

```go
myStackOfInts.Push(1)
myStackOfInts.Push(2)
firstNum, _ := myStackOfInts.Pop()
secondNum, _ := myStackOfInts.Pop()

// get our ints from out interface{}
reallyFirstNum, ok := firstNum.(int)
AssertTrue(ok) // need to check we definitely got an int out of the interface{}

reallySecondNum, ok := secondNum.(int)
AssertTrue(ok) // and again!

AssertEqual(reallyFirstNum+reallySecondNum, 3)
```

The unpleasantness radiating from this test would be repeated for every potential user of our `Stack` implementation, yuck.

对于我们的`Stack` 实现的每个潜在用户来说，这个测试带来的不愉快都会重复，哎呀。

### Generic data structures to the rescue

### 通用数据结构来救援

Just like you can define generic arguments to functions, you can define generic data structures.

就像您可以为函数定义通用参数一样，您也可以定义通用数据结构。

Here's our new `Stack` implementation, featuring a generic data type and the tests, showing them working how we'd like them to work, with full type-safety. ([Full code listing here](https://go2goplay.golang.org/p/xAWcaMelgQV))

这是我们新的 `Stack` 实现，具有通用数据类型和测试，向它们展示我们希望它们如何工作，并具有完整的类型安全性。 （[此处列出完整代码](https://go2goplay.golang.org/p/xAWcaMelgQV))

```go
package main

import (
    "log"
)

type Stack[T any] struct {
    values []T
}

func (s *Stack[T]) Push(value T) {
    s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.values)==0
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }

    index := len(s.values) -1
    el := s.values[index]
    s.values = s.values[:index]
    return el, true
}

func main() {
    myStackOfInts := new(Stack[int])

    // check stack is empty
    AssertTrue(myStackOfInts.IsEmpty())

    // add a thing, then check it's not empty
    myStackOfInts.Push(123)
    AssertFalse(myStackOfInts.IsEmpty())

    // add another thing, pop it back again
    myStackOfInts.Push(456)
    value, _ := myStackOfInts.Pop()
    AssertEqual(value, 456)
    value, _ = myStackOfInts.Pop()
    AssertEqual(value, 123)
    AssertTrue(myStackOfInts.IsEmpty())

    // can get the numbers we put in as numbers, not untyped interface{}
    myStackOfInts.Push(1)
    myStackOfInts.Push(2)
    firstNum, _ := myStackOfInts.Pop()
    secondNum, _ := myStackOfInts.Pop()
    AssertEqual(firstNum+secondNum, 3)
}
```

You'll notice the syntax for defining generic data structures is consistent with defining generic arguments to functions.

您会注意到定义通用数据结构的语法与定义函数的通用参数一致。

```go
type Stack[T any] struct {
    values []T
}
```

It's _almost_ the same as before, it's just that what we're saying is the **type of the stack constrains what type of values you can work with**.

它_几乎_和以前一样，只是我们所说的是**堆栈的类型限制了您可以使用的值的类型**。

Once you create a `Stack[Orange]` or a `Stack[Apple]` the methods defined on our stack will only let you pass in and will only return the particular type of the stack you're working with:

一旦你创建了一个 `Stack[Orange]` 或一个 `Stack[Apple]`，在我们的堆栈上定义的方法只会让你传入并且只会返回你正在使用的堆栈的特定类型：

```go
func (s *Stack[T]) Pop() (T, bool) {
```

You can imagine the types of implementation being somehow generated for you, depending on what type of stack you create:

您可以想象以某种方式为您生成的实现类型，具体取决于您创建的堆栈类型：

```go
func (s *Stack[Orange]) Pop() (Orange, bool) {
```

```go
func (s *Stack[Apple]) Pop() (Apple, bool) {
```

Now that we have done this refactoring, we can safely remove the string stack test because we don't need to prove the same logic over and over.

现在我们已经完成了这个重构，我们可以安全地移除字符串堆栈测试，因为我们不需要一遍又一遍地证明相同的逻辑。

Using a generic data type we have:

使用通用数据类型，我们有：

- Reduced duplication of important logic. 

- 减少重要逻辑的重复。

- Made `Pop` return `T` so that if we create a `Stack[int]` we in practice get back `int` from `Pop`; we can now use `+` without the need for type assertion gymnastics.
- Prevented misuse at compile time. You cannot `Push` oranges to an apple stack.

- 使 `Pop` 返回 `T`，这样如果我们创建一个 `Stack[int]`，我们实际上会从 `Pop` 返回 `int`；我们现在可以在不需要类型断言的情况下使用 `+`。
- 在编译时防止误用。您不能将橙子“推”到苹果堆中。

## Wrapping up

##  总结

This chapter should have given you a taste of generics syntax, and some ideas as to why generics might be helpful. We've written our own `Assert` functions which we can safely re-use to experiment with other ideas around generics, and we've implemented a simple data structure to store any type of data we wish, in a type-safe manner.

本章应该让您体验一下泛型语法，以及一些关于为什么泛型可能有用的想法。我们已经编写了我们自己的 `Assert` 函数，我们可以安全地重复使用这些函数来试验关于泛型的其他想法，并且我们已经实现了一个简单的数据结构来以类型安全的方式存储我们希望的任何类型的数据。

### Generics are simpler than using `interface{}` in most cases

### 泛型在大多数情况下比使用 `interface{}` 更简单

If you're inexperienced with statically-typed languages, the point of generics may not be immediately obvious, but I hope the examples in this chapter have illustrated where the Go language isn't as expressive as we'd like. In particular using `interface{}` makes your code:

如果您对静态类型语言没有经验，那么泛型的意义可能不是很明显，但我希望本章中的示例已经说明了 Go 语言不像我们希望的那样具有表现力。特别是使用`interface{}` 可以让你的代码：

- Less safe (mix apples and oranges), requires more error handling
- Less expressive, `interface{}` tells you nothing about the data
- More likely to rely on [reflection](https://github.com/quii/learn-go-with-tests/blob/main/reflection.md), type-assertions etc which makes your code more difficult to work with and more error prone as it pushes checks from compile-time to runtime

- 不太安全（混合苹果和橙子），需要更多的错误处理
- 表达能力较差，`interface{}` 不会告诉您有关数据的任何信息
- 更可能依赖 [reflection](https://github.com/quii/learn-go-with-tests/blob/main/reflection.md)、类型断言等，这会使您的代码更难使用并且更容易出错，因为它将检查从编译时推送到运行时

Using statically typed languages is an act of describing constraints. If you do it well, you create code that is not only safe and simple to use but also simpler to write because the possible solution space is smaller.

使用静态类型语言是一种描述约束的行为。如果你做得好，你创建的代码不仅安全且易于使用，而且编写起来也更简单，因为可能的解决方案空间更小。

Generics gives us a new way to express constraints in our code, which as demonstrated will allow us to consolidate and simplify code that is not possible to do today.

泛型为我们提供了一种在代码中表达约束的新方法，正如演示的那样，这将使我们能够整合和简化今天无法做到的代码。

### Will generics turn Go into Java?

### 泛型会变成 Java 吗？

- No.

- 不。

There's a lot of [FUD (fear, uncertainty and doubt)](https://en.wikipedia.org/wiki/Fear,_uncertainty,_and_doubt) in the Go community about generics leading to nightmare abstractions and baffling code bases. This is usually caveatted with "they must be used carefully".

Go 社区中有很多 [FUD（恐惧、不确定和怀疑）](https://en.wikipedia.org/wiki/Fear,_uncertainty,_and_doubt) 关于泛型导致噩梦般的抽象和令人困惑的代码库。这通常带有“必须小心使用”的警告。

Whilst this is true, it's not especially useful advice because this is true of any language feature.

虽然这是真的，但它并不是特别有用的建议，因为任何语言功能都是如此。

Not many people complain about our ability to define interfaces which, like generics is a way of describing constraints within our code. When you describe an interface you are making a design choice that _could be poor_, generics are not unique in their ability to make confusing, annoying to use code.

没有多少人抱怨我们定义接口的能力，就像泛型一样，它是在我们的代码中描述约束的一种方式。当你描述一个接口时，你正在做出一个_可能很差_的设计选择，泛型并不是唯一的，因为它们能够使使用代码变得混乱、烦人。

### You're already using generics

### 你已经在使用泛型了

When you consider that if you've used arrays, slices or maps; you've _already been a consumer of generic code_.

当您考虑使用数组、切片或映射时；你_已经是通用代码的消费者_了。

```go
var myApples []Apples
// You cant do this!
append(myApples, Orange{})
```

### Abstraction is not a dirty word

### 抽象不是一个肮脏的词

It's easy to dunk on [AbstractSingletonProxyFactoryBean](https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/aop/framework/AbstractSingletonProxyFactoryBean.html) but let's not pretend a code base with no abstraction at all isn't also bad. It's your job to _gather_ related concepts when appropriate, so your system is easier to understand and change; rather than being a collection of disparate functions and types with a lack of clarity.

在 [AbstractSingletonProxyFactoryBean](https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/aop/framework/AbstractSingletonProxyFactoryBean.html) 上扣篮很容易，但我们不要假装一个代码库完全没有抽象也不错。在适当的时候_收集_相关概念是你的工作，这样你的系统就更容易理解和改变；而不是缺乏清晰度的不同功能和类型的集合。

### [Make it work, make it right, make it fast](https://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast#:~:text=%22Make%20it%20work%2C%20make%20it,to%20DesignForPerformance%20ahead%20of%20time.)

### [让它工作，让它正确，让它快速](https://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast#:~:text=%22Make%20it%20work%2C%20make%20it,to%20DesignForPerformance%20ahead%20of%20time。)

People run in to problems with generics when they're abstracting too quickly without enough information to make good design decisions.

当人们在没有足够的信息来做出良好的设计决策的情况下过快地抽象时，他们会遇到泛型问题。

The TDD cycle of red, green, refactor means that you have more guidance as to what code you _actually need_ to deliver your behaviour, **rather than imagining abstractions up front**; but you still need to be careful. 

红色、绿色、重构的 TDD 循环意味着你有更多关于你_实际需要_来交付你的行为的代码的指导，**而不是预先想象抽象**；但你仍然需要小心。

There's no hard and fast rules here but resist making things generic until you can see that you have a useful generalisation. When we created the various `Stack` implementations we importantly started with _concrete_ behaviour like `StackOfStrings` and `StackOfInts` backed by tests. From our _real_ code we could start to see real patterns, and backed by our tests, we could explore refactoring toward a more general-purpose solution.

这里没有硬性规定，但不要把事情变得通用，直到你能看到你有一个有用的概括。当我们创建各种 `Stack` 实现时，我们重要的是从 _concrete_ 行为开始，例如由测试支持的 `StackOfStrings` 和 `StackOfInts`。从我们的 _real_ 代码中，我们可以开始看到真实的模式，在我们的测试的支持下，我们可以探索重构以实现更通用的解决方案。

People often advise you to only generalise when you see the same code three times, which seems like a good starting rule of thumb.

人们经常建议您只在看到相同代码 3 次时才进行概括，这似乎是一个很好的入门经验法则。

A common path I've taken in other programming languages has been:

我在其他编程语言中采用的常见路径是：

- One TDD cycle to drive some behaviour
- Another TDD cycle to exercise some other related scenarios

- 一个 TDD 周期来驱动一些行为
- 另一个 TDD 周期来演练一些其他相关场景

> Hmm, these things look similar - but a little duplication is better than coupling to a bad abstraction

> 嗯，这些东西看起来很相似——但是一点点重复总比耦合到一个糟糕的抽象要好

- Sleep on it
- Another TDD cycle

- 睡在上面
- 另一个 TDD 周期

> OK, I'd like to try to see if I can generalise this thing. Thank goodness I am so smart and good-looking because I use TDD, so I can refactor whenever I wish, and the process has helped me understand what behaviour I actually need before designing too much.

> 好的，我想试试看能不能概括这个东西。谢天谢地，我这么聪明好看，因为我用 TDD，所以我可以随时重构，这个过程帮助我在设计太多之前了解我真正需要的行为。

- This abstraction feels nice! The tests are still passing, and the code is simpler
- I can now delete a number of tests, I've captured the _essence_ of the behaviour and removed unnecessary detail 

- 这种抽象感觉很好！测试还在通过，代码更简单
- 我现在可以删除一些测试，我已经捕获了行为的本质并删除了不必要的细节

