# Maps



**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/maps)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/maps)**

In [arrays & slices](arrays-and-slices.md), you saw how to store values in order. Now, we will look at a way to store items by a `key` and look them up quickly.

在 [arrays & slices](arrays-and-slices.md) 中，您看到了如何按顺序存储值。现在，我们将研究一种通过“键”存储项目并快速查找它们的方法。

Maps allow you to store items in a manner similar to a dictionary. You can think of the `key` as the word and the `value` as the definition. And what better way is there to learn about Maps than to build our own dictionary?

 Map 允许您以类似于字典的方式存储项目。您可以将“键”视为单词，将“值”视为定义。还有什么比构建我们自己的字典更好的方式来了解 Map ？

First, assuming we already have some words with their definitions in the dictionary, if we search for a word, it should return the definition of it.

首先，假设我们在字典中已经有一些词及其定义，如果我们搜索一个词，它应该返回它的定义。

## Write the test first

## 先写测试

In `dictionary_test.go`

```go
package main

import "testing"

func TestSearch(t *testing.T) {
    dictionary := map[string]string{"test": "this is just a test"}

    got := Search(dictionary, "test")
    want := "this is just a test"

    if got != want {
        t.Errorf("got %q want %q given, %q", got, want, "test")
    }
}
```

Declaring a Map is somewhat similar to an array. Except, it starts with the `map` keyword and requires two types. The first is the key type, which is written inside the `[]`. The second is the value type, which goes right after the `[]`.

声明 Map 有点类似于数组。除此之外，它以`map` 关键字开头并且需要两种类型。第一个是键类型，写在`[]`里面。第二个是值类型，它紧跟在 `[]` 之后。

The key type is special. It can only be a comparable type because without the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value. Comparable types are explained in depth in the [language spec](https://golang.org/ref/spec#Comparison_operators).

密钥类型是特殊的。它只能是可比较的类型，因为无法判断 2 个键是否相等，我们无法确保获得正确的值。 [语言规范](https://golang.org/ref/spec#Comparison_operators) 中对可比类型进行了深入解释。

The value type, on the other hand, can be any type you want. It can even be another map.

另一方面，值类型可以是您想要的任何类型。它甚至可以是另一个 Map 。

Everything else in this test should be familiar.

此测试中的其他所有内容都应该很熟悉。

## Try to run the test

## 尝试运行测试

By running `go test` the compiler will fail with `./dictionary_test.go:8:9: undefined: Search`.

通过运行 `go test`，编译器将失败并返回 `./dictionary_test.go:8:9: undefined: Search`。

## Write the minimal amount of code for the test to run and check the output

## 编写最少的代码来运行测试并检查输出

In `dictionary.go`

```go
package main

func Search(dictionary map[string]string, word string) string {
    return ""
}
```

Your test should now fail with a *clear error message*

您的测试现在应该会失败并显示 *clear 错误消息 *

`dictionary_test.go:12: got '' want 'this is just a test' given, 'test'`.

`dictionary_test.go:12：得到''想要'这只是一个测试'，'test''。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func Search(dictionary map[string]string, word string) string {
    return dictionary[word]
}
```

Getting a value out of a Map is the same as getting a value out of Array `map[key]`.

从 Map 中获取值与从数组 `map[key]` 中获取值相同。

## Refactor

## 重构

```go
func TestSearch(t *testing.T) {
    dictionary := map[string]string{"test": "this is just a test"}

    got := Search(dictionary, "test")
    want := "this is just a test"

    assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

I decided to create an `assertStrings` helper to make the implementation more general.

我决定创建一个 `assertStrings` 帮助器来使实现更加通用。

### Using a custom type

### 使用自定义类型

We can improve our dictionary's usage by creating a new type around map and making `Search` a method.

我们可以通过围绕 Map 创建一个新类型并使“搜索”成为一种方法来改进字典的使用。

In `dictionary_test.go`:

```go
func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    got := dictionary.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
}
```

We started using the `Dictionary` type, which we have not defined yet. Then called `Search` on the `Dictionary` instance.

我们开始使用尚未定义的`Dictionary` 类型。然后在`Dictionary` 实例上调用`Search`。

We did not need to change `assertStrings`.

我们不需要更改 `assertStrings`。

In `dictionary.go`:

```go
type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
    return d[word]
}
```

Here we created a `Dictionary` type which acts as a thin wrapper around `map`. With the custom type defined, we can create the `Search` method.

在这里，我们创建了一个 `Dictionary` 类型，它充当 `map` 的薄包装器。定义自定义类型后，我们可以创建`Search` 方法。

## Write the test first

## 先写测试

The basic search was very easy to implement, but what will happen if we supply a word that's not in our dictionary?

基本搜索很容易实现，但是如果我们提供字典中没有的单词会发生什么？

We actually get nothing back. This is good because the program can continue to run, but there is a better approach. The function can report that the word is not in the dictionary. This way, the user isn't left wondering if the word doesn't exist or if there is just no definition (this might not seem very useful for a dictionary. However, it's a scenario that could be key in other usecases).

我们实际上一无所获。这很好，因为程序可以继续运行，但还有更好的方法。该函数可以报告该词不在字典中。这样，用户就不会想知道该词是否不存在或是否没有定义（这对于字典来说似乎不是很有用。但是，这是一个可能在其他用例中很关键的场景）。

```go
func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

        if err == nil {
            t.Fatal("expected to get an error.")
        }

        assertStrings(t, err.Error(), want)
    })
}
```

The way to handle this scenario in Go is to return a second argument which is an `Error` type.

在 Go 中处理这种情况的方法是返回第二个参数，它是一个 `Error` 类型。

`Error`s can be converted to a string with the `.Error()` method, which we do when passing it to the assertion. We are also protecting `assertStrings` with `if` to ensure we don't call `.Error()` on `nil`.

可以使用 `.Error()` 方法将 `Error` 转换为字符串，我们在将其传递给断言时会这样做。我们还使用 `if` 保护 `assertStrings`，以确保我们不会在 `nil` 上调用 `.Error()`。

## Try and run the test

## 尝试并运行测试

This does not compile

这不编译

```
./dictionary_test.go:18:10: assignment mismatch: 2 variables but 1 values
```

## Write the minimal amount of code for the test to run and check the output

## 编写最少的代码来运行测试并检查输出

```go
func (d Dictionary) Search(word string) (string, error) {
    return d[word], nil
}
```

Your test should now fail with a much clearer error message.

您的测试现在应该会失败并显示更清晰的错误消息。

`dictionary_test.go:22: expected to get an error.`

`dictionary_test.go:22：预计会出错。`

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", errors.New("could not find the word you were looking for")
    }

    return definition, nil
}
```

In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values. The second value is a boolean which indicates if the key was found successfully.

为了完成这个过程，我们使用了 Map 查找的一个有趣属性。它可以返回 2 个值。第二个值是一个布尔值，表示是否成功找到了密钥。

This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.

这个属性允许我们区分一个不存在的词和一个没有定义的词。

## Refactor

## 重构

```go
var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }

    return definition, nil
}
```

We can get rid of the magic error in our `Search` function by extracting it into a variable. This will also allow us to have a better test.

我们可以通过将它提取到一个变量中来摆脱我们的“搜索”函数中的魔法错误。这也能让我们进行更好的测试。

```go
t.Run("unknown word", func(t *testing.T) {
    _, got := dictionary.Search("unknown")

    assertError(t, got, ErrNotFound)
})

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}
```

By creating a new helper we were able to simplify our test, and start using our `ErrNotFound` variable so our test doesn't fail if we change the error text in the future.

通过创建一个新的帮助器，我们能够简化我们的测试，并开始使用我们的 `ErrNotFound` 变量，这样如果我们将来更改错误文本，我们的测试就不会失败。

## Write the test first

## 先写测试

We have a great way to search the dictionary. However, we have no way to add new words to our dictionary.

我们有一个很好的方法来搜索字典。但是，我们无法在字典中添加新词。

```go
func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    dictionary.Add("test", "this is just a test")

    want := "this is just a test"
    got, err := dictionary.Search("test")
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

In this test, we are utilizing our `Search` function to make the validation of the dictionary a little easier.

在这个测试中，我们利用我们的“搜索”功能使字典的验证更容易一些。

## Write the minimal amount of code for the test to run and check output

## 为测试运行和检查输出编写最少的代码

In `dictionary.go`

在`dictionary.go`

```go
func (d Dictionary) Add(word, definition string) {
}
```

Your test should now fail

你的测试现在应该失败

```
dictionary_test.go:31: should find added word: could not find the word you were looking for
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}
```

Adding to a map is also similar to an array. You just need to specify a key and set it equal to a value.

添加到 Map 也类似于数组。您只需要指定一个键并将其设置为等于一个值。

### Pointers, copies, et al

### 指针、副本等

An interesting property of maps is that you can modify them without passing as an address to it (e.g `&myMap`)

 Map 的一个有趣特性是您可以修改它们而无需将其作为地址传递给它（例如`&myMap`）

This may make them _feel_ like a "reference type", [but as Dave Cheney describes](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) they are not.

这可能让他们_感觉_像一个“引用类型”，[但正如 Dave Cheney 描述的那样](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-这是什么)他们不是。

> A map value is a pointer to a runtime.hmap structure.

> 映射值是指向 runtime.hmap 结构的指针。

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

因此，当您将映射传递给函数/方法时，您确实在复制它，但只是指针部分，而不是包含数据的底层数据结构。

A gotcha with maps is that they can be a `nil` value. A `nil` map behaves like an empty map when reading, but attempts to write to a `nil` map will cause a runtime panic. You can read more about maps [here](https://blog.golang.org/go-maps-in-action).

 Map 的一个问题是它们可以是一个 `nil` 值。 `nil` 映射在读取时表现得像一个空映射，但尝试写入一个 `nil` 映射会导致运行时恐慌。您可以在 [此处](https://blog.golang.org/go-maps-in-action) 阅读有关 Map 的更多信息。

Therefore, you should never initialize an empty map variable:

因此，你永远不应该初始化一个空的 map 变量：

```go
var m map[string]string
```

Instead, you can initialize an empty map like we were doing above, or use the `make` keyword to create a map for you:

相反，您可以像我们上面所做的那样初始化一个空 Map ，或者使用 `make` 关键字为您创建一个 Map ：

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Both approaches create an empty `hash map` and point `dictionary` at it. Which ensures that you will never get a runtime panic.

这两种方法都会创建一个空的“哈希图”并将“字典”指向它。这确保您永远不会遇到运行时恐慌。

## Refactor 

## 重构

There isn't much to refactor in our implementation but the test could use a little simplification.

在我们的实现中没有太多要重构的，但测试可以使用一些简化。

```go
func TestAdd(t *testing.T) {
    dictionary := Dictionary{}
    word := "test"
    definition := "this is just a test"

    dictionary.Add(word, definition)

    assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got %q want %q", got, definition)
    }
}
```

We made variables for word and definition, and moved the definition assertion into its own helper function.

我们为单词和定义创建了变量，并将定义断言移动到它自己的辅助函数中。

Our `Add` is looking good. Except, we didn't consider what happens when the value we are trying to add already exists!

我们的“添加”看起来不错。除了，我们没有考虑当我们尝试添加的值已经存在时会发生什么！

Map will not throw an error if the value already exists. Instead, they will go ahead and overwrite the value with the newly provided value. This can be convenient in practice, but makes our function name less than accurate. `Add` should not modify existing values. It should only add new words to our dictionary.

如果值已经存在，Map 不会抛出错误。相反，他们将继续使用新提供的值覆盖该值。这在实践中很方便，但会使我们的函数名称不够准确。 `Add` 不应修改现有值。它应该只在我们的字典中添加新词。

## Write the test first

## 先写测试

```go
func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        definition := "this is just a test"

        err := dictionary.Add(word, definition)

        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
    })

    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dictionary := Dictionary{word: definition}
        err := dictionary.Add(word, "new test")

        assertError(t, err, ErrWordExists)
        assertDefinition(t, dictionary, word, definition)
    })
}
...
func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

For this test, we modified `Add` to return an error, which we are validating against a new error variable, `ErrWordExists`. We also modified the previous test to check for a `nil` error, as well as the `assertError` function.

对于这个测试，我们修改了 `Add` 以返回一个错误，我们正在针对一个新的错误变量 `ErrWordExists` 进行验证。我们还修改了之前的测试以检查 `nil` 错误，以及 `assertError` 函数。

## Try to run test

## 尝试运行测试

The compiler will fail because we are not returning a value for `Add`.

编译器将失败，因为我们没有返回 `Add` 的值。

```
./dictionary_test.go:30:13: dictionary.Add(word, definition) used as value
./dictionary_test.go:41:13: dictionary.Add(word, "new test") used as value
```

## Write the minimal amount of code for the test to run and check the output

## 编写最少的代码来运行测试并检查输出

In `dictionary.go`



```go
var (
    ErrNotFound   = errors.New("could not find the word you were looking for")
    ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Add(word, definition string) error {
    d[word] = definition
    return nil
}
```

Now we get two more errors. We are still modifying the value, and returning a `nil` error.

现在我们又得到两个错误。我们仍在修改该值，并返回一个 `nil` 错误。

```
dictionary_test.go:43: got error '%!q(<nil>)' want 'cannot add word because it already exists'
dictionary_test.go:44: got 'new test' want 'this is just a test'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (d Dictionary) Add(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        d[word] = definition
    case nil:
        return ErrWordExists
    default:
        return err
    }

    return nil
}
```

Here we are using a `switch` statement to match on the error. Having a `switch` like this provides an extra safety net, in case `Search` returns an error other than `ErrNotFound`.

在这里，我们使用 `switch` 语句来匹配错误。拥有像这样的 `switch` 提供了一个额外的安全网，以防 `Search` 返回除 `ErrNotFound` 以外的错误。

## Refactor

## 重构

We don't have too much to refactor, but as our error usage grows we can make a few modifications.

我们没有太多要重构的内容，但是随着错误使用量的增加，我们可以进行一些修改。

```go
const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}
```

We made the errors constant; this required us to create our own `DictionaryErr` type which implements the `error` interface. You can read more about the details in [this excellent article by Dave Cheney](https://dave.cheney.net/2016/04/07/constant-errors). Simply put, it makes the errors more reusable and immutable.

我们使错误保持不变；这需要我们创建我们自己的 DictionaryErr 类型，它实现了 `error` 接口。您可以在 [Dave Cheney 撰写的这篇优秀文章](https://dave.cheney.net/2016/04/07/constant-errors) 中阅读有关详细信息的更多信息。简而言之，它使错误更具可重用性和不可变性。

Next, let's create a function to `Update` the definition of a word.

接下来，让我们创建一个函数来“更新”一个词的定义。

## Write the test first

## 先写测试

```go
func TestUpdate(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{word: definition}
    newDefinition := "new definition"

    dictionary.Update(word, newDefinition)

    assertDefinition(t, dictionary, word, newDefinition)
}
```

`Update` is very closely related to `Add` and will be our next implementation.

`Update` 与 `Add` 密切相关，将是我们的下一个实现。

## Try and run the test

## 尝试并运行测试

```
./dictionary_test.go:53:2: dictionary.Update undefined (type Dictionary has no field or method Update)
```

## Write minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

We already know how to deal with an error like this. We need to define our function.

我们已经知道如何处理这样的错误。我们需要定义我们的函数。

```go
func (d Dictionary) Update(word, definition string) {}
```

With that in place, we are able to see that we need to change the definition of the word.

有了这个，我们就可以看到我们需要改变这个词的定义。

```
dictionary_test.go:55: got 'this is just a test' want 'new definition'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We already saw how to do this when we fixed the issue with `Add`. So let's implement something really similar to `Add`.

当我们使用`Add`解决问题时，我们已经看到了如何做到这一点。所以让我们实现一些与`Add`非常相似的东西。

```go
func (d Dictionary) Update(word, definition string) {
    d[word] = definition
}
```

There is no refactoring we need to do on this since it was a simple change. However, we now have the same issue as with `Add`. If we pass in a new word, `Update` will add it to the dictionary.

我们不需要对此进行重构，因为这是一个简单的更改。但是，我们现在遇到了与“添加”相同的问题。如果我们传入一个新单词，`Update` 会将它添加到字典中。

## Write the test first

## 先写测试

```go
t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    newDefinition := "new definition"
    dictionary := Dictionary{word: definition}

    err := dictionary.Update(word, newDefinition)

    assertError(t, err, nil)
    assertDefinition(t, dictionary, word, newDefinition)
})

t.Run("new word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{}

    err := dictionary.Update(word, definition)

    assertError(t, err, ErrWordDoesNotExist)
})
```

We added yet another error type for when the word does not exist. We also modified `Update` to return an `error` value.

当单词不存在时，我们添加了另一种错误类型。我们还修改了 `Update` 以返回一个 `error` 值。

## Try and run the test

## 尝试并运行测试

```
./dictionary_test.go:53:16: dictionary.Update(word, newDefinition) used as value
./dictionary_test.go:64:16: dictionary.Update(word, definition) used as value
./dictionary_test.go:66:23: undefined: ErrWordDoesNotExist
```

We get 3 errors this time, but we know how to deal with these.

这次我们遇到了 3 个错误，但我们知道如何处理这些错误。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

```go
const (
    ErrNotFound         = DictionaryErr("could not find the word you were looking for")
    ErrWordExists       = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (d Dictionary) Update(word, definition string) error {
    d[word] = definition
    return nil
}
```

We added our own error type and are returning a `nil` error.

我们添加了自己的错误类型并返回了一个 `nil` 错误。

With these changes, we now get a very clear error:

通过这些更改，我们现在得到一个非常明显的错误：

```
dictionary_test.go:66: got error '%!q(<nil>)' want 'cannot update word because it does not exist'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (d Dictionary) Update(word, definition string) error {
    _, err := d.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[word] = definition
    default:
        return err
    }

    return nil
}
```

This function looks almost identical to `Add` except we switched when we update the `dictionary` and when we return an error.

这个函数看起来几乎和 `Add` 一样，除了我们在更新 `dictionary` 和返回错误时切换。

### Note on declaring a new error for Update

### 关于为更新声明新错误的注意事项

We could reuse `ErrNotFound` and not add a new error. However, it is often better to have a precise error for when an update fails.

我们可以重用 `ErrNotFound` 而不会添加新的错误。但是，当更新失败时，最好有一个精确的错误。

Having specific errors gives you more information about what went wrong. Here is an example in a web app:

具有特定错误可为您提供有关出错的更多信息。以下是 Web 应用程序中的示例：

> You can redirect the user when `ErrNotFound` is encountered, but display an error message when `ErrWordDoesNotExist` is encountered.

> 您可以在遇到 `ErrNotFound` 时重定向用户，但在遇到 `ErrWordDoesNotExist` 时显示错误消息。

Next, let's create a function to `Delete` a word in the dictionary.

接下来，让我们创建一个函数来“删除”字典中的单词。

## Write the test first

## 先写测试

```go
func TestDelete(t *testing.T) {
    word := "test"
    dictionary := Dictionary{word: "test definition"}

    dictionary.Delete(word)

    _, err := dictionary.Search(word)
    if err != ErrNotFound {
        t.Errorf("Expected %q to be deleted", word)
    }
}
```

Our test creates a `Dictionary` with a word and then checks if the word has been removed.

我们的测试创建了一个带有单词的“词典”，然后检查该单词是否已被删除。

## Try to run the test

## 尝试运行测试

By running `go test` we get:

通过运行`go test`，我们得到：

```
./dictionary_test.go:74:6: dictionary.Delete undefined (type Dictionary has no field or method Delete)
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

```go
func (d Dictionary) Delete(word string) {

}
```

After we add this, the test tells us we are not deleting the word.

添加这个之后，测试告诉我们我们没有删除这个词。

```
dictionary_test.go:78: Expected 'test' to be deleted
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func (d Dictionary) Delete(word string) {
    delete(d, word)
}
```

Go has a built-in function `delete` that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.

Go 有一个内置函数 `delete` 可以处理 Map 。它需要两个参数。第一个是 Map ，第二个是要删除的键。

The `delete` function returns nothing, and we based our `Delete` method on the same notion. Since deleting a value that's not there has no effect, unlike our `Update` and `Add` methods, we don't need to complicate the API with errors.

`delete` 函数不返回任何内容，我们的 `Delete` 方法基于相同的概念。由于删除一个不存在的值没有任何效果，不像我们的 `Update` 和 `Add` 方法，我们不需要用错误使 API 复杂化。

## Wrapping up

##  总结

In this section, we covered a lot. We made a full CRUD (Create, Read, Update and Delete) API for our dictionary. Throughout the process we learned how to:

在本节中，我们涵盖了很多内容。我们为我们的字典制作了一个完整的 CRUD（创建、读取、更新和删除）API。在整个过程中，我们学会了如何：

* Create maps
* Search for items in maps
* Add new items to maps
* Update items in maps
* Delete items from a map
* Learned more about errors
   * How to create errors that are constants
   * Writing error wrappers 

* 创建
* 在 Map 中搜索项目
* 向 Map 添加新项目
* 更新 Map 中的项目
* 从 Map 中删除项目
* 了解有关错误的更多信息
  * 如何创建常量错误
  * 编写错误包装器

