# A tour of the Kubernetes source code

# Kubernetes 源代码之旅

From kubectl to API Server  从 kubectl 到 API 服务器

SaveLike By [Brad Topol](https://developer.ibm.com/profiles/btopol), Published June 21, 2017

https://developer.ibm.com/articles/a-tour-of-the-kubernetes-source-code/

* * *

## Overview

##  概述

Kubernetes continues to experience explosive growth and software developers that are able to understand and contribute to the Kubernetes code base are in high demand. Learning the Kubernetes code base is not easy. Kubernetes is written Go which is a fairly new programming language and it has a large amount of source code.

Kubernetes 继续经历爆炸式增长，能够理解并为 Kubernetes 代码库做出贡献的软件开发人员需求量很大。学习 Kubernetes 代码库并不容易。 Kubernetes 是用 Go 编写的，这是一种相当新的编程语言，它有大量的源代码。

In this multi-part series of articles I will dig in and explain key portions of the Kubernetes code base and also explain the techniques I have used to help me understand the code. My goal is to provide a set of articles that will enable software developers new to Kubernetes to more quickly learn the Kubernetes source code.

在这个多部分的系列文章中，我将深入研究并解释 Kubernetes 代码库的关键部分，并解释我用来帮助我理解代码的技术。我的目标是提供一组文章，让刚接触 Kubernetes 的软件开发人员能够更快地学习 Kubernetes 源代码。

In this first article, I will cover the flow through the code from running a simple kubectl command to sending a REST call to the API Server. Before using this article to dig into the Kubernetes code, I recommend you read an outstanding high level [overview](https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/) of the Kubernetes architecture by [Julia Evans](https://twitter.com/b0rk).

在第一篇文章中，我将介绍从运行简单的 kubectl 命令到向 API 服务器发送 REST 调用的整个代码流程。在使用本文深入研究 Kubernetes 代码之前，我建议您阅读 Kubernetes 架构的优秀高级 [overview](https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/)[朱莉娅·埃文斯](https://twitter.com/b0rk)。

## Running a basic kubectl command

## 运行一个基本的 kubectl 命令

The command line interface for Kubernetes is called kubectl. It is used for running commands against Kubernetes clusters. When attempting to learn the Kubernetes source code, the portion of the source code that implements the command line interface is a great place to start. The command we will use to trace through the source code is the `kubectl create -f` command which creates a resource from a file. The resource we are creating is a single replica pod with a basic `nginx` container image. The specification for this resource is shown below and I placed this specification in a file called `~/nginx_kube_example/nginx_pod.yaml`.

Kubernetes 的命令行界面称为 kubectl。它用于针对 Kubernetes 集群运行命令。在尝试学习 Kubernetes 源代码时，实现命令行界面的源代码部分是一个很好的起点。我们将用来跟踪源代码的命令是“kubectl create -f”命令，它从文件中创建资源。我们正在创建的资源是具有基本“nginx”容器映像的单个副本 pod。此资源的规范如下所示，我将此规范放在名为 `~/nginx_kube_example/nginx_pod.yaml` 的文件中。

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
name: nginx
spec:
replicas: 1
selector:
    app: nginx
template:
    metadata:
      name: nginx
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80

```

From a Kubernetes development environment we can invoke kubectl as shown in the figure below:

在 Kubernetes 开发环境中，我们可以调用 kubectl，如下图所示：

![screen capture of kubectl command to create a resource from a file](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_cmd_trimmed.jpg)

Now that we know what kubectl command we are running and how to run it, let's look at where we can find the implementation of this command in the Kubernetes source code.

现在我们知道我们正在运行什么 kubectl 命令以及如何运行它，让我们看看在 Kubernetes 源代码中可以找到该命令的实现。

## Locating the implementation of kubectl commands in the Kubernetes source code

## 在Kubernetes源码中定位kubectl命令的实现

The entry point for all the kubectl commands can be found in the [`github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd`](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd) folder. In this folder there is a name of a Go file that matches the name of the kubectl command that is implemented. For example, the `kubectl create` command has an initial entry point in a file named `create.go`. The folder and the example go files implementing the various commands are shown in the figure below.

所有 kubectl 命令的入口点都可以在 [`github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd`](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd) 文件夹。在此文件夹中，有一个 Go 文件的名称，该名称与已实现的 kubectl 命令的名称相匹配。例如，“kubectl create”命令在名为“create.go”的文件中有一个初始入口点。实现各种命令的文件夹和示例 go 文件如下图所示。

![screen capture of directory containing code entry points for all the kubectl commands](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_cmd_dir.jpg)

## Kubernetes loves the Cobra command framework 

## Kubernetes 喜欢 Cobra 命令框架

Kubernetes commands are implemented using the [Cobra command framework](https://github.com/spf13/cobra). Cobra provides a lot of great features for building command line interfaces and a basic overview of Cobra's capabilities can be found [here](https://blog.gopheracademy.com/advent-2014/introducing-cobra/). As shown in the Figure, one of the nice features of how Kubernetes utilizes Cobra is that it is very easy to locate which file implements each command line option. Furthermore, the Cobra structure puts the command usage message and command descriptions adjacent to the code that runs the command. This is shown in the figure and the [actual lines of code](https://github.com/kubernetes/kubernetes/blob/fd9a91e0b57face905c4225b8a6633b2ea9c832d/pkg/kubectl/cmd/create.go#L62-%2376). What's great about this structure is that you can go through and look at the descriptions for all the Kubernetes kubectl commands and then quickly jump to the code that implements the commands. As shown in lines 62-76 in the figure, the strings `Use`, `Short`, `Long`, and `Example` all hold information describing the command and `Run` points to a function that actually runs the command.

Kubernetes 命令是使用 [Cobra 命令框架](https://github.com/spf13/cobra) 实现的。 Cobra 提供了许多用于构建命令行界面的强大功能，并且可以在 [此处](https://blog.gopheracademy.com/advent-2014/introducing-cobra/) 找到 Cobra 功能的基本概述。如图所示，Kubernetes 如何利用 Cobra 的一个很好的特点是很容易找到哪个文件实现了每个命令行选项。此外，Cobra 结构将命令使用消息和命令描述与运行命令的代码相邻。如图和 [实际代码行数](https://github.com/kubernetes/kubernetes/blob/fd9a91e0b57face905c4225b8a6633b2ea9c832d/pkg/kubectl/cmd/create.go#L62-%2376)所示。这种结构的好处在于，您可以浏览并查看所有 Kubernetes kubectl 命令的描述，然后快速跳转到实现这些命令的代码。如图中第 62-76 行所示，字符串 `Use`、`Short`、`Long` 和 `Example` 都包含描述命令的信息，而 `Run` 指向实际运行命令的函数。

![screen capture of RunCreate function, which performs the bulk of the kubectl create command](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_NewCMDCreate.jpg)

The `RunCreate` function invoked on line 74 in the above figure is where the bulk of the kubectl create command is implemented. The implementation of this function can be found in the same `create.go` file. The figure below shows the `RunCreate` function. On line 132 I added a `fmt.Println` just to confirm this code was being called when I thought it would be called. In the [Compiling and running Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes) section below I show how you can speed up recompiling the Kubernetes code base when adding debugging statements solely to the kubectl source code.

上图中第 74 行调用的 `RunCreate` 函数是 kubectl create 命令的大部分实现的地方。这个函数的实现可以在同一个 `create.go` 文件中找到。下图显示了 `RunCreate` 函数。在第 132 行，我添加了一个`fmt.Println`，只是为了确认当我认为它会被调用时，这段代码被调用了。在下面的 [编译和运行 Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes) 部分中，我展示了在仅向 kubectl 源添加调试语句时如何加快重新编译 Kubernetes 代码库代码。

![screen capture of RunCreate function in create.go](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_RunCreate.jpg)

## Builders and visitors abound in Kubernetes 

## Kubernetes 中的建设者和访客比比皆是

The `resource.NewBuilder` function code shown on lines 133-140 is particularly intimidating to someone new to Go and Kubernetes. It's worth taking some time to explain this section of code in more detail. At a high level, what this code does is take the arguments and parameters from the command line and converts them into a list of resources. It's also responsible for creating a visitor construct that can be used to iterate across all the resources. The code is complex because it uses a variant of the [builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) where individual functions are each doing a separate portion of the data initialization. The functions `Schema`, `ContinueOnError`, `NamespaceParam`, `DefaultNamespace`, `FilenameParam`, `SelectorParam`, and `Flatten` all take in a pointer to a `Builder struct`, perform some form of modification on the ` Builder struct`, and then return the pointer to the `Builder struct` for the next method in the chain to use when it performs its modifications. All of these methods can be found in the `builder.go` file but I have included a few below so you can see how they work.

第 133-140 行显示的 `resource.NewBuilder` 函数代码对于刚接触 Go 和 Kubernetes 的人来说尤其令人生畏。值得花一些时间更详细地解释这部分代码。在高层次上，这段代码所做的是从命令行获取参数和参数，并将它们转换为资源列表。它还负责创建可用于遍历所有资源的访问者构造。代码很复杂，因为它使用 [builder 模式](https://en.wikipedia.org/wiki/Builder_pattern) 的变体，其中各个函数都在执行数据初始化的单独部分。 `Schema`、`ContinueOnError`、`NamespaceParam`、`DefaultNamespace`、`FilenameParam`、`SelectorParam` 和 `Flatten` 函数都接受一个指向 `Builder struct` 的指针，对`Builder struct` 执行某种形式的修改Builder struct`，然后返回指向 `Builder struct` 的指针，供链中的下一个方法在执行修改时使用。所有这些方法都可以在 `builder.go` 文件中找到，但我在下面包含了一些，以便您了解它们是如何工作的。

```go
func (b *Builder) Schema(schema validation.Schema) *Builder {
b.schema = schema
return b
}

func (b *Builder) ContinueOnError() *Builder {
b.continueOnError = true
return b
}

func (b *Builder) Flatten() *Builder {
b.flatten = true
return b
}
```



Once all the initializers have completed, the `resource.NewBuilder` function finally invokes a `Do` function. The `Do` function is a critical piece as it returns a `Result` object that will be used to drive the creation of our resource. The `Do` function also creates a `Visitor` object that can be used to traverse the list of resources that were associated with this invocation of `resource.NewBuilder`. The `Do` function implementation is shown below.

一旦所有的初始化程序都完成了，`resource.NewBuilder` 函数最终会调用 `Do` 函数。 `Do` 函数是一个关键部分，因为它返回一个 `Result` 对象，该对象将用于驱动我们的资源的创建。 `Do` 函数还创建了一个 `Visitor` 对象，该对象可用于遍历与此 `resource.NewBuilder` 调用相关联的资源列表。 `Do` 函数实现如下所示。

![screen capture of the Builder Do function that creates a DecoratedVisitor and returns a Result object](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/builder_Do.jpg)



As shown above on line 816, a new `DecoratedVisitor` is created and stored as part of the `Result` object that is returned by the `Builder Do` function. The `DecoratedVisitor` has a `Visit` function that will call the `Visitor` function that is passed into it. The implementation of this can be found at [github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/visitor.go#L306](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/visitor.go#L306) and is shown below.

如上面第 816 行所示，一个新的 `DecoratedVisitor` 被创建并存储为 `Builder Do` 函数返回的 `Result` 对象的一部分。 `DecoratedVisitor` 有一个 `Visit` 函数，它将调用传递给它的 `Visitor` 函数。可以在 [github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/visitor.go#L306](https://github.com/kubernetes/kubernetes/blob/6b52d8f13a4a4a769b403a04f812c99ed98815/visitor.go#L306)如下所示。

![screen capture of DecoratedVisitor Visit function that will eventually invoke createAndRefresh](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/visitor_NewDecoratedVisitor.jpg)

)

The `Result` object returned by the `Do` function has a `Visit` function that is used to invoke the `DecoratedVisitor Visit` function. This provides us a path from line 150 in the `RunCreate` function in `create.go` to eventually calling the anonymous function that is passed in on line 150 and contains the `createAndRefresh` function that will lead us to the code making a REST call to the API server. The implementation of the `Result Visit` function that is called on line 150 of `RunCreate` function in `create.go` is shown below.

`Do` 函数返回的 `Result` 对象有一个 `Visit` 函数，用于调用 `DecoratedVisitor Visit` 函数。这为我们提供了从 `create.go` 中的 `RunCreate` 函数中的第 150 行到最终调用在第 150 行传入并包含 `createAndRefresh` 函数的匿名函数的路径，该函数将引导我们生成 REST 的代码调用 API 服务器。 `create.go` 中`RunCreate` 函数的第 150 行调用的`Result Visit` 函数的实现如下所示。

![screen capture of Result Visit function that takes as parameter the function to invoke when visiting resources](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/result_Visit.jpg)

Now that we have seen how everything is connected through `Visit` functions and `DecoratedVisitor` classes, we see that the inline visitor function on line 150 below has a `createAndRefresh` function on line 165.

现在我们已经看到了一切是如何通过 `Visit` 函数和 `DecoratedVisitor` 类连接起来的，我们看到下面第 150 行的内联访问者函数在第 165 行有一个 `createAndRefresh` 函数。

![screen capture of enabling createAndRefresh to be invoked from a Result's Visitor object](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_RunCreate_callCreateAndRefresh.jpg) 

The `createAndRefresh` function invokes the `Resource.NewHelper`function found in `helper.go` and this function returns a new `Helper` object.

`createAndRefresh` 函数调用 `helper.go` 并且这个函数返回一个新的 `Helper` 对象。

![screen capture of the createAndRefresh function creating a new Helper object that performs the REST call to the Kubernetes APIServer](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_createAndRefresh.jpg)

Here is the [code](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go) that returns a new `Helper` object. It is actually pretty straightforward.

这是返回一个新的 `Helper` 对象的 [代码](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go)。它实际上非常简单。

```
func NewHelper(client RESTClient, mapping *meta.RESTMapping) *Helper {
return &Helper{
    Resource:        mapping.Resource,
    RESTClient:      client,
    Versioner:       mapping.MetadataAccessor,
    NamespaceScoped: mapping.Scope.Name() == meta.RESTScopeNameNamespace,
}
}

```

With the `Helper` created and its `Create` function invoked on line 217 of `createAndRefresh` we finally see that the `Create` function invokes a `createResource` function [on line 119 of the `Helper Create` function](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go). This is shown below. The `Helper createResource` function, also shown below, performs the actual REST call to the API server to create the resource we defined in our YAML file.

创建了 `Helper` 并在 `createAndRefresh` 的第 217 行调用了它的 `Create` 函数，我们终于看到 `Create` 函数调用了 `createResource` 函数[在 `Helper Create` 函数的第 119 行](https：//github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go)。这如下所示。 `Helper createResource` 函数（如下所示)执行对 API 服务器的实际 REST 调用，以创建我们在 YAML 文件中定义的资源。

![screen capture of the helper `create` and `createResource` functions that actually perform the REST call to the API server to create the resource](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/helper_create.jpg)

## Compiling and running Kubernetes

## 编译和运行 Kubernetes

Now that we have reviewed the code its time to learn how to compile and run the code. In many of the code samples provided above you will see `fmt.Println()` calls in the code. All of those are debugging statements that I added to the code and you can add them to your copy of the source code as well. To compile the code we are going to use a special option that informs the Kubernetes build process to only compile the kubectl portion of the code. This will speed up the Kubernetes compilation process dramatically. The make command for doing the optimize compile is:

现在我们已经查看了代码，是时候学习如何编译和运行代码了。在上面提供的许多代码示例中，您将在代码中看到“fmt.Println()”调用。所有这些都是我添加到代码中的调试语句，您也可以将它们添加到源代码副本中。为了编译代码，我们将使用一个特殊选项，通知 Kubernetes 构建过程只编译代码的 kubectl 部分。这将大大加快 Kubernetes 的编译过程。进行优化编译的 make 命令是：

```
make WHAT='cmd/kubectl'
```

Once we have recompiled the kubectl portion of the code with our debugging print statements added, we can then then start up our Kubernetes development environment using the following command:

一旦我们重新编译了代码的 kubectl 部分并添加了我们的调试打印语句，我们就可以使用以下命令启动我们的 Kubernetes 开发环境：

```
PATH=$PATH KUBERNETES_PROVIDER=local hack/local-up-cluster.sh
```



In another terminal window we can go ahead and run the kubectl commnand and watch it run with our `fmt.Printlns` included. We do this with the following command:

在另一个终端窗口中，我们可以继续运行 kubectl commnand 并观察它在包含我们的 `fmt.Printlns` 的情况下运行。我们使用以下命令执行此操作：

```
cluster/kubectl.sh create -f ~/nginx_kube_example/nginx_pod.yaml

```

The figure below shows what the output looks like with our debugging print statements included:

下图显示了包含我们的调试打印语句后的输出：

![screen capture of running kubectl create with debugging Printlns included](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_create_cmd_with_printfs.jpg)

## Code learning tools 

## 代码学习工具

I know what you are thinking. Brad, you are new to Kube and Go and you quickly learned all this. You must be a genius! Well, sadly I have a list of Twitter followers who would be more than happy to jump in with solid concrete evidence that would easily refute that statement. But through the help of others I have identified several tools and techniques that can really help accelerate your ability to learn the Kubernetes source code. In this section I describe my favorite techniques: use of the Chrome Sourcegraph plugin, properly formatted print statements, the use of a Go `panic` to get desperately needed stack traces, and GitHub `blame` to travel back in time.

我知道你在想什么。 Brad，你是 Kube 和 Go 的新手，你很快就学会了这一切。你一定是个天才！好吧，可悲的是，我有一个 Twitter 追随者名单，他们非常乐意加入可以轻松反驳该声明的可靠具体证据。但是通过其他人的帮助，我发现了一些工具和技术，它们可以真正帮助你提高学习 Kubernetes 源代码的能力。在本节中，我描述了我最喜欢的技术：使用 Chrome Sourcegraph 插件、格式正确的打印语句、使用 Go `panic` 来获取迫切需要的堆栈跟踪，以及 GitHub `blame` 来回到过去。

### Chrome Sourcegraph plugin

### Chrome Sourcegraph 插件

[Morgan Bauer](https://twitter.com/ibmhb) introduced me to one of the coolest tools I have seen for learning Kubernetes code. The [Chrome Sourcegraph plugin](https://chrome.google.com/webstore/detail/sourcegraph-for-github/dgjhfomjieaadpoljlnidmbgkdffpack?hl=en) provides several advanced IDE features that make it dramatically easier to understand Kubernetes Go code when browsing GitHub repositories. Here is an example of how it can help. When I first started looking at Kubernetes code I found the following code snippet absolutely depressing to parse through and understand. It had a ton of functions and it was just overwhelming.

[Morgan Bauer](https://twitter.com/ibmhb) 向我介绍了我见过的用于学习 Kubernetes 代码的最酷的工具之一。 [Chrome Sourcegraph 插件](https://chrome.google.com/webstore/detail/sourcegraph-for-github/dgjhfomjieaadpoljlnidmbgkdffpack?hl=en) 提供了几个高级 IDE 功能，使浏览时更容易理解 Kubernetes Go 代码GitHub 存储库。这是它如何提供帮助的示例。当我第一次开始查看 Kubernetes 代码时，我发现以下代码片段绝对令人沮丧，难以解析和理解。它有很多功能，而且简直是压倒性的。

![Screen capture of code section I found bewildering and depressing when new to Go and Kubernetes programming styles](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph0.jpg)



When looking at this same piece of code in a Chrome browser with the Sourcegraph extension installed you can hover the mouse over each function and quickly get a description of the function, what is passed into the function and what it returns. This is a huge time saver as you can avoid having to grep the code base to understand where a function is defined and what it does. An example of this is shown in the figure below.

在安装了 Sourcegraph 扩展程序的 Chrome 浏览器中查看同一段代码时，您可以将鼠标悬停在每个函数上并快速获取函数的描述、传递给函数的内容以及返回的内容。这可以节省大量时间，因为您可以避免使用 grep 代码库来了解函数的定义位置和功能。下图显示了一个示例。

![Screen capture of Sourcegraph hover view, which makes it obvious that ContinueOnError operates on a Builder object and returns a Builder object and describes what the function does](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph1.jpg)



The Chrome Sourcegraph extension also has an advanced view that provides the ability to peek into the function that is being invoked. This extremely useful capability is shown here.

Chrome Sourcegraph 扩展还具有一个高级视图，可以查看正在调用的函数。这里展示了这个非常有用的功能。

![Screen capture of Chrome Sourcegraph advanced view that provides the ability to peek into the function that is being invoked](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph3.jpg)



One issue I noticed with Chrome Sourcegraph is that sometimes it hangs and fails to pop up the code details. My experience has been that this is easily fixed by simply hitting the refresh button on the browser.

我注意到 Chrome Sourcegraph 的一个问题是它有时会挂起并且无法弹出代码详细信息。我的经验是，只需点击浏览器上的刷新按钮即可轻松解决此问题。

### Print statements never go out of style

### 打印语句永远不会过时

Adding print statements as I have shown throughout this article is a huge help to validating that the code is executing in fashion that matches how you are interpreting it. The `%#v` formatting option shown below typically provides the best debugging information. Don't forget that you may have to add `"fmt"` to your list of imports if it is not already included in the module.

正如我在本文中所展示的那样，添加打印语句对于验证代码是否以与您解释它的方式相匹配的方式执行非常有帮助。下面显示的 `%#v` 格式化选项通常提供最好的调试信息。不要忘记，如果模块中尚未包含“fmt”，则可能必须将其添加到导入列表中。

```
fmt.Println("\n createAndRefresh Info = %#v", info)
```

### When in doubt, PANIC 

### 如有疑问，恐慌

I was having a very difficult time determining how the `createAndRefresh` function in [`Create.go`](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/cmd/create.go#L219) was invoked. Finally, I decided to throw in a `panic` into the code to force a stack trace to be generated and printed to the screen. The code below shows how I added a `panic` to the function. This was a huge help as it helped me to determine which type of `Visitor` was actually being used to invoke the `createAndRefresh` function.

我很难确定 [`Create.go`](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/cmd/create.go#L219 中的`createAndRefresh` 函数) 被调用。最后，我决定在代码中加入“恐慌”，以强制生成堆栈跟踪并将其打印到屏幕上。下面的代码显示了我如何向函数添加“恐慌”。这是一个巨大的帮助，因为它帮助我确定了实际使用哪种类型的“访问者”来调用“createAndRefresh”函数。

```go
func createAndRefresh(info *resource.Info) error {
fmt.Println("\n createAndRefresh Info = %#v", info)
panic("Want Stack Trace")
obj, err := resource.NewHelper(info.Client, info.Mapping).Create(info.Namespace, true, info.Object)
if err != nil {
    return err
}
info.Refresh(obj, true)
return nil
}

```

### Visit the past with GitHub blame

### 用GitHub责备访问过去

Sometimes you look at some lines of source code and you think to yourself, what was the person thinking when they committed those lines of code. Thankfully, the GitHub browser interface has a `blame` option available as a button on the user interface. The figure below shows the location of the **blame** button.

有时您查看某些源代码行，然后您会想，当他们提交这些代码行时，这个人在想什么。值得庆幸的是，GitHub 浏览器界面有一个“责备”选项，可作为用户界面上的按钮使用。下图显示了**责备**按钮的位置。

![Screen capture of blame button on the GitHub browser interface](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/githubblame1.jpg)

When you push the **blame** button, you are given a view of the code that has the commits responsible for each line of code in the source file. This allows you to go back in time and look at the commit that added a particular line of code and determine what the developer was trying to accomplish when that line of code was added. The figure below illustrates the use of the **blame** option and on the left hand side all the commits are listed.

当您按下 **blame** 按钮时，您会看到代码的视图，其中包含负责源文件中每一行代码的提交。这使您可以及时返回并查看添加特定代码行的提交，并确定添加该代码行时开发人员试图完成的工作。下图说明了 **blame** 选项的使用，左侧列出了所有提交。

![Screen capture of Github blame option illustrating which commit was responsible for each line of code](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/githubblame2.jpg)

## Summary

##  概括

In this article, we have examined several key portions of the Kubernetes code base responsible for running a simple kubectl command and the flow through the code that actually sends a REST call to the API Server. We have also provided examples on how to compile and run the command from a Kubernetes development environment. We concluded with a section that describes several useful tools and techniques for learning the Kubernetes source code.

在本文中，我们检查了负责运行简单 kubectl 命令的 Kubernetes 代码库的几个关键部分，以及实际向 API 服务器发送 REST 调用的代码流。我们还提供了有关如何从 Kubernetes 开发环境编译和运行命令的示例。我们最后用一节描述了学习 Kubernetes 源代码的几个有用的工具和技术。

In the next article in this series we will examine another key portion of the Kubernetes code base. In the meantime, hopefully this article will give you the courage to roll up your sleeves and start learning the Kubernetes code base. Every journey starts with a first step!

在本系列的下一篇文章中，我们将研究 Kubernetes 代码库的另一个关键部分。同时，希望这篇文章能让你鼓起勇气，卷起袖子开始学习 Kubernetes 代码库。每一段旅程都从第一步开始！



Table of Contents

目录

- [Overview](http://webcache.googleusercontent.com#overview)
- [Running a basic kubectl command](http://webcache.googleusercontent.com#running-a-basic-kubectl-command)
- [Locating the implementation of kubectl commands in the Kubernetes source code](http://webcache.googleusercontent.com#locating-the-implementation-of-kubectl-commands-in-the-kubernetes-source-code)
- [Kubernetes loves the Cobra command framework](http://webcache.googleusercontent.com#kubernetes-loves-the-cobra-command-framework)
- [Builders and visitors abound in Kubernetes](http://webcache.googleusercontent.com#builders-and-visitors-abound-in-kubernetes)
- [Compiling and running Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes)
- [Code learning tools](http://webcache.googleusercontent.com#code-learning-tools)
- [概述](http://webcache.googleusercontent.com#overview)
- [运行基本的 kubectl 命令](http://webcache.googleusercontent.com#running-a-basic-kubectl-command)
-  [定位Kubernetes源码中kubectl命令的实现](http://webcache.googleusercontent.com#locating-the-implementation-of-kubectl-commands-in-the-kubernetes-source-code)
- [Kubernetes 喜欢 Cobra 命令框架](http://webcache.googleusercontent.com#kubernetes-loves-the-cobra-command-framework)
- [Kubernetes 中的构建者和访问者比比皆是](http://webcache.googleusercontent.com#builders-and-visitors-abound-in-kubernetes)
- [编译和运行Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes)
- [代码学习工具](http://webcache.googleusercontent.com#code-learning-tools)

  - [Chrome Sourcegraph plugin](http://webcache.googleusercontent.com#chrome-sourcegraph-plugin)
   - [Print statements never go out of style](http://webcache.googleusercontent.com#print-statements-never-go-out-of-style)
   - [When in doubt, PANIC](http://webcache.googleusercontent.com#when-in-doubt-panic)
   - [Visit the past with GitHub blame](http://webcache.googleusercontent.com#visit-the-past-with-github-blame)

- [Chrome Sourcegraph 插件](http://webcache.googleusercontent.com#chrome-sourcegraph-plugin)
  - [打印语句永远不会过时](http://webcache.googleusercontent.com#print-statements-never-go-out-of-style)
  - [如有疑问，恐慌](http://webcache.googleusercontent.com#when-in-doubt-panic)
  - [与GitHub责备一起访问过去](http://webcache.googleusercontent.com#visit-the-past-with-github-blame)

- [Summary](http://webcache.googleusercontent.com#summary) 

- [摘要](http://webcache.googleusercontent.com#summary)

