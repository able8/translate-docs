# [Faking stdin and stdout in Go](https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/)

# [在 Go 中伪造标准输入和标准输出](https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/)

May 02, 2020 at 05:36

In this post I want to discuss *faking* (or *redirecting*) standard input and output (`os.Stdin` and `os.Stdout`) in Go programs. This is often done in tests, but may also be useful in other scenarios.

在这篇文章中，我想讨论 Go 程序中的*伪造*（或*重定向*）标准输入和输出（`os.Stdin` 和 `os.Stdout`）。这通常在测试中完成，但在其他场景中也可能有用。

The basic idea is demonstrated in the following pseudocode:

下面的伪代码演示了基本思想：

```
StartFakingIO(stdin to feed)

FunctionUnderTest(...)

out := GetCapturedOutput()
if out != expected { ... }
```


We assume `FunctionUnderTest` reads from `os.Stdin` and writes to `os.Stdout` directly (this whole post is unnecessary if `FunctionUnderTest` uses dependency injection to take an `io.Reader` and `io.Writer` instead) . Therefore `StartFakingIO` should redirect the `os.Stdin` and `os.Stdout` globals, such that the code of `FunctionUnderTest` remains unchanged.

我们假设`FunctionUnderTest`从`os.Stdin`读取并直接写入`os.Stdout`（如果`FunctionUnderTest`使用依赖注入来代替`io.Reader`和`io.Writer`，则不需要整篇文章）。因此`StartFakingIO` 应该重定向`os.Stdin` 和`os.Stdout` 全局变量，这样`FunctionUnderTest` 的代码保持不变。

## Pipes

## 管道

Since `os.Stdin` and `os.Stdout` are of type `*os.File`, we can't just use `io` interfaces to replace them; we need concrete `*os.File`s. Luckily, this is exactly what the `os.Pipe()` function provides:

由于`os.Stdin` 和`os.Stdout` 属于`*os.File` 类型，我们不能仅仅使用`io` 接口来替换它们；我们需要具体的`*os.File`s。幸运的是，这正是 `os.Pipe()` 函数所提供的：

```
func Pipe() (r *File, w *File, err error)

Pipe returns a connected pair of Files;reads from r return bytes written to
w.It returns the files and an error, if any.
```


In graphic form:

以图形形式：

![Go pipe - write at one end, read from another](https://eli.thegreenplace.net/images/2020/go-pipe.png)

Here's a simple code snippet to demonstrate it:

这是一个简单的代码片段来演示它：

```
func main() {
  r, w, err := os.Pipe()
  if err != nil {
    log.Fatal(err)
  }
  w.Write([]byte("hello"))

  buf := make([]byte, 1024)
  n, err := r.Read(buf)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(buf[:n]))
}
```


This prints "hello".

这将打印“你好”。

If `os.Pipe` reminds you of Unix pipes, that's because it's exactly what it is. [Under the hood](https://golang.org/src/os/pipe_linux.go),`os.Pipe` is a thin wrapper around the `pipe(2)` syscall.

如果`os.Pipe` 让你想起Unix 管道，那是因为它正是它的本质。 [幕后](https://golang.org/src/os/pipe_linux.go)，`os.Pipe`是对`pipe(2)` 系统调用的一个薄包装。

Given this pipe construct, it's easy to come up with a faking scheme for `os.Stdout`:

鉴于这种管道结构，很容易为 `os.Stdout` 提出一个伪造方案：

```
package main

import (
  "fmt"
  "log"
  "os"
)

func main() {
  r, w, err := os.Pipe()
  if err != nil {
    log.Fatal(err)
  }
  origStdout := os.Stdout
  os.Stdout = w

  fmt.Print("hello to stdout")

  buf := make([]byte, 1024)
  n, err := r.Read(buf)
  if err != nil {
    log.Fatal(err)
  }

  // Restore
  os.Stdout = origStdout

  fmt.Println("Written to stdout:", string(buf[:n]))
}
```


Faking `os.Stdin` is very similar, except the direction is inverted. See the full code at the end of the post.

伪造`os.Stdin` 非常相似，只是方向相反。请参阅帖子末尾的完整代码。

This approach is the basis of most stdio faking packages you'll find online. Unfortunately, this approach has a serious problem in some scenarios.

这种方法是您可以在网上找到的大多数 stdio 伪造软件包的基础。不幸的是，这种方法在某些场景中存在严重问题。

## Pipe buffers

## 管道缓冲区

Pipes on Linux have limited capacities. Here's what `man 7 pipe` has to say about it:

Linux 上的管道容量有限。以下是“man 7 pipe”对此的评价：

> A pipe has a limited capacity. If the pipe is full, then a write(2) will block or fail, depending on whether the O_NONBLOCK flag is set (see below). Different implementations have different limits for the pipe capacity. Applications should not  rely  on  a particular capacity: an application should be designed so that a reading process consumes data as soon as it is available, so that a writing process does not remain blocked.
>
> In Linux versions before 2.6.11, the capacity of a pipe was the same as the system page size (e.g., 4096 bytes on i386). Since Linux 2.6.11, the pipe capacity is 65536 bytes. Since Linux 2.6.35, the default pipe capacity is 65536 bytes, but the capacity can be queried and set using the fcntl(2) F_GETPIPE_SZ and F_SETPIPE_SZ operations. See fcntl(2) for more information.

> 管道的容量有限。如果管道已满，则 write(2) 将阻塞或失败，具体取决于是否设置了 O_NONBLOCK 标志（见下文）。不同的实现对管道容量有不同的限制。应用程序不应依赖于特定的容量：应用程序的设计应使读取进程在数据可用时立即消耗数据，以便写入进程不会一直阻塞。
>
> 在 2.6.11 之前的 Linux 版本中，管道的容量与系统页面大小相同（例如，i386 上为 4096 字节）。从 Linux 2.6.11 开始，管道容量为 65536 字节。从 Linux 2.6.35 开始，默认管道容量为 65536 字节，但可以使用 fcntl(2) F_GETPIPE_SZ 和 F_SETPIPE_SZ 操作来查询和设置容量。有关更多信息，请参阅 fcntl(2)。

We can easily test this by extending the previous example to print out much more to `os.Stdout` before trying to read from it:

我们可以通过扩展前面的例子来轻松测试这一点，在尝试读取它之前将更多内容打印到 `os.Stdout`：

```
func main() {
  r, w, err := os.Pipe()
  if err != nil {
    log.Fatal(err)
  }
  origStdout := os.Stdout
  os.Stdout = w

  for i := 0;i < 5000;i++ {
    fmt.Print("hello to stdout")
  }

  buf := make([]byte, 1024)
  n, err := r.Read(buf)
  if err != nil {
    log.Fatal(err)
  }

  // Restore
  os.Stdout = origStdout

  fmt.Println("Written to stdout:", string(buf[:n]))
}
```

It's exactly the same code, except that now we print out "hello to stdout" 5000 times, for a total of 75,000 bytes, which should overflow the buffer. 
这是完全相同的代码，除了现在我们打印出 5000 次“hello to stdout”，总共 75,000 个字节，这应该会溢出缓冲区。

Indeed, if you run this program, it hangs. Sending SIGQUIT to the program shows it's stuck in the call to `fmt.Print`. Without anything reading from the pipe's other end, the program can't proceed once the pipe buffer has been filled. Obviously, this problem may not apply to most scenarios - you don't typically print out this much data, especially in unit tests. But it's still fairly common to get bitten by it.

事实上，如果你运行这个程序，它就会挂起。向程序发送 SIGQUIT 表明它卡在对 `fmt.Print` 的调用中。如果没有从管道的另一端读取任何内容，一旦管道缓冲区被填满，程序就无法继续。显然，这个问题可能不适用于大多数场景——您通常不会打印出这么多数据，尤其是在单元测试中。但被它咬伤仍然很常见。

To solve this problem, we have to ensure that something is reading from the pipe continuously, to prevent the overflow. This can be easily done with a separate goroutine, as the next section will demonstrate.

为了解决这个问题，我们必须确保不断有东西从管道中读取，以防止溢出。这可以使用单独的 goroutine 轻松完成，下一节将演示。

## A complete stdio faker

##一个完整的stdio faker

I'll now show the code a complete "stdio faker" type, that will enable us writing code just like the "basic idea" pseudocode at the top of this post. [The full code with tests and an example is available on GitHub](https://github.com/eliben/code-for-blog/tree/master/2020/go-fake-stdio). Let's start with the type; all its fields are private, since the user code only interacts with the faker via its methods:

我现在将向代码展示一个完整的“stdio faker”类型，这将使​​我们能够像本文顶部的“基本思想”伪代码一样编写代码。 [包含测试和示例的完整代码可在 GitHub 上找到](https://github.com/eliben/code-for-blog/tree/master/2020/go-fake-stdio)。让我们从类型开始；它的所有字段都是私有的，因为用户代码仅通过其方法与伪造者进行交互：

```
// FakeStdio can be used to fake stdin and capture stdout.
// Between creating a new FakeStdio and calling ReadAndRestore on it,
// code reading os.Stdin will get the contents of stdinText passed to New.
// Output to os.Stdout will be captured and returned from ReadAndRestore.
// FakeStdio is not reusable;don't attempt to use it after calling
// ReadAndRestore, but it should be safe to create a new FakeStdio.
type FakeStdio struct {
  origStdout   *os.File
  stdoutReader *os.File

  outCh chan []byte

  origStdin   *os.File
  stdinWriter *os.File
}
```


This is the constructor:

这是构造函数：

```
func New(stdinText string) (*FakeStdio, error) {
  // Pipe for stdin.
  //
  //                 ======
  //  w ------------->||||------> r
  // (stdinWriter)   ======      (os.Stdin)
  stdinReader, stdinWriter, err := os.Pipe()
  if err != nil {
    return nil, err
  }

  // Pipe for stdout.
  //
  //               ======
  //  w ----------->||||------> r
  // (os.Stdout)   ======      (stdoutReader)
  stdoutReader, stdoutWriter, err := os.Pipe()
  if err != nil {
    return nil, err
  }

  origStdin := os.Stdin
  os.Stdin = stdinReader

  _, err = stdinWriter.Write([]byte(stdinText))
  if err != nil {
    stdinWriter.Close()
    os.Stdin = origStdin
    return nil, err
  }

  origStdout := os.Stdout
  os.Stdout = stdoutWriter

  outCh := make(chan []byte)

  // This goroutine reads stdout into a buffer in the background.
  go func() {
    var b bytes.Buffer
    if _, err := io.Copy(&b, stdoutReader);err != nil {
      log.Println(err)
    }
    outCh <- b.Bytes()
  }()

  return &FakeStdio{
    origStdout:   origStdout,
    stdoutReader: stdoutReader,
    outCh:        outCh,
    origStdin:    origStdin,
    stdinWriter:  stdinWriter,
  }, nil
}
```


Of particular interest in this code:


1. The ASCII diagrams showing how the different pipes are hooked together.
2. A goroutine that runs in the background throughout the lifetime of a `FakeStdio`. This goroutine continuously reads from the reading end of the fake stdout to drain the buffer, ensuring that large writes don't block.

对这段代码特别感兴趣：
1. ASCII 图显示了不同的管道是如何连接在一起的。
2. 在“FakeStdio”的整个生命周期中在后台运行的 goroutine。该 goroutine 不断从伪 stdout 的读取端读取以耗尽缓冲区，确保不会阻塞大量写入。

And this is the `ReadAndRestore` method:

这是“ReadAndRestore”方法：

```
// ReadAndRestore collects all captured stdout and returns it;it also restores
// os.Stdin and os.Stdout to their original values.
func (sf *FakeStdio) ReadAndRestore() ([]byte, error) {
  if sf.stdoutReader == nil {
    return nil, fmt.Errorf("ReadAndRestore from closed FakeStdio")
  }

  // Close the writer side of the faked stdout pipe.This signals to the
  // background goroutine that it should exit.
  os.Stdout.Close()
  out := <-sf.outCh

  os.Stdout = sf.origStdout
  os.Stdin = sf.origStdin

  if sf.stdoutReader != nil {
    sf.stdoutReader.Close()
    sf.stdoutReader = nil
  }

  if sf.stdinWriter != nil {
    sf.stdinWriter.Close()
    sf.stdinWriter = nil
  }

  return out, nil
}
```


[Here's a usage example](https://github.com/eliben/code-for-blog/blob/master/2020/go-fake-stdio/example_test.go):

[这是一个使用示例](https://github.com/eliben/code-for-blog/blob/master/2020/go-fake-stdio/example_test.go)：

```
func ExampleFakeInOut() {
  // Create a new fakestdio with some input to feed into Stdin.
  fs, err := New("input text")
  if err != nil {
    log.Fatal(err)
  }
  var scanned string
  fmt.Scanf("%s", &scanned)

  // Emit text to Stdout - it will be captured.
  fmt.Print("some output")

  b, err := fs.ReadAndRestore()
  if err != nil {
    log.Fatal(err)
  }

  // This will go to the actual os.Stdout because we're no longer capturing.
  fmt.Printf("Scanned: %q, Captured: %q", scanned, string(b))

  // Output: Scanned: "input", Captured: "some output"
}
```


## Variations

## 变化

Writing a fully general package for faking stdio requires handling a whole bunch of different requirements and scenarios; I'm not attempting to do so here, but I'll mention some.

为伪造 stdio 编写一个完全通用的包需要处理一大堆不同的需求和场景；我不打算在这里这样做，但我会提到一些。

My implementation has an additional method I haven't shown so far:

我的实现有一个额外的方法，到目前为止我还没有展示：

```
// CloseStdin closes the fake stdin.This may be necessary if the process has
// logic for reading stdin until EOF; otherwise such code would block forever.
func (sf *FakeStdio) CloseStdin() {
   if sf.stdinWriter != nil {
     sf.stdinWriter.Close()
     sf.stdinWriter = nil
   }
}
```

As its comment explains, this is important to test code that reads `os.Stdin` until it's closed - think a standard Unix line filter program.

正如它的注释所解释的那样，这对于测试读取 `os.Stdin` 直到它关闭的代码很重要——想想一个标准的 Unix 行过滤程序。

Another feature that could be added is a method to feed more data to the faked `os.Stdin`; in the current approach, the only data to stdin is provided in the constructor. To test interactive code we may want to send more data to stdin after we've seen some of the output. This should be fairly easy to add - try it as an exercise!

另一个可以添加的功能是一种向伪造的 `os.Stdin` 提供更多数据的方法；在当前方法中，构造函数中提供了 stdin 的唯一数据。为了测试交互式代码，我们可能希望在看到一些输出后向 stdin 发送更多数据。这应该很容易添加 - 作为练习试试吧！

A similar variant is reading captured stdout data before `FakeStdio` is restored; this could also be useful for testing interactive code. For this to work, a slightly more significant rework of the code would be required. The stdout draining goroutine will need to have its `io.Copy` broken up to individual `Read` operations, and a synchronized way to access the buffer it's filling will have to be added.

一个类似的变体是在“FakeStdio”恢复之前读取捕获的标准输出数据；这对于测试交互式代码也很有用。为此，需要对代码进行稍微更重要的返工。 stdout 排空 goroutine 需要将其 `io.Copy` 分解为单独的 `Read` 操作，并且必须添加一种同步方式来访问它正在填充的缓冲区。

So far the code takes care not to overflow stdout; but what about stdin? If you need to feed more than 64KiB into stdin, the current approach will hang. This requires a similar goroutine, but on the user code side.

到目前为止，代码注意不要溢出标准输出；但是标准输入呢？如果您需要向 stdin 输入超过 64KiB 的数据，则当前方法将挂起。这需要一个类似的 goroutine，但在用户代码方面。

Finally, this code only handles `os.Stdin` and `os.Stdout`; there's also `os.Stderr` we could capture. That should be trivial to add, if needed.

最后，这段代码只处理 `os.Stdin` 和 `os.Stdout`；还有我们可以捕获的`os.Stderr`。如果需要，添加这应该是微不足道的。


