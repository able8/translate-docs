# How to Write Go Code

# 如何编写 Go 代码

## Introduction

##  介绍

This document demonstrates the development of a simple Go package inside a module and introduces the [go tool](https://golang.org/cmd/go/), the standard way to fetch, build, and install Go modules, packages, and commands.

本文档演示了在模块内开发一个简单的 Go 包，并介绍了 [go 工具](https://golang.org/cmd/go/)，这是获取、构建和安装 Go 模块、包、和命令。

Note: This document assumes that you are using Go 1.13 or later and the `GO111MODULE` environment variable is not set. If you are looking for the older, pre-modules version of this document, it is archived [here](https://golang.org/doc/gopath_code.html).

注意：本文档假设您使用的是 Go 1.13 或更高版本，并且未设置 `GO111MODULE` 环境变量。如果您正在寻找本文档的较旧的模块前版本，它已存档 [此处](https://golang.org/doc/gopath_code.html)。

## Code organization

## 代码组织

Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

Go 程序被组织成包。包是同一目录下编译在一起的源文件的集合。在一个源文件中定义的函数、类型、变量和常量对同一包中的所有其他源文件可见。

A repository contains one or more modules. A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).

一个存储库包含一个或多个模块。模块是一起发布的相关 Go 包的集合。 Go 存储库通常只包含一个模块，位于存储库的根目录。一个名为 go.mod 的文件在那里声明了模块路径：模块中所有包的导入路径前缀。该模块在包含其 `go.mod` 文件的目录中包含包以及该目录的子目录，直到包含另一个 `go.mod` 文件（如果有）的下一个子目录。

Note that you don't need to publish your code to a remote repository before you can build it. A module can be defined locally without belonging to a repository. However, it's a good habit to organize your code as if you will publish it someday.

请注意，在构建代码之前，您无需将代码发布到远程存储库。模块可以在本地定义而不属于存储库。然而，组织你的代码就像你有一天会发布它是一个好习惯。

Each module's path not only serves as an import path prefix for its packages, but also indicates where the `go` command should look to download it. For example, in order to download the module `golang.org/x/tools`, the `go` command would consult the repository indicated by `https://golang.org/x/tools` (described more [here]( https://golang.org/cmd/go/#hdr-Relative_import_paths)).

每个模块的路径不仅用作其包的导入路径前缀，而且还指示 `go` 命令应该在哪里下载它。例如，为了下载模块 `golang.org/x/tools`，`go` 命令将查询由 `https://golang.org/x/tools` 指示的存储库（详细描述 [此处]( https://golang.org/cmd/go/#hdr-Relative_import_paths))。

An import path is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. For example, the module `github.com/google/go-cmp` contains a package in the directory `cmp/`. That package's import path is `github.com/google/go-cmp/cmp`. Packages in the standard library do not have a module path prefix.

导入路径是用于导入包的字符串。包的导入路径是其模块路径与其在模块内的子目录相连。例如，模块 `github.com/google/go-cmp` 在目录 `cmp/` 中包含一个包。该包的导入路径是`github.com/google/go-cmp/cmp`。标准库中的包没有模块路径前缀。

## Your first program

## 你的第一个程序

To compile and run a simple program, first choose a module path (we'll use `example/user/hello`) and create a `go.mod` file that declares it:

要编译和运行一个简单的程序，首先选择一个模块路径（我们将使用 `example/user/hello`）并创建一个声明它的 `go.mod` 文件：

```sh
$ mkdir hello # Alternatively, clone it if it already exists in version control.
$ cd hello
$ go mod init example/user/hello
go: creating new go.mod: module example/user/hello
$ cat go.mod
module example/user/hello

go 1.16
$
```

The first statement in a Go source file must be `package name`. Executable commands must always use `package main`.

Go 源文件中的第一条语句必须是 `package name`。可执行命令必须始终使用`package main`。

Next, create a file named `hello.go` inside that directory containing the following Go code:

接下来，在该目录中创建一个名为“hello.go”的文件，其中包含以下 Go 代码：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
```

Now you can build and install that program with the `go` tool:

现在您可以使用 `go` 工具构建和安装该程序：

```
$ go install example/user/hello
$
```

This command builds the `hello` command, producing an executable binary. It then installs that binary as `$HOME/go/bin/hello` (or, under Windows, `%USERPROFILE%\go\bin\hello.exe`).

此命令构建 `hello` 命令，生成可执行二进制文件。然后它会将该二进制文件安装为 `$HOME/go/bin/hello`（或者，在 Windows 下，是 `%USERPROFILE%\go\bin\hello.exe`）。

The install directory is controlled by the `GOPATH` and `GOBIN` [environment variables](https://golang.org/cmd/go/#hdr-Environment_variables). If `GOBIN` is set, binaries are installed to that directory. If `GOPATH` is set, binaries are installed to the `bin` subdirectory of the first directory in the `GOPATH` list. Otherwise, binaries are installed to the `bin` subdirectory of the default `GOPATH` (`$HOME/go` or `%USERPROFILE%\go`).

安装目录由`GOPATH`和`GOBIN`[环境变量](https://golang.org/cmd/go/#hdr-Environment_variables)控制。如果设置了`GOBIN`，二进制文件将安装到该目录中。如果设置了 `GOPATH`，二进制文件将安装到 `GOPATH` 列表中第一个目录的 `bin` 子目录中。否则，二进制文件将安装到默认 `GOPATH`（`$HOME/go` 或 `%USERPROFILE%\go`)的 `bin` 子目录中。

You can use the `go env` command to portably set the default value for an environment variable for future `go` commands:

您可以使用 `go env` 命令为将来的 `go` 命令可移植地设置环境变量的默认值：

```
$ go env -w GOBIN=/somewhere/else/bin
$
```

To unset a variable previously set by `go env -w`, use `go env -u`:

要取消设置先前由 `go env -w` 设置的变量，请使用 `go env -u`：

```
$ go env -u GOBIN
$
```

Commands like `go install` apply within the context of the module containing the current working directory. If the working directory is not within the `example/user/hello` module, `go install` may fail. 

`go install` 之类的命令适用于包含当前工作目录的模块的上下文。如果工作目录不在 `example/user/hello` 模块中，`go install` 可能会失败。

For convenience, `go` commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given. So in our working directory, the following commands are all equivalent:

为方便起见，`go` 命令接受相对于工作目录的路径，如果没有给出其他路径，则默认为当前工作目录中的包。所以在我们的工作目录中，以下命令都是等价的：

```
$ go install example/user/hello
$ go install .
$ go install
```

Next, let's run the program to ensure it works. For added convenience, we'll add the install directory to our `PATH` to make running binaries easy:

接下来，让我们运行该程序以确保其正常工作。为了更加方便，我们将安装目录添加到我们的 `PATH` 中，以方便运行二进制文件：

```
# Windows users should consult https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

If you're using a source control system, now would be a good time to initialize a repository, add the files, and commit your first change. Again, this step is optional: you do not need to use source control to write Go code.

如果您使用的是源代码控制系统，现在是初始化存储库、添加文件和提交第一个更改的好时机。同样，这一步是可选的：您不需要使用源代码控制来编写 Go 代码。

```
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```

The `go` command locates the repository containing a given  module path by requesting a corresponding HTTPS URL and reading metadata embedded in the HTML response (see `go help importpath`). Many hosting services already provide that metadata for repositories containing Go code, so the easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.

`go` 命令通过请求相应的 HTTPS URL 并读取嵌入在 HTML 响应中的元数据来定位包含给定模块路径的存储库（参见 `go help importpath`）。许多托管服务已经为包含 Go 代码的存储库提供了元数据，因此让您的模块可供其他人使用的最简单方法通常是使其模块路径与存储库的 URL 匹配。

### Importing packages from your module

### 从你的模块导入包

Let's write a `morestrings` package and use it from the `hello` program. First, create a directory for the package named `$HOME/hello/morestrings`, and then a file named `reverse.go` in that directory with the following contents:

让我们编写一个 `morestrings` 包并在 `hello` 程序中使用它。首先，为名为“$HOME/hello/morestrings”的包创建一个目录，然后在该目录中创建一个名为“reverse.go”的文件，其内容如下：

```go
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1;i < len(r)/2;i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

Because our `ReverseRunes` function begins with an upper-case letter, it is [exported](https://golang.org/ref/spec#Exported_identifiers), and can be used in other packages that import our `morestrings` package.

因为我们的 `ReverseRunes` 函数以大写字母开头，所以它是 [exported](https://golang.org/ref/spec#Exported_identifiers)，并且可以在导入我们的 `morestrings` 包的其他包中使用。

Let's test that the package compiles with `go build`:

让我们测试包是否使用 `go build` 编译：

```
$ cd $HOME/hello/morestrings
$ go build
$
```

This won't produce an output file. Instead it saves the compiled package in the local build cache.

这不会产生输出文件。相反，它将编译的包保存在本地构建缓存中。

After confirming that the `morestrings` package builds, let's use it from the `hello` program. To do so, modify your original `$HOME/hello/hello.go` to use the morestrings package:

在确认`morestrings` 包构建完成后，让我们从`hello` 程序中使用它。为此，请修改您原来的 `$HOME/hello/hello.go` 以使用 morestrings 包：

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
```

Install the `hello` program:

安装 `hello` 程序：

```
$ go install example/user/hello
```

Running the new version of the program, you should see a new, reversed message:

运行新版本的程序，您应该会看到一条新的反向消息：

```
$ hello
Hello, Go!
```

### Importing packages from remote modules

### 从远程模块导入包

An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The `go` tool uses this property to automatically fetch packages from remote repositories. For instance, to use `github.com/google/go-cmp/cmp` in your program:

导入路径可以描述如何使用 Git 或 Mercurial 等版本控制系统获取包源代码。 `go` 工具使用此属性自动从远程存储库中获取包。例如，要在您的程序中使用 `github.com/google/go-cmp/cmp`：

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

Now that you have a dependency on an external module, you need to download that module and record its version in your `go.mod` file. The `go mod tidy` command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.

现在您依赖于外部模块，您需要下载该模块并将其版本记录在您的 `go.mod` 文件中。 `go mod tidy` 命令为导入的包添加缺少的模块要求，并删除不再使用的模块的要求。

```sh
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
- "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

Module dependencies are automatically downloaded to the `pkg/mod` subdirectory of the directory indicated by the `GOPATH` environment variable. The downloaded contents for a given version of a module are shared among all other modules that `require` that version, so the `go` command marks those files and directories as read-only. To remove all downloaded modules, you can pass the `-modcache` flag to `go clean`:

模块依赖会自动下载到 `GOPATH` 环境变量指示的目录的 `pkg/mod` 子目录中。模块的给定版本的下载内容在“需要”该版本的所有其他模块之间共享，因此“go”命令将这些文件和目录标记为只读。要删除所有下载的模块，您可以将 `-modcache` 标志传递给 `go clean`：

```
$ go clean -modcache
$
```

## Testing

## 测试

Go has a lightweight test framework composed of the `go test` command and the `testing` package.

Go 有一个轻量级的测试框架，由 `go test` 命令和 `testing` 包组成。

You write a test by creating a file with a name ending in `_test.go` that contains functions named `TestXXX` with signature `func (t *testing.T)`. The test framework runs each such function; if the function calls a failure function such as `t.Error` or `t.Fail`, the test is considered to have failed.

您可以通过创建一个名称以“_test.go”结尾的文件来编写测试，该文件包含名为“TestXXX”且签名为“func (t *testing.T)”的函数。测试框架运行每个这样的功能；如果函数调用了失败函数，例如 `t.Error` 或 `t.Fail`，则认为测试失败。

Add a test to the `morestrings` package by creating the file `$HOME/hello/morestrings/reverse_test.go` containing the following Go code.

通过创建包含以下 Go 代码的文件 `$HOME/hello/morestrings/reverse_test.go`，向 `morestrings` 包添加测试。

```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

Then run the test with `go test`:

然后使用 `go test` 运行测试：

```
$ cd $HOME/hello/morestrings
$ go test
PASS
ok      example/user/hello/morestrings 0.165s
$
```

Run `go help test` and see the [testing package documentation](https://golang.org/pkg/testing/) for more detail.

运行 `go help test` 并查看 [测试包文档](https://golang.org/pkg/testing/) 以获取更多详细信息。

## What's next

##  下一步是什么

Subscribe to the [golang-announce](https://groups.google.com/group/golang-announce) mailing list to be notified when a new stable version of Go is released.

订阅 [golang-announce](https://groups.google.com/group/golang-announce) 邮件列表，以便在发布新的稳定版 Go 时收到通知。

See [Effective Go](https://golang.org/doc/effective_go.html) for tips on writing clear, idiomatic Go code.

有关编写清晰、惯用的 Go 代码的提示，请参阅 [Effective Go](https://golang.org/doc/effective_go.html)。

Take  [A Tour of Go](https://tour.golang.org/) to learn the language proper.

参加 [A Tour of Go](https://tour.golang.org/) 学习正确的语言。

Visit the [documentation page](https://golang.org/doc/#articles) for a set of in-depth articles about the Go language and its libraries and tools.

访问 [文档页面](https://golang.org/doc/#articles) 获取一组关于 Go 语言及其库和工具的深入文章。

## Getting help

## 获得帮助

For real-time help, ask the helpful gophers in the community-run [gophers Slack server](https://gophers.slack.com/messages/general/) (grab an invite [here](https://invite.slack.golangbridge.org/)).

如需实时帮助，请在社区运行的 [gophers Slack 服务器](https://gophers.slack.com/messages/general/) 中询问乐于助人的 gophers（获取邀请 [此处](https://invite.slack.golangbridge.org/))。

The official mailing list for discussion of the Go language is [Go Nuts](https://groups.google.com/group/golang-nuts). 

官方讨论 Go 语言的邮件列表是 [Go Nuts](https://groups.google.com/group/golang-nuts)。

