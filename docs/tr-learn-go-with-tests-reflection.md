# Reflection

#  反射

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/reflection)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/reflection)**

[From Twitter](https://twitter.com/peterbourgon/status/1011403901419937792?s=09)

[来自推特](https://twitter.com/peterbourgon/status/1011403901419937792?s=09)

> golang challenge: write a function `walk(x interface{}, fn func(string))` which takes a struct `x` and calls `fn` for all strings fields found inside. difficulty level: recursively.

> golang 挑战：编写一个函数 `walk(x interface{}, fn func(string))`，它接受一个结构体 `x` 并为内部找到的所有字符串字段调用 `fn`。难度级别：递归。

To do this we will need to use _reflection_.

为此，我们需要使用 _reflection_。

> Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.

> 计算中的反射是程序检查自身结构的能力，尤其是通过类型；它是元编程的一种形式。这也是一个很大的混乱来源。

From [The Go Blog: Reflection](https://blog.golang.org/laws-of-reflection)

来自 [The Go 博客：反思](https://blog.golang.org/laws-of-reflection)

## What is `interface`?

## 什么是`接口`？

We have enjoyed the type-safety that Go has offered us in terms of functions that work with known types, such as `string`, `int` and our own types like `BankAccount`.

我们很享受 Go 为我们提供的类型安全性，它可以处理已知类型的函数，例如 `string`、`int` 和我们自己的类型，比如 `BankAccount`。

This means that we get some documentation for free and the compiler will complain if you try and pass the wrong type to a function.

这意味着我们可以免费获得一些文档，如果您尝试将错误的类型传递给函数，编译器会抱怨。

You may come across scenarios though where you want to write a function where you don't know the type at compile time.

尽管您想编写一个在编译时不知道类型的函数，但您可能会遇到这样的情况。

Go lets us get around this with the type `interface{}` which you can think of as just _any_ type.

Go 让我们使用类型 `interface{}` 来解决这个问题，你可以将其视为 _any_ 类型。

So `walk(x interface{}, fn func(string))` will accept any value for `x`.

所以 `walk(x interface{}, fn func(string))` 将接受 `x` 的任何值。

### So why not use `interface` for everything and have really flexible functions?

### 那么为什么不使用`interface` 来实现所有功能并拥有非常灵活的功能呢？

- As a user of a function that takes `interface` you lose type safety. What if you meant to pass `Foo.bar` of type `string` into a function but instead did `Foo.baz` which is an `int`? The compiler won't be able to inform you of your mistake. You also have no idea _what_ you're allowed to pass to a function. Knowing that a function takes a `UserService` for instance is very useful.
- As a writer of such a function, you have to be able to inspect _anything_ that has been passed to you and try and figure out what the type is and what you can do with it. This is done using _reflection_. This can be quite clumsy and difficult to read and is generally less performant (as you have to do checks at runtime).

- 作为使用`interface` 函数的用户，你失去了类型安全。如果您打算将类型为“string”的“Foo.bar”传递给一个函数，而将“Foo.baz”传递给一个“int”，该怎么办？编译器将无法通知您您的错误。你也不知道_what_你可以传递给一个函数。例如，知道一个函数需要一个 `UserService` 是非常有用的。
- 作为这样一个函数的作者，你必须能够检查传递给你的_任何东西_并尝试找出类型是什么以及你可以用它做什么。这是使用 _reflection_ 完成的。这可能非常笨拙且难以阅读，并且通常性能较低（因为您必须在运行时进行检查）。

In short only use reflection if you really need to.

简而言之，只有在确实需要时才使用反射。

If you want polymorphic functions, consider if you could design it around an interface (not `interface`, confusingly) so that users can use your function with multiple types if they implement whatever methods you need for your function to work.

如果您想要多态函数，请考虑是否可以围绕接口（而不是“接口”，令人困惑）设计它，以便用户在实现函数工作所需的任何方法时，可以将您的函数用于多种类型。

Our function will need to be able to work with lots of different things. As always we'll take an iterative approach, writing tests for each new thing we want to support and refactoring along the way until we're done.

我们的函数需要能够处理许多不同的事情。与往常一样，我们将采用迭代方法，为我们想要支持的每个新事物编写测试并在此过程中重构，直到我们完成。

## Write the test first

## 先写测试

We'll want to call our function with a struct that has a string field in it (`x`). Then we can spy on the function (`fn`) passed in to see if it is called.

我们将要使用一个结构体调用我们的函数，该结构体中包含一个字符串字段（`x`）。然后我们可以窥探传入的函数（`fn`）是否被调用。

```go
func TestWalk(t *testing.T) {

    expected := "Chris"
    var got []string

    x := struct {
        Name string
    }{expected}

    walk(x, func(input string) {
        got = append(got, input)
    })

    if len(got) != 1 {
        t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
    }
}
```

- We want to store a slice of strings (`got`) which stores which strings were passed into `fn` by `walk`. Often in previous chapters, we have made dedicated types for this to spy on function/method invocations but in this case, we can just pass in an anonymous function for `fn` that closes over `got`.
- We use an anonymous `struct` with a `Name` field of type string to go for the simplest "happy" path.
- Finally, call `walk` with `x` and the spy and for now just check the length of `got`, we'll be more specific with our assertions once we've got something very basic working.

- 我们想要存储一段字符串（`got`），其中存储了哪些字符串被 `walk` 传递给了 `fn`。通常在前面的章节中，我们已经为此创建了专用类型来监视函数/方法调用，但在这种情况下，我们可以只为在 `got` 上关闭的 `fn` 传入一个匿名函数。
- 我们使用匿名的 `struct` 和类型为 string 的 `Name` 字段来寻找最简单的“快乐”路径。
- 最后，使用 `x` 和 spy 调用 `walk`，现在只检查 `got` 的长度，一旦我们有了一些非常基本的工作，我们将更具体地使用我们的断言。

## Try to run the test

## 尝试运行测试

```
./reflection_test.go:21:2: undefined: walk
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to define `walk`

我们需要定义`walk`

```go
func walk(x interface{}, fn func(input string)) {

}
```

Try and run the test again

尝试并再次运行测试

```
=== RUN   TestWalk
--- FAIL: TestWalk (0.00s)
    reflection_test.go:19: wrong number of function calls, got 0 want 1
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We can call the spy with any string to make this pass.

我们可以使用任何字符串调用 spy 来完成此操作。

```go
func walk(x interface{}, fn func(input string)) {
    fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
}
```

The test should now be passing. The next thing we'll need to do is make a more specific assertion on what our `fn` is being called with.

测试现在应该通过了。接下来我们需要做的是对我们的 `fn` 被调用的内容做出更具体的断言。

## Write the test first

## 先写测试

Add the following to the existing test to check the string passed to `fn` is correct

将以下内容添加到现有测试中以检查传递给 `fn` 的字符串是否正确

```go
if got[0] != expected {
    t.Errorf("got %q, want %q", got[0], expected)
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk
--- FAIL: TestWalk (0.00s)
    reflection_test.go:23: got 'I still can't believe South Korea beat Germany 2-0 to put them last in their group', want 'Chris'
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    field := val.Field(0)
    fn(field.String())
}
```

This code is _very unsafe and very naive_, but remember: our goal when we are in "red" (the tests failing) is to write the smallest amount of code possible. We then write more tests to address our concerns.

这段代码_非常不安全且非常幼稚_，但请记住：当我们处于“红色”状态（测试失败）时，我们的目标是编写尽可能少的代码。然后我们编写更多测试来解决我们的担忧。

We need to use reflection to have a look at `x` and try and look at its properties.

我们需要使用反射来查看 `x` 并尝试查看它的属性。

The [reflect package](https://godoc.org/reflect) has a function `ValueOf` which returns us a `Value` of a given variable. This has ways for us to inspect a value, including its fields which we use on the next line.

[reflect 包](https://godoc.org/reflect) 有一个函数 `ValueOf`，它返回给定变量的 `Value`。这为我们提供了检查值的方法，包括我们在下一行使用的字段。

We then make some very optimistic assumptions about the value passed in

然后我们对传入的值做出一些非常乐观的假设

- We look at the first and only field, there may be no fields at all which would cause a panic
- We then call `String()` which returns the underlying value as a string but we know it would be wrong if the field was something other than a string.

- 我们看第一个也是唯一的字段，可能根本没有会引起恐慌的字段
- 然后我们调用`String()`，它将底层值作为字符串返回，但我们知道如果该字段不是字符串，那将是错误的。

## Refactor

## 重构

Our code is passing for the simple case but we know our code has a lot of shortcomings.

我们的代码通过了简单的情况，但我们知道我们的代码有很多缺点。

We're going to be writing a number of tests where we pass in different values and checking the array of strings that `fn` was called with.

我们将编写一些测试，在这些测试中我们传递不同的值并检查调用 `fn` 的字符串数组。

We should refactor our test into a table based test to make this easier to continue testing new scenarios.

我们应该将我们的测试重构为基于表的测试，以便更轻松地继续测试新场景。

```go
func TestWalk(t *testing.T) {

    cases := []struct{
        Name string
        Input interface{}
        ExpectedCalls []string
    } {
        {
            "Struct with one string field",
            struct {
                Name string
            }{ "Chris"},
            []string{"Chris"},
        },
    }

    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }
}
```

Now we can easily add a scenario to see what happens if we have more than one string field.

现在我们可以轻松添加一个场景，看看如果我们有多个字符串字段会发生什么。

## Write the test first

## 先写测试

Add the following scenario to the `cases`.

将以下场景添加到 `cases` 中。

```go
{
    "Struct with two string fields",
    struct {
        Name string
        City string
    }{"Chris", "London"},
    []string{"Chris", "London"},
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Struct_with_two_string_fields
    --- FAIL: TestWalk/Struct_with_two_string_fields (0.00s)
        reflection_test.go:40: got [Chris], want [Chris London]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i:=0;i<val.NumField();i++ {
        field := val.Field(i)
        fn(field.String())
    }
}
```

`val` has a method `NumField` which returns the number of fields in the value. This lets us iterate over the fields and call `fn` which passes our test.

`val` 有一个方法 `NumField`，它返回值中的字段数。这让我们可以遍历字段并调用通过我们的测试的 `fn`。

## Refactor

## 重构

It doesn't look like there's any obvious refactors here that would improve the code so let's press on.

看起来这里没有任何明显的重构可以改进代码，所以让我们继续。

The next shortcoming in `walk` is that it assumes every field is a `string`. Let's write a test for this scenario.

`walk` 的下一个缺点是它假设每个字段都是一个 `string`。让我们为这个场景编写一个测试。

## Write the test first

## 先写测试

Add the following case

添加以下案例

```go
{
    "Struct with non string field",
    struct {
        Name string
        Age  int
    }{"Chris", 33},
    []string{"Chris"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Struct_with_non_string_field
    --- FAIL: TestWalk/Struct_with_non_string_field (0.00s)
        reflection_test.go:46: got [Chris <int Value>], want [Chris]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We need to check that the type of the field is a `string`.

我们需要检查该字段的类型是否为“字符串”。

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        if field.Kind() == reflect.String {
            fn(field.String())
        }
    }
}
```

We can do that by checking its [`Kind`](https://godoc.org/reflect#Kind).

我们可以通过检查它的 [`Kind`](https://godoc.org/reflect#Kind) 来做到这一点。

## Refactor

## 重构

Again it looks like the code is reasonable enough for now.

再次看起来代码现在足够合理。

The next scenario is what if it isn't a "flat" `struct`? In other words, what happens if we have a `struct` with some nested fields?

下一个场景是如果它不是“平面”`struct`怎么办？换句话说，如果我们有一个带有一些嵌套字段的 `struct` 会发生什么？

## Write the test first

## 先写测试

We have been using the anonymous struct syntax to declare types ad-hocly for our tests so we could continue to do that like so

我们一直在使用匿名结构体语法为我们的测试特别声明类型，所以我们可以继续这样做

```go
{
    "Nested fields",
    struct {
        Name string
        Profile struct {
            Age  int
            City string
        }
    }{"Chris", struct {
        Age  int
        City string
    }{33, "London"}},
    []string{"Chris", "London"},
},
```

But we can see that when you get inner anonymous structs the syntax gets a little messy. [There is a proposal to make it so the syntax would be nicer](https://github.com/golang/go/issues/12854).

但是我们可以看到，当您获得内部匿名结构时，语法会变得有些混乱。 [有一个建议使它的语法更好](https://github.com/golang/go/issues/12854)。

Let's just refactor this by making a known type for this scenario and reference it in the test. There is a little indirection in that some of the code for our test is outside the test but readers should be able to infer the structure of the `struct` by looking at the initialisation.

让我们通过为此场景创建一个已知类型并在测试中引用它来重构它。有一点间接，因为我们测试的一些代码在测试之外，但读者应该能够通过查看初始化来推断 `struct` 的结构。

Add the following type declarations somewhere in your test file

在测试文件的某处添加以下类型声明

```go
type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age  int
    City string
}
```

Now we can add this to our cases which reads a lot clearer than before

现在我们可以将它添加到我们的案例中，这比以前更清晰

```go
{
    "Nested fields",
    Person{
        "Chris",
        Profile{33, "London"},
    },
    []string{"Chris", "London"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Nested_fields
    --- FAIL: TestWalk/Nested_fields (0.00s)
        reflection_test.go:54: got [Chris], want [Chris London]
```

The problem is we're only iterating on the fields on the first level of the type's hierarchy.

问题是我们只在类型层次结构的第一级上迭代字段。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        if field.Kind() == reflect.String {
            fn(field.String())
        }

        if field.Kind() == reflect.Struct {
            walk(field.Interface(), fn)
        }
    }
}
```

The solution is quite simple, we again inspect its `Kind` and if it happens to be a `struct` we just call `walk` again on that inner `struct`.

解决方案非常简单，我们再次检查它的 `Kind`，如果它恰好是一个 `struct`，我们只需在该内部 `struct` 上再次调用 `walk`。

## Refactor

## 重构

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)
        }
    }
}
```

When you're doing a comparison on the same value more than once _generally_ refactoring into a `switch` will improve readability and make your code easier to extend.

当你对同一个值进行多次比较时_通常_重构为一个`switch`将提高可读性并使你的代码更容易扩展。

What if the value of the struct passed in is a pointer?

如果传入的结构体的值是一个指针呢？

## Write the test first

## 先写测试

Add this case

添加此案例

```go
{
    "Pointers to things",
    &Person{
        "Chris",
        Profile{33, "London"},
    },
    []string{"Chris", "London"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Pointers_to_things
panic: reflect: call of reflect.Value.NumField on ptr Value [recovered]
    panic: reflect: call of reflect.Value.NumField on ptr Value
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)
        }
    }
}
```

You can't use `NumField` on a pointer `Value`, we need to extract the underlying value before we can do that by using `Elem()`.

你不能在指针 `Value` 上使用 `NumField`，我们需要在使用 `Elem()` 之前提取底层值。

## Refactor

## 重构

Let's encapsulate the responsibility of extracting the `reflect.Value` from a given `interface{}` into a function.

让我们把从给定的 `interface{}` 中提取 `reflect.Value` 的职责封装到一个函数中。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)
        }
    }
}

func getValue(x interface{}) reflect.Value {
    val := reflect.ValueOf(x)

    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    return val
}
```

This actually adds _more_ code but I feel the abstraction level is right.

这实际上添加了_more_ 代码，但我觉得抽象级别是正确的。

- Get the `reflect.Value` of `x` so I can inspect it, I don't care how.
- Iterate over the fields, doing whatever needs to be done depending on its type.

- 获取 `x` 的 `reflect.Value` 以便我可以检查它，我不在乎如何。
- 遍历字段，根据其类型做任何需要做的事情。

Next, we need to cover slices.

接下来，我们需要覆盖切片。

## Write the test first

## 先写测试

```go
{
    "Slices",
    []Profile {
        {33, "London"},
        {34, "Reykjavík"},
    },
    []string{"London", "Reykjavík"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Slices
panic: reflect: call of reflect.Value.NumField on slice Value [recovered]
    panic: reflect: call of reflect.Value.NumField on slice Value
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

This is similar to the pointer scenario before, we are trying to call `NumField` on our `reflect.Value` but it doesn't have one as it's not a struct.

这类似于之前的指针场景，我们试图在我们的 `reflect.Value` 上调用 `NumField`，但它没有，因为它不是一个结构体。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    if val.Kind() == reflect.Slice {
        for i:=0;i< val.Len();i++ {
            walk(val.Index(i).Interface(), fn)
        }
        return
    }

    for i := 0;i < val.NumField();i++ {
        field := val.Field(i)

        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)
        }
    }
}
```

## Refactor

## 重构

This works but it's yucky. No worries, we have working code backed by tests so we are free to tinker all we like.

这有效，但很恶心。不用担心，我们有由测试支持的工作代码，所以我们可以随意修改我们喜欢的所有内容。

If you think a little abstractly, we want to call `walk` on either

如果你想得有点抽象，我们想在任何一个上调用 `walk`

- Each field in a struct
- Each _thing_ in a slice

- 结构中的每个字段
- 切片中的每个 _thing_

Our code at the moment does this but doesn't reflect it very well. We just have a check at the start to see if it's a slice (with a `return` to stop the rest of the code executing) and if it's not we just assume it's a struct.

我们目前的代码做到了这一点，但没有很好地反映它。我们只是在开始时检查它是否是一个切片（使用 `return` 来停止其余代码的执行），如果不是，我们就假设它是一个结构体。

Let's rework the code so instead we check the type _first_ and then do our work.

让我们重新编写代码，而不是检查类型 _first_ 然后做我们的工作。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    switch val.Kind() {
    case reflect.Struct:
        for i:=0;i<val.NumField();i++ {
            walk(val.Field(i).Interface(), fn)
        }
    case reflect.Slice:
        for i:=0;i<val.Len();i++ {
            walk(val.Index(i).Interface(), fn)
        }
    case reflect.String:
        fn(val.String())
    }
}
```

Looking much better! If it's a struct or a slice we iterate over its values calling `walk` on each one. Otherwise, if it's a `reflect.String` we can call `fn`.

看起来好多了！如果它是一个结构体或一个切片，我们遍历它的值，对每个值调用 `walk`。否则，如果它是一个 `reflect.String`，我们可以调用 `fn`。

Still, to me it feels like it could be better. There's repetition of the operation of iterating over fields/values and then calling `walk` but conceptually they're the same.

不过，对我来说，感觉可能会更好。重复遍历字段/值然后调用 `walk` 的操作，但概念上它们是相同的。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    numberOfValues := 0
    var getField func(int) reflect.Value

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
        getField = val.Field
    case reflect.Slice:
        numberOfValues = val.Len()
        getField = val.Index
    }

    for i:=0;i< numberOfValues;i++ {
        walk(getField(i).Interface(), fn)
    }
}
```

If the `value` is a `reflect.String` then we just call `fn` like normal.

如果 `value` 是一个 `reflect.String`，那么我们就像平常一样调用 `fn`。

Otherwise, our `switch` will extract out two things depending on the type

否则，我们的 `switch` 将根据类型提取出两件事

- How many fields there are
- How to extract the `Value` (`Field` or `Index`)

- 有多少个字段
- 如何提取 `Value`（`Field` 或 `Index`）

Once we've determined those things we can iterate through `numberOfValues` calling `walk` with the result of the `getField` function.

一旦我们确定了这些事情，我们就可以通过 `numberOfValues` 调用 `walk` 和 `getField` 函数的结果进行迭代。

Now we've done this, handling arrays should be trivial.

现在我们已经完成了，处理数组应该是微不足道的。

## Write the test first

## 先写测试

Add to the cases

添加到案例

```go
{
    "Arrays",
    [2]Profile {
        {33, "London"},
        {34, "Reykjavík"},
    },
    []string{"London", "Reykjavík"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Arrays
    --- FAIL: TestWalk/Arrays (0.00s)
        reflection_test.go:78: got [], want [London Reykjavík]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Arrays can be handled the same way as slices, so just add it to the case with a comma

数组的处理方式与切片的处理方式相同，因此只需使用逗号将其添加到大小写中

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    numberOfValues := 0
    var getField func(int) reflect.Value

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
        getField = val.Field
    case reflect.Slice, reflect.Array:
        numberOfValues = val.Len()
        getField = val.Index
    }

    for i:=0;i< numberOfValues;i++ {
        walk(getField(i).Interface(), fn)
    }
}
```

The next type we want to handle is `map`.

我们要处理的下一个类型是`map`。

## Write the test first

## 先写测试

```go
{
    "Maps",
    map[string]string{
        "Foo": "Bar",
        "Baz": "Boz",
    },
    []string{"Bar", "Boz"},
},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestWalk/Maps
    --- FAIL: TestWalk/Maps (0.00s)
        reflection_test.go:86: got [], want [Bar Boz]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Again if you think a little abstractly you can see that `map` is very similar to `struct`, it's just the keys are unknown at compile time.

同样，如果您稍微抽象地思考一下，您会发现 `map` 与 `struct` 非常相似，只是在编译时键是未知的。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    numberOfValues := 0
    var getField func(int) reflect.Value

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
        getField = val.Field
    case reflect.Slice, reflect.Array:
        numberOfValues = val.Len()
        getField = val.Index
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walk(val.MapIndex(key).Interface(), fn)
        }
    }

    for i:=0;i< numberOfValues;i++ {
        walk(getField(i).Interface(), fn)
    }
}
```

However, by design you cannot get values out of a map by index. It's only done by _key_, so that breaks our abstraction, darn.

但是，根据设计，您无法通过索引从地图中获取值。它只能由 _key_ 完成，所以这打破了我们的抽象，该死。

## Refactor

## 重构

How do you feel right now? It felt like maybe a nice abstraction at the time but now the code feels a little wonky.

你现在感觉如何？当时感觉这可能是一个不错的抽象，但现在代码感觉有点不稳定。

_This is OK!_ Refactoring is a journey and sometimes we will make mistakes. A major point of TDD is it gives us the freedom to try these things out.

_没关系！_重构是一段旅程，有时我们会犯错误。 TDD 的一个重点是它让我们可以自由地尝试这些东西。

By taking small steps backed by tests this is in no way an irreversible situation. Let's just put it back to how it was before the refactor.

通过采取由测试支持的小步骤，这绝不是不可逆转的情况。让我们把它放回重构之前的样子。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0;i< val.NumField();i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i:= 0;i<val.Len();i++ {
            walkValue(val.Index(i))
        }
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walkValue(val.MapIndex(key))
        }
    }
}
```

We've introduced `walkValue` which DRYs up the calls to `walk` inside our `switch` so that they only have to extract out the `reflect.Value`s from `val`.

我们已经引入了 `walkValue`，它在我们的 `switch` 中干掉了对 `walk` 的调用，这样他们只需从 `val` 中提取出 `reflect.Value`。

### One final problem

### 最后一个问题

Remember that maps in Go do not guarantee order. So your tests will sometimes fail because we assert that the calls to `fn` are done in a particular order.

请记住，Go 中的地图不保证顺序。所以你的测试有时会失败，因为我们断言对 `fn` 的调用是按特定顺序完成的。

To fix this, we'll need to move our assertion with the maps to a new test where we do not care about the order.

为了解决这个问题，我们需要将带有映射的断言移动到我们不关心顺序的新测试中。

```go
t.Run("with maps", func(t *testing.T) {
    aMap := map[string]string{
        "Foo": "Bar",
        "Baz": "Boz",
    }

    var got []string
    walk(aMap, func(input string) {
        got = append(got, input)
    })

    assertContains(t, got, "Bar")
    assertContains(t, got, "Boz")
})
```

Here is how `assertContains` is defined

这里是如何定义 `assertContains`

```go
func assertContains(t testing.TB, haystack []string, needle string)  {
    t.Helper()
    contains := false
    for _, x := range haystack {
        if x == needle {
            contains = true
        }
    }
    if !contains {
        t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
    }
}
```

The next type we want to handle is `chan`.

我们要处理的下一个类型是 `chan`。

## Write the test first

## 先写测试

```go
t.Run("with channels", func(t *testing.T) {
        aChannel := make(chan Profile)

        go func() {
            aChannel <- Profile{33, "Berlin"}
            aChannel <- Profile{34, "Katowice"}
            close(aChannel)
        }()

        var got []string
        want := []string{"Berlin", "Katowice"}

        walk(aChannel, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v, want %v", got, want)
        }
    })
```

## Try to run the test

## 尝试运行测试

```
--- FAIL: TestWalk (0.00s)
    --- FAIL: TestWalk/with_channels (0.00s)
        reflection_test.go:115: got [], want [Berlin Katowice]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We can iterate through all values sent through channel until it was closed with Recv()

我们可以遍历所有通过 channel 发送的值，直到它被 Recv() 关闭

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0;i < val.NumField();i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0;i < val.Len();i++ {
            walkValue(val.Index(i))
        }
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walkValue(val.MapIndex(key))
        }
    case reflect.Chan:
        for v, ok := val.Recv();ok;v, ok = val.Recv() {
            walkValue(v)
        }
    }
}
```

The next type we want to handle is `func`.

我们要处理的下一个类型是 `func`。

## Write the test first

## 先写测试

```go
t.Run("with function", func(t *testing.T) {
        aFunction := func() (Profile, Profile) {
            return Profile{33, "Berlin"}, Profile{34, "Katowice"}
        }

        var got []string
        want := []string{"Berlin", "Katowice"}

        walk(aFunction, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v, want %v", got, want)
        }
    })
```

## Try to run the test

## 尝试运行测试

```
--- FAIL: TestWalk (0.00s)
    --- FAIL: TestWalk/with_function (0.00s)
        reflection_test.go:132: got [], want [Berlin Katowice]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Non zero-argument functions do not seem to make a lot of sense in this scenario. But we should allow for arbitrary return values.

在这种情况下，非零参数函数似乎没有多大意义。但是我们应该允许任意返回值。

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0;i < val.NumField();i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0;i < val.Len();i++ {
            walkValue(val.Index(i))
        }
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walkValue(val.MapIndex(key))
        }
    case reflect.Chan:
        for v, ok := val.Recv();ok;v, ok = val.Recv() {
            walkValue(v)
        }
    case reflect.Func:
        valFnResult := val.Call(nil)
        for _, res := range valFnResult {
            walkValue(res)
        }
    }
}
```

## Wrapping up

##  总结

- Introduced some of the concepts from the `reflect` package.
- Used recursion to traverse arbitrary data structures.
- Did an in retrospect bad refactor but didn't get too upset about it. By working iteratively with tests it's not such a big deal.
- This only covered a small aspect of reflection. [The Go blog has an excellent post covering more details](https://blog.golang.org/laws-of-reflection).
- Now that you know about reflection, do your best to avoid using it. 

- 引入了 `reflect` 包中的一些概念。
- 使用递归遍历任意数据结构。
- 回想起来，做了一个糟糕的重构，但并没有为此感到太沮丧。通过反复进行测试，这不是什么大问题。
- 这仅涵盖了反射的一小部分。 [Go 博客有一篇很棒的文章，涵盖了更多细节](https://blog.golang.org/laws-of-reflection)。
- 既然您了解反射，请尽量避免使用它。

