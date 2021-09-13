# Pointers & errors

# 指针和错误

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/pointers)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/pointers)**

We learned about structs in the last section which let us capture a number of values related around a concept.

我们在上一节中学习了结构体，它让我们捕获了一些与概念相关的值。

At some point you may wish to use structs to manage state, exposing methods to let users change the state in a way that you can control.

在某些时候，您可能希望使用结构来管理状态，公开方法让用户以您可以控制的方式更改状态。

**Fintech loves Go** and uhhh bitcoins? So let's show what an amazing banking system we can make.

**金融科技喜欢 Go **和比特币？因此，让我们展示我们可以制作出多么了不起的银行系统。

Let's make a `Wallet` struct which lets us deposit `Bitcoin`.

让我们创建一个“钱包”结构，让我们存入“比特币”。

## Write the test first

## 先写测试

```go
func TestWallet(t *testing.T) {

    wallet := Wallet{}

    wallet.Deposit(10)

    got := wallet.Balance()
    want := 10

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

In the [previous example](./structs-methods-and-interfaces.md) we accessed fields directly with the field name, however in our _very secure wallet_ we don't want to expose our inner state to the rest of the world . We want to control access via methods.

在[上一个示例](./structs-methods-and-interfaces.md) 中，我们直接使用字段名称访问字段，但是在我们_非常安全的钱包_中，我们不想将我们的内部状态暴露给世界其他地方.我们想通过方法控制访问。

## Try to run the test

## 尝试运行测试

`./wallet_test.go:7:12: undefined: Wallet`

`./wallet_test.go:7:12: 未定义：钱包`

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

The compiler doesn't know what a `Wallet` is so let's tell it.

编译器不知道什么是“钱包”，所以让我们告诉它。

```go
type Wallet struct { }
```

Now we've made our wallet, try and run the test again

现在我们已经制作了我们的钱包，再次尝试运行测试

```go
./wallet_test.go:9:8: wallet.Deposit undefined (type Wallet has no field or method Deposit)
./wallet_test.go:11:15: wallet.Balance undefined (type Wallet has no field or method Balance)
```

We need to define these methods.

我们需要定义这些方法。

Remember to only do enough to make the tests run. We need to make sure our test fails correctly with a clear error message.

记住只做足以使测试运行的工作。我们需要确保我们的测试正确失败并显示明确的错误消息。

```go
func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
    return 0
}
```

If this syntax is unfamiliar go back and read the structs section.

如果这个语法不熟悉，请返回并阅读结构部分。

The tests should now compile and run

测试现在应该编译并运行

`wallet_test.go:15: got 0 want 10`

## Write enough code to make it pass

## 编写足够的代码使其通过

We will need some kind of _balance_ variable in our struct to store the state

我们将需要在我们的结构中使用某种 _balance_ 变量来存储状态

```go
type Wallet struct {
    balance int
}
```

In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private _outside the package it's defined in_.

在 Go 中，如果一个符号（变量、类型、函数等）以小写符号开头，那么它是私有的_在它定义的包之外_。

In our case we want our methods to be able to manipulate this value, but no one else.

在我们的例子中，我们希望我们的方法能够操作这个值，但没有其他方法。

Remember we can access the internal `balance` field in the struct using the "receiver" variable.

请记住，我们可以使用“receiver”变量访问结构中的内部 `balance` 字段。

```go
func (w Wallet) Deposit(amount int) {
    w.balance += amount
}

func (w Wallet) Balance() int {
    return w.balance
}
```

With our career in fintech secured, run the test suite and bask in the passing test

确保我们在金融科技领域的职业生涯得到保障，运行测试套件并享受通过的测试

`wallet_test.go:15: got 0 want 10`



Well this is confusing, our code looks like it should work.
We add the new amount onto our balance and then the balance method should return the current state of it.

好吧，这令人困惑，我们的代码看起来应该可以工作。
我们将新金额添加到我们的余额中，然后 balance 方法应该返回它的当前状态。

In Go, **when you call a function or a method the arguments are** _**copied**_.

在 Go 中，**当你调用一个函数或方法时，参数是** _**复制的**_。

When calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.

当调用 `func (w Wallet) Deposit(amount int)` 时，`w` 是我们调用该方法的任何内容的副本。

Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the _address_ of that bit of memory with `&myVal`.

无需过于计算机科学，当您创建一个价值时——比如钱包，它会存储在内存中的某个地方。你可以用 `&myVal` 找出那块内存的 _address_ 是什么。

Experiment by adding some prints to your code

通过在代码中添加一些打印进行实验

```go
func TestWallet(t *testing.T) {

    wallet := Wallet{}

    wallet.Deposit(10)

    got := wallet.Balance()

    fmt.Printf("address of balance in test is %v \n", &wallet.balance)

    want := 10

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

```go
func (w Wallet) Deposit(amount int) {
    fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
    w.balance += amount
}
```

The `\n` escape character prints a new line after outputting the memory address.
We get the pointer (memory address) of something by placing an `&` character at the beginning of the symbol.

`\n` 转义字符在输出内存地址后打印一个新行。
我们通过在符号的开头放置一个 `&` 字符来获取某事物的指针（内存地址）。

Now re-run the test

现在重新运行测试

```text
address of balance in Deposit is 0xc420012268
address of balance in test is 0xc420012260
```

You can see that the addresses of the two balances are different. So when we change the value of the balance inside the code, we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.

可以看到两个余额的地址是不同的。因此，当我们更改代码中余额的值时，我们正在处理来自测试的内容的副本。因此，测试中的余额不变。

We can fix this with _pointers_. [Pointers](https://gobyexample.com/pointers) let us _point_ to some values and then let us change them. 

我们可以用 _pointers_ 解决这个问题。 [指针](https://gobyexample.com/pointers) 让我们_指向_一些值，然后让我们更改它们。

So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.

因此，我们不是获取整个钱包的副本，而是获取指向该钱包的指针，以便我们可以更改其中的原始值。

```go
func (w *Wallet) Deposit(amount int) {
    w.balance += amount
}

func (w *Wallet) Balance() int {
    return w.balance
}
```

The difference is the receiver type is `*Wallet` rather than `Wallet` which you can read as "a pointer to a wallet".

不同之处在于接收者类型是 `*Wallet` 而不是 `Wallet`，你可以将其读作“指向钱包的指针”。

Try and re-run the tests and they should pass.

尝试并重新运行测试，它们应该会通过。

Now you might wonder, why did they pass? We didn't dereference the pointer in the function, like so:

现在你可能想知道，为什么他们通过了？我们没有取消引用函数中的指针，如下所示：

```go
func (w *Wallet) Balance() int {
    return (*w).balance
}
```

and seemingly addressed the object directly. In fact, the code above using `(*w)` is absolutely valid. However, the makers of Go deemed this notation cumbersome, so the language permits us to write `w.balance`, without an explicit dereference.
These pointers to structs even have their own name: _struct pointers_ and they are [automatically dereferenced](https://golang.org/ref/spec#Method_values).

并且似乎直接针对对象。事实上，上面使用`(*w)` 的代码是绝对有效的。然而，Go 的开发者认为这种表示法很麻烦，因此该语言允许我们编写 `w.balance`，而无需显式取消引用。
这些指向结构体的指针甚至有自己的名字：_struct pointers_ 并且它们[自动取消引用](https://golang.org/ref/spec#Method_values)。

Technically you do not need to change `Balance` to use a pointer receiver as taking a copy of the balance is fine. However, by convention you should keep your method receiver types the same for consistency.

从技术上讲，您不需要更改 `Balance` 来使用指针接收器，因为复制余额就可以了。但是，按照惯例，您应该保持方法接收器类型相同以保持一致性。

## Refactor

## 重构

We said we were making a Bitcoin wallet but we have not mentioned them so far. We've been using `int` because they're a good type for counting things!

我们说我们正在制作一个比特币钱包，但到目前为止我们还没有提到它们。我们一直在使用 `int` 因为它们是一种很好的计数类型！

It seems a bit overkill to create a `struct` for this. `int` is fine in terms of the way it works but it's not descriptive.

为此创建一个 `struct` 似乎有点矫枉过正。 `int` 的工作方式很好，但它不是描述性的。

Go lets you create new types from existing ones.

Go 允许您从现有类型创建新类型。

The syntax is `type MyName OriginalType`

语法是`type MyName OriginalType`

```go
type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
```

```go
func TestWallet(t *testing.T) {

    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))

    got := wallet.Balance()

    want := Bitcoin(10)

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

To make `Bitcoin` you just use the syntax `Bitcoin(999)`.

要制作`Bitcoin`，您只需使用语法`Bitcoin(999)`。

By doing this we're making a new type and we can declare _methods_ on them. This can be very useful when you want to add some domain specific functionality on top of existing types.

通过这样做，我们正在创建一个新类型，我们可以在它们上声明 _methods_。当您想在现有类型之上添加一些特定于域的功能时，这非常有用。

Let's implement [Stringer](https://golang.org/pkg/fmt/#Stringer) on Bitcoin

让我们在比特币上实现 [Stringer](https://golang.org/pkg/fmt/#Stringer)

```go
type Stringer interface {
        String() string
}
```

This interface is defined in the `fmt` package and lets you define how your type is printed when used with the `%s` format string in prints.

该接口在 `fmt` 包中定义，并允许您定义在打印中与 `%s` 格式字符串一起使用时如何打印类型。

```go
func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
```

As you can see, the syntax for creating a method on a type declaration is the same as it is on a struct.

如您所见，在类型声明上创建方法的语法与在结构上创建的语法相同。

Next we need to update our test format strings so they will use `String()` instead.

接下来我们需要更新我们的测试格式字符串，以便它们使用 `String()` 代替。

```go
     if got != want {
        t.Errorf("got %s want %s", got, want)
    }
```

To see this in action, deliberately break the test so we can see it

要看到它的实际效果，请故意打破测试以便我们可以看到它

`wallet_test.go:18: got 10 BTC want 20 BTC`

`wallet_test.go:18：得到 10 BTC 想要 20 BTC`

This makes it clearer what's going on in our test.

这使我们的测试中发生的事情更加清晰。

The next requirement is for a `Withdraw` function.

下一个要求是“提款”功能。

## Write the test first

## 先写测试

Pretty much the opposite of `Deposit()`

几乎与`Deposit()`相反

```go
func TestWallet(t *testing.T) {

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}

        wallet.Deposit(Bitcoin(10))

        got := wallet.Balance()

        want := Bitcoin(10)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}

        wallet.Withdraw(Bitcoin(10))

        got := wallet.Balance()

        want := Bitcoin(10)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })
}
```

## Try to run the test

## 尝试运行测试

`./wallet_test.go:26:9: wallet.Withdraw undefined (type Wallet has no field or method Withdraw)`



## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

```go
func (w *Wallet) Withdraw(amount Bitcoin) {

}
```

`wallet_test.go:33: got 20 BTC want 10 BTC`



## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (w *Wallet) Withdraw(amount Bitcoin) {
    w.balance -= amount
}
```

## Refactor

## 重构

There's some duplication in our tests, lets refactor that out.

我们的测试中有一些重复，让我们重构一下。

```go
func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

}
```

What should happen if you try to `Withdraw` more than is left in the account? For now, our requirement is to assume there is not an overdraft facility.

如果您尝试“提款”超过帐户中剩余的金额，会发生什么？目前，我们的要求是假设没有透支设施。

How do we signal a problem when using `Withdraw`?

使用“提款”时我们如何发出问题信号？

In Go, if you want to indicate an error it is idiomatic for your function to return an `err` for the caller to check and act on.

在 Go 中，如果你想指出一个错误，你的函数通常会返回一个 `err` 供调用者检查并采取行动。

Let's try this out in a test.

让我们在测试中试试这个。

## Write the test first

## 先写测试

```go
t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertBalance(t, wallet, startingBalance)

    if err == nil {
        t.Error("wanted an error but didn't get one")
    }
})
```

We want `Withdraw` to return an error _if_ you try to take out more than you have and the balance should stay the same.

我们希望 `Withdraw` 返回错误 _if_ 您尝试取出的数量超过您拥有的数量，并且余额应该保持不变。

We then check an error has returned by failing the test if it is `nil`.

然后我们通过测试失败来检查返回的错误是否为“nil”。

`nil` is synonymous with `null` from other programming languages. Errors can be `nil` because the return type of `Withdraw` will be `error`, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

`nil` 与其他编程语言中的 `null` 同义。错误可以是 `nil`，因为 `Withdraw` 的返回类型将是 `error`，它是一个接口。如果您看到一个接受参数或返回接口值的函数，则它们可以为 nillable。

Like `null` if you try to access a value that is `nil` it will throw a **runtime panic**. This is bad! You should make sure that you check for nils.

就像 `null` 一样，如果你尝试访问一个 `nil` 的值，它会抛出一个 **runtime panic**。这不好！您应该确保检查 nils。

## Try and run the test

## 尝试并运行测试

`./wallet_test.go:31:25: wallet.Withdraw(Bitcoin(100)) used as value`

The wording is perhaps a little unclear, but our previous intent with `Withdraw` was just to call it, it will never return a value. To make this compile we will need to change it so it has a return type.

措辞可能有点不清楚，但我们之前使用 `Withdraw` 的意图只是调用它，它永远不会返回值。为了使这个编译，我们需要改变它，使它有一个返回类型。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
    w.balance -= amount
    return nil
}
```

Again, it is very important to just write enough code to satisfy the compiler. We correct our `Withdraw` method to return `error` and for now we have to return _something_ so let's just return `nil`.

同样，编写足够的代码以满足编译器的要求非常重要。我们更正我们的 `Withdraw` 方法以返回 `error`，现在我们必须返回 _something_ 所以让我们只返回 `nil`。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return errors.New("oh no")
    }

    w.balance -= amount
    return nil
}
```

Remember to import `errors` into your code.

请记住将 `errors` 导入到您的代码中。

`errors.New` creates a new `error` with a message of your choosing.

`errors.New` 使用您选择的消息创建一个新的 `error`。

## Refactor

## 重构

Let's make a quick test helper for our error check to improve the test's readability

让我们为我们的错误检查制作一个快速测试助手，以提高测试的可读性

```go
assertError := func(t testing.TB, err error) {
    t.Helper()
    if err == nil {
        t.Error("wanted an error but didn't get one")
    }
}
```

And in our test

在我们的测试中

```go
t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertError(t, err)
    assertBalance(t, wallet, startingBalance)
})
```

Hopefully when returning an error of "oh no" you were thinking that we _might_ iterate on that because it doesn't seem that useful to return.

希望在返回“哦不”的错误时，您认为我们可能会对其进行迭代，因为返回似乎没有那么有用。

Assuming that the error ultimately gets returned to the user, let's update our test to assert on some kind of error message rather than just the existence of an error.

假设错误最终返回给用户，让我们更新我们的测试以断言某种错误消息，而不仅仅是错误的存在。

## Write the test first

## 先写测试

Update our helper for a `string` to compare against.

更新我们的 helper 来比较一个 `string`。

```go
assertError := func(t testing.TB, got error, want string) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got.Error() != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

And then update the caller

然后更新调用者

```go
t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertError(t, err, "cannot withdraw, insufficient funds")
    assertBalance(t, wallet, startingBalance)
})
```

We've introduced `t.Fatal` which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around. Without this the test would carry on to the next step and panic because of a nil pointer.

我们引入了`t.Fatal`，如果它被调用，它将停止测试。这是因为如果周围没有错误，我们不想对返回的错误做出更多断言。如果没有这个，测试将继续进行下一步并由于 nil 指针而恐慌。

## Try to run the test

## 尝试运行测试

`wallet_test.go:61: got err 'oh no' want 'cannot withdraw, insufficient funds'`

`wallet_test.go:61：得到错误'哦不'想要'不能提款，资金不足'`

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return errors.New("cannot withdraw, insufficient funds")
    }

    w.balance -= amount
    return nil
}
```

## Refactor

## 重构

We have duplication of the error message in both the test code and the `Withdraw` code. 

我们在测试代码和 `Withdraw` 代码中都有重复的错误消息。

It would be really annoying for the test to fail if someone wanted to re-word the error and it's just too much detail for our test. We don't _really_ care what the exact wording is, just that some kind of meaningful error around withdrawing is returned given a certain condition.

如果有人想改写错误的话，测试失败会很烦人，这对我们的测试来说太详细了。我们并不_真的_关心确切的措辞是什么，只是在特定条件下返回有关撤回的某种有意义的错误。

In Go, errors are values, so we can refactor it out into a variable and have a single source of truth for it.

在 Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供单一的真实来源。

```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}
```

The `var` keyword allows us to define values global to the package.

`var` 关键字允许我们定义包的全局值。

This is a positive change in itself because now our `Withdraw` function looks very clear.

这本身就是一个积极的变化，因为现在我们的 `Withdraw` 函数看起来非常清晰。

Next we can refactor our test code to use this value instead of specific strings.

接下来我们可以重构我们的测试代码以使用这个值而不是特定的字符串。

```go
func TestWallet(t *testing.T) {

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw with funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, Bitcoin(20))
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

And now the test is easier to follow too.

现在测试也更容易遵循。

I have moved the helpers out of the main test function just so when someone opens up a file they can start reading our assertions first, rather than some helpers.

我已经将助手从主测试函数中移出，这样当有人打开文件时，他们可以先开始阅读我们的断言，而不是一些助手。

Another useful property of tests is that they help us understand the _real_ usage of our code so we can make sympathetic code. We can see here that a developer can simply call our code and do an equals check to `ErrInsufficientFunds` and act accordingly.

测试的另一个有用特性是它们帮助我们了解代码的 _real_ 用法，因此我们可以编写有同情心的代码。我们在这里可以看到，开发人员可以简单地调用我们的代码并对 `ErrInsufficientFunds` 进行相等检查并采取相应的行动。

### Unchecked errors

### 未检查的错误

Whilst the Go compiler helps you a lot, sometimes there are things you can still miss and error handling can sometimes be tricky.

虽然 Go 编译器对你有很大帮助，但有时你仍然会错过一些事情，错误处理有时会很棘手。

There is one scenario we have not tested. To find it, run the following in a terminal to install `errcheck`, one of many linters available for Go.

我们还没有测试过一种场景。要找到它，请在终端中运行以下命令以安装 `errcheck`，这是 Go 可用的众多短绒工具之一。

`go get -u github.com/kisielk/errcheck`



Then, inside the directory with your code run `errcheck .`

然后，在包含您的代码的目录中运行 `errcheck .`

You should get something like

你应该得到类似的东西



`wallet_test.go:17:18: wallet.Withdraw(Bitcoin(10))`

What this is telling us is that we have not checked the error being returned on that line of code. That line of code on my computer corresponds to our normal withdraw scenario because we have not checked that if the `Withdraw` is successful that an error is _not_ returned.

这告诉我们的是，我们还没有检查该行代码返回的错误。我电脑上的那行代码对应于我们的正常提款场景，因为我们还没有检查“提款”是否成功，是否_不_返回错误。

Here is the final test code that accounts for this.

这是解释这一点的最终测试代码。

```go
func TestWallet(t *testing.T) {

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw with funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, Bitcoin(20))
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error but didn't want one")
    }
}

func assertError(t testing.TB, got error, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %s, want %s", got, want)
    }
}
```

## Wrapping up

##  总结

### Pointers 

### 指针

* Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
* The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary \(like database connection pools\).

* 当您将值传递给函数/方法时，Go 会复制值，因此如果您正在编写一个需要改变状态的函数，您将需要它来获取指向您想要更改的内容的指针。
* Go 获取值的副本这一事实在很多时候很有用，但有时您不希望系统复制某些内容，在这种情况下，您需要传递引用。示例包括引用非常大的数据结构或仅需要一个实例的事物\（如数据库连接池\）。

### nil



* Pointers can be nil
* When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
* Useful for when you want to describe a value that could be missing

* 指针可以为零
* 当一个函数返回一个指向某物的指针时，您需要确保检查它是否为 nil 或者您可能会引发运行时异常 - 编译器在这里不会帮助您。
* 当你想描述一个可能缺失的值时很有用

### Errors

### 错误

* Errors are the way to signify failure when calling a function/method.
* By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
* This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
* [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

* 错误是调用函数/方法时表示失败的方式。
* 通过聆听我们的测试，我们得出结论，检查错误中的字符串会导致测试不稳定。所以我们重构了我们的实现以使用有意义的值来代替，这导致更容易测试代码并得出结论，这对于我们 API 的用户来说也会更容易。
* 这不是错误处理的结束，你可以做更复杂的事情，但这只是一个介绍。后面的部分将介绍更多策略。
* [不要只检查错误，优雅地处理它们](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

### Create new types from existing ones

### 从现有类型创建新类型

* Useful for adding more domain specific meaning to values
* Can let you implement interfaces

* 用于为值添加更多特定于领域的含义
* 可以让你实现接口

Pointers and errors are a big part of writing Go that you need to get comfortable with. Thankfully the compiler will _usually_ help you out if you do something wrong, just take your time and read the error. 

指针和错误是编写 Go 的重要组成部分，您需要熟悉它们。幸运的是，如果你做错了什么，编译器会_通常_帮助你，花点时间阅读错误。

