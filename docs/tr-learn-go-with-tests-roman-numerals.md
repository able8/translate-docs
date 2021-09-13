# Roman Numerals

#  罗马数字

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/roman-numerals)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/roman-numerals)**

Some companies will ask you to do the [Roman Numeral Kata](http://codingdojo.org/kata/RomanNumerals/) as part of the interview process. This chapter will show how you can tackle it with TDD.

作为面试过程的一部分，一些公司会要求您做[罗马数字卡塔](http://codingdojo.org/kata/RomanNumerals/)。本章将展示如何使用 TDD 解决它。

We are going to write a function which converts an [Arabic number](https://en.wikipedia.org/wiki/Arabic_numerals) (numbers 0 to 9) to a Roman Numeral.

我们将编写一个函数，将 [阿拉伯数字](https://en.wikipedia.org/wiki/Arabic_numerals)（数字 0 到 9)转换为罗马数字。

If you haven't heard of [Roman Numerals](https://en.wikipedia.org/wiki/Roman_numerals) they are how the Romans wrote down numbers.

如果您还没有听说过 [罗马数字](https://en.wikipedia.org/wiki/Roman_numerals)，它们就是罗马人记下数字的方式。

You build them by sticking symbols together and those symbols represent numbers

您通过将符号粘在一起来构建它们，这些符号代表数字

So `I` is "one". `III` is three.

所以`I`是“一”。 `III`是三。

Seems easy but there's a few interesting rules. `V` means five, but `IV` is 4 (not `IIII`).

看起来很简单，但有一些有趣的规则。 `V` 表示 5，但 `IV` 是 4（不是 `IIII`）。

`MCMLXXXIV` is 1984. That looks complicated and it's hard to imagine how we can write code to figure this out right from the start.

`MCMLXXXIV` 是 1984 年。这看起来很复杂，很难想象我们如何从一开始就编写代码来解决这个问题。

As this book stresses, a key skill for software developers is to try and identify "thin vertical slices" of _useful_ functionality and then **iterating**. The TDD workflow helps facilitate iterative development.

正如本书所强调的，软件开发人员的一项关键技能是尝试识别 _有用_ 功能的“垂直薄片”，然后**迭代**。 TDD 工作流有助于促进迭代开发。

So rather than 1984, let's start with 1.

因此，让我们从 1 开始，而不是 1984 年。

## Write the test first

## 先写测试

```go
func TestRomanNumerals(t *testing.T) {
    got := ConvertToRoman(1)
    want := "I"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

If you've got this far in the book this is hopefully feeling very boring and routine to you. That's a good thing.

如果你在书中读到了这些，希望你会觉得非常无聊和例行公事。这是好事。

## Try to run the test

## 尝试运行测试

`./numeral_test.go:6:9: undefined: ConvertToRoman`

Let the compiler guide the way

让编译器带路

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Create our function but don't make the test pass yet, always make sure the tests fails how you expect

创建我们的函数，但不要让测试通过，始终确保测试按预期失败

```go
func ConvertToRoman(arabic int) string {
    return ""
}
```

It should run now

它现在应该运行

```go
=== RUN   TestRomanNumerals
--- FAIL: TestRomanNumerals (0.00s)
    numeral_test.go:10: got '', want 'I'
FAIL
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func ConvertToRoman(arabic int) string {
    return "I"
}
```

## Refactor

## 重构

Not much to refactor yet.

还没有多少重构。

_I know_ it feels weird just to hard-code the result but with TDD we want to stay out of "red" for as long as possible. It may _feel_ like we haven't accomplished much but we've defined our API and got a test capturing one of our rules; even if the "real" code is pretty dumb.

_我知道_只是对结果进行硬编码感觉很奇怪，但是使用 TDD 我们希望尽可能长时间地远离“红色”。可能_感觉_我们没有完成太多工作，但我们已经定义了 API 并进行了测试以捕获我们的规则之一；即使“真正的”代码非常愚蠢。

Now use that uneasy feeling to write a new test to force us to write slightly less dumb code.

现在使用这种不安的感觉来编写一个新的测试来迫使我们编写稍微不那么愚蠢的代码。

## Write the test first

## 先写测试

We can use subtests to nicely group our tests

我们可以使用子测试很好地对我们的测试进行分组

```go
func TestRomanNumerals(t *testing.T) {
    t.Run("1 gets converted to I", func(t *testing.T) {
        got := ConvertToRoman(1)
        want := "I"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

    t.Run("2 gets converted to II", func(t *testing.T) {
        got := ConvertToRoman(2)
        want := "II"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRomanNumerals/2_gets_converted_to_II
    --- FAIL: TestRomanNumerals/2_gets_converted_to_II (0.00s)
        numeral_test.go:20: got 'I', want 'II'
```

Not much surprise there

没有太多惊喜

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func ConvertToRoman(arabic int) string {
    if arabic == 2 {
        return "II"
    }
    return "I"
}
```

Yup, it still feels like we're not actually tackling the problem. So we need to write more tests to drive us forward.

是的，仍然感觉我们并没有真正解决这个问题。所以我们需要编写更多的测试来推动我们前进。

## Refactor

## 重构

We have some repetition in our tests. When you're testing something which feels like it's a matter of "given input X, we expect Y" you should probably use table based tests.

我们的测试中有一些重复。当您测试的东西感觉像是“给定输入 X，我们期望 Y”的问题时，您可能应该使用基于表格的测试。

```go
func TestRomanNumerals(t *testing.T) {
    cases := []struct {
        Description string
        Arabic      int
        Want        string
    }{
        {"1 gets converted to I", 1, "I"},
        {"2 gets converted to II", 2, "II"},
    }

    for _, test := range cases {
        t.Run(test.Description, func(t *testing.T) {
            got := ConvertToRoman(test.Arabic)
            if got != test.Want {
                t.Errorf("got %q, want %q", got, test.Want)
            }
        })
    }
}
```

We can now easily add more cases without having to write any more test boilerplate.

我们现在可以轻松添加更多案例，而无需编写更多测试样板。

Let's push on and go for 3

让我们继续前进 3

## Write the test first

## 先写测试

Add the following to our cases

将以下内容添加到我们的案例中

```go
{"3 gets converted to III", 3, "III"},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRomanNumerals/3_gets_converted_to_III
    --- FAIL: TestRomanNumerals/3_gets_converted_to_III (0.00s)
        numeral_test.go:20: got 'I', want 'III'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func ConvertToRoman(arabic int) string {
    if arabic == 3 {
        return "III"
    }
    if arabic == 2 {
        return "II"
    }
    return "I"
}
```

## Refactor 

## 重构

OK so I'm starting to not enjoy these if statements and if you look at the code hard enough you can see that we're building a string of `I` based on the size of `arabic`.

好的，所以我开始不喜欢这些 if 语句了，如果你仔细查看代码，你会发现我们正在根据 `arabic` 的大小构建一个 `I` 字符串。

We "know" that for more complicated numbers we will be doing some kind of arithmetic and string concatenation.

我们“知道”对于更复杂的数字，我们将进行某种算术和字符串连接。

Let's try a refactor with these thoughts in mind, it _might not_ be suitable for the end solution but that's OK. We can always throw our code away and start afresh with the tests we have to guide us.

让我们考虑到这些想法尝试重构，它_可能不_适合最终解决方案，但没关系。我们总是可以扔掉我们的代码，重新开始我们必须指导我们的测试。

```go
func ConvertToRoman(arabic int) string {

    var result strings.Builder

    for i:=0;i<arabic;i++ {
        result.WriteString("I")
    }

    return result.String()
}
```

You may not have used [`strings.Builder`](https://golang.org/pkg/strings/#Builder) before

您之前可能没有使用过 [`strings.Builder`](https://golang.org/pkg/strings/#Builder)

> A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.

> Builder 用于使用 Write 方法有效地构建字符串。它最大限度地减少了内存复制。

Normally I wouldn't bother with such optimisations until I have an actual performance problem but the amount of code is not much larger than a "manual" appending on a string so we may as well use the faster approach.

通常，在遇到实际性能问题之前，我不会考虑进行此类优化，但代码量并不比附加在字符串上的“手动”大多少，因此我们不妨使用更快的方法。

The code looks better to me and describes the domain _as we know it right now_.

代码对我来说看起来更好，并描述了域_正如我们现在所知_。

### The Romans were into DRY too...

### 罗马人也很干燥......

Things start getting more complicated now. The Romans in their wisdom thought repeating characters would become hard to read and count. So a rule with Roman Numerals is you can't have the same character repeated more than 3 times in a row.

现在事情开始变得更加复杂。罗马人以他们的智慧认为重复的字符会变得难以阅读和计数。因此，罗马数字的规则是同一字符不能连续重复超过 3 次。

Instead you take the next highest symbol and then "subtract" by putting a symbol to the left of it. Not all symbols can be used as subtractors; only I (1), X (10) and C (100).

相反，您取下一个最高的符号，然后通过在其左侧放置一个符号来“减去”。并非所有符号都可以用作减法器；只有 I (1)、X (10) 和 C (100)。

For example `5` in Roman Numerals is `V`. To create 4 you do not do `IIII`, instead you do `IV`.

例如，罗马数字中的“5”是“V”。要创建 4，您不执行 `IIII`，而是执行 `IV`。

## Write the test first

## 先写测试

```go
{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRomanNumerals/4_gets_converted_to_IV_(cant_repeat_more_than_3_times)
    --- FAIL: TestRomanNumerals/4_gets_converted_to_IV_(cant_repeat_more_than_3_times) (0.00s)
        numeral_test.go:24: got 'IIII', want 'IV'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
func ConvertToRoman(arabic int) string {

    if arabic == 4 {
        return "IV"
    }

    var result strings.Builder

    for i:=0;i<arabic;i++ {
        result.WriteString("I")
    }

    return result.String()
}
```

## Refactor

## 重构

I don't "like" that we have broken our string building pattern and I want to carry on with it.

我不“喜欢”我们打破了我们的弦乐构建模式，我想继续下去。

```go
func ConvertToRoman(arabic int) string {

    var result strings.Builder

    for i := arabic;i > 0;i-- {
        if i == 4 {
            result.WriteString("IV")
            break
        }
        result.WriteString("I")
    }

    return result.String()
}
```

In order for 4 to "fit" with my current thinking I now count down from the Arabic number, adding symbols to our string as we progress. Not sure if this will work in the long run but let's see!

为了让 4 与我目前的想法“契合”，我现在从阿拉伯数字开始倒数，随着我们的进展向我们的字符串添加符号。不确定从长远来看这是否会奏效，但让我们看看！

Let's make 5 work

让我们做 5 个工作

## Write the test first

## 先写测试

```go
{"5 gets converted to V", 5, "V"},
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRomanNumerals/5_gets_converted_to_V
    --- FAIL: TestRomanNumerals/5_gets_converted_to_V (0.00s)
        numeral_test.go:25: got 'IIV', want 'V'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Just copy the approach we did for 4

只需复制我们为 4 所做的方法

```go
func ConvertToRoman(arabic int) string {

    var result strings.Builder

    for i := arabic;i > 0;i-- {
        if i == 5 {
            result.WriteString("V")
            break
        }
        if i == 4 {
            result.WriteString("IV")
            break
        }
        result.WriteString("I")
    }

    return result.String()
}
```

## Refactor

## 重构

Repetition in loops like this are usually a sign of an abstraction waiting to be called out. Short-circuiting loops can be an effective tool for readability but it could also be telling you something else.

像这样的循环中的重复通常是等待被调用的抽象的标志。短路循环可以是提高可读性的有效工具，但它也可以告诉您其他信息。

We are looping over our Arabic number and if we hit certain symbols we are calling `break` but what we are _really_ doing is subtracting over `i` in a ham-fisted manner.

我们正在循环我们的阿拉伯数字，如果我们遇到某些符号，我们将调用“break”，但我们_真正_所做的是以一种笨拙的方式减去“i”。

```go
func ConvertToRoman(arabic int) string {

    var result strings.Builder

    for arabic > 0 {
        switch {
        case arabic > 4:
            result.WriteString("V")
            arabic -= 5
        case arabic > 3:
            result.WriteString("IV")
            arabic -= 4
        default:
            result.WriteString("I")
            arabic--
        }
    }

    return result.String()
}

```

- Given the signals I'm reading from our code, driven from our tests of some very basic scenarios I can see that to build a Roman Numeral I need to subtract from `arabic` as I apply symbols
- The `for` loop no longer relies on an `i` and instead we will keep building our string until we have subtracted enough symbols away from `arabic`. 

- 鉴于我从我们的代码中读取的信号，从我们对一些非常基本场景的测试驱动，我可以看到要构建一个罗马数字，我需要在应用符号时从“阿拉伯语”中减去
- `for` 循环不再依赖于 `i`，而是我们将继续构建我们的字符串，直到我们从 `arabic` 中减去足够多的符号。

I'm pretty sure this approach will be valid for 6 (VI), 7 (VII) and 8 (VIII) too. Nonetheless add the cases in to our test suite and check (I won't include the code for brevity, check the github for samples if you're unsure).

我很确定这种方法也适用于 6 (VI)、7 (VII) 和 8 (VIII)。尽管如此，将案例添加到我们的测试套件中并进行检查（为简洁起见，我不会包含代码，如果您不确定，请查看 github 中的示例）。

9 follows the same rule as 4 in that we should subtract `I` from the representation of the following number. 10 is represented in Roman Numerals with `X`; so therefore 9 should be `IX`.

9 遵循与 4 相同的规则，因为我们应该从以下数字的表示中减去“I”。 10 在罗马数字中用‘X’表示；所以因此 9 应该是“IX”。

## Write the test first

## 先写测试

```go
{"9 gets converted to IX", 9, "IX"}
```

## Try to run the test

## 尝试运行测试

```
=== RUN   TestRomanNumerals/9_gets_converted_to_IX
    --- FAIL: TestRomanNumerals/9_gets_converted_to_IX (0.00s)
        numeral_test.go:29: got 'VIV', want 'IX'
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We should be able to adopt the same approach as before

我们应该能够采用与以前相同的方法

```go
case arabic > 8:
    result.WriteString("IX")
    arabic -= 9
```

## Refactor

## 重构

It _feels_ like the code is still telling us there's a refactor somewhere but it's not totally obvious to me, so let's keep going.

_感觉_代码仍在告诉我们某处有重构，但对我来说并不完全明显，所以让我们继续。

I'll skip the code for this too, but add to your test cases a test for `10` which should be `X` and make it pass before reading on.

我也会跳过这方面的代码，但是在你的测试用例中添加一个对“10”的测试，它应该是“X”，并在继续阅读之前让它通过。

Here are a few tests I added as I'm confident up to 39 our code should work

这是我添加的一些测试，因为我有信心我们的代码最多可以工作 39

```go
{"10 gets converted to X", 10, "X"},
{"14 gets converted to XIV", 14, "XIV"},
{"18 gets converted to XVIII", 18, "XVIII"},
{"20 gets converted to XX", 20, "XX"},
{"39 gets converted to XXXIX", 39, "XXXIX"},
```

If you've ever done OO programming, you'll know that you should view `switch` statements with a bit of suspicion. Usually you are capturing a concept or data inside some imperative code when in fact it could be captured in a class structure instead.

如果你曾经做过面向对象编程，你就会知道你应该带着一点怀疑看待 `switch` 语句。通常，您在一些命令式代码中捕获概念或数据，而实际上它可以在类结构中捕获。

Go isn't strictly OO but that doesn't mean we ignore the lessons OO offers entirely (as much as some would like to tell you).

Go 不是严格的面向对象，但这并不意味着我们完全忽略面向对象提供的课程（正如一些人想告诉你的那样）。

Our switch statement is describing some truths about Roman Numerals along with behaviour.

我们的 switch 语句描述了一些关于罗马数字和行为的真相。

We can refactor this by decoupling the data from the behaviour.

我们可以通过将数据与行为解耦来重构它。

```go
type RomanNumeral struct {
    Value  int
    Symbol string
}

var allRomanNumerals = []RomanNumeral {
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func ConvertToRoman(arabic int) string {

    var result strings.Builder

    for _, numeral := range allRomanNumerals {
        for arabic >= numeral.Value {
            result.WriteString(numeral.Symbol)
            arabic -= numeral.Value
        }
    }

    return result.String()
}
```

This feels much better. We've declared some rules around the numerals as data rather than hidden in an algorithm and we can see how we just work through the Arabic number, trying to add symbols to our result if they fit.

这感觉好多了。我们已经将围绕数字的一些规则声明为数据而不是隐藏在算法中，我们可以看到我们如何只处理阿拉伯数字，如果它们合适，则尝试将符号添加到我们的结果中。

Does this abstraction work for bigger numbers? Extend the test suite so it works for the Roman number for 50 which is `L`.

这种抽象是否适用于更大的数字？扩展测试套件，使其适用于罗马数字 50，即“L”。

Here are some test cases, try and make them pass.

这里有一些测试用例，试着让它们通过。

```go
{"40 gets converted to XL", 40, "XL"},
{"47 gets converted to XLVII", 47, "XLVII"},
{"49 gets converted to XLIX", 49, "XLIX"},
{"50 gets converted to L", 50, "L"},
```

Need help? You can see what symbols to add in [this gist](https://gist.github.com/pamelafox/6c7b948213ba55332d86efd0f0b037de).

需要帮忙？您可以在 [this gist](https://gist.github.com/pamelafox/6c7b948213ba55332d86efd0f0b037de) 中查看要添加的符号。

## And the rest!

## 剩下的！

Here are the remaining symbols

这是剩下的符号

| Arabic | Roman |
| ------ |:---: |
| 100    | C   |
| 500    | D   |
| 1000   | M   |



Take the same approach for the remaining symbols, it should just be a matter of adding data to both the tests and our array of symbols.

对其余符号采取相同的方法，只需将数据添加到测试和我们的符号数组即可。

Does your code work for `1984`: `MCMLXXXIV` ?

您的代码是否适用于“1984”：“MCMLXXXIV”？

Here is my final test suite

这是我的最终测试套件

```go
func TestRomanNumerals(t *testing.T) {
    cases := []struct {
        Arabic int
        Roman  string
    }{
        {Arabic: 1, Roman: "I"},
        {Arabic: 2, Roman: "II"},
        {Arabic: 3, Roman: "III"},
        {Arabic: 4, Roman: "IV"},
        {Arabic: 5, Roman: "V"},
        {Arabic: 6, Roman: "VI"},
        {Arabic: 7, Roman: "VII"},
        {Arabic: 8, Roman: "VIII"},
        {Arabic: 9, Roman: "IX"},
        {Arabic: 10, Roman: "X"},
        {Arabic: 14, Roman: "XIV"},
        {Arabic: 18, Roman: "XVIII"},
        {Arabic: 20, Roman: "XX"},
        {Arabic: 39, Roman: "XXXIX"},
        {Arabic: 40, Roman: "XL"},
        {Arabic: 47, Roman: "XLVII"},
        {Arabic: 49, Roman: "XLIX"},
        {Arabic: 50, Roman: "L"},
        {Arabic: 100, Roman: "C"},
        {Arabic: 90, Roman: "XC"},
        {Arabic: 400, Roman: "CD"},
        {Arabic: 500, Roman: "D"},
        {Arabic: 900, Roman: "CM"},
        {Arabic: 1000, Roman: "M"},
        {Arabic: 1984, Roman: "MCMLXXXIV"},
        {Arabic: 3999, Roman: "MMMCMXCIX"},
        {Arabic: 2014, Roman: "MMXIV"},
        {Arabic: 1006, Roman: "MVI"},
        {Arabic: 798, Roman: "DCCXCVIII"},
    }
    for _, test := range cases {
        t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
            got := ConvertToRoman(test.Arabic)
            if got != test.Roman {
                t.Errorf("got %q, want %q", got, test.Roman)
            }
        })
    }
}
```

- I removed `description` as I felt the _data_ described enough of the information.
- I added a few other edge cases I found just to give me a little more confidence. With table based tests this is very cheap to do.

- 我删除了 `description`，因为我觉得 _data_ 描述了足够的信息。
- 我添加了一些我发现的其他边缘情况，只是为了让我更有信心。使用基于表的测试，这是非常便宜的。

I didn't change the algorithm, all I had to do was update the `allRomanNumerals` array.

我没有改变算法，我所要做的就是更新 `allRomanNumerals` 数组。

```go
var allRomanNumerals = []RomanNumeral{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}
```

## Parsing Roman Numerals

## 解析罗马数字

We're not done yet. Next we're going to write a function that converts _from_ a Roman Numeral to an `int`

我们还没有完成。接下来我们将编写一个函数将_from_罗马数字转换为`int`

## Write the test first

## 先写测试

We can re-use our test cases here with a little refactoring

我们可以通过一些重构来重用我们的测试用例

Move the `cases` variable outside of the test as a package variable in a `var` block.

将 `cases` 变量作为 `var` 块中的包变量移出测试。

```go
func TestConvertingToArabic(t *testing.T) {
    for _, test := range cases[:1] {
        t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
            got := ConvertToArabic(test.Roman)
            if got != test.Arabic {
                t.Errorf("got %d, want %d", got, test.Arabic)
            }
        })
    }
}
```

Notice I am using the slice functionality to just run one of the tests for now (`cases[:1]`) as trying to make all of those tests pass all at once is too big a leap

请注意，我现在使用切片功能只运行其中一个测试（`cases[:1]`），因为试图让所有这些测试一次全部通过是一个太大的飞跃

## Try to run the test

## 尝试运行测试

```
./numeral_test.go:60:11: undefined: ConvertToArabic
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Add our new function definition

添加我们的新函数定义

```go
func ConvertToArabic(roman string) int {
    return 0
}
```

The test should now run and fail

测试现在应该运行并失败

```
--- FAIL: TestConvertingToArabic (0.00s)
    --- FAIL: TestConvertingToArabic/'I'_gets_converted_to_1 (0.00s)
        numeral_test.go:62: got 0, want 1
```

## Write enough code to make it pass

## 编写足够的代码使其通过

You know what to do

你知道该做什么

```go
func ConvertToArabic(roman string) int {
    return 1
}
```

Next, change the slice index in our test to move to the next test case (e.g. `cases[:2]`). Make it pass yourself with the dumbest code you can think of, continue writing dumb code (best book ever right?) for the third case too. Here's my dumb code.

接下来，更改我们测试中的切片索引以移动到下一个测试用例（例如`cases[:2]`）。用你能想到的最愚蠢的代码让它通过，继续为第三种情况编写愚蠢的代码（有史以来最好的书吗？）。这是我的愚蠢代码。

```go
func ConvertToArabic(roman string) int {
    if roman == "III" {
        return 3
    }
    if roman == "II" {
        return 2
    }
    return 1
}
```

Through the dumbness of _real code that works_ we can start to see a pattern like before. We need to iterate through the input and build _something_, in this case a total.

通过_真正有效的代码的愚蠢_我们可以开始看到像以前一样的模式。我们需要遍历输入并构建 _something_，在这种情况下是一个总数。

```go
func ConvertToArabic(roman string) int {
    total := 0
    for range roman {
        total++
    }
    return total
}
```

## Write the test first

## 先写测试

Next we move to `cases[:4]` (`IV`) which now fails because it gets 2 back as that's the length of the string.

接下来我们转到 `cases[:4]` (`IV`)，它现在失败了，因为它返回 2，因为这是字符串的长度。

## Write enough code to make it pass

## 编写足够的代码使其通过

```go
// earlier..
type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
    for _, s := range r {
        if s.Symbol == symbol {
            return s.Value
        }
    }

    return 0
}

// later..
func ConvertToArabic(roman string) int {
    total := 0

    for i := 0;i < len(roman);i++ {
        symbol := roman[i]

        // look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
        if i+1 < len(roman) && symbol == 'I' {
            nextSymbol := roman[i+1]

            // build the two character string
            potentialNumber := string([]byte{symbol, nextSymbol})

            // get the value of the two character string
            value := allRomanNumerals.ValueOf(potentialNumber)

            if value != 0 {
                total += value
                i++ // move past this character too for the next loop
            } else {
                total++
            }
        } else {
            total++
        }
    }
    return total
}
```

This is horrible but it does work. It's so bad I felt the need to add comments.

这很可怕，但确实有效。太糟糕了，我觉得有必要添加评论。

- I wanted to be able to look up an integer value for a given roman numeral so I made a type from our array of `RomanNumeral`s and then added a method to it, `ValueOf`
- Next in our loop we need to look ahead _if_ the string is big enough _and the current symbol is a valid subtractor_. At the moment it's just `I` (1) but can also be `X` (10) or `C` (100).
     - If it satisfies both of these conditions we need to lookup the value and add it to the total _if_ it is one of the special subtractors, otherwise ignore it
     - Then we need to further increment `i` so we don't count this symbol twice

- 我希望能够查找给定罗马数字的整数值，所以我从我们的 `RomanNumeral`s 数组中创建了一个类型，然后向它添加了一个方法，`ValueOf`
- 接下来在我们的循环中，我们需要向前看_如果_字符串足够大_并且当前符号是有效的减法器_。目前它只是`I`（1），但也可以是`X`（10）或`C`（100）。
    - 如果它同时满足这两个条件，我们需要查找该值并将其添加到总数中 _if_ 它是特殊减法器之一，否则忽略它
    - 然后我们需要进一步增加`i`，这样我们就不会计算这个符号两次

## Refactor 

## 重构

I'm not entirely convinced this will be the long-term approach and there's potentially some interesting refactors we could do, but I'll resist that in case our approach is totally wrong. I'd rather make a few more tests pass first and see. For the meantime I made the first `if` statement slightly less horrible.

我并不完全相信这将是一种长期的方法，并且我们可能会进行一些有趣的重构，但我会拒绝这样做，以防我们的方法完全错误。我宁愿先通过更多的测试然后看看。与此同时，我做了第一个“if”语句，但不那么可怕。

```go
func ConvertToArabic(roman string) int {
    total := 0

    for i := 0;i < len(roman);i++ {
        symbol := roman[i]

        if couldBeSubtractive(i, symbol, roman) {
            nextSymbol := roman[i+1]

            // build the two character string
            potentialNumber := string([]byte{symbol, nextSymbol})

            // get the value of the two character string
            value := allRomanNumerals.ValueOf(potentialNumber)

            if value != 0 {
                total += value
                i++ // move past this character too for the next loop
            } else {
                total++
            }
        } else {
            total++
        }
    }
    return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
    return index+1 < len(roman) && currentSymbol == 'I'
}
```

## Write the test first

## 先写测试

Let's move on to `cases[:5]`

让我们继续讨论`cases[:5]`

```
=== RUN   TestConvertingToArabic/'V'_gets_converted_to_5
    --- FAIL: TestConvertingToArabic/'V'_gets_converted_to_5 (0.00s)
        numeral_test.go:62: got 1, want 5
```

## Write enough code to make it pass

## 编写足够的代码使其通过

Apart from when it is subtractive our code assumes that every character is a `I` which is why the value is 1. We should be able to re-use our `ValueOf` method to fix this.

除了当它是减法时，我们的代码假设每个字符都是一个“I”，这就是值是 1 的原因。我们应该能够重新使用我们的“ValueOf”方法来解决这个问题。

```go
func ConvertToArabic(roman string) int {
    total := 0

    for i := 0;i < len(roman);i++ {
        symbol := roman[i]

        // look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
        if couldBeSubtractive(i, symbol, roman) {
            nextSymbol := roman[i+1]

            // build the two character string
            potentialNumber := string([]byte{symbol, nextSymbol})

            if value := allRomanNumerals.ValueOf(potentialNumber);value != 0 {
                total += value
                i++ // move past this character too for the next loop
            } else {
                total++ // this is fishy...
            }
        } else {
            total+=allRomanNumerals.ValueOf(string([]byte{symbol}))
        }
    }
    return total
}
```

## Refactor

## 重构

When you index strings in Go, you get a `byte`. This is why when we build up the string again we have to do stuff like `string([]byte{symbol})`. It's repeated a couple of times, let's just move that functionality so that `ValueOf` takes some bytes instead.

当你在 Go 中索引字符串时，你会得到一个 `byte`。这就是为什么当我们再次构建字符串时，我们必须执行诸如“string([]byte{symbol})”之类的操作。它重复了几次，让我们移动这个功能，以便 `ValueOf` 取一些字节。

```go
func (r RomanNumerals) ValueOf(symbols ...byte) int {
    symbol := string(symbols)
    for _, s := range r {
        if s.Symbol == symbol {
            return s.Value
        }
    }

    return 0
}
```

Then we can just pass in the bytes as is, to our function

然后我们可以按原样将字节传递给我们的函数

```go
func ConvertToArabic(roman string) int {
    total := 0

    for i := 0;i < len(roman);i++ {
        symbol := roman[i]

        if couldBeSubtractive(i, symbol, roman) {
            if value := allRomanNumerals.ValueOf(symbol, roman[i+1]);value != 0 {
                total += value
                i++ // move past this character too for the next loop
            } else {
                total++ // this is fishy...
            }
        } else {
            total+=allRomanNumerals.ValueOf(symbol)
        }
    }
    return total
}
```

It's still pretty nasty, but it's getting there.

它仍然很讨厌，但它正在到达那里。

If you start moving our `cases[:xx]` number through you'll see that quite a few are passing now. Remove the slice operator entirely and see which ones fail, here's some examples from my suite

如果你开始移动我们的 `cases[:xx]` 数字，你会发现现在有很多都通过了。完全删除切片运算符并查看哪些失败，这是我套件中的一些示例

```
=== RUN   TestConvertingToArabic/'XL'_gets_converted_to_40
    --- FAIL: TestConvertingToArabic/'XL'_gets_converted_to_40 (0.00s)
        numeral_test.go:62: got 60, want 40
=== RUN   TestConvertingToArabic/'XLVII'_gets_converted_to_47
    --- FAIL: TestConvertingToArabic/'XLVII'_gets_converted_to_47 (0.00s)
        numeral_test.go:62: got 67, want 47
=== RUN   TestConvertingToArabic/'XLIX'_gets_converted_to_49
    --- FAIL: TestConvertingToArabic/'XLIX'_gets_converted_to_49 (0.00s)
        numeral_test.go:62: got 69, want 49
```

I think all we're missing is an update to `couldBeSubtractive` so that it accounts for the other kinds of subtractive symbols

我认为我们缺少的只是对“couldBeSubtractive”的更新，以便它考虑其他类型的减法符号

```go
func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
    isSubtractiveSymbol := currentSymbol == 'I' ||currentSymbol == 'X' ||currentSymbol =='C'
    return index+1 < len(roman) && isSubtractiveSymbol
}
```

Try again, they still fail. However we left a comment earlier...

再试一次，他们仍然失败。但是我们之前留下了评论......

```go
total++ // this is fishy...
```

We should never be just incrementing `total` as that implies every symbol is a `I`. Replace it with:

我们永远不应该只是增加 `total`，因为这意味着每个符号都是一个 `I`。替换为：

```go
total += allRomanNumerals.ValueOf(symbol)
```

And all the tests pass! Now that we have fully working software we can indulge ourselves in some refactoring, with confidence.

并且所有的测试都通过了！现在我们有了完全可以工作的软件，我们可以放心地进行一些重构。

## Refactor

## 重构

Here is all the code I finished up with. I had a few failed attempts but as I keep emphasising, that's fine and the tests help me play around with the code freely.

这是我完成的所有代码。我有一些失败的尝试，但正如我一直强调的那样，这很好，测试帮助我自由地玩弄代码。

```go
import "strings"

func ConvertToArabic(roman string) (total int) {
    for _, symbols := range windowedRoman(roman).Symbols() {
        total += allRomanNumerals.ValueOf(symbols...)
    }
    return
}

func ConvertToRoman(arabic int) string {
    var result strings.Builder

    for _, numeral := range allRomanNumerals {
        for arabic >= numeral.Value {
            result.WriteString(numeral.Symbol)
            arabic -= numeral.Value
        }
    }

    return result.String()
}

type romanNumeral struct {
    Value  int
    Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
    symbol := string(symbols)
    for _, s := range r {
        if s.Symbol == symbol {
            return s.Value
        }
    }

    return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
    symbol := string(symbols)
    for _, s := range r {
        if s.Symbol == symbol {
            return true
        }
    }
    return false
}

var allRomanNumerals = romanNumerals{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
    for i := 0;i < len(w);i++ {
        symbol := w[i]
        notAtEnd := i+1 < len(w)

        if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
            symbols = append(symbols, []byte{symbol, w[i+1]})
            i++
        } else {
            symbols = append(symbols, []byte{symbol})
        }
    }
    return
}

func isSubtractive(symbol uint8) bool {
    return symbol == 'I' ||symbol == 'X' ||symbol == 'C'
}
```

My main problem with the previous code is similar to our refactor from earlier. We had too many concerns coupled together. We wrote an algorithm which was trying to extract Roman Numerals from a string _and_ then find their values.

我之前代码的主要问题类似于我们之前的重构。我们有太多的担忧交织在一起。我们编写了一个算法，试图从字符串 _and_ 中提取罗马数字，然后找到它们的值。

So I created a new type `windowedRoman` which took care of extracting the numerals, offering a `Symbols` method to retrieve them as a slice. This meant our `ConvertToArabic` function could simply iterate over the symbols and total them.

所以我创建了一个新的类型 `windowedRoman`，它负责提取数字，提供一个 `Symbols` 方法来将它们作为切片检索。这意味着我们的 `ConvertToArabic` 函数可以简单地迭代符号并将它们汇总。

I broke the code down a bit by extracting some functions, especially around the wonky if statement to figure out if the symbol we are currently dealing with is a two character subtractive symbol.

我通过提取一些函数来稍微分解代码，特别是在不稳定的 if 语句周围，以确定我们当前处理的符号是否是两个字符的减法符号。

There's probably a more elegant way but I'm not going to sweat it. The code is there and it works and it is tested. If I (or anyone else) finds a better way they can safely change it - the hard work is done.

可能有一种更优雅的方式，但我不会出汗。代码就在那里，它可以工作并经过测试。如果我（或其他任何人）找到更好的方法，他们可以安全地更改它 - 艰苦的工作已经完成。

## An intro to property based tests

## 基于属性的测试介绍

There have been a few rules in the domain of Roman Numerals that we have worked with in this chapter

我们在本章中使用了罗马数字领域的一些规则

- Can't have more than 3 consecutive symbols
- Only I (1), X (10) and C (100) can be "subtractors"
- Taking the result of `ConvertToRoman(N)` and passing it to `ConvertToArabic` should return us `N`

- 不能有超过 3 个连续符号
- 只有 I (1), X (10) 和 C (100) 可以是“减法器”
- 获取 `ConvertToRoman(N)` 的结果并将其传递给 `ConvertToArabic` 应该返回我们的 `N`

The tests we have written so far can be described as "example" based tests where we provide the tooling some examples around our code to verify.

到目前为止，我们编写的测试可以被描述为基于“示例”的测试，其中我们提供了一些围绕我们代码的示例来验证的工具。

What if we could take these rules that we know about our domain and somehow exercise them against our code?

如果我们可以采用这些我们知道的关于我们的领域的规则，并以某种方式针对我们的代码执行这些规则会怎样？

Property based tests help you do this by throwing random data at your code and verifying the rules you describe always hold true. A lot of people think property based tests are mainly about random data but they would be mistaken. The real challenge about property based tests is having a _good_ understanding of your domain so you can write these properties.

基于属性的测试通过向您的代码抛出随机数据并验证您描述的规则始终适用来帮助您做到这一点。很多人认为基于属性的测试主要是关于随机数据，但他们会错的。基于属性的测试的真正挑战是对您的领域有一个_良好的_了解，以便您可以编写这些属性。

Enough words, let's see some code

话不多说，看代码

```go
func TestPropertiesOfConversion(t *testing.T) {
    assertion := func(arabic int) bool {
        roman := ConvertToRoman(arabic)
        fromRoman := ConvertToArabic(roman)
        return fromRoman == arabic
    }

    if err := quick.Check(assertion, nil);err != nil {
        t.Error("failed checks", err)
    }
}
```

### Rationale of property

### 财产的理由

Our first test will check that if we transform a number into Roman, when we use our other function to convert it back to a number that we get what we originally had.

我们的第一个测试将检查我们是否将数字转换为罗马数字，当我们使用其他函数将其转换回我们得到的原始数字时。

- Given random number (e.g `4`).
- Call `ConvertToRoman` with random number (should return `IV` if `4`).
- Take the result of above and pass it to `ConvertToArabic`.
- The above should give us our original input (`4`). 

- 给定随机数（例如`4`）。
- 使用随机数调用 `ConvertToRoman`（如果为 `4`，则应返回 `IV`）。
- 将上面的结果传递给 `ConvertToArabic`。
- 以上应该给我们原始输入（`4`）。

This feels like a good test to build us confidence because it should break if there's a bug in either. The only way it could pass is if they have the same kind of bug; which isn't impossible but feels unlikely.

这感觉是建立我们信心的一个很好的测试，因为如果其中任何一个存在错误，它就会中断。它可以通过的唯一方法是它们是否有相同类型的错误；这并非不可能，但感觉不太可能。

### Technical explanation

  ### 技术说明

 We're using the [testing/quick](https://golang.org/pkg/testing/quick/) package from the standard library

  我们正在使用标准库中的 [testing/quick](https://golang.org/pkg/testing/quick/) 包

 Reading from the bottom, we provide `quick.Check` a function that it will run against a number of random inputs, if the function returns `false` it will be seen as failing the check.

  从底部看，我们提供了 `quick.Check` 一个函数，它将针对一些随机输入运行，如果函数返回 `false`，它将被视为未通过检查。

 Our `assertion` function above takes a random number and runs our functions to test the property.

  我们上面的 `assertion` 函数接受一个随机数并运行我们的函数来测试属性。

 ### Run our test

  ### 运行我们的测试

 Try running it; your computer may hang for a while, so kill it when you're bored :)

  尝试运行它；你的电脑可能会挂一段时间，所以当你无聊的时候把它关掉:)

 What's going on? Try adding the following to the assertion code.

  这是怎么回事？尝试将以下内容添加到断言代码中。

 ```go
assertion := func(arabic int) bool {
    if arabic <0 ||arabic > 3999 {
        log.Println(arabic)
        return true
    }
    roman := ConvertToRoman(arabic)
    fromRoman := ConvertToArabic(roman)
    return fromRoman == arabic
}
 ```

You should see something like this:

您应该会看到如下内容：

```
=== RUN   TestPropertiesOfConversion
2019/07/09 14:41:27 6849766357708982977
2019/07/09 14:41:27 -7028152357875163913
2019/07/09 14:41:27 -6752532134903680693
2019/07/09 14:41:27 4051793897228170080
2019/07/09 14:41:27 -1111868396280600429
2019/07/09 14:41:27 8851967058300421387
2019/07/09 14:41:27 562755830018219185
```

Just running this very simple property has exposed a flaw in our implementation. We used `int` as our input but:
- You can't do negative numbers with Roman Numerals
- Given our rule of a max of 3 consecutive symbols we can't represent a value greater than 3999 ([well, kinda](https://www.quora.com/Which-is-the-maximum-number-in-Roman-numerals)) and `int` has a much higher maximum value than 3999.

仅仅运行这个非常简单的属性就暴露了我们实现中的一个缺陷。我们使用 `int` 作为我们的输入，但是：
- 你不能用罗马数字做负数
- 鉴于我们最多 3 个连续符号的规则，我们不能表示大于 3999 的值（[好吧，有点](https://www.quora.com/Which-is-the-maximum-number-in-Roman-numerals)) 和 `int` 的最大值比 3999 高得多。

This is great! We've been forced to think more deeply about our domain which is a real strength of property based tests.

这很棒！我们被迫更深入地思考我们的领域，这是基于属性的测试的真正优势。

Clearly `int` is not a great type. What if we tried something a little more appropriate?

显然 `int` 不是一个很好的类型。如果我们尝试一些更合适的方法呢？

### [`uint16`](https://golang.org/pkg/builtin/#uint16)



Go has types for _unsigned integers_, which means they cannot be negative; so that rules out one class of bug in our code immediately. By adding 16, it means it is a 16 bit integer which can store a max of `65535`, which is still too big but gets us closer to what we need.

Go 有 _unsigned integers_ 的类型，这意味着它们不能为负数；这样就可以立即排除我们代码中的一类错误。通过添加 16，这意味着它是一个 16 位整数，最多可以存储“65535”，这仍然太大，但让我们更接近我们需要的东西。

Try updating the code to use `uint16` rather than `int`. I updated `assertion` in the test to give a bit more visibility.

尝试更新代码以使用 `uint16` 而不是 `int`。我在测试中更新了 `assertion` 以提供更多可见性。

```go
assertion := func(arabic uint16) bool {
    if arabic > 3999 {
        return true
    }
    t.Log("testing", arabic)
    roman := ConvertToRoman(arabic)
    fromRoman := ConvertToArabic(roman)
    return fromRoman == arabic
}
```

If you run the test they now actually run and you can see what is being tested. You can run multiple times to see our code stands up well to the various values! This gives me a lot of confidence that our code is working how we want.

如果您运行测试，它们现在实际运行，您可以看到正在测试的内容。您可以多次运行以查看我们的代码对各种值的支持情况！这给了我很大的信心，我们的代码正在按照我们想要的方式工作。

The default number of runs `quick.Check` performs is 100 but you can change that with a config.

`quick.Check` 执行的默认运行次数是 100，但您可以通过配置更改它。

```go
if err := quick.Check(assertion, &quick.Config{
    MaxCount:1000,
});err != nil {
    t.Error("failed checks", err)
}
```

### Further work

###  进一步的工作

- Can you write property tests that check the other properties we described?
- Can you think of a way of making it so it's impossible for someone to call our code with a number greater than 3999?
     - You could return an error
     - Or create a new type that cannot represent > 3999
         - What do you think is best?

- 你能编写属性测试来检查我们描述的其他属性吗？
- 你能想出一种方法来使某人无法使用大于 3999 的号码调用我们的代码吗？
    - 你可能会返回一个错误
    - 或者创建一个不能代表> 3999的新类型
        - 你认为什么是最好的？

## Wrapping up

##  总结

### More TDD practice with iterative development

### 更多带有迭代开发的 TDD 实践

Did the thought of writing code that converts 1984 into MCMLXXXIV feel intimidating to you at first? It did to me and I've been writing software for quite a long time.

一开始写代码将 1984 转换成 MCMLXXXIV 的想法让你感到害怕吗？对我来说确实如此，而且我已经编写软件很长时间了。

The trick, as always, is to **get started with something simple** and take **small steps**.

与往常一样，诀窍是**从简单的事情开始**并采取**小步骤**。

At no point in this process did we make any large leaps, do any huge refactorings, or get in a mess. 

在这个过程中，我们没有进行任何大的飞跃，没有进行任何大规模的重构，也没有陷入混乱。

I can hear someone cynically saying "this is just a kata". I can't argue with that, but I still take this same approach for every project I work on. I never ship a big distributed system in my first step, I find the simplest thing the team could ship (usually a "Hello world" website) and then iterate on small bits of functionality in manageable chunks, just like how we did here.

我可以听到有人愤世嫉俗地说“这只是一个套路”。我无法反驳，但我仍然对我从事的每个项目采取同样的方法。在我的第一步中，我从未发布过大型分布式系统，我找到了团队可以发布的最简单的东西（通常是“Hello world”网站），然后在可管理的块中迭代一小部分功能，就像我们在这里所做的那样。

The skill is knowing _how_ to split work up, and that comes with practice and with some lovely TDD to help you on your way.

技能是知道_如何_拆分工作，这伴随着练习和一些可爱的 TDD 来帮助你前进。

### Property based tests

### 基于属性的测试

- Built into the standard library
- If you can think of ways to describe your domain rules in code, they are an excellent tool for giving you more confidence
- Force you to think about your domain deeply
- Potentially a nice complement to your test suite

- 内置于标准库中
- 如果你能想办法用代码描述你的领域规则，它们是一个让你更有信心的绝佳工具
- 强迫你深入思考你的领域
- 可能是您测试套件的一个很好的补充

## Postscript

## 后记

This book is reliant on valuable feedback from the community.
[Dave](http://github.com/gypsydave5) is an enormous help in practically every chapter. But he had a real rant about my use of 'Arabic numerals' in this chapter so, in the interests of full disclosure, here's what he said.

这本书依赖于来自社区的宝贵反馈。
[Dave](http://github.com/gypsydave5) 对几乎所有章节。但他对我在这本书中使用“阿拉伯数字”大发雷霆章如此，为了充分披露，这是他所说的。
