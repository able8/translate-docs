# OS Exec

# 操作系统执行

**[You can find all the code here](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/os-exec)**

**[你可以在这里找到所有代码](https://github.com/quii/learn-go-with-tests/tree/main/q-and-a/os-exec)**

[keith6014](https://www.reddit.com/user/keith6014) asks on [reddit](https://www.reddit.com/r/golang/comments/aaz8ji/testdata_and_function_setup_help/)

[keith6014](https://www.reddit.com/user/keith6014) 在 [reddit](https://www.reddit.com/r/golang/comments/aaz8ji/testdata_and_function_setup_help/) 上提问

> I am executing a command using os/exec.Command() which generated XML data. The command will be executed in a function called GetData().

> 我正在使用生成 XML 数据的 os/exec.Command() 执行命令。该命令将在名为 GetData() 的函数中执行。

> In order to test GetData(), I have some testdata which I created.

> 为了测试 GetData()，我创建了一些测试数据。

> In my _test.go I have a TestGetData which calls GetData() but that will use os.exec, instead I would like for it to use my testdata.

> 在我的 _test.go 中，我有一个调用 GetData() 的 TestGetData，但它将使用 os.exec，而我希望它使用我的 testdata。

> What is a good way to achieve this? When calling GetData should I have a "test" flag mode so it will read a file ie GetData(mode string)?

> 实现这一目标的好方法是什么？当调用 GetData 时，我应该有一个“测试”标志模式，以便它会读取一个文件，即 GetData（模式字符串）？

A few things

一些东西

- When something is difficult to test, it's often due to the separation of concerns not being quite right
- Don't add "test modes" into your code, instead use [Dependency Injection](/dependency-injection.md) so that you can model your dependencies and separate concerns.

- 当某事难以测试时，通常是由于关注点分离不完全正确
- 不要将“测试模式”添加到您的代码中，而是使用 [依赖注入](/dependency-injection.md) 以便您可以对依赖项进行建模并分离关注点。

I have taken the liberty of guessing what the code might look like

我冒昧地猜测代码可能是什么样子

```go
type Payload struct {
    Message string `xml:"message"`
}

func GetData() string {
    cmd := exec.Command("cat", "msg.xml")

    out, _ := cmd.StdoutPipe()
    var payload Payload
    decoder := xml.NewDecoder(out)

    // these 3 can return errors but I'm ignoring for brevity
    cmd.Start()
    decoder.Decode(&payload)
    cmd.Wait()

    return strings.ToUpper(payload.Message)
}
```

- It uses `exec.Command` which allows you to execute an external command to the process
- We capture the output in `cmd.StdoutPipe` which returns us a `io.ReadCloser` (this will become important)
- The rest of the code is more or less copy and pasted from the [excellent documentation](https://golang.org/pkg/os/exec/#example_Cmd_StdoutPipe).
     - We capture any output from stdout into an `io.ReadCloser` and then we `Start` the command and then wait for all the data to be read by calling `Wait`. In between those two calls we decode the data into our `Payload` struct.

- 它使用`exec.Command`，它允许您对进程执行外部命令
- 我们在 `cmd.StdoutPipe` 中捕获输出，它返回一个 `io.ReadCloser`（这将变得很重要）
- 其余代码或多或少是从[优秀文档](https://golang.org/pkg/os/exec/#example_Cmd_StdoutPipe)复制粘贴的。
    - 我们将 stdout 的任何输出捕获到 `io.ReadCloser` 中，然后我们 `Start` 命令，然后通过调用 `Wait` 等待所有数据被读取。在这两次调用之间，我们将数据解码到我们的 `Payload` 结构中。

Here is what is contained inside `msg.xml`

这是 `msg.xml` 中包含的内容

```xml
<payload>
    <message>Happy New Year!</message>
</payload>
```

I wrote a simple test to show it in action

我写了一个简单的测试来展示它的实际效果

```go
func TestGetData(t *testing.T) {
    got := GetData()
    want := "HAPPY NEW YEAR!"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

## Testable code

## 可测试的代码

Testable code is decoupled and single purpose. To me it feels like there are two main concerns for this code

可测试代码是解耦和单一用途的。对我来说，这段代码有两个主要问题

1. Retrieving the raw XML data
2. Decoding the XML data and applying our business logic (in this case `strings.ToUpper` on the `<message>`)

1. 检索原始 XML 数据
2. 解码 XML 数据并应用我们的业务逻辑（在本例中为 `<message>` 上的 `strings.ToUpper`）

The first part is just copying the example from the standard lib.

第一部分只是从标准库中复制示例。

The second part is where we have our business logic and by looking at the code we can see where the "seam" in our logic starts; it's where we get our `io.ReadCloser`. We can use this existing abstraction to separate concerns and make our code testable.

第二部分是我们的业务逻辑，通过查看代码，我们可以看到逻辑中的“接缝”从哪里开始；这就是我们获得 `io.ReadCloser` 的地方。我们可以使用这个现有的抽象来分离关注点并使我们的代码可测试。

**The problem with GetData is the business logic is coupled with the means of getting the XML. To make our design better we need to decouple them**

**GetData 的问题在于业务逻辑与获取 XML 的方法相结合。为了使我们的设计更好，我们需要将它们解耦**

Our `TestGetData` can act as our integration test between our two concerns so we'll keep hold of that to make sure it keeps working.

我们的“TestGetData”可以作为我们两个关注点之间的集成测试，因此我们将保持它以确保它继续工作。

Here is what the newly separated code looks like

这是新分离的代码的样子

```go
type Payload struct {
    Message string `xml:"message"`
}

func GetData(data io.Reader) string {
    var payload Payload
    xml.NewDecoder(data).Decode(&payload)
    return strings.ToUpper(payload.Message)
}

func getXMLFromCommand() io.Reader {
    cmd := exec.Command("cat", "msg.xml")
    out, _ := cmd.StdoutPipe()

    cmd.Start()
    data, _ := ioutil.ReadAll(out)
    cmd.Wait()

    return bytes.NewReader(data)
}

func TestGetDataIntegration(t *testing.T) {
    got := GetData(getXMLFromCommand())
    want := "HAPPY NEW YEAR!"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```

Now that `GetData` takes its input from just an `io.Reader` we have made it testable and it is no longer concerned how the data is retrieved; people can re-use the function with anything that returns an `io.Reader` (which is extremely common). For example we could start fetching the XML from a URL instead of the command line.

现在`GetData` 只从`io.Reader` 获取它的输入，我们已经使它可测试，它不再关心如何检索数据；人们可以将这个函数重用于任何返回 `io.Reader` 的东西（这是非常常见的）。例如，我们可以开始从 URL 而不是命令行获取 XML。

```go
func TestGetData(t *testing.T) {
    input := strings.NewReader(`
<payload>
    <message>Cats are the best animal</message>
</payload>`)

    got := GetData(input)
    want := "CATS ARE THE BEST ANIMAL"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

```

Here is an example of a unit test for `GetData`. 

以下是“GetData”的单元测试示例。

By separating the concerns and using existing abstractions within Go testing our important business logic is a breeze. 

通过分离关注点并在 Go 中使用现有的抽象来测试我们重要的业务逻辑是轻而易举的。

