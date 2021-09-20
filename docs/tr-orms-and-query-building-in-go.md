# ORMs and Query Building in Go

# 在 Go 中构建 ORM 和查询

*Posted on Sat 13 2019 to [Programming](https://andrewpillar.com/programming)*

Recently, I have been looking into various solutions for interacting with databases with ease in Go. My go to library for database work in Go is [sqlx](https://github.com/jmoiron/sqlx), this makes unmarshalling the data from the database into structs a cinch. You write out your SQL query, tag your structs using the `db` tag, and let [sqlx](https://github.com/jmoiron/sqlx) handle the rest. However, the main problem I have encountered, was with idiomatic query building. This led me to investigate this problem, and jot down some of my thoughts in this post.

最近，我一直在研究在 Go 中轻松与数据库交互的各种解决方案。我在 Go 中使用的数据库工作库是 [sqlx](https://github.com/jmoiron/sqlx)，这使得将数据库中的数据解组为结构体变得轻而易举。你写出你的 SQL 查询，使用 `db` 标签标记你的结构，然后让 [sqlx](https://github.com/jmoiron/sqlx) 处理剩下的事情。但是，我遇到的主要问题是惯用查询构建。这促使我调查这个问题，并在这篇文章中记下我的一些想法。

**TL;DR** First class functions are an idiomatic way of doing SQL query building in Go. Check out the repository containing some example code I wrote testing this out: https://github.com/andrewpillar/query.

**TL;DR** First class 函数是在 Go 中进行 SQL 查询构建的惯用方式。查看包含我编写的一些示例代码的存储库：https://github.com/andrewpillar/query。

## GORM, Layered Complexity, and the Active Record Pattern

## GORM、分层复杂性和 Active Record 模式

Most people who dip their toe into database work in Go, will most likely be pointed towords [gorm](https://gorm.io) for working with databases. It is a fairly fully featured ORM, that supports migrations, relations, transactions, and much more. For those who have worked with ActiveRecord, or Eloquent, GORM's usage would be some what familiar to you.

大多数在 Go 中涉足数据库工作的人很可能会使用 [gorm](https://gorm.io) 来处理数据库。它是一个功能相当齐全的 ORM，支持迁移、关系、事务等等。对于那些使用过 ActiveRecord 或 Eloquent 的人来说，GORM 的用法对您来说可能有些熟悉。

I have used GORM briefly before, and for simple CRUD based applications this is fine. However, when it comes to more layered complexity, I find that it falls short. Assume we are building a blogging application, and we allow users to search for posts via the `search` query string in a URL. If this is present, we want to constrain the query with a `WHERE title LIKE`, otherwise we do not.

我之前简单地使用过 GORM，对于简单的基于 CRUD 的应用程序，这很好。然而，当涉及到更多层次的复杂性时，我发现它不够用。假设我们正在构建一个博客应用程序，我们允许用户通过 URL 中的 `search` 查询字符串搜索帖子。如果存在，我们希望使用“WHERE title LIKE”来约束查询，否则我们不这样做。

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")

db := gorm.Open("postgres", "...")

if search != "" {
    db = db.Where("title LIKE ?", "%" + search + "%")
}

db.Find(&posts)
```

Nothing to controversial, we simply check to see if we have a value and modify the invocation to GORM itself. However, what if we wanted to allow searching of posts after a certain date? We would need to add some more checks, first to see if the `after` query string is present in the URL, and if so modify the query appropriately.

没有什么可争议的，我们只是检查我们是否有一个值并将调用修改为 GORM 本身。但是，如果我们想允许在某个日期之后搜索帖子怎么办？我们需要添加更多检查，首先查看 URL 中是否存在 `after` 查询字符串，如果存在，则适当修改查询。

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")
after := r.URL.Query().Get("after")

db := gorm.Open("postgres", "...")

if search != "" {
    db = db.Where("title LIKE ?", "%" + search + "%")
}

if after != "" {
    db = db.Where("created_at > ?", after)
}

db.Find(&posts)
```

So we add another check to determine if the invocation should be modified. This is working fine so far, but things could start getting out of hand. Ideally what we would want is some way of extending GORM with some custom callbacks that would accept the `search`, and `after` variables regardless of their value, and defer the logic to that custom callback. GORM does support a [plugin system](https://gorm.io/docs/write_plugins.html), for writing custom callbacks, however it seems this is more suited for modifying table state upon certain operations.

所以我们添加另一个检查来确定是否应该修改调用。到目前为止，这一切正常，但事情可能会开始失控。理想情况下，我们想要的是使用一些自定义回调扩展 GORM 的某种方式，这些回调将接受 `search` 和 `after` 变量而不管它们的值，并将逻辑推迟到该自定义回调。 GORM 确实支持 [插件系统](https://gorm.io/docs/write_plugins.html)，用于编写自定义回调，但是这似乎更适合在某些操作时修改表状态。

As demonstrated above, I find GORM's biggest drawback is how cumbersone it can be to do layered complexity. More often than not when writing SQL queries, you will want this. Trying to determine if you want to add a `WHERE` clause to a query based off of some user input, or how you should order the records.

如上所述，我发现 GORM 的最大缺点是它可以实现分层复杂性。通常在编写 SQL 查询时，您会想要这个。试图确定是否要根据某些用户输入向查询添加“WHERE”子句，或者您应该如何对记录进行排序。

I believe this comes down to one thing, and I made a [comment](https://news.ycombinator.com/item?id=19851753) about this some time ago on HN:

我相信这归结为一件事，我前段时间在 HN 上做了一个 [评论](https://news.ycombinator.com/item?id=19851753)：

> Personally I think an active record style ORM for Go like gorm is a poor fit for a language that doesn't come across as inherently OOP. Going through some of the documentation for gorm, it seems to rely heavily on method chaining which for Go seems wrong considering how errors are handled in that language. In my opinion, an ORM should be as idiomatic to the language as possible.

> 就我个人而言，我认为像 gorm 这样的 Go 的活动记录样式 ORM 不适合一种天生就不是 OOP 的语言。通过阅读 gorm 的一些文档，它似乎严重依赖方法链，考虑到该语言中的错误是如何处理的，这对于 Go 来说似乎是错误的。在我看来，ORM 应该尽可能地符合语言习惯。

This comment was made on a submission of the blog post [To ORM or not to ORM](https://eli.thegreenplace.net/2019/to-orm-or-not-to-orm/), which I highly recommend you read. The author of the post approaches the same conclusion about GORM that I did.

此评论是在博客文章 [To ORM or not to ORM](https://eli.thegreenplace.net/2019/to-orm-or-not-to-orm/) 的提交上做出的，我强烈推荐你读。这篇文章的作者对 GORM 得出的结论与我得出的结论相同。

## Idiomatic Query Building in Go 

## Go 中的惯用查询构建

The `database/sql` package in the standard library is great for interacting with databases. And [sqlx](https://github.com/jmoiron/sqlx) is a fine extension on top of that for handling the return of data. However, this still doesn't fully solve the problem at hand. How can we effectively build complex queries programmatically that is idiomatic to Go. Assume we were using [sqlx](https://github.com/jmoiron/sqlx) for the same query above, what would that look like?

标准库中的`database/sql` 包非常适合与数据库交互。 [sqlx](https://github.com/jmoiron/sqlx) 是一个很好的扩展，用于处理数据的返回。然而，这仍然不能完全解决手头的问题。我们如何以编程方式有效地构建 Go 惯用的复杂查询。假设我们使用 [sqlx](https://github.com/jmoiron/sqlx) 进行上述相同的查询，那会是什么样子？

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")
after := r.URL.Query().Get("after")

db := sqlx.Open("postgres", "...")

query := "SELECT * FROM posts"
args := make([]interface{}, 0)

if search != "" {
    query += " WHERE title LIKE ?"
    args = append(args, search)
}

if after != "" {
    if search != "" {
        query += " AND "
    } else {
        query += " WHERE "
    }
    
    query += "created_at > ?"
    
    args = append(args, after)
}

err := db.Select(&posts, sqlx.Rebind(query), args...)
```

Not much better than what we did with GORM, in fact much uglier. We're checking if `search` exists twice just so we can have the correct SQL grammar in place for the query, storing our arguments in an `[]interface{}` slice, and concatenating onto a string. This too, is not as extensible, or easy to maintain. Ideally we want to be able to build up the query, and hand it off to [sqlx](https://github.com/jmoiron/sqlx) to handle the rest. So, what would an idiomatic query builder in Go look like? Well, it would come in one of two forms in my opinion, the first being one that utilises option structs, and the other that utilises first class functions.

并不比我们对 GORM 所做的好多少，实际上要丑得多。我们正在检查 `search` 是否存在两次，以便我们可以为查询设置正确的 SQL 语法，将我们的参数存储在一个 `[]interface{}` 切片中，并连接到一个字符串上。这也不是可扩展或易于维护的。理想情况下，我们希望能够构建查询，并将其交给 [sqlx](https://github.com/jmoiron/sqlx) 来处理其余部分。那么，Go 中惯用的查询构建器会是什么样子呢？好吧，在我看来，它将采用两种形式之一，第一种是利用选项结构的，另一种是利用一流的函数。

Let's take a look at [squirrel](https://github.com/masterminds/squirrel). This library offers the ability to build up queries, and execute them directly in a way that I find rather idiomatic to Go. Here though, we will only be focussing on the query building aspect.

我们来看看[松鼠](https://github.com/masterminds/squirrel)。这个库提供了构建查询的能力，并以我认为非常适合 Go 的方式直接执行它们。不过，在这里，我们将只关注查询构建方面。

With [squirrel](https://github.com/masterminds/squirrel), we can implement our above logic like so.

使用 [squirrel](https://github.com/masterminds/squirrel)，我们可以像这样实现上面的逻辑。

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")
after := r.URL.Query().Get("after")

eqs := make([]sq.Eq, 0)

if search != "" {
    eqs = append(eqs, sq.Like{"title", "%" + search + "%"})
}

if after != "" {
    eqs = append(eqs, sq.Gt{"created_at", after})
}

q := sq.Select("*").From("posts")

for _, eq := range eqs {
    q = q.Where(eq)
}

query, args, err := q.ToSql()

if err != nil {
    return
}

err := db.Select(&posts, query, args...)
```

This is slightly better than what we had with GORM, and miles better than the string concatenation we were doing before. However, it still comes across as slightly cumbersone to write. [squirrel](https://github.com/masterminds/squirrel) uses option structs for some of the clauses in an SQL query. Optional structs are common pattern in Go for APIs that aim to be highly configurable.

这比我们使用 GORM 的略好，比我们之前做的字符串连接好几英里。然而，写起来仍然有点麻烦。 [squirrel](https://github.com/masterminds/squirrel) 对 SQL 查询中的某些子句使用选项结构。可选结构是 Go for API 中的常见模式，旨在实现高度可配置。

An API for query building in Go should fulfill both of these needs:

在 Go 中构建查询的 API 应该满足这两个需求：

- Idiomacy
- Extensibility

- 习语
- 可扩展性

How can this be achieved with Go?

如何用 Go 实现这一点？

## First Class Functions for Query Building

## 用于查询构建的一流函数

Dave Cheney has written two blog posts about first class functions, based off of a post made by Rob Pike about the same topic. For those interested they can be found here:

Dave Cheney 写了两篇关于一流函数的博客文章，基于 Rob Pike 发表的关于同一主题的文章。有兴趣的可以在这里找到它们：

- [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Do not fear the first class functions](https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions)

- [自引用功能及选项设计](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [友好 API 的功能选项](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [不要害怕一流的功能](https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions)

I'd highly recommend reading the above three posts, and using the pattern they suggest when you next come to implement an API that needs to be highly configurable.

我强烈建议您阅读以上三篇文章，并在您下次实现需要高度可配置的 API 时使用他们建议的模式。

Below is an example of what this might look like for query building:

以下是查询构建的示例：

```go
posts := make([]*Post, 0)

db := sqlx.Open("postgres", "...")

q := Select(
    Columns("*"),
    Table("posts"),
)

err := db.Select(&posts, q.Build(), q.Args()...)
```

A naive example, I know. But let's take a look at how we might implement an API like this so that it can be used for query building. First, we should implement a query struct to keep track of the query's state whilst it's being built.

一个天真的例子，我知道。但是让我们来看看我们如何实现这样的 API，以便它可以用于查询构建。首先，我们应该实现一个查询结构来跟踪正在构建的查询的状态。

```go
type statement uint8

type Query struct {
    stmt  statement
    table []string
    cols  []string
    args  []interface{}
}

const (
    _select statement = iota
)
```

The above struct will keep track of the statement we're building, whether it's `SELECT`, `UPDATE`, `INSERT`, or `DELETE`, table being operated on, the columns we're working with, and the arguments that will be passed to the final query. To keep this simple, let's focus on implementing the `SELECT` statement for the query builder.

上面的结构将跟踪我们正在构建的语句，无论是“SELECT”、“UPDATE”、“INSERT”还是“DELETE”、正在操作的表、我们正在使用的列以及参数将传递给最终查询。为了简单起见，让我们专注于为查询构建器实现 `SELECT` 语句。

Next, we need to define a type that can be used for modifying the query we're building. This is the type that would be passed numerous times as a first class function. Each time this function is called, it should return the newly modified query, if applicable.

接下来，我们需要定义一个可用于修改我们正在构建的查询的类型。这是将作为第一类函数多次传递的类型。每次调用此函数时，它都应返回新修改的查询（如果适用）。

```go
type Option func(q Query) Query
```

We can now implement the first part of the builder, the `Select` function. This will begin building a query for the `SELECT` statement we want to build up.

我们现在可以实现构建器的第一部分，`Select` 函数。这将开始为我们想要构建的 `SELECT` 语句构建一个查询。

```go
func Select(opts ...Option) Query {
    q := Query{
        stmt: select_,
    }

    for _, opt := range opts {
        q = opt(q)
    }

    return q
}
```

You should now be able to see how everything is slowly coming together, and how the `UPDATE`, `INSERT`, and `DELETE` statements could be trivially implemented too. Without actually implementing some options to pass to `Select`, the above function is fairly useless, so let's do that.

您现在应该能够看到所有内容是如何慢慢融合在一起的，以及如何轻松实现 `UPDATE`、`INSERT` 和 `DELETE` 语句。没有实际实现一些选项来传递给`Select`，上面的函数是相当无用的，所以让我们这样做。

```go
func Columns(cols ...string) Option {
    return func(q Query) Query {
        q.cols = cols

        return q
    }
}

func Table(table string) Option {
    return func(q Query) Query {
        q.table = table

        return q
    }
}
```

As you can see, we implement these first class functions in a way so that they return the underlying option function that will be called. It would be typically expected for the option function to modify the query passed to it, and for a copy to be returned.

如您所见，我们以某种方式实现了这些第一类函数，以便它们返回将被调用的底层期权函数。通常期望选项函数修改传递给它的查询，并返回一个副本。

For this to be useful for our use case of building complex queries, we ought to implement the ability to add `WHERE` clauses to our query. This will require having to keep track of the various `WHERE` clauses in the query too.

为了使这对我们构建复杂查询的用例有用，我们应该实现将“WHERE”子句添加到我们的查询的能力。这也需要跟踪查询中的各种“WHERE”子句。

```go
type where struct {
    col string
    op  string
    val interface{}
}

type Query struct {
    stmt   statement
    table  []string
    cols   []string
    wheres []where
    args   []interface{}
}
```

We define a custom type for a `WHERE` clause, and add a `wheres` property to the original `Query` struct. Let's implement two types of `WHERE` clauses for our needs, the first being `WHERE LIKE`, and the other being `WHERE >`.

我们为 `WHERE` 子句定义了一个自定义类型，并将一个 `wheres` 属性添加到原始的 `Query` 结构中。让我们根据需要实现两种类型的`WHERE` 子句，第一种是`WHERE LIKE`，另一种是`WHERE >`。

```go
func WhereLike(col string, val interface{}) Option {
    return func(q Query) Query {
        w := where{
            col: col,
            op:  "LIKE",
            val: fmt.Sprintf("$%d", len(q.args) + 1),
        }

        q.wheres = append(q.wheres, w)
        q.args = append(q.args, val)

        return q
    }
}

func WhereGt(col string, val interface{}) Option {
    return func(q Query) Query {
        w := where{
            col: col,
            op:  ">",
            val: fmt.Sprintf("$%d", len(q.args) + 1),
        }

        q.wheres = append(q.wheres, w)
        q.args = append(q.args, val)

        return q
    }
}
```

When handling the addition of a `WHERE` clause to a query, we appropriately handle the bindvar syntax for the underlying SQL driver, Postgres in this case, and store the actual value itself in the `args` slice on the query.

在处理向查询添加`WHERE` 子句时，我们适当地处理底层SQL 驱动程序的bindvar 语法，在本例中为Postgres，并将实际值本身存储在查询的`args` 切片中。

So, with what little we have implemented we should be able to achieve what we want in an idiomatic way.

因此，通过我们实施的很少，我们应该能够以惯用的方式实现我们想要的。

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")
after := r.URL.Query().Get("after")

db := sqlx.Open("postgres", "...")

opts := []Option{
    Columns("*"),
    Table("posts"),
}

if search != "" {
    opts = append(opts, WhereLike("title", "%" + search + "%"))
}

if after != "" {
    opts = append(opts, WhereGt("created_at", after))
}

q := Select(opts...)

err := db.Select(&posts, q.Build(), q.Args()...)
```

Slightly better, but still not great. However, we can extend the functionality to get what we want. So, let's implement some functions that will return options for our specific needs.

稍微好一点，但仍然不是很好。但是，我们可以扩展功能以获得我们想要的。因此，让我们实现一些函数，这些函数将返回满足我们特定需求的选项。

```go
func Search(col, val string) Option {
    return func(q Query) Query {
        if val == "" {
            return q
        }

        return WhereLike(col, "%" + val + "%")(q)
    }
}

func After(val string) Option {
    return func(q Query) Query {
        if val == "" {
            return q
        }

        return WhereGt("created_at", val)(q)
    }
}
```

With the above two functions implemented we can now cleanly build up a somewhat complex query for our use case. Both of these functions will only modify the query, if the values passed to them are considered correct.

实现上述两个函数后，我们现在可以为我们的用例干净地构建一个有点复杂的查询。如果传递给它们的值被认为是正确的，这两个函数只会修改查询。

```go
posts := make([]Post, 0)

search := r.URL.Query().Get("search")
after := r.URL.Query().Get("after")

db := sqlx.Open("postgres", "...")

q := Select(
    Columns("*"),
    Table("posts"),
    Search("title", search),
    After(after),
)

err := db.Select(&posts, q.Build(), q.Args()...)
```

I find this to be a rather idiomatic way of building up complex queries in Go. Now, of course you've made it this far in the post, and must be wondering, "That's good and all, but you didn't imeplement the Build(), or Args() methods". This is true, to an extent. In the interest of not wanting to prolong this post any further than needed, I didn't bother. So, if you are interested in some of the ideas presented here, take a look at the [code](https://github.com/andrewpillar/query), I submitted to GitHub. It's nothing to rigorous, and doesn't cover everything a query builder would need, it's lacking `JOIN`, for example and supports only the Postgres bindvar.

我发现这是在 Go 中构建复杂查询的一种相当惯用的方式。现在，当然您已经在帖子中做到了这一点，并且一定想知道，“这很好，但您没有实现 Build() 或 Args() 方法”。在某种程度上，这是真的。为了不想把这篇文章拖得比需要的还要长，我没有打扰。因此，如果您对这里提出的一些想法感兴趣，请查看我提交给 GitHub 的 [代码](https://github.com/andrewpillar/query)。它没有什么严格的，并且没有涵盖查询构建器需要的所有内容，例如它缺少`JOIN`，并且仅支持 Postgres bindvar。

If you have any disagreements with what I have said in this post, or would like to discuss this further, then please reach out to me at [me@andrewpillar.com](mailto:me@andrewpillar.com). 

如果您对我在这篇文章中所说的有任何异议，或者想进一步讨论，请通过 [me@andrewpillar.com](mailto:me@andrewpillar.com) 与我联系。

