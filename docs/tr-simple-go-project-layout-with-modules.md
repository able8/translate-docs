# Simple Go project layout with modules

# 带有模块的简单 Go 项目布局

October 01, 2019

*[Updated and verified on 2020-10-21]*

*[2020-10-21更新验证]*

A very common question Go beginners have is "how do I organize my code?". Some of the things folks are wondering about are:

Go 初学者有一个非常常见的问题是“我如何组织我的代码？”。人们想知道的一些事情是：

- How does my repository structure reflect the way users import my code?
- How do I distribute *commands* (command-line programs that users can install) in addition to code?
- How do modules change the way I organize my code?
- How do multiple packages coexist in a single module?

- 我的存储库结构如何反映用户导入我的代码的方式？
- 除了代码之外，我如何分发*命令*（用户可以安装的命令行程序）？
- 模块如何改变我组织代码的方式？
- 多个包如何共存于一个模块中？

Unfortunately, there is some easy-to-find advice online that's outdated and over-complicated, so I wanted to create an example that's both minimal and up-to-date. I believe that in these cases it's better to provide an example that's small and easy to understand. Advanced users can grow their projects from a simple starting point, if needed.

不幸的是，网上有一些过时且过于复杂的易于查找的建议，因此我想创建一个既简约又最新的示例。我相信在这些情况下，最好提供一个小而易于理解的示例。如果需要，高级用户可以从一个简单的起点发展他们的项目。

The concepts demonstrated here:

这里展示的概念：

- Splitting a module into multiple packages, each importable by users; some of these packages import others (from within the same module).
- *Internal* packages, only importable from other packages in their module, not by outside users.
- Commands/programs that users can install with `go get`.

- 将一个模块拆分成多个包，每个包都可由用户导入；其中一些包导入其他包（从同一模块中）。
- *内部*包，只能从其模块中的其他包导入，不能由外部用户导入。
- 用户可以使用 `go get` 安装的命令/程序。

Just one small definition and we'll get started: when I say *user* I mean the developer who is using my module, either by `import`-ing it in their code, or by `go get`-ing a program.

只需一个小定义，我们就可以开始了：当我说 *user* 时，我指的是使用我的模块的开发人员，通过在他们的代码中通过 `import`-ing 或通过 `go get`-ing 一个程序。

## Getting started

##  入门

The sample project this post describes is on GitHub: https://github.com/eliben/modlib

这篇文章描述的示例项目位于 GitHub 上：https://github.com/eliben/modlib

The project path is the *module name*. The `go.mod` file for the project contains this line:

项目路径是*模块名称*。该项目的 `go.mod` 文件包含以下行：

```
module github.com/eliben/modlib
```

It is very common for Go projects to be named by their GitHub path. Go also supports [custom names](https://golang.org/cmd/go/#hdr-Remote_import_paths), but that is outside the scope of this post. Throughout the post, you can substitute `github.com/eliben/modlib` with `github.com/your-handle/your-project` or `your-project-domain.io`, whatever works for you.

Go 项目以它们的 GitHub 路径命名是很常见的。 Go 还支持 [自定义名称](https://golang.org/cmd/go/#hdr-Remote_import_paths)，但这超出了本文的范围。在整篇文章中，你可以用 `github.com/your-handle/your-project` 或 `your-project-domain.io` 替换 `github.com/eliben/modlib`，任何适合你的方法。

The module name is extremely important, because it serves as the basis of imported names in user code:

模块名称非常重要，因为它是用户代码中导入名称的基础：

![Import path example with arrows showing module name and package](https://eli.thegreenplace.net/images/2019/import-module-arrows.png)

## Project layout

## 项目布局

Here is the directory and file layout of the modlib repository:

这是 modlib 存储库的目录和文件布局：

```
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib
│   ├── lib.go
│   └── lib_test.go
├── cmd
│   ├── modlib-client
│   │   └── main.go
│   └── modlib-server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── serverlib
    └── lib.go
```

Let's start with the files in the root directory.

让我们从根目录中的文件开始。

**LICENSE** and **README.md** are fairly obvious and I won't spend time on them here.

**LICENSE** 和 **README.md** 相当明显，我不会在这里花时间讨论它们。

**go.mod** is the module definition file. It contains the module name shown above and that's it - my project has no dependencies. Dependencies are a whole different topic, quite unrelated to project layout. There's a lot of good documentation online. I suggest starting with the official blog posts - [part 1](https://blog.golang.org/using-go-modules), [part 2](https://blog.golang.org/migrating-to-go-modules), and [part 3](https://blog.golang.org/publishing-go-modules).

**go.mod** 是模块定义文件。它包含上面显示的模块名称，就是这样 - 我的项目没有依赖项。依赖是一个完全不同的话题，与项目布局完全无关。网上有很多很好的文档。我建议从官方博客文章开始 - [part 1](https://blog.golang.org/using-go-modules), [part 2](https://blog.golang.org/migrating-to-go-modules) 和 [第 3 部分](https://blog.golang.org/publishing-go-modules)。

**go.sum** contains all the dependency checksums, and is managed by the `go` tools. You don't have to worry about it, but keep it checked into source control alongside **go.mod**.

**go.sum** 包含所有的依赖校验和，由 `go` 工具管理。您不必担心它，但将其与 **go.mod** 一起检查到源代码管理中。

**config.go** this is the first code file we're examining; it contains a single trivial function [[1\]](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/#id3):

**config.go** 这是我们正在检查的第一个代码文件；它包含一个简单的函数 [[1\]](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/#id3)：

```go
package modlib

func Config() string {
  return "modlib config"
}
```

The most important part here is the `package modlib`. Since this file is at the top level of the module, its package name is considered to be the module name. This is what you get when you just import `github.com/eliben/modlib`. The user code can look like this ([Playground link](https://play.golang.org/p/tXawUZ9j502)):

这里最重要的部分是`package modlib`。由于这个文件在模块的顶层，它的包名被认为是模块名。这就是你导入 `github.com/eliben/modlib` 时得到的结果。用户代码可能如下所示（[Playground 链接](https://play.golang.org/p/tXawUZ9j502))：

```go
package main

import "fmt"
import "github.com/eliben/modlib"

func main() {
  fmt.Println(modlib.Config())
}
```

So the rule is simple: if your module provides a single package, or you want to export code from the top-level package of the module, place all the code for this at the top-level directory of the module, and name the package as the last part of the module's path (unless you're using vanity imports, in which case it's more flexible).

所以规则很简单：如果你的模块提供单个包，或者你想从模块的顶级包中导出代码，请将所有代码放在模块的顶级目录中，并命名包作为模块路径的最后一部分（除非您使用虚导入，在这种情况下它更灵活）。

## Additional packages

## 附加包

Now moving on to the **clientlib** directory.

现在转到 **clientlib** 目录。

**clientlib/lib.go** is a file in the `clientlib` package of our module. It doesn't matter what the file is called, and many packages consist of multiple files. What's important is that the `package` declaration at the top of the file says `clientlib`:

**clientlib/lib.go** 是我们模块的 `clientlib` 包中的一个文件。文件叫什么并不重要，许多包由多个文件组成。重要的是文件顶部的“package”声明是“clientlib”：

```go
package clientlib

func Hello() string {
  return "clientlib hello"
}
```

User code will import this package with `github.com/eliben/modlib/clientlib`, as follows ([Playground link](https://play.golang.org/p/pe_uAr52Kdy)):

用户代码将使用 `github.com/eliben/modlib/clientlib` 导入这个包，如下（[Playground 链接](https://play.golang.org/p/pe_uAr52Kdy))：

```go
package main

import "fmt"
import "github.com/eliben/modlib"
import "github.com/eliben/modlib/clientlib"

func main() {
  fmt.Println(modlib.Config())
  fmt.Println(clientlib.Hello())
}
```

The **serverlib** directory contains another package users can import. There's nothing new there - just showing how multiple packages live alongside each other.

**serverlib** 目录包含用户可以导入的另一个包。没有什么新东西 - 只是展示了多个包是如何相互并存的。

A quick word on nesting of packages: it can go as deep as you need. The package name visible to users is determined by the *relative path* from the module root. For example, if we have a subdirectory called `clientlib/tokens` with some code in the `tokens` package, the user will import that with `import "github.com/eliben/modlib/clientlib/tokens`.

关于包嵌套的简短说明：它可以根据需要深入。用户可见的包名由模块根目录的*相对路径*决定。例如，如果我们有一个名为 `clientlib/tokens` 的子目录，在 `tokens` 包中有一些代码，用户将使用 `import "github.com/eliben/modlib/clientlib/tokens` 导入它。

It's also important to highlight that for some modules a single top-level package is sufficient. In the case of `modlib` this would mean no subdirectories with user-importable packages, but all code being in the top directory in a single or multiple Go files all in `package modlib`.

还需要强调的是，对于某些模块，单个顶级包就足够了。在 `modlib` 的情况下，这意味着没有包含用户可导入包的子目录，但所有代码都位于 `package modlib` 中的单个或多个 Go 文件中的顶级目录中。

## Commands / programs

## 命令/程序

Some Go projects distribute *programs*, or *commands*, instead of (or in addition to) importable packages. If this isn't relevant to your project, feel free to skip this section and don't add a **cmd** directory.

一些 Go 项目分发 *programs* 或 *commands*，而不是（或除了）可导入包。如果这与您的项目无关，请随意跳过本节，不要添加 **cmd** 目录。

The **cmd** directory is the conventional location of all the command-line programs made available by the project. The naming scheme for programs is typically:

**cmd** 目录是项目提供的所有命令行程序的常规位置。程序的命名方案通常是：

![Path for commands in a repository](https://eli.thegreenplace.net/images/2019/command-paths.png)

Such commands can be installed by the user using the `go` tool as follows:

用户可以使用 go 工具安装这些命令，如下所示：

```
$ go get github.com/eliben/modlib/cmd/cmd-name

# Go downloads, builds and installs cmd-name into the default location.
# The bin/ directory in the default location is often in $PATH, so we can
# just invoke cmd-name now

$ cmd-name ...
```

In modlib, there are two different command-line programs provided, as an example: `modlib-client` and `modlib-server`. In each of them, the code is in `package main`; the filename is also called `main.go`, but this isn't a requirement. It doesn't matter what the file names are called, as long as they're in `package main`.

在 modlib 中，提供了两个不同的命令行程序，例如：`modlib-client` 和 `modlib-server`。在每一个中，代码都在`package main`中；文件名也称为“main.go”，但这不是必需的。文件名如何命名并不重要，只要它们在 `package main` 中即可。

In fact, since modlib is a real repository, you can install and run these tools on your machine:

事实上，由于 modlib 是一个真正的存储库，你可以在你的机器上安装和运行这些工具：

```go
$ go get github.com/eliben/modlib/cmd/modlib-client
$ modlib-client
Running client
Config: modlib config
clientlib hello

$ go get github.com/eliben/modlib/cmd/modlib-server
$ modlib-server
Running server
Config: modlib config
Auth: thou art authorized
serverlib hello

# Clean up...
$ rm -f `which modlib-server` `which modlib-client`
```

It's instructional to take a look at the [code of modlib-server](https://github.com/eliben/modlib/blob/master/cmd/modlib-server/main.go):

看看[modlib-server的代码](https://github.com/eliben/modlib/blob/master/cmd/modlib-server/main.go)是有指导意义的：

```go
package main

import (
  "fmt"

  "github.com/eliben/modlib"
  "github.com/eliben/modlib/internal/auth"
  "github.com/eliben/modlib/serverlib"
)

func main() {
  fmt.Println("Running server")
  fmt.Println("Config:", modlib.Config())
  fmt.Println("Auth:", auth.GetAuth())
  fmt.Println(serverlib.Hello())
}
```

The important thing I want to highlight here is how it imports other code from modlib. In Go, absolute imports are the way to go - note how this is used here. This applies to packages as well, not just commands. If code in package `clientlib` needs to import the main `modlib` package, it will do so by `import github.com/eliben/modlib`.

我想在这里强调的重要一点是它如何从 modlib 导入其他代码。在 Go 中，绝对导入是要走的路——注意这里是如何使用的。这也适用于包，而不仅仅是命令。如果`clientlib` 包中的代码需要导入主`modlib` 包，它将通过`import github.com/eliben/modlib` 来实现。

## Internal packages

## 内部包

Another important concept is internal (or private) packages - packages that are used internally by a project, but which we don't want to export to users. This is especially important in Go with modules, due to semantic versioning. Everything exported by your project in `v1` becomes a *public API*, and has to abide by semantic versioning compatibility guarantees. Therefore, it's imperative to export only the minimal API surface that's essential for users of your project. All the other code which your package needs for its implementation should live in **internal**.

另一个重要的概念是内部（或私有）包 - 项目内部使用的包，但我们不想导出给用户。由于语义版本控制，这在带有模块的 Go 中尤为重要。您的项目在 v1 中导出的所有内容都成为*公共 API*，并且必须遵守语义版本控制兼容性保证。因此，必须仅导出对项目用户至关重要的最小 API 表面。您的包实现所需的所有其他代码都应位于 **internal** 中。

The Go tooling recognizes **internal** as a special path. Packages in the same module can import it as usual (see the previous code snippet, for example). Users (that is, code outside the module) cannot import it, though. If we try to do this, [we get an error](https://play.golang.org/p/VdvEEcxrJiO):

Go 工具将 **internal** 识别为特殊路径。同一个模块中的包可以像往常一样导入它（例如，参见前面的代码片段）。但是，用户（即模块外的代码）无法导入它。如果我们尝试这样做，[我们得到一个错误](https://play.golang.org/p/VdvEEcxrJiO)：

```
use of internal package github.com/eliben/modlib/internal/auth not allowed
```

In the modlib project, there's a single package in **internal**. In real projects, there is often a whole tree of packages there.

在 modlib 项目中，**internal** 中有一个包。在实际项目中，通常有一个完整的包树。

If you're wondering whether some package belongs in **internal**, it's prudent to begin by answering "yes". It's easy to take an internal API and export it to users - just a quick renaming/refactoring commit. It's very painful to take an external API and un-export it (user code may depend on it); at stable module versions (`v1` and beyond), this requires a major version bump to break compatibility [[2\]](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/#id4).

如果您想知道某个包是否属于 **internal**，最好先回答“是”。使用内部 API 并将其导出给用户很容易 - 只需快速重命名/重构提交。获取外部 API 并取消导出它是非常痛苦的（用户代码可能依赖于它）；在稳定的模块版本（`v1` 及更高版本）中，这需要一个主要的版本来破坏兼容性 [[2\]](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-模块/#id4)。

I really like to put as much as possible in **internal**, not only private Go packages needed by my module. For example, if the repository contains the source code of the website of the project, I'd place that in `internal/website`. The same goes for internal tools/scripts needed to work on the project. The idea is that the root directory of a project should be minimal and clear to *users*. In a way, it's self-documentation. A user looking at my project's GitHub page should get an immediate sense of where the things they need are located. Since users don't typically really need the stuff I use to *develop* the project, hiding it in **internal** makes sense.

我真的很喜欢将尽可能多的放在 **internal** 中，而不仅仅是我的模块需要的私有 Go 包。例如，如果存储库包含项目网站的源代码，我会将其放在 `internal/website` 中。处理项目所需的内部工具/脚本也是如此。这个想法是一个项目的根目录应该是最小的并且对*用户*来说是清晰的。在某种程度上，它是自我记录。查看我项目的 GitHub 页面的用户应该可以立即了解他们需要的东西在哪里。由于用户通常并不真正需要我用来*开发*项目的东西，因此将其隐藏在 **internal** 中是有意义的。

## But what about a pkg/ directory?

## 但是 pkg/ 目录呢？

In some Go repositories you'll find a `pkg/` directory with importable packages, and some online guides recommend having such a directory in your hierarchy. Why haven't I mentioned it so far?

在一些 Go 存储库中，您会找到一个包含可导入包的 `pkg/` 目录，并且一些在线指南建议在您的层次结构中拥有这样一个目录。为什么我到现在还没有提到它？

In my personal view, while you may want a `pkg/` directory in some rare scenarios, **in the majority of cases it's an antipattern**. It's much better to start your project without it. Here's why.

在我个人看来，虽然在一些罕见的情况下你可能需要一个 `pkg/` 目录，**在大多数情况下它是一个反模式**。最好在没有它的情况下开始您的项目。这是为什么。

A `pkg/` directory is commonly found/recommended in large projects where a complete application lives in a single repository; this application may contain Go packages, but also tools, static assets (HTML, CSS etc.), configuration and deployment scripts, and so on. In these cases it may seem unwise to scatter a bunch of Go package directories around in the repository, creating confusion about what's where.

`pkg/` 目录通常在大型项目中找到/推荐，其中完整的应用程序位于单个存储库中；该应用程序可能包含 Go 包，但也包含工具、静态资产（HTML、CSS 等）、配置和部署脚本等。在这些情况下，将一堆 Go 包目录分散在存储库中似乎是不明智的，这会造成对什么位置的混淆。

That could certainly happen, but I'd argue that in such applications the code you place in `pkg/` should almost certainly be in `internal/` instead. If your project is a large top-level application, it shouldn't have importable packages; instead, importable packages should be split out to separate repositories which are small, self-contained and reusable. Don't forget that Go's semantic versioning applies at the module level. 

这肯定会发生，但我认为在这样的应用程序中，你放在 `pkg/` 中的代码几乎肯定应该放在 `internal/` 中。如果你的项目是一个大型的顶级应用程序，它不应该有可导入的包；相反，应将可导入的包拆分到独立的小、自包含和可重用的存储库中。不要忘记 Go 的语义版本控制适用于模块级别。

What about projects that truly contain only importable packages? Well, then you most likely don't need `pkg/` either, because it's just empty filling adding 4 characters to every import path using your project without any real benefit. If your project is an importable module, just follow the advice from the rest of this post. Many of the most popular Go modules like [gorilla/mux](https://github.com/gorilla/mux) and [cobra](https://github.com/spf13/cobra) do just fine without a `pkg /` directory.

那些真正只包含可导入包的项目呢？好吧，那么您很可能也不需要`pkg/`，因为它只是空填充，使用您的项目向每个导入路径添加 4 个字符，没有任何实际好处。如果您的项目是可导入模块，只需遵循本文其余部分的建议即可。许多最流行的 Go 模块，如 [gorilla/mux](https://github.com/gorilla/mux) 和 [cobra](https://github.com/spf13/cobra) 在没有 pkg 的情况下也能正常运行 目录。

To conclude, if you believe you need a `pkg/` directory, spend some time thinking whether you *really* need it. In my experience, 90% of Go projects don't need a separate directory for their packages at all; out of those that do need one, 90% should choose to place their packages in `internal/`. If your project is truly in the 1% that could benefit from `pkg/`, that's absolutely fine! Just keep in mind that the odds for this are low.

总而言之，如果您认为您需要一个 `pkg/` 目录，请花一些时间思考您是否*真的*需要它。根据我的经验，90% 的 Go 项目根本不需要单独的包目录；在那些确实需要一个的人中，90% 应该选择将他们的包裹放在 `internal/` 中。如果您的项目确实是可以从 `pkg/` 中受益的 1%，那绝对没问题！请记住，这种情况的可能性很低。

Most importantly, start simple. 

最重要的是，从简单开始。

