# Type, value and equality of interfaces

yourbasic.org/golang

- [Interface type](http://yourbasic.org#interface-type)
- [Structural typing](http://yourbasic.org#structural-typing)
- [The empty interface](http://yourbasic.org#the-empty-interface)
- [Interface values](http://yourbasic.org#interface-values)
- [Equality](http://yourbasic.org#equality)

## Interface type

> An interface type consists of a set of method signatures.
> A variable of interface type can hold any value that implements these methods.

In this example both `Temp` and `*Point` implement the `MyStringer` interface.

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

Actually, `*Temp` also implements `MyStringer`, since the method set of
a pointer type `*T` is the set of all methods with receiver `*T` or `T`.

When you call a method on an interface value, the method of its underlying type is executed.

```
var x MyStringer

x = Temp(24)
fmt.Println(x.String()) <span class="comment">// 24 °C</span>

x = &Point{1, 2}
fmt.Println(x.String()) <span class="comment">// (1,2)</span>
```

## Structural typing

> A type implements an interface by implementing its methods.
> No explicit declaration is required.

In fact, the `Temp`, `*Temp` and `*Point` types also implement
the standard library [`fmt.Stringer`](https://golang.org/pkg/fmt/#Stringer) interface.
The `String` method in this interface is used to print values passed as an operand
to functions such as [`fmt.Println`](https://golang.org/pkg/fmt/#Println).

```
var x MyStringer

x = Temp(24)
fmt.Println(x) <span class="comment">// 24 °C</span>

x = &Point{1, 2}
fmt.Println(x) <span class="comment">// (1,2)</span>
```

## The empty interface

The interface type that specifies no methods is known as the empty interface.

```
interface{}
```

An empty interface can hold values of any type since every type implements at least zero methods.

```
var x interface{}

x = 2.4
fmt.Println(x) <span class="comment">// 2.4</span>

x = &Point{1, 2}
fmt.Println(x) <span class="comment">// (1,2)</span>
```

The [`fmt.Println`](https://golang.org/pkg/fmt/#Println) function is a chief example.
It takes any number of arguments of any type.

```
func Println(a ...interface{}) (n int, err error)
```

## Interface values

> An **interface value** consists of a **concrete value**
> and a **dynamic type**: `[Value, Type]`

In a call to [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf), you can use `%v` to print the concrete value and `%T` to print the dynamic type.

```
var x MyStringer
fmt.Printf("%v %T\n", x, x) <span class="comment">// <nil> <nil></span>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) <span class="comment">// 24 °C main.Temp</span>

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) <span class="comment">// (1,2) *main.Point</span>

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) <span class="comment">// <nil> *main.Point</span>
```

The **zero value** of an interface type is nil, which is represented as `[nil, nil]`.

Calling a method on a nil interface is a run-time error.
However, it’s quite common to write methods that can handle
a receiver value `[nil, Type]`, where `Type` isn’t nil.

You can use **[type assertions](http://yourbasic.org/golang/type-assertion-switch/)** or
**[type switches](http://yourbasic.org/golang/type-assertion-switch/)**
to access the dynamic type of an interface value.
See [Find the type of an object](http://yourbasic.org/golang/find-type-of-object/) for more details.

## Equality

Two interface values are equal

- if they have equal concrete values**and** identical dynamic types,
- or if both are nil.

A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if

- `t`’s concrete value is equal to `x`
- **and** `t`’s dynamic type is identical to `X`.

```
var x MyStringer
fmt.Println(x == nil)<span class="comment"> // true</span>

x = (*Point)(nil)
fmt.Println(x == nil) <span class="comment">// false</span>
```

In the second print statement, the concrete value of `x` equals `nil`,
but its dynamic type is `*Point`, which is not `nil`.

> **Warning:** See [Nil is not nil](http://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/)
> for a real-world example where this definition of equality leads to puzzling results.

### Further reading

[![](http://yourbasic.org/golang/bland-face-tiny.png)](http://yourbasic.org/golang/generics/)

[Generics (alternatives and workarounds)](http://yourbasic.org/golang/generics/) discusses how interfaces, multiple functions, type assertions, reflection and code generation can be use in place of parametric polymorphism in Go.

### Go step by step

[![](http://yourbasic.org/golang/step-by-step-thumb.jpg)](http://yourbasic.org/golang/nutshells/)

Core Go concepts:
[interfaces](http://yourbasic.org/golang/interfaces-explained/),
[structs](http://yourbasic.org/golang/structs-explained/),
[slices](http://yourbasic.org/golang/slices-explained/),
[maps](http://yourbasic.org/golang/maps-explained/),
[for loops](http://yourbasic.org/golang/for-loop/),
[switch statements](http://yourbasic.org/golang/switch-statement/),
[packages](http://yourbasic.org/golang/packages-explained/).
