# A tour of the Kubernetes source code

From kubectl to API Server

SaveLike By [Brad Topol](https://developer.ibm.com/profiles/btopol), Published June 21, 2017

https://developer.ibm.com/articles/a-tour-of-the-kubernetes-source-code/

* * *

## Overview

Kubernetes continues to experience explosive growth and software developers that are able to understand and contribute to the Kubernetes code base are in high demand. Learning the Kubernetes code base is not easy. Kubernetes is written Go which is a fairly new programming language and it has a large amount of source code.

In this multi-part series of articles I will dig in and explain key portions of the Kubernetes code base and also explain the techniques I have used to help me understand the code. My goal is to provide a set of articles that will enable software developers new to Kubernetes to more quickly learn the Kubernetes source code.

In this first article, I will cover the flow through the code from running a simple kubectl command to sending a REST call to the API Server. Before using this article to dig into the Kubernetes code, I recommend you read an outstanding high level [overview](https://jvns.ca/blog/2017/06/04/learning-about-kubernetes/) of the Kubernetes architecture by [Julia Evans](https://twitter.com/b0rk).

## Running a basic kubectl command

The command line interface for Kubernetes is called kubectl. It is used for running commands against Kubernetes clusters. When attempting to learn the Kubernetes source code, the portion of the source code that implements the command line interface is a great place to start. The command we will use to trace through the source code is the `kubectl create -f` command which creates a resource from a file. The resource we are creating is a single replica pod with a basic `nginx` container image. The specification for this resource is shown below and I placed this specification in a file called `~/nginx_kube_example/nginx_pod.yaml`.

```
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

Show moreShow more icon

From a Kubernetes development environment we can invoke kubectl as shown in the figure below:

![screen capture of kubectl command to create a resource from a file](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_cmd_trimmed.jpg)

Now that we know what kubectl command we are running and how to run it, let's look at where we can find the implementation of this command in the Kubernetes source code.

## Locating the implementation of kubectl commands in the Kubernetes source code

The entry point for all the kubectl commands can be found in the [`github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd`](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/cmd) folder. In this folder there is a name of a Go file that matches the name of the kubectl command that is implemented. For example, the `kubectl create` command has an initial entry point in a file named `create.go`. The folder and the example go files implementing the various commands are shown in the figure below.

![screen capture of directory containing code entry points for all the kubectl commands](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_cmd_dir.jpg)

## Kubernetes loves the Cobra command framework

Kubernetes commands are implemented using the [Cobra command framework](https://github.com/spf13/cobra). Cobra provides a lot of great features for building command line interfaces and a basic overview of Cobra's capabilities can be found [here](https://blog.gopheracademy.com/advent-2014/introducing-cobra/). As shown in the Figure, one of the nice features of how Kubernetes utilizes Cobra is that it is very easy to locate which file implements each command line option. Furthermore, the Cobra structure puts the command usage message and command descriptions adjacent to the code that runs the command. This is shown in the figure and the [actual lines of code](https://github.com/kubernetes/kubernetes/blob/fd9a91e0b57face905c4225b8a6633b2ea9c832d/pkg/kubectl/cmd/create.go#L62-%2376). What's great about this structure is that you can go through and look at the descriptions for all the Kubernetes kubectl commands and then quickly jump to the code that implements the commands. As shown in lines 62-76 in the figure, the strings `Use`, `Short`, `Long`, and `Example` all hold information describing the command and `Run` points to a function that actually runs the command.

![screen capture of RunCreate function, which performs the bulk of the kubectl create command](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_NewCMDCreate.jpg)

The `RunCreate` function invoked on line 74 in the above figure is where the bulk of the kubectl create command is implemented. The implementation of this function can be found in the same `create.go` file. The figure below shows the `RunCreate` function. On line 132 I added a `fmt.Println` just to confirm this code was being called when I thought it would be called. In the [Compiling and running Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes) section below I show how you can speed up recompiling the Kubernetes code base when adding debugging statements solely to the kubectl source code.

![screen capture of RunCreate function in create.go](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_RunCreate.jpg)

## Builders and visitors abound in Kubernetes

The `resource.NewBuilder` function code shown on lines 133-140 is particularly intimidating to someone new to Go and Kubernetes. It's worth taking some time to explain this section of code in more detail. At a high level, what this code does is take the arguments and parameters from the command line and converts them into a list of resources. It's also responsible for creating a visitor construct that can be used to iterate across all the resources. The code is complex because it uses a variant of the [builder pattern](https://en.wikipedia.org/wiki/Builder_pattern) where individual functions are each doing a separate portion of the data initialization. The functions `Schema`, `ContinueOnError`, `NamespaceParam`, `DefaultNamespace`, `FilenameParam`, `SelectorParam`, and `Flatten` all take in a pointer to a `Builder struct`, perform some form of modification on the `Builder struct`, and then return the pointer to the `Builder struct` for the next method in the chain to use when it performs its modifications. All of these methods can be found in the `builder.go` file but I have included a few below so you can see how they work.

```
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

Show moreShow more icon

Once all the initializers have completed, the `resource.NewBuilder` function finally invokes a `Do` function. The `Do` function is a critical piece as it returns a `Result` object that will be used to drive the creation of our resource. The `Do` function also creates a `Visitor` object that can be used to traverse the list of resources that were associated with this invocation of `resource.NewBuilder`. The `Do` function implementation is shown below.

![screen capture of the Builder Do function that creates a DecoratedVisitor and returns a Result object](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/builder_Do.jpg)

As shown above on line 816, a new `DecoratedVisitor` is created and stored as part of the `Result` object that is returned by the `Builder Do` function. The `DecoratedVisitor` has a `Visit` function that will call the `Visitor` function that is passed into it. The implementation of this can be found at [github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/visitor.go#L306](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/visitor.go#L306) and is shown below.

![screen capture of DecoratedVisitor Visit function that will eventually invoke createAndRefresh](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/visitor_NewDecoratedVisitor.jpg)

The `Result` object returned by the `Do` function has a `Visit` function that is used to invoke the `DecoratedVisitor Visit` function. This provides us a path from line 150 in the `RunCreate` function in `create.go` to eventually calling the anonymous function that is passed in on line 150 and contains the `createAndRefresh` function that will lead us to the code making a REST call to the API server. The implementation of the `Result Visit` function that is called on line 150 of `RunCreate` function in `create.go` is shown below.

![screen capture of Result Visit function that takes as parameter the function to invoke when visiting resources](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/result_Visit.jpg)

Now that we have seen how everything is connected through `Visit` functions and `DecoratedVisitor` classes, we see that the inline visitor function on line 150 below has a `createAndRefresh` function on line 165.

![screen capture of enabling createAndRefresh to be invoked from a Result's Visitor object](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_RunCreate_callCreateAndRefresh.jpg)

The `createAndRefresh` function invokes the `Resource` [`NewHelper`](https://github.com/kubernetes/kubernetes/blob/a3501fb9948f6a4834919b6778227cbd75f8c400/pkg/kubectl/resource/helper.go#L46) function found in `helper.go` and this function returns a new `Helper` object.

![screen capture of the createAndRefresh function creating a new Helper object that performs the REST call to the Kubernetes APIServer](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/create_createAndRefresh.jpg)

Here is the [code](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go) that returns a new `Helper` object. It is actually pretty straightforward.

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

Show moreShow more icon

With the `Helper` created and its `Create` function invoked on line 217 of `createAndRefresh` we finally see that the `Create` function invokes a `createResource` function [on line 119 of the `Helper Create` function](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/resource/helper.go). This is shown below. The `Helper createResource` function, also shown below, performs the actual REST call to the API server to create the resource we defined in our YAML file.

![screen capture of the helper `create` and `createResource` functions that actually perform the REST call to the API server to create the resource](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/helper_create.jpg)

## Compiling and running Kubernetes

Now that we have reviewed the code its time to learn how to compile and run the code. In many of the code samples provided above you will see `fmt.Println()` calls in the code. All of those are debugging statements that I added to the code and you can add them to your copy of the source code as well. To compile the code we are going to use a special option that informs the Kubernetes build process to only compile the kubectl portion of the code. This will speed up the Kubernetes compilation process dramatically. The make command for doing the optimize compile is:

```
make WHAT='cmd/kubectl'

```

Show moreShow more icon

Once we have recompiled the kubectl portion of the code with our debugging print statements added, we can then then start up our Kubernetes development environment using the following command:

```
PATH=$PATH KUBERNETES_PROVIDER=local hack/local-up-cluster.sh

```

Show moreShow more icon

In another terminal window we can go ahead and run the kubectl commnand and watch it run with our `fmt.Printlns` included. We do this with the following command:

```
cluster/kubectl.sh create -f ~/nginx_kube_example/nginx_pod.yaml

```

Show moreShow more icon

The figure below shows what the output looks like with our debugging print statements included:

![screen capture of running kubectl create with debugging Printlns included](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/kubectl_create_cmd_with_printfs.jpg)

## Code learning tools

I know what you are thinking. Brad, you are new to Kube and Go and you quickly learned all this. You must be a genius! Well, sadly I have a list of Twitter followers who would be more than happy to jump in with solid concrete evidence that would easily refute that statement. But through the help of others I have identified several tools and techniques that can really help accelerate your ability to learn the Kubernetes source code. In this section I describe my favorite techniques: use of the Chrome Sourcegraph plugin, properly formatted print statements, the use of a Go `panic` to get desperately needed stack traces, and GitHub `blame` to travel back in time.

### Chrome Sourcegraph plugin

[Morgan Bauer](https://twitter.com/ibmhb) introduced me to one of the coolest tools I have seen for learning Kubernetes code. The [Chrome Sourcegraph plugin](https://chrome.google.com/webstore/detail/sourcegraph-for-github/dgjhfomjieaadpoljlnidmbgkdffpack?hl=en) provides several advanced IDE features that make it dramatically easier to understand Kubernetes Go code when browsing GitHub repositories. Here is an example of how it can help. When I first started looking at Kubernetes code I found the following code snippet absolutely depressing to parse through and understand. It had a ton of functions and it was just overwhelming.

![Screen capture of code section I found bewildering and depressing when new to Go and Kubernetes programming styles](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph0.jpg)

When looking at this same piece of code in a Chrome browser with the Sourcegraph extension installed you can hover the mouse over each function and quickly get a description of the function, what is passed into the function and what it returns. This is a huge time saver as you can avoid having to grep the code base to understand where a function is defined and what it does. An example of this is shown in the figure below.

![Screen capture of Sourcegraph hover view, which makes it obvious that ContinueOnError operates on a Builder object and returns a Builder object and describes what the function does](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph1.jpg)

The Chrome Sourcegraph extension also has an advanced view that provides the ability to peek into the function that is being invoked. This extremely useful capability is shown here.

![Screen capture of Chrome Sourcegraph advanced view that provides the ability to peek into the function that is being invoked](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/sourcegraph3.jpg)

One issue I noticed with Chrome Sourcegraph is that sometimes it hangs and fails to pop up the code details. My experience has been that this is easily fixed by simply hitting the refresh button on the browser.

### Print statements never go out of style

Adding print statements as I have shown throughout this article is a huge help to validating that the code is executing in fashion that matches how you are interpreting it. The `%#v` formatting option shown below typically provides the best debugging information. Don't forget that you may have to add `"fmt"` to your list of imports if it is not already included in the module.

```
fmt.Println("\n createAndRefresh Info = %#v", info)

```

Show moreShow more icon

### When in doubt, PANIC

I was having a very difficult time determining how the `createAndRefresh` function in [`Create.go`](https://github.com/kubernetes/kubernetes/blob/6b52d8f1383d3a4a769b403a04f812c99ed98815/pkg/kubectl/cmd/create.go#L219) was invoked. Finally, I decided to throw in a `panic` into the code to force a stack trace to be generated and printed to the screen. The code below shows how I added a `panic` to the function. This was a huge help as it helped me to determine which type of `Visitor` was actually being used to invoke the `createAndRefresh` function.

```
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

Show moreShow more icon

### Visit the past with GitHub blame

Sometimes you look at some lines of source code and you think to yourself, what was the person thinking when they committed those lines of code. Thankfully, the GitHub browser interface has a `blame` option available as a button on the user interface. The figure below shows the location of the **blame** button.

![Screen capture of blame button on the GitHub browser interface](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/githubblame1.jpg)

When you push the **blame** button, you are given a view of the code that has the commits responsible for each line of code in the source file. This allows you to go back in time and look at the commit that added a particular line of code and determine what the developer was trying to accomplish when that line of code was added. The figure below illustrates the use of the **blame** option and on the left hand side all the commits are listed.

![Screen capture of Github blame option illustrating which commit was responsible for each line of code](https://developer.ibm.com/developer/default/articles/a-tour-of-the-kubernetes-source-code/images/githubblame2.jpg)

## Summary

In this article, we have examined several key portions of the Kubernetes code base responsible for running a simple kubectl command and the flow through the code that actually sends a REST call to the API Server. We have also provided examples on how to compile and run the command from a Kubernetes development environment. We concluded with a section that describes several useful tools and techniques for learning the Kubernetes source code.

In the next article in this series we will examine another key portion of the Kubernetes code base. In the meantime, hopefully this article will give you the courage to roll up your sleeves and start learning the Kubernetes code base. Every journey starts with a first step!



Table of Contents

- [Overview](http://webcache.googleusercontent.com#overview)
- [Running a basic kubectl command](http://webcache.googleusercontent.com#running-a-basic-kubectl-command)
- [Locating the implementation of kubectl commands in the Kubernetes source code](http://webcache.googleusercontent.com#locating-the-implementation-of-kubectl-commands-in-the-kubernetes-source-code)
- [Kubernetes loves the Cobra command framework](http://webcache.googleusercontent.com#kubernetes-loves-the-cobra-command-framework)
- [Builders and visitors abound in Kubernetes](http://webcache.googleusercontent.com#builders-and-visitors-abound-in-kubernetes)
- [Compiling and running Kubernetes](http://webcache.googleusercontent.com#compiling-and-running-kubernetes)
- [Code learning tools](http://webcache.googleusercontent.com#code-learning-tools)

  - [Chrome Sourcegraph plugin](http://webcache.googleusercontent.com#chrome-sourcegraph-plugin)
  - [Print statements never go out of style](http://webcache.googleusercontent.com#print-statements-never-go-out-of-style)
  - [When in doubt, PANIC](http://webcache.googleusercontent.com#when-in-doubt-panic)
  - [Visit the past with GitHub blame](http://webcache.googleusercontent.com#visit-the-past-with-github-blame)

- [Summary](http://webcache.googleusercontent.com#summary)

