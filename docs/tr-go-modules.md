# Package management in Go

# Go 中的包管理

## Brief overview of package management in Go — pre and post Go modules

## 简要概述 Go 中的包管理——前和后 Go 模块

By Jai on August 26, 2019

Package management is one of the things Go has always missed. One of the major drawbacks of the previous (pre 1.11) `go get` was lack of support for managing dependency versions and enabling  reproducible builds. The community has developed package managers and  tools like Glide, dep and [many others](https://github.com/golang/go/wiki/PackageManagementTools) serving as de-facto solutions for versioning dependencies.

包管理是 Go 一直错过的事情之一。之前（1.11 之前）`go get` 的主要缺点之一是缺乏对管理依赖项版本和启用可重复构建的支持。社区已经开发了包管理器和工具，如 Glide、dep 和 [许多其他](https://github.com/golang/go/wiki/PackageManagementTools) 作为版本控制依赖项的事实上的解决方案。

> “I use go get for production builds.” — said no one ever.

> “我使用 go get 进行生产构建。” - 说从来没有人。

Go’s implementation of package management traces its origins back to Google  (which has a giant monolithic repository for all their source code). Let’s break down on what’s wrong with ‘pre - go module’ package  management tooling.

Go 的包管理实现可以追溯到谷歌（谷歌拥有一个巨大的单体存储库，用于存放所有源代码）。让我们分解一下“pre-go module”包管理工具有什么问题。

1. Versioning dependencies
2. Vendoring dependencies
3. The necessity of `GOPATH`

1. 版本依赖
2. 供应商依赖
3. `GOPATH`的必要性

## Versioning dependencies

## 版本依赖

`go get` by default didn’t support module versioning. The idea behind the first  version of go’s package management was — no need for module versioning,  no need for 3rd-party module repositories, you build everything from  your current branch.

`go get` 默认不支持模块版本控制。 go 包管理的第一个版本背后的想法是 - 不需要模块版本控制，不需要 3rd-party 模块存储库，您可以从当前分支构建所有内容。

Pre Go 1.11, adding a dependency meant cloning that dependency’s source code repo in your `GOPATH`. That was about it. There was no concept of versions. Rather, it always  pointed to the current master branch at the time of cloning. Another  major issue cropped up when different projects needed different versions of a dependency — which wasn’t possible either.

在 Go 1.11 之前，添加依赖项意味着在您的“GOPATH”中克隆该依赖项的源代码存储库。就是这样。没有版本的概念。相反，它总是在克隆时指向当前的主分支。当不同的项目需要不同版本的依赖项时，另一个主要问题出现了——这也是不可能的。

## Vendoring dependencies

## 供应商依赖

Package vendoring is commonly referred to as the case where dependent packages  are stored in the same place as your project. That usually means your  dependencies are checked into your source management system, such as  Git.

Package vendoring 通常被称为依赖包与您的项目存储在同一位置的情况。这通常意味着您的依赖项会被检入您的源代码管理系统，例如 Git。

Consider this case — A uses dependency B, which uses a  feature of dependency C introduced in version 1.5 of C, B must be able  to ensure that A’s build uses C 1.5 or later. Pre Go 1.5, there was no  mechanism for carrying dependency code alongside commands without  rewriting import paths.

考虑这种情况——A 使用依赖 B，它使用 C 的 1.5 版中引入的依赖 C 的特性，B 必须能够确保 A 的构建使用 C 1.5 或更高版本。在 Go 1.5 之前，没有在不重写导入路径的情况下将依赖代码与命令一起携带的机制。

## Necessity of `GOPATH`

## `GOPATH` 的必要性

`GOPATH` exists for two main reasons:

`GOPATH` 的存在有两个主要原因：

1. In Go, the `import` declaration references a package via its fully qualified import path. `GOPATH` exist so that from any directory inside `GOPATH/src` the go tool can compute the absolute import path of the package in question.
2. A location to store dependencies fetched by `go get.`

1. 在 Go 中，`import` 声明通过其完全限定的导入路径引用包。 `GOPATH` 存在以便从 `GOPATH/src` 中的任何目录，go 工具可以计算相关包的绝对导入路径。
2. 一个存放`go get`获取的依赖项的位置。

What’s wrong with this?

这有什么问题？

1. `GOPATH` doesn’t allow checking out the source of a project in a directory of choice like they are used to with other languages.
2. Additionally, `GOPATH` does not let the developer have more than one copy of a project (or its dependencies) checked out at the same time.

1. `GOPATH` 不允许像其他语言那样在选择的目录中检出项目的源代码。
2. 此外，`GOPATH` 不允许开发人员同时签出一个以上的项目（或其依赖项）副本。

## Introducing Go Modules

## 介绍 Go 模块

Go 1.11 introduces preliminary support for Go modules. From Go Wiki,

Go 1.11 引入了对 Go 模块的初步支持。从 Go 维基，

> A *module* is a collection of related Go packages that are versioned together as a single unit. Modules record precise dependency requirements and create  reproducible builds.

> *module* 是相关 Go 包的集合，它们作为一个单元一起进行版本控制。模块记录精确的依赖需求并创建可重现的构建。

Go modules brings three important features built-in,

Go 模块带来了三个重要的内置特性，

1. `go.mod` file similar to `package.json` or `Pipfile`.
2. A machine-generated transitive dependency description - `go.sum`.
3. No more `GOPATH` limitation. Modules can be in any path.

1. `go.mod` 文件类似于 `package.json` 或 `Pipfile`。
5. 机器生成的传递依赖描述 - `go.sum`。
6. 不再有`GOPATH`限制。模块可以在任何路径中。

```bash
$ go help mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'.For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

    go mod <command> [arguments]

The commands are:

    download    download modules to local cache
    edit        edit go.mod from tools or scripts
    graph       print module requirement graph
    init        initialize new module in current directory
    tidy        add missing and remove unused modules
    vendor      make vendored copy of dependencies
    verify      verify dependencies have expected content
    why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

Relevant [discussion thread](https://groups.google.com/forum/#!topic/golang-dev/a5PqQuBljF4).

相关[讨论线程](https://groups.google.com/forum/#!topic/golang-dev/a5PqQuBljF4)。

## Migrating to Go Modules 

## 迁移到 Go 模块

To use Go modules, update Go to version `>= 1.11`. Since `GOPATH` is going away, one can activate module support in one of these two ways:

要使用 Go 模块，请将 Go 更新到版本 `>= 1.11`。由于 GOPATH 即将消失，您可以通过以下两种方式之一激活模块支持：

- Invoke the `go` command in a directory outside of the `GOPATH/src` tree, with a valid `go.mod` file in the current directory.
- Go modules don’t work if source is under `GOPATH`. To override this behaviour, invoke the `go` command with `GO111MODULE=on` environment variable set.

- 在 `GOPATH/src` 树之外的目录中调用 `go` 命令，并在当前目录中使用有效的 `go.mod` 文件。
- 如果源代码在 `GOPATH` 下，Go 模块将不起作用。要覆盖此行为，请使用“GO111MODULE=on”环境变量集调用“go”命令。

Let’s start porting by following these simple steps:

让我们按照以下简单步骤开始移植：

- As `GOPATH` isn’t necessary anymore, move the module out of `GOPATH`.
- From the project root, create the initial module definition - `go mod init github.com/username/repository`. The best part is, `go mod` automatically converts dependencies from existing package managers like `dep`, `Gopkg`, `glide` and [six others](https://tip.golang.org/pkg/cmd/go/internal/modconv/?m=all#pkg-variables). This will create a file called `go.mod` with the module name and dependencies with its versions.

- 由于不再需要`GOPATH`，将模块移出`GOPATH`。
- 从项目根目录，创建初始模块定义 - `go mod init github.com/username/repository`。最好的部分是，`go mod` 会自动转换现有包管理器的依赖项，如 `dep`、`Gopkg`、`glide` 和 [其他六个](https://tip.golang.org/pkg/cmd/go/内部/modconv/?m=all#pkg-variables)。这将创建一个名为 `go.mod` 的文件，其中包含模块名称和依赖项及其版本。

```bash
$ cat go.mod
module github.com/deepsourcelabs/cli

go 1.12

require (
    github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e
    github.com/getsentry/raven-go v0.2.0
    github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9
```

- Run `go build` to create a `go.sum` file which contains the expected cryptographic checksums of the content of specific module versions. This is to ensure that future downloads of these modules retrieve the same bits as the first download. Note that  go.sum is not a lock file.

- 运行 `go build` 以创建一个 `go.sum` 文件，其中包含特定模块版本内容的预期加密校验和。这是为了确保这些模块的未来下载检索与第一次下载相同的位。请注意， go.sum 不是锁定文件。

```bash
$ cat go.sum
github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e h1:9574pc8MX6rF/QyO14SPHhM5KKIOo9fkb/1ifuYMTKU=
github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e/go.mod h1:GJKEexRPVJrBSOjoqN5VNOIKJ5Q3RViH6eu3puDRwx4=
github.com/getsentry/raven-go v0.2.0 h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=
github.com/getsentry/raven-go v0.2.0/go.mod h1:KungGk8q33+aIAZUIVWZDr2OfAEBsO49PX4NzFV5kcQ=
github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9 h1:dIsTcVF0w9viTLHXUEkDI7cXITMe+M/MRRM2MwisVow=
github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
```

> Note on versioning: To maintain backward compatibility, if the module is version v2 or higher, the major version of the module *must* be included as a `/vN` at the end of the module paths used in `go.mod ` files (eg, `module github.com/username/repository/v2`

> 版本注意事项：为了保持向后兼容性，如果模块是 v2 或更高版本，模块的主要版本*必须*作为 `/vN` 包含在 `go.mod ` 中使用的模块路径的末尾文件（例如，`module github.com/username/repository/v2`

## Everyday commands

## 日常命令

### List dependencies

### 列出依赖项

`go list -m all` lists the current module and all its dependencies.

`go list -m all` 列出当前模块及其所有依赖项。

```bash
$ go list -m all
github.com/deepsourcelabs/cli
github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e
github.com/getsentry/raven-go v0.2.0
github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9
```

> In the `go list` output, the current module, also known as the *main module*, is always the first line, followed by dependencies sorted by module path.

> 在 `go list` 输出中，当前模块，也称为 *main 模块 *，始终位于第一行，其后是按模块路径排序的依赖项。

### List available versions of a package

### 列出包的可用版本

`go list -m -versions github.com/username/repository` lists available versions of a package.

`go list -m -versions github.com/username/repository` 列出包的可用版本。

```bash
$ go list -m -versions github.com/getsentry/raven-go
github.com/getsentry/raven-go v0.1.0 v0.1.1 v0.1.2 v0.2.0
```

### Add a dependency

### 添加依赖

Adding a dependency is implicit. After importing a dependency in code, running `go build` or `go test` command gets the latest version of the module and adds it to `go.mod` file. If you would like to add a dependency explicitly, run`go get github.com/username/repository`.

添加依赖项是隐式的。在代码中导入依赖项后，运行 `go build` 或 `go test` 命令获取模块的最新版本并将其添加到 `go.mod` 文件中。如果您想显式添加依赖项，请运行`go get github.com/username/repository`。

### Upgrade/downgrade a dependency

### 升级/降级依赖项

`go get github.com/username/repository@vx.x.x` downloads and sets the specific version of the dependency and updates `go.mod` file.

`go get github.com/username/repository@vx.x.x` 下载并设置依赖的特定版本并更新 `go.mod` 文件。

```bash
$ go get github.com/getsentry/raven-go@v0.1.2
go: finding github.com/getsentry/raven-go v0.1.2
go: downloading github.com/getsentry/raven-go v0.1.2
go: extracting github.com/getsentry/raven-go v0.1.2
$ cat go.mod
module github.com/deepsourcelabs/marvin-go

go 1.12

require (
    github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e
    github.com/getsentry/raven-go v0.1.2
    github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9
)
$ cat go.sum
github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e h1:9574pc8MX6rF/QyO14SPHhM5KKIOo9fkb/1ifuYMTKU=
github.com/certifi/gocertifi v0.0.0-20190410005359-59a85de7f35e/go.mod h1:GJKEexRPVJrBSOjoqN5VNOIKJ5Q3RViH6eu3puDRwx4=
github.com/getsentry/raven-go v0.1.2 h1:4V0z512S5mZXiBvmW2RbuZBSIY1sEdMNsPjpx2zwtSE=
github.com/getsentry/raven-go v0.1.2/go.mod h1:KungGk8q33+aIAZUIVWZDr2OfAEBsO49PX4NzFV5kcQ=
github.com/getsentry/raven-go v0.2.0 h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=
github.com/getsentry/raven-go v0.2.0/go.mod h1:KungGk8q33+aIAZUIVWZDr2OfAEBsO49PX4NzFV5kcQ=
github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9 h1:dIsTcVF0w9viTLHXUEkDI7cXITMe+M/MRRM2MwisVow=
github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
```

### Vendoring dependencies

### 供应商依赖

When using modules, the go command completely ignores vendor directories. For backward compatibility with older versions of Go, or to ensure that  all files used for a build are stored together in a single file tree,  run`go mod vendor`.

使用模块时， go 命令完全忽略供应商目录。为了与旧版本的 Go 向后兼容，或确保用于构建的所有文件都存储在单个文件树中，请运行 `go mod vendor`。

This creates a directory named `vendor` in the root directory of the main module and stores all the packages from dependency modules there.

这会在主模块的根目录中创建一个名为“vendor”的目录，并将所有依赖模块的包存储在那里。

> Note: To build using the main module’s top-level vendor directory, run ‘go build -mod=vendor’.

> 注意：要使用主模块的顶级供应商目录进行构建，请运行“go build -mod=vendor”。

### Remove unused dependencies

### 删除未使用的依赖项

`go mod tidy` trims unused dependencies and updates `go.mod` file.

`go mod tidy` 修剪未使用的依赖项并更新 `go.mod` 文件。

## FAQs

## 常见问题

#### Is `GOPATH` not needed anymore?

#### 不再需要`GOPATH`了吗？

No.  Farewell `GOPATH`.

不。再见`GOPATH`。

#### Which version is pulled by default?

#### 默认拉取哪个版本？

The go.mod file and the go command more generally use semantic versions as  the standard form for describing module versions, so that versions can  be compared to determine which should be considered earlier or later  than another. A module version like `v1.2.3` is introduced by tagging a revision in the underlying source repository. Untagged  revisions can be referred to using a “pseudo-version” like `v0.0.0-yyyymmddhhmmss-abcdefabcdef`, where the time is the commit time in UTC and the final suffix is the prefix of the commit hash.

go.mod 文件和 go 命令更普遍地使用语义版本作为描述模块版本的标准形式，以便可以比较版本以确定哪个应该比另一个更早或更晚。通过在底层源存储库中标记修订版来引入像“v1.2.3”这样的模块版本。可以使用“v0.0.0-yyyymmddhhmmss-abcdefabcdef”之类的“伪版本”来引用未标记的修订，其中时间是 UTC 中的提交时间，最终后缀是提交哈希的前缀。

#### Should `go.sum` be checked into version control?

#### 应该将 `go.sum` 签入版本控制吗？

Yes.

是的。

About DeepSource

关于深源

DeepSource helps you automatically find and fix issues in your code during code  reviews, such as bug risks, anti-patterns, performance issues, and  security flaws. It takes less than 5 minutes to set up with your  Bitbucket, GitHub, or GitLab account. It works for Python, Go, Ruby, and JavaScript. 

DeepSource 帮助您在代码审查期间自动查找和修复代码中的问题，例如错误风险、反模式、性能问题和安全缺陷。设置您的 Bitbucket、GitHub 或 GitLab 帐户只需不到 5 分钟。它适用于 Python、Go、Ruby 和 JavaScript。

