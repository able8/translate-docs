# Clean Go Code

# 干净的 Go 代码

From: https://github.com/Pungyeon/clean-go-article

## Preface: Why Write Clean Code?

## 前言：为什么要编写干净的代码？

This document is a reference for the Go community that aims to help developers write cleaner code. Whether you're working on a personal project or as part of a larger team, writing clean code is an important skill to have. Establishing good paradigms and consistent, accessible standards for writing clean code can help prevent developers from wasting many meaningless hours on trying to understand their own (or others') work.

本文档是 Go 社区的参考，旨在帮助开发人员编写更简洁的代码。无论您是在从事个人项目还是作为更大团队的一部分，编写干净的代码都是一项重要的技能。为编写干净的代码建立良好的范例和一致的、可访问的标准可以帮助防止开发人员浪费许多无意义的时间来尝试理解他们自己（或其他人）的工作。

> <em>We don’t read code, we <b>decode</b> it – Peter Seibel</em>

> <em>我们不阅读代码，我们<b>解码</b>它——Peter Seibel</em>

As developers, we're sometimes tempted to write code in a way that's convenient for the time being without regard for best practices; this makes code reviews and testing more difficult. In a sense, we're <em>encoding</em>—and, in doing so, making it more difficult for others to decode our work. But we want our code to be usable, readable, and maintainable. And that requires coding the <em>right</em> way, not the easy way.

作为开发人员，我们有时会想以一种暂时方便的方式编写代码，而不考虑最佳实践；这使得代码审查和测试更加困难。从某种意义上说，我们正在<em>编码</em>——这样做会让其他人更难解码我们的工作。但是我们希望我们的代码可用、可读和可维护。这需要以<em>正确</em>的方式进行编码，而不是简单的方式。

This document begins with a simple and short introduction to the fundamentals of writing clean code. Later, we'll discuss concrete refactoring examples specific to Go.

本文档首先简要介绍了编写干净代码的基础知识。稍后，我们将讨论特定于 Go 的具体重构示例。

##### A short word on `gofmt`
I'd like to take a few sentences to clarify my stance on `gofmt` because there are plenty of things I disagree with when it comes to this tool. I prefer snake case over camel case, and I quite like my constant variables to be uppercase. And, naturally, I also have many opinions on bracket placement. *That being said*, `gofmt` does allow us to have a common standard for writing Go code, and that's a great thing. As a developer myself, I can certainly appreciate that Go programmers may feel somewhat restricted by `gofmt`, especially if they disagree with some of its rules. But in my opinion, homogeneous code is more important than having complete expressive freedom.

##### 关于`gofmt`的简短词
我想用几句话来澄清我对 `gofmt` 的立场，因为在谈到这个工具时，我有很多不同意的地方。我更喜欢蛇形大小写而不是驼色大小写，而且我非常喜欢我的常量变量是大写的。而且，自然地，我对支架放置也有很多意见。 *话虽如此*，`gofmt` 确实让我们有一个共同的标准来编写 Go 代码，这是一件很棒的事情。作为一名开发人员，我当然可以理解 Go 程序员可能会受到 `gofmt` 的限制，特别是如果他们不同意它的一些规则。但在我看来，同构代码比拥有完全的表达自由更重要。

## Table of Contents
* [Introduction to Clean Code](#Introduction-to-Clean-Code)
     * [Test-Driven Development](#Test-Driven-Development)
     * [Naming Conventions](#Naming-Conventions)
     * * [Comments](#Comments)
         * [Function Naming](#Function-Naming)
         * [Variable Naming](#Variable-Naming)
     * [Cleaning Functions](#Cleaning-Functions)
       * [Function Length](#Function-Length)
       * [Function Signatures](#Function-Signatures)
     * [Variable Scope](#Variable-Scope)
     * [Variable Declaration](#Variable-Declaration)
    
* [Clean Go](#Clean-Go)
     * [Return Values](#Return-Values)
       * [Returning Defined Errors](#Returning-Defined-Errors)
       * [Returning Dynamic Errors](#Returning-Dynamic-Errors)
     * [Pointers in Go](#Pointers-in-Go)
     * [Closures Are Function Pointers](#Closures-are-Function-Pointers)
     * [Interfaces in Go](#Interfaces-in-Go)
     * [The Empty `interface{}`](#The-Empty-Interface)
* [Summary](#Summary)

##  目录
* [清洁代码简介](#Introduction-to-Clean-Code)
    * [测试驱动开发](#Test-Driven-Development)
    * [命名约定](#Naming-Conventions)
    * * [评论](#Comments)
        * [函数命名](#Function-Naming)
        * [变量命名](#Variable-Naming)
    * [清洁功能](#Cleaning-Functions)
      * [函数长度](#Function-Length)
      * [函数签名](#Function-Signatures)
    * [变量范围](#Variable-Scope)
    * [变量声明](#Variable-Declaration)
    
* [Clean Go](#Clean-Go)
    * [返回值](#Return-Values)
      * [返回定义的错误](#Returning-Defined-Errors)
      * [返回动态错误](#Returning-Dynamic-Errors)
    * [Go 中的指针](#Pointers-in-Go)
    * [闭包是函数指针](#Closures-are-Function-Pointers)
    * [Go 中的接口](#Interfaces-in-Go)
    * [空`接口{}`](#The-Empty-Interface)
* [总结](# 总结)

## Introduction to Clean Code

## 干净代码简介

Clean code is the pragmatic concept of promoting readable and maintainable software. Clean code establishes trust in the codebase and helps minimize the chances of careless bugs being introduced. It also helps developers maintain their agility, which typically plummets as the codebase expands due to the increased risk of introducing bugs.

干净的代码是促进可读和可维护软件的实用概念。干净的代码建立了对代码库的信任，并有助于最大限度地减少引入粗心错误的机会。它还可以帮助开发人员保持他们的敏捷性，由于引入错误的风险增加，敏捷性通常会随着代码库的扩展而下降。

### Test-Driven Development

### 测试驱动开发

Test-driven development is the practice of testing your code frequently throughout short development cycles or sprints. It ultimately contributes to code cleanliness by inviting developers to question the functionality and purpose of their code. To make testing easier, developers are encouraged to write short functions that only do one thing. For example, it's arguably much easier to test (and understand) a function that's only 4 lines long than one that's 40.

测试驱动开发是在较短的开发周期或冲刺中频繁测试代码的做法。通过邀请开发人员质疑其代码的功能和目的，它最终有助于代码清洁。为了使测试更容易，我们鼓励开发人员编写只做一件事的简短函数。例如，可以说测试（和理解）一个只有 4 行长的函数比一个 40 行长的函数要容易得多。

Test-driven development consists of the following cycle:

测试驱动的开发包括以下循环：

1. Write (or execute) a test
2. If the test fails, make it pass
3. Refactor your code accordingly
4. Repeat

1. 编写（或执行）一个测试
2.如果测试失败，让它通过
3. 相应地重构你的代码
4. 重复

Testing and refactoring are intertwined in this process. As you refactor your code to make it more understandable or maintainable, you need to test your changes thoroughly to ensure that you haven't altered the behavior of your functions. This can be incredibly useful as the codebase grows.

测试和重构在这个过程中交织在一起。当您重构代码以使其更易于理解或维护时，您需要彻底测试您的更改以确保您没有更改函数的行为。随着代码库的增长，这非常有用。

###  Naming Conventions

### 命名约定

#### Comments 

####  注释

I'd like to first address the topic of commenting code, which is an essential practice but tends to be misapplied. Unnecessary comments can indicate problems with the underlying code, such as the use of poor naming conventions. However, whether or not a particular comment is "necessary" is somewhat subjective and depends on how legibly the code was written. For example, the logic of well-written code may still be so complex that it requires a comment to clarify what is going on. In that case, one might argue that the comment is <em>helpful</em> and therefore necessary.

我想首先解决评论代码的话题，这是一个基本的做法，但往往会被误用。不必要的注释可能表明底层代码存在问题，例如使用了糟糕的命名约定。但是，特定注释是否“必要”有点主观，取决于代码编写的清晰程度。例如，编写良好的代码的逻辑可能仍然如此复杂，以至于需要注释来阐明发生了什么。在这种情况下，人们可能会争辩说该评论<em>有帮助</em>，因此是必要的。

In Go, according to `gofmt`, <em>all</em> public variables and functions should be annotated. I think this is absolutely fine, as it gives us consistent rules for documenting our code. However, I always want to distinguish between comments that enable auto-generated documentation and <em>all other</em> comments. Annotation comments, for documentation, should be written like documentation—they should be at a high level of abstraction and concern the logical implementation of the code as little as possible.

在 Go 中，根据 `gofmt`，<em>所有</em> 公共变量和函数都应该被注解。我认为这绝对没问题，因为它为我们记录代码提供了一致的规则。但是，我总是想区分启用自动生成文档的注释和<em>所有其他</em>注释。注释注释，对于文档，应该像文档一样编写——它们应该处于高抽象级别，并尽可能少地关注代码的逻辑实现。

I say this because there are other ways to explain code and ensure that it's being written comprehensibly and expressively. If the code is neither of those, some people find it acceptable to introduce a comment explaining the convoluted logic. Unfortunately, that doesn't really help. For one, most people simply won't read comments, as they tend to be very intrusive to the experience of reviewing code. Additionally, as you can imagine, a developer won't be too happy if they're forced to review unclear code that's been slathered with comments. The less that people have to read to understand what your code is doing, the better off they'll be.

我这样说是因为还有其他方法可以解释代码并确保其编写得易于理解和表达。如果代码不是这两者，有些人认为引入解释复杂逻辑的注释是可以接受的。不幸的是，这并没有真正的帮助。一方面，大多数人根本不会阅读评论，因为他们往往非常干扰审查代码的体验。此外，正如您可以想象的那样，如果开发人员被迫审查被评论覆盖的不明确的代码，他们将不会太高兴。人们为了解您的代码在做什么而需要阅读的内容越少，他们就会越好。

Let's take a step back and look at some concrete examples. Here's how you <em>shouldn't</em> comment your code:

让我们退后一步，看一些具体的例子。以下是您<em>不应该</em>评论您的代码的方式：

```go
// iterate over the range 0 to 9
// and invoke the doSomething function
// for each iteration
for i := 0;i < 10;i++ {
  doSomething(i)
}
```

This is what I like to call a <strong>tutorial comment</strong>; it's fairly common in tutorials, which often explain the low-level functionality of a language (or programming in general). While these comments may be helpful for beginners, they're absolutely useless in production code. Hopefully, we aren't collaborating with programmers who don't understand something as simple as a looping construct by the time they've begun working on a development team. As programmers, we shouldn't have to read the comment to understand what's going on—we know that we're iterating over the range 0 to 9 because we can simply read the code. Hence the proverb:

这就是我喜欢称之为<strong>教程评论</strong>的东西；这在教程中很常见，教程通常解释语言（或一般编程）的低级功能。虽然这些注释可能对初学者有帮助，但它们在生产代码中绝对没用。希望我们不会与那些在他们开始在开发团队中工作时不了解循环结构这样简单的东西的程序员合作。作为程序员，我们不应该通过阅读注释来理解发生了什么——我们知道我们在 0 到 9 的范围内进行迭代，因为我们可以简单地阅读代码。因此有谚语：

> <em>Document why, not how. – Venkat Subramaniam</em>

> <em>记录原因，而不是如何。 – Venkat Subramaniam</em>

Following this logic, we can now change our comment to explain <em>why</em> we are iterating from the range 0 to 9:

按照这个逻辑，我们现在可以改变我们的注释来解释<em>为什么</em>我们从 0 到 9 的范围内迭代：

```go
// instatiate 10 threads to handle upcoming work load
for i := 0;i < 10;i++ {
  doSomething(i)
}
```

Now we understand <em>why</em> we have a loop and can tell <em>what</em> we're doing by simply reading the code... Sort of.

现在我们明白了<em>为什么</em>我们有一个循环，并且可以通过简单地阅读代码来告诉<em>什么</em>我们在做什么......有点。

This still isn't what I'd consider clean code. The comment is worrying because it probably should not be necessary to express such an explanation in prose, assuming the code is well written (which it isn't). Technically, we're still saying what we're doing, not why we're doing it. We can easily express this "what" directly in our code by using more meaningful names:

这仍然不是我认为干净的代码。该评论令人担忧，因为假设代码编写得很好（事实并非如此），可能没有必要在散文中表达这样的解释。从技术上讲，我们仍然在说我们在做什么，而不是我们为什么要这样做。我们可以通过使用更有意义的名称直接在我们的代码中轻松表达这个“什么”：

```go
for workerID := 0;workerID < 10;workerID++ {
  instantiateThread(workerID)
}
```

With just a few changes to our variable and function names, we've managed to explain what we're doing directly in our code. This is much clearer for the reader because they won't have to read the comment and then map the prose to the code. Instead, they can simply read the code to understand what it's doing.

只需对我们的变量和函数名称进行一些更改，我们就设法直接在代码中解释了我们正在做什么。这对读者来说更清晰，因为他们不必阅读注释，然后将散文映射到代码。相反，他们可以简单地阅读代码以了解它在做什么。

Of course, this was a relatively trivial example. Writing clear and expressive code is unfortunately not always so easy; it can become increasingly difficult as the codebase itself grows in complexity. The more you practice writing comments in this mindset and avoid explaining what you're doing, the cleaner your code will become.

当然，这是一个相对微不足道的例子。不幸的是，编写清晰而富有表现力的代码并不总是那么容易。随着代码库本身的复杂性增加，它会变得越来越困难。以这种心态练习编写注释并避免解释您在做什么的次数越多，您的代码就会变得越干净。

#### Function Naming 

#### 函数命名

Let's now move on to function naming conventions. The general rule here is really simple: the more specific the function, the more general its name. In other words, we want to start with a very broad and short function name, such as `Run` or `Parse`, that describes the general functionality. Let's imagine that we are creating a configuration parser. Following this naming convention, our top level of abstraction might look something like the following:

现在让我们继续讨论函数命名约定。这里的一般规则非常简单：函数越具体，其名称就越笼统。换句话说，我们希望从一个非常广泛和简短的函数名称开始，例如 `Run` 或 `Parse`，它描述了一般功能。假设我们正在创建一个配置解析器。遵循此命名约定，我们的顶级抽象可能如下所示：

```go
func main() {
    configpath := flag.String("config-path", "", "configuration file path")
    flag.Parse()

    config, err := configuration.Parse(*configpath)
    
    ...
}
```

We'll focus on the naming of the `Parse` function. Despite this function's very short and general name, it's actually quite clear what it attempts to achieve.

我们将重点关注 `Parse` 函数的命名。尽管这个函数的名称非常简短和通用，但实际上它试图实现的目标非常清楚。

When we go one layer deeper, our function naming will become slightly more specific:

当我们深入一层时，我们的函数命名将变得更加具体：

```go
func Parse(filepath string) (Config, error) {
    switch fileExtension(filepath) {
    case "json":
        return parseJSON(filepath)
    case "yaml":
        return parseYAML(filepath)
    case "toml":
        return parseTOML(filepath)
    default:
        return Config{}, ErrUnknownFileExtension
    }
}
```

Here, we've clearly distinguished the nested function calls from their parent without being overly specific. This allows each nested function call to make sense on its own as well as within the context of the parent. On the other hand, if we had named the `parseJSON` function `json` instead, it couldn't possibly stand on its own. The functionality would become lost in the name, and we would no longer be able to tell whether this function is parsing, creating, or marshalling JSON.

在这里，我们已经清楚地区分了嵌套函数调用与其父函数调用，但并不过分具体。这允许每个嵌套函数调用在其自身以及父级上下文中都有意义。另一方面，如果我们将 `parseJSON` 函数命名为 `json`，它就不可能独立存在。功能将在名称中丢失，我们将不再能够判断此函数是解析、创建还是编组 JSON。

Notice that `fileExtension` is actually a little more specific. However, this is because its functionality is in fact quite specific in nature:

请注意，`fileExtension` 实际上更具体一些。然而，这是因为它的功能实际上非常具体：

```go
func fileExtension(filepath string) string {
    segments := strings.Split(filepath, ".")
    return segments[len(segments)-1]
}
```

This kind of logical progression in our function names—from a high level of abstraction to a lower, more specific one—makes the code easier to follow and read. Consider the alternative: If our highest level of abstraction is too specific, then we'll end up with a name that attempts to cover all bases, like `DetermineFileExtensionAndParseConfigurationFile`. This is horrendously difficult to read; we are trying to be too specific too soon and end up confusing the reader, despite trying to be clear!

我们函数名称中的这种逻辑进展——从高级抽象到低级、更具体的抽象——使代码更容易理解和阅读。考虑替代方案：如果我们的最高抽象级别过于具体，那么我们最终会得到一个试图涵盖所有基础的名称，例如“DetermineFileExtensionAndParseConfigurationFile”。这非常难以阅读；我们试图过早地过于具体，最终使读者感到困惑，尽管我们试图说清楚！

#### Variable Naming
Rather interestingly, the opposite is true for variables. Unlike functions, our variables should be named from more to less specific the deeper we go into nested scopes.

#### 变量命名
有趣的是，变量的情况正好相反。与函数不同的是，我们的变量应该随着嵌套作用域的深入而从多到少命名。

> <em>You shouldn’t name your variables after their types for the same reason you wouldn’t name your pets 'dog' or 'cat'. – Dave Cheney</em>

> <em>你不应该以变量的类型命名你的变量，就像你不会给你的宠物命名“狗”或“猫”一样。 – 戴夫·切尼</em>

Why should our variable names become less specific as we travel deeper into a function's scope? Simply put, as a variable's scope becomes smaller, it becomes increasingly clear for the reader what that variable represents, thereby eliminating the need for specific naming. In the example of the previous function `fileExtension`, we could even shorten the name of the variable `segments` to `s` if we wanted to. The context of the variable is so clear that it's unnecessary to explain it any further with longer variable names. Another good example of this is in nested for loops:

为什么当我们深入到函数的作用域时，我们的变量名称会变得不那么具体？简而言之，随着变量的作用域变小，读者会越来越清楚该变量代表什么，从而消除了对特定命名的需要。在前面函数`fileExtension`的例子中，如果我们愿意，我们甚至可以将变量`segments`的名称缩短为`s`。变量的上下文非常清楚，没有必要用更长的变量名进一步解释它。另一个很好的例子是嵌套的 for 循环：

```go
func PrintBrandsInList(brands []BeerBrand) {
    for _, b := range brands {
        fmt.Println(b)
    }
}
```

In the above example, the scope of the variable `b` is so small that we don't need to spend any additional brain power on remembering what exactly it represents. However, because the scope of `brands` is slightly larger, it helps for it to be more specific. When expanding the variable scope in the function below, this distinction becomes even more apparent:

在上面的例子中，变量‘b’的范围很小，我们不需要花费任何额外的脑力来记住它到底代表什么。但是，由于“品牌”的范围稍大，因此有助于使其更加具体。在下面的函数中扩展变量范围时，这种区别变得更加明显：

```go
func BeerBrandListToBeerList(beerBrands []BeerBrand) []Beer {
    var beerList []Beer
    for _, brand := range beerBrands {
        for _, beer := range brand {
            beerList = append(beerList, beer)
        }
    }
    return beerList
}
```

Great! This function is easy to read. Now, let's apply the opposite (i.e., wrong) logic when naming our variables:

伟大的！这个函数很容易阅读。现在，让我们在命名变量时应用相反的（即错误的）逻辑：

```go
func BeerBrandListToBeerList(b []BeerBrand) []Beer {
    var bl []Beer
    for _, beerBrand := range b {
        for _, beerBrandBeerName := range beerBrand {
            bl = append(bl, beerBrandBeerName)
        }
    }
    return bl
}
```

Even though it's possible to figure out what this function is doing, the excessive brevity of the variable names makes it difficult to follow the logic as we travel deeper. This could very well spiral into full-blown confusion because we're mixing short and long variable names inconsistently.

尽管有可能弄清楚这个函数在做什么，但变量名过于简洁使得我们在深入研究时很难遵循逻辑。这很可能会演变成全面的混乱，因为我们不一致地混合了短变量名和长变量名。

### Cleaning Functions

### 清洁功能

Now that we know some best practices for naming our variables and functions, as well as clarifying our code with comments, let's dive into some specifics of how we can refactor functions to make them cleaner.

现在我们知道了一些命名变量和函数的最佳实践，以及用注释阐明我们的代码，让我们深入研究如何重构函数以使其更清晰的一些细节。

#### Function Length

#### 函数长度

> <em>How small should a function be? Smaller than that! – Robert C. Martin</em>

> <em>一个函数应该有多小？比那个还小！ – 罗伯特 C. 马丁</em>

When writing clean code, our primary goal is to make our code easily digestible. The most effective way to do this is to make our functions as short as possible. It's important to understand that we don't necessarily do this to avoid code duplication. The more important reason is to improve <em>code comprehension</em>.

在编写干净的代码时，我们的主要目标是使我们的代码易于理解。做到这一点最有效的方法是使我们的函数尽可能短。重要的是要了解我们不一定要这样做以避免代码重复。更重要的原因是提高<em>代码理解</em>。

It can help to look at a function's description at a very high level to understand this better:

它可以帮助在非常高的层次上查看函数的描述以更好地理解这一点：

```
fn GetItem:
    - parse json input for order id
    - get user from context
    - check user has appropriate role
    - get order from database
```

By writing short functions (which are typically 5–8 lines in Go), we can create code that reads almost as naturally as our description above:

通过编写短函数（在 Go 中通常为 5-8 行），我们可以创建与上面描述的一样自然的代码：

```go
var (
    NullItem = Item{}
    ErrInsufficientPrivileges = errors.New("user does not have sufficient privileges")
)

func GetItem(ctx context.Context, json []bytes) (Item, error) {
    order, err := NewItemFromJSON(json)
    if err != nil {
        return NullItem, err
    }
    if !GetUserFromContext(ctx).IsAdmin() {
          return NullItem, ErrInsufficientPrivileges
    }
    return db.GetItem(order.ItemID)
}
```

Using smaller functions also eliminates another horrible habit of writing code: indentation hell. <strong>Indentation hell</strong> typically occurs when a chain of `if` statements are carelessly nested in a function. This makes it <em>very</em> difficult for human beings to parse the code and should be eliminated whenever spotted. Indentation hell is particularly common when working with `interface{}` and using type casting:

使用较小的函数还消除了编写代码的另一个可怕习惯：缩进地狱。 <strong>缩进地狱</strong>通常发生在不小心嵌套在函数中的一系列 `if` 语句时。这使得人类<em>非常</em>难以解析代码，一旦发现就应该消除。使用 `interface{}` 和类型转换时，缩进地狱特别常见：

```go
func GetItem(extension string) (Item, error) {
    if refIface, ok := db.ReferenceCache.Get(extension);ok {
        if ref, ok := refIface.(string);ok {
            if itemIface, ok := db.ItemCache.Get(ref);ok {
                if item, ok := itemIface.(Item);ok {
                    if item.Active {
                        return Item, nil
                    } else {
                      return EmptyItem, errors.New("no active item found in cache")
                    }
                } else {
                  return EmptyItem, errors.New("could not cast cache interface to Item")
                }
            } else {
              return EmptyItem, errors.New("extension was not found in cache reference")
            }
        } else {
          return EmptyItem, errors.New("could not cast cache reference interface to Item")
        }
    }
    return EmptyItem, errors.New("reference not found in cache")
}
```

First, indentation hell makes it difficult for other developers to understand the flow of your code. Second, if the logic in our `if` statements expands, it'll become exponentially more difficult to figure out which statement returns what (and to ensure that all paths return some value). Yet another problem is that this deep nesting of conditional statements forces the reader to frequently scroll and keep track of many logical states in their head. It also makes it more difficult to test the code and catch bugs because there are so many different nested possibilities that you have to account for.

首先，缩进地狱使其他开发人员难以理解您的代码流程。其次，如果我们的 `if` 语句中的逻辑扩展，找出哪个语句返回什么（并确保所有路径返回某个值）将变得更加困难。另一个问题是条件语句的这种深度嵌套迫使读者经常滚动并跟踪他们头脑中的许多逻辑状态。它还使测试代码和捕获错误变得更加困难，因为您必须考虑许多不同的嵌套可能性。

Indentation hell can result in reader fatigue if a developer has to constantly parse unwieldy code like the sample above. Naturally, this is something we want to avoid at all costs.

如果开发人员必须像上面的示例那样不断地解析笨拙的代码，那么缩进地狱可能会导致读者疲劳。当然，这是我们想要不惜一切代价避免的事情。

So, how do we clean this function? Fortunately, it's actually quite simple. On our first iteration, we will try to ensure that we are returning an error as soon as possible. Instead of nesting the `if` and `else` statements, we want to "push our code to the left," so to speak. Take a look:

那么，我们如何清理这个函数呢？幸运的是，它实际上非常简单。在我们的第一次迭代中，我们将尝试确保尽快返回错误。可以这么说，我们不想嵌套 `if` 和 `else` 语句，而是想“将我们的代码向左推”。看一看：

```go
func GetItem(extension string) (Item, error) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, errors.New("reference not found in cache")
    }

    ref, ok := refIface.(string)
    if !ok {
        // return cast error on reference
    }

    itemIface, ok := db.ItemCache.Get(ref)
    if !ok {
        // return no item found in cache by reference
    }

    item, ok := itemIface.(Item)
    if !ok {
        // return cast error on item interface
    }

    if !item.Active {
        // return no item active
    }

    return Item, nil
}
```

Once we're done with our first attempt at refactoring the function, we can proceed to split up the function into smaller functions. Here's a good rule of thumb:  If the `value, err :=` pattern is repeated more than once in a function, this is an indication that we can split the logic of our code into smaller pieces:

一旦我们完成了对函数的第一次重构尝试，我们就可以继续将函数拆分为更小的函数。这是一个很好的经验法则：如果 `value, err :=` 模式在一个函数中重复多次，这表明我们可以将代码逻辑拆分为更小的部分：

```go
func GetItem(extension string) (Item, error) {
    ref, ok := getReference(extension)
    if !ok {
        return EmptyItem, ErrReferenceNotFound
    }
    return getItemByReference(ref)
}

func getReference(extension string) (string, bool) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, false
    }
    return refIface.(string)
}

func getItemByReference(reference string) (Item, error) {
    item, ok := getItemFromCache(reference)
    if !item.Active ||!ok {
        return EmptyItem, ErrItemNotFound
    }
    return Item, nil
}

func getItemFromCache(reference string) (Item, bool) {
    if itemIface, ok := db.ItemCache.Get(ref);ok {
        return EmptyItem, false
    }
    return itemIface.(Item), true
}
```

As mentioned previously, indentation hell can make it difficult to test our code. When we split up our `GetItem` function into several helpers, we make it easier to track down bugs when testing our code. Unlike the original version, which consisted of several `if` statements in the same scope, the refactored version of `GetItem` has just two branching paths that we must consider. The helper functions are also short and digestible, making them easier to read.

如前所述，缩进地狱会使测试我们的代码变得困难。当我们将 `GetItem` 函数拆分为多个帮助程序时，我们可以在测试代码时更轻松地跟踪错误。与由同一范围内的多个 if 语句组成的原始版本不同，GetItem 的重构版本只有两个我们必须考虑的分支路径。辅助函数也简短易懂，使它们更易于阅读。

> Note: For production code, one should elaborate on the code even further by returning errors instead of `bool` values. This makes it much easier to understand where the error is originating from. However, as these are just example functions, returning `bool` values will suffice for now. Examples of returning errors more explicitly will be explained in more detail later.

> 注意：对于生产代码，应该通过返回错误而不是 `bool` 值来进一步详细说明代码。这使得更容易理解错误的来源。然而，由于这些只是示例函数，现在返回 `bool` 值就足够了。稍后将更详细地解释更明确地返回错误的示例。

Notice that cleaning the `GetItem` function resulted in more lines of code overall. However, the code itself is now much easier to read. It's layered in an onion-style fashion, where we can ignore "layers" that we aren't interested in and simply peel back the ones that we do want to examine. This makes it easier to understand low-level functionality because we only have to read maybe 3–5 lines at a time.

请注意，清理`GetItem` 函数会导致整体代码行数增加。但是，代码本身现在更容易阅读。它以洋葱式的方式分层，我们可以忽略我们不感兴趣的“层”，而简单地剥离我们想要检查的那些。这使我们更容易理解低级功能，因为我们一次只需要阅读 3-5 行。

This example illustrates that we cannot measure the cleanliness of our code by the number of lines it uses. The first version of the code was certainly much shorter. However, it was <em>artificially</em> short and very difficult to read. In most cases, cleaning code will initially expand the existing codebase in terms of the number of lines. But this is highly preferable to the alternative of having messy, convoluted logic. If you're ever in doubt about this, just consider how you feel about the following function, which does exactly the same thing as our code but only uses two lines:

这个例子说明我们不能通过代码使用的行数来衡量代码的清洁度。代码的第一个版本肯定要短得多。然而，它<em>人为</em>很短，而且很难阅读。在大多数情况下，清理代码最初会在行数方面扩展现有代码库。但这比拥有凌乱、令人费解的逻辑的替代方案更可取。如果您对此有疑问，请考虑您对以下函数的感受，该函数与我们的代码完全相同，但只使用了两行：

```go
func GetItemIfActive(extension string) (Item, error) {
    if refIface,ok := db.ReferenceCache.Get(extension);ok {if ref,ok := refIface.(string);ok { if itemIface,ok := db.ItemCache.Get(ref);ok { if item,ok := itemIface.(Item);ok { if item.Active { return Item,nil }}}}} return EmptyItem, errors.New("reference not found in cache")
}
```

#### Function Signatures 

#### 函数签名

Creating a good function naming structure makes it easier to read and understand the intent of the code. As we saw above, making our functions shorter helps us understand the function's logic. The last part of cleaning our functions involves understanding the context of the function input. With this comes another easy-to-follow rule: <strong>Function signatures should only contain one or two input parameters</strong>. In certain exceptional cases, three can be acceptable, but this is where we should start considering a refactor. Much like the rule that our functions should only be 5–8 lines long, this can seem quite extreme at first. However, I feel that this rule is much easier to justify.

创建一个好的函数命名结构可以更容易阅读和理解代码的意图。正如我们在上面看到的，让我们的函数更短有助于我们理解函数的逻辑。清理函数的最后一部分涉及理解函数输入的上下文。随之而来的是另一个易于遵循的规则：<strong>函数签名应该只包含一个或两个输入参数</strong>。在某些特殊情况下，三个是可以接受的，但这是我们应该开始考虑重构的地方。就像我们的函数应该只有 5-8 行长的规则一样，这乍一看似乎非常极端。但是，我觉得这个规则更容易证明。

Take the following function from [RabbitMQ's introduction tutorial to its Go library](https://www.rabbitmq.com/tutorials/tutorial-one-go.html):

从[RabbitMQ的Go库介绍教程](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)中获取以下函数：

```go
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
```

The function `QueueDeclare` takes six input parameters, which is quite a lot. With some effort, it's possible to understand what this code does thanks to the comments. However, the comments are actually part of the problem—as mentioned earlier, they should be substituted with descriptive code whenever possible. After all, there's nothing preventing us from invoking the `QueueDeclare` function <em>without</em> comments:

函数 QueueDeclare 需要六个输入参数，这是相当多的。通过一些努力，可以通过注释理解这段代码的作用。然而，注释实际上是问题的一部分——正如前面提到的，它们应该尽可能用描述性代码代替。毕竟，没有什么可以阻止我们在没有</em>注释的情况下调用 `QueueDeclare` 函数：

```go
q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
```

Now, without looking at the commented version, try to remember what the fourth and fifth `false` arguments represent. It's impossible, right? You will inevitably forget at some point. This can lead to costly mistakes and bugs that are difficult to correct. The mistakes might even occur through incorrect comments—imagine labeling the wrong input parameter. Correcting this mistake will be unbearably difficult to correct, especially when familiarity with the code has deteriorated over time or was low to begin with. Therefore, it is recommended to replace these input parameters with an 'Options' `struct` instead:

现在，不要看评论版本，试着记住第四和第五个“false”参数代表什么。这是不可能的，对吧？你不可避免地会在某个时候忘记。这可能导致代价高昂的错误和难以纠正的错误。错误甚至可能通过不正确的注释发生——想象一下标记错误的输入参数。纠正这个错误将是难以忍受的，尤其是当对代码的熟悉程度随着时间的推移而恶化或一开始就很低的时候。因此，建议将这些输入参数替换为“选项”`struct`：

```go
type QueueOptions struct {
    Name string
    Durable bool
    DeleteOnExit bool
    Exclusive bool
    NoWait bool
    Arguments []interface{}
}

q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
    Durable: false,
    DeleteOnExit: false,
    Exclusive: false,
    NoWait: false,
    Arguments: nil,
})
```

This solves two problems: misusing comments, and accidentally labeling the variables incorrectly. Of course, we can still confuse properties with the wrong value, but in these cases, it will be much easier to determine where our mistake lies within the code. The ordering of the properties also doesn't matter anymore, so incorrectly ordering the input values is no longer a concern. The last added bonus of this technique is that we can use our `QueueOptions` struct to infer the default values of our function's input parameters. When structures in Go are declared, all properties are initialised to their default value. This means that our `QueueDeclare` option can actually be invoked in the following way:

这解决了两个问题：误用注释，以及意外地错误地标记变量。当然，我们仍然可以将属性与错误的值混淆，但在这些情况下，确定我们的错误在代码中的位置会容易得多。属性的排序也不再重要，因此输入值的错误排序不再是问题。这项技术的最后一个额外好处是我们可以使用我们的 `QueueOptions` 结构来推断函数输入参数的默认值。当 Go 中的结构被声明时，所有的属性都被初始化为它们的默认值。这意味着我们的 `QueueDeclare` 选项实际上可以通过以下方式调用：

```go
q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
})
```

The rest of the values are initialised to their default value of `false` (except for `Arguments`, which as an interface has a default value of `nil`). Not only are we much safer with this approach, but we are also much clearer with our intentions. In this case, we could actually write less code. This is an all-around win for everyone on the project.

其余的值被初始化为它们的默认值 `false`（除了 `Arguments`，它作为一个接口的默认值是 `nil`）。使用这种方法不仅我们更安全，而且我们的意图也更加清晰。在这种情况下，我们实际上可以编写更少的代码。这对项目中的每个人来说都是一个全面的胜利。

One final note on this: It's not always possible to change a function's signature. In this case, for example, we don't actually have control over our `QueueDeclare` function signature because it's from the RabbitMQ library. It's not our code, so we can't change it. However, we can wrap these functions to suit our purposes:

关于此的最后一个说明：并非总是可以更改函数的签名。例如，在这种情况下，我们实际上无法控制我们的 `QueueDeclare` 函数签名，因为它来自 RabbitMQ 库。这不是我们的代码，所以我们不能改变它。但是，我们可以包装这些函数以满足我们的目的：

```go
type RMQChannel struct {
    channel *amqp.Channel
}

func (rmqch *RMQChannel) QueueDeclare(opts QueueOptions) (Queue, error) {
    return rmqch.channel.QueueDeclare(
        opts.Name,
        opts.Durable,
        opts.DeleteOnExit,
        opts.Exclusive,
        opts.NoWait,
        opts.Arguments,
    )
}
```

Basically, we create a new structure named `RMQChannel` that contains the `amqp.Channel` type, which has the `QueueDeclare` method. We then create our own version of this method, which essentially just calls the old version of the RabbitMQ library function. Our new method has all the advantages described before, and we achieved this without actually having to change any of the code in the RabbitMQ library.

基本上，我们创建了一个名为 `RMQChannel` 的新结构，它包含 `amqp.Channel` 类型，它具有 `QueueDeclare` 方法。然后我们创建我们自己的这个方法版本，它本质上只是调用旧版本的 RabbitMQ 库函数。我们的新方法具有之前描述的所有优点，而且我们无需实际更改 RabbitMQ 库中的任何代码即可实现这一点。

We'll use this idea of wrapping functions to introduce more clean and safe code later when discussing `interface{}`.

稍后在讨论 `interface{}` 时，我们将使用这种包装函数的想法来引入更干净和安全的代码。

### Variable Scope
Now, let's take a step back and revisit the idea of writing smaller functions. This has another nice side effect that we didn't cover in the previous chapter: Writing smaller functions can typically eliminate reliance on mutable variables that leak into the global scope.

### 变量范围
现在，让我们退后一步，重新审视编写更小的函数的想法。这有另一个我们在前一章中没有涉及的很好的副作用：编写更小的函数通常可以消除对泄漏到全局作用域的可变变量的依赖。

Global variables are problematic and don't belong in clean code; they make it very difficult for programmers to understand the current state of a variable. If a variable is global and mutable, then by definition, its value can be changed by any part of the codebase. At no point can you guarantee that this variable is going to be a specific value... And that's a headache for everyone. This is yet another example of a trivial problem that's exacerbated when the codebase expands.

全局变量有问题，不属于干净的代码；它们使程序员很难理解变量的当前状态。如果一个变量是全局的和可变的，那么根据定义，它的值可以被代码库的任何部分改变。您在任何时候都不能保证这个变量将是一个特定的值......这对每个人来说都是一个头疼的问题。这是另一个在代码库扩展时会加剧的小问题的例子。

Let's look at a short example of how non-global variables with a large scope can cause problems. These variables also introduce the issue of <strong>variable shadowing</strong>, as demonstrated in the code taken from an article titled [Golang scope issue](https://idiallo.com/blog/golang-scopes):

让我们看一个简短的例子，说明具有大范围的非全局变量如何导致问题。这些变量还引入了<strong>变量阴影</strong>的问题，如标题为 [Golang 范围问题](https://idiallo.com/blog/golang-scopes) 的文章中的代码所示：

```go
func doComplex() (string, error) {
    return "Success", nil
}

func main() {
    var val string
    num := 32

    switch num {
    case 16:
    // do nothing
    case 32:
        val, err := doComplex()
        if err != nil {
            panic(err)
        }
        if val == "" {
            // do something else
        }
    case 64:
        // do nothing
    }

    fmt.Println(val)
}
```

What's the problem with this code? From a quick skim, it seems the `var val string` value should be printed out as `Success` by the end of the `main` function. Unfortunately, this is not the case. The reason for this lies in the following line:

这段代码有什么问题？快速浏览一下，似乎在 `main` 函数结束时 `var val string` 值应该打印为 `Success`。不幸的是，这种情况并非如此。原因在于以下几行：

```go
val, err := doComplex()
```

This declares a new variable `val` in the switch's `case 32` scope and has nothing to do with the variable declared in the first line of `main`. Of course, it can be argued that Go syntax is a little tricky, which I don't necessarily disagree with, but there is a much worse issue at hand. The declaration of `var val string` as a mutable, largely scoped variable is completely unnecessary. If we do a <strong>very</strong> simple refactor, we will no longer have this issue:

这在 switch 的 `case 32` 作用域中声明了一个新变量 `val`，与 `main` 第一行中声明的变量无关。当然，可以说 Go 语法有点棘手，我不一定不同意，但手头有一个更糟糕的问题。将 `var val string` 声明为一个可变的、范围很大的变量是完全没有必要的。如果我们进行<strong>非常</strong>的重构，我们将不再有这个问题：

```go
func getStringResult(num int) (string, error) {
    switch num {
    case 16:
    // do nothing
    case 32:
       return doComplex()
    case 64:
        // do nothing
    }
    return "", nil
}

func main() {
    val, err := getStringResult(32)
    if err != nil {
        panic(err)
    }
    if val == "" {
        // do something else
    }
    fmt.Println(val)
}
```

After our refactor, `val` is no longer modified, and the scope has been reduced. Again, keep in mind that these functions are very simple. Once this kind of code style becomes a part of larger, more complex systems, it can be impossible to figure out why errors are occurring. We don't want this to happen—not only because we generally dislike software errors but also because it's disrespectful to our colleagues, and ourselves; we are potentially wasting each other's time having to debug this type of code. Developers need to take responsibility for their own code rather than blaming these issues on the variable declaration syntax of a particular language like Go.

在我们重构之后，`val` 不再被修改，范围也缩小了。同样，请记住，这些功能非常简单。一旦这种代码风格成为更大、更复杂系统的一部分，就不可能弄清楚为什么会发生错误。我们不希望这种情况发生——不仅因为我们通常不喜欢软件错误，还因为它不尊重我们的同事和我们自己；我们可能会浪费彼此的时间来调试这种类型的代码。开发人员需要对自己的代码负责，而不是将这些问题归咎于 Go 等特定语言的变量声明语法。

On a side note, if the `// do something else` part is another attempt to mutate the `val` variable, we should extract that logic out as its own self-contained function, as well as the previous part of it. This way, instead of expanding the mutable scope of our variables, we can just return a new value:

附带说明一下，如果 `// do something else` 部分是对 `val` 变量进行变异的另一种尝试，我们应该将该逻辑作为它自己的自包含函数以及它的前一部分提取出来。这样，我们可以只返回一个新值，而不是扩展变量的可变范围：

```go
func getVal(num int) (string, error) {
    val, err := getStringResult(num)
    if err != nil {
        return "", err
    }
    if val == "" {
        return NewValue() // pretend function
    }
}

func main() {
    val, err := getVal(32)
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}
```

### Variable Declaration 

### 变量声明

Other than avoiding issues with variable scope and mutability, we can also improve readability by declaring variables as close to their usage as possible. In C programming, it's common to see the following approach to declaring variables:

除了避免变量作用域和可变性的问题外，我们还可以通过尽可能接近其用法来声明变量来提高可读性。在 C 编程中，通常会看到以下声明变量的方法：

```go
func main() {
  var err error
  var items []Item
  var sender, receiver chan Item
  
  items = store.GetItems()
  sender = make(chan Item)
  receiver = make(chan Item)
  
  for _, item := range items {
    ...
  }
}
```

This suffers from the same symptom as described in our discussion of variable scope. Even though these variables might not actually be reassigned at any point, this kind of coding style keeps the readers on their toes, in all the wrong ways. Much like computer memory, our brain's short-term memory has a limited capacity. Having to keep track of which variables are mutable and whether or not a particular fragment of code will mutate them makes it more difficult to understand what the code is doing. Figuring out the eventually returned value can be a nightmare. Therefore, to makes this easier for our readers (and our future selves), it's recommended that you declare variables as close to their usage as possible:

这与我们在讨论变量范围时所描述的症状相同。尽管这些变量实际上可能不会在任何时候重新分配，但这种编码风格以所有错误的方式让读者保持警觉。就像计算机记忆一样，我们大脑的短期记忆容量有限。必须跟踪哪些变量是可变的，以及特定的代码片段是否会改变它们，这使得理解代码在做什么变得更加困难。找出最终返回的值可能是一场噩梦。因此，为了让我们的读者（以及我们未来的自己）更容易做到这一点，建议您声明变量尽可能接近它们的用法：

```go
func main() {
    var sender chan Item
    sender = make(chan Item)

    go func() {
        for {
            select {
            case item := <-sender:
                // do something
            }
        }
    }()
}
```

However, we can do even better by invoking the function directly after its declaration. This makes it much clearer that the function logic is associated with the declared variable:

但是，我们可以通过在声明之后直接调用函数来做得更好。这使得函数逻辑与声明的变量相关联变得更加清晰：

```go
func main() {
  sender := func() chan Item {
    channel := make(chan Item)
    go func() {
      for {
        select { ... }
      }
    }()
    return channel
  }
}
```

And coming full circle, we can move the anonymous function to make it a named function instead:

又回到原点，我们可以移动匿名函数使其成为命名函数：

```go
func main() {
  sender := NewSenderChannel()
}

func NewSenderChannel() chan Item {
  channel := make(chan Item)
  go func() {
    for {
      select { ... }
    }
  }()
  return channel
}
```

It is still clear that we are declaring a variable, and the logic associated with the returned channel is simple, unlike in the first example. This makes it easier to traverse the code and understand the role of each variable.

仍然很明显我们声明了一个变量，与返回的通道关联的逻辑很简单，与第一个示例不同。这样可以更轻松地遍历代码并了解每个变量的作用。

Of course, this doesn't actually prevent us from mutating our `sender` variable. There is nothing that we can do about this, as there is no way of declaring a `const struct` or `static` variables in Go. This means that we'll have to restrain ourselves from modifying this variable at a later point in the code.

当然，这实际上并不能阻止我们改变我们的 `sender` 变量。我们对此无能为力，因为在 Go 中无法声明 `const struct` 或 `static` 变量。这意味着我们将不得不限制自己在代码的稍后点修改这个变量。

> NOTE: The keyword `const` does exist but is limited in use to primitive types only.

> 注意：关键字`const` 确实存在，但仅限于用于原始类型。

One way of getting around this can at least limit the mutability of a variable to the package level. The trick involves creating a structure with the variable as a private property. This private property is thenceforth only accessible through other methods provided by this wrapping structure. Expanding on our channel example, this would look something like the following:

解决这个问题的一种方法至少可以将变量的可变性限制在包级别。诀窍涉及创建一个将变量作为私有属性的结构。此私有属性此后只能通过此包装结构提供的其他方法访问。扩展我们的频道示例，这将类似于以下内容：

```go
type Sender struct {
  sender chan Item
}

func NewSender() *Sender {
  return &Sender{
    sender: NewSenderChannel(),
  }
}

func (s *Sender) Send(item Item) {
  s.sender <- item
}
```

We have now ensured that the `sender` property of our `Sender` struct is never mutated—at least not from outside of the package. As of writing this document, this is the only way of creating publicly immutable non-primitive variables. It's a little verbose, but it's truly worth the effort to ensure that we don't end up with strange bugs resulting from accidental variable modification.

我们现在已经确保我们的 `Sender` 结构的 `sender` 属性永远不会发生变化——至少不会从包的外部发生变化。在撰写本文档时，这是创建公开不可变的非原始变量的唯一方法。这有点冗长，但确实值得努力确保我们不会因意外的变量修改而导致奇怪的错误。

```go
func main() {
  sender := NewSender()
  sender.Send(&Item{})
}
```

Looking at the example above, it's clear how this also simplifies the usage of our package. This way of hiding the implementation is beneficial not only for the maintainers of the package but also for the users. Now, when initialising and using the `Sender` structure, there is no concern over its implementation. This opens up for a much looser architecture. Because our users aren't concerned with the implementation, we are free to change it at any point, since we have reduced the point of contact that users have with the package. If we no longer wish to use a channel implementation in our package, we can easily change this without breaking the usage of the `Send` method (as long as we adhere to its current function signature).

看看上面的例子，很明显这也简化了我们包的使用。这种隐藏实现的方式不仅有利于包的维护者，也有利于用户。现在，在初始化和使用 `Sender` 结构时，无需担心其实现。这为更松散的架构开辟了道路。因为我们的用户不关心实现，我们可以随时更改它，因为我们减少了用户与包的接触点。如果我们不再希望在我们的包中使用通道实现，我们可以在不破坏 `Send` 方法的使用的情况下轻松更改它（只要我们坚持其当前的函数签名）。

>  NOTE: There is a fantastic explanation of how to handle the abstraction in client libraries, taken from the talk [AWS re:Invent 2017: Embracing Change without Breaking the World (DEV319)](https://www.youtube.com/watch?v=kJq81Y7OEx4).

> 注意：从演讲 [AWS re:Invent 2017: Embracing Change without Breaking the World (DEV319)](https://www.youtube.com/看？v=kJq81Y7OEx4)。

## Clean Go

## 干净利落

This section focuses less on the generic aspects of writing clean Go code and more on the specifics, with an emphasis on the underlying clean code principles.

本节较少关注编写干净的 Go 代码的通用方面，而更多地关注细节，重点是底层的干净代码原则。

### Return Values

### 返回值

#### Returning Defined Errors

#### 返回定义的错误

We'll start things off nice and easy by describing a cleaner way to return errors. As we discussed earlier, our main goal with writing clean code is to ensure readability, testability, and maintainability of the codebase. The technique for returning errors that we'll discuss here will achieve all three of those goals with very little effort.

我们将通过描述一种更简洁的返回错误的方式来开始事情。正如我们之前所讨论的，我们编写干净代码的主要目标是确保代码库的可读性、可测试性和可维护性。我们将在这里讨论的返回错误的技术将毫不费力地实现所有这三个目标。

Let's consider the normal way to return a custom error. This is a hypothetical example taken from a thread-safe map implementation that we've named `Store`:

让我们考虑返回自定义错误的正常方法。这是一个取自我们命名为“Store”的线程安全映射实现的假设示例：

```go
package smelly

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return Item{}, errors.New("item could not be found in the store")
    }
    return item, nil
}
```

There is nothing inherently smelly about this function when we consider it in isolation. We look into the `items` map of our `Store` struct to see if we already have an item with the given `id`. If we do, we return it; otherwise, we return an error. Pretty standard. So, what is the issue with returning custom errors as string values? Well, let's look at what happens when we use this function inside another package:

当我们孤立地考虑它时，这个函数本身并没有什么臭味。我们查看“Store”结构的“items”映射，看看我们是否已经有一个带有给定“id”的项目。如果我们这样做，我们将其退回；否则，我们返回一个错误。很标准。那么，将自定义错误作为字符串值返回有什么问题？好吧，让我们看看当我们在另一个包中使用这个函数时会发生什么：

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := smelly.GetItem("123")
    if err != nil {
        if err.Error() == "item could not be found in the store" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, errr.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(item)
}
```

This is actually not too bad. However, there is one glaring problem: An error in Go is simply an `interface` that implements a function (`Error()`) returning a string; thus, we are now hardcoding the expected error code into our codebase, which isn't ideal. This hardcoded string is known as a <strong>magic string</strong>. And its main problem is flexibility: If at some point we decide to change the string value used to represent an error, our code will break (softly) unless we update it in possibly many different places. Our code is tightly coupled—it relies on that specific magic string and the assumption that it will never change as the codebase grows.

这实际上还不算太糟糕。然而，有一个明显的问题：Go 中的错误只是一个实现函数（`Error()`）的`接口`，返回一个字符串；因此，我们现在将预期的错误代码硬编码到我们的代码库中，这并不理想。这个硬编码的字符串被称为<strong>魔法字符串</strong>。它的主要问题是灵活性：如果在某个时候我们决定更改用于表示错误的字符串值，我们的代码将（软）中断，除非我们在许多不同的地方更新它。我们的代码是紧密耦合的——它依赖于那个特定的魔法字符串，并且假设它永远不会随着代码库的增长而改变。

An even worse situation would arise if a client were to use our package in their own code. Imagine that we decided to update our package and changed the string that represents an error—the client's software would now suddenly break. This is quite obviously something that we want to avoid. Fortunately, the fix is very simple:

如果客户在他们自己的代码中使用我们的包，则会出现更糟糕的情况。想象一下，我们决定更新我们的包并更改表示错误的字符串——客户端的软件现在会突然中断。这显然是我们想要避免的。幸运的是，修复非常简单：

```go
package clean

var (
    NullItem = Item{}

    ErrItemNotFound = errors.New("item could not be found in the store")
)

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, ErrItemNotFound
    }
    return item, nil
}
```

By simply representing the error as a variable (`ErrItemNotFound`), we've ensured that anyone using this package can check against the variable rather than the actual string that it returns:

通过简单地将错误表示为一个变量（`ErrItemNotFound`），我们确保使用这个包的任何人都可以检查变量而不是它返回的实际字符串：

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err, clean.ErrItemNotFound) {
           http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(item)
}
```

This feels much nicer and is also much safer. Some would even say that it's easier to read as well. In the case of a more verbose error message, it certainly would be preferable for a developer to simply read `ErrItemNotFound` rather than a novel on why a certain error has been returned.

这感觉好多了，也更安全。有些人甚至会说它也更容易阅读。在更详细的错误消息的情况下，对于开发人员来说，简单地阅读“ErrItemNotFound”，而不是关于为什么返回某个错误的小说当然更可取。

This approach is not limited to errors and can be used for other returned values. As an example, we are also returning a `NullItem` instead of `Item{}` as we did before. There are many different scenarios in which it might be preferable to return a defined object, rather than initialising it on return.

这种方法不仅限于错误，还可用于其他返回值。例如，我们还返回了一个 `NullItem` 而不是我们之前所做的 `Item{}`。在许多不同的场景中，最好返回一个定义的对象，而不是在返回时初始化它。

Returning default `NullItem` values like we did in the previous examples can also be safer in certain cases. As an example, a user of our package could forget to check for errors and end up initialising a variable that points to an empty struct containing a default value of `nil` as one or more property values. When attempting to access this `nil` value later in the code, the client software would panic. However, when we return our custom default value instead, we can ensure that all values that would otherwise default to `nil` are initialised. Thus, we'd ensure that we do not cause panics in our users' software.

在某些情况下，像我们在前面的示例中所做的那样返回默认的 `NullItem` 值也可以更安全。例如，我们包的用户可能忘记检查错误并最终初始化一个变量，该变量指向一个包含默认值“nil”的空结构作为一个或多个属性值。当稍后在代码中尝试访问这个 `nil` 值时，客户端软件会崩溃。但是，当我们返回自定义默认值时，我们可以确保初始化所有默认为“nil”的值。因此，我们将确保不会在用户的软件中引起恐慌。

This also benefits us. Consider this: If we wanted to achieve the same safety without returning a default value, we would have to change our code everywhere we return this type of empty value. However, with our default value approach, we now only have to change our code in a single place:

这也让我们受益匪浅。考虑一下：如果我们想在不返回默认值的情况下实现相同的安全性，则必须在返回此类空值的任何地方更改代码。然而，使用我们的默认值方法，我们现在只需要在一个地方更改我们的代码：

```go
var NullItem = Item{
    itemMap: map[string]Item{},
}
```

> NOTE: In many scenarios, invoking a panic will actually be preferable to indicate that there is an error check missing.

> 注意：在许多情况下，调用恐慌实际上更可取地表明缺少错误检查。

> NOTE: Every interface property in Go has a default value of `nil`. This means that this is useful for any struct that has an interface property. This is also true for structs that contain channels, maps, and slices, which could potentially also have a `nil` value.

> 注意：Go 中的每个接口属性都有一个默认值“nil”。这意味着这对于具有接口属性的任何结构都很有用。对于包含通道、映射和切片的结构也是如此，它们也可能具有 nil 值。

#### Returning Dynamic Errors
There are certainly some scenarios where returning an error variable might not actually be viable. In cases where the information in customised errors is dynamic, if we want to describe error events more specifically, we can no longer define and return our static errors. Here's an example:

#### 返回动态错误
在某些情况下，返回错误变量可能实际上并不可行。在自定义错误中的信息是动态的情况下，如果我们想更具体地描述错误事件，我们不能再定义和返回我们的静态错误。下面是一个例子：

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, fmt.Errorf("Could not find item with ID: %s", id)
    }
    return item, nil
}
```

So, what to do? There is no well-defined or standard method for handling and returning these kinds of dynamic errors. My personal preference is to return a new interface, with a bit of added functionality:

那么该怎么办？没有明确定义或标准的方法来处理和返回这些类型的动态错误。我个人的偏好是返回一个新界面，增加一些功能：

```go
type ErrorDetails interface {
    Error() string
    Type() string
}

type errDetails struct {
    errtype error
    details interface{}
}

func NewErrorDetails(err error, details ...interface{}) ErrorDetails {
    return &errDetails{
        errtype: err,
        details: details,
    }
}

func (err *errDetails) Error() string {
    return fmt.Sprintf("%v: %v", err.errtype, err.details)
}

func (err *errDetails) Type() error {
    return err.errtype
}
```

This new data structure still works as our standard error. We can still compare it to `nil` since it's an interface implementation, and we can still call `.Error()` on it, so it won't break any existing implementations. However, the advantage is that we can now check our error type as we could previously, despite our error now containing the <em>dynamic</em> details:

这个新的数据结构仍然作为我们的标准错误。我们仍然可以将它与 `nil` 进行比较，因为它是一个接口实现，我们仍然可以对其调用 `.Error()`，因此它不会破坏任何现有的实现。然而，优点是我们现在可以像以前一样检查我们的错误类型，尽管我们的错误现在包含<em>动态</em>详细信息：

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, NewErrorDetails(
            ErrItemNotFound,
            fmt.Sprintf("could not find item with id: %s", id))
    }
    return item, nil
}
```

And our HTTP handler function can then be refactored to check for a specific error again:

然后可以重构我们的 HTTP 处理程序函数以再次检查特定错误：

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err.Type(), clean.ErrItemNotFound) {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(item)
}
```

### Nil Values 

### 零值

A controversial aspect of Go is the addition of `nil`. This value corresponds to the value `NULL` in C and is essentially an uninitialised pointer. We've already seen some of the problems that `nil` can cause, but to sum up: Things break when you try to access methods or properties of a `nil` value. Thus, it's recommended to avoid returning a  `nil` value when possible. This way, the users of our code are less likely to accidentally access `nil` values.

Go 的一个有争议的方面是添加了 `nil`。该值对应于 C 中的值“NULL”，本质上是一个未初始化的指针。我们已经看到了 nil 可能导致的一些问题，但总结一下：当您尝试访问 nil 值的方法或属性时，事情会中断。因此，建议尽可能避免返回 `nil` 值。这样，我们代码的用户就不太可能意外访问“nil”值。

There are other scenarios in which it is common to find `nil` values that can cause some unnecessary pain. An example of this is incorrectly initialising a `struct` (as in the example below), which can lead to it containing `nil` properties. If accessed, those `nil`s will cause a panic.

在其他情况下，通常会发现“nil”值可能会导致一些不必要的痛苦。这方面的一个例子是错误地初始化了一个 `struct`（如下例所示），这可能导致它包含 `nil` 属性。如果访问，那些 `nil` 会引起恐慌。

```go
type App struct {
    Cache *KVCache
}

type KVCache struct {
  mtx sync.RWMutex
    store map[string]string
}

func (cache *KVCache) Add(key, value string) {
  cache.mtx.Lock()
  defer cache.mtx.Unlock()
  
    cache.store[key] = value
}
```

This code is absolutely fine. However, the danger is that our `App` can be initialised incorrectly, without initialising the `Cache` property within. Should the following code be invoked, our application will panic:

这段代码绝对没问题。然而，危险在于我们的 `App` 可能会被错误地初始化，而没有初始化其中的 `Cache` 属性。如果调用以下代码，我们的应用程序将发生恐慌：

```go
     app := App{}
    app.Cache.Add("panic", "now")
```

The `Cache` property has never been initialised and is therefore a `nil` pointer. Thus, invoking the `Add` method like we did here will cause a panic, with the following message:

`Cache` 属性从未被初始化，因此是一个 `nil` 指针。因此，像我们在此处所做的那样调用 `Add` 方法将导致恐慌，并显示以下消息：

> panic: runtime error: invalid memory address or nil pointer dereference

> 恐慌：运行时错误：无效的内存地址或空指针取消引用

Instead, we can turn the `Cache` property of our `App` structure into a private property and create a getter-like method to access it. This gives us more control over what we are returning; specifically, it ensures that we aren't returning a `nil` value:

相反，我们可以将“App”结构的“Cache”属性转换为私有属性，并创建一个类似 getter 的方法来访问它。这使我们可以更好地控制返回的内容；具体来说，它确保我们不会返回一个 `nil` 值：

```go
type App struct {
    cache *KVCache
}

func (app *App) Cache() *KVCache {
    if app.cache == nil {
        app.cache = NewKVCache()
    }
    return app.cache
}
```

The code that previously panicked will now be refactored to the following:

之前恐慌的代码现在将重构为以下内容：

```go
app := App{}
app.Cache().Add("panic", "now")
```

This ensures that users of our package don't have to worry about the implementation and whether they're using our package in an unsafe manner. All they need to worry about is writing their own clean code.

这确保了我们包的用户不必担心实现以及他们是否以不安全的方式使用我们的包。他们需要担心的只是编写自己的干净代码。

> NOTE: There are other methods of achieving a similarly safe outcome. However, I believe this is the most straightforward approach.

> 注意：还有其他方法可以获得类似的安全结果。但是，我相信这是最直接的方法。

### Pointers in Go
Pointers in Go are a rather extensive topic. They're a very big part of working with the language—so much so that it is essentially impossible to write Go without some knowledge of pointers and their workings in the language. Therefore, it is important to understand how to use pointers without adding unnecessary complexity (and thereby keeping your codebase clean). Note that we will not review the details of how pointers are implemented in Go. Instead, we will focus on the quirks of Go pointers and how we can handle them.

### Go 中的指针
Go 中的指针是一个相当广泛的主题。它们是语言工作的重要组成部分——如此之多以至于如果没有指针及其在语言中的工作原理的一些知识，基本上不可能编写 Go。因此，了解如何使用指针而不增加不必要的复杂性（从而保持代码库干净）非常重要。请注意，我们不会回顾在 Go 中如何实现指针的细节。相反，我们将专注于 Go 指针的怪癖以及我们如何处理它们。

Pointers add complexity to code. If we aren't cautious, incorrectly using pointers can introduce nasty side effects or bugs that are particularly difficult to debug. By sticking to the basic principles of writing clean code that we covered in the first part of this document, we can at least reduce the chances of introducing unnecessary complexity to our code.

指针增加了代码的复杂性。如果我们不小心，错误地使用指针可能会引入令人讨厌的副作用或特别难以调试的错误。通过坚持我们在本文档第一部分中介绍的编写干净代码的基本原则，我们至少可以减少将不必要的复杂性引入到我们的代码中的机会。

#### Pointer Mutability
We've already looked at the problem of mutability in the context of globally or largely scoped variables. However, mutability is not necessarily always a bad thing, and I am by no means an advocate for writing 100% pure functional programs. Mutability is a powerful tool, but we should really only ever use it when it's necessary. Let's have a look at a code example illustrating why:

#### 指针可变性
我们已经在全局变量或大范围变量的上下文中研究了可变性问题。然而，可变性并不一定总是坏事，我绝不提倡编写 100% 纯函数式程序。可变性是一个强大的工具，但我们真的应该只在必要时使用它。让我们看一个说明原因的代码示例：

```go
func (store *UserStore) Insert(user *User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = user
    return nil
}

func (store *UserStore) userExists(id int64) bool {
    _, ok := store.users[id]
    return ok
}
```

At first glance, this doesn't seem too bad. In fact, it might even seem like a rather simple insert function for a common list structure. We accept a pointer as input, and if no other users with this `id` exist, then we insert the provided user pointer into our list. Then, we use this functionality in our public API for creating new users:

乍一看，这似乎还不错。事实上，对于公共列表结构，它甚至可能看起来像是一个相当简单的插入函数。我们接受一个指针作为输入，如果没有其他具有这个 `id` 的用户存在，那么我们将提供的用户指针插入到我们的列表中。然后，我们在公共 API 中使用此功能来创建新用户：

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {
    user, err := parseUserFromRequest(r)
    if err != nil {
        http.Error(w, err, http.StatusBadRequest)
        return
    }
    if err := insertUser(w, user);err != nil {
      http.Error(w, err, http.StatusInternalServerError)
      return
    }
}

func insertUser(w http.ResponseWriter, user User) error {
      if err := store.Insert(user);err != nil {
        return err
    }
      user.Password = ""
      return json.NewEncoder(w).Encode(user)
}
```

Once again, at first glance, everything looks fine. We parse the user from the received request and insert the user struct into our store. Once we have successfully inserted our user into the store, we then set the password to be an empty string before returning the user as a JSON object to our client. This is all quite common practice, typically when returning a user object whose password has been hashed, since we don't want to return the hashed password.

再一次，乍一看，一切看起来都很好。我们从收到的请求中解析用户并将用户结构插入我们的商店。一旦我们成功地将我们的用户插入到商店中，我们就会在将用户作为 JSON 对象返回给我们的客户端之前将密码设置为一个空字符串。这是很常见的做法，通常在返回密码已散列的用户对象时，因为我们不想返回散列的密码。

However, imagine that we are using an in-memory store based on a `map`. This code will produce some unexpected results. If we check our user store, we'll see that the change we made to the users password in the HTTP handler function also affected the object in our store. This is because the pointer address returned by `parseUserFromRequest` is what we populated our store with, rather than an actual value. Therefore, when making changes to the dereferenced password value, we end up changing the value of the object we are pointing to in our store.

但是，想象一下我们正在使用基于 `map` 的内存存储。这段代码会产生一些意想不到的结果。如果我们检查我们的用户存储，我们会看到我们在 HTTP 处理函数中对用户密码所做的更改也影响了我们存储中的对象。这是因为 `parseUserFromRequest` 返回的指针地址是我们填充我们存储的，而不是实际值。因此，当对解除引用的密码值进行更改时，我们最终会更改指向我们存储中的对象的值。

This is a great example of why both mutability and variable scope can cause some serious issues and bugs when used incorrectly. When passing pointers as an input parameter of a function, we are expanding the scope of the variable whose data is being pointed to. Even more worrying is the fact that we are expanding the scope to an undefined level. We are *almost* expanding the scope of the variable to the global level. As demonstrated by the above example, this can lead to disastrous bugs that are particularly difficult to find and eradicate.

这是一个很好的例子，说明了为什么可变性和变量范围在使用不当时会导致一些严重的问题和错误。当将指针作为函数的输入参数传递时，我们正在扩展其数据所指向的变量的范围。更令人担忧的是，我们正在将范围扩大到一个未定义的水平。我们*几乎*将变量的范围扩展到全局级别。如上例所示，这可能会导致特别难以发现和消除的灾难性错误。

Fortunately, the fix for this is rather simple:

幸运的是，解决这个问题很简单：

```go
func (store *UserStore) Insert(user User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = &user
    return nil
}
```

Instead of passing a pointer to a `User` struct, we are now passing in a copy of a `User`. We are still storing a pointer to our store; however, instead of storing the pointer from outside of the function, we are storing the pointer to the copied value, whose scope is inside the function. This fixes the immediate problem but might still cause issues further down the line if we aren't careful. Consider this code:

我们现在不是传递一个指向 `User` 结构的指针，而是传入一个 `User` 的副本。我们仍然存储一个指向我们商店的指针；然而，我们不是从函数外部存储指针，而是存储指向复制值的指针，其作用域在函数内部。这解决了眼前的问题，但如果我们不小心，可能仍然会导致问题进一步恶化。考虑这个代码：

```go
func (store *UserStore) Get(id int64) (*User, error) {
    user, ok := store.users[id]
    if !ok {
        return EmptyUser, ErrUserNotFound
    }
    return store.users[id], nil
}
```

Again, this is a very standard implementation of a getter function for our store. However, it's still bad code because we are once again expanding the scope of our pointer, which may end up causing unexpected side effects. When returning the actual pointer value, which we are storing in our user store, we are essentially giving other parts of our application the ability to change our store values. This is bound to cause confusion. Our store should be the only entity allowed to make changes to its values. The easiest fix for this is to return a value of `User` rather than returning a pointer.

同样，这是我们商店的 getter 函数的非常标准的实现。然而，它仍然是糟糕的代码，因为我们再次扩大了我们的指针的范围，这可能最终导致意想不到的副作用。当返回我们存储在用户存储中的实际指针值时，我们本质上是让应用程序的其他部分能够更改我们的存储值。这势必引起混乱。我们的商店应该是唯一允许更改其值的实体。最简单的解决方法是返回一个 `User` 的值而不是返回一个指针。

> NOTE: Consider the case where our application uses multiple threads. In this scenario, passing pointers to the same memory location can also potentially result in a race condition. In other words, we aren't only potentially corrupting our data—we could also cause a panic from a data race. 

> 注意：考虑我们的应用程序使用多线程的情况。在这种情况下，将指针传递到同一内存位置也可能导致竞争条件。换句话说，我们不仅可能破坏我们的数据——我们还可能因数据竞争而引起恐慌。

Please keep in mind that there is intrinsically nothing wrong with returning pointers. However, the expanded scope of variables (and the number of owners pointing to those variables) is the most important consideration when working with pointers. This is what categorises our previous example as a smelly operation. This is also why common Go constructors are absolutely fine:

请记住，返回指针本质上没有任何问题。然而，在使用指针时，变量的扩展范围（以及指向这些变量的所有者的数量）是最重要的考虑因素。这就是将我们之前的示例归类为臭操作的原因。这也是常见的 Go 构造函数绝对没问题的原因：

```go
func AddName(user *User, name string) {
    user.Name = name
}
```

This is *okay* because the variable scope, which is defined by whoever invokes the function, remains the same after the function returns. Combined with the fact that the function invoker remains the sole owner of the variable, this means that the pointer cannot be manipulated in an unexpected manner.

这是*好的*，因为由调用函数的任何人定义的变量范围在函数返回后保持不变。结合函数调用者仍然是变量的唯一所有者这一事实，这意味着不能以意外的方式操作指针。

### Closures Are Function Pointers

### 闭包是函数指针

Before we get into the next topic of using interfaces in Go, I would like to introduce a common alternative. It's what C programmers know as "function pointers" and what most other programming languages call <strong>closures</strong>. A closure is simply an input parameter like any other, except it represents (points to) a function that can be invoked. In JavaScript, it's quite common to use closures as callbacks, which are just functions that are invoked after some asynchronous operation has finished. In Go, we don't really have this notion. We can, however, use closures to partially overcome a different hurdle: The lack of generics.

在我们进入下一个在 Go 中使用接口的主题之前，我想介绍一个常见的替代方案。这就是 C 程序员所说的“函数指针”以及大多数其他编程语言所称的<strong>闭包</strong>。闭包与其他任何参数一样只是一个输入参数，只是它表示（指向）一个可以调用的函数。在 JavaScript 中，使用闭包作为回调是很常见的，回调只是在一些异步操作完成后调用的函数。在 Go 中，我们真的没有这个概念。然而，我们可以使用闭包来部分克服一个不同的障碍：缺乏泛型。

Consider the following function signature:

考虑以下函数签名：

```go
func something(closure func(float64) float64) float64 { ... }
```

Here, `something` takes another function (a closure) as input and returns a `float64`. The input function takes a `float64` as input and also returns a `float64`. This pattern can be particularly useful for creating a loosely coupled architecture, making it easier to to add functionality without affecting other parts of the code. Suppose we have a struct containing data that we want to manipulate in some form. Through this structure's `Do()` method, we can perform operations on that data. If we know the operation ahead of time, we can obviously handle that logic directly in our `Do()` method:

这里，`something` 接受另一个函数（一个闭包）作为输入并返回一个 `float64`。输入函数接受一个 `float64` 作为输入，并返回一个 `float64`。这种模式对于创建松散耦合的体系结构特别有用，可以更轻松地添加功能而不影响代码的其他部分。假设我们有一个包含要以某种形式操作的数据的结构体。通过这个结构的`Do()`方法，我们可以对该数据执行操作。如果我们提前知道操作，我们显然可以直接在我们的 Do() 方法中处理该逻辑：

```go
func (datastore *Datastore) Do(operation Operation, data []byte) error {
  switch(operation) {
  case COMPARE:
    return datastore.compare(data)
  case CONCAT:
    return datastore.add(data)
  default:
    return ErrUnknownOperation
  }
}
```

But as you can imagine, this function is quite rigid—it performs a predetermined operation on the data contained in the `Datastore` struct. If at some point we would like to introduce more operations, we'd end up bloating our `Do` method with quite a lot of irrelevant logic that would be hard to maintain. The function would have to always care about what operation it's performing and to cycle through a number of nested options for each operation. It might also be an issue for developers wanting to use our `Datastore` object who don't have access to edit our package code, since there is no way of extending structure methods in Go as there is in most OOP languages.

但是你可以想象，这个函数是相当死板的——它对包含在 `Datastore` 结构中的数据执行预定的操作。如果在某个时候我们想引入更多操作，我们最终会用大量难以维护的不相关逻辑来膨胀我们的“Do”方法。该函数必须始终关心它正在执行的操作，并为每个操作循环访问多个嵌套选项。对于想要使用我们的 Datastore 对象但无权编辑我们的包代码的开发人员来说，这也可能是一个问题，因为在 Go 中无法像大多数 OOP 语言那样扩展结构方法。

So instead, let's try a different approach using closures:

因此，让我们尝试使用闭包的不同方法：

```go
func (datastore *Datastore) Do(operation func(data []byte, data []byte) ([]byte, error), data []byte) error {
  result, err := operation(datastore.data, data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func concat(a []byte, b []byte) ([]byte, error) {
  ...
}

func main() {
  ...
  datastore.Do(concat, data)
  ...
}
```

You'll notice immediately that the function signature for `Do` ends up being quite messy. We also have another issue: The closure isn't particularly generic. What happens if we find out that we actually want the `concat` to be able to take more than just two byte arrays as input? Or if we want to add some completely new functionality that may also need more or fewer input values than `(data []byte, data []byte)`?

您会立即注意到`Do` 的函数签名非常混乱。我们还有另一个问题：闭包不是特别通用。如果我们发现我们实际上希望 `concat` 能够接受两个以上的字节数组作为输入，会发生什么？或者，如果我们想要添加一些全新的功能，这些功能可能还需要比 `(data []byte, data []byte)` 更多或更少的输入值？

One way to solve this issue is to change our `concat` function. In the example below, I have changed it to only take a single byte array as an input argument, but it could just as well have been the opposite case:

解决此问题的一种方法是更改我们的 `concat` 函数。在下面的示例中，我已将其更改为仅将单个字节数组作为输入参数，但也可能是相反的情况：

```go
func concat(data []byte) func(data []byte) ([]byte, error) {
  return func(concatting []byte) ([]byte, error) {
    return append(data, concatting), nil
  }
}

func (datastore *Datastore) Do(operation func(data []byte) ([]byte, error)) error {
  result, err := operation(datastore.data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func main() {
  ...
  datastore.Do(compare(data))
  ...
}
```

Notice how we've effectively moved some of the clutter out of the `Do` method signature and into the `concat` method signature. Here, the `concat` function returns yet another function. Within the returned function, we store the input values originally passed in to our `concat` function. The returned function can therefore now take a single input parameter; within our function logic, we will append it to our original input value. As a newly introduced concept, this may seem quite strange. However, it's good to get used to having this as an option; it can help loosen up logic coupling and get rid of bloated functions.

请注意我们如何有效地将一些混乱从`Do` 方法签名移到`concat` 方法签名中。在这里，`concat` 函数返回另一个函数。在返回的函数中，我们存储最初传递给我们的 `concat` 函数的输入值。因此，返回的函数现在可以采用单个输入参数；在我们的函数逻辑中，我们会将它附加到我们的原始输入值。作为一个新引入的概念，这可能看起来很奇怪。但是，习惯将其作为一种选择是件好事；它可以帮助放松逻辑耦合并摆脱臃肿的功能。

In the next section, we'll get into interfaces. Before we do so, let's take a short moment to discuss the difference between interfaces and closures. First, it's worth noting that interfaces and closures definitely solve some common problems. However, the way that interfaces are implemented in Go can sometimes make it tricky to decide whether to use interfaces or closures for a particular problem. Usually, whether an interface or a closure is used isn't really of importance; the right choice is whichever one solves the problem at hand. Typically, closures will be simpler to implement if the operation is simple by nature. However, as soon as the logic contained within a closure becomes complex, one should strongly consider using an interface instead.

在下一节中，我们将进入接口。在我们这样做之前，让我们花点时间讨论一下接口和闭包之间的区别。首先，值得注意的是接口和闭包确实解决了一些常见问题。然而，在 Go 中实现接口的方式有时会使决定是使用接口还是闭包来解决特定问题变得棘手。通常，使用接口还是闭包并不重要；正确的选择是解决手头问题的任何一个。通常，如果操作本质上是简单的，那么闭包会更容易实现。然而，一旦包含在闭包中的逻辑变得复杂，就应该强烈考虑使用接口来代替。

Dave Cheney has an excellent write-up on this topic, as well as a talk:

戴夫·切尼 (Dave Cheney) 有一篇关于这个话题的精彩文章，还有一次演讲：

* https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
* https://www.youtube.com/watch?v=5buaPyJ0XeQ&t=9s

* https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
* https://www.youtube.com/watch?v=5buaPyJ0XeQ&t=9s

Jon Bodner also has a related talk:

Jon Bodner 也有一个相关的演讲：

* https://www.youtube.com/watch?v=5IKcPMJXkKs

* https://www.youtube.com/watch?v=5IKcPMJXkKs

### Interfaces in Go

### Go 中的接口

In general, Go's approach to handling `interface`s is quite different from those of other languages. Interfaces aren't explicitly implemented like they would be in Java or C#; rather, they are implicitly created if they fulfill the contract of the interface. As an example, this means that any `struct` that has an `Error()` method implements (or "fulfills") the `Error` interface and can be returned as an `error`. This manner of implementing interfaces is extremely easy and makes Go feel more fast paced and dynamic.

一般来说，Go 处理接口的方法与其他语言的处理方法大不相同。接口不像在 Java 或 C# 中那样显式实现；相反，如果它们满足接口的契约，它们就会被隐式创建。例如，这意味着任何具有 `Error()` 方法的 `struct` 都实现（或“实现”）了 `Error` 接口，并且可以作为 `error` 返回。这种实现接口的方式非常简单，并且让 Go 感觉更加快节奏和动态。

However, there are certainly disadvantages with this approach. As the interface implementation is no longer explicit, it can be difficult to see which interfaces are implemented by a struct. Therefore, it's common to define interfaces with as few methods as possible; this makes it easier to understand whether a particular struct fulfills the contract of the interface.

但是，这种方法肯定有缺点。由于接口实现不再显式，因此很难看出结构实现了哪些接口。因此，通常用尽可能少的方法定义接口；这使得更容易理解特定结构是否满足接口的约定。

An alternative is to create constructors that return an interface rather than the concrete type:

另一种方法是创建返回接口而不是具体类型的构造函数：

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}

type NullWriter struct {}

func (writer *NullWriter) Write(data []byte) (n int, err error) {
    // do nothing
    return len(data), nil
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

The above function ensures that the `NullWriter` struct implements the `Writer` interface. If we were to delete the `Write` method from `NullWriter`, we would get a compilation error. This is a good way of ensuring that our code behaves as expected and that we can rely on the compiler as a safety net in case we try to write invalid code.

上述函数确保 `NullWriter` 结构实现了 `Writer` 接口。如果我们从 `NullWriter` 中删除 `Write` 方法，我们会得到一个编译错误。这是确保我们的代码按预期运行的好方法，并且我们可以依靠编译器作为安全网，以防我们尝试编写无效代码。

In certain cases, it might not be desirable to write a constructor, or perhaps we would like for our constructor to return the concrete type, rather than the interface. As an example, the `NullWriter` struct has no properties to populate on initialisation, so writing a constructor is a little redundant. Therefore, we can use the less verbose method of checking interface compatibility:

在某些情况下，可能不需要编写构造函数，或者我们可能希望构造函数返回具体类型，而不是接口。例如，`NullWriter` 结构没有要在初始化时填充的属性，因此编写构造函数有点多余。因此，我们可以使用不那么冗长的检查接口兼容性的方法：

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}

type NullWriter struct {}
var _ io.Writer = &NullWriter{}
```

In the above code, we are initialising a variable with the Go `blank identifier`, with the type assignment of `io.Writer`. This results in our variable being checked to fulfill the `io.Writer` interface contract, before being discarded. This method of checking interface fulfillment also makes it possible to check that several interface contracts are fulfilled:

在上面的代码中，我们使用 Go `blank identifier` 初始化一个变量，类型赋值为 `io.Writer`。这导致我们的变量在被丢弃之前被检查以履行 `io.Writer` 接口契约。这种检查接口履行的方法还可以检查是否履行了几个接口合同：

```go
type NullReaderWriter struct{}
var _ io.Writer = &NullWriter{}
var _ io.Reader = &NullWriter{}
```

From the above code, it's very easy to understand which interfaces must be fulfilled; this ensures that the compiler will help us out during compile time. Therefore, this is generally the preferred solution for checking interface contract fulfillment.

从上面的代码，很容易理解必须满足哪些接口；这确保编译器会在编译时帮助我们。因此，这通常是检查接口合同履行情况的首选解决方案。

There's yet another method of trying to be more explicit about which interfaces a given struct implements. However, this third method actually achieves the opposite of what we want. It involves using embedded interfaces as a struct property.

还有另一种方法可以更明确地说明给定结构实现了哪些接口。然而，这第三种方法实际上实现了我们想要的相反。它涉及使用嵌入式接口作为结构属性。

> <em>Wait what? – Presumably most people</em>

> <em>等什么？ – 想必大多数人</em>

Let's rewind a bit before we dive deep into the forbidden forest of smelly Go. In Go, we can use embedded structs as a type of inheritance in our struct definitions. This is really nice, as we can decouple our code by defining reusable structs.

在我们深入到恶臭的围棋禁林之前，让我们先倒退一下。在 Go 中，我们可以在结构定义中使用嵌入式结构作为继承类型。这真的很好，因为我们可以通过定义可重用的结构来解耦我们的代码。

```go
type Metadata struct {
    CreatedBy types.User
}

type Document struct {
    *Metadata
    Title string
    Body string
}

type AudioFile struct {
    *Metadata
    Title string
    Body string
}
```

Above, we are defining a `Metadata` object that will provide us with property fields that we are likely to use on many different struct types. The neat thing about using the embedded struct, rather than explicitly defining the properties directly in our struct, is that it has decoupled the `Metadata` fields. Should we choose to update our `Metadata` object, we can change it in just a single place. As we've seen several times so far, we want to ensure that a change in one place in our code doesn't break other parts. Keeping these properties centralised makes it clear that structures with an embedded `Metadata` have the same properties—much like how structures that fulfill interfaces have the same methods.

上面，我们定义了一个 `Metadata` 对象，它将为我们提供可能在许多不同结构类型上使用的属性字段。使用嵌入式结构而不是直接在我们的结构中显式定义属性的巧妙之处在于它已经解耦了“元数据”字段。如果我们选择更新我们的 `Metadata` 对象，我们可以在一个地方更改它。到目前为止，我们已经多次看到，我们希望确保代码中某个地方的更改不会破坏其他部分。保持这些属性集中可以清楚地表明，带有嵌入“元数据”的结构具有相同的属性——就像实现接口的结构具有相同的方法一样。

Now, let's look at an example of how we can use a constructor to further prevent breaking our code when making changes to our `Metadata` struct:

现在，让我们看一个示例，说明如何在更改“元数据”结构时使用构造函数进一步防止破坏代码：

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
    }
}

func NewDocument(title string, body string) Document {
    return Document{
        Metadata: NewMetadata(),
        Title: title,
        Body: body,
    }
}
```

Suppose that at a later point in time, we decide that we'd also like a `CreatedAt` field on our `Metadata` object. We can now easily achieve this by simply updating our `NewMetadata` constructor:

假设在稍后的某个时间点，我们决定在我们的 `Metadata` 对象上也需要一个 `CreatedAt` 字段。我们现在可以通过简单地更新我们的 NewMetadata 构造函数来轻松实现这一点：

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
        CreatedAt: time.Now(),
    }
}
```

Now, both our `Document` and `AudioFile` structures are updated to also populate these fields on construction. This is the core principle behind decoupling and an excellent example of ensuring maintainability of code. We can also add new methods without breaking our existing code:

现在，我们的 `Document` 和 `AudioFile` 结构都更新为在构建时也填充这些字段。这是解耦背后的核心原则，也是确保代码可维护性的一个很好的例子。我们还可以在不破坏现有代码的情况下添加新方法：

```go
type Metadata struct {
    CreatedBy types.User
    CreatedAt time.Time
    UpdatedBy types.User
    UpdatedAt time.Time
}

func (metadata *Metadata) AddUpdateInfo(user types.User) {
    metadata.UpdatedBy = user
    metadata.UpdatedAt = time.Now()
}
```

Again, without breaking the rest of our codebase, we've managed to introduce new functionality. This kind of programming makes implementing new features very quick and painless, which is exactly what we are trying to achieve by writing clean code.

同样，在不破坏我们代码库的其余部分的情况下，我们已经设法引入了新功能。这种编程使得实现新功能非常快速和轻松，这正是我们试图通过编写干净的代码来实现的。

Let's return to the topic of interface contract fulfillment using embedded interfaces. Consider the following code as an example:

让我们回到使用嵌入式接口履行接口契约的话题。以以下代码为例：

```go
type NullWriter struct {
    Writer
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

The above code compiles. Technically, we are implementing the interface of `Writer` in our `NullWriter`, as `NullWriter` will inherit all the functions that are associated with this interface. Some see this as a clear way of showing that our `NullWriter` is implementing the `Writer` interface. However, we must be careful when using this technique.

上面的代码编译通过。从技术上讲，我们正在我们的 `NullWriter` 中实现 `Writer` 的接口，因为 `NullWriter` 将继承与该接口关联的所有函数。有些人认为这是一种明确的方式，表明我们的 `NullWriter` 正在实现 `Writer` 接口。但是，我们在使用这种技术时必须小心。

```go
func main() {
    w := NewNullWriter()

    w.Write([]byte{1, 2, 3})
}
```

As mentioned before, the above code will compile. The `NewNullWriter` returns a `Writer`, and everything is hunky-dory according to the compiler because `NullWriter` fulfills the contract of `io.Writer`, via the embedded interface. However, running the code above will result in the following:

如前所述，上面的代码将被编译。 `NewNullWriter` 返回一个 `Writer`，根据编译器，一切都是笨拙的，因为 `NullWriter` 通过嵌入式接口履行了 `io.Writer` 的契约。但是，运行上面的代码将导致以下结果：

> panic: runtime error: invalid memory address or nil pointer dereference 

> 恐慌：运行时错误：无效的内存地址或空指针取消引用

What happened? An interface method in Go is essentially a function pointer. In this case, since we are pointing to the function of an interface, rather than an actual method implementation, we are trying to invoke a function that's actually a `nil` pointer. To prevent this from happening, we would have to provide the `NulllWriter` with a struct that fulfills the interface contract, with actual implemented methods.

发生了什么？ Go 中的接口方法本质上是一个函数指针。在这种情况下，由于我们指向的是一个接口的函数，而不是一个实际的方法实现，我们试图调用一个实际上是一个 `nil` 指针的函数。为了防止这种情况发生，我们必须为 `NullWriter` 提供一个满足接口契约的结构，以及实际实现的方法。

```go
func main() {
  w := NullWriter{
    Writer: &bytes.Buffer{},
  }

    w.Write([]byte{1, 2, 3})
}
```

> NOTE: In the above example, `Writer` is referring to the embedded `io.Writer` interface. It is also possible to invoke the `Write` method by accessing this property with `w.Writer.Write()`.

> 注意：在上面的例子中，`Writer` 指的是内嵌的 `io.Writer` 接口。也可以通过使用 `w.Writer.Write()` 访问此属性来调用 `Write` 方法。

We are no longer triggering a panic and can now use the `NullWriter` as a `Writer`. This initialisation process is not much different from having properties that are initialised as `nil`, as discussed previously. Therefore, logically, we should try to handle them in a similar way. However, this is where embedded interfaces become a little difficult to work with. In a previous section, it was explained that the best way to handle potential `nil` values is to make the property in question private and create a public *getter* method. This way, we could ensure that our property is, in fact, not `nil`. Unfortunately, this is simply not possible with embedded interfaces, as they are by nature always public.

我们不再触发恐慌，现在可以使用 `NullWriter` 作为 `Writer`。如前所述，此初始化过程与将属性初始化为“nil”并没有太大区别。因此，从逻辑上讲，我们应该尝试以类似的方式处理它们。然而，这是嵌入式接口变得有点难以使用的地方。在上一节中，解释了处理潜在的 nil 值的最佳方法是将相关属性设为私有并创建公共 *getter* 方法。这样，我们可以确保我们的财产实际上不是“nil”。不幸的是，这对于嵌入式接口来说是不可能的，因为它们本质上总是公开的。

Another concern raised by using embedded interfaces is the potential confusion caused by partially overwritten interface methods:

使用嵌入式接口引起的另一个问题是部分覆盖接口方法引起的潜在混淆：

```go
type MyReadCloser struct {
  io.ReadCloser
}

func (closer *ReadCloser) Read(data []byte) { ... }

func main() {
  closer := MyReadCloser{}
  
  closer.Read([]byte{1, 2, 3})     // works fine
  closer.Close()         // causes panic
  closer.ReadCloser.Closer()         // no panic
}
```

Even though this might look like we're overriding methods, which is common in languages such as C# and Java, we actually aren't. Go doesn't support inheritance (and thus has no notion of a superclass). We can imitate the behaviour, but it is not a built-in part of the language. By using methods such as interface embedding without caution, we can create confusing and potentially buggy code, just to save a few more lines.

尽管这看起来像是我们覆盖了方法，这在 C# 和 Java 等语言中很常见，但实际上并非如此。 Go 不支持继承（因此没有超类的概念）。我们可以模仿这种行为，但它不是语言的内置部分。通过不小心使用诸如接口嵌入之类的方法，我们可能会创建令人困惑且可能存在错误的代码，只是为了节省更多行。

> NOTE: Some argue that using embedded interfaces is a good way of creating a mock structure for testing a subset of interface methods. Essentially, by using an embedded interface, you won't have to implement all of the methods of the interface; rather, you can choose to implement only the few methods that you'd like to test. Within the context of testing/mocking, I can see this argument, but I am still not a fan of this approach.

> 注意：有些人认为使用嵌入式接口是创建用于测试接口方法子集的模拟结构的好方法。本质上，通过使用嵌入式接口，您不必实现接口的所有方法；相反，您可以选择仅实现您想要测试的少数方法。在测试/模拟的背景下，我可以看到这个论点，但我仍然不喜欢这种方法。

Let's quickly get back to clean code and proper usage of interfaces. It's time to discuss using interfaces as function parameters and return values. The most common proverb for interface usage with functions in Go is the following:

让我们快速回到干净的代码和接口的正确使用。是时候讨论使用接口作为函数参数和返回值了。 Go 中函数的接口使用最常见的谚语如下：

> <em>Be conservative in what you do; be liberal in what you accept from others – Jon Postel</em>

> <em>做事要保守；在接受别人的东西时保持自由——乔恩·波斯特尔</em>

> FUN FACT: This proverb actually has nothing to do with Go. It's taken from an early specification of the TCP networking protocol.

> 小花絮：这句谚语实际上与 Go 无关。它取自 TCP 网络协议的早期规范。

In other words, you should write functions that accept an interface and return a concrete type. This is generally good practice and is especially useful when doing tests with mocking. As an example, we can create a function that takes a writer interface as its input and invokes the `Write` method of that interface:

换句话说，您应该编写接受接口并返回具体类型的函数。这通常是一种很好的做法，并且在使用模拟进行测试时特别有用。例如，我们可以创建一个函数，该函数将编写器接口作为其输入并调用该接口的“Write”方法：

```go
type Pipe struct {
    writer io.Writer
    buffer bytes.Buffer
}

func NewPipe(w io.Writer) *Pipe {
    return &Pipe{
        writer: w,
    }
}

func (pipe *Pipe) Save() error {
    if _, err := pipe.writer.Write(pipe.FlushBuffer());err != nil {
        return err
    }
    return nil
}
```

Let's assume that we are writing to a file when our application is running, but we don't want to write to a new file for all tests that invoke this function. We can implement a new mock type that will basically do nothing. Essentially, this is just basic dependency injection and mocking, but the point is that it is extremely easy to achieve in Go:

假设我们在应用程序运行时写入文件，但我们不想为调用此函数的所有测试写入新文件。我们可以实现一个新的模拟类型，它基本上什么都不做。本质上，这只是基本的依赖注入和模拟，但重点是在 Go 中实现它非常容易：

```go
type NullWriter struct {}

func (w *NullWriter) Write(data []byte) (int, error) {
    return len(data), nil
}

func TestFn(t *testing.T) {
    ...
    pipe := NewPipe(NullWriter{})
    ...
}
```

> NOTE: There is actually already a null writer implementation built into the `ioutil` package named `Discard`. 

> 注意：实际上已经在名为“Discard”的 `ioutil` 包中内置了一个 null 编写器实现。

When constructing our `Pipe` struct with `NullWriter` (rather than a different writer), when invoking our `Save` function, nothing will happen. The only thing we had to do was add four lines of code. This is why it is encouraged to make interfaces as small as possible in idiomatic Go—it makes it especially easy to implement patterns like the one we just saw. However, this implementation of interfaces also comes with a <em>huge</em> downside.

当使用 `NullWriter`（而不是不同的编写器）构造我们的 `Pipe` 结构时，当调用我们的 `Save` 函数时，什么都不会发生。我们唯一要做的就是添加四行代码。这就是为什么鼓励在惯用的 Go 中使接口尽可能小——这使得实现我们刚刚看到的模式特别容易。然而，这种接口的实现也有一个<em>巨大</em>的缺点。

### The Empty `interface{}`
Unlike other languages, Go does not have an implementation for generics. There have been many proposals for one, but all have been turned down by the Go language team. Unfortunately, without generics, developers must try to find creative alternatives, which very often involves using the empty `interface{}`. This section describes why these often <em>too</em> creative implementations should be considered bad practice and unclean code. There will also be examples of appropriate usage of the empty `interface{}` and how to avoid some pitfalls of writing code with it.

### 空的`接口{}`
与其他语言不同，Go 没有泛型的实现。曾经有过很多提议，但都被 Go 语言团队拒绝了。不幸的是，如果没有泛型，开发人员必须尝试寻找创造性的替代方案，这通常涉及使用空的“接口{}”。本节描述了为什么这些通常<em>太</em> 创造性的实现应该被视为不好的做法和不干净的代码。还将提供适当使用空`interface{}` 的示例以及如何避免使用它编写代码的一些陷阱。

As mentioned in a previous section, Go determines whether a concrete type implements a particular interface by checking whether the type implements the <em>methods</em> of that interface. So what happens if our interface declares no methods, as is the case with the empty interface?

如前一节所述，Go 通过检查该类型是否实现了该接口的<em>方法</em> 来确定该具体类型是否实现了该特定接口。那么如果我们的接口没有声明任何方法会发生什么，就像空接口的情况一样？

```go
type EmptyInterface interface {}
```

The above is equivalent to the built-in type `interface{}`. A natural consequence of this is that we can write generic functions that accept any type as arguments. This is extremely useful for certain kinds of functions, such as print helpers. Interestingly, this is actually what makes it possible to pass in any type to the `Println` function from the `fmt` package:

以上相当于内置类型`interface{}`。这样做的一个自然结果是我们可以编写接受任何类型作为参数的泛型函数。这对于某些类型的功能非常有用，例如打印助手。有趣的是，这实际上使得可以将任何类型从 `fmt` 包中传递给 `Println` 函数：

```go
func Println(v ...interface{}) {
    ...
}
```

In this case, `Println` isn't just accepting a single `interface{}`; rather, the function accepts a <em>slice</em> of types that implement the empty `interface{}`. As there are no methods associated with the empty `interface{}`, <em>all</em> types are accepted, even making it possible to feed `Println` with a slice of different types. This is a very common pattern when handling string conversion (both from and to a string). Good examples of this come from the `json` standard library package:

在这种情况下，`Println` 不只是接受单个 `interface{}`；相反，该函数接受一个 <em>slice</em> 类型，这些类型实现了空的 `interface{}`。由于没有与空的 `interface{}` 相关联的方法，<em>所有</em> 类型都被接受，甚至可以使用不同类型的切片提供给 `Println`。这是处理字符串转换（从字符串和到字符串）时非常常见的模式。很好的例子来自`json` 标准库包：

```go
func InsertItemHandler(w http.ResponseWriter, r *http.Request) {
    var item Item
    if err := json.NewDecoder(r.Body).Decode(&item);err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.InsertItem(item);err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatsOK)
}
```

All the less elegant code is contained within the `Decode` function. Thus, developers using this functionality won't have to worry about type reflection or type casting; we just have to worry about providing a pointer to a concrete type. This is good because the `Decode()` function is technically returning a concrete type. We are passing in our `Item` value, which will be populated from the body of the HTTP request. This means we won't have to deal with the potential risks of handling the `interface{}` value ourselves.

所有不太优雅的代码都包含在“解码”函数中。因此，使用此功能的开发人员不必担心类型反射或类型转换；我们只需要担心提供一个指向具体类型的指针。这很好，因为`Decode()` 函数在技术上返回了一个具体的类型。我们正在传递我们的 `Item` 值，该值将从 HTTP 请求的正文中填充。这意味着我们不必自己处理处理 `interface{}` 值的潜在风险。

However, even when using the empty `interface{}` with good programming practices, we still have some issues. If we pass in a JSON string that has nothing to do with our `Item` type but is still valid JSON, we won't receive an error—our `item` variable will just be left with the default values. So, while we don't have to worry about reflection and casting errors, we will still have to make sure that the message sent from our client is a valid `Item` type. Unfortunately, as of writing this document, there is no simple or good way to implement these types of generic decoders without using the empty `interface{}` type. 

然而，即使使用具有良好编程习惯的空 `interface{}`，我们仍然存在一些问题。如果我们传入一个 JSON 字符串，它与我们的 `Item` 类型无关但仍然是有效的 JSON，我们不会收到错误——我们的 `item` 变量将保留默认值。因此，虽然我们不必担心反射和转换错误，但我们仍然必须确保从我们的客户端发送的消息是有效的 `Item` 类型。不幸的是，在编写本文档时，没有简单或好的方法来实现这些类型的通用解码器而不使用空的 `interface{}` 类型。

The problem with using `interface{}` in this manner is that we are leaning towards using Go, a statically typed language, as a dynamically typed language. This becomes even clearer when looking at poor implementations of the `interface{}` type. The most common example of this comes from developers trying to implement a generic store or list of some sort.

以这种方式使用 `interface{}` 的问题是我们倾向于使用 Go，一种静态类型语言，作为一种动态类型语言。在查看 `interface{}` 类型的糟糕实现时，这一点变得更加清晰。最常见的例子来自试图实现某种通用存储或列表的开发人员。

Let's look at an example of trying to implement a generic HashMap package that can store any type using `interface{}`.

让我们看一个尝试实现通用 HashMap 包的示例，该包可以使用 `interface{}` 存储任何类型。

```go
type HashMap struct {
    store map[string]interface{}
}

func (hashmap *HashMap) Insert(key string, value interface{}) {
    hashmap.store[key] = value
}

func (hashmap *HashMap) Get(key string) (interface{}, error) {
    value, ok := hashmap.store[key]
    if !ok {
        return nil, ErrKeyNotFoundInHashMap
    }
    return value
}
```

> NOTE: I have omitted thread safety from this example to keep it simple.

> 注意：为了简单起见，我在这个例子中省略了线程安全。

Please keep in mind that the implementation pattern shown above is actually used in quite a lot of Go packages. It is even used in the standard library `sync` package for the `sync.Map` type. So what's the problem with this implementation? Well, let's have a look at an example of using the package:

请记住，上面显示的实现模式实际上在相当多的 Go 包中使用。它甚至用于“sync.Map”类型的标准库“sync”包中。那么这个实现有什么问题呢？好吧，让我们看一个使用包的例子：

```go
func SomeFunction(id string) (Item, error) {
    itemIface, err := hashmap.Get(id)
    if err != nil {
        return EmptyItem, err
    }
    item, ok := itemIface.(Item)
    if !ok {
        return EmptyItem, ErrCastingItem
    }
    return item, nil
}
```

At first glance, this looks fine. However, we'll start getting into trouble if we add <em>different</em> types to our store, something that's currently allowed. There is nothing preventing us from adding something other than the `Item` type. So what happens when someone starts adding other types into our HashMap, like a pointer `*Item` instead of an `Item`? Our function now might return an error. Worst of all, this might not even be caught by our tests. Depending on the complexity of the system, this could introduce some bugs that are particularly difficult to debug.

乍一看，这看起来不错。但是，如果我们将 <em>不同</em> 类型添加到我们的商店，我们将开始遇到麻烦，这是目前允许的。除了“Item”类型之外，没有什么可以阻止我们添加其他内容。那么当有人开始向我们的 HashMap 添加其他类型时会发生什么，比如指针 `*Item` 而不是 `Item`？我们的函数现在可能会返回错误。最糟糕的是，这甚至可能不会被我们的测试发现。根据系统的复杂性，这可能会引入一些特别难以调试的错误。

This type of code should never reach production. Remember: Go does not (yet) support generics. That's just a fact that developers must accept for the time being. If we want to use generics, then we should use a different language that does support generics rather than relying on dangerous hacks.

这种类型的代码永远不应该进入生产环境。请记住：Go（还）不支持泛型。这只是开发人员暂时必须接受的事实。如果我们想使用泛型，那么我们应该使用支持泛型的不同语言，而不是依赖危险的黑客。

So, how do we prevent this code from reaching production? The simplest solution is to just write the functions with concrete types instead of using `interface{}` values. Of course, this is not always the best approach, as there might be some functionality within the package that is not trivial to implement ourselves. Therefore, a better approach may be to create wrappers that expose the functionality we need but still ensure type safety:

那么，我们如何防止这段代码进入生产环境呢？最简单的解决方案是只用具体类型编写函数，而不是使用 `interface{}` 值。当然，这并不总是最好的方法，因为包中可能有一些功能对于我们自己来说并不容易实现。因此，更好的方法可能是创建包装器来公开我们需要的功能，但仍能确保类型安全：

```go
type ItemCache struct {
  kv tinykv.KV
}

func (cache *ItemCache) Get(id string) (Item, error) {
  value, ok := cache.kv.Get(id)
  if !ok {
    return EmptyItem, ErrItemNotFound
  }
  return interfaceToItem(value)
}

func interfaceToItem(v interface{}) (Item, error) {
  item, ok := v.(Item)
  if !ok {
    return EmptyItem, ErrCouldNotCastItem
  }
  return item, nil
}

func (cache *ItemCache) Put(id string, item Item) error {
  return cache.kv.Put(id, item)
}
```

> NOTE: Implementations of other functionalities of the `tinykv.KV` cache have been omitted for the sake of brevity.

> 注意：为简洁起见，省略了 `tinykv.KV` 缓存的其他功能的实现。

The wrapper above now ensures that we are using the actual types and that we are no longer passing in `interface{}` types. It is therefore no longer possible to accidentally populate our store with a wrong value type, and we have restricted our casting of types as much as possible. This is a very straightforward way of solving our issue, even if somewhat manually.

上面的包装器现在确保我们使用的是实际类型，并且不再传入 `interface{}` 类型。因此，不再可能意外地使用错误的值类型填充我们的商店，并且我们尽可能地限制了我们的类型转换。这是解决我们问题的一种非常直接的方法，即使有点手动。

## Summary

##  概括

First of all, thank you for making it all the way through this article! I hope it has provided some insight into clean code and how it helps ensure maintainability, readability, and stability in any codebase.

首先，感谢您让这篇文章一路走来！我希望它提供了一些关于干净代码的见解，以及它如何帮助确保任何代码库的可维护性、可读性和稳定性。

Let's briefly sum up all the topics we've covered: 

让我们简要总结一下我们涵盖的所有主题：

- <strong>Functions</strong>—A function's name should reflect its scope; the smaller the scope of a function, the more specific its name. Ensure that all functions serve a single purpose in as few lines as possible. A good rule of thumb is to limit your functions to 5–8 lines and to only accept 2–3 arguments.

- <strong>函数</strong>——函数的名称应该反映它的作用域；函数的范围越小，其名称就越具体。确保所有功能在尽可能少的行中服务于单一目的。一个好的经验法则是将你的函数限制在 5-8 行，并且只接受 2-3 个参数。

- <strong>Variables</strong>—Unlike functions, variables should assume more generic names as their scope becomes smaller. It's also recommended that you limit the scope of a variable as much as possible to prevent unintentional modification. On a similar note, you should keep the modification of variables to a minimum; this becomes an especially important consideration as the scope of a variable grows.

- <strong>变量</strong>——与函数不同，随着作用域变小，变量应该采用更通用的名称。还建议您尽可能限制变量的范围，以防止意外修改。同样，您应该尽量减少对变量的修改；随着变量范围的扩大，这成为一个特别重要的考虑因素。

- <strong>Return Values</strong>—Concrete types should be returned whenever possible. Make it as difficult as possible for users of your package to make mistakes and as easy as possible for them to understand the values returned by your functions.

- <strong>返回值</strong>——应尽可能返回具体类型。让你的包的用户尽可能地难于犯错，让他们尽可能容易地理解你的函数返回的值。

- <strong>Pointers</strong>—Use pointers with caution, and limit their scope and mutability to an absolute minimum. Remember: Garbage collection only assists with memory management; it does not assist with all of the other complexities associated with pointers.

- <strong>指针</strong>——谨慎使用指针，并将其范围和可变性限制在绝对最小值。请记住：垃圾收集仅有助于内存管理；它对与指针相关的所有其他复杂性没有帮助。

- <strong>Interfaces</strong>—Use interfaces as much as possible to loosen the coupling of your code. Hide any code using the empty `interface{}` as much as possible from end users to prevent it from being exposed.

- <strong>接口</strong>——尽可能多地使用接口来放松代码的耦合。使用空的“interface{}”尽可能对最终用户隐藏任何代码，以防止其暴露。

As a final note, it's worth mentioning that the notion of clean code is particularly subjective, and that likely won't ever change. However, much like my statement concerning `gofmt`, I think it's more important to find a common standard than something that everyone agrees with; the latter is extremely difficult to achieve.

最后，值得一提的是，干净代码的概念特别主观，而且可能永远不会改变。然而，就像我关于 `gofmt` 的声明一样，我认为找到一个共同的标准比每个人都同意的标准更重要；后者极难实现。

It's also important to understand that fanaticism is never the goal with clean code. A codebase will most likely never be fully 'clean,' in the same way that your office desk probably isn't either. There's certainly room for you to step outside the rules and boundaries covered in this article. However, remember that the most important reason for writing clean code is to help yourself and other developers. We support engineers by ensuring stability in the software we produce and by making it easier to debug faulty code. We help our fellow developers by ensuring that our code is readable and easily digestible. We help <em>everyone</em> involved in the project by establishing a flexible codebase that allows us to quickly introduce new features without breaking our current platform. We move quickly by going slowly, and everyone is satisfied.

同样重要的是要理解狂热从来都不是干净代码的目标。代码库很可能永远不会完全“干净”，就像您的办公桌可能也不是一样。您当然有空间超越本文所涵盖的规则和界限。但是，请记住，编写干净代码的最重要原因是帮助自己和其他开发人员。我们通过确保我们生产的软件的稳定性和更容易地调试错误代码来支持工程师。我们通过确保我们的代码可读且易于理解来帮助我们的开发人员。我们通过建立灵活的代码库来帮助参与项目的<em>每个人</em>，该代码库使我们能够在不破坏当前平台的情况下快速引入新功能。我们走得慢，走得快，每个人都满意。

I hope you will join this discussion to help the Go community define (and refine) the concept of clean code. Let's establish a common ground so that we can improve software—not only for ourselves but for the sake of everyone. 

我希望你能加入这个讨论来帮助 Go 社区定义（和完善）干净代码的概念。让我们建立一个共同点，这样我们就可以改进软件——不仅为了我们自己，而且为了每个人。

