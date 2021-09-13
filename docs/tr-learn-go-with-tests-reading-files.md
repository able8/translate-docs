# Reading files

# 读取文件

- **[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/reading-files)**
- [Here is a video of me working through the problem and taking questions from the Twitch stream](https://www.youtube.com/watch?v=nXts4dEJnkU)

- **[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/reading-files)**
- [这是我解决问题并从 Twitch 流中回答问题的视频](https://www.youtube.com/watch?v=nXts4dEJnkU)

In this chapter we're going to learn how to read some files, get some data out of them, and do something useful.

在本章中，我们将学习如何读取一些文件，从中获取一些数据，并做一些有用的事情。

Pretend you're working with your friend to create some blog software. The idea is an author will write their posts in markdown, with some metadata at the top of the file. On startup, the web server will read a folder to create some `Post`s, and then a separate `NewHandler` function will use those `Post`s as a datasource for the blog's webserver.

假设您正在与您的朋友一起创建一些博客软件。这个想法是作者将用 Markdown 写他们的帖子，在文件的顶部有一些元数据。在启动时，Web 服务器将读取一个文件夹以创建一些 `Post`，然后一个单独的 `NewHandler` 函数将使用这些 `Post` 作为博客 Web 服务器的数据源。

We've been asked to create the package that converts a given folder of blog post files into a collection of `Post`s.

我们被要求创建一个包，将给定的博客帖子文件文件夹转换为一个“帖子”集合。

### Example data

### 示例数据

hello world.md
```markdown
Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd, go
---
Hello world!

The body of posts starts after the `---`
```

### Expected data

### 预期数据

```go
type Post struct {
    Title, Description, Body string
    Tags []string
}
```

## Iterative, test-driven development

## 迭代的、测试驱动的开发

We'll take an iterative approach where we're always taking simple, safe steps toward our goal.

我们将采用迭代方法，始终朝着我们的目标采取简单、安全的步骤。

This requires us to break up our work, but we should be careful not to fall into the trap of taking a ["bottom up"](https://en.wikipedia.org/wiki/Top-down_and_bottom-up_design) approach.

这需要我们分解我们的工作，但我们应该小心不要陷入采取[“自下而上”](https://en.wikipedia.org/wiki/Top-down_and_bottom-up_design)方法的陷阱。

We should not trust our over-active imaginations when we start work. We could be tempted into making some kind of abstraction that is only validated once we stick everything together, such as some kind of `BlogPostFileParser`.

当我们开始工作时，我们不应该相信我们过度活跃的想象力。我们可能会尝试制作某种抽象，这种抽象只有在我们将所有东西粘在一起后才能得到验证，例如某种“BlogPostFileParser”。

This is _not_ iterative and is missing out on the tight feedback loops that TDD is supposed to bring us.

这_不是_迭代，并且错过了 TDD 应该给我们带来的紧密反馈循环。

Kent Beck says:

肯特贝克 说：

> Optimism is an occupational hazard of programming. Feedback is the treatment.

> 乐观是编程的职业危害。反馈就是治疗。

Instead, our approach should strive to be as close to delivering _real_ consumer value as quickly as possible (often called a "happy path"). Once we have delivered a small amount of consumer value end-to-end, further iteration of the rest of the requirements is usually straightforward.

相反，我们的方法应该努力尽可能快地交付 _real_ 消费者价值（通常称为“快乐路径”）。一旦我们端到端地交付了少量消费者价值，其余需求的进一步迭代通常很简单。

## Thinking about the kind of test we want to see

## 思考我们想要看到的测试类型

Let's remind ourselves of our mindset and goals when starting:

让我们在开始时提醒自己我们的心态和目标：

- **Write the test we want to see**. Think about how we'd like to use the code we're going to write from a consumer's point of view.
- Focus on _what_ and _why_, but don't get distracted by _how_.

- **写我们想看的测试**。从消费者的角度考虑我们希望如何使用我们将要编写的代码。
- 专注于_what_和_why_，但不要被_how_分心。

Our package needs to offer a function that can be pointed at a folder, and return us some posts.

我们的包需要提供一个可以指向文件夹的函数，并返回一些帖子给我们。

```go
var posts []blogposts.Post
posts = blogposts.NewPostsFromFS("some-folder")
```

To write a test around this, we'd need some kind of test folder with some example posts in it. _There's nothing terribly wrong with this_, but you are making some trade-offs:

要围绕此编写测试，我们需要某种测试文件夹，其中包含一些示例帖子。 _这并没有什么大错_，但是您正在做出一些权衡：

- for each test you may need to create new files to test a particular behaviour
- some behaviour will be challenging to test, such as failing to load files
- the tests will run a little slower because they will need to access the file system

- 对于每个测试，您可能需要创建新文件来测试特定行为
- 某些行为将难以测试，例如无法加载文件
- 测试运行速度会慢一些，因为它们需要访问文件系统

We're also unnecessarily coupling ourselves to a specific implementation of the file system.

我们也不必要地将自己耦合到文件系统的特定实现。

### File system abstractions introduced in Go 1.16

### Go 1.16 中引入的文件系统抽象

Go 1.16 introduced an abstraction for file systems; the [io/fs](https://golang.org/pkg/io/fs/) package.

Go 1.16 引入了文件系统的抽象； [io/fs](https://golang.org/pkg/io/fs/) 包。

> Package fs defines basic interfaces to a file system. A file system can be provided by the host operating system but also by other packages.

> 包 fs 定义了文件系统的基本接口。文件系统可以由主机操作系统提供，也可以由其他包提供。

This lets us loosen our coupling to a specific file system, which will then let us inject different implementations according to our needs.

这让我们可以放松与特定文件系统的耦合，然后让我们根据需要注入不同的实现。

> [On the producer side of the interface, the new embed.FS type implements fs.FS, as does zip.Reader. The new os.DirFS function provides an implementation of fs.FS backed by a tree of operating system files.](https://golang.org/doc/go1.16#fs) 

> [在接口的生产者端，新的 embed.FS 类型实现了 fs.FS，zip.Reader 也是如此。新的 os.DirFS 函数提供了由操作系统文件树支持的 fs.FS 实现。](https://golang.org/doc/go1.16#fs)

If we use this interface, users of our package have a number of options baked-in to the standard library to use. Learning to leverage interfaces defined in Go's standard library (eg `io.fs`, [`io.Reader`](https://golang.org/pkg/io/#Reader), [`io.Writer`](https://golang.org/pkg/io/#Writer)), is vital to writing loosely coupled packages. These packages can then be re-used in contexts different to those you imagined, with minimal fuss from your consumers.

如果我们使用这个接口，我们包的用户有许多内置到标准库中的选项可以使用。学习利用 Go 标准库中定义的接口（例如 `io.fs`、[`io.Reader`](https://golang.org/pkg/io/#Reader)、[`io.Writer`](https://golang.org/pkg/io/#Writer))，对于编写松散耦合的包至关重要。然后，这些包可以在与您想象的环境不同的环境中重复使用，而您的消费者不会大惊小怪。

In our case, maybe our consumer wants the posts to be embedded into the Go binary rather than files in a "real" filesystem? Either way, _our code doesn't need to care_.

在我们的例子中，也许我们的消费者希望将帖子嵌入 Go 二进制文件而不是“真实”文件系统中的文件？无论哪种方式，_我们的代码都不需要关心_。

For our tests, the package [testing/fstest](https://golang.org/pkg/testing/fstest/) offers us an implementation of [io/FS](https://golang.org/pkg/io/fs/#FS) to use, similar to the tools we're familiar with in [net/http/httptest](https://golang.org/pkg/net/http/httptest/).

对于我们的测试，包 [testing/fstest](https://golang.org/pkg/testing/fstest/) 为我们提供了 [io/FS](https://golang.org/pkg/io/fs/#FS) 来使用，类似于我们在 [net/http/httptest](https://golang.org/pkg/net/http/httptest/) 中熟悉的工具。

Given this information, the following feels like a better approach,

鉴于这些信息，以下感觉是更好的方法，

```go
var posts blogposts.Post
posts = blogposts.NewPostsFromFS(someFS)
```

## Write the test first

## 先写测试

We should keep scope as small and useful as possible. If we prove that we can read all the files in a directory, that will be a good start. This will give us confidence in the software we're writing. We can check that the count of `[]Post` returned is the same as the number of files in our fake file system.

我们应该保持范围尽可能小和有用。如果我们证明我们可以读取目录中的所有文件，那将是一个好的开始。这将使我们对正在编写的软件充满信心。我们可以检查返回的 `[]Post` 的数量是否与我们的假文件系统中的文件数量相同。

Create a new project to work through this chapter.

创建一个新项目来完成本章。

- `mkdir blogposts`
- `cd blogposts`
- `go mod init github.com/{your-name}/blogposts`
- `touch blogposts_test.go`

- 

```go
package blogposts_test

import (
    "testing"
    "testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
    fs := fstest.MapFS{
        "hello world.md":  {Data: []byte("hi")},
        "hello-world2.md": {Data: []byte("hola")},
    }

    posts := blogposts.NewPostsFromFS(fs)

    if len(posts) != len(fs) {
        t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
    }
}

```

Notice that the package of our test is `blogposts_test`. Remember, when TDD is practiced well we take a _consumer-driven_ approach: we don't want to test internal details because _consumers_ don't care about them. By appending `_test` to our intended package name, we only access exported members from our package - just like a real user of our package.

请注意，我们的测试包是“blogposts_test”。请记住，当 TDD 实践得很好时，我们采取_消费者驱动_方法：我们不想测试内部细节，因为_消费者_不关心它们。通过将 `_test` 附加到我们想要的包名称，我们只能访问包中导出的成员 - 就像我们包的真实用户一样。

We've imported [`testing/fstest`](https://golang.org/pkg/testing/fstest/) which gives us access to the [`fstest.MapFS`](https://golang.org/pkg/testing/fstest/#MapFS) type. Our fake file system will pass `fstest.MapFS` to our package.

我们已经导入了 [`testing/fstest`](https://golang.org/pkg/testing/fstest/) 这让我们可以访问 [`fstest.MapFS`](https://golang.org/pkg/testing/fstest/#MapFS) 类型。我们的假文件系统会将 `fstest.MapFS` 传递给我们的包。

> A MapFS is a simple in-memory file system for use in tests, represented as a map from path names (arguments to Open) to information about the files or directories they represent.

> MapFS 是一个用于测试的简单内存文件系统，表示为从路径名（Open 的参数）到有关它们所代表的文件或目录的信息的映射。

This feels simpler than maintaining a folder of test files, and it will execute quicker.

这感觉比维护一个测试文件的文件夹更简单，而且它会执行得更快。

Finally, we codified the usage of our API from a consumer's point of view, then checked if it creates the correct number of posts.

最后，我们从消费者的角度对 API 的使用进行了编码，然后检查它是否创建了正确数量的帖子。

## Try to run the test

## 尝试运行测试

```
./blogpost_test.go:15:12: undefined: blogposts
```

## Write the minimal amount of code for the test to run and _check the failing test output_

## 为要运行的测试编写最少的代码并_检查失败的测试输出_

The package doesn't exist. Create a new file `blogposts.go` and put `package blogposts` inside it. You'll need to then import that package into your tests. For me, the imports now look like:

该包不存在。创建一个新文件 `blogposts.go` 并将 `package blogposts` 放入其中。然后您需要将该包导入到您的测试中。对我来说，进口现在看起来像：

```go
import (
    blogposts "github.com/quii/learn-go-with-tests/reading-files"
    "testing"
    "testing/fstest"
)
```

Now the tests won't compile because our new package does not have a `NewPostsFromFS` function, that returns some kind of collection.

现在测试将无法编译，因为我们的新包没有返回某种集合的 `NewPostsFromFS` 函数。

```
./blogpost_test.go:16:12: undefined: blogposts.NewPostsFromFS
```

This forces us to make the skeleton of our function to make the test run. Remember not to overthink the code at this point; we're only trying to get a running test, and to make sure it fails as we'd expect. If we skip this step we may skip over assumptions and, write a test which is not useful.

这迫使我们制作函数的骨架来运行测试。记住此时不要过度考虑代码；我们只是试图进行运行测试，并确保它按我们预期的那样失败。如果我们跳过这一步，我们可能会跳过假设并编写一个无用的测试。

```go
package blogposts

import "testing/fstest"

type Post struct {

}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
    return nil
}
```

The test should now correctly fail

测试现在应该正确失败

```
=== RUN   TestNewBlogPosts
    blogposts_test.go:48: got 0 posts, wanted 2 posts
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We _could_ ["slime"](https://deniseyu.github.io/leveling-up-tdd/) this to make it pass:

我们_可以_ ["slime"](https://deniseyu.github.io/leveling-up-tdd/) 让它通过：

```go
func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
    return []Post{{},{}}
}
```

But, as Denise Yu wrote: 

但是，正如 Denise Yu 所写：

>Sliming is useful for giving a “skeleton” to your object. Designing an interface and executing logic are two concerns, and sliming tests strategically lets you focus on one at a time.

>Sliming 可用于为您的对象提供“骨架”。设计界面和执行逻辑是两个问题，而瘦身测试战略性地让您一次专注于一个。

We already have our structure. So, what do we do instead?

我们已经有了我们的结构。那么，我们该怎么做呢？

As we've cut scope, all we need to do is read the directory and create a post for each file we encounter. We don't have to worry about opening files and parsing them just yet.

由于我们已经缩小了范围，我们需要做的就是读取目录并为我们遇到的每个文件创建一个帖子。我们不必担心打开文件和解析它们。

```go
func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
    dir, _ := fs.ReadDir(fileSystem, ".")
    var posts []Post
    for range dir {
        posts = append(posts, Post{})
    }
    return posts
}
```

[`fs.ReadDir`](https://golang.org/pkg/io/fs/#ReadDir) reads a directory inside a given `fs.FS` returning [`[]DirEntry`](https://golang.org/pkg/io/fs/#DirEntry).

[`fs.ReadDir`](https://golang.org/pkg/io/fs/#ReadDir) 读取给定 `fs.FS` 中的目录，返回 [`[]DirEntry`](https://golang.org/pkg/io/fs/#DirEntry)。

Already our idealised view of the world has been foiled because errors can happen, but remember now our focus is _making the test pass_, not changing design, so we'll ignore the error for now.

我们的理想化世界观已经被挫败，因为错误可能发生，但现在请记住，我们的重点是_使测试通过_，而不是改变设计，所以我们现在将忽略错误。

The rest of the code is straightforward: iterate over the entries, create a `Post` for each one and, return the slice.

其余代码很简单：迭代条目，为每个条目创建一个“Post”，然后返回切片。

## Refactor

## 重构

Even though our tests are passing, we can't use our new package outside of this context, because it is coupled to a concrete implementation `fstest.MapFS`. But, it doesn't have to be. Change the argument to our `NewPostsFromFS` function to accept the interface from the standard library.

即使我们的测试通过了，我们也不能在这个上下文之外使用我们的新包，因为它耦合到一个具体的实现 `fstest.MapFS`。但是，不必如此。将参数更改为我们的 `NewPostsFromFS` 函数以接受来自标准库的接口。

```go
func NewPostsFromFS(fileSystem fs.FS) []Post {
    dir, _ := fs.ReadDir(fileSystem, ".")
    var posts []Post
    for range dir {
        posts = append(posts, Post{})
    }
    return posts
}
```

Re-run the tests: everything should be working.

重新运行测试：一切正常。

### Error handling

### 错误处理

We parked error handling earlier when we focused on making the happy-path work. Before continuing to iterate on the functionality, we should acknowledge that errors can happen when working with files. Beyond reading the directory, we can run into problems when we open individual files. Let's change our API (via our tests first, naturally) so that it can return an `error`.

当我们专注于使快乐路径工作时，我们更早地停止了错误处理。在继续迭代功能之前，我们应该承认在处理文件时可能会发生错误。除了读取目录之外，我们在打开单个文件时还会遇到问题。让我们改变我们的 API（首先通过我们的测试，自然地），以便它可以返回一个 `error`。

```go
func TestNewBlogPosts(t *testing.T) {
    fs := fstest.MapFS{
        "hello world.md":  {Data: []byte("hi")},
        "hello-world2.md": {Data: []byte("hola")},
    }

    posts, err := blogposts.NewPostsFromFS(fs)

    if err != nil {
        t.Fatal(err)
    }

    if len(posts) != len(fs) {
        t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
    }
}
```

Run the test: it should complain about the wrong number of return values. Fixing the code is straightforward.

运行测试：它应该抱怨返回值的数量错误。修复代码很简单。

```go
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
    dir, err := fs.ReadDir(fileSystem, ".")
    if err != nil {
        return nil, err
    }
    var posts []Post
    for range dir {
        posts = append(posts, Post{})
    }
    return posts, nil
}
```

This will make the test pass. The TDD practitioner in you might be annoyed we didn't see a failing test before writing the code to propagate the error from `fs.ReadDir`. To do this "properly", we'd need a new test where we inject a failing `fs.FS` test-double to make `fs.ReadDir` return an `error`.

这将使测试通过。您中的 TDD 从业者可能会感到恼火，我们在编写代码以从 fs.ReadDir 传播错误之前没有看到失败的测试。为了“正确地”做到这一点，我们需要一个新的测试，在其中我们注入一个失败的 `fs.FS` test-double 以使 `fs.ReadDir` 返回一个 `error`。

```go
type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
    return nil, errors.New("oh no, i always fail")
}

// later
_, err := blogposts.NewPostsFromFS(StubFailingFS{})
```

This should give you confidence in our approach. The interface we're using has one method, which makes creating test-doubles to test different scenarios trivial.

这应该会让您对我们的方法充满信心。我们使用的接口有一种方法，这使得创建测试替身来测试不同的场景变得微不足道。

In some cases, testing error handling is the pragmatic thing to do but, in our case, we're not doing anything _interesting_ with the error, we're just propagating it, so it's not worth the hassle of writing a new test.

在某些情况下，测试错误处理是务实的事情，但在我们的例子中，我们没有对错误做任何_有趣的_，我们只是传播它，所以编写新测试的麻烦不值得。

Logically, our next iterations will be around expanding our `Post` type so that it has some useful data.

从逻辑上讲，我们的下一次迭代将围绕扩展我们的“Post”类型，以便它有一些有用的数据。

## Write the test first

## 先写测试

We'll start with the first line in the proposed blog post schema, the title field.

我们将从建议的博客文章架构中的第一行标题字段开始。

We need to change the contents of the test files so they match what was specified, and then we can make an assertion that it is parsed correctly.
```go
func TestNewBlogPosts(t *testing.T) {
    fs := fstest.MapFS{
        "hello world.md":  {Data: []byte("Title: Post 1")},
        "hello-world2.md": {Data: []byte("Title: Post 2")},
    }

    // rest of test code cut for brevity
    got := posts[0]
    want := blogposts.Post{Title: "Post 1"}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %+v, want %+v", got, want)
    }
}
```

## Try to run the test
```
./blogpost_test.go:58:26: unknown field 'Title' in struct literal of type blogposts.Post
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Add the new field to our `Post` type so that the test will run

将新字段添加到我们的 `Post` 类型，以便测试运行

```go
type Post struct {
    Title string
}
```

Re-run the test, and you should get a clear, failing test

重新运行测试，你应该得到一个清晰的、失败的测试

```
=== RUN   TestNewBlogPosts
=== RUN   TestNewBlogPosts/parses_the_post
    blogpost_test.go:61: got {Title:}, want {Title:Post 1}
```

## Write enough code to make it pass

## 编写足够的代码使其通过

We'll need to open each file and then extract the title

我们需要打开每个文件，然后提取标题

```go
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
    dir, err := fs.ReadDir(fileSystem, ".")
    if err != nil {
        return nil, err
    }
    var posts []Post
    for _, f := range dir {
        post, err := getPost(fileSystem, f)
        if err != nil {
            return nil, err //todo: needs clarification, should we totally fail if one file fails?or just ignore?
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
    postFile, err := fileSystem.Open(f.Name())
    if err != nil {
        return Post{}, err
    }
    defer postFile.Close()

    postData, err := io.ReadAll(postFile)
    if err != nil {
        return Post{}, err
    }

    post := Post{Title: string(postData)[7:]}
    return post, nil
}
```

Remember our focus at this point is not to write elegant code, it's just to get to a point where we have working software.

请记住，此时我们的重点不是编写优雅的代码，而只是为了让我们拥有可运行的软件。

Even though this feels like a small increment forward it still required us to write a fair amount of code and make some assumptions in respect to error handling. This would be a point where you should talk to your colleagues and decide the best approach.

尽管这感觉像是向前的一个小增量，但它仍然需要我们编写大量代码并就错误处理做出一些假设。这将是您应该与您的同事交谈并决定最佳方法的地方。

The iterative approach has given us fast feedback that our understanding of the requirements is incomplete.

迭代方法给了我们快速的反馈，表明我们对需求的理解是不完整的。

`fs.FS` gives us a way of opening a file within it by name with its `Open` method. From there we read the data from the file and, for now, we do not need any sophisticated parsing, just cutting out the `Title: ` text by slicing the string.

`fs.FS` 为我们提供了一种使用 `Open` 方法按名称打开文件的方法。从那里我们从文件中读取数据，现在，我们不需要任何复杂的解析，只需通过对字符串进行切片来切掉 `Title:` 文本。

## Refactor

## 重构

Separating the 'opening file code' from the 'parsing file contents code' will make the code simpler to understand and work with.

将“打开文件代码”与“解析文件内容代码”分开将使代码更易于理解和使用。

```go
func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
    postFile, err := fileSystem.Open(f.Name())
    if err != nil {
        return Post{}, err
    }
    defer postFile.Close()
    return newPost(postFile)
}

func newPost(postFile fs.File) (Post, error) {
    postData, err := io.ReadAll(postFile)
    if err != nil {
        return Post{}, err
    }

    post := Post{Title: string(postData)[7:]}
    return post, nil
}
```

When you refactor out new functions or methods, take care and think about the arguments. You're designing here, and are free to think deeply about what is appropriate because you have passing tests. Think about coupling and cohesion. In this case you should ask yourself:

当您重构新函数或方法时，请注意并考虑参数。您在这里进行设计，并且可以自由地深入思考什么是合适的，因为您已经通过了测试。考虑耦合和内聚。在这种情况下，您应该问自己：

> Does `newPost` have to be coupled to an `fs.File` ? Do we use all the methods and data from this type? What do we _really_ need?

> `newPost` 必须与 `fs.File` 耦合吗？我们是否使用了这种类型的所有方法和数据？我们_真正_需要什么？

In our case we only use it as an argument to `io.ReadAll` which needs an `io.Reader`. So we should loosen the coupling in our function and ask for an `io.Reader`.

在我们的例子中，我们只将它用作需要一个 `io.Reader` 的 `io.ReadAll` 的参数。所以我们应该放松我们函数中的耦合并要求一个`io.Reader`。

```go
func newPost(postFile io.Reader) (Post, error) {
    postData, err := io.ReadAll(postFile)
    if err != nil {
        return Post{}, err
    }

    post := Post{Title: string(postData)[7:]}
    return post, nil
}
```

You can make a similar argument for our `getPost` function, which takes an `fs.DirEntry` argument but simply calls `Name()` to get the file name. We don't need all that; let's decouple from that type and pass the file name through as a string. Here's the fully refactored code:

你可以为我们的 `getPost` 函数创建一个类似的参数，它接受一个 `fs.DirEntry` 参数，但只需调用 `Name()` 来获取文件名。我们不需要所有这些；让我们从该类型中解耦并将文件名作为字符串传递。这是完全重构的代码：

```go
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
    dir, err := fs.ReadDir(fileSystem, ".")
    if err != nil {
        return nil, err
    }
    var posts []Post
    for _, f := range dir {
        post, err := getPost(fileSystem, f.Name())
        if err != nil {
            return nil, err //todo: needs clarification, should we totally fail if one file fails?or just ignore?
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
    postFile, err := fileSystem.Open(fileName)
    if err != nil {
        return Post{}, err
    }
    defer postFile.Close()
    return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
    postData, err := io.ReadAll(postFile)
    if err != nil {
        return Post{}, err
    }

    post := Post{Title: string(postData)[7:]}
    return post, nil
}
```

From now on, most of our efforts can be neatly contained within `newPost`. The concerns of opening and iterating over files are done, and now we can focus on extracting the data for our `Post` type. Whilst not technically necessary, files are a nice way to logically group related things together, so I moved the `Post` type and `newPost` into a new `post.go` file.

从现在开始，我们的大部分工作都可以整齐地包含在 `newPost` 中。打开和迭代文件的问题已经完成，现在我们可以专注于为我们的“Post”类型提取数据。虽然在技术上不是必需的，但文件是一种将相关内容逻辑组合在一起的好方法，所以我将 `Post` 类型和 `newPost` 移动到一个新的 `post.go` 文件中。

### Test helper

### 测试助手

We should take care of our tests too. We're going to be making assertions on `Posts` a lot, so we should write some code to help with that

我们也应该注意我们的测试。我们将在“帖子”上做很多断言，所以我们应该编写一些代码来帮助它

```go
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %+v, want %+v", got, want)
    }
}
```

```go
assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})
```

## Write the test first

## 先写测试

Let's extend our test further to extract the next line from the file, the description. Up until making it pass should now feel comfortable and familiar.

让我们进一步扩展我们的测试以从文件中提取下一行，即描述。直到通过它现在应该感到舒适和熟悉。

```go
func TestNewBlogPosts(t *testing.T) {
    const (
        firstBody = `Title: Post 1
Description: Description 1`
        secondBody = `Title: Post 2
Description: Description 2`
    )

    fs := fstest.MapFS{
        "hello world.md":  {Data: []byte(firstBody)},
        "hello-world2.md": {Data: []byte(secondBody)},
    }

    // rest of test code cut for brevity
    assertPost(t, posts[0], blogposts.Post{
        Title: "Post 1",
        Description: "Description 1",
    })

}
```

## Try to run the test

## 尝试运行测试

```
./blogpost_test.go:47:58: unknown field 'Description' in struct literal of type blogposts.Post
```

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Add the new field to `Post`.

将新字段添加到“发布”。

```go
type Post struct {
    Title       string
    Description string
}
```

The tests should now compile, and fail.

测试现在应该可以编译并失败。

```
=== RUN   TestNewBlogPosts
    blogpost_test.go:47: got {Title:Post 1
        Description: Description 1 Description:}, want {Title:Post 1 Description:Description 1}
```

## Write enough code to make it pass

## 编写足够的代码使其通过

The standard library has a handy library for helping you scan through data, line by line; [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner)

标准库有一个方便的库，可帮助您逐行扫描数据； [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner)

> Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text.

> Scanner 提供了一个方便的界面来读取数据，例如由换行符分隔的文本行文件。

```go
func newPost(postFile io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postFile)

    scanner.Scan()
    titleLine := scanner.Text()

    scanner.Scan()
    descriptionLine := scanner.Text()

    return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil
}
```

Handily, it also takes an `io.Reader` to read through (thank you again, loose-coupling), we don't need to change our function arguments.

方便地，它也需要一个 `io.Reader` 来阅读（再次感谢你，松耦合），我们不需要更改我们的函数参数。

Call `Scan` to read a line, and then extract the data using `Text`.

调用 Scan 读取一行，然后使用 Text 提取数据。

This function could never return an `error`. It would be tempting at this point to remove it from the return type, but we know we'll have to handle invalid file structures later so, we may as well leave it.

这个函数永远不会返回一个 `error`。在这一点上将它从返回类型中删除是很诱人的，但我们知道我们稍后必须处理无效的文件结构，所以我们不妨保留它。

## Refactor

## 重构

We have repetition around scanning a line and then reading the text. We know we're going to do this operation at least one more time, it's a simple refactor to DRY up so let's start with that.

我们围绕扫描一行然后阅读文本进行重复。我们知道我们将至少再做一次这个操作，这是一个简单的重构来干燥，所以让我们从它开始。

```go
func newPost(postFile io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postFile)

    readLine := func() string {
        scanner.Scan()
        return scanner.Text()
    }

    title := readLine()[7:]
    description := readLine()[13:]

    return Post{Title: title, Description: description}, nil
}
```

This has barely saved any lines of code, but that's rarely the point of refactoring. What I'm trying to do here is just separating the _what_ from the _how_ of reading lines to make the code a little more declarative to the reader.

这几乎没有节省任何代码行，但这很少是重构的重点。我在这里试图做的只是将阅读行的 _what_ 与 _how_ 分开，使代码对读者更具说明性。

Whilst the magic numbers of 7 and 13 get the job done, they're not awfully descriptive.

虽然 7 和 13 的神奇数字可以完成工作，但它们的描述性并不强。

```go
const (
    titleSeparator       = "Title: "
    descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postFile)

    readLine := func() string {
        scanner.Scan()
        return scanner.Text()
    }

    title := readLine()[len(titleSeparator):]
    description := readLine()[len(descriptionSeparator):]

    return Post{Title: title, Description: description}, nil
}
```

Now that I'm staring at the code with my creative refactoring mind, I'd like to try making our readLine function take care of removing the tag. There's also a more readable way of trimming a prefix from a string with the function `strings.TrimPrefix`.

现在，我正以创造性的重构思维盯着代码，我想尝试让我们的 readLine 函数负责删除标记。还有一种更具可读性的方法，可以使用函数“strings.TrimPrefix”从字符串中修剪前缀。

```go
func newPost(postBody io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postBody)

    readMetaLine := func(tagName string) string {
        scanner.Scan()
        return strings.TrimPrefix(scanner.Text(), tagName)
    }

    return Post{
        Title:       readMetaLine(titleSeparator),
        Description: readMetaLine(descriptionSeparator),
    }, nil
}
```

You may or may not like this idea, but I do. The point is in the refactoring state we are free to play with the internal details, and you can keep running your tests to check things still behave correctly. We can always go back to previous states if we're not happy. The TDD approach gives us this license to frequently experiment with ideas, so we have more shots at writing great code.

你可能喜欢也可能不喜欢这个主意，但我喜欢。关键是在重构状态，我们可以自由地处理内部细节，您可以继续运行测试以检查事情是否仍然正确运行。如果我们不开心，我们总是可以回到以前的状态。 TDD 方法为我们提供了频繁试验想法的许可，因此我们有更多机会编写出色的代码。

The next requirement is extracting the post's tags. If you're following along, I'd recommend trying to implement it yourself before reading on. You should now have a good, iterative rhythm and feel confident to extract the next line and parse out the data.

下一个要求是提取帖子的标签。如果您正在跟进，我建议您在继续阅读之前尝试自己实现它。您现在应该有一个良好的迭代节奏，并且有信心提取下一行并解析出数据。

For brevity, I will not go through the TDD steps, but here's the test with tags added.

为简洁起见，我不会完成 TDD 步骤，但这里是添加了标签的测试。

```go
func TestNewBlogPosts(t *testing.T) {
    const (
        firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
        secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
    )

    // rest of test code cut for brevity
    assertPost(t, posts[0], blogposts.Post{
        Title:       "Post 1",
        Description: "Description 1",
        Tags:        []string{"tdd", "go"},
    })
}
```

You're only cheating yourself if you just copy and paste what I write. To make sure we're all on the same page, here's my code which includes extracting the tags.

如果你只是复制和粘贴我写的东西，你只是在欺骗自己。为了确保我们都在同一页面上，这是我的代码，其中包括提取标签。

```go
const (
    titleSeparator       = "Title: "
    descriptionSeparator = "Description: "
    tagsSeparator        = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postBody)

    readMetaLine := func(tagName string) string {
        scanner.Scan()
        return strings.TrimPrefix(scanner.Text(), tagName)
    }

    return Post{
        Title:       readMetaLine(titleSeparator),
        Description: readMetaLine(descriptionSeparator),
        Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
    }, nil
}
```

Hopefully no surprises here. We were able to re-use `readMetaLine` to get the next line for the tags and then split them up using `strings.Split`.

希望这里没有惊喜。我们能够重新使用 `readMetaLine` 来获取标签的下一行，然后使用 `strings.Split` 将它们拆分。

The last iteration on our happy path is to extract the body.

我们快乐路径上的最后一次迭代是提取身体。

Here's a reminder of the proposed file format.

这是建议的文件格式的提醒。

```
Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd, go
---
Hello world!

The body of posts starts after the `---`
```

We've read the first 3 lines already. We then need to read one more line, discard it and then the remainder of the file contains the post's body.

我们已经阅读了前 3 行。然后我们需要再读一行，丢弃它，然后文件的其余部分包含帖子的正文。

## Write the test first

## 先写测试

Change the test data to have the separator, and a body with a few newlines to check we grab all the content.

将测试数据更改为具有分隔符和带有几个换行符的正文，以检查我们是否获取了所有内容。

```go
     const (
        firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
        secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
    )
```

Add to our assertion like the others

像其他人一样添加到我们的断言中

```go
     assertPost(t, posts[0], blogposts.Post{
        Title:       "Post 1",
        Description: "Description 1",
        Tags:        []string{"tdd", "go"},
        Body: `Hello
World`,
    })
```

## Try to run the test

## 尝试运行测试

```
./blogpost_test.go:60:3: unknown field 'Body' in struct literal of type blogposts.Post
```

As we'd expect.

正如我们所料。

## Write the minimal amount of code for the test to run and check the failing test output

## 为测试编写最少的代码以运行并检查失败的测试输出

Add `Body` to `Post` and the test should fail.

将 `Body` 添加到 `Post`，测试应该会失败。

```
=== RUN   TestNewBlogPosts
    blogposts_test.go:38: got {Title:Post 1 Description:Description 1 Tags:[tdd go] Body:}, want {Title:Post 1 Description:Description 1 Tags:[tdd go] Body:Hello
        World}
```

## Write enough code to make it pass

## 编写足够的代码使其通过

1. Scan the next line to ignore the `---` separator.
2. Keep scanning until there's nothing left to scan.

1. 扫描下一行忽略`---`分隔符。

2. 继续扫描，直到没有任何东西可以扫描。

```go
func newPost(postBody io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postBody)

    readMetaLine := func(tagName string) string {
        scanner.Scan()
        return strings.TrimPrefix(scanner.Text(), tagName)
    }

    title := readMetaLine(titleSeparator)
    description := readMetaLine(descriptionSeparator)
    tags := strings.Split(readMetaLine(tagsSeparator), ", ")

    scanner.Scan() // ignore a line

    buf := bytes.Buffer{}
    for scanner.Scan() {
        fmt.Fprintln(&buf, scanner.Text())
    }
    body := strings.TrimSuffix(buf.String(), "\n")

    return Post{
        Title:       title,
        Description: description,
        Tags:        tags,
        Body:        body,
    }, nil
}
```

- `scanner.Scan()` returns a `bool` which indicates whether there's more data to scan, so we can use that with a `for` loop to keep reading through the data until the end.
- After every `Scan()` we write the data into the buffer using `fmt.Fprintln`. We use the version that adds a newline because the scanner removes the newlines from each line, but we need to maintain them.
- Because of the above, we need to trim the final newline, so we don't have a trailing one.

- `scanner.Scan()` 返回一个 `bool`，指示是否有更多数据要扫描，因此我们可以使用它和 `for` 循环来继续读取数据直到结束。
- 在每次 `Scan()` 之后，我们使用 `fmt.Fprintln` 将数据写入缓冲区。我们使用添加换行符的版本，因为扫描器会从每一行中删除换行符，但我们需要维护它们。
- 由于上述原因，我们需要修剪最后的换行符，所以我们没有尾随。

## Refactor

## 重构

Encapsulating the idea of getting the rest of the data into a function will help future readers quickly understand _what_ is happening in `newPost`, without having to concern themselves with implementation specifics.

将获取其余数据的想法封装到一个函数中将有助于未来的读者快速了解“newPost”中发生的_什么_，而不必担心实现细节。

```go
func newPost(postBody io.Reader) (Post, error) {
    scanner := bufio.NewScanner(postBody)

    readMetaLine := func(tagName string) string {
        scanner.Scan()
        return strings.TrimPrefix(scanner.Text(), tagName)
    }

    return Post{
        Title:       readMetaLine(titleSeparator),
        Description: readMetaLine(descriptionSeparator),
        Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
        Body:        readBody(scanner),
    }, nil
}

func readBody(scanner *bufio.Scanner) string {
    scanner.Scan() // ignore a line
    buf := bytes.Buffer{}
    for scanner.Scan() {
        fmt.Fprintln(&buf, scanner.Text())
    }
    return strings.TrimSuffix(buf.String(), "\n")
}
```

## Iterating further

## 进一步迭代

We've made our "steel thread" of functionality, taking the shortest route to get to our happy path, but clearly there's some distance to go before it is production ready.

我们已经制作了功能的“钢线”，以最短的路线到达我们的幸福之路，但显然在生产准备好之前还有一段距离。

We haven't handled:

我们还没有处理：

- when the file's format is not correct
- the file is not a `.md`
- what if the order of the metadata fields is different? Should that be allowed? Should we be able to handle it?

- 当文件格式不正确时
- 该文件不是`.md`
- 如果元数据字段的顺序不同怎么办？应该允许吗？我们应该能够处理它吗？

Crucially though, we have working software, and we have defined our interface. The above are just further iterations, more tests to write and drive our behaviour. To support any of the above we shouldn't have to change our _design_, just implementation details.

但至关重要的是，我们有可以运行的软件，并且我们已经定义了我们的界面。以上只是进一步的迭代，更多的测试来编写和驱动我们的行为。为了支持上述任何一项，我们不必更改我们的 _design_，只需更改实现细节。

Keeping focused on the goal means we made the important decisions, and validated them against the desired behaviour, rather than getting bogged down on matters that won't affect the overall design.

专注于目标意味着我们做出了重要的决定，并根据所需的行为验证了它们，而不是陷入不会影响整体设计的事情上。

## Wrapping up

##  总结

`fs.FS`, and the other changes in Go 1.16 give us some elegant ways of reading data from file systems and testing them simply.

`fs.FS` 和 Go 1.16 中的其他更改为我们提供了一些从文件系统读取数据并简单地测试它们的优雅方法。

If you wish to try out the code "for real":

如果您想“真正”试用代码：

- Create a `cmd` folder within the project, add a `main.go` file
- Add the following code

- 在项目中创建一个 `cmd` 文件夹，添加一个 `main.go` 文件
- 添加以下代码

```go
import (
    blogposts "github.com/quii/fstest-spike"
    "log"
    "os"
)

func main() {
    posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
    if err != nil {
        log.Fatal(err)
    }
    log.Println(posts)
}
```

- Add some markdown files into a `posts` folder and run the program!

- 将一些 markdown 文件添加到 `posts` 文件夹中并运行程序！

Notice the symmetry between the production code

注意生产代码之间的对称性

```go
posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
```

And the tests

和测试

```go
posts, err := blogposts.NewPostsFromFS(fs)
```

This is when consumer-driven, top-down TDD _feels correct_.

这是消费者驱动的、自上而下的 TDD _感觉正确的时候。

A user of our package can look at our tests and quickly get up to speed with what it's supposed to do and how to use it. As maintainers, we can be _confident our tests are useful because they're from a consumer's point of view_. We're not testing implementation details or other incidental details, so we can be reasonably confident that our tests will help us, rather than hinder us when refactoring.

我们包的用户可以查看我们的测试并快速了解它应该做什么以及如何使用它。作为维护者，我们可以_确信我们的测试是有用的，因为它们是从消费者的角度来看的_。我们不测试实现细节或其他附带细节，因此我们可以合理地相信我们的测试将帮助我们，而不是在重构时阻碍我们。

By relying on good software engineering practices like  [**dependency injection**](dependency-injection.md) our code is simple to test and re-use.

通过依赖于良好的软件工程实践，例如 [**dependency injection**](dependency-injection.md)，我们的代码易于测试和重用。

When you're creating packages, even if they're only internal to your project, prefer a top-down consumer driven approach. This will stop you over-imagining designs and making abstractions you may not even need and will help ensure the tests you write are useful.

当您创建包时，即使它们仅在您的项目内部，也更喜欢自上而下的消费者驱动方法。这将阻止您过度想象设计并进行您甚至可能不需要的抽象，并有助于确保您编写的测试有用。

The iterative approach kept every step small, and the continuous feedback helped us uncover unclear requirements possibly sooner than with other, more ad-hoc approaches.

迭代方法使每一步都保持在较小的范围内，并且持续的反馈帮助我们可能比其他更临时的方法更快地发现不明确的需求。

### Writing? 

###  写作？

It's important to note that these new features only have operations for _reading_ files. If your work needs to do writing, you'll need to look elsewhere. Remember to keep thinking about what the standard library offers currently, if you're writing data you should probably look into leveraging existing interfaces such as `io.Writer` to keep your code loosely-coupled and re-usable.

需要注意的是，这些新功能只对 _reading_ 文件有操作。如果你的工作需要写作，你需要去别处寻找。请记住继续考虑标准库当前提供的内容，如果您正在编写数据，您可能应该考虑利用现有的接口，例如 `io.Writer` 来保持您的代码松散耦合和可重用。

### Further reading

### 进一步阅读

- This was a light intro to `io/fs`. [Ben Congdon has done an excellent write-up](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/) which was a lot of help for writing this chapter.
- [Discussion on the file system interfaces](https://github.com/golang/go/issues/41190) 

- 这是对 `io/fs` 的简单介绍。 [Ben Congdon 写了一篇很棒的文章](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/)写这一章。
- [文件系统接口讨论](https://github.com/golang/go/issues/41190)

