## map[string]interface{} in Go

From: https://bitfieldconsulting.com/golang/map-string-interface

What is a `map[string]interface{}` in Go, and why is it so useful? How do we deal with maps of `string` to `interface{}` in our programs? What the heck is an `interface{}`, anyway? Let's find out, with renowned Go teacher John Arundel, aka [@bitfield](https://twitter.com/bitfield), author of perhaps the friendliest and most helpful introductory Go book series, [For the Love of Go](https ://bitfieldconsulting.com/books).

什么是 Go 中的“map[string]interface{}”，为什么它如此有用？我们如何处理程序中从 `string` 到 `interface{}` 的映射？无论如何，“interface{}”到底是什么？让我们和著名的 Go 老师约翰·阿伦德尔一起找出来，又名 [@bitfield](https://twitter.com/bitfield)，也许是最友好和最有帮助的围棋入门系列丛书 [For the Love of Go](https://twitter.com/bitfield) 的作者://bitfieldconsulting.com/books）。

‘For the Love of Go’ is a series of fun, easy-to-follow ebooks introducing software engineering in Go.

“For the Love of Go”是一系列有趣、易于阅读的电子书，介绍了 Go 中的软件工程。

This is part 6 of a [Golang tutorial](https://bitfieldconsulting.com/golang) series on [maps](https://bitfieldconsulting.com/golang/tag/map). Check out the other posts in the series:
- [Declaring and initializing maps](https://bitfieldconsulting.com/golang/map-declaring-initializing)
- [Map types in Go](https://bitfieldconsulting.com/golang/map-types)
- [Storing and retrieving map values](https://bitfieldconsulting.com/golang/storing-retrieving-map-values)
- [Finding whether a Go map key exists](https://bitfieldconsulting.com/golang/map-key-exists)
- [Iterating over Golang maps](https://bitfieldconsulting.com/golang/map-iteration)
- [Frequently asked questions about Go maps](https://bitfieldconsulting.com/golang/map-faq)

We've talked about lots of different kinds of Go maps in this series, but here's one you'll meet quite often: the `map[string]interface{}`. Let's take a look at what this does and how to use it.

我们在本系列中讨论了许多不同类型的 Go maps，但这里有一个你会经常遇到的：`map[string]interface{}`。让我们来看看它的作用以及如何使用它。

## Golang 'map string interface' example

Following our diner theme for these tutorials, or perhaps channeling [Ron Swanson](https://www.youtube.com/watch?v=gvvaZCuvHnc), here's an example of a `map[string]interface{}` literal:

遵循我们这些教程的晚餐主题，或者可能引导 [Ron Swanson](https://www.youtube.com/watch?v=gvvaZCuvHnc)，这里有一个 `map[string]interface{}` 文字示例：

```
foods := map[string]interface{}{
   "bacon": "delicious",
   "eggs": struct {
     source string
     price  float64
   }{"chicken", 1.75},
   "steak": true,
 }
```


## What is a `map[string]interface{}`?

If you've read the earlier tutorial in this series on [map types](https://bitfieldconsulting.com/golang/map-types), you'll know how to read this code right away. The type of the `foods` variable in the above example is a map where the keys are strings, and the values are of type `interface{}`.

如果您已阅读本系列中有关 [map 类型](https://bitfieldconsulting.com/golang/map-types) 的早期教程，您将立即知道如何阅读此代码。上面例子中`foods`变量的类型是一个map，其中键是字符串，值是`interface{}`类型。

So what's that? Go *interfaces* are worthy of a tutorial series in themselves, though it's one of those topics that seems a lot  more complicated than it actually is; it's just a little unfamiliar to  most of us at first. Suffice it to say here that an interface is a way  of referring to a value without specifying its type. Instead, the  interface specifies what *methods* it has; for example, the widely-used `io.Reader` interface type tells you that a value of that type has a `Read()` method with a certain signature.

那是什么？ Go *interfaces* 本身就值得一个教程系列，尽管它是那些看起来比实际复杂得多的主题之一；刚开始我们大多数人都有些陌生。这里只要说接口是一种引用值而不指定其类型的方式就足够了。相反，接口指定了它拥有的*方法*；例如，广泛使用的`io.Reader` 接口类型告诉您该类型的值具有带有特定签名的`Read()` 方法。

So what is `interface{}`? Pronounced 'empty interface', it's the interface that specifies no methods at all! (Note that this doesn't mean that `interface{}` values must *have* no methods; it simply doesn't say anything at all about what methods they may or may not have. In the words of a [Go proverb] (https://go-proverbs.github.io/), `interface{}` says nothing.)

那么什么是“interface{}”？发音为“空接口”，它是根本不指定任何方法的接口！ （请注意，这并不意味着`interface{}` 值必须*有*没有方法；它根本没有说明它们可能有或可能没有什么方法。用[Go 谚语]的话来说（https://go-proverbs.github.io/），`interface{}` 什么也没说。）

## Why is `interface{}` so useful?

## 为什么`interface{}` 如此有用？

What's the point of `interface{}`, then, if it doesn't  tell us anything about the value? Well, that's precisely why it's  useful: it can refer to anything! The type `interface{}` applies to *any* value. A variable declared as `interface{}` can hold a string value, an integer, any kind of struct, a pointer to an `os.File`, or indeed anything you can think of.

那么，如果`interface{}` 没有告诉我们任何有关价值的信息，它的意义何在？嗯，这正是它有用的原因：它可以指代任何东西！ `interface{}` 类型适用于 *any* 值。声明为 `interface{}` 的变量可以保存字符串值、整数、任何类型的结构、指向 `os.File` 的指针，或者你能想到的任何东西。

Suppose we need to write a function that prints out the value passed  to it, but we don't know in advance what type this value would be. This  is a job for `interface{}`:

假设我们需要编写一个函数来打印传递给它的值，但我们事先不知道这个值是什么类型。这是`interface{}` 的工作：

```
func printAnything(v interface{})
```


Indeed, `fmt.Println` is defined in a very similar way, for exactly this reason:

事实上，`fmt.Println` 的定义方式非常相似，正是出于这个原因：

```
func Println(a ...interface{}) ...
```


## `map[string]interface{}` and arbitrary data

## `map[string]interface{}` 和任意数据

Similarly, if we want a collection of *any type of thing*, each one identified by a string, which is a convenient way to organise arbitrary data, we can do that with a `map[string]interface{}`. In fact, we just described the schema of JSON objects, for example. Take this raw JSON data:

类似地，如果我们想要一个*任何类型的事物*的集合，每个事物都由一个字符串标识，这是一种组织任意数据的便捷方式，我们可以使用`map[string]interface{}`来实现。事实上，我们只是描述了 JSON 对象的模式，例如。获取这个原始 JSON 数据：

```
{
    "name":"John",
    "age":29,
    "hobbies":[
       "martial arts",
       "breakfast foods",
       "piano"
    ]
 }
```


Overlooking the obviously fictitious age for the  moment, we can see that this is a collection of things identified by  string keys, but what kind of things? We have a string, an integer, and  an array of strings.

暂时放过这个明显是虚构的时代，我们可以看出这是一个字符串键标识的东西的集合，但到底是什么东西呢？我们有一个字符串、一个整数和一个字符串数组。

Supposing we needed to translate this into a Go struct value, we could define a type like this:

假设我们需要将其转换为 Go 结构体值，我们可以定义这样的类型：

```
type Person struct {
     Name    string
     Age     int
     Hobbies []string
 }
```


Great. But this requires that we know the schema of the object in advance. What if someone gives us *arbitrary* JSON data, and we need to unmarshal it into a Go value? How can we  possibly do that, given that all we know is that it's a map of strings  to objects of any type?

太好了。但这需要我们提前知道对象的模式。如果有人给了我们*任意*的 JSON 数据，而我们需要将其解组为 Go 值怎么办？我们怎么可能做到这一点，因为我们只知道它是一个字符串到任何类型对象的映射？

## Decoding JSON data to `map[string]interface{}`

## 将 JSON 数据解码为 `map[string]interface{}`

Suppose that we have the biographically questionable JSON data about me stored in a variable called `data`. How can we unmarshal this into a Go variable so that we can start looking at it? What type would that variable need to be?

假设我们有关于我的传记上有问题的 JSON 数据存储在一个名为“data”的变量中。我们如何将其解组为 Go 变量以便我们可以开始查看它？该变量需要是什么类型？

```
var p map[string]interface{}
err = json.Unmarshal(data, &p)
```


Provided there are no errors, the `p` variable now contains our arbitrary data. Success! But, given that we  know nothing at all about the type of each value in the map, what can we usefully do with it?

如果没有错误，`p` 变量现在包含我们的任意数据。成功！但是，鉴于我们对地图中每个值的类型一无所知，我们可以用它做什么？

## Using `map[string]interface{}` data

## 使用 `map[string]interface{}` 数据

One thing we can do is use a *type switch* to do different things depending on the type of the value. Here's an example:

我们可以做的一件事是使用 *type switch* 根据值的类型做不同的事情。下面是一个例子：

```
for k, v := range p {
   switch c := v.(type) {
   case string:
     fmt.Printf("Item %q is a string, containing %q\n", k, c)
   case float64:
     fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
   default:
     fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
   }
 }
```


The special syntax `switch c := v.(type)` tells us that this is a type switch, meaning that Go will try to match the type of `v` to each case in the switch statement. For example, the first case will be executed if `v` is a string:

特殊语法 `switch c := v.(type)` 告诉我们这是一个类型 switch，这意味着 Go 将尝试将 `v` 的类型与 switch 语句中的每个 case 匹配。例如，如果 `v` 是一个字符串，将执行第一种情况：

```
Item "name" is a string, containing "John"
```


In each case, the variable `c` receives the value of `v`, but converted to the relevant type. So in the `string` case, `c` will be of type `string`.

在每种情况下，变量 c 接收 v 的值，但转换为相关类型。所以在 `string` 的情况下，`c` 将是 `string` 类型。

The `float64` case will match when `v` is a `float64`:

当 `v` 是 `float64` 时，`float64` 大小写将匹配：

```
Looks like item "age" is a number, specifically 29.000000
```


You might be puzzled that the whole-number value `29` was unmarshaled into a `float64`, but that's normal. All JSON numbers are treated as `float64` by `json.Unmarshal`. It's the most general of Go's numeric types.

您可能对整数值“29”被解组为“float64”感到困惑，但这是正常的。所有 JSON 数字都被 `json.Unmarshal` 视为 `float64`。它是 Go 中最通用的数字类型。

Finally, if no other case matches, the `default` case is activated:

最后，如果没有其他 case 匹配，则激活“default” case：

```
Not sure what type item "hobbies" is, but I think it might be []interface {}
```


(The format specifier `%T` to `fmt.Printf` prints the *type* of its value, which is sometimes handy. In this case we can see that the value of `"hobbies"` is a slice of arbitrary data, which makes sense.)

（格式说明符 `%T` 到 `fmt.Printf` 打印其值的 *type*，这有时很方便。在这种情况下，我们可以看到 `"hobbies"` 的值是任意数据的切片，这是有道理的。）

## When to use `map[string]interface{}`

## 何时使用 `map[string]interface{}`

As we've seen, the "map of string to empty interface" type is very  useful when we need to deal with data that comes from outside the Go  world; for example, arbitrary JSON data of unknown schema. Many web APIs return data like this, for example.

正如我们所见，当我们需要处理来自 Go 世界之外的数据时，“字符串到空接口的映射”类型非常有用；例如，未知模式的任意 JSON 数据。例如，许多 Web API 返回这样的数据。

It's also extremely common when writing Terraform providers, which  makes sense; Terraform resources are also essentially maps of strings to arbitrary data. It's recursive, too; the 'arbitrary data' is also often a map of strings to more arbitrary data. It's `map[string]interface{}` all the way down!

在编写 Terraform 提供程序时也非常常见，这是有道理的； Terraform 资源本质上也是字符串到任意数据的映射。它也是递归的； “任意数据”通常也是字符串到更多任意数据的映射。一直是`map[string]interface{}`！

Configuration files, too, generally have this kind of schema. You can think of YAML or [CUE](https://bitfieldconsulting.com/golang/cuelang-exciting) files as being maps of string to empty interface, just like JSON. So  when we're dealing with structured data of any kind, we'll often use  this type in Go programs.

配置文件也通常具有这种模式。您可以将 YAML 或 [CUE](https://bitfieldconsulting.com/golang/cuelang-exciting) 文件视为字符串到空接口的映射，就像 JSON 一样。所以当我们处理任何类型的结构化数据时，我们会经常在 Go 程序中使用这种类型。

## And when not to

## 什么时候不

​            ![Go;is there anything it  can’t  do?](https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/1591464379231-98VVVIYN4N30XAV27XXS/travel_adapter.jpg?format=1000w)

Go; is there anything it *can’t* do?

去；有什么它*不能*做的吗？

A `map[string]interface{}` is like one of those universal travel adapters, that plugs into any  kind of socket and works with any voltage. You can use it to protect  your own vulnerable programs from damage caused by weird, alien data.

`map[string]interface{}` 就像一种通用旅行适配器，可以插入任何类型的插座并在任何电压下工作。您可以使用它来保护自己的易受攻击的程序免受奇怪的外来数据造成的损害。

Should you use `map[string]interface{}` values within your own programs, when there's no need to handle arbitrary input data? No,  you shouldn't. While it might seem convenient to not have to explicitly  define the schema of your objects, that can lead to all kinds of  problems.

当不需要处理任意输入数据时，您是否应该在自己的程序中使用 `map[string]interface{}` 值？不，你不应该。虽然不必显式定义对象的模式看起来很方便，但这可能会导致各种问题。

For one thing, since `interface{}` proverbially says  nothing, whenever we deal with a value of this type, we have to use  protective type assertions to prevent panics:

一方面，因为众所周知`interface{}` 什么也没说，每当我们处理这种类型的值时，我们必须使用保护性类型断言来防止恐慌：

```go
if _, ok := x.(string); !ok {
   log.Fatal("oh no")
 }
```


In other words, it's a lot more difficult to write safe, reliable  programs that operate on such maps. If your library produces data of  this kind, it will hardly endear you to users. Instead, just use a plain old struct, which enables compile-time type checking and is much more  convenient to deal with.

换句话说，编写在此类地图上运行的安全、可靠的程序要困难得多。如果您的图书馆产生此类数据，则用户几乎不会喜欢您。相反，只需使用一个普通的旧结构，它启用编译时类型检查并且更方便处理。

Just like a travel adapter, `map[string]interface{}` is a  bit wonky and awkward to use when you're at home and you can rely on  your sockets all having the expected voltage and pin schema. But when  you're in contact with other worlds, outside the warm, safe cocoon of  Go's type system, `map[string]interface{}` is perhaps the ultimate travel accessory. Use it well!

就像旅行适配器一样，“map[string]interface{}”在您在家时使用起来有点不稳定和笨拙，您可以依靠所有具有预期电压和引脚架构的插座。但是，当您与其他世界接触时，在 Go 类型系统温暖、安全的茧之外，`map[string]interface{}` 可能是终极旅行配件。好好利用吧！

## Next

## 下一个

To wrap up this series, we'll look at a bunch of [frequently asked questions about Go maps](https://bitfieldconsulting.com/golang/map-faq).

为了结束这个系列，我们将看看一堆 [关于 Go maps的常见问题](https://bitfieldconsulting.com/golang/map-faq)。

## Looking for help with Go?

## 正在寻找 Go 方面的帮助？

If you found this article useful, and you'd like to learn more about  Go, then I can help! I offer one-to-one or group mentoring in Go  development, for complete beginners through to experienced Gophers. Check out my [Learn Golang with mentoring](https://bitfieldconsulting.com/golang/learn) page to find out more, or email [go@bitfieldconsulting.com](mailto:go@bitfieldconsulting.com). I've also written a series of friendly little ebooks called [For the Love of Go](https://bitfieldconsulting.com/books) which are available to buy now.

如果你觉得这篇文章有用，并且你想了解更多关于 Go 的知识，那么我可以提供帮助！我在 Go 开发中提供一对一或小组指导，从初学者到经验丰富的 Gophers。查看我的 [通过指导学习 Golang](https://bitfieldconsulting.com/golang/learn) 页面以了解更多信息，或发送电子邮件至 [go@bitfieldconsulting.com](mailto:go@bitfieldconsulting.com)。我还编写了一系列名为 [For the Love of Go](https://bitfieldconsulting.com/books) 的友好小电子书，现在可以购买。

Gopher image by [egonelbre](https://github.com/egonelbre/gophers)

[John Arundel](https://bitfieldconsulting.com/golang?author=5e10bdc11264f20181591485) 
[约翰·阿伦德尔](https://bitfieldconsulting.com/golang?author=5e10bdc11264f20181591485)


