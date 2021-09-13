# Arrays and slices

# 数组和切片

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/arrays)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/arrays)**

Arrays allow you to store multiple elements of the same type in a variable in a particular order.

数组允许您以特定顺序将多个相同类型的元素存储在变量中。

When you have an array, it is very common to have to iterate over them. So let's use [our new-found knowledge of `for`](https://github.com/quii/learn-go-with-tests/blob/main/iteration.md) to make a `Sum` function. `Sum` will take an array of numbers and return the total.

当你有一个数组时，必须遍历它们是很常见的。因此，让我们使用 [我们对 `for` 的新知识](https://github.com/quii/learn-go-with-tests/blob/main/iteration.md) 来创建一个 `Sum` 函数。 `Sum` 将采用一组数字并返回总数。

Let's use our TDD skills

让我们使用我们的 TDD 技能

## Write the test first

## 先写测试

Create a new folder to work in. Create a new file called `sum_test.go` and insert the following:

创建一个新文件夹以进行工作。创建一个名为 `sum_test.go` 的新文件并插入以下内容：

```
package main

import "testing"

func TestSum(t *testing.T) {

    numbers := [5]int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if got != want {
        t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
}
```

Arrays have a *fixed capacity* which you define when you declare the variable. We can initialize an array in two ways:

数组具有*固定容量*，您在声明变量时定义该容量。我们可以通过两种方式初始化数组：

- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

- [N]type{value1, value2, ..., valueN} 例如`数字：= [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} 例如`数字：= [...] int{1, 2, 3, 4, 5}`

It is sometimes useful to also print the inputs to the function in the error message. Here, we are using the `%v` placeholder to print the "default" format, which works well for arrays.

在错误消息中打印函数的输入有时也很有用。在这里，我们使用 `%v` 占位符来打印“默认”格式，它适用于数组。

[Read more about the format strings](https://golang.org/pkg/fmt/)

[阅读有关格式字符串的更多信息](https://golang.org/pkg/fmt/)

## Try to run the test

## 尝试运行测试

By running `go test` the compiler will fail with `./sum_test.go:10:15: undefined: Sum`

通过运行 `go test`，编译器将失败并返回 `./sum_test.go:10:15: undefined: Sum`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

In `sum.go`

在`sum.go`

```
package main

func Sum(numbers [5]int) int {
    return 0
}
```

Your test should now fail with *a clear error message*

您的测试现在应该会失败并显示 *清晰的错误消息*

```
sum_test.go:13: got 0 want 15 given, [1 2 3 4 5]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```
func Sum(numbers [5]int) int {
    sum := 0
    for i := 0;i < 5;i++ {
        sum += numbers[i]
    }
    return sum
}
```

To get the value out of an array at a particular index, just use `array[index]` syntax. In this case, we are using `for` to iterate 5 times to work through the array and add each item onto `sum`.

要从特定索引处的数组中获取值，只需使用 `array[index]` 语法。在这种情况下，我们使用 `for` 迭代 5 次以处理数组并将每个项目添加到 `sum`。

## Refactor

## 重构

Let's introduce [`range`](https://gobyexample.com/range) to help clean up our code

让我们引入 [`range`](https://gobyexample.com/range) 来帮助清理我们的代码

```
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```

`range` lets you iterate over an array. On each iteration, `range` returns two values - the index and the value. We are choosing to ignore the index value by using `_` [blank identifier](https://golang.org/doc/effective_go.html#blank).

`range` 可以让你遍历一个数组。在每次迭代中，`range` 返回两个值 - 索引和值。我们选择使用 `_` [空白标识符](https://golang.org/doc/effective_go.html#blank) 来忽略索引值。

#
### Arrays and their type

#
### 数组及其类型

An interesting property of arrays is that the size is encoded in its type. If you try to pass an `[4]int` into a function that expects `[5]int`, it won't compile. They are different types so it's just the same as trying to pass a `string` into a function that wants an `int`.

数组的一个有趣的特性是大小以它的类型编码。如果您尝试将 `[4]int` 传递给需要 `[5]int` 的函数，它将无法编译。它们是不同的类型，所以这与尝试将一个 `string` 传递给一个需要 `int` 的函数是一样的。

You may be thinking it's quite cumbersome that arrays have a fixed length, and most of the time you probably won't be using them!

您可能会认为数组具有固定长度很麻烦，而且大多数时候您可能不会使用它们！

Go has *slices* which do not encode the size of the collection and instead can have any size.

Go 有 *slices*，它不编码集合的大小，而是可以有任何大小。

The next requirement will be to sum collections of varying sizes.

下一个要求是对不同大小的集合求和。

## Write the test first

## 先写测试

We will now use the [slice type](https://golang.org/doc/effective_go.html#slices) which allows us to have collections of any size. The syntax is very similar to arrays, you just omit the size when declaring them

我们现在将使用 [slice type](https://golang.org/doc/effective_go.html#slices)，它允许我们拥有任何大小的集合。语法与数组非常相似，您只需在声明它们时省略大小

```
mySlice := []int{1,2,3}` rather than `myArray := [3]int{1,2,3}
func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := [5]int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

}
```

## Try and run the test

## 尝试并运行测试

This does not compile

这不编译

```
./sum_test.go:22:13: cannot use numbers (type []int) as type [5]int in argument to Sum
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

The problem here is we can either

这里的问题是我们可以

- Break the existing API by changing the argument to `Sum` to be a slice rather than an array. When we do this, we will potentially ruin someone's day because our *other* test will no longer compile!
- Create a new function 

- 通过将 `Sum` 的参数更改为切片而不是数组来打破现有的 API。当我们这样做时，我们可能会毁了某人的一天，因为我们的 *other* 测试将不再编译！
- 创建一个新函数

In our case, no one else is using our function, so rather than having two functions to maintain, let's have just one.

在我们的例子中，没有其他人在使用我们的函数，所以与其维护两个函数，不如只维护一个。

```
func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```

If you try to run the tests they will still not compile, you will  have to change the first test to pass in a slice rather than an array.

如果您尝试运行它们仍然无法编译的测试，则必须将第一个测试更改为传入切片而不是数组。

## Write enough code to make it pass

## 编写足够的代码使其通过

It turns out that fixing the compiler problems were all we need to do here and the tests pass!

事实证明，我们在这里只需要修复编译器问题，并且测试通过了！

## Refactor

## 重构

We already refactored `Sum` - all we did was replace arrays with slices, so no extra changes are required. Remember that we must not neglect our test code in the refactoring stage - we can further improve our `Sum` tests.

我们已经重构了 `Sum` - 我们所做的只是用切片替换数组，因此不需要额外的更改。请记住，我们不能在重构阶段忽略我们的测试代码——我们可以进一步改进我们的“Sum”测试。

```
func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

}
```

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much *confidence* as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. **Every test has a cost**.

质疑测试的价值很重要。目标不应该是尽可能多的测试，而应该是对您的代码库有尽可能多的*信心*。太多的测试会变成一个真正的问题，它只会增加更多的维护开销。 **每个测试都有成本**。

In our case, you can see that having two tests for this function is redundant. If it works for a slice of one size it's very likely it'll work for a slice of any size (within reason).

在我们的例子中，您可以看到对这个函数进行两次测试是多余的。如果它适用于一个大小的切片，它很可能适用于任何大小的切片（在合理范围内）。

Go's built-in testing toolkit features a [coverage tool](https://blog.golang.org/cover). Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.

Go 的内置测试工具包具有一个 [覆盖工具](https://blog.golang.org/cover)。虽然争取 100% 的覆盖率不应该是您的最终目标，但覆盖率工具可以帮助识别测试未覆盖的代码区域。如果您对 TDD 一直很严格，那么无论如何您很可能都会接近 100% 的覆盖率。

Try running

尝试跑步

```
go test -cover
```

You should see

你应该看到

```
PASS
coverage: 100.0% of statements
```

Now delete one of the tests and check the coverage again.

现在删除其中一个测试并再次检查覆盖率。

Now that we are happy we have a well-tested function you should commit your great work before taking on the next challenge.

现在我们很高兴我们有一个经过良好测试的功能，您应该在接受下一个挑战之前做出出色的工作。

We need a new function called `SumAll` which will take a varying number of slices, returning a new slice containing the totals for each slice passed in.

我们需要一个名为“SumAll”的新函数，它将采用不同数量的切片，返回一个包含传入的每个切片的总数的新切片。

For example

例如

```
SumAll([]int{1,2}, []int{0,9})` would return `[]int{3, 9}
```

or

或者

```
SumAll([]int{1,1,1})` would return `[]int{3}
```

## Write the test first

## 先写测试

```
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if got != want {
        t.Errorf("got %v want %v", got, want)
    }
}
```

## Try and run the test

## 尝试并运行测试

```
./sum_test.go:23:9: undefined: SumAll
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We need to define `SumAll` according to what our test wants.

我们需要根据我们的测试需要定义 `SumAll`。

Go can let you write [*variadic functions*](https://gobyexample.com/variadic-functions) that can take a variable number of arguments.

Go 可以让你编写 [*variadic 函数*](https://gobyexample.com/variadic-functions)，它可以接受可变数量的参数。

```
func SumAll(numbersToSum ...[]int) (sums []int) {
    return
}
```

This is valid, but our tests still won't compile!

这是有效的，但我们的测试仍然无法编译！

```
./sum_test.go:26:9: invalid operation: got != want (slice can only be compared to nil)
```

Go does not let you use equality operators with slices. You *could* write a function to iterate over each `got` and `want` slice and check their values but for convenience sake, we can use [`reflect.DeepEqual`](https://golang.org/pkg/reflect/#DeepEqual) which is useful for seeing if *any* two variables are the same.

Go 不允许你在切片中使用相等运算符。您*可以*编写一个函数来迭代每个 `got` 和 `want` 切片并检查它们的值，但为了方便起见，我们可以使用 [`reflect.DeepEqual`](https://golang.org/pkg/reflect/#DeepEqual) 这对于查看 *any* 两个变量是否相同很有用。

```
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

(make sure you `import reflect` in the top of your file to have access to `DeepEqual`)

（确保您在文件顶部“导入反射”以访问“DeepEqual”）

It's important to note that `reflect.DeepEqual` is not "type safe" - the code will compile even if you did something a bit silly. To see this in action, temporarily change the test to:

重要的是要注意 `reflect.DeepEqual` 不是“类型安全的”——即使你做了一些愚蠢的事情，代码也会编译。要查看此操作，请暂时将测试更改为：

```
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := "bob"

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

What we have done here is try to compare a `slice` with a `string`. This makes no sense, but the test compiles! So while using `reflect.DeepEqual` is a convenient way of comparing slices (and other things) you must be careful when using it.

我们在这里所做的是尝试将 `slice` 与 `string` 进行比较。这没有任何意义，但测试编译通过！因此，虽然使用 `reflect.DeepEqual` 是一种比较切片（和其他东西）的便捷方法，但在使用它时必须小心。

Change the test back again and run it. You should have test output like the following

再次更改测试并运行它。你应该有如下的测试输出

```
sum_test.go:30: got [] want [3 9]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

What we need to do is iterate over the varargs, calculate the sum using our existing `Sum` function, then add it to the slice we will return

我们需要做的是迭代可变参数，使用我们现有的 `Sum` 函数计算总和，然后将其添加到我们将返回的切片中

```
func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numbersToSum {
        sums[i] = Sum(numbers)
    }

    return sums
}
```

Lots of new things to learn!

很多新东西要学习！

There's a new way to create a slice. `make` allows you to create a slice with a starting capacity of the `len` of the `numbersToSum` we need to work through.

有一种创建切片的新方法。 `make` 允许您创建一个切片，其起始容量为我们需要处理的 `numbersToSum` 的 `len`。

You can index slices like arrays with `mySlice[N]` to get the value out or assign it a new value with `=`

您可以使用 `mySlice[N]` 索引切片之类的数组以获取值或使用 `=` 为其分配一个新值

The tests should now pass.

测试现在应该通过了。

## Refactor

## 重构

As mentioned, slices have a capacity. If you have a slice with a capacity of 2 and try to do `mySlice[10] = 1` you will get a *runtime* error.

如前所述，切片具有容量。如果您有一个容量为 2 的切片并尝试执行 `mySlice[10] = 1`，您将收到 *runtime* 错误。

However, you can use the `append` function which takes a slice and a new value, then returns a new slice with all the items in it.

但是，您可以使用 `append` 函数，它接受一个切片和一个新值，然后返回一个包含所有项目的新切片。

```
func SumAll(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}
```

In this implementation, we are worrying less about capacity. We start with an empty slice `sums` and append to it the result of `Sum` as we work through the varargs.

在这个实现中，我们对容量的担心较少。我们从一个空切片 `sums` 开始，并在处理可变参数时将 `Sum` 的结果附加到它上面。

Our next requirement is to change `SumAll` to `SumAllTails`, where it will calculate the totals of the "tails" of each slice. The tail of a collection is all items in the collection except the first one (the "head").

我们的下一个要求是将 `SumAll` 更改为 `SumAllTails`，它将计算每个切片的“尾部”的总数。集合的尾部是集合中除第一个（“头”）之外的所有项目。

## Write the test first

## 先写测试

```
func TestSumAllTails(t *testing.T) {
    got := SumAllTails([]int{1, 2}, []int{0, 9})
    want := []int{2, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

## Try and run the test

## 尝试并运行测试

```
./sum_test.go:26:9: undefined: SumAllTails
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Rename the function to `SumAllTails` and re-run the test

将函数重命名为 `SumAllTails` 并重新运行测试

```
sum_test.go:30: got [3 9] want [2 9]
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```
func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        tail := numbers[1:]
        sums = append(sums, Sum(tail))
    }

    return sums
}
```

Slices can be sliced! The syntax is `slice[low:high]`. If you omit the value on one of the sides of the `:` it captures everything to that side of it. In our case, we are saying "take from 1 to the end" with `numbers[1:]`. You may wish to spend some time writing other tests around slices and experiment with the slice operator to get more familiar with it.

切片可以切片！语法是`slice[low:high]`。如果省略 `:` 一侧的值，它会捕获该一侧的所有内容。在我们的例子中，我们用 `numbers[1:]` 说“从 1 到最后”。您可能希望花一些时间围绕切片编写其他测试并尝试使用切片运算符以更熟悉它。

## Refactor

## 重构

Not a lot to refactor this time.

这次重构的内容不多。

What do you think would happen if you passed in an empty slice into our function? What is the "tail" of an empty slice? What happens when you tell Go to capture all elements from `myEmptySlice[1:]`?

如果您将一个空切片传入我们的函数，您认为会发生什么？空切片的“尾巴”是什么？当你告诉 Go 从 `myEmptySlice[1:]` 中捕获所有元素时会发生什么？

## Write the test first

## 先写测试

```
func TestSumAllTails(t *testing.T) {

    t.Run("make the sums of some slices", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

}
```

## Try and run the test

## 尝试并运行测试

```
panic: runtime error: slice bounds out of range [recovered]
    panic: runtime error: slice bounds out of range
```

Oh no! It's important to note the test *has compiled*, it is a runtime error. Compile time errors are our friend because they help us write software that works, runtime errors are our enemies because they affect our users.

不好了！重要的是要注意测试*已编译*，这是一个运行时错误。编译时错误是我们的朋友，因为它们帮助我们编写有效的软件，运行时错误是我们的敌人，因为它们会影响我们的用户。

## Write enough code to make it pass

## 编写足够的代码使其通过

```
func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        if len(numbers) == 0 {
            sums = append(sums, 0)
        } else {
            tail := numbers[1:]
            sums = append(sums, Sum(tail))
        }
    }

    return sums
}
```

## Refactor

## 重构

Our tests have some repeated code around the assertions again, so let's extract those into a function

我们的测试在断言周围有一些重复的代码，所以让我们将它们提取到一个函数中

```
func TestSumAllTails(t *testing.T) {

    checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of tails of", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checkSums(t, got, want)
    })

}
```

A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly adds a new test with `checkSums(t, got, "dave")` the compiler will stop them in their tracks.

这样做的一个方便的副作用是这为我们的代码增加了一点类型安全性。如果开发人员错误地添加了一个带有 `checkSums(t, got, "dave")` 的新测试，编译器将阻止他们继续前进。

```
$ go test
./sum_test.go:52:21: cannot use "dave" (type string) as type []int in argument to checkSums
```

## Wrapping up

##  总结

We have covered

我们已经涵盖

- Arrays
- Slices
   - The various ways to make them
   - How they have a *fixed* capacity but you can create new slices from old ones using `append`
   - How to slice, slices!
- `len` to get the length of an array or slice
- Test coverage tool
- `reflect.DeepEqual` and why it's useful but can reduce the type-safety of your code

- 数组
- 切片
  - 制作它们的各种方法
  - 它们如何具有*固定*容量，但您可以使用`append`从旧切片创建新切片
  - 如何切片，切片！
- `len` 获取数组或切片的长度
- 测试覆盖率工具
- `reflect.DeepEqual` 以及为什么它有用但会降低代码的类型安全性

We've used slices and arrays with integers but they work with any other type too, including arrays/slices themselves. So you can declare a variable of `[][]string` if you need to.

我们使用了带有整数的切片和数组，但它们也适用于任何其他类型，包括数组/切片本身。因此，如果需要，您可以声明一个 `[][]string` 变量。

[Check out the Go blog post on slices](https://blog.golang.org/go-slices-usage-and-internals) for an in-depth look into slices. Try writing more tests to solidify what you learn from reading it.

[查看关于切片的 Go 博客文章](https://blog.golang.org/go-slices-usage-and-internals) 深入了解切片。尝试编写更多测试以巩固您从阅读中学到的知识。

Another handy way to experiment with Go other than writing tests is the Go playground. You can try most things out and you can easily share your code if you need to ask questions. [I have made a go playground with a slice in it for you to experiment with.](https://play.golang.org/p/ICCWcRGIO68)

除了编写测试之外，另一种方便的 Go 实验方法是 Go 游乐场。您可以尝试大多数事情，如果您需要提问，您可以轻松共享您的代码。 [我制作了一个带有切片的 go 游乐场供您试验。](https://play.golang.org/p/ICCWcRGIO68)

[Here is an example](https://play.golang.org/p/bTrRmYfNYCp) of slicing an array and how changing the slice affects the original array; but a "copy" of the slice will not affect the original array. [Another example](https://play.golang.org/p/Poth8JS28sc) of why it's a good idea to make a copy of a slice after slicing a very large slice. 

[这里是一个例子](https://play.golang.org/p/bTrRmYfNYCp) 切片数组以及改变切片如何影响原始数组；但是切片的“副本”不会影响原始数组。 [另一个例子](https://play.golang.org/p/Poth8JS28sc) 关于为什么在切割一个非常大的切片后制作一个切片的副本是个好主意。

