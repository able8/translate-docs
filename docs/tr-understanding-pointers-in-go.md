# Understanding Pointers in Go

# 理解 Go 中的指针

October  4, 2019 56.6k views

2019 年 10 月 4 日 56.6k 次观看

### Introduction

###  介绍

When you write software in Go you’ll be writing functions and methods. You pass data to these functions as *arguments*. Sometimes, the function needs a local copy of the data, and you want  the original to remain unchanged. For example, if you’re a bank, and you have a function that shows the user the changes to their balance  depending on the savings plan they choose, you don’t want to change the  customer’s actual balance before they choose a plan; you just want to  use it in calculations. This is called *passing by value*, because you’re sending the value of the variable to the function, but not the variable itself.

当您使用 Go 编写软件时，您将编写函数和方法。您将数据作为*参数*传递给这些函数。有时，该函数需要数据的本地副本，而您希望原始数据保持不变。例如，如果您是一家银行，并且您有一个功能可以根据用户选择的储蓄计划向用户显示其余额的变化，那么您不想在客户选择计划之前更改其实际余额；您只想在计算中使用它。这称为*按值传递*，因为您将变量的值发送给函数，而不是变量本身。

Other times, you may want the function to be able to alter the data  in the original variable. For instance, when the bank customer makes a  deposit to their account, you want the deposit function to be able to  access the actual balance, not a copy. In this case, you don’t need to  send the actual data to the function; you just need to tell the function where the data is located in memory. A data type called a *pointer* holds the memory address of the data, but not the data itself. The  memory address tells the function where to find the data, but not the  value of the data. You can pass the pointer to the function instead of  the data, and the function can then alter the original variable in  place. This is called *passing by reference*, because the value of the variable isn’t passed to the function, just its location.

其他时候，您可能希望函数能够更改原始变量中的数据。例如，当银行客户向其帐户存款时，您希望存款功能能够访问实际余额，而不是副本。在这种情况下，您不需要将实际数据发送给函数；你只需要告诉函数数据在内存中的位置。称为 *pointer* 的数据类型保存数据的内存地址，但不保存数据本身。内存地址告诉函数在哪里可以找到数据，而不是数据的值。您可以将指针传递给函数而不是数据，然后函数可以就地更改原始变量。这称为*通过引用*，因为变量的值没有传递给函数，只是传递给它的位置。

In this article, you will create and use pointers to share access to the memory space for a variable.

在本文中，您将创建并使用指针来共享对变量内存空间的访问。

## Defining and Using Pointers

## 定义和使用指针

When you use a pointer to a variable, there are a couple of different syntax elements that you need to understand. The first one is the use  of the ampersand (`&`). If you place an ampersand in front of a variable name, you are stating that you want to get the *address*, or a pointer to that variable. The second syntax element is the use of the asterisk (`*`) or *dereferencing* operator. When you declare a pointer variable, you follow the variable  name with the type of the variable that the pointer points to, prefixed  with an `*`, like this:

当您使用指向变量的指针时，您需要了解几个不同的语法元素。第一个是与号 (`&`) 的使用。如果您在变量名前放置一个 & 符号，则表示您要获取 *address* 或指向该变量的指针。第二个语法元素是星号 (`*`) 或 *dereferencing* 运算符的使用。当你声明一个指针变量时，你在变量名后面加上指针指向的变量的类型，加上一个 `*` 前缀，像这样：

```go
var myPointer *int32 = &someint
```

 

This creates `myPointer` as a pointer to an `int32` variable, and initializes the pointer with the address of `someint`. The pointer doesn’t actually contain an `int32`, just the address of one.

这将创建 `myPointer` 作为指向 `int32` 变量的指针，并使用 `someint` 的地址初始化该指针。指针实际上并不包含 `int32`，只是一个的地址。

Let’s take a look at a pointer to a `string`. The following code declares both a value of a string, and a pointer to a string:

让我们看一下指向 `string` 的指针。下面的代码声明了一个字符串的值和一个指向字符串的指针：

main.go

main.go

```go
package main

import "fmt"

func main() {
    var creature string = "shark"
    var pointer *string = &creature

    fmt.Println("creature =", creature)
    fmt.Println("pointer =", pointer)
}
```

 

Run the program with the following command:

使用以下命令运行程序：

```bash
go run main.go
```

 

When you run the program, it will print out the value of the  variable, as well as the address of where the variable is stored (the  pointer address). The memory address is a hexadecimal number, and not  meant to be human-readable. In practice, you’ll probably never output a  memory address to look at it. We’re showing you for illustrative  purposes. Because each program is created in its own memory space when  it is run, the value of the pointer will be different each time you run  it, and will be different than the output shown here:

当你运行程序时，它会打印出变量的值，以及变量存储的地址（指针地址）。内存地址是一个十六进制数，并不意味着是人类可读的。实际上，您可能永远不会输出内存地址来查看它。我们向您展示是为了说明目的。因为每个程序在运行时都是在自己的内存空间中创建的，所以每次运行时指针的值都会不同，并且会与此处显示的输出不同：

```
Outputcreature = shark
pointer = 0xc0000721e0
```

The first variable we defined we named `creature`, and set it equal to a `string` with the value of `shark`. We then created another variable named `pointer`. This time, we set the value of the `pointer` variable to the address of the `creature` variable. We store the address of a value in a variable by using the ampersand (`&`) symbol. This means that the `pointer` variable is storing the **address** of the `creature` variable, not the actual value.

我们定义的第一个变量命名为 `creature`，并将其设置为一个值为 `shark` 的 `string`。然后我们创建了另一个名为 `pointer` 的变量。这一次，我们将 `pointer` 变量的值设置为 `creature` 变量的地址。我们使用与号 (`&`) 符号将值的地址存储在变量中。这意味着 `pointer` 变量存储的是 `creature` 变量的 **address**，而不是实际值。

This is why when we printed out the value of `pointer`, we received the value of `0xc0000721e0`, which is the address of where the `creature` variable is currently stored in computer memory. 

这就是为什么当我们打印出 `pointer` 的值时，我们收到了 `0xc0000721e0` 的值，这是 `creature` 变量当前存储在计算机内存中的地址。

If you want to print out the value of the variable being pointed at from the `pointer` variable, you need to *dereference* that variable. The following code uses the `*` operator to dereference the `pointer` variable and retrieve its value:

如果你想打印出从 `pointer` 变量指向的变量的值，你需要*取消引用*那个变量。以下代码使用 `*` 运算符取消引用 `pointer` 变量并检索其值：

main.go

main.go

```go
package main

import "fmt"

func main() {
    var creature string = "shark"
    var pointer *string = &creature

    fmt.Println("creature =", creature)
    fmt.Println("pointer =", pointer)

    fmt.Println("*pointer =", *pointer)
}
```

 

If you run this code, you’ll see the following output:

如果您运行此代码，您将看到以下输出：

```
Outputcreature = shark
pointer = 0xc000010200
*pointer = shark
```

The last line we added now dereferences the `pointer` variable, and prints out the value that is stored at that address.

我们添加的最后一行现在取消引用 `pointer` 变量，并打印出存储在该地址的值。

If you want to modify the value stored at the `pointer` variable’s location, you can use the dereference operator as well:

如果要修改存储在 `pointer` 变量位置的值，也可以使用取消引用运算符：

main.go

main.go

```go
package main

import "fmt"

func main() {
    var creature string = "shark"
    var pointer *string = &creature

    fmt.Println("creature =", creature)
    fmt.Println("pointer =", pointer)

    fmt.Println("*pointer =", *pointer)

    *pointer = "jellyfish"
    fmt.Println("*pointer =", *pointer)
}
```

 

Run this code to see the output:

运行此代码以查看输出：

```
Outputcreature = shark
pointer = 0xc000094040
*pointer = shark
*pointer = jellyfish
```

We set the value the `pointer` variable refers to by using the asterisk (`*`) in front of the variable name, and then providing a new value of `jellyfish`. As you can see, when we print the dereferenced value, it is now set to `jellyfish`.

我们通过在变量名前面使用星号（`*`）来设置`pointer`变量引用的值，然后提供一个新的`jellyfish`值。如您所见，当我们打印取消引用的值时，它现在被设置为 `jellyfish`。

You may not have realized it, but we actually changed the value of the `creature` variable as well. This is because the `pointer` variable is actually pointing at the `creature` variable’s address. This means that if we change the value pointed at from the `pointer` variable, we also change the value of the `creature` variable.

你可能没有意识到，但我们实际上也改变了 `creature` 变量的值。这是因为 `pointer` 变量实际上指向的是 `creature` 变量的地址。这意味着如果我们改变了从 `pointer` 变量指向的值，我们也会改变 `creature` 变量的值。

main.go

main.go

```go
package main

import "fmt"

func main() {
    var creature string = "shark"
    var pointer *string = &creature

    fmt.Println("creature =", creature)
    fmt.Println("pointer =", pointer)

    fmt.Println("*pointer =", *pointer)

    *pointer = "jellyfish"
    fmt.Println("*pointer =", *pointer)

    fmt.Println("creature =", creature)
}
```

 

The output looks like this:

输出如下所示：

```
Outputcreature = shark
pointer = 0xc000010200
*pointer = shark
*pointer = jellyfish
creature = jellyfish
```

Although this code illustrates how a pointer works, this is not the  typical way in which you would use pointers in Go. It is more common to  use them when defining function arguments and return values, or using  them when defining methods on custom types. Let’s look at how you would use pointers with functions to share access to a variable.

尽管此代码说明了指针的工作原理，但这并不是您在 Go 中使用指针的典型方式。更常见的是在定义函数参数和返回值时使用它们，或者在自定义类型上定义方法时使用它们。让我们看看如何将指针与函数一起使用来共享对变量的访问。

Again, keep in mind that we are printing the value of `pointer` to illustrate that it is a pointer. In practice, you wouldn’t use the  value of a pointer, other than to reference the underlying value to  retrieve or update that value.

同样，请记住，我们正在打印 `pointer` 的值以说明它是一个指针。在实践中，除了引用底层值来检索或更新该值外，您不会使用指针的值。

## Function Pointer Receivers

## 函数指针接收器

When you write a function, you can define arguments to be passed ether by *value*, or by *reference*. Passing by *value* means that a copy of that value is sent to the function, and any changes to that argument within that function *only* effect that variable within that function, and not where it was passed from. However, if you pass by *reference*, meaning you pass a pointer to that argument, you can change the value  from within the function, and also change the value of the original  variable that was passed in. You can read more about how to define  functions in our [How To Define and Call Functions in Go](https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go).

编写函数时，您可以定义要通过 *value* 或 *reference* 传递的参数。传递 *value* 意味着该值的副本被发送到函数，并且对该函数内该参数的任何更改*仅*影响该函数内的该变量，而不是从哪里传递它。但是，如果通过 *reference* 传递，意味着传递指向该参数的指针，则可以从函数内部更改该值，也可以更改传入的原始变量的值。您可以阅读更多关于如何在我们的 [How To Define and Call Functions in Go](https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go) 中定义函数。

Deciding when to pass a pointer as opposed when to send a value is  all about knowing if you want the value to change or not. If you don’t  want the value to change, send it as a value. If you want the function  you are passing your variable to be able to change it, then you would  pass it as a pointer.

决定什么时候传递一个指针而不是什么时候发送一个值就是要知道你是否想要改变这个值。如果您不想更改值，请将其作为值发送。如果您希望传递变量的函数能够更改它，那么您可以将它作为指针传递。

To see the difference, let’s first look at a function that is passing in an argument by `value`:

为了看到不同之处，让我们首先看一个通过 `value` 传入参数的函数：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func main() {
    var creature Creature = Creature{Species: "shark"}

    fmt.Printf("1) %+v\n", creature)
    changeCreature(creature)
    fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature Creature) {
    creature.Species = "jellyfish"
    fmt.Printf("2) %+v\n", creature)
}
```

 

The output looks like this:

输出如下所示：

```
Output1) {Species:shark}
2) {Species:jellyfish}
3) {Species:shark}
```

First we created a custom type named `Creature`. It has one field named `Species`, which is a string. In the `main` function, we created an instance of our new type named `creature` and set the `Species` field to `shark`. We then printed out the variable to show the current value stored within the `creature` variable.

首先，我们创建了一个名为“Creature”的自定义类型。它有一个名为“Species”的字段，它是一个字符串。在 `main` 函数中，我们创建了一个名为 `creature` 的新类型的实例，并将 `Species` 字段设置为 `shark`。然后我们打印出变量以显示存储在 `creature` 变量中的当前值。

Next, we call `changeCreature` and pass in a copy of the `creature` variable.

接下来，我们调用 `changeCreature` 并传入 `creature` 变量的副本。

The function `changeCreature` is defined as taking one argument named `creature`, and it is of type `Creature` that we defined earlier. We then change the value of the `Species` field to `jellyfish` and print it out. Notice that within the `changeCreature` function, the value of `Species` is now `jellyfish`, and it prints out `2) {Species:jellyfish}`. This is because we are allowed to change the value within our function scope.

函数 `changeCreature` 被定义为接受一个名为 `creature` 的参数，它是我们之前定义的 `Creature` 类型。然后我们将 `Species` 字段的值更改为 `jellyfish` 并将其打印出来。请注意，在 `changeCreature` 函数中，`Species` 的值现在是 `jellyfish`，并打印出 `2) {Species:jellyfish}`。这是因为我们可以在函数范围内更改值。

However, when the last line of the `main` function prints the value of `creature`, the value of `Species` is still `shark`. The reason that the value didn’t change is because we passed the variable by *value*. This means that a copy of the value was created in memory, and passed to the `changeCreature` function. This allows us to have a function that can make changes to  any arguments passed in as needed, but will not affect any of those  variables outside of the function.

但是，当`main`函数的最后一行打印`creature`的值时，`Species`的值仍然是`shark`。值没有改变的原因是因为我们通过 *value* 传递了变量。这意味着在内存中创建了该值的副本，并传递给了 `changeCreature` 函数。这使我们可以拥有一个函数，可以根据需要更改传入的任何参数，但不会影响函数之外的任何这些变量。

Next, let’s change the `changeCreature` function to take an argument by *reference*. We can do this by changing the type from `creature` to a pointer by using the asterisk (`*`) operator. Instead of passing a `creature`, we’re now passing a pointer to a `creature`, or a `*creature`. In the previous example, `creature` is a `struct` that has a `Species` value of `shark`. `*creature` is a pointer, not a struct, so its value is a memory location, and that’s what we pass to `changeCreature()`.

接下来，让我们更改 `changeCreature` 函数以通过 *reference* 获取参数。我们可以通过使用星号 (`*`) 运算符将类型从 `creature` 更改为指针来实现。我们现在传递一个指向“creature”或“*creature”的指针，而不是传递一个 `creature`。在前面的例子中，`creature` 是一个 `struct`，它的 `Species` 值为 `shark`。 `*creature` 是一个指针，而不是一个结构体，所以它的值是一个内存位置，这就是我们传递给 `changeCreature()` 的内容。

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func main() {
    var creature Creature = Creature{Species: "shark"}

    fmt.Printf("1) %+v\n", creature)
    changeCreature(&creature)
    fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
    creature.Species = "jellyfish"
    fmt.Printf("2) %+v\n", creature)
}
```

 

Run this code to see the following output:

运行此代码以查看以下输出：

```
Output1) {Species:shark}
2) &{Species:jellyfish}
3) {Species:jellyfish}
```

Notice that now when we change the value of `Species` to `jellyfish` in the `changeCreature` function, it changes the original value defined in the `main` function as well. This is because we passed the `creature` variable by *reference*, which allows access to the original value and can change it as needed.

请注意，现在当我们在 `changeCreature` 函数中将 `Species` 的值更改为 `jellyfish` 时，它也会更改在 `main` 函数中定义的原始值。这是因为我们通过 *reference* 传递了 `creature` 变量，它允许访问原始值并可以根据需要更改它。

Therefore, if you want a function to be able to change a value, you  need to pass it by reference. To pass by reference, you pass the pointer to the variable, and not the variable itself.

因此，如果您希望函数能够更改值，则需要通过引用传递它。要通过引用传递，您将指针传递给变量，而不是变量本身。

However, sometimes you may not have an actual value defined for a pointer. In those cases, it is possible to have a [panic](https://www.digitalocean.com/community/tutorials/handling-panics-in-go) in the program. Let’s look at how that happens and how to plan for that potential problem.

但是，有时您可能没有为指针定义实际值。在这些情况下，程序中可能会出现 [恐慌](https://www.digitalocean.com/community/tutorials/handling-panics-in-go)。让我们看看这是如何发生的，以及如何为这个潜在问题制定计划。

## Nil Pointers

## 空指针

All variables in Go have a [zero value](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go#zero-values). This is true even for a pointer. If you declare a pointer to a type, but assign no value, the zero value will be `nil`. `nil` is a way to say that “nothing has been initialized” for the variable.

Go 中的所有变量都有一个 [零值](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go#zero-values)。即使对于指针也是如此。如果您声明一个指向类型的指针，但没有赋值，则零值将为“nil”。 `nil` 是表示变量“没有初始化”的一种方式。

In the following program, we are defining a pointer to a `Creature` type, but we are never instantiating that actual instance of a `Creature` and assigning the address of it to the `creature` pointer variable. The value will be `nil` and we can’t reference any of the fields or methods that would be defined on the `Creature` type:

在下面的程序中，我们定义了一个指向 `Creature` 类型的指针，但我们从来没有实例化一个 `Creature` 的实际实例并将它的地址分配给 `creature` 指针变量。该值将为 `nil`，我们不能引用将在 `Creature` 类型上定义的任何字段或方法：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func main() {
    var creature *Creature

    fmt.Printf("1) %+v\n", creature)
    changeCreature(creature)
    fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
    creature.Species = "jellyfish"
    fmt.Printf("2) %+v\n", creature)
}
```

 

The output looks like this:

输出如下所示：

```
Output1) <nil>
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x109ac86]

goroutine 1 [running]:
main.changeCreature(0x0)
        /Users/corylanou/projects/learn/src/github.com/gopherguides/learn/_training/digital-ocean/pointers/src/nil.go:18 +0x26
    main.main()
            /Users/corylanou/projects/learn/src/github.com/gopherguides/learn/_training/digital-ocean/pointers/src/nil.go:13 +0x98
        exit status 2
```

When we run the program, it printed out the value of the `creature` variable, and  the value is `<nil>`. We then call the `changeCreature` function, and when that function tries to set the value of the `Species` field, it *panics*. This is because there is no instance of the variable actually created. Because of this, the program has no where to actually store the value, so the program panics.

当我们运行程序时，它会打印出 `creature` 变量的值，该值是 `<nil>`。然后我们调用 `changeCreature` 函数，当该函数尝试设置 `Species` 字段的值时，它*恐慌*。这是因为没有实际创建的变量实例。因此，程序没有实际存储值的位置，因此程序会出现混乱。

It is common in Go that if you are receiving an argument as a  pointer, you check to see if it was nil or not before performing any  operations on it to prevent the program from panicking.

在 Go 中，如果你接收一个参数作为一个指针，你会在对其执行任何操作之前检查它是否为 nil 以防止程序恐慌，这在 Go 中很常见。

This is a common approach for checking for `nil`:

这是检查 `nil` 的常用方法：

```go
if someVariable == nil {
    // print an error or return from the method or fuction
}
```

 

Effectively you want to make sure you don’t have a `nil`  pointer that was passed into your function or method. If you do, you’ll likely just want to return, or return an error to show that an invalid  argument was passed to the function or method. The following code  demonstrates checking for `nil`:

实际上，您要确保没有传递给函数或方法的 `nil` 指针。如果这样做，您可能只想返回，或返回一个错误以表明传递给函数或方法的参数无效。以下代码演示了对 `nil` 的检查：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func main() {
    var creature *Creature

    fmt.Printf("1) %+v\n", creature)
    changeCreature(creature)
    fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
    if creature == nil {
        fmt.Println("creature is nil")
        return
    }

    creature.Species = "jellyfish"
    fmt.Printf("2) %+v\n", creature)
}
```

 

We added a check in the `changeCreature` to see if the value of the `creature` argument was `nil`. If it was, we print out “creature is nil”, and return out of the  function. Otherwise, we continue and change the value of the `Species` field. If we run the program, we will now get the following output:

我们在 `changeCreature` 中添加了一个检查，以查看 `creature` 参数的值是否为 `nil`。如果是，我们打印出“creature is nil”，然后从函数中返回。否则，我们继续并更改“物种”字段的值。如果我们运行程序，我们现在将得到以下输出：

```
Output1) <nil>
creature is nil
3) <nil>
```

Notice that while we still had a `nil` value for the `creature` variable, we are no longer panicking because we are checking for that scenario.

请注意，虽然我们仍然有 `creature` 变量的 `nil` 值，但我们不再恐慌，因为我们正在检查这种情况。

Finally, if we create an instance of the `Creature` type and assign it to the `creature` variable, the program will now change the value as expected:

最后，如果我们创建一个 `Creature` 类型的实例并将其分配给 `creature` 变量，程序现在将按预期更改该值：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func main() {
    var creature *Creature
    creature = &Creature{Species: "shark"}

    fmt.Printf("1) %+v\n", creature)
    changeCreature(creature)
    fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
    if creature == nil {
        fmt.Println("creature is nil")
        return
    }

    creature.Species = "jellyfish"
    fmt.Printf("2) %+v\n", creature)
}
```

 

Now that we have an instance of the `Creature` type, the program will run and we will get the following expected output:

现在我们有一个 `Creature` 类型的实例，程序将运行，我们将获得以下预期输出：

```
Output1) &{Species:shark}
2) &{Species:jellyfish}
3) &{Species:jellyfish}
```

When you are working with pointers, there is a potential for the  program to panic. To avoid panicking, you should check to see if a  pointer value is `nil` prior to trying to access any of the fields or methods defined on it.

当您使用指针时，程序可能会出现恐慌。为避免恐慌，在尝试访问定义在其上的任何字段或方法之前，您应该检查指针值是否为“nil”。

Next, let’s look at how using pointers and values affects defining methods on a type.

接下来，让我们看看使用指针和值如何影响在类型上定义方法。

## Method Pointer Receivers

## 方法指针接收器

A *receiver* in go is the argument that is defined in a method declaration. Take a look at the following code:

go 中的 *receiver* 是在方法声明中定义的参数。看看下面的代码：

```go
type Creature struct {
    Species string
}

func (c Creature) String() string {
    return c.Species
}
```

 

The receiver in this method is `c Creature`. It is stating that the instance of `c` is of type `Creature` and you will reference that type via that instance variable.

这个方法中的接收者是`c Creature`。它表明 `c` 的实例属于 `Creature` 类型，您将通过该实例变量引用该类型。

Just like the behavior of functions is different based on whether you send in an argument as a pointer or a value, methods also have  different behavior. The big difference is that if you define a method  with a value receiver, you are not able to make changes to the instance  of that type that the method was defined on.

就像函数的行为根据您将参数作为指针还是值发送而不同一样，方法也有不同的行为。最大的区别在于，如果您定义了一个带有值接收器的方法，您将无法对定义该方法的该类型的实例进行更改。

There will be times that you would like your method to be able to  update the instance of the variable that you are using. To allow for  this, you would want to make the receiver a pointer.

有时您希望您的方法能够更新您正在使用的变量的实例。为了实现这一点，您需要让接收者成为一个指针。

Let’s add a `Reset` method to our `Creature` type that will set the `Species` field to an empty string:

让我们向我们的 `Creature` 类型添加一个 `Reset` 方法，它将把 `Species` 字段设置为空字符串：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func (c Creature) Reset() {
    c.Species = ""
}

func main() {
    var creature Creature = Creature{Species: "shark"}

    fmt.Printf("1) %+v\n", creature)
    creature.Reset()
    fmt.Printf("2) %+v\n", creature)
}
```

 

If we run the program, we will get the following output:

如果我们运行程序，我们将得到以下输出：

```
Output1) {Species:shark}
2) {Species:shark}
```

Notice that even though in the `Reset` method we set the value of `Species` to an empty string, that when we print out the value of our `creature` variable in the `main` function, the value is still set to ` shark`. This is because we defined the `Reset` method has having a `value` receiver. This means that the method will only have access to a *copy* of the `creature` variable.

请注意，即使在 `Reset` 方法中我们将 `Species` 的值设置为空字符串，当我们在 `main` 函数中打印出 `creature` 变量的值时，该值仍然设置为 `鲨鱼`。这是因为我们定义了 `Reset` 方法有一个 `value` 接收器。这意味着该方法只能访问 `creature` 变量的 *copy*。

If we want to be able to modify the instance of the `creature` variable in the methods, we need to define them as having a `pointer` receiver:

如果我们希望能够在方法中修改 `creature` 变量的实例，我们需要将它们定义为具有 `pointer` 接收器：

main.go

main.go

```go
package main

import "fmt"

type Creature struct {
    Species string
}

func (c *Creature) Reset() {
    c.Species = ""
}

func main() {
    var creature Creature = Creature{Species: "shark"}

    fmt.Printf("1) %+v\n", creature)
    creature.Reset()
    fmt.Printf("2) %+v\n", creature)
}
```

 

Notice that we now added an asterisk (`*`) in front of the `Creature` type in when we defined the `Reset` method. This means that the instance of `Creature` that is passed to the `Reset` method is now a pointer, and as such when we make changes it will affect the original instance of that variables.

请注意，我们现在在定义 `Reset` 方法时在 `Creature` 类型前面添加了一个星号（`*`）。这意味着传递给 `Reset` 方法的 `Creature` 实例现在是一个指针，因此当我们进行更改时，它将影响该变量的原始实例。

```
Output1) {Species:shark}
2) {Species:}
```

The `Reset` method has now changed the value of the `Species` field.

`Reset` 方法现在更改了 `Species` 字段的值。

## Conclusion

##  结论

Defining a function or method as a pass by *value* or pass by *reference* will affect what parts of your program are able to make changes to  other parts. Controlling when that variable can be changed will allow  you to write more robust and predictable software. Now that you have  learned about pointers, you can see how they are used in interfaces as  well. 

将函数或方法定义为通过*值* 或通过*引用* 将影响程序的哪些部分能够对其他部分进行更改。控制何时可以更改该变量将使您能够编写更健壮和可预测的软件。现在您已经了解了指针，您也可以了解它们在接口中的使用方式。

