# [Exposing interfaces in Go](https://www.efekarakus.com/golang/2019/12/29/working-with-interfaces-in-go.html)    

​      Published: Dec 29, 2019       

[Interfaces](https://golang.org/doc/effective_go.html#interfaces_and_types) are my favorite feature in Go. An interface type represents a set of  methods. Unlike most other languages, you don’t have to explicitly  declare that a type *implements* an interface. A struct `S` implements the interface `I` *implicitly* if `S` defines the methods that `I` requires.

Writing good interfaces is difficult. It’s easy to [pollute](https://rakyll.org/interface-pollution/) the “API” of a package by exposing broad or unnecessary interfaces. In  this article, we’ll explain the reasoning behind existing guidelines for interfaces and supplement them with examples from the standard library.

#### **[“The bigger the interface, the weaker the abstraction”](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)**

It’s unlikely that you’ll be able to find multiple types that can  implement a large interface. For that reason, ”interfaces with only one  or two methods are common in Go code”. Instead of declaring large *public* interfaces, consider depending on or returning an explicit type.

The [`io.Reader`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) and [`io.Writer`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L77-L92) interfaces are the usual examples for powerful interfaces.

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

After grepping the std lib, I found 81 structs across 30 packages that implement an `io.Reader`, and 99 methods or functions that consume it across 39 packages.

#### **[“Go interfaces \*generally\* belong in the package that uses values of the interface type, not the package that implements those values”](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)**

By defining the interface in the package that actually uses it, we  can let the client define the abstraction rather than the provider  dictating the abstraction to all of its clients.

An example is the [`io.Copy`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/io/io.go#L363) function. It accepts both the `Writer` and `Reader` interfaces as arguments defined in the same package.

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

Another example from a different package is the[`color.Color`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L10-L19) interface. The  [`Index`](https://github.com/golang/go/blob/2d6f8cc2cdd5993eb8dc80655735a38ef067af6e/src/image/color/color.go#L292-L308) method of the  `color.Palette` type depends on it, allowing it to accept any struct that implements the `Color` interface.

```
func (p Palette) Index(c Color) int
```

#### **[“If  a type exists only to implement an interface and will never have  exported methods beyond that interface, there is no need to export the  type itself”](https://golang.org/doc/effective_go.html#generality)**

The previous guideline from the [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces) mentions that:

> The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations  without requiring extensive refactoring.

Complementing it with the statement from [EffectiveGo](https://golang.org/doc/effective_go.html#generality), we can see the full picture where it’s okay to define an interface in the producer package:

> If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the  type itself.”

An example is the [`rand.Source`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L23-L28) interface which is returned by [`rand.NewSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L41-L48). The underlying struct [`rngSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rng.go#L180) within the constructor only exports the methods needed for the `Source` and `Source64` interfaces so the type itself is not exposed.

The `rand` package has two other types that implement the interface: [`lockedSource`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L382), and  [`Rand`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/math/rand/rand.go#L51) (the latter is exposed because it has other public methods).

**What’s the benefit of returning an interface over a concrete type anyway?**

Returning an interface allows you to have functions that can return multiple concrete types. For example, the [`aes.NewCipher`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher.go#L32) constructor returns a [`cipher.Block`](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/cipher/cipher.go#L15) interface. If you look [within the constructor](https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/crypto/aes/cipher_asm.go#L33-L54), you can see that two different structs are returned.

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

**In my short experience …**

… this pattern is a lot more difficult to execute than the previous  ones. In the early phases of development the needs of a client package  evolves quickly. Hence, it results in modifying the producer package. If the returned type is an interface, it  slowly becomes too big to the  point where returning a concrete type would make more sense.

My hypothesis for the mechanics of this pattern is:

1. The interface returned needs to be small so that there can be multiple implementations.
2. Hold off on returning an interface until you have multiple types  in your package implementing only the interface. Multiple types with the same behavior signature gives confidence that you’ve the right  abstraction.

#### **Consider creating a separate interfaces-only package for namespacing and standardization**

This is **not** an official guideline from the Go team,  it’s just an observation as having packages that contain only interfaces is a common pattern in the standard library.

An example is the [`hash.Hash`](https://golang.org/pkg/hash/) interface that’s implemented by the packages under the subdirectories of `hash/` such as [`hash/crc32` ](https://golang.org/pkg/hash/crc32/) and [`hash/adler32`](https://golang.org/pkg/hash/adler32/). The hash package only exposes interfaces.

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

1. A better namespace for the interfaces. `hash.Hash` is easier to understand than `adler32.Hash`.
2. Standardizing how to implement a functionality. A separate package with only interfaces hints that hash functions should have the methods  required by the `hash.Hash` interface.

Another package with only interfaces is [`encoding`](https://golang.org/pkg/encoding/).

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

So why is it exposed?

I believe it’s because they want to hint to developers a standard  method signature for (un)marshalling a binary into an object. Existing  packages that test if a value implements the `encoding.BinaryMarshaler` interface won’t need to change their implementation if a new struct implements the interface.

```
if m, ok := v.(encoding.BinaryMarshaler); ok {
    return m.MarshalBinary()
}
```

It’s worth noting that this pattern is not followed with the `Resetter` interface in [compress/zlib](https://golang.org/pkg/compress/zlib/#Resetter) and [compress/flate](https://golang.org/pkg/compress/flate/#Resetter) packages as it’s duplicated in both packages. However, this appears to be a point of discussion even with Go maintainers ([see CR comment#27](https://codereview.appspot.com/97140043#msg27)).

#### **Finally, private interfaces don’t have to deal with these considerations as they’re not exposed.**

We can have larger interfaces such as [`gobType`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/encoding/gob/type.go#L167-L173) from the `encoding/gob` package without worrying about its contents. Interfaces can be duplicated across packages, such as the `timeout` interface that exists in both the [`os`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/os/error.go#L35-L37) and [`net`](https://github.com/golang/go/blob/c170b14c2c1cfb2fd853a37add92a82fd6eb4318/src/net/net.go#L494-L496) packages, without thinking about placing them in a separate location.

#### **Takeaways**

Defer, defer, and defer writing an interface to when you have a better understanding of the abstraction needed.

A good signal as a producer is when you have multiple types that  implements the same method signatures. Then you can refactor and return  an interface. As a consumer keep your interfaces tiny so that multiple  types can implement it.

*Thanks to Nick Fischer for reviewing early drafts of this post!*

​              Last modified: Dec 30, 2019          
