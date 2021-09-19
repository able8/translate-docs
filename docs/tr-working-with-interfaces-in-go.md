# [Exposing interfaces in Go](https://www.efekarakus.com/golang/2019/12/29/working-with-interfaces-in-go.html)

# [在 Go 中暴露接口](https://www.efekarakus.com/golang/2019/12/29/working-with-interfaces-in-go.html)

      Published: Dec 29, 2019

发布时间：2019 年 12 月 29 日

[Interfaces](https://golang.org/doc/effective_go.html#interfaces_and_types) are my favorite feature in Go. An interface type represents a set of  methods. Unlike most other languages, you don’t have to explicitly  declare that a type *implements* an interface. A struct `S` implements the interface `I` *implicitly* if `S` defines the methods that `I` requires.

[接口](https://golang.org/doc/effective_go.html#interfaces_and_types) 是我最喜欢的 Go 功能。接口类型表示一组方法。与大多数其他语言不同，您不必显式声明类型*实现*一个接口。如果`S` 定义了`I` 所需的方法，则结构`S` 将*隐式* 实现接口`I`。

Writing good interfaces is difficult. It’s easy to [pollute](https://rakyll.org/interface-pollution/) the “API” of a package by exposing broad or unnecessary interfaces. In  this article, we’ll explain the reasoning behind existing guidelines for interfaces and supplement them with examples from the standard library.

编写好的接口很困难。通过暴露广泛的或不必要的接口，很容易[污染](https://rakyll.org/interface-pollution/) 包的“API”。在本文中，我们将解释现有接口指南背后的原因，并用标准库中的示例对其进行补充。

#### **[“The bigger the interface, the weaker the abstraction”](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)**

#### **[“界面越大，抽象越弱”](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)**

It’s unlikely that you’ll be able to find multiple types that can  implement a large interface. For that reason, ”interfaces with only one  or two methods are common in Go code”. Instead of declaring large *public* interfaces, consider depending on or returning an explicit type.

您不太可能找到可以实现大型接口的多种类型。出于这个原因，“只有一种或两种方法的接口在 Go 代码中很常见”。与其声明大型 *public* 接口，不如考虑依赖或返回显式类型。

The [`io.Reader`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) and [`io.Writer`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) interfaces are the usual examples for powerful interfaces.

[`io.Reader`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) 和 [`io.Writer`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) 接口是强大接口的常见示例。

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

After grepping the std lib, I found 81 structs across 30 packages that implement an `io.Reader`, and 99 methods or functions that consume it across 39 packages.

在对标准库进行 grep 之后，我发现 30 个包中的 81 个结构实现了 `io.Reader`，以及 99 个方法或函数在 39 个包中使用它。

#### **[“Go interfaces \*generally\* belong in the package that uses values of the interface type, not the package that implements those values”](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)**

#### **[“Go 接口\*一般\*属于使用接口类型值的包，而不是实现这些值的包”](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)**

By defining the interface in the package that actually uses it, we  can let the client define the abstraction rather than the provider  dictating the abstraction to all of its clients.

通过在实际使用它的包中定义接口，我们可以让客户定义抽象，而不是提供者将抽象规定给它的所有客户。

An example is the [`io.Copy`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L363) function. It accepts both the `Writer` and `Reader` interfaces as arguments defined in the same package.

一个例子是 [`io.Copy`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L363) 函数。它接受`Writer` 和`Reader` 接口作为在同一个包中定义的参数。

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

Another example from a different package is the[`color.Color`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L10-L19) interface. The  [`Index`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L292-L308) method of the  `color.Palette` type depends on it, allowing it to accept any struct that implements the `Color` interface.

另一个来自不同包的示例是 [`color.Color`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L10-L19) 接口。 `color.Palette` 类型的 [`Index`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L292-L308) 方法依赖于它，允许它接受任何实现 `Color` 接口的结构。

```
func (p Palette) Index(c Color) int
```

#### **[“If  a type exists only to implement an interface and will never have  exported methods beyond that interface, there is no need to export the  type itself”](https://golang.org/doc/effective_go.html#generality)**

#### **[“如果一个类型只是为了实现一个接口而存在，并且永远不会在该接口之外导出方法，则不需要导出类型本身”](https://golang.org/doc/effective_go.html#generality)**

The previous guideline from the [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces) mentions that:

[CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces) 的先前指南提到：

> The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations  without requiring extensive refactoring.

> 实现包应该返回具体的（通常是指针或结构）类型：这样，新方法可以添加到实现中，而无需大量重构。

Complementing it with the statement from [EffectiveGo](https://golang.org/doc/effective_go.html#generality), we can see the full picture where it’s okay to define an interface in the producer package:

与[EffectiveGo](https://golang.org/doc/effective_go.html#generality)的声明相辅相成，我们可以看到在生产者包中定义接口的全貌：

> If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the  type itself.” 

> 如果一个类型只是为了实现一个接口而存在，并且永远不会在该接口之外导出方法，那么就没有必要导出该类型本身。”

An example is the [`rand.Source`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L23-L28) interface which is returned by [`rand .NewSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L41-L48). The underlying struct [`rngSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rng.go#L180) within the constructor only exports the methods needed for the `Source` and `Source64` interfaces so the type itself is not exposed.

一个例子是由 [`rand. .NewSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L41-L48)。构造函数中的底层结构[`rngSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rng.go#L180) 仅导出 `Source` 所需的方法和 `Source64` 接口，因此类型本身不会公开。

The `rand` package has two other types that implement the interface: [`lockedSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L382), and [`Rand`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L51) (the latter is exposed because it has other public methods).

`rand` 包有另外两种实现接口的类型：[`lockedSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L382)，和[`Rand`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L51)（后者暴露是因为它有其他公共方法)。

**What’s the benefit of returning an interface over a concrete type anyway?**

**无论如何通过具体类型返回接口有什么好处？**

Returning an interface allows you to have functions that can return multiple concrete types. For example, the [`aes.NewCipher`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher.go#L32) constructor returns a [`cipher.Block`] (https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/cipher/cipher.go#L15) interface. If you look [within the constructor](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher_asm.go#L33-L54), you can see that two different structs are returned.

返回接口允许您拥有可以返回多个具体类型的函数。例如，[`aes.NewCipher`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher.go#L32) 构造函数返回一个 [`cipher.Block`] （https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/cipher/cipher.go#L15)接口。如果您查看[在构造函数中](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher_asm.go#L33-L54)，您可以看到返回了两个不同的结构。

```
func newCipher(key []byte) (cipher.Block, error) {
  ...
  c := aesCipherAsm{aesCipher{make([]uint32, n), make([]uint32, n)}}
  ...
  if supportsAES && supportsGFMUL {
    // Returned type is aesCipherGCM.
    return &aesCipherGCM{c}, nil
  }
  // Returned type is aesCipherAsm.
  return &c, nil
}
```

Note that in the previous example, the interface was defined in the producer package `rand`. However in this example, the returned type is defined in a different package `cipher`.

请注意，在前面的示例中，接口是在生产者包“rand”中定义的。然而，在这个例子中，返回的类型是在不同的包 `cipher` 中定义的。

**In my short experience …**

**在我短暂的经验中......**

… this pattern is a lot more difficult to execute than the previous  ones. In the early phases of development the needs of a client package  evolves quickly. Hence, it results in modifying the producer package. If the returned type is an interface, it  slowly becomes too big to the  point where returning a concrete type would make more sense.

……这种模式比以前的模式更难执行。在开发的早期阶段，客户端包的需求会迅速发展。因此，它导致修改生产者包。如果返回的类型是一个接口，它会慢慢变得太大，以至于返回一个具体类型会更有意义。

My hypothesis for the mechanics of this pattern is:

我对这种模式的机制的假设是：

1. The interface returned needs to be small so that there can be multiple implementations.
2. Hold off on returning an interface until you have multiple types  in your package implementing only the interface. Multiple types with the same behavior signature gives confidence that you’ve the right  abstraction.

1.返回的接口要小，这样才能有多个实现。
2. 推迟返回接口，直到您的包中有多个类型仅实现该接口。具有相同行为签名的多个类型让您确信您拥有正确的抽象。

#### **Consider creating a separate interfaces-only package for namespacing and standardization**

#### **考虑为命名空间和标准化创建一个单独的仅接口包**

This is **not** an official guideline from the Go team,  it’s just an observation as having packages that contain only interfaces is a common pattern in the standard library.

这 ** 不是 ** Go 团队的官方指南，这只是一种观察，因为仅包含接口的包是标准库中的常见模式。

An example is the [`hash.Hash`](https://golang.org/pkg/hash/) interface that's implemented by the packages under the subdirectories of `hash/` such as [`hash/crc32` ](https://golang.org/pkg/hash/crc32/) and [`hash/adler32`](https://golang.org/pkg/hash/adler32/). The hash package only exposes interfaces.

一个例子是 [`hash.Hash`](https://golang.org/pkg/hash/) 接口，它是由 `hash/` 子目录下的包实现的，例如 [`hash/crc32` ](https://golang.org/pkg/hash/crc32/) 和 [`hash/adler32`](https://golang.org/pkg/hash/adler32/)。 hash 包只暴露接口。

```
package hash

type Hash interface {
  ...
}

type Hash32 interface {
    Hash
    Sum32() uint32
}

type Hash64 interface {
    Hash
    Sum64() uint64
}
```

I suspect the benefits of moving interfaces to a separate package instead of exposing them in the subdirectories are two-fold:

我怀疑将接口移动到单独的包而不是将它们暴露在子目录中的好处有两个：

1. A better namespace for the interfaces. `hash.Hash` is easier to understand than `adler32.Hash`.
2. Standardizing how to implement a functionality. A separate package with only interfaces hints that hash functions should have the methods  required by the `hash.Hash` interface.

1. 更好的接口命名空间。 `hash.Hash` 比 `adler32.Hash` 更容易理解。
2. 标准化如何实现功能。一个只有接口的单独包提示散列函数应该具有`hash.Hash` 接口所需的方法。

Another package with only interfaces is [`encoding`](https://golang.org/pkg/encoding/).

另一个只有接口的包是 [`encoding`](https://golang.org/pkg/encoding/)。

```
package encoding

type BinaryMarshaler interface {
    MarshalBinary() (data []byte, err error)
}

type BinaryUnmarshaler interface {
    UnmarshalBinary(data []byte) error
}

type TextMarshaler interface {
    MarshalText() (text []byte, err error)
}

type TextUnmarshaler interface {
    UnmarshalText(text []byte) error
}
```

There are many structs in the std lib that implement the `encoding` interfaces. However, unlike `hash` which is consumed by the packages under `crypto/`, there are no functions in the std lib that accept or return an `encoding` interface.

标准库中有许多结构体实现了 `encoding` 接口。然而，与由 `crypto/` 下的包使用的 `hash` 不同，std lib 中没有接受或返回 `encoding` 接口的函数。

So why is it exposed?

那为什么会暴露呢？

I believe it’s because they want to hint to developers a standard  method signature for (un)marshalling a binary into an object. Existing  packages that test if a value implements the `encoding.BinaryMarshaler` interface won’t need to change their implementation if a new struct implements the interface.

我相信这是因为他们想向开发人员提示一个标准的方法签名，用于（取消）将二进制文件编组到对象中。如果新结构实现了该接口，则测试值是否实现了 `encoding.BinaryMarshaler` 接口的现有包不需要更改它们的实现。

```
if m, ok := v.(encoding.BinaryMarshaler);ok {
    return m.MarshalBinary()
}
```

It's worth noting that this pattern is not followed with the `Resetter` interface in [compress/zlib](https://golang.org/pkg/compress/zlib/#Resetter) and [compress/flate](https://golang.org/pkg/compress/flate/#Resetter) packages as it's duplicated in both packages. However, this appears to be a point of discussion even with Go maintainers ([see CR comment#27](https://codereview.appspot.com/97140043#msg27)).

值得注意的是，[compress/zlib](https://golang.org/pkg/compress/zlib/#Resetter)和[compress/flate](https://golang.org/pkg/compress/zlib/#Resetter)中的`Resetter`接口没有遵循这种模式golang.org/pkg/compress/flate/#Resetter) 包，因为它在两个包中都是重复的。然而，这似乎是 Go 维护者的一个讨论点（[参见 CR 评论#27](https://codereview.appspot.com/97140043#msg27))。

#### **Finally, private interfaces don’t have to deal with these considerations as they’re not exposed.**

#### **最后，私有接口不必处理这些注意事项，因为它们没有公开。**

We can have larger interfaces such as [`gobType`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/encoding/gob/type.go#L167-L173) from the `encoding/gob` package without worrying about its contents. Interfaces can be duplicated across packages, such as the `timeout` interface that exists in both the [`os`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/os/error.go#L35-L37) and [`net`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/net/net.go#L494-L496) packages, without thinking about placing them in a separate location.

我们可以有更大的接口，例如来自 `encoding/gob包而不必担心其内容。接口可以跨包复制，例如在 [`os`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/os/error.go#L35-L37) 和 [`net`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/net/net.go#L494-L496) 包，而不考虑将它们放在单独的位置。

#### **Takeaways**

#### **外卖**

Defer, defer, and defer writing an interface to when you have a better understanding of the abstraction needed.

推迟，推迟，再推迟编写接口，直到您对所需的抽象有了更好的理解。

A good signal as a producer is when you have multiple types that  implements the same method signatures. Then you can refactor and return  an interface. As a consumer keep your interfaces tiny so that multiple  types can implement it.

作为生产者的一个好信号是当您有多种类型实现相同的方法签名时。然后你可以重构并返回一个接口。作为消费者，保持你的接口很小，以便多种类型可以实现它。

*Thanks to Nick Fischer for reviewing early drafts of this post!*

*感谢 Nick Fischer 审阅这篇文章的早期草稿！*

              Last modified: Dec 30, 2019 

最后修改时间：2019 年 12 月 30 日

