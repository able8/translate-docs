# Writing your first kubectl plugin with Go

# 用 Go 编写你的第一个 kubectl 插件

https://bmuschko.com/blog/writing-your-first-kubectl-plugin/

https://bmuschko.com/blog/writing-your-first-kubectl-plugin/

June 20, 2019

2019 年 6 月 20 日

The `kubectl` command line tool offers a powerful, client-side mechanism for communicating with the Kubernetes API server. `kubectl` comes with a lot of fine-grained commands, subcommands and options. Getting familiar with all of these options takes practice and time.

`kubectl` 命令行工具提供了一种强大的客户端机制，用于与 Kubernetes API 服务器进行通信。 `kubectl` 带有许多细粒度的命令、子命令和选项。熟悉所有这些选项需要练习和时间。

Despite its wide array of options, you sometimes wish that `kubectl` would provide higher-level functions or would simplify specific  operations. For example, you may have to shell into a running container  to inspect the environment in interactive mode. That’s easy to achieve  with the command `kubectl exec mypod -it --namespace=abc  — /bin/sh`, however, there are various options you have to remember. Wouldn’t it be easier if there would be a simplied command say `kubectl iexec mypod`? That’s where [kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) come in. You can write your own commands and the behavior they should expose.

尽管有多种选择，但有时您还是希望 `kubectl` 能够提供更高级别的功能或简化特定操作。例如，您可能需要 shell 进入正在运行的容器以交互模式检查环境。使用命令 `kubectl exec mypod -it --namespace=abc — /bin/sh` 很容易实现，但是，您必须记住各种选项。如果有一个简单的命令说 `kubectl iexec mypod` 不是更容易吗？这就是 [kubectl 插件](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) 的用武之地。您可以编写自己的命令以及它们应该公开的行为。

In this blog post, I want to show how to write a simple `kubectl` plugin that leverages the Go client to communicate with the API server. You will learn how to stand up the Go project, declare the dependency  on the Go client library with the help of Go Modules and write the  source code. You will be able to find the [source code on GitHub](https://github.com/bmuschko/kubectl-server-version) if you’d like to have a look at the details.

在这篇博文中，我想展示如何编写一个简单的 `kubectl` 插件，该插件利用 Go 客户端与 API 服务器进行通信。您将学习如何建立 Go 项目，在 Go Modules 的帮助下声明对 Go 客户端库的依赖并编写源代码。如果您想查看详细信息，您将能够找到 [GitHub 上的源代码](https://github.com/bmuschko/kubectl-server-version)。

> The Go library isn’t the only option for writing a `kubectl` plugin. Kubernetes supports [other languages](https://kubernetes.io/docs/reference/using-api/client-libraries/) though they won’t be subject of this blog post. I personally haven’t  tried any of the other client libraries. I would think that the Go  library is likely the most well-maintained as Kubernetes itself is  written in Go.

> Go 库并不是编写 `kubectl` 插件的唯一选择。 Kubernetes 支持 [其他语言](https://kubernetes.io/docs/reference/using-api/client-libraries/)，尽管它们不是本博文的主题。我个人还没有尝试过任何其他客户端库。我认为 Go 库可能是维护得最好的，因为 Kubernetes 本身是用 Go 编写的。

## Defining the plugin functionality

## 定义插件功能

The purpose of our plugin is to render the Kubernetes server version  on the command line. We will want to make this functionality available  with the command `server-version`. The result should look similar to the console output below.

我们插件的目的是在命令行上呈现 Kubernetes 服务器版本。我们希望通过命令 `server-version` 来提供这个功能。结果应该类似于下面的控制台输出。

```bash
$ kubectl server-version
Hello from Kubernetes server with version v1.10.11!
```

Furthermore, it’s a good idea to provide a subcommand that renders  the version of the plugin so that the end user knows which feature set  to expect. It will help tremendously with the upgrade process if you  know which plugin version has been installed on your machine already. The following console output indicates that the plugin with version `0.2.0` is used.

此外，最好提供一个子命令来呈现插件的版本，以便最终用户知道期望的功能集。如果您知道您的机器上已经安装了哪个插件版本，它将极大地帮助升级过程。以下控制台输出表明使用了版本为 0.2.0 的插件。

```bash
$ kubectl server-version version
kubectl server-version v0.2.0
```

We defined the commands we want to expose and their functionality. Let’s start by setting up the project structure.

我们定义了要公开的命令及其功能。让我们从设置项目结构开始。

## Setting up the project

## 设置项目

There’s really nothing special about setting up a Go project that implements a `kubectl` plugin. You’d follow the same conventions as you usually would for any other Go project.

设置一个实现 `kubectl` 插件的 Go 项目真的没有什么特别之处。您将遵循与任何其他 Go 项目通常相同的约定。

We’ll start by adding a `main.go` file, the entrypoint into the application. Next, we’ll generate the Go Module with `go mod init github.com/bmuschko/kubectl-server-version`. You’ll end up with a `go.mod` and `go.sum` file that we will populate later. The package `cmd` is going to contain all relevant CLI command handling. Moreover, I  personally like to provide a license for the project and a README file  that explains how to use the plugin. In the end, you should end up with a project layout shown below.

我们将首先添加一个 main.go 文件，这是应用程序的入口点。接下来，我们将使用 `go mod init github.com/bmuschko/kubectl-server-version` 生成 Go 模块。你最终会得到一个 `go.mod` 和 `go.sum` 文件，我们稍后会填充它们。包 `cmd` 将包含所有相关的 CLI 命令处理。此外，我个人喜欢为项目提供许可证和解释如何使用插件的 README 文件。最后，您应该得到如下所示的项目布局。

```
.
├── LICENSE
├── README.adoc
├── cmd
├── go.mod
├── go.sum
└── main.go
```

In the next section, we’ll add the necessary dependencies.

在下一节中，我们将添加必要的依赖项。

## Declaring the dependencies 

## 声明依赖

The initial structure of the `go.mod` file has already been set up when we ran the `go init` command. Remember that you have to set the environment variable `GO111MODULES=on` to enable Go Modules before adding any dependencies. To get started with the Go client simply run the command `go get k8s.io/client-go`. There are other portions to the Go client library that you might need e.g. `k8s.io/api`, `k8s.io/apimachinery`. You can retrieve them with the same Go command. Have a look at the [documentation](https://github.com/kubernetes/client-go/) for more information.

当我们运行 `go init` 命令时，`go.mod` 文件的初始结构已经建立。请记住，在添加任何依赖项之前，您必须设置环境变量 `GO111MODULES=on` 以启用 Go Modules。要开始使用 Go 客户端，只需运行命令 `go get k8s.io/client-go`。您可能需要 Go 客户端库的其他部分，例如`k8s.io/api`、`k8s.io/apimachinery`。您可以使用相同的 Go 命令检索它们。查看 [文档](https://github.com/kubernetes/client-go/) 了解更多信息。

The Go Modules file below shows my final list of dependencies. As you can see, we are building the project with Go 1.12. I also added the  dependency `github.com/spf13/cobra` for exposing the CLI commands and `github.com/stretchr/testify` to help with simplifying test code.

下面的 Go Modules 文件显示了我的最终依赖列表。如您所见，我们正在使用 Go 1.12 构建项目。我还添加了依赖 `github.com/spf13/cobra` 来公开 CLI 命令和 `github.com/stretchr/testify` 以帮助简化测试代码。

*go.mod*

*go.mod*

```
module github.com/bmuschko/kubectl-server-version

go 1.12

require (
    github.com/evanphx/json-patch v4.5.0+incompatible // indirect
    github.com/gogo/protobuf v1.2.1 // indirect
    github.com/golang/protobuf v1.3.1 // indirect
    github.com/google/btree v1.0.0 // indirect
    github.com/google/gofuzz v1.0.0 // indirect
    github.com/googleapis/gnostic v0.3.0 // indirect
    github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
    github.com/imdario/mergo v0.3.7 // indirect
    github.com/json-iterator/go v1.1.6 // indirect
    github.com/modern-go/reflect2 v1.0.1 // indirect
    github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
    github.com/pkg/errors v0.8.1 // indirect
    github.com/spf13/cobra v0.0.4
    github.com/stretchr/testify v1.3.0
    golang.org/x/net v0.0.0-20190619014844-b5b0513f8c1b // indirect
    golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
    golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
    gopkg.in/inf.v0 v0.9.1 // indirect
    k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b // indirect
    k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d // indirect
    k8s.io/cli-runtime v0.0.0-20190409023024-d644b00f3b79
    k8s.io/client-go v11.0.0+incompatible
    k8s.io/klog v0.3.3 // indirect
    k8s.io/kube-openapi v0.0.0-20190603182131-db7b694dc208 // indirect
    k8s.io/utils v0.0.0-20190607212802-c55fbcfc754a // indirect
    sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
    sigs.k8s.io/yaml v1.1.0 // indirect
)
```

Next up, we’ll implement the plugin functionality.

接下来，我们将实现插件功能。

## Implementing the plugin functionality

## 实现插件功能

Not surprisingly, we will start by filling the `main.go` with life. The implementation logic is straightforward. First, we’ll create a new command by calling the function `NewServerVersionCommand`. Standard input, standard output and standard error have been wrapped with `genericclioptions.IOStreams` and passed in as parameter. Next, we’ll execute the newly created command and handle a potential error case.

毫不奇怪，我们将从用生命填充 `main.go` 开始。实现逻辑很简单。首先，我们将通过调用函数“NewServerVersionCommand”来创建一个新命令。标准输入、标准输出和标准错误已经用`genericclioptions.IOStreams` 包裹并作为参数传入。接下来，我们将执行新创建的命令并处理潜在的错误情况。

*main.go*

*main.go*

```go
package main

import (
    "github.com/bmuschko/kubectl-server-version/cmd"
    "k8s.io/cli-runtime/pkg/genericclioptions"
    "os"
)

var version = "undefined"

func main() {
    cmd.SetVersion(version)

    serverVersionCmd := cmd.NewServerVersionCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
    if err := serverVersionCmd.Execute();err != nil {
        os.Exit(1)
    }
}
```

The value of the variable `version` can be injected at build-time and represents the plugin version. We’ll pass on the value to the package `cmd` for later use in the `version` subcommand.

变量 `version` 的值可以在构建时注入，代表插件版本。我们会将值传递给包 `cmd` 以供稍后在 `version` 子命令中使用。

Let’s have a look at the meat of the plugin functionality, the  command that communicates with the API server and renders the returned  Kubernetes server version. I split up the code into multiple,  digestiable pieces. The code block below uses the [Cobra library](https://github.com/spf13/cobra) to create the new command. Apart from the Kubernetes package imports,  there’s nothing special about this implementation. You can easily  imagine similar code as part of any other CLI tool written in Go.

让我们来看看插件功能的核心，即与 API 服务器通信并呈现返回的 Kubernetes 服务器版本的命令。我将代码分成多个可消化的部分。下面的代码块使用 [Cobra 库](https://github.com/spf13/cobra) 创建新命令。除了 Kubernetes 包导入之外，这个实现没有什么特别之处。您可以轻松地将类似的代码想象为任何其他用 Go 编写的 CLI 工具的一部分。

*cmd/server_version.go*

*cmd/server_version.go*

```go
import (
    "errors"
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "k8s.io/cli-runtime/pkg/genericclioptions"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

type serverVersionCmd struct {
    out io.Writer
}

// NewServerVersionCommand creates the command for rendering the Kubernetes server version.
func NewServerVersionCommand(streams genericclioptions.IOStreams) *cobra.Command {
    helloWorldCmd := &serverVersionCmd{
        out: streams.Out,
    }

    cmd := &cobra.Command{
        Use:          "server-version",
        Short:        "Prints Kubernetes server version",
        SilenceUsage: true,
        RunE: func(c *cobra.Command, args []string) error {
            if len(args) != 0 {
                return errors.New("this command does not accept arguments")
            }
            return helloWorldCmd.run()
        },
    }

    cmd.AddCommand(newVersionCmd(streams.Out))
    return cmd
}

func (sv *serverVersionCmd) run() error {
    serverVersion, err := getServerVersion()
    if err != nil {
        return err
    }

    _, err = fmt.Fprintf(sv.out, "Hello from Kubernetes server with version %s!\n", serverVersion)
    if err != nil {
        return err
    }
    return nil
}
```

The most important portion of this file is the call to the function `getServerVersion()`. In here, we are actually interacting with the Go client.

这个文件最重要的部分是调用函数`getServerVersion()`。在这里，我们实际上是在与 Go 客户端进行交互。

The "client" instance is called a `clientset` and is represented by the struct `kubernetes.Clientset`. For the creation of the the struct, you will have to provide additional client configuration in the form of a `restclient.Config`. Once created, the `clientset` gives access to the whole Kubernetes API and all relevant operations  for querying, creating, updating and deleting Kubernetes objects.

“客户端”实例称为“clientset”，由结构“kubernetes.Clientset”表示。为了创建结构体，您必须以 `restclient.Config` 的形式提供额外的客户端配置。一旦创建，`clientset` 就可以访问整个 Kubernetes API 以及用于查询、创建、更新和删除 Kubernetes 对象的所有相关操作。

The function below retrieves the `DiscoveryClient` which can discover server-supported API groups, versions and resources. For example, the function `ServerVersion()` returns the Kubernetes server version.

下面的函数检索“DiscoveryClient”，它可以发现服务器支持的 API 组、版本和资源。例如，函数`ServerVersion()` 返回 Kubernetes 服务器版本。

*cmd/server_version.go*

*cmd/server_version.go*

```go
func getServerVersion() (string, error) {
    loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
    configOverrides := &clientcmd.ConfigOverrides{}
    kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

    config, err := kubeConfig.ClientConfig()
    if err != nil {
        return "", err
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return "", err
    }

    sv, err := clientset.Discovery().ServerVersion()
    if err != nil {
        return "", err
    }

    return sv.String(), nil
}
```

That’s really all there is to the implementation. The most important  aspect is to get familiar with Go client API. Everything else is really  just plain old Go programming.

这就是实现的全部内容。最重要的方面是熟悉 Go 客户端 API。其他一切都只是普通的 Go 编程。

## Installing and using the plugin

## 安装和使用插件

There are a couple of requirements for installing `kubectl` plugins on a machine.

在机器上安装 `kubectl` 插件有几个要求。

1. A plugin can be installed as bash script or as binary.
2. The plugin name has to have the prefix `kubectl-`.
3. The portion of the plugin name after `kubectl-` represents the command name.
4. A dash character in the command name has to be changed to an underscore.

1. 插件可以安装为 bash 脚本或二进制文件。
2.插件名称必须有前缀`kubectl-`。
3.插件名中`kubectl-`后面的部分代表命令名。
4. 命令名称中的破折号字符必须更改为下划线。

In this post, we didn’t even talk about the option of writing a plugin as bash script. For more information, see the [Kubernetes documentation](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#example-plugin). Our plugin has been written in Go so we’ll have to build the binary first. The easiest option is to use the `go build` command. The "installation process" isn’t complicated at all. You simply copy the binary into the `PATH`, as described [here](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#installing-kubectl-plugins).

在这篇文章中，我们甚至没有讨论将插件编写为 bash 脚本的选项。有关更多信息，请参阅 [Kubernetes 文档](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#example-plugin)。我们的插件是用 Go 编写的，所以我们必须先构建二进制文件。最简单的选择是使用 `go build` 命令。 “安装过程”一点也不复杂。您只需将二进制文件复制到 `PATH` 中，如 [此处](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#installing-kubectl-plugins) 所述。

The following commands build the binary of the plugin and copy it into the `PATH`. Under MacOSX the `PATH` is represented by `/usr/local/bin`.

以下命令构建插件的二进制文件并将其复制到“PATH”中。在 MacOSX 下，`PATH` 由 `/usr/local/bin` 表示。

```bash
$ go build -o kubectl-server_version
$ cp kubectl-server_version /usr/local/bin
$ kubectl plugin list
The following kubectl-compatible plugins are available:

/usr/local/bin/kubectl-server_version
```

Awesome, the plugin has been installed. We can now go ahead and  execute it as described earlier. As you can see, the development and  installation process is not very burdensome. But what if you wanted to  share the plugin with the team or the Kubernetes community? Let’s talk  about distribution the plugin in the next section.

太好了，插件已经安装了。我们现在可以继续执行它，如前所述。可以看到，开发安装过程并不是很繁琐。但是，如果您想与团队或 Kubernetes 社区共享插件怎么办？让我们在下一节中讨论分发插件。

## Distributing the plugin on GitHub Releases

## 在 GitHub Releases 上分发插件

There are various options for distributing the binary of your plugin. In this post, we’ll discuss the distribution of the plugin on [GitHub Releases](https://help.github.com/en/articles/creating-releases).

分发插件的二进制文件有多种选择。在这篇文章中，我们将在 [GitHub Releases](https://help.github.com/en/articles/creating-releases) 上讨论插件的分发。

> Another popular option is to use the package manager for `kubectl` plugins called [krew](https://github.com/kubernetes-sigs/krew). It’s your best option when it comes to easy of use and making the  plugin available to the world. I will leave this topic for another blog  post. 

> 另一个流行的选择是使用名为 [krew](https://github.com/kubernetes-sigs/krew) 的 `kubectl` 插件的包管理器。在易于使用和向全世界提供插件方面，这是您的最佳选择。我将把这个话题留给另一篇博文。

My tool of choice for publishing binaries to GitHub Releases is [GoReleaser](https://github.com/goreleaser/goreleaser). It helps with defining, building and publishing the binaries you want  to release with a given Git tag. GoReleaser reads the configuration for  the binaries in a file named `.goreleaser.yml`. You can find [an example](https://github.com/bmuschko/kubectl-server-version/blob/master/.goreleaser.yml) for such a file in the plugin repository. Once a version has been  released, any consumer can download the binary for a target platform and install the plugin directly into the `PATH`.

我选择将二进制文件发布到 GitHub Releases 的工具是 [GoReleaser](https://github.com/goreleaser/goreleaser)。它有助于定义、构建和发布您想要使用给定 Git 标签发布的二进制文件。 GoReleaser 在名为“.goreleaser.yml”的文件中读取二进制文件的配置。您可以在插件存储库中找到此类文件的 [示例](https://github.com/bmuschko/kubectl-server-version/blob/master/.goreleaser.yml)。版本发布后，任何消费者都可以下载目标平台的二进制文件并将插件直接安装到“PATH”中。

## Conclusion

##  结论

The blog post taught you how to get started with writing a `kubectl` plugin using the Go language. We created a new project from scratch and took it all the way from inception to distribution. I hope I could  inspire you to write your own plugins. In the course of the next couple  of months, I intend to write more blog posts on this topic so stay  tuned! In the meantime, check out the all the [kubectl plugins that have been published by the community](https://github.com/topics/kubectl-plugins). Maybe you will find the functionality you’ve been missing for so long from the standard `kubectl` command. 

这篇博文教你如何开始使用 Go 语言编写 `kubectl` 插件。我们从头开始创建了一个新项目，并从开始到分发一路进行。我希望我能激励你编写自己的插件。在接下来的几个月里，我打算写更多关于这个主题的博客文章，敬请期待！同时，查看所有[社区已发布的kubectl插件](https://github.com/topics/kubectl-plugins)。也许你会从标准的 `kubectl` 命令中找到你长期以来一直缺少的功能。

