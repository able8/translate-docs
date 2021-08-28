# Developing price and currency handling for Go

  # 为 Go 开发价格和货币处理

 September 2020 From: https://bojanz.github.io/price-currency-handling-go/

Now that [bojanz/currency](https://github.com/bojanz/currency)is listed on [Awesome Go](https://awesome-go.com/), it's a good time to reflect on the ideas that made it a reality. Back in March I started using Go for a few side projects, and I ran into the need to represent and handle currency amounts (prices). After some research I realized that the ecosystem was missing a complete solution.

现在 [bojanz/currency](https://github.com/bojanz/currency)已在 [Awesome Go](https://awesome-go.com/) 上列出，现在是反思这些想法的好时机使它成为现实。早在 3 月份，我就开始在一些副项目中使用 Go，我遇到了表示和处理货币金额（价格）的需求。经过一些研究，我意识到生态系统缺少一个完整的解决方案。

### What do we need?

###  我们需要什么？

Let’s sketch out our requirements.

让我们勾勒出我们的要求。

```c
type Amount struct {
    number       decimal
    currencyCode string
}

// Methods: NewAmount(), Add, Sub, Mul, Div, Round, Cmp...
```


- A currency amount consists of a decimal number and a currency code.
- Trying to combine or compare amounts of different currencies returns an error.
- Like time.Time, a currency.Amount has value semantics, which means  that amounts are immutable. Adding amount A to amount B produces a third amount C instead of modifying B.
- The number and currencyCode are unexported to require usage of the appropriate methods.

- 货币金额由十进制数和货币代码组成。
- 尝试合并或比较不同货币的数量会返回错误。
- 与 time.Time 一样，currency.Amount 具有值语义，这意味着数量是不可变的。将数量 A 添加到数量 B 会产生第三个数量 C，而不是修改 B。
- 数字和货币代码未导出，需要使用适当的方法。

We should be able to get basic information about a currency: numeric code, number of digits (used for rounding), and symbol ($, £, €, etc).

我们应该能够获得有关货币的基本信息：数字代码、位数（用于四舍五入）和符号（$、£、€ 等）。

We should be able to format an amount for display, getting “$10”, “10 €”, etc. Formatting is locale specific, not currency specific (“10 €” for fr-CH, “€ 10” for de-CH) .

我们应该能够格式化显示的金额，获得“$10”、“10 €”等。格式化是特定于语言环境的，而不是特定于货币的（“10 €”代表 fr-CH，“€10”代表 de-CH） .

### A decimal journey

### 十进制旅程

Our first problem is that Go doesn’t have a builtin type for decimal numbers. Developers learn early on that [floats must never be used](https://husobee.github.io/money/float/2016/09/23/never-use-floats-for-currency.html) instead, because they are imprecise, and as amounts are multiplied, divided, rounded and summed up, those  imprecisions add up, quickly becoming real business problems.

我们的第一个问题是 Go 没有十进制数的内置类型。开发人员很早就了解到 [永远不能使用浮动](https://husobee.github.io/money/float/2016/09/23/never-use-floats-for-currency.html) 相反，因为它们是不精确，随着数量的乘法、除法、四舍五入和求和，这些不精确性加起来，很快成为真正的业务问题。

An old and common workaround is to store the amount in its minor  units (e.g. cents) as an integer, representing $5.99 as 599. But every trick has its cost. No amount can  have sub-minor-unit precision (e.g. “5.884”), which is needed for certain kinds of products (e.g. selling in bulk) and in certain tax jurisdictions (e.g. EU VAT). Handling multiple currencies becomes more difficult, as different  currencies have different numbers of decimals (JPY has 0, KWD has 3), making it harder to order by amount in the database.

一种古老而常见的解决方法是将其次要单位（例如美分）中的金额存储为整数，将 5.99 美元表示为 599。但每个技巧都有其成本。任何金额都不能具有次要单位精度（例如“5.884”），这是某些类型的产品（例如批量销售）和某些税收管辖区（例如欧盟增值税）所需要的。处理多种货币变得更加困难，因为不同的货币有不同的小数位数（日元有 0，KWD 有 3），这使得在数据库中按金额排序变得更加困难。

Luckily, Go has two solid packages that implement decimals in userspace. The first one is [cockroachdb/apd](https://github.com/cockroachdb/apd). It is well maintained and fast enough, solving our need. The API is not very friendly:

幸运的是，Go 有两个在用户空间中实现小数的可靠包。第一个是 [cockroachdb/apd](https://github.com/cockroachdb/apd)。它维护良好且速度足够快，解决了我们的需求。 API 不是很友好：

```c
// a + b = c
c := apd.New(0, 0)
ctx := apd.BaseContext.WithPrecision(16)
ctx.Add(c, a, b)

// round d to 2 decimals.
result := apd.New(0, 0)
ctx := apd.BaseContext.WithPrecision(16)
ctx.Rounding = apd.RoundHalfUp
ctx.Quantize(result, d, -2)
```


However, since we have our own methods for arithmetic and comparisons, we can wrap the apd logic, never even exposing the underlying implementation to the user. We accept strings, and use them to instantiate the underlying type:

然而，由于我们有自己的算术和比较方法，我们可以包装 apd 逻辑，甚至永远不会将底层实现暴露给用户。我们接受字符串，并使用它们来实例化底层类型：

```
amount, _ := currency.NewAmount("20.99", "USD")
taxAmount, _ := amount.Mul("0.20")
// Methods use apd.NewFromString(n) to get a decimal.
```


This will also serve us well if we decide to switch the underlying decimal implementation, for example to [ericlagergren/decimal](https://github.com/ericlagergren/decimal)which is faster but has seen [instability](https:/ /github.com/ericlagergren/decimal/issues/154) due to slower maintanance this year.

如果我们决定切换底层的十进制实现，这也会很好地为我们服务，例如切换到 [ericlagergren/decimal](https://github.com/ericlagergren/decimal)，它更快但已经看到[不稳定](https:/ /github.com/ericlagergren/decimal/issues/154) 由于今年的维护速度较慢。

### Where do currencies come from? 

### 货币从何而来？

Inflation happens, old currencies get deprecated, new currencies get introduced. It pays off to generate currency data from an external source, so that new data is always one `go generate` away. Currency codes and their numeric codes can be retrieved from [ISO](https://www.currency-iso.org/dam/downloads/lists/list_one.xml).Locale-specific data, such as currency names, symbols, formatting rules are taken from [CLDR](http://cldr.unicode.org/), a rare case of the entire industry cooperating on a common problem.

通货膨胀发生，旧货币被弃用，新货币被引入。从外部来源生成货币数据是值得的，因此新数据总是“一去不复返”。货币代码及其数字代码可以从 [ISO](https://www.currency-iso.org/dam/downloads/lists/list_one.xml)中检索。特定于语言环境的数据，例如货币名称、符号、格式规则，取自 [CLDR](http://cldr.unicode.org/)，这是整个行业在共同问题上合作的罕见案例。

The problem with CLDR data is that there’s megabytes of it, adding to binary sizes and memory usage. Let’s try to reduce this weight.

CLDR 数据的问题在于它有数兆字节，增加了二进制大小和内存使用量。让我们试着减轻这个重量。

The first trick is to reduce the number of locales for which data is  generated. CLDR has 542 locales, but it is not likely that an  application will need to format prices in Church Slavic or Esperanto. Chrome uses an allowlist, while we opted for a denylist listing each ignored locale,  allowing community members to re-include a locale if they end up needing it.

第一个技巧是减少生成数据的语言环境的数量。 CLDR 有 542 个语言环境，但应用程序不太可能需要在教堂斯拉夫语或世界语中格式化价格。 Chrome 使用许可名单，而我们选择了一个拒绝名单，列出每个被忽略的区域设置，允许社区成员在最终需要时重新包含区域设置。

The second trick is to stop shipping currency names, since they are  rarely used on the backend and can be retrieved on the frontend. Currencies tend to be identified by their code (USD) or their symbol  ($), while currency names are usually left for certain lists in the UI. With a few lines of javascript ([Intl.DisplayNames](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DisplayNames)), the frontend can retrieve a localized currency name for each code.

第二个技巧是停止发送货币名称，因为它们很少在后端使用并且可以在前端检索。货币往往由它们的代码 (USD) 或它们的符号 ($) 来标识，而货币名称通常会留给 UI 中的某些列表。使用几行 javascript ([Intl.DisplayNames](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DisplayNames))，前端可以检索本地化的货币每个代码的名称。

The third trick is to deduplicate locales by parent, relying on the package performing locale fallback. If “fr-FR” and “fr” have the same data, “fr-FR” is removed, and the package selects “fr” instead.

第三个技巧是通过父级对语言环境进行重复数据删除，依赖于执行语言环境回退的包。如果“fr-FR”和“fr”有相同的数据，则去掉“fr-FR”，包选择“fr”。

Finaly, symbols are grouped, to reduce repetition:

最后，符号被分组，以减少重复：

```c
"CAD": {
    {"CA$", []string{"en"}},
    {"$", []string{"en-CA", "fr-CA"}},
    {"$CA", []string{"fa", "fr"}},
    {"C$", []string{"nl"}},
},
```


Our [gen.go](https://github.com/bojanz/currency/blob/master/gen.go)is 800 lines of scary code, but the result is worth it. The generated [data.go](https://github.com/bojanz/currency/blob/master/data.go) is only 30kb, adding around 128kb to binary size.

我们的 [gen.go](https://github.com/bojanz/currency/blob/master/gen.go)是 800 行可怕的代码，但结果是值得的。生成的 [data.go](https://github.com/bojanz/currency/blob/master/data.go) 只有 30kb，二进制大小增加了大约 128kb。

### Putting it all together

### 把它们放在一起

We now have an amount struct, formatting data, symbols. The final step is to create a formatter.

我们现在有一个数量结构，格式化数据，符号。最后一步是创建格式化程序。

The formatter is about 200 lines of code long and respects locale-specific symbol positioning, grouping and decimal separators, group sizes, numbering systems, etc. It has the full [set of options](https://github.com/bojanz/currency/blob/master/formatter.go#L40) offered by NumberFormatter APIs in programming languages such as PHP, Java, Swift, etc.

格式化程序长约 200 行代码，并考虑特定于语言环境的符号定位、分组和小数分隔符、组大小、编号系统等。它具有完整的 [选项集](https://github.com/bojanz/货币/blob/master/formatter.go#L40) 由 NumberFormatter API 以 PHP、Java、Swift 等编程语言提供。

```c
locale := currency.NewLocale("tr")
formatter := currency.NewFormatter(locale)
amount, _ := currency.NewAmount("1245.988", "EUR")
fmt.Println(formatter.Format(amount)) // €1.245,988

formatter.MaxDigits = 2
fmt.Println(formatter.Format(amount)) // €1.245,99

formatter.NoGrouping = true
fmt.Println(formatter.Format(amount)) // €1245,99

formatter.CurrencyDisplay = currency.DisplayCode
fmt.Println(formatter.Format(amount)) // EUR 1245,99

// Different numbering system.
amount, _ := currency.NewAmount("1234.59", "IRR")
locale := currency.NewLocale("fa")
formatter := currency.NewFormatter(locale)
fmt.Println(formatter.Format(amount)) // ‎ریال ۱٬۲۳۴٫۵۹
```


### Conclusion

###  结论

With the right approach to data and the right set of constraints, we manage to solve currency handling with minimum cost (~2500 lines of code, ~30kb of data). However, the use case would be greatly helped by Go having a decimal type built in. I remain hopeful. 

通过正确的数据处理方法和正确的约束集，我们设法以最低成本（约 2500 行代码，约 30kb 数据）解决货币处理问题。但是，Go 内置十进制类型会对用例有很大帮助。我仍然充满希望。

