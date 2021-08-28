# Tests Make Your Code Inherently Better

# 测试使您的代码本质上更好

From: https://www.mitchdennett.com/tests-make-your-code-inherently-better/

I have been developing software for about 8 years now but my journey into testing has only been within the last 2 years. In college we had half a semester on testing. Not even a whole class on it. It was  definitely an after thought and not a main focus it should be. Writing  tests has made my code cleaner and more concise than any other endeavor. Let's take a quick look at why.

我已经开发软件大约 8 年了，但我的测试之旅只发生在最近 2 年。在大学里，我们有半个学期的考试时间。甚至整个班级都没有。这绝对是事后的想法，而不是它应该的主要关注点。编写测试使我的代码比任何其他努力都更简洁、更简洁。让我们快速了解一下原因。

In this short example we are going to look at a REST API that returns a list of Recipes.

在这个简短的示例中，我们将查看一个返回食谱列表的 REST API。

## The "Bad" Code

## “坏”代码

Let's take a look at some quote unquote bad code.

让我们来看看一些引用 unquote 的坏代码。

```go
func (h *Handler) handleListRecipes(w http.ResponseWriter, r *http.Request) {
    pages, ok := r.URL.Query()["page"]

    if !ok ||len(pages[0]) < 1 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    page, err := strconv.Atoi(pages[0])

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    queryDocStmt := `SELECT recipe_id, title from recipe limit 50 offset $1`

    var offset int
    if page-1 < 0 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    offset = (page - 1) * 50

    rows, err := h.DB.Query(queryDocStmt, offset)
    itemsList := make([]*better.Recipe, 0)

    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    for rows.Next() {
        var item better.Recipe

        if err := rows.Scan(&item.ID, &item.Title);err != nil {
            // Check for a scan error.
            // Query rows will be closed with defer.
            log.Println(err)
            continue
        }

        itemsList = append(itemsList, &item)

    }

    retJSON, err := json.Marshal(itemsList)
    fmt.Fprintf(w, string(retJSON))
}
```


This code is a  mess. It is not testable at all. When unit testing we don't want to  actually make a call to the database. We can assume the standard library has been tested exhaustively.

这段代码一团糟。它根本无法测试。在进行单元测试时，我们不想实际调用数据库。我们可以假设标准库已经过详尽的测试。

So back to our code. There is no easy way to mock out our database calls so we can write a good test.

回到我们的代码。没有简单的方法来模拟我们的数据库调用，所以我们可以编写一个好的测试。

## Refactoring Out Our Database Calls

## 重构我们的数据库调用

So we need to make the database call mockable. This just means that we  need a way to replace the actual database call with something else  during our test so we don't actually make to request to the database. So let's pull it out and create a `Store` that is responsible for all interactions with the database.

所以我们需要使数据库调用可模拟。这只是意味着我们需要一种方法来在测试期间用其他东西替换实际的数据库调用，这样我们实际上不会向数据库发出请求。因此，让我们将其拉出并创建一个负责与数据库的所有交互的 `Store`。

```go
type Store struct {
    DB *sql.DB
}

//ListRecipes will list all the recipes for a given page
func (d *DB) ListRecipes(page int) ([]*better.Recipe, error) {
    queryDocStmt := `SELECT recipe_id, title from recipe limit 50 offset $1`

    var offset int
    if page-1 < 0 {
        return nil, errors.New("Bad Request")
    }

    offset = (page - 1) * 50

    rows, err := d.DB.Query(queryDocStmt, offset)
    itemsList := make([]*better.Recipe, 0)

    if err != nil ||rows == nil {
        log.Println(err)
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var item Recipe

        if err := rows.Scan(&item.ID, &item.Title);err != nil {
            // Check for a scan error.
            // Query rows will be closed with defer.
            log.Println(err)
            continue
        }

        itemsList = append(itemsList, &item)

    }

    return itemsList, nil

}
```


So this is a step in the right direction. With that in place we can go and edit our `handleListRecipes` func.

所以这是朝着正确方向迈出的一步。有了这个，我们可以去编辑我们的 `handleListRecipes` 函数。

```go
func (h *Handler) handleListRecipes(w http.ResponseWriter, r *http.Request) {
    page, err := strconv.Atoi(r.URL.Query().Get("page"))
    if err != nil {
        http.Error(w, "invalid page", http.StatusBadRequest)
        return
    }

    items, err := h.RecipeStore.ListRecipes(page)
    if err != nil {
        log.Print("http error", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(items);err != nil {
        log.Print("http json encoding error", err)
    }
}
```


As you can see  this made our handler function way cleaner, more readable and more  concise. We are still not there yet though. In order to be able to mock  (replace) our `RecipeStore` in our test we need to create an  Interface. If the only thing you take away from this is to use  Interfaces more I will consider this a success.

如您所见，这使我们的处理程序函数变得更清晰、更易读和更简洁。不过，我们还没有到那里。为了能够在我们的测试中模拟（替换）我们的 `RecipeStore`，我们需要创建一个接口。如果您从中获得的唯一一件事就是更多地使用接口，我会认为这是成功的。

## Interfaces, Interfaces, Interfaces

## 接口，接口，接口

So let's create our interface so we can mock out our `RecipeStore` when testing.

所以让我们创建我们的界面，这样我们就可以在测试时模拟我们的“RecipeStore”。

```go
type RecipeService interface {
    ListRecipes(page int) ([]*Recipe, error)
}
```


Super simple. But super powerful.

超级简单。但是超级强大。

With that, just change our Handler to take a `RecipeService` interface instead of the concrete `RecipeStore`. And voila we can now get started on writing our test.

有了这个，只需更改我们的 Handler 以采用“RecipeService”接口而不是具体的“RecipeStore”。瞧，我们现在可以开始编写我们的测试了。

```go
items, err := h.RecipeService.ListRecipes(page)
```


## Tests 

## 测试

So now that we have our code refactored we can start with our tests. It is now going to be super easy to replace our actual Store with a mocked  service that we can now control in our tests. No more database calls.

所以现在我们重构了我们的代码，我们可以开始我们的测试。现在用我们现在可以在测试中控制的模拟服务替换我们的实际 Store 将变得非常容易。没有更多的数据库调用。

First let's create a new `mock` package to hold all our mocks and create a `recipe_service.go` file in there with the following.

首先让我们创建一个新的 `mock` 包来保存我们所有的模拟，并在其中创建一个包含以下内容的 `recipe_service.go` 文件。

```go
package mock

import (
    better "github.com/mitchdennett/tests-make-your-code-inherently-better"
)

type MockRecipeService struct {
    ListRecipesFunc func(page int) ([]*better.Recipe, error)
}

func (s *MockRecipeService) ListRecipes(page int) ([]*better.Recipe, error) {
    return s.ListRecipesFunc(page)
}
```


This mock allows our test to inject a function into it to test different outcomes. Now we can write our first test.

这个模拟允许我们的测试向其中注入一个函数来测试不同的结果。现在我们可以编写我们的第一个测试。

```go
func TestListRecipes(t *testing.T) {
    req, err := http.NewRequest("GET", "/recipes?page=1", nil)
    if err != nil {
        t.Fatal(err)
    }

    var store mock.MockRecipeService
    store.ListRecipesFunc = func(page int) ([]*better.Recipe, error) {
        return []*better.Recipe{{ID: 1, Title: "Pasta"}}, nil
    }
    handler := Handler{RecipeService: &store}

    rr := httptest.NewRecorder()

    handler.ServeHTTP(rr, req)

    if status := rr.Code;status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}
```


## Why is this better?

## 为什么这更好？

It's better for multiple reasons. The obvious one is that we can now  correctly test our code which is obviously better. Secondly, our code is more concise, clear and readable. Finally, we have removed our  dependency with the database from all but one package. Should we ever  need to change our data store we can simply implement the `RecipeService` interface and nothing else.

出于多种原因，这更好。显而易见的是，我们现在可以正确测试我们的代码，这显然更好。其次，我们的代码更加简洁、清晰和可读。最后，我们从除了一个包之外的所有包中删除了对数据库的依赖。如果我们需要更改我们的数据存储，我们可以简单地实现`RecipeService` 接口而不是其他任何东西。

Now there is definitely more we can do to this code to improve it. But by  simply focusing on testing we inherently have to write better and  cleaner code. Check out this Github Repo with the final code https://github.com/mitchdennett/tests-make-your-code-inherently-better.

现在我们肯定可以对这段代码做更多的事情来改进它。但是通过简单地专注于测试，我们本质上必须编写更好、更清晰的代码。使用最终代码 https://github.com/mitchdennett/tests-make-your-code-inherently-better 查看此 Github Repo。

# Join the Newsletter

# 加入时事通讯

Subscribe to get our latest content by email. We won't send you spam. Unsubscribe at any time. 

订阅以通过电子邮件获取我们的最新内容。我们不会向您发送垃圾邮件。随时退订。


