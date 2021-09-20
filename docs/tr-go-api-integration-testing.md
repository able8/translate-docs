# Web API Integration Testing with Go

  # 使用 Go 进行 Web API 集成测试

 Tuesday. June 11, 2019 -  31 mins

周二。 2019 年 6 月 11 日 - 31 分钟

I’m learning Go by building a small API-backed web application, and wanted  to share the process in case it helps someone else. In this post, we'll  continue where we left off last time with the [Go web API](https://rshipp.com/go-web-api) for managing [GitHub stars](https://help.github.com/en/articles/about-stars), adding automated tests to ensure our code functions as expected. If  you'd like to follow along with this post without going through the  previous one, you can grab a copy of the API (`main.go`) from [this GitHub repo](https://github.com/rshipp/StarManager/tree/0d7cdd291711c7dd6a706bc38844f250343c7b1f).

我正在通过构建一个由 API 支持的小型 Web 应用程序来学习 Go，并希望分享该过程以帮助其他人。在这篇文章中，我们将继续上次中断的地方，使用 [Go web API](https://rshipp.com/go-web-api) 来管理 [GitHub 星](https://help.github.com/en/articles/about-stars)，添加自动化测试以确保我们的代码按预期运行。如果您想在不阅读前一篇文章的情况下继续阅读这篇文章，您可以从 [this GitHub repo](https://github.com/rshipp/) 获取 API (`main.go`) 的副本StarManager/tree/0d7cdd291711c7dd6a706bc38844f250343c7b1f)。

We'll start by writing [integration tests](https://en.wikipedia.org/wiki/Integration_testing), which will rely on a real database backend rather than a [mock](https://en.wikipedia.org/wiki/Mock_object) or [stub](https://en.wikipedia.org/wiki/Unit_testing). Compared to [unit tests](https://en.wikipedia.org/wiki/Unit_testing), integration tests have a few drawbacks and benefits for our purposes.

我们将首先编写 [集成测试](https://en.wikipedia.org/wiki/Integration_testing)，它将依赖于真实的数据库后端而不是 [mock](https://en.wikipedia.org/wiki/Mock_object) 或 [存根](https://en.wikipedia.org/wiki/Unit_testing)。与[单元测试](https://en.wikipedia.org/wiki/Unit_testing) 相比，集成测试对于我们的目的来说有一些缺点和好处。

Drawbacks:

缺点：

1. Using a real database makes integration tests slower than unit tests that rely on stubbed methods or mocked interfaces.
2. Testing the full stack at once instead of each small piece at a time can make it harder to tell where bugs are coming from.

1. 使用真实数据库使集成测试比依赖存根方法或模拟接口的单元测试更慢。
2. 一次测试整个堆栈而不是一次测试每个小部分会使判断错误来自哪里变得更加困难。

Benefits:

好处：

1. Integration tests allow us to ensure our SQL queries and schema are correct  (especially important here because that’s where most of our API’s  functionality comes from).
2. Since we don’t have to use stubs or mocks, we can write less code for the tests and get them up and running faster.

1. 集成测试让我们能够确保我们的 SQL 查询和模式是正确的（在这里尤其重要，因为这是我们大部分 API 功能的来源）。
2. 由于我们不必使用存根或模拟，我们可以为测试编写更少的代码并使它们更快地启动和运行。

We will eventually want unit tests, for the reasons mentioned above, but  starting with integration gives us the biggest return for the time  being.

由于上述原因，我们最终会需要单元测试，但从集成开始，我们暂时获得了最大的回报。

## What You’ll Need

## 你需要什么

Before we get started, you’ll need a few things:

在我们开始之前，您需要准备一些东西：

- [Go](https://en.wikipedia.org/wiki/Code_smell) installed on your computer.
- The [Go web API](https://github.com/rshipp/StarManager/tree/0d7cdd291711c7dd6a706bc38844f250343c7b1f)(`main.go`) from the [previous post](https://rshipp.com/go-web-api).
- A text editor.

- [Go](https://en.wikipedia.org/wiki/Code_smell) 安装在您的计算机上。
- 来自 [上一篇文章](https://rshipp.com/go-web-api)。
- 一个文本编辑器。

You may also want to skim through the [tour of Go](https://tour.golang.org/welcome/1) or another introduction to the Go language if you haven't already, though you shouldn't need more than a basic understanding.

您可能还想浏览 [tour of Go](https://tour.golang.org/welcome/1) 或其他 Go 语言介绍（如果您还没有)，但您不应该需要更多一个基本的了解。

If you run into anything unclear in this post, feel free to [open an issue](https://github.com/rshipp/rshipp.github.io/issues) on GitHub and let me know!

如果您在本文中遇到任何不清楚的地方，请随时在 GitHub 上 [open an issue](https://github.com/rshipp/rshipp.github.io/issues) 告诉我！

## Create Handler

## 创建处理程序

Let’s go through and write an integration test function for each of our five HTTP request handlers from `main.go`, starting with the Create handler.

让我们为 main.go 中的五个 HTTP 请求处理程序中的每一个编写集成测试函数，从 Create 处理程序开始。

Go tests are expected to be in a file ending with `_test.go`. Since we’re writing tests for `main.go`, the conventional test file name is `main_test.go`. (Make sure `main_test.go` and `main.go` are in the same folder.)

Go 测试应该在以 `_test.go` 结尾的文件中。由于我们正在为 `main.go` 编写测试，所以常规的测试文件名为 `main_test.go`。 （确保 `main_test.go` 和 `main.go` 在同一个文件夹中。）

Set up a basic outline for our first test in `main_test.go`:

在 main_test.go 中为我们的第一个测试设置一个基本大纲：

```gogo
package main

import (
    "testing"
)

func setup() *App {
    // Initialize an in-memory database for full integration testing.
    app := &App{}
    app.Initialize("sqlite3", ":memory:")
    return app
}

func teardown(app *App) {
    // Closing the connection discards the in-memory database.
    app.DB.Close()
}

func TestCreateHandler(t *testing.T) {
    app := setup()

    // Test body will be here!

    teardown(app)
}
```

The documentation for the Go [testing](https://golang.org/pkg/testing/#pkg-overview) package goes over the requirements for test functions: they must each be named like `TestXxx` (where `X` is a capital letter), and have an argument `t *testing.T`. 

Go [testing](https://golang.org/pkg/testing/#pkg-overview) 包的文档详细介绍了测试函数的要求：它们都必须命名为 `TestXxx`（其中 `X` 是一个大写字母)，并有一个参数`t *testing.T`。

There are no special “setup” or “teardown” functions as we might be used to  from other languages, nor is there a testing class with instance  variables we can use to access our database. To get around this, we  define a `setup` function to initialize an [in-memory SQLite database](https://sqlite.org/inmemorydb.html) and return an `App` pointer (we defined `App` in ` main.go`), and `teardown` to accept that same pointer and close our database connection. We then call `app := setup()` at the start of every integration test function, and `teardown(app)` at the end of each function. This ensures that our database is always  clean and in a consistent state at the start of each test.

没有我们可能在其他语言中习惯使用的特殊“设置”或“拆卸”功能，也没有我们可以用来访问数据库的带有实例变量的测试类。为了解决这个问题，我们定义了一个 `setup` 函数来初始化一个 [in-memory SQLite 数据库](https://sqlite.org/inmemorydb.html) 并返回一个 `App` 指针（我们在 ` main.go`) 和 `teardown` 接受相同的指针并关闭我们的数据库连接。然后我们在每个集成测试函数开始时调用`app := setup()`，在每个函数结束时调用`teardown(app)`。这可确保我们的数据库在每次测试开始时始终干净且处于一致状态。

If we run the tests with `go test`, they should pass:

如果我们使用 go test 运行测试，它们应该通过：

```go
PASS
ok      _/<...>/StarManager     0.004s
```

Let’s start filling out the test body for `TestCreateHandler`:

让我们开始填写 TestCreateHandler 的测试主体：

```go
     testStar := &Star{
        ID:          1,
        Name:        "test/name",
        Description: "test desc",
        URL:         "test url",
    }

    // Transform Star record into *strings.Reader suitable for use in HTTP POST forms.
    data := url.Values{
        "name":        {testStar.Name},
        "description": {testStar.Description},
        "url":         {testStar.URL},
    }

    form := strings.NewReader(data.Encode())

    // Set up a new request.
    req, err := http.NewRequest("POST", "/stars", form)
    if err != nil {
        t.Fatal(err)
    }
    // Our API expects a form body, so set the content-type header to make sure it's treated as one.
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    http.HandlerFunc(app.CreateHandler).ServeHTTP(rr, req)
```

We’re using a few new Go packages here, so we’ll have to add them to the import list at the top of `main_test.go`:

我们在这里使用了一些新的 Go 包，所以我们必须将它们添加到 `main_test.go` 顶部的导入列表中：

```go
import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"
    "testing"
)
```

The [url.Values.Encode()](https://golang.org/pkg/net/url/#Values) and [strings.NewReader](https://golang.org/pkg/strings/#NewReader) code takes our `testStar` record and converts it into the string format that [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) expects. We then use that `form` string with `http.NewRequest` to set up a request to the Create endpoint defined in `main.go`, `POST /stars`. Note that we are not actually making an HTTP request here, just preparing one.

[url.Values.Encode()](https://golang.org/pkg/net/url/#Values) 和 [strings.NewReader](https://golang.org/pkg/strings/#NewReader)代码获取我们的 `testStar` 记录并将其转换为 [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) 期望的字符串格式。然后我们使用带有 `http.NewRequest` 的 `form` 字符串来设置对 `main.go`、`POST /stars` 中定义的 Create 端点的请求。请注意，我们实际上并不是在这里发出 HTTP 请求，只是准备一个。

This is where Go's built-in `httptest` package comes in handy: we set up a [httptest.ResponseRecorder](https://golang.org/pkg/net/http/httptest/#ResponseRecorder), then pass it in to `http.HandlerFunc().ServeHTTP()`. With both this response record and the request we prepared earlier, we can test our `app.CreateHandler()` directly, without needing to set up a local HTTP server or client. In  essence, we’re passing variables around in Go’s internal functions  without using network requests or responses at all.

这就是 Go 内置的 `httptest` 包派上用场的地方：我们设置了一个 [httptest.ResponseRecorder](https://golang.org/pkg/net/http/httptest/#ResponseRecorder)，然后将它传递给`http.HandlerFunc().ServeHTTP()`。有了这个响应记录和我们之前准备的请求，我们可以直接测试我们的`app.CreateHandler()`，不需要设置本地HTTP服务器或客户端。本质上，我们在 Go 的内部函数中传递变量，根本不使用网络请求或响应。

If there was an unexpected error forming the request, we can call [t.Fatal](https://golang.org/pkg/testing/#T.Fatal) to stop executing the test function immediately and print the error message.

如果形成请求时出现意外错误，我们可以调用 [t.Fatal](https://golang.org/pkg/testing/#T.Fatal) 立即停止执行测试函数并打印错误消息。

With our POST request “sent” to our handler, and the response recorded in `rr`, we can continue filling out `TestCreateHandler`, checking that our API works as expected:

随着我们的 POST 请求“发送”到我们的处理程序，并且响应记录在 `rr` 中，我们可以继续填写 `TestCreateHandler`，检查我们的 API 是否按预期工作：

```go
     // Test that the status code is correct.
    if status := rr.Code;status != http.StatusCreated {
        t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusCreated, status)
    }

    // Test that the Location header is correct.
    expectedURL := fmt.Sprintf("/stars/%s", testStar.Name)
    if location := rr.Header().Get("Location");location != expectedURL {
        t.Errorf("Location header is invalid. Expected %s. Got %s instead", expectedURL, location)
    }

    // Test that the created star is correct.
    // Note: There is only one star in the database.
    createdStar := Star{}
    app.DB.First(&createdStar)
    if createdStar != *testStar {
        t.Errorf("Created star is invalid. Expected %+v. Got %+v instead", testStar, createdStar)
    }
```

First up, we expect to see `201 Created` as the status code. Go provides some nicely named [aliases](https://golang.org/pkg/net/http/#pkg-constants) for HTTP status codes, so we can reference this as `http.StatusCreated`, and compare it to the actual response code we got in `rr.Code`. If they’re different, we use [t.Errorf](https://golang.org/pkg/testing/#T.Errorf) to print out a useful message and then fail the test.

首先，我们希望看到“201 Created”作为状态代码。 Go 为 HTTP 状态代码提供了一些很好命名的 [别名](https://golang.org/pkg/net/http/#pkg-constants)，因此我们可以将其引用为 `http.StatusCreated`，并将其与我们在 `rr.Code` 中得到的实际响应代码。如果它们不同，我们使用 [t.Errorf](https://golang.org/pkg/testing/#T.Errorf) 打印出有用的消息，然后测试失败。

Next, the `Location` header: we expect this to be set to the URL of the star that was just created, which we defined in `main.go` as `/stars/{star.Name}`. The actual header is in `rr.Header()`, so we can compare that to the expected URL to verify correctness.

接下来是 `Location` 标头：我们希望将其设置为刚刚创建的星星的 URL，我们在 `main.go` 中将其定义为 `/stars/{star.Name}`。实际的标头在 `rr.Header()` 中，因此我们可以将其与预期的 URL 进行比较以验证正确性。

Finally, we want to see if a star was actually created. Since we're working with a real database, we can use our [GORM](http://gorm.io/) database pointer `app.DB` directly to fetch the first (and only) star in the database, and compare it to our original `testStar`. (Using GORM directly like this is a bit of a [code smell](https://en.wikipedia.org/wiki/Code_smell), but we’ll worry about refactoring later.)

最后，我们想看看星星是否真的被创造出来了。由于我们正在使用一个真实的数据库，我们可以直接使用我们的 [GORM](http://gorm.io/) 数据库指针 `app.DB` 来获取数据库中的第一个（也是唯一一个）星，并进行比较它到我们原来的`testStar`。 （像这样直接使用 GORM 有点[代码味道](https://en.wikipedia.org/wiki/Code_smell)，但我们稍后会担心重构。)

When we run `go test` again, it should report a `PASS`. Looks like our Create handler passed the test!

当我们再次运行 `go test` 时，它应该报告一个 `PASS`。看起来我们的 Create 处理程序通过了测试！

For a more complete project, we’d want to have additional tests for edge  cases: what happens if we try to create a star that already exists? what if the database is down? are there invalid characters that break the  SQL query? Our API in `main.go` is  pretty naive right now, so a lot of these will probably fail in  unexpected ways. As we build more functionality into the API, we’ll  continue to add integration and unit tests that make sure everything  works correctly.

对于一个更完整的项目，我们希望对边缘情况进行额外的测试：如果我们尝试创建一个已经存在的星星会发生什么？如果数据库宕机了怎么办？是否存在破坏 SQL 查询的无效字符？我们在 `main.go` 中的 API 现在非常幼稚，所以其中很多可能会以意想不到的方式失败。随着我们在 API 中构建更多功能，我们将继续添加集成和单元测试，以确保一切正常工作。

## Update Handler

## 更新处理程序

The  Update handler, like the Create handler, expects a form-encoded request  body with star attributes. The code we're using to do that is a little  fragile (we may have to manually update it when we add new fields to the Star struct), so let's start by refactoring it out into a function  inside `main_test.go` so we only have to maintain it in once place:

Update 处理程序与 Create 处理程序一样，需要一个带有星号属性的表单编码请求正文。我们用来做这件事的代码有点脆弱（当我们向 Star 结构添加新字段时，我们可能必须手动更新它），所以让我们首先将它重构为 `main_test.go` 中的一个函数，这样我们只需要在一个地方维护它：

```go
func StarFormValues(star Star) *strings.Reader {
    // Transforms Star record into *strings.Reader suitable for use in HTTP POST forms.
    data := url.Values{
        "name":        {star.Name},
        "description": {star.Description},
        "url":         {star.URL},
    }

    return strings.NewReader(data.Encode())
}
```

In the Create handler, remove the code we pulled out into `StarFormValues`, and update the `NewRequest` call:

在 Create 处理程序中，删除我们提取到 `StarFormValues` 中的代码，并更新 `NewRequest` 调用：

```go
     // Set up a new request.
    req, err := http.NewRequest("POST", "/stars", StarFormValues(*testStar))
```

Now we can reuse that function in the Update handler too.

现在我们也可以在更新处理程序中重用该函数。

Since we want to test updating a star, let’s start by putting one in the database:

由于我们要测试更新一颗星，让我们先在数据库中放入一颗星：

```go
func TestUpdateHandler(t *testing.T) {
    app := setup()

    // Create a star for us to update.
    testStar := &Star{
        ID:          1,
        Name:        "test/name",
        Description: "test desc",
        URL:         "test url",
    }
    app.DB.Create(testStar)
```

There are two main things we want to test here: updating a star’s name (which changes the URL used to  reference it), and updating other fields in a star. We could do this  manually, but that sounds like a lot of duplicated code. Luckily, Go has a pattern called [table-driven tests](https://github.com/golang/go/wiki/TableDrivenTests) that will save us a lot of effort.

我们要在这里测试两件主要的事情：更新星号的名称（更改用于引用它的 URL），以及更新星号中的其他字段。我们可以手动执行此操作，但这听起来像是很多重复的代码。幸运的是，Go 有一个名为 [table-driven tests](https://github.com/golang/go/wiki/TableDrivenTests) 的模式，可以为我们节省很多精力。

The basic pattern for a table-driven test looks something like this:

表驱动测试的基本模式如下所示：

```go
myTests = []struct {
    input    int
    expected int
}{
    {1, 1},
    {2, 4},
    {4, 16},
}

for _, tt := range myTests {
    if actual := MySquareFunction(tt.input);actual != tt.expected {
        t.Errorf("MySquareFunction(%d): expected: %d, actual: %d", tt.input, tt.expected, actual)
    }
}
```

The “table” is a [slice](https://gobyexample.com/slices) of anonymous (unnamed) [structs](https://gobyexample.com/structs). We can define as many fields as we need - in this case `input` and `expected` - and fill the slice with as many records as we want with different  values for those fields. Here, we have 3 records, each representing `{input, expected}`. We then loop over the table with `range` (`tt` is the conventional name for table-driven test elements, but you could  call it something else if you wanted), and run the same test on each  record in the table.

“表”是匿名（未命名）[结构](https://gobyexample.com/structs)的[切片](https://gobyexample.com/slices)。我们可以根据需要定义任意数量的字段——在本例中为“input”和“expected”——并用我们想要的尽可能多的记录填充切片，这些字段的值不同。在这里，我们有 3 条记录，每条记录代表 `{input, expected}`。然后我们用 `range` 遍历表（`tt` 是表驱动测试元素的常规名称，但如果您愿意，可以将其称为其他名称)，并对表中的每条记录运行相同的测试。

Applying this to our use case, back in `TestUpdateHandler` in `main_test.go`, we can set up a table of stars:

将此应用到我们的用例中，回到`main_test.go`中的`TestUpdateHandler`，我们可以设置一个星表：

```go
     // Set up a test table.
    starTests := []struct {
        original Star
        update   Star
    }{
        {original: *testStar,
            update: Star{ID: 1, Name: "test/name", Description: "updated desc", URL: "test URL"},
        },
        {original: Star{ID: 1, Name: "test/name", Description: "updated desc", URL: "test URL"},
            update: Star{ID: 1, Name: "updated name", Description: "updated desc", URL: "test URL"},
        },
    }

    for _, tt := range starTests {
```

The “original” star is what we know will be in the database when the test runs (note the second `original` is the same as the first `update`), and the “update” star is what we want to update it to .

“原始”星是我们知道在测试运行时将在数据库中的内容（注意第二个“原始”与第一个“更新”相同），“更新”星是我们想要将其更新为的.

Inside that loop, we send a PUT request to the update endpoint `/stars/{star.Name}`, with the contents of the updated fields:

在该循环中，我们向更新端点 `/stars/{star.Name}` 发送 PUT 请求，其中包含更新字段的内容：

```go
         // Set up a new request.
        req, err := http.NewRequest("PUT", fmt.Sprintf("/stars/%s", tt.original.Name), StarFormValues(tt.update))
        if err != nil {
            t.Fatal(err)
        }
        // Our API expects a form body, so set the content-type header appropriately.
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        rr := httptest.NewRecorder()
        // We need a mux router in order to pass in the `name` variable.
        r := mux.NewRouter()

        r.HandleFunc("/stars/{name:.*}", app.UpdateHandler).Methods("PUT")
        r.ServeHTTP(rr, req)
```

One difference here from the Create test: we need a custom router, since the Update handler expects a `name` variable with the name of the star we want to update. We use the same `{name:.*}` pattern here as we do in the routes at the bottom of `main.go`.

这里与 Create 测试的一个不同之处是：我们需要一个自定义路由器，因为 Update 处理程序需要一个 `name` 变量，其中包含我们要更新的星星的名称。我们在这里使用相同的 `{name:.*}` 模式，就像我们在 `main.go` 底部的路由中所做的那样。

Be sure to add mux (`"github.com/gorilla/mux"`) to the import list at the top of `main_test.go`, since we’re calling `mux.NewRouter()`.

确保将 mux (`"github.com/gorilla/mux"`) 添加到 `main_test.go` 顶部的导入列表中，因为我们正在调用 `mux.NewRouter()`。

The rest of the test function is about the same as it was for create; we check the return code (`204 No Content`) and make sure the database was updated successfully:

测试函数的其余部分与用于创建的大致相同；我们检查返回码（`204 No Content`）并确保数据库更新成功：

```go
         // Test that the status code is correct.
        if status := rr.Code;status != http.StatusNoContent {
            t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusNoContent, status)
        }

        // Test that the updated star is correct.
        // Note: There is only one star in the database.
        updatedStar := Star{}
        app.DB.First(&updatedStar)
        if updatedStar != tt.update {
            t.Errorf("Updated star is invalid. Expected %+v. Got %+v instead", tt.update, updatedStar)
        }
    }

    teardown(app)
}
```

## View Handler

## 视图处理程序

In the View handler test, we’ll need a couple new techniques: reading the HTTP response body, and [unmarshalling](https://en.wikipedia.org/wiki/Unmarshalling) JSON. Let’s go ahead and add the imports we’ll use to the list at the top of `main_test.go`:

在视图处理程序测试中，我们需要一些新技术：读取 HTTP 响应正文和 [unmarshalling](https://en.wikipedia.org/wiki/Unmarshalling) JSON。让我们继续将我们将使用的导入添加到 `main_test.go` 顶部的列表中：

```go
     "encoding/json"
    "io/ioutil"
```

For the test function, we’ll use a loop  that’s similar to our table-driven tests. We don’t really need the  anonymous struct though, so we can simplify it a bit to just a slice of  Star records:

对于测试功能，我们将使用类似于我们的表驱动测试的循环。不过，我们并不真正需要匿名结构，因此我们可以将其简化为仅一部分 Star 记录：

```go
func TestViewHandler(t *testing.T) {
    app := setup()

    // Set up a test table.
    starTests := []Star{
        Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"},
        Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"},
    }

    for _, star := range starTests {
        // Create a star for us to view.
        app.DB.Create(star)

        // Set up a new request.
        req, err := http.NewRequest("GET", fmt.Sprintf("/stars/%s", star.Name), nil)
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        // We need a mux router in order to pass in the `name` variable.
        r := mux.NewRouter()

        r.HandleFunc("/stars/{name:.*}", app.ViewHandler).Methods("GET")
        r.ServeHTTP(rr, req)

        // Test that the status code is correct.
        if status := rr.Code;status != http.StatusOK {
            t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
        }
```

We haven’t needed it before now, but `rr` does hold the complete response body returned from our API. We can access it through [rr.Result().Body](https://golang.org/pkg/net/http/httptest/#ResponseRecorder.Result), by [reading](https://stackoverflow.com/questions/39945968/most-efficient-way-to-convert-io-readcloser-to-byte-array) with [ioutil.ReadAll()](https://golang.org/pkg/io/ioutil/#ReadAll) :

我们之前不需要它，但 `rr` 确实保存了从我们的 API 返回的完整响应主体。我们可以通过[rr.Result().Body](https://golang.org/pkg/net/http/httptest/#ResponseRecorder.Result)，通过[阅读](https://stackoverflow.com/问题/39945968/most-efficient-way-to-convert-io-readcloser-to-byte-array) 与 [ioutil.ReadAll()](https://golang.org/pkg/io/ioutil/#ReadAll) ：

```go
         // Read the response body.
        data, err := ioutil.ReadAll(rr.Result().Body)
        if err != nil {
            t.Fatal(err)
        }
```

Now that we have the response content in `data`, we can [unmarshal it](https://golang.org/pkg/encoding/json/#Unmarshal) into a Star struct and compare with the star we created in the database :

现在我们在`data`中有响应内容，我们可以[解组它](https://golang.org/pkg/encoding/json/#Unmarshal)到一个Star结构中，并与我们在数据库中创建的star进行比较：

```go
         // Test that the updated star is correct.
        returnedStar := Star{}
        if err := json.Unmarshal(data, &returnedStar);err != nil {
            t.Errorf("Returned star is invalid JSON. Got: %s", data)
        }
        if returnedStar != star {
            t.Errorf("Returned star is invalid. Expected %+v. Got %+v instead", star, returnedStar)
        }
    }

    teardown(app)
}
```

## List Handler

## 列表处理程序

The List handler test is pretty similar to the View test, in that we make a GET request and check the JSON from the response body. Here, we don’t  need a custom `mux` router, since there aren’t any variables to pass in to the List handler.

List 处理程序测试与 View 测试非常相似，因为我们发出 GET 请求并检查响应正文中的 JSON。在这里，我们不需要自定义的 `mux` 路由器，因为没有任何变量可以传递给 List 处理程序。

```go
func TestListHandler(t *testing.T) {
    app := setup()

    // Create a couple stars to list.
    stars := []Star{
        Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"},
        Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"},
    }

    for _, star := range stars {
        app.DB.Create(star)
    }

    // Set up a new request.
    req, err := http.NewRequest("GET", "/stars", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()

    http.HandlerFunc(app.ListHandler).ServeHTTP(rr, req)

    // Test that the status code is correct.
    if status := rr.Code;status != http.StatusOK {
        t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
    }

    // Read the response body.
    data, err := ioutil.ReadAll(rr.Result().Body)
    if err != nil {
        t.Fatal(err)
    }

    // Test that our stars list is the same as what was returned.
    returnedStars := []Star{}
    if err := json.Unmarshal(data, &returnedStars);err != nil {
        t.Errorf("Returned star list is invalid JSON. Got: %s", data)
    }
    if len(returnedStars) != len(stars) {
        t.Errorf("Returned star list is an invalid length. Expected %d. Got %d instead", len(stars), len(returnedStars))
    }
    for index, returnedStar := range returnedStars {
        if returnedStar != stars[index] {
            t.Errorf("Returned star is invalid. Expected %+v. Got %+v instead", stars[index], returnedStar)
        }
    }

    teardown(app)
}
```

Note the loop at the bottom of this  test, where we make sure each item in the JSON matches each of our test  stars, in order. If the order returned from the List handler ever  changes, we’ll have to revisit this test.

请注意此测试底部的循环，我们确保 JSON 中的每个项目按顺序与我们的每个测试星匹配。如果从 List 处理程序返回的顺序发生变化，我们将不得不重新审视这个测试。

## Delete Handler

## 删除处理程序

The test for the Delete handler doesn’t really have anything new either,  just recycling the same concepts used above in a slightly different way:

对 Delete 处理程序的测试也没有任何新内容，只是以稍微不同的方式回收了上面使用的相同概念：

```go
func TestDeleteHandler(t *testing.T) {
    app := setup()

    // Set up a test table.
    starTests := []struct {
        star Star
    }{
        {star: Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"}},
        {star: Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"}},
    }

    for _, tt := range starTests {
        // Create a star for us to delete.
        app.DB.Create(tt.star)

        // Set up a new request.
        req, err := http.NewRequest("DELETE", fmt.Sprintf("/stars/%s", tt.star.Name), nil)
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        // We need a mux router in order to pass in the `name` variable.
        r := mux.NewRouter()

        r.HandleFunc("/stars/{name:.*}", app.DeleteHandler).Methods("DELETE")
        r.ServeHTTP(rr, req)

        // Test that the status code is correct.
        if status := rr.Code;status != http.StatusNoContent {
            t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusNoContent, status)
        }

        // Test that the star is no longer in the db.
        deletedStar := Star{}
        app.DB.Where("name = ?", tt.star.Name).First(&deletedStar)
        if deletedStar != (Star{}) {
            t.Errorf("Star still exists in db: %+v", tt.star)
        }
    }

    teardown(app)
}
```

Run the tests with `go test`, and watch them all pass!

使用 `go test` 运行测试，并观察它们全部通过！

```go
PASS
ok      _/<...>/StarManager     0.008s
```

If you want to see the results of each individual test, use the “verbose” flag - `go test -v`:

如果您想查看每个单独测试的结果，请使用“详细”标志 - `go test -v`：

```go
=== RUN   TestCreateHandler
--- PASS: TestCreateHandler (0.00s)
=== RUN   TestUpdateHandler
--- PASS: TestUpdateHandler (0.00s)
=== RUN   TestViewHandler
--- PASS: TestViewHandler (0.00s)
=== RUN   TestListHandler
--- PASS: TestListHandler (0.00s)
=== RUN   TestDeleteHandler
--- PASS: TestDeleteHandler (0.00s)
PASS
ok      _/home/ryan/dev/StarManager     0.008s
```

You can check out the complete code for this post [on GitHub](https://github.com/rshipp/StarManager/tree/38bb56732e89dc5c70d486cf71ce0dfa9ee88d2c).

你可以查看这篇文章的完整代码 [在 GitHub 上](https://github.com/rshipp/StarManager/tree/38bb56732e89dc5c70d486cf71ce0dfa9ee88d2c)。

## Conclusion

##  结论

A quick recap:

快速回顾：

1. We wrote setup/teardown functions for our integration tests.
2. Tested HTTP response codes for all five HTTP request handlers.
3. Tested HTTP response headers, and database contents for the Create handler.
4. Tested database contents for the Update and Delete handlers.
5. Tested HTTP response body JSON for the View and List handlers.

1. 我们为集成测试编写了设置/拆卸功能。
2. 测试了所有五个 HTTP 请求处理程序的 HTTP 响应代码。
3. 测试 HTTP 响应标头和 Create 处理程序的数据库内容。
4. 测试更新和删除处理程序的数据库内容。
5. 测试了视图和列表处理程序的 HTTP 响应正文 JSON。

In the process, we covered several Go features and concepts:

在这个过程中，我们介绍了几个 Go 特性和概念：

- Writing Go test files and functions.
- Integration tests vs unit tests.
- Basic use of the `testing` and `httptest` packages.
- Table-driven tests.
- Anonymous structs.
- JSON unmarshalling.
- Reading a byte stream with `ioutil.ReadAll()`.
- Using custom routes in HTTP handler tests to pass in variables.
- And more!

- 编写 Go 测试文件和函数。
- 集成测试与单元测试。
- `testing` 和 `httptest` 包的基本使用。
- 表驱动测试。
- 匿名结构。
- JSON 解组。
- 使用 `ioutil.ReadAll()` 读取字节流。
- 在 HTTP 处理程序测试中使用自定义路由来传递变量。
- 和更多！

In future posts, I’ll revisit this API and walk through adding some new functionality and creating a frontend for the star app.

在以后的文章中，我将重新访问此 API，并逐步介绍添加一些新功能并为 Star 应用程序创建前端。

## Additional References

## 其他参考资料

I found these three resources especially helpful! If you’re new to  testing in Go and want to learn more, I highly recommend them as a  starting place.

我发现这三个资源特别有用！如果您不熟悉 Go 测试并想了解更多信息，我强烈建议您将其作为起点。

- [Integration Tests in Go - Philosophical Hacker](https://www.philosophicalhacker.com/post/integration-tests-in-go/)
- [A Quick Guide to Testing in Golang - CaitieM](https://caitiem.com/2016/08/18/a-quick-guide-to-testing-in-golang/)
- [Testing HTTP handlers in Go - Lanre Adelowo](https://lanre.wtf/blog/2017/04/08/testing-http-handlers-go/)

- [Go 中的集成测试 - Philosophical Hacker](https://www.philosophicalhacker.com/post/integration-tests-in-go/)
- [Golang 测试快速指南 - CaitieM](https://caitiem.com/2016/08/18/a-quick-guide-to-testing-in-golang/)
- [在 Go 中测试 HTTP 处理程序 - Lanre Adelowo](https://lanre.wtf/blog/2017/04/08/testing-http-handlers-go/)

There’s also the free ebook “Learn Go With Tests” on GitHub that looks really  nice, though I only used it a little for this post:

GitHub 上还有一本免费的电子书“Learn Go With Tests”，看起来很不错，虽然我在这篇文章中只使用了一点：

- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

Other references I used while researching, but didn’t mention in the post: 

我在研究时使用的其他参考资料，但在帖子中没有提到：

