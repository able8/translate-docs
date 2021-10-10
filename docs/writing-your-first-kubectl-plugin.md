# Writing your first kubectl plugin with Go

https://bmuschko.com/blog/writing-your-first-kubectl-plugin/


June 20, 2019          

The `kubectl` command line tool offers a powerful, client-side mechanism for communicating with the Kubernetes API server. `kubectl` comes with a lot of fine-grained commands, subcommands and options.  Getting familiar with all of these options takes practice and time.

Despite its wide array of options, you sometimes wish that `kubectl` would provide higher-level functions or would simplify specific  operations. For example, you may have to shell into a running container  to inspect the environment in interactive mode. That’s easy to achieve  with the command `kubectl exec mypod -it --namespace=abc  — /bin/sh`, however, there are various options you have to remember. Wouldn’t it be easier if there would be a simplied command say `kubectl iexec mypod`? That’s where [kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) come in. You can write your own commands and the behavior they should expose.

In this blog post, I want to show how to write a simple `kubectl` plugin that leverages the Go client to communicate with the API server. You will learn how to stand up the Go project, declare the dependency  on the Go client library with the help of Go Modules and write the  source code. You will be able to find the [source code on GitHub](https://github.com/bmuschko/kubectl-server-version) if you’d like to have a look at the details.

> The Go library isn’t the only option for writing a `kubectl` plugin. Kubernetes supports [other languages](https://kubernetes.io/docs/reference/using-api/client-libraries/) though they won’t be subject of this blog post. I personally haven’t  tried any of the other client libraries. I would think that the Go  library is likely the most well-maintained as Kubernetes itself is  written in Go.

## Defining the plugin functionality

The purpose of our plugin is to render the Kubernetes server version  on the command line. We will want to make this functionality available  with the command `server-version`. The result should look similar to the console output below.

```bash
$ kubectl server-version
Hello from Kubernetes server with version v1.10.11!
```

Furthermore, it’s a good idea to provide a subcommand that renders  the version of the plugin so that the end user knows which feature set  to expect. It will help tremendously with the upgrade process if you  know which plugin version has been installed on your machine already.  The following console output indicates that the plugin with version `0.2.0` is used.

```bash
$ kubectl server-version version
kubectl server-version v0.2.0
```

We defined the commands we want to expose and their functionality. Let’s start by setting up the project structure.

## Setting up the project

There’s really nothing special about setting up a Go project that implements a `kubectl` plugin. You’d follow the same conventions as you usually would for any other Go project.

We’ll start by adding a `main.go` file, the entrypoint into the application. Next, we’ll generate the Go Module with `go mod init github.com/bmuschko/kubectl-server-version`. You’ll end up with a `go.mod` and `go.sum` file that we will populate later. The package `cmd` is going to contain all relevant CLI command handling. Moreover, I  personally like to provide a license for the project and a README file  that explains how to use the plugin. In the end, you should end up with a project layout shown below.

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

## Declaring the dependencies

The initial structure of the `go.mod` file has already been set up when we ran the `go init` command. Remember that you have to set the environment variable `GO111MODULES=on` to enable Go Modules before adding any dependencies. To get started with the Go client simply run the command `go get k8s.io/client-go`. There are other portions to the Go client library that you might need e.g. `k8s.io/api`, `k8s.io/apimachinery`. You can retrieve them with the same Go command. Have a look at the [documentation](https://github.com/kubernetes/client-go/) for more information.

The Go Modules file below shows my final list of dependencies. As you can see, we are building the project with Go 1.12. I also added the  dependency `github.com/spf13/cobra` for exposing the CLI commands and `github.com/stretchr/testify` to help with simplifying test code.

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

## Implementing the plugin functionality

Not surprisingly, we will start by filling the `main.go` with life. The implementation logic is straightforward. First, we’ll create a new command by calling the function `NewServerVersionCommand`. Standard input, standard output and standard error have been wrapped with `genericclioptions.IOStreams` and passed in as parameter. Next, we’ll execute the newly created command and handle a potential error case.

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
	if err := serverVersionCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

The value of the variable `version` can be injected at build-time and represents the plugin version. We’ll pass on the value to the package `cmd` for later use in the `version` subcommand.

Let’s have a look at the meat of the plugin functionality, the  command that communicates with the API server and renders the returned  Kubernetes server version. I split up the code into multiple,  digestiable pieces. The code block below uses the [Cobra library](https://github.com/spf13/cobra) to create the new command. Apart from the Kubernetes package imports,  there’s nothing special about this implementation. You can easily  imagine similar code as part of any other CLI tool written in Go.

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

The "client" instance is called a `clientset` and is represented by the struct `kubernetes.Clientset`. For the creation of the the struct, you will have to provide additional client configuration in the form of a `restclient.Config`. Once created, the `clientset` gives access to the whole Kubernetes API and all relevant operations  for querying, creating, updating and deleting Kubernetes objects.

The function below retrieves the `DiscoveryClient` which can discover server-supported API groups, versions and resources. For example, the function `ServerVersion()` returns the Kubernetes server version.

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

## Installing and using the plugin

There are a couple of requirements for installing `kubectl` plugins on a machine.

1. A plugin can be installed as bash script or as binary.
2. The plugin name has to have the prefix `kubectl-`.
3. The portion of the plugin name after `kubectl-` represents the command name.
4. A dash character in the command name has to be changed to an underscore.

In this post, we didn’t even talk about the option of writing a plugin as bash script. For more information, see the [Kubernetes documentation](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#example-plugin). Our plugin has been written in Go so we’ll have to build the binary first. The easiest option is to use the `go build` command. The "installation process" isn’t complicated at all. You simply copy the binary into the `PATH`, as described [here](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#installing-kubectl-plugins).

The following commands build the binary of the plugin and copy it into the `PATH`. Under MacOSX the `PATH` is represented by `/usr/local/bin`.

```bash
$ go build -o kubectl-server_version
$ cp kubectl-server_version /usr/local/bin
$ kubectl plugin list
The following kubectl-compatible plugins are available:

/usr/local/bin/kubectl-server_version
```

Awesome, the plugin has been installed. We can now go ahead and  execute it as described earlier. As you can see, the development and  installation process is not very burdensome. But what if you wanted to  share the plugin with the team or the Kubernetes community? Let’s talk  about distribution the plugin in the next section.

## Distributing the plugin on GitHub Releases

There are various options for distributing the binary of your plugin. In this post, we’ll discuss the distribution of the plugin on [GitHub Releases](https://help.github.com/en/articles/creating-releases).

> Another popular option is to use the package manager for `kubectl` plugins called [krew](https://github.com/kubernetes-sigs/krew). It’s your best option when it comes to easy of use and making the  plugin available to the world. I will leave this topic for another blog  post.

My tool of choice for publishing binaries to GitHub Releases is [GoReleaser](https://github.com/goreleaser/goreleaser). It helps with defining, building and publishing the binaries you want  to release with a given Git tag. GoReleaser reads the configuration for  the binaries in a file named `.goreleaser.yml`. You can find [an example](https://github.com/bmuschko/kubectl-server-version/blob/master/.goreleaser.yml) for such a file in the plugin repository. Once a version has been  released, any consumer can download the binary for a target platform and install the plugin directly into the `PATH`.

## Conclusion

The blog post taught you how to get started with writing a `kubectl` plugin using the Go language. We created a new project from scratch and took it all the way from inception to distribution. I hope I could  inspire you to write your own plugins. In the course of the next couple  of months, I intend to write more blog posts on this topic so stay  tuned! In the meantime, check out the all the [kubectl plugins that have been published by the community](https://github.com/topics/kubectl-plugins). Maybe you will find the functionality you’ve been missing for so long from the standard `kubectl` command.