# Why Go? – Key advantages you may have overlooked

# 为什么用 Go？ – 您可能忽略的主要优势

yourbasic.org/golang

Go makes it easier (than Java or Python) to write correct, clear and efficient code.

Go 使（比 Java 或 Python）更容易编写正确、清晰和高效的代码。

![Unicorn racing towards the Rainbow](https://yourbasic.org/golang/unicorn.png)

- Minimalism
   - [Features for the future](https://yourbasic.org/golang/advantages-over-java-python/#features-for-the-future)
- [Code transparency](https://yourbasic.org/golang/advantages-over-java-python/#code-transparency)
- [Compatibility](https://yourbasic.org/golang/advantages-over-java-python/#compatibility)
- [Performance](https://yourbasic.org/golang/advantages-over-java-python/#performance)

- 极简主义
  - [未来功能](https://yourbasic.org/golang/advantages-over-java-python/#features-for-the-future)
- [代码透明度](https://yourbasic.org/golang/advantages-over-java-python/#code-transparency)
- [兼容性](https://yourbasic.org/golang/advantages-over-java-python/#compatibility)
- [性能](https://yourbasic.org/golang/advantages-over-java-python/#performance)

Choosing a programming language isn’t easy. The separate features of a language may look great at first, but it takes *time* and *experience* to spot the drawbacks.

选择一种编程语言并不容易。一门语言的独立特性乍一看可能很棒，但需要*时间*和*经验*才能发现缺点。

As a CS professor and longtime Go and Java developer, I'd like to share some of my thoughts and explain why I prefer **[Go](https://golang.org/)** to **[Java]( https://www.java.com/)** or **[Python](https://www.python.org/)** – Go makes it much easier for me to write good code.

作为一名 CS 教授和长期的 Go 和 Java 开发人员，我想分享一些我的想法并解释为什么我更喜欢 **[Go](https://golang.org/)** 而不是 **[Java]( https://www.java.com/)** 或 **[Python](https://www.python.org/)** – Go 让我更容易编写好的代码。

Go has been my main programming tool since 2012, replacing Java, which in turn replaced C in 1998. I use Python mostly for teaching. My journey as a pro­gram­mer started back in 1978 on a [TI-57](https://en.wikipedia.org/wiki/TI-57) with 50 program steps and 8 registers.

自 2012 年以来，Go 一直是我的主要编程工具，取代了 Java，后者又在 1998 年取代了 C。我主要使用 Python 进行教学。我的程序员之旅始于 1978 年的 [TI-57](https://en.wikipedia.org/wiki/TI-57)，有 50 个程序步骤和 8 个寄存器。

## Minimalism

## 极简主义

![Unicorn racing towards the Rainbow](https://yourbasic.org/golang/minimalist-field-bw.jpg)

Go is a **minimalist language**, and that’s (mostly) a blessing.

Go 是一种**极简主义语言**，这（主要）是一种祝福。

The formal [Go language specification](https://golang.org/ref/spec) is only 50 pages, has plenty of examples, and is fairly easy to read. A skilled programmer could probably learn Go from the specification alone.

正式的 [Go 语言规范](https://golang.org/ref/spec) 只有 50 页，有大量示例，并且相当容易阅读。一个熟练的程序员可能只从规范中学习 Go。

The core language consists of a few **simple**, **[orthogonal](https://en.wikipedia.org/wiki/Orthogonality_(programming))** features that can be combined in a relatively small number of ways . This makes it easier to learn the language, and to read and write programs.

核心语言由几个**简单**、**[正交](https://en.wikipedia.org/wiki/Orthogonality_(programming))** 特性组成，这些特性可以通过相对较少的方式进行组合.这使得学习语言以及阅读和编写程序变得更加容易。

When you add new features to a language, the complexity doesn’t just add up, it often multiplies: language features can interact in many ways. This is a significant problem – language complexity affects **all developers** (not just the ones writing the spec and implementing the compiler).

当您向语言添加新功能时，复杂性不仅会增加，而且通常会成倍增加：语言功能可以通过多种方式进行交互。这是一个重大问题——语言复杂性影响**所有开发人员**（不仅仅是编写规范和实现编译器的开发人员）。

Here are some core Go features:

以下是 Go 的一些核心功能：

- The **built-in** frameworks for [testing](https://golang.org/doc/code.html#Testing) and [profiling](https://blog.golang.org/profiling-go-programs) are small and easy to learn, but still fully functional. There are plenty of third-party add-ons, but chances are you won’t need them.
- It's possible to **[debug](https://blog.golang.org/debugging-what-you-deploy)** and **[profile](https://blog.golang.org/profiling-go-programs)** an optimized binary running in production through an HTTP server.
- Go has [automatically generated documentation](https://blog.golang.org/godoc-documenting-go-code) with [testable examples](https://blog.golang.org/examples). Once again, the **interface is minimal**, and there is very little to learn.
- Go is **strongly** and **statically** typed with **[no implicit conversions](https://yourbasic.org/golang/conversions/)**, but the syntactic overhead is still surprisingly small. This is achieved by simple [type inference in assign­ments](https://tour.golang.org/basics/14) together with [untyped numeric constants](https://yourbasic.org/golang/untyped-constants/). This gives Go stronger type safety than Java (which has implicit conversions), but the code reads more like Python (which has untyped variables).
- Programs are constructed from [packages](https://yourbasic.org/golang/packages-explained/) that offer clear **[code separation](https://yourbasic.org/golang/public-private/)** and allow efficient management of dependencies. The package mechanism is perhaps the single most well-designed feature of the language, and certainly one of the most overlooked.
- Structurally typed [interfaces](https://yourbasic.org/golang/interfaces-explained/) provide runtime **polymorphism** through [dynamic dispatch](https://en.wikipedia.org/wiki/Runtime_polymorphism). 

- [testing](https://golang.org/doc/code.html#Testing) 和 [profiling](https://blog.golang.org/profiling-go-)的**内置**框架程序体积小且易于学习，但仍然功能齐全。有很多第三方附加组件，但您可能不需要它们。
- 可以 **[debug](https://blog.golang.org/debugging-what-you-deploy)** 和 **[profile](https://blog.golang.org/profiling-go-programs)** 通过 HTTP 服务器在生产中运行的优化二进制文件。
- Go 有 [自动生成的文档](https://blog.golang.org/godoc-documenting-go-code) 和 [可测试的例子](https://blog.golang.org/examples)。再一次，**界面极小**，可以学习的东西很少。
- Go 是 **strongly** 和 **statically** 类型的**[无隐式转换](https://yourbasic.org/golang/conversions/)**，但语法开销仍然非常小。这是通过简单的[赋值中的类型推断](https://tour.golang.org/basics/14)和[无类型数字常量](https://yourbasic.org/golang/untyped-constants/)来实现的。这为 Go 提供了比 Java（具有隐式转换）更强的类型安全性，但代码读起来更像 Python（具有无类型变量)。
- 程序由 [packages](https://yourbasic.org/golang/packages-explained/)构建，提供清晰的**[代码分离](https://yourbasic.org/golang/public-private/)* * 并允许有效管理依赖项。包机制可能是该语言设计得最好的单一特性，当然也是最容易被忽视的特性之一。
- 结构类型[接口](https://yourbasic.org/golang/interfaces-explained/)通过[动态调度](https://en.wikipedia.org/wiki/Runtime_polymorphism)提供运行时**多态性**。

- [Concurrency](https://yourbasic.org/golang/concurrent-programming/) is an **integral part** of Go, supported by [goroutines](https://yourbasic.org/golang/goroutines-explained/), [channels](https://yourbasic.org/golang/channels-explained/) and the [select statement](https://yourbasic.org/golang/select-explained/).

- [并发](https://yourbasic.org/golang/concurrent-programming/) 是 Go 的**组成部分**，由 [goroutines](https://yourbasic.org/golang/goroutines-explained)、[频道](https://yourbasic.org/golang/channels-explained/) 和 [select 语句](https://yourbasic.org/golang/select-explained/) 支持。

> See [Go vs. Java: 15 main differences](https://yourbasic.org/golang/go-vs-java/) for a small code example, and a sample of basic data types, methods and control structures.

> 请参阅 [Go 与 Java：15 个主要差异](https://yourbasic.org/golang/go-vs-java/) 以获取小代码示例以及基本数据类型、方法和控制结构的示例。

### Features for the future

### 未来的功能

![Futurist mouse trap](https://yourbasic.org/golang/futurist-mouse-trap.jpg)

Go **omits several features** found in other modern languages.

Go **省略了其他现代语言中的一些功能**。

Here’s the [language designers’ answer](https://golang.org/doc/faq#Design) to *Why does Go not have feature X?*

这是 [语言设计师的答案](https://golang.org/doc/faq#Design) *为什么 Go 没有功能 X？*

> Every language contains novel features and omits someone’s favorite feature. Go was designed with an eye on felicity of programming, speed of compilation, orthogonality of concepts, and the need to support features such as concurrency and garbage collection. Your favorite feature may be missing because it doesn’t fit, because it affects compilation speed or clarity of design, or because it would make the fundamental system model too difficult.

> 每种语言都包含新颖的功能，并省略了某人最喜欢的功能。 Go 的设计着眼于编程的便利性、编译速度、概念的正交性以及支持并发和垃圾收集等功能的需要。您最喜欢的功能可能会因为不适合，因为它影响编译速度或设计的清晰度，或者因为它会使基本系统模型变得太困难而丢失。

New features are considered only if there is a pressing need demon­strated by [experience reports](https://github.com/golang/go/wiki/ExperienceReports) from real-world projects.

仅当现实世界项目中的 [经验报告](https://github.com/golang/go/wiki/ExperienceReports) 表明存在紧迫需求时，才会考虑新功能。

There are a few likely major additions in the pipeline:

管道中有一些可能的主要新增内容：

- Package management through [modules](https://blog.golang.org/modules2019) were [preliminary introduced in Go 1.11](https://golang.org/doc/go1.11#modules).
- There is a [generics draft design](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md), which may be implemented in Go 2. For now, have a look at [Generics (alternatives and workarounds)](https://yourbasic.org/golang/generics/).
- Similarly, there is an [error handling draft design](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling.md) that extends the current minimalist [error handling](https://yourbasic.org/golang/errors-explained/).

- 通过 [modules](https://blog.golang.org/modules2019) 进行包管理 [在 Go 1.11 中初步引入](https://golang.org/doc/go1.11#modules)。
- 有一个[泛型设计草案](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)，可能会在 Go 2 中实现。现在，看看[泛型（替代方案和解决方法)](https://yourbasic.org/golang/generics/)。
- 同样，有一个[错误处理草案设计](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling.md) 扩展了当前的极简[错误处理](https://yourbasic.org/golang/errors-explained/)。

Currently, some [minor additions](https://blog.golang.org/go2-here-we-come) are considered to help create and test a [community-driven process](https://blog.golang.org/toward-go2) for developing Go.

目前，一些[次要添加](https://blog.golang.org/go2-here-we-come)被认为有助于创建和测试[社区驱动的流程](https://blog.golang.org/toward-go2) 用于开发 Go。

However, features such as [optional parameters, default parameter values and method overloading](https://yourbasic.org/golang/overload-overwrite-optional-parameter/) probably won't be part of a future [Go 2]( https://blog.golang.org/go2-here-we-come).

但是，诸如[可选参数、默认参数值和方法重载](https://yourbasic.org/golang/overload-overwrite-optional-parameter/) 等功能可能不会成为未来 [Go 2]( https://blog.golang.org/go2-here-we-come)。

### Java comparison

### Java 比较

[The Java® Language Specification](https://docs.oracle.com/javase/specs/jls/se12/jls12.pdf) is currently 750 pages. Much of the complexity is due to [feature creep](https://en.wikipedia.org/wiki/Feature_creep). Here are but three examples.

[Java® 语言规范](https://docs.oracle.com/javase/specs/jls/se12/jls12.pdf) 目前有 750 页。大部分复杂性是由于 [功能蠕变](https://en.wikipedia.org/wiki/Feature_creep)。这里仅举三个例子。

Java [inner classes](https://en.wikipedia.org/wiki/Inner_class) suddenly appeared in 1997; it took more than a year to update the specification, and it became almost twice as big as a result. That’s a high price to pay for a non-essential feature.

Java [内部类](https://en.wikipedia.org/wiki/Inner_class) 1997 年突然出现；更新规范花了一年多的时间，结果变成了几乎两倍。这是为非必要功能付出的高昂代价。

[Generics in Java](https://en.wikipedia.org/wiki/Generics_in_Java), implemented by [type erasure](https://en.wikipedia.org/wiki/Type_erasure), make some code cleaner and allow for additional runtime checks, but quickly become complex when you move beyond basic examples: [generic arrays](http://www.tothenew.com/blog/why-is-generic-array-creation-not-allowed-in-java/) aren't supported and [type wildcards](https://docs.oracle.com/javase/tutorial/extra/generics/wildcards.html) with upper and lower bounds are quite complicated. The string “generic” appears 280 times in the specification. It’s not clear to me if this feature is worth its cost.

[Java 中的泛型](https://en.wikipedia.org/wiki/Generics_in_Java)，由[typeerasure](https://en.wikipedia.org/wiki/Type_erasure) 实现，使一些代码更清晰并允许额外的运行时检查，但当您超越基本示例时，很快就会变得复杂：[通用数组](http://www.tothenew.com/blog/why-is-generic-array-creation-not-allowed-in-java/) 不受支持，并且具有上限和下限的 [类型通配符](https://docs.oracle.com/javase/tutorial/extra/generics/wildcards.html) 非常复杂。字符串“generic”在规范中出现了 280 次。我不清楚这个功能是否值得付出代价。

A Java [enum](https://docs.oracle.com/javase/tutorial/java/javaOO/enum.html), introduced in 2004, is a special type of class that represents a group of constants. It’s certainly nice to have, but offers little that couldn’t be done with ordinary classes. The string “enum” appears 241 times in the specification.

2004 年引入的 Java [enum](https://docs.oracle.com/javase/tutorial/java/javaOO/enum.html) 是一种特殊类型的类，表示一组常量。拥有它当然很好，但提供了一些普通课程无法做到的东西。字符串“enum”在规范中出现了 241 次。

## Code transparency

## 代码透明度

![Tangled cables](https://yourbasic.org/golang/eniac.jpg)

Glen Beck and Betty Snyder programming the ENIAC

Your project is doomed if you can’t understand your code.

如果你不能理解你的代码，你的项目就注定失败。

- You **always** need to know **exactly what** your coding is doing,
- and **sometimes** need to **estimate the resources** (time and memory) it uses.

- 你**总是**需要知道**确切地**你的编码在做什么，
- 并且**有时**需要**估计它使用的资源**（时间和内存）。

Go tries to meet both of these goals. 

Go 试图同时满足这两个目标。

The syntax is designed to be transparent and there is **one** [standard code format](https://golang.org/doc/effective_go.html#formatting), automatically generated by the [fmt tool](https://blog.golang.org/go-fmt-your-code).

语法设计透明，有**一个** [标准代码格式](https://golang.org/doc/effective_go.html#formatting)，由[fmt工具](https://golang.org/doc/effective_go.html#formatting)自动生成/blog.golang.org/go-fmt-your-code)。

Another example is that Go programs with [unused package imports do not compile](https://yourbasic.org/golang/unused-imports/). This improves code clarity and long-term performance.

另一个例子是带有 [未使用的包导入的 Go 程序不会编译](https://yourbasic.org/golang/unused-imports/)。这提高了代码清晰度和长期性能。

I suspect that the Go designers made some things **difficult on purpose**. You need to jump through hoops to [catch a panic](https://yourbasic.org/golang/recover-from-panic/)(exception) and you are forced to sprinkle your code with the word “unsafe” to step around type safety.

我怀疑 Go 的设计师故意制造了一些 **困难的东西**。你需要跳过箍来[抓住恐慌](https://yourbasic.org/golang/recover-from-panic/)（例外)并且你被迫在你的代码中撒上“不安全”这个词来绕过类型安全。

### Python comparison

### Python 比较

The Python code snippet `del a[i]` deletes the element at index `i` from a list `a`. This code certainly is quite **readable**, but not so **transparent**: it's easy to miss that the [time complexity](https://yourbasic.org/algorithms/time-complexity-explained/) is *O*(*n*), where *n* is the number of elements in the list.

Python 代码片段 `del a[i]` 从列表 `a` 中删除索引 `i` 处的元素。这段代码当然是**可读**，但不是那么**透明**：很容易错过 [时间复杂度](https://yourbasic.org/algorithms/time-complexity-explained/) 是 * O*(*n*)，其中*n* 是列表中的元素数。

Go doesn’t have a similar utility function. This is definitely **less convenient**, but also **more transparent**. Your code needs to explicitly state if it copies a section of the list. See [2 ways to delete an element from a slice](https://yourbasic.org/golang/delete-element-slice/) for a Go code example.

Go 没有类似的效用函数。这绝对是**不那么方便**，但也**更透明**。您的代码需要明确说明它是否复制了列表的一部分。有关 Go 代码示例，请参阅 [2 种从切片中删除元素的方法](https://yourbasic.org/golang/delete-element-slice/)。

### Java comparison

### Java 比较

Code transparency is not just a syntactic issue. Here are two examples where the rules for [Go package initialization and program execution order](https://yourbasic.org/golang/package-init-function-main-execution-order/) make it easier to reason about and maintain a project.

代码透明度不仅仅是一个语法问题。这里有两个例子，其中 [Go 包初始化和程序执行顺序](https://yourbasic.org/golang/package-init-function-main-execution-order/) 的规则更容易推理和维护项目。

- [Circular dependencies](https://en.wikipedia.org/wiki/Circular_dependency) can cause many unwanted effects. As opposed to Java, a Go program with initialization cycles will not compile.
- When the main function of a Go program returns, the program exits. A Java programs exits when all user non-daemon threads finish.

- [循环依赖](https://en.wikipedia.org/wiki/Circular_dependency) 会导致许多不良影响。与 Java 不同，具有初始化周期的 Go 程序将无法编译。
- 当 Go 程序的 main 函数返回时，程序退出。当所有用户非守护线程完成时，Java 程序退出。

This means that you may need to study large parts of a Java program to understand some of its behavior. This may even be impossible if you use third-party libraries.

这意味着您可能需要研究 Java 程序的大部分内容以了解其某些行为。如果您使用第三方库，这甚至可能是不可能的。

## Compatibility

## 兼容性

![Wrench with perfect fit](https://yourbasic.org/golang/wrench.jpg)

A language that **changes abruptly**, or becomes **unavailable**, can end your project.

**突然改变**或变得**不可用**的语言可以结束您的项目。

Go 1 has succinct and strict **[compatibility guarantees](https://golang.org/doc/go1compat)** for the core language and standard packages – Go programs that work today should continue to work with future releases of Go 1 . Backward compatibility has been excellent so far.

Go 1 为核心语言和标准包提供了简洁而严格的**[兼容性保证](https://golang.org/doc/go1compat)** - 今天运行的 Go 程序应该继续与 Go 1 的未来版本一起使用. 到目前为止，向后兼容性一直非常出色。

Go is an **open source** project and comes with a [BSD-style license](https://golang.org/LICENSE) that permits commercial use, modification, distribution, and private use. Copyright belongs to [The Go Authors](https://golang.org/AUTHORS), those of us who [contributed](https://golang.org/doc/contribute.html) to the project. There is also a [patent grant](https://golang.org/PATENTS) by Google.

Go 是一个 **开源** 项目，并带有 [BSD 风格的许可证](https://golang.org/LICENSE)，允许商业使用、修改、分发和私人使用。版权所有属于 [The Go Authors](https://golang.org/AUTHORS)，我们这些[贡献](https://golang.org/doc/contribute.html) 项目的人。谷歌还有一项[专利授权](https://golang.org/PATENTS)。

### Python comparison

### Python 比较

If you’re a Python developer, you know the pain of having to deal with the [differ­ences between Python 2.7.x and Python 3.x](https://sebastianraschka.com/Articles/2014_python_2_3_key_diff.html). There are [strong reasons](https://wiki.python.org/moin/Python2orPython3) for choosing Python 3, but if you depend on libraries that are only available for an older version, you may not be able.

如果您是一名 Python 开发人员，您就会知道必须处理 [Python 2.7.x 和 Python 3.x 之间的差异](https://sebastianraschka.com/Articles/2014_python_2_3_key_diff.html) 的痛苦。选择 Python 3 有[强有力的理由](https://wiki.python.org/moin/Python2orPython3)，但如果您依赖仅适用于旧版本的库，您可能无法使用。

### Java comparison

### Java 比较

![Lightning And Dark Clouds](https://yourbasic.org/golang/dark-clouds-lightning.jpg)

Java has a very good history of backward compatibility and the [Compatibility Guide for JDK 8](https://www.oracle.com/technetwork/java/javase/8-compatibility-guide-2156366.html) is extensive. Also, Java has been freely available to developers for a long time.

Java 具有良好的向后兼容性历史，[JDK 8 兼容性指南](https://www.oracle.com/technetwork/java/javase/8-compatibility-guide-2156366.html) 内容广泛。此外，Java 长期以来一直免费供开发人员使用。

Unfortunately, there are some dark clouds on the horizon with the [Oracle America, Inc. v. Google, Inc.](https://en.wikipedia.org/wiki/Oracle_America,_Inc._v._Google,_Inc.) legal case about the nature of computer code and copyright law, and Oracle's new [Java licensing model](https://www.infoworld.com/article/3284164/oracle-now-requires-a-subscription-to-use-java-se.html).

不幸的是，随着 [Oracle America, Inc. v. Google, Inc.](https://en.wikipedia.org/wiki/Oracle_America,_Inc._v._Google,_Inc.) 合法关于计算机代码和版权法性质的案例，以及 Oracle 新的 [Java 许可模式](https://www.infoworld.com/article/3284164/oracle-now-requires-a-subscription-to-use-java-se.html)。

## Performance

##  表现

![Lightning And Dark Clouds](https://yourbasic.org/golang/motorbike-racer.jpg)

The exterior of Go is far from flashy, but there is a **fine-tuned engine** underneath. 

Go 的外观远非华而不实，但下面有一个**微调的引擎**。

It makes little sense to discuss performance issues out of context. Running time and memory use is heavily influenced by factors such as algorithms, data structures, input, coding skill, operating systems, and hardware.

脱离上下文讨论性能问题毫无意义。运行时间和内存使用受算法、数据结构、输入、编码技能、操作系统和硬件等因素的影响很大。

Still, **language**, **runtime** and **standard libraries** can have a large effect on perfor­mance. This discussion is limited to high-level issues and design decisions. See the [Go FAQ](https://golang.org/doc/faq) for a more detailed look at the [implementation](https://golang.org/doc/faq#Implementation) and its [performance]( https://golang.org/doc/faq#Performance).

尽管如此，**语言**、**运行时**和**标准库**会对性能产生很大影响。此讨论仅限于高级问题和设计决策。有关 [实现](https://golang.org/doc/faq#Implementation) 及其 [性能]( https://golang.org/doc/faq#Performance)。

First, Go is a **compiled** language. An executable Go program typically consists of a **single standalone binary**, with no separate dynamic libraries or virtual machines, which can be **directly deployed**.

首先，Go 是一种**编译**语言。一个可执行的 Go 程序通常由一个**单个独立二进制**组成，没有单独的动态库或虚拟机，可以**直接部署**。

**Size and speed of generated code** will vary depending on target architecture. Go code generation is [fairly mature](https://about.sourcegraph.com/go/generating-better-machine-code-with-ssa/) and the major OSes (Linux, macOS, Windows) and architectures (Intel x86 /x86-64, ARM64, WebAssembly, ARM), as well as many others, are supported. You can expect performance to be on a similar level to that of C++ or Java. Compared to interpreted Python code, the improvement can be huge.

**生成代码的大小和速度**会因目标架构而异。 Go 代码生成 [相当成熟](https://about.sourcegraph.com/go/generating-better-machine-code-with-ssa/) 和主要操作系统 (Linux, macOS, Windows) 和架构 (Intel x86 /x86-64、ARM64、WebAssembly、ARM) 以及许多其他文件都受支持。您可以期望性能与 C++ 或 Java 处于相似的水平。与解释型 Python 代码相比，改进可能是巨大的。

Go is **garbage collected**, protecting against memory leaks. The collection has **[very low latency](https://blog.golang.org/ismmkeynote)**. In fact, you may never notice that the GC thread is there.

Go 是 **垃圾收集**，防止内存泄漏。该集合具有**[非常低的延迟](https://blog.golang.org/ismmkeynote)**。事实上，您可能永远不会注意到 GC 线程在那里。

The **standard libraries** are typically of high quality, with optimized code using efficient algorithms. As an example, [regular expressions](https://yourbasic.org/golang/regexp-cheat-sheet/) are very efficient in Go with running time linear in the size of the input. Unfortunately, this is [not true for Java and Python](https://swtch.com/~rsc/regexp/regexp1.html).

**标准库** 通常是高质量的，使用高效算法优化代码。例如，[正则表达式](https://yourbasic.org/golang/regexp-cheat-sheet/) 在 Go 中非常高效，运行时间与输入大小成线性关系。不幸的是，这 [不适用于 Java 和 Python](https://swtch.com/~rsc/regexp/regexp1.html)。

**Build speeds**, in absolute terms, are currently fairly good. More importantly, Go is **designed** to make compilation and dependency analysis easy, making it possible to create programming tools that **scales well** with growing projects. 

**构建速度**，从绝对值来看，目前相当不错。更重要的是，Go 的**设计**是为了使编译和依赖项分析变得容易，从而可以创建**可扩展**与不断增长的项目的编程工具。

