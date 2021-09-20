# How to Write Go Code

## Introduction 

This document demonstrates the development of a simple Go package inside a module and introduces the [go tool](https://golang.org/cmd/go/), the standard way to fetch, build, and install Go modules, packages, and commands.

Note: This document assumes that you are using Go 1.13 or later and the `GO111MODULE` environment variable is not set. If you are looking for the older, pre-modules version of this document, it is archived [here](https://golang.org/doc/gopath_code.html).

## Code organization 

Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A repository contains one or more modules. A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).

Note that you don't need to publish your code to a remote repository before you can build it. A module can be defined locally without belonging to a repository. However, it's a good habit to organize your code as if you will publish it someday.

Each module's path not only serves as an import path prefix for its packages, but also indicates where the `go` command should look to download it. For example, in order to download the module `golang.org/x/tools`, the `go` command would consult the repository indicated by `https://golang.org/x/tools` (described more [here](https://golang.org/cmd/go/#hdr-Relative_import_paths)).

An import path is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. For example, the module `github.com/google/go-cmp` contains a package in the directory `cmp/`. That package's import path is `github.com/google/go-cmp/cmp`. Packages in the standard library do not have a module path prefix.

## Your first program 

To compile and run a simple program, first choose a module path (we'll use `example/user/hello`) and create a `go.mod` file that declares it:

```
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

Next, create a file named `hello.go` inside that directory containing the following Go code:

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

Now you can build and install that program with the `go` tool:

```
$ go install example/user/hello
$
```

This command builds the `hello` command, producing an executable binary. It then installs that binary as `$HOME/go/bin/hello` (or, under Windows, `%USERPROFILE%\go\bin\hello.exe`).

The install directory is controlled by the `GOPATH` and `GOBIN` [environment variables](https://golang.org/cmd/go/#hdr-Environment_variables). If `GOBIN` is set, binaries are installed to that directory. If `GOPATH` is set, binaries are installed to the `bin` subdirectory of the first directory in the `GOPATH` list. Otherwise, binaries are installed to the `bin` subdirectory of the default `GOPATH` (`$HOME/go` or `%USERPROFILE%\go`).

You can use the `go env` command to portably set the default value for an environment variable for future `go` commands:

```
$ go env -w GOBIN=/somewhere/else/bin
$
```

To unset a variable previously set by `go env -w`, use `go env -u`:

```
$ go env -u GOBIN
$
```

Commands like `go install` apply within the context of the module containing the current working directory. If the working directory is not within the `example/user/hello` module, `go install` may fail.

For convenience, `go` commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given. So in our working directory, the following commands are all equivalent:

```
$ go install example/user/hello
$ go install .
$ go install
```

Next, let's run the program to ensure it works. For added convenience, we'll add the install directory to our `PATH` to make running binaries easy:

```
# Windows users should consult https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

If you're using a source control system, now would be a good time to initialize a repository, add the files, and commit your first change. Again, this step is optional: you do not need to use source control to write Go code.

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

### Importing packages from your module 

Let's write a `morestrings` package and use it from the `hello` program. First, create a directory for the package named `$HOME/hello/morestrings`, and then a file named `reverse.go` in that directory with the following contents:

```
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

Because our `ReverseRunes` function begins with an upper-case letter, it is [exported](https://golang.org/ref/spec#Exported_identifiers), and can be used in other packages that import our `morestrings` package.

Let's test that the package compiles with `go build`:

```
$ cd $HOME/hello/morestrings
$ go build
$
```

This won't produce an output file. Instead it saves the compiled package in the local build cache.

After confirming that the `morestrings` package builds, let's use it from the `hello` program. To do so, modify your original `$HOME/hello/hello.go` to use the morestrings package:

```
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

```
$ go install example/user/hello
```

Running the new version of the program, you should see a new, reversed message:

```
$ hello
Hello, Go!
```

### Importing packages from remote modules 

An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The `go` tool uses this property to automatically fetch packages from remote repositories. For instance, to use `github.com/google/go-cmp/cmp` in your program:

```
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

```
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
- 	"Hello World",
+ 	"Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

Module dependencies are automatically downloaded to the `pkg/mod` subdirectory of the directory indicated by the `GOPATH` environment variable. The downloaded contents for a given version of a module are shared among all other modules that `require` that version, so the `go` command marks those files and directories as read-only. To remove all downloaded modules, you can pass the `-modcache` flag to `go clean`:

```
$ go clean -modcache
$
```

## Testing 

Go has a lightweight test framework composed of the `go test` command and the `testing` package.

You write a test by creating a file with a name ending in `_test.go` that contains functions named `TestXXX` with signature `func (t *testing.T)`. The test framework runs each such function; if the function calls a failure function such as `t.Error` or `t.Fail`, the test is considered to have failed.

Add a test to the `morestrings` package by creating the file `$HOME/hello/morestrings/reverse_test.go` containing the following Go code.

```
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

```
$ cd $HOME/hello/morestrings
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

Run `go help test` and see the [testing package documentation](https://golang.org/pkg/testing/) for more detail.

## What's next 

Subscribe to the [golang-announce](https://groups.google.com/group/golang-announce) mailing list to be notified when a new stable version of Go is released.

See [Effective Go](https://golang.org/doc/effective_go.html) for tips on writing clear, idiomatic Go code.

Take  [A Tour of Go](https://tour.golang.org/) to learn the language proper.

Visit the [documentation page](https://golang.org/doc/#articles) for a set of in-depth articles about the Go language and its libraries and tools.

## Getting help 

For real-time help, ask the helpful gophers in the community-run [gophers Slack server](https://gophers.slack.com/messages/general/) (grab an invite [here](https://invite.slack.golangbridge.org/)).

The official mailing list for discussion of the Go language is [Go Nuts](https://groups.google.com/group/golang-nuts).
