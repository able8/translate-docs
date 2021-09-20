# Using Go instead of bash for scripts

# 在脚本中使用 Go 而不是 bash

From: https://blog.kowalczyk.info/article/4b1f9201181340099b698246857ea98d/using-go-instead-of-bash-for-scripts.html

I like to automate my programming work.

我喜欢自动化我的编程工作。

In every programming project I ended up writing bash (on Unix and Mac) and batch / PowerShell (on Windows) scripts.

在每个编程项目中，我最终都编写了 bash（在 Unix 和 Mac 上）和批处理/PowerShell（在 Windows 上）脚本。

I settled on a convention to put scripts in a directory `s` (it's short, fast to type and a shortcut for `scripts`).

我确定了一个约定，将脚本放在目录 `s` 中（它很短，输入速度很快，并且是 `scripts` 的快捷方式）。

I had `./s/run.sh` to run the program locally.`./s/tests.sh` to run tests. `./s/deploy.sh` to deploy web apps to the server etc.

我有`./s/run.sh` 在本地运行程序。`./s/tests.sh` 来运行测试。 `./s/deploy.sh` 将 Web 应用程序部署到服务器等。

It worked but I wasn't quite happy.

它奏效了，但我不太高兴。

For cross-platform projects I had to write the same script twice (`./s/run.sh` and `./s/run.bat`).

对于跨平台项目，我必须编写两次相同的脚本（`./s/run.sh` 和 `./s/run.bat`）。

I write those scripts so infrequently that I every time I need to re-learn basics. How do I declare a function? How do I write `if`? How do I write a loop?

我很少编写这些脚本，以至于每次我都需要重新学习基础知识。如何声明一个函数？我怎么写“如果”？我如何写一个循环？

In bash, simple things are complicated and non-simple things are very complicated.

在bash中，简单的事情很复杂，不简单的事情很复杂。

This article describes how I replaced bash scripts with a single, multiple-purpose Go program.

本文描述了我如何用一个单一的、多用途的 Go 程序替换 bash 脚本。

You can see a full example at https://github.com/kjk/notionapi/tree/master/do

你可以在 https://github.com/kjk/notionapi/tree/master/do 看到一个完整的例子

# Replacing bash with Go

# 用 Go 替换 bash

One day it hit me: I would rather write the helper scripts in Go.

有一天，它打动了我：我宁愿用 Go 编写辅助脚本。

Go is cross-platform; I don't have to write the same thing twice.

Go 是跨平台的；我不必两次写同样的东西。

I write in Go daily so I can implement simple things quickly.

我每天都用 Go 编写代码，这样我就可以快速实现简单的事情。

In Go simple things are simple and complicated things are possible.

在 Go 中，简单的事情是简单的，复杂的事情是可能的。

The one drawback is more lines of code but the difference is immaterial. Those are short programs either way.

一个缺点是代码行数更多，但差异并不重要。无论哪种方式，这些都是简短的节目。

This article describes a system I refined by using it in multiple projects.

本文描述了我在多个项目中使用它改进的系统。

## Establishing conventions

## 建立约定

It's all about convenience so less typing is better.

这一切都是为了方便，所以打字越少越好。

Establishing conventions to share between multiple projects frees mental energy for more important things.

建立约定以在多个项目之间共享可以为更重要的事情释放精神能量。

The system I settled on is:

我选择的系统是：

- `do` directory contains a single, multiple-purpose Go program. A single program that does many things is better suited to Go  than multiple programs as it makes it easy to share helper functions

- `do` 目录包含一个单一的、多用途的 Go 程序。一个可以做很多事情的程序比多个程序更适合 Go，因为它可以很容易地共享辅助函数

- to run it: `cd do; go run . ${flags}`

- 运行它：`cd do;去跑步 。 ${flags}`

- to make things easier to type, I have

   - 为了使输入更容易，我有

  ```
   do/do.sh
  ```

  
:

   ：

  ```
   #!/bin/bash
  
  cd do
  go run -race .$@
  ```

- for Windows I have

   - 对于 Windows，我有

  ```
   do\do.bat
  ```

  
:

   ：

  ```
   @cd do
  @go run -race .%*
  ```

- to make things even easier to type, I add the following to

   - 为了使输入更容易，我将以下内容添加到

  ```
   ~/.bash_profile
  ```

   
(on Mac):

   （在 Mac 上）：

  ```
   function doit() {
      if [ -f ./do.sh ];then
          ./do.sh $@
      elif [ -f ./do/do.sh ];then
          ./do/do.sh $@
      else
          echo "no do.sh or do/do.sh found"
      fi
  }
  ```

  
I can then type `doit ${args}` to launch either `./do.sh` or `./do/do.sh` (whichever exists).

然后我可以输入 `doit ${args}` 来启动 `./do.sh` 或 `./do/do.sh`（以存在者为准）。

In every project I can type  `doit -run` which executes `./do/do.sh -run` which executes `cd do; go run . -run`.

在每个项目中，我都可以输入执行 `./do/do.sh -run` 的 `doit -run`，执行 `cd do;去跑步 。 -运行`。

In the old system, that would be `./s/run.sh` or `.\s\run.bat`.

在旧系统中，这将是`./s/run.sh` 或`.\s\run.bat`。

Other cmd-line arguments trigger other actions e.g. `doit -test`, `doit -deploy` etc.

其他 cmd 行参数触发其他操作，例如`doit -test`、`doit -deploy` 等。

If I forget which flags are available, `do` without arguments prints them all.

如果我忘记了哪些标志可用，则不带参数的 `do` 会将它们全部打印出来。

## A structure of `do` program

##`do`程序的结构

Main function checks cmd-line arguments and calls the right function to perform a given command.

Main 函数检查 cmd 行参数并调用正确的函数来执行给定的命令。

Here's an implementation of dispatching two commands: `-run` and `-deploy`.

这是分派两个命令的实现：`-run` 和 `-deploy`。

```
func main() {
    cdToTopDir()
    fmt.Printf("topDir: '%s'\n", topDir())

    var (
        flgRun    bool
        flgDeploy bool
    }
    
    {
        flag.BoolVar(&flgRun, "run", false, "runs the program")
        flag.BoolVar(&flgDeploy, "deploy", false, "deploys to production")
        flag.Parse()
    }

    if flgRun {
        doRun()
        return
    }

    if flgDeploy {
        doDeploy()
        return
    }

  // this prints available flags
    flag.Usage()
}
```

## Running from a known current directory

## 从已知的当前目录运行

When we run the program, we're inside `do` directory

当我们运行程序时，我们在 `do` 目录中

It's important to know what is the current director so that when we refer to files in the project, we know their path.

知道当前目录是什么很重要，这样当我们引用项目中的文件时，我们就知道它们的路径。

By convention I set current directory to be top directory of the project.

按照惯例，我将当前目录设置为项目的顶级目录。

The first thing that the program does is call `cdToTopDir()` which fixes the current directory to this known location.

程序做的第一件事是调用`cdToTopDir()`，它将当前目录固定到这个已知位置。

The simplest implementation:

最简单的实现：

```
func cdToTopDir() {
    err := os.Chdir("..")
    must(err)
}
```

This relies on knowledge that we execute the program with `cd do; go run . ${args}`.

这依赖于我们使用 `cd do; 执行程序的知识；去跑步 。 ${args}`。

I also print the absolute path of current directory at the beginning to make sure it is correct.

我还在开头打印当前目录的绝对路径以确保它是正确的。

## Crashing on errors is fine

## 出错时崩溃很好

In a regular Go program, handling errors by propagating them to callers is key for writing robust software. 

在常规 Go 程序中，通过将错误传播给调用者来处理错误是编写健壮软件的关键。

In short scripts it's ok to `panic` when error happens. It makes for shorter code. `panic` prints the callstack which is handy when investigating unexpected errors.

简而言之，发生错误时可以“恐慌”。它使代码更短。 `panic` 打印调用堆栈，这在调查意外错误时很方便。

I have a helper function `must(err error)` that panics if err is not nil:

我有一个辅助函数 `must(err error)`，如果 err 不是 nil，它就会恐慌：

```
func must(err error) {
    if err != nil {
        fmt.Printf("err: %s\n", err)
        panic(err)
    }
}
```

Here's how to use it:

以下是如何使用它：

```
func readFile(path string) []byte {
    d, err := ioutil.ReadFile(path)
    must(err)
    return d
}
```

## Executing programs

## 执行程序

A common thing to do is executing other programs. For example,  `do -run` would typically execute `go build . -o myapp` and `./myapp.`

通常要做的事情是执行其他程序。例如，`do -run` 通常会执行 `go build 。 -o myapp` 和 `./myapp.`

Go has an excellent `os/exec` package for that:

Go 有一个很好的 `os/exec` 包：

```
cmd := exec.Command("go", "build", ".", "-o", "myapp")
err := cmd.Run()
must(err)

cmd = exec.Command("./myapp")
err = cmd.Run()
must(err)
```

Other useful things possible with `exec.Cmd`:

`exec.Cmd` 可能还有其他有用的东西：

- set working directory of the executed program

   - 设置执行程序的工作目录

  ```
   cmd := exec.Command("./myapp")
  cmd.Dir = "working/directory"
  ```

- setting environment variables

   - 设置环境变量

  ```
   cmd.Env = os.Environ()
  cmd.Env = append(cmd.Env, "GOOS=linux", "GOARCH=amd64")
  ```

- get the output of the command

   - 获取命令的输出

  ```
   cmd := exec.Command("ls", "-lah")
  // CombinedOutput() calls Run() and returns captured stdout / stderr as []byte
  out, err := cmd.CombinedOutput()
  must(err)
  fmt.Printf("output of ls:\n%s\n", string(out))
  ```

- sometimes you don't want to block waiting for the program to finish. Very true for launching Windows GUI programs:

   - 有时您不想阻塞等待程序完成。非常适合启动 Windows GUI 程序：

  ```
   func openNotepadWithFile(path string) {
      cmd := exec.Command("notepad.exe", path)
      err := cmd.Start() // this starts the programs but doesn't wait for it to finish
      must(err)
  }
  ```

- if you

   - 如果你

  ```
   Start()
  ```

   
a process, you might want to ensure it's killed on exit:

   一个进程，您可能希望确保它在退出时被杀死：

  ```
   func main() {
      // ...
      err := cmd.Start()
      must(err)
  
      // ensure to kill the process upon exit
      defer cmd.Process.Kill()
  }
  ```

- logging the commands we execute

   - 记录我们执行的命令

  ```
   fmt.Printf("Running: %s\n", strings.Join(cmd.Args[1:], " "))
  ```

- seeing program's stdout and stderr while it's executing

   - 在程序执行时查看程序的 stdout 和 stderr

  ```
   // set program's stdout / stderr ot console's stdout/stderr to
  // see what it prints
  // incompatible with capturing stdout / stderr with `CombinedOutput()`
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  ```

# Helper functions

# 辅助函数

Scripts in different projects often need the same functionality.

不同项目中的脚本通常需要相同的功能。

I keep common functions in a separate file `util.go` so that I can quickly bootstrap new project.

我将常用功能保存在单独的文件 `util.go` 中，以便我可以快速启动新项目。

Here are a few common helper functions.

下面是一些常见的辅助函数。

### assert

### 断言

Inspired by C, panics if condition is not true. Use to verify you get expected results.

受 C 的启发，如果条件不为真则恐慌。用于验证您获得了预期的结果。

```
func assert(cond bool, format string, args ...interface{}) {
    if cond {
        return
    }
    s := fmt.Sprintf(format, args...)
    panic(s)
}
```

### logf

### 日志

To be used instead of `fmt.Printf`. The advantage is that if we want to e.g. start logging to file, we need to change just `logf` function.

用于代替 `fmt.Printf`。优点是，如果我们想要例如开始记录到文件，我们只需要更改 `logf` 函数。

```
// a centralized place allows us to tweak logging, if need be
func logf(format string, args ...interface{}) {
    if len(args) == 0 {
        fmt.Print(format)
        return
    }
    fmt.Printf(format, args...)
}
```

### openBrowser

### 打开浏览器

When working on backends for web apps it's convenient to auto-open the web site in the browser when starting the app locally.

在处理 Web 应用程序的后端时，在本地启动应用程序时，可以方便地在浏览器中自动打开网站。

```
// openBrowsers open web browser with a given url
// (can be http:// or file://)
func openBrowser(url string) {
    var err error
    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }
    must(err)
}
```

### readZipFile

### 读取压缩文件

Reads all files in a zip file and returns them as a map from file name to content.

读取 zip 文件中的所有文件，并将它们作为从文件名到内容的映射返回。

```
func readZipFile(path string) map[string][]byte {
    r, err := zip.OpenReader(path)
    must(err)
    defer r.Close()
    res := map[string][]byte{}
    for _, f := range r.File {
        rc, err := f.Open()
        must(err)
        d, err := ioutil.ReadAll(rc)
        must(err)
        rc.Close()
        res[f.Name] = d
    }
    return res
}
```

### readFile, writeFile

### 读取文件，写入文件

Shorter way to read / write files.

更短的读/写文件的方法。

```
func readFile(path string) []byte {
    d, err := ioutil.ReadFile(path)
    must(err)
    return d
}

func writeFile(path string, data []byte) {
    err := ioutil.WriteFile(path, data, 0666)
    must(err)
}
```

### getHomeDir

### getHomeDir

Returns a path of the user's home directory.

返回用户主目录的路径。

```
func getHomeDir() string {
    s, err := os.UserHomeDir()
    must(err)
    return s
}
```

### cpFile

### cp文件

Equivalent of `cp` in bash.

相当于 bash 中的 `cp`。

```
func cpFile(dstPath, srcPath string) {
    d, err := ioutil.ReadFile(srcPath)
    must(err)
    err = ioutil.WriteFile(dstPath, d, 0666)
    must(err)
}
```

### checkGitClean

### checkGitClean

To prevent accidental deploys, my scripts use `checkGitClean` and refuse to deploy if there are un-commited changes to working directory:

为了防止意外部署，我的脚本使用 `checkGitClean` 并在工作目录有未提交的更改时拒绝部署：

```
var (
    verbose bool
)

func runCmd(cmd *exec.Cmd) string {
    if verbose {
        fmt.Printf("> %s\n", strings.Join(cmd.Args, " "))
    }
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("%s failed with '%s'. Output:\n%s\n", strings.Join(cmd.Args, " "), err, string(out))
    }
    must(err)
    if verbose && len(out) > 0 {
        fmt.Printf("%s\n", out)
    }
    return string(out)
}

func gitStatus(dir string) string {
    cmd := exec.Command("git", "status")
    if dir != "" {
        cmd.Dir = dir
    }
    return runCmd(cmd)
}

func checkGitClean(dir string) {
    s := gitStatus(dir)
    expected := []string{
        "On branch master",
        "Your branch is up to date with 'origin/master'.",
        "nothing to commit, working tree clean",
    }
    for _, exp := range expected {
        if !strings.Contains(s, exp) {
            fmt.Printf("Git repo in '%s' not clean.\nDidn't find '%s' in output of git status:\n%s\n", dir, exp, s)
            os.Exit(1)
        }
    }
}
```

### createZipFile

### createZipFile

This is a helper to create a .zip archive with the content of one or more directories or files.

这是使用一个或多个目录或文件的内容创建 .zip 存档的助手。

Example use: `createZipFile("archive.zip", ".", "myapp", "www")`

使用示例：`createZipFile("archive.zip", ".", "myapp", "www")`

This creates [`archive.zip`](http://archive.zip) with the content of `myapp` file and `www` directory. Those files are located in current (`.`) directory.

这将创建 [`archive.zip`](http://archive.zip) 包含 `myapp` 文件和 `www` 目录的内容。这些文件位于当前 (`.`) 目录中。

In fairness, would be shorter to sub-launch `zip` program, but I like the control.

公平地说，子启动`zip`程序会更短，但我喜欢控件。

```
func zipAddFile(zw *zip.Writer, zipName string, path string) {
    zipName = filepath.ToSlash(zipName)
    d, err := ioutil.ReadFile(path)
    must(err)
    w, err := zw.Create(zipName)
    _, err = w.Write(d)
    must(err)
    if verbose {
        fmt.Printf("  added %s from %s\n", zipName, path)
    }
}

func zipDirRecur(zw *zip.Writer, baseDir string, dirToZip string) {
    dir := filepath.Join(baseDir, dirToZip)
    files, err := ioutil.ReadDir(dir)
    must(err)
    for _, fi := range files {
        if fi.IsDir() {
            zipDirRecur(zw, baseDir, filepath.Join(dirToZip, fi.Name()))
        } else if fi.Mode().IsRegular() {
            zipName := filepath.Join(dirToZip, fi.Name())
            path := filepath.Join(baseDir, zipName)
            zipAddFile(zw, zipName, path)
        } else {
            path := filepath.Join(baseDir, fi.Name())
            s := fmt.Sprintf("%s is not a dir or regular file", path)
            panic(s)
        }
    }
}

func createZipFile(dst string, baseDir string, toZip ...string) {
    removeFile(dst)
    if len(toZip) == 0 {
        panic("must provide toZip args")
    }
    if verbose {
        fmt.Printf("Creating zip file %s\n", dst)
    }
    w, err := os.Create(dst)
    must(err)
    defer w.Close()
    zw := zip.NewWriter(w)
    must(err)
    for _, name := range toZip {
        path := filepath.Join(baseDir, name)
        fi, err := os.Stat(path)
        must(err)
        if fi.IsDir() {
            zipDirRecur(zw, baseDir, name)
        } else if fi.Mode().IsRegular() {
            zipAddFile(zw, name, path)
        } else {
            s := fmt.Sprintf("%s is not a dir or regular file", path)
            panic(s)
        }
    }
    err = zw.Close()
    must(err)
}
```

### wc -l

### wc -l

I like to know how big my programs are as measured by lines of code.

我想知道按代码行数衡量我的程序有多大。

On Unix simple stats can be done with `find . -name "*.go" | xargs wc -l`

在 Unix 上，可以使用 `find 完成简单的统计。 -name "*.go" | xargs wc -l`

Similar  functionality in Go is significantly larger. The good thing is that it's cross-platform, more flexible and once written can be easily added to  more projects.

Go 中的类似功能要大得多。好处是它是跨平台的，更灵活，一旦编写就可以轻松添加到更多项目中。

Different projects want to count different files / directories so I built a flexible system that allows combining (with `and` and `or`) file filter functions.

不同的项目想要计算不同的文件/目录，所以我构建了一个灵活的系统，允许组合（与 `and` 和 `or`）文件过滤器功能。

I wrote a helper library https://github.com/kjk/u.

我写了一个辅助库 https://github.com/kjk/u。

Here's how I use it in a [real project](https://github.com/kjk/notionapi/blob/master/do/wc.go):

这是我在 [真实项目](https://github.com/kjk/notionapi/blob/master/do/wc.go) 中使用它的方法：

`filter` is a file filter function that tells us to count `.go`, `.js`, `.html` and `.css` files in all sub-directories but to  exclude `node_modules` and `tmpdata` directories because they contain files not written by me:

`filter` 是一个文件过滤函数，它告诉我们计算所有子目录中的 `.go`、`.js`、`.html` 和 `.css` 文件，但排除 `node_modules` 和 `tmpdata` 目录，因为它们包含不是我写的文件：

```
package main

import (
    "fmt"

    "github.com/kjk/u"
)

var srcFiles = u.MakeAllowedFileFilterForExts(".go", ".js", ".html", ".css")
var excludeDirs = u.MakeExcludeDirsFilter("node_modules", "tmpdata")
var filter = u.MakeFilterAnd(srcFiles, excludeDirs)

func doLineCount() int {
    stats := u.NewLineStats()
    recursive := true
    err := stats.CalcInDir(".", filter, recursive)
    if err != nil {
        fmt.Printf("doLineCount: stats.wcInDir() failed with '%s'\n", err)
        return 1
    }
    u.PrintLineStats(stats)
    return 0
}
```

# More Go resources

# 更多围棋资源

- [Essential Go](https://www.programming-books.io/essential/go/) is a free, comprehensive book about Go that I maintain 

- [Essential Go](https://www.programming-books.io/essential/go/) 是我维护的一本关于围棋的免费、全面的书籍

                                                            Written on Aug 16 2021.                                        Topics: [go](https://blog.kowalczyk.info/tag/go).

写于 2021 年 8 月 16 日。主题：[go](https://blog.kowalczyk.info/tag/go)。

                                                            [home](https://blog.kowalczyk.info/)

[主页](https://blog.kowalczyk.info/)

                    Found a mistake, have a comment? [Let me know](https://blog.kowalczyk.info/article/4b1f9201181340099b698246857ea98d/using-go-instead-of-bash-for-scripts.html#). 

发现错误，有意见？ [让我知道](https://blog.kowalczyk.info/article/4b1f9201181340099b698246857ea98d/using-go-instead-of-bash-for-scripts.html#)。

