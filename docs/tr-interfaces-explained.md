# Type, value and equality of interfaces

# 接口的类型、值和相等性

yourbasic.org/golang

- [Interface type](http://yourbasic.org#interface-type)
- [Structural typing](http://yourbasic.org#structural-typing)
- [The empty interface](http://yourbasic.org#the-empty-interface)
- [Interface values](http://yourbasic.org#interface-values)
- [Equality](http://yourbasic.org#equality)


## Interface type

## 接口类型

> An interface type consists of a set of method signatures.
> A variable of interface type can hold any value that implements these methods.

> 接口类型由一组方法签名组成。
> 接口类型的变量可以保存实现这些方法的任何值。

In this example both `Temp` and `*Point` implement the `MyStringer` interface.

在这个例子中，`Temp` 和 `*Point` 都实现了 `MyStringer` 接口。

```
type MyStringer interface {
    String() string
}
```


```
type Temp int

func (t Temp) String() string {
    return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
    x, y int
}

func (p *Point) String() string {
    return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
```


Actually, `*Temp` also implements `MyStringer`, since the method set of a pointer type `*T` is the set of all methods with receiver `*T` or `T`.

实际上，`*Temp` 也实现了 `MyStringer`，因为方法集指针类型 `*T` 是具有接收器 `*T` 或 `T` 的所有方法的集合。

When you call a method on an interface value, the method of its underlying type is executed.

当您对接口值调用方法时，会执行其基础类型的方法。

```
var x MyStringer

x = Temp(24)
fmt.Println(x.String()) // 24 °C

x = &Point{1, 2}
fmt.Println(x.String()) // (1,2)
```


## Structural typing

## 结构类型

> A type implements an interface by implementing its methods.
> No explicit declaration is required.

> 类型通过实现其方法来实现接口。
> 无需明确声明。

In fact, the `Temp`, `*Temp` and `*Point` types also implement the standard library [`fmt.Stringer`](https://golang.org/pkg/fmt/#Stringer) interface.
The `String` method in this interface is used to print values passed as an operand to functions such as [`fmt.Println`](https://golang.org/pkg/fmt/#Println).

事实上，`Temp`、`*Temp` 和`*Point` 类型也实现了标准库 [`fmt.Stringer`](https://golang.org/pkg/fmt/#Stringer) 接口。
此接口中的 String 方法用于打印作为操作数传递的值到诸如 [`fmt.Println`](https://golang.org/pkg/fmt/#Println) 之类的函数。

```
var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```


## The empty interface

## 空接口

The interface type that specifies no methods is known as the empty interface.

不指定方法的接口类型称为空接口。

```
interface{}
```


An empty interface can hold values of any type since every type implements at least zero methods.

空接口可以保存任何类型的值，因为每种类型都至少实现零个方法。

```
var x interface{}

x = 2.4
fmt.Println(x) // 2.4

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```


The [`fmt.Println`](https://golang.org/pkg/fmt/#Println) function is a chief example. It takes any number of arguments of any type.

[`fmt.Println`](https://golang.org/pkg/fmt/#Println) 函数就是一个主要的例子。它接受任意数量的任意类型的参数。

```
func Println(a ...interface{}) (n int, err error)
```


## Interface values

## 接口值

> An **interface value** consists of a **concrete value** and a **dynamic type**: `[Value, Type]`

> **接口值**由**具体值**组成和**动态类型**：`[Value, Type]`

In a call to [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf), you can use `%v` to print the concrete value and `%T` to print the dynamic type .

在调用 [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) 时，可以使用 `%v` 打印具体值，使用 `%T` 打印动态类型.

```
var x MyStringer
fmt.Printf("%v %T\n", x, x) // <nil> <nil>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
```


The **zero value** of an interface type is nil, which is represented as `[nil, nil]`.

接口类型的**零值**为nil，表示为`[nil, nil]`。

Calling a method on a nil interface is a run-time error. However, it’s quite common to write methods that can handle  a receiver value `[nil, Type]`, where `Type` isn’t nil.

在 nil 接口上调用方法是一个运行时错误。然而，编写可以处理的方法是很常见的接收器值 `[nil, Type]`，其中 `Type` 不是 nil。

You can use **[type assertions](http://yourbasic.org/golang/type-assertion-switch/)** or **[type switches](http://yourbasic.org/golang/type-assertion-switch/)**
to access the dynamic type of an interface value.
See [Find the type of an object](http://yourbasic.org/golang/find-type-of-object/) for more details.

您可以使用 **[type assertions](http://yourbasic.org/golang/type-assertion-switch/)** 或**[类型开关](http://yourbasic.org/golang/type-assertion-switch/)**
访问接口值的动态类型。
有关更多详细信息，请参阅[查找对象的类型](http://yourbasic.org/golang/find-type-of-object/)。

## Equality

##  平等

Two interface values are equal

两个接口值相等

- if they have equal concrete values **and** identical dynamic types,
- or if both are nil.

- 如果它们具有相同的具体值**和**相同的动态类型，
- 或者如果两者都为零。

A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if

接口类型为 T 的值 t 和非接口类型 X 的值 x 相等，如果

- `t`’s concrete value is equal to `x`
- **and** `t`’s dynamic type is identical to `X`.

- `t` 的具体值等于 `x`
- **和** `t` 的动态类型与 `X` 相同。

```
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
```


In the second print statement, the concrete value of `x` equals `nil`, but its dynamic type is `*Point`, which is not `nil`.

在第二个打印语句中，`x` 的具体值等于 `nil`，但它的动态类型是 `*Point`，而不是 `nil`。

> **Warning:** See [Nil is not nil](http://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/)
> for a real-world example where this definition of equality leads to puzzling results.

> **警告：** 参见 [Nil is not nil](http://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/)
> 一个真实世界的例子，在这个例子中，平等的定义导致了令人费解的结果。

### Further reading 

### 进一步阅读



[Generics (alternatives and workarounds)](http://yourbasic.org/golang/generics/) discusses how interfaces, multiple functions, type assertions, reflection and code generation can be use in place of parametric polymorphism in Go.

[泛型（替代方案和解决方法）](http://yourbasic.org/golang/generics/) 讨论了如何使用接口、多个函数、类型断言、反射和代码生成来代替 Go 中的参数多态性。

### Go step by step

### 一步一步来

[![](http://yourbasic.org/golang/step-by-step-thumb.jpg)](http://yourbasic.org/golang/nutshells/)

Core Go concepts:
[interfaces](http://yourbasic.org/golang/interfaces-explained/),
[structs](http://yourbasic.org/golang/structs-explained/),
[slices](http://yourbasic.org/golang/slices-explained/),
[maps](http://yourbasic.org/golang/maps-explained/),
[for loops](http://yourbasic.org/golang/for-loop/),
[switch statements](http://yourbasic.org/golang/switch-statement/),
[packages](http://yourbasic.org/golang/packages-explained/). 

