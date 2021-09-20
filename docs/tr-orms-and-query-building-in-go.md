# ORMs and Query Building in Go

# 在 Go 中构建 ORM 和查询

_Posted on Sat 13 2019 to [Programming](http://andrewpillar.com/programming)_ _Updated at 22:30 on Mon 23 Mar 2020_

_2019 年 3 月 13 日星期六发布到 [Programming](http://andrewpillar.com/programming)_ _2020 年 3 月 23 日星期一 22:30 更新_

Recently, I have been looking into various solutions for interacting with
databases with ease in Go. My go to library for database work in Go is
[sqlx](https://github.com/jmoiron/sqlx), this makes unmarshalling the data from
the database into structs a cinch. You write out your SQL query, tag your
structs using the `db` tag, and let [sqlx](https://github.com/jmoiron/sqlx)
handle the rest. However, the main problem I have encountered, was with
idiomatic query building. This led me to investigate this problem, and jot down
some of my thoughts in this post.

最近，我一直在寻找各种解决方案来与
在 Go 中轻松使用数据库。我去图书馆在 Go 中进行数据库工作是
[sqlx](https://github.com/jmoiron/sqlx)，这使得解组来自
数据库变成结构很简单。你写出你的 SQL 查询，标记你的
使用 `db` 标签构建结构，并让 [sqlx](https://github.com/jmoiron/sqlx)
处理剩下的。但是，我遇到的主要问题是
惯用查询构建。这让我调查了这个问题，并记下了
我在这篇文章中的一些想法。

**TL;DR** First class functions are an idiomatic way of doing SQL query building
in Go. Check out the repository containing some example code I wrote testing
this out: [https://github.com/andrewpillar/query](https://github.com/andrewpillar/query).

**TL;DR** First class 函数是构建 SQL 查询的惯用方法
在去。查看包含我编写测试的一些示例代码的存储库
这出：[https://github.com/andrewpillar/query](https://github.com/andrewpillar/query)。

## GORM, Layered Complexity, and the Active Record Pattern

## GORM、分层复杂性和 Active Record 模式

Most people who dip their toe into database work in Go, will most likely be
pointed towords [gorm](https://gorm.io) for working with databases. It is a
fairly fully featured ORM, that supports migrations, relations, transactions,
and much more. For those who have worked with ActiveRecord, or Eloquent, GORM's
usage would be some what familiar to you.

大多数在 Go 中涉足数据库工作的人很可能是
指向词 [gorm](https://gorm.io) 用于处理数据库。它是一个
功能相当齐全的 ORM，支持迁移、关系、事务、
以及更多。对于那些使用过 ActiveRecord 或 Eloquent 的人，GORM 的
用法对您来说有些熟悉。

I have used GORM briefly before, and for simple CRUD based applications this is
fine. However, when it comes to more layered complexity, I find that it falls
short. Assume we are building a blogging application, and we allow users to
search for posts via the `search` query string in a URL. If this is present, we
want to constrain the query with a `WHERE title LIKE`, otherwise we do not.

我之前简单地使用过 GORM，对于简单的基于 CRUD 的应用程序，这是
美好的。然而，当涉及到更多层次的复杂性时，我发现它下降了
短的。假设我们正在构建一个博客应用程序，我们允许用户
通过 URL 中的 `search` 查询字符串搜索帖子。如果存在这种情况，我们
想用“WHERE title LIKE”来约束查询，否则我们不这样做。

```
posts := make([]Post, 0)

search := r.URL.Query().Get("search")

db := gorm.Open("postgres", "...")

if search != "" {
    db = db.Where("title LIKE ?", "%" + search + "%")
}

db.Find(&posts)

```

Nothing to controversial, we simply check to see if we have a value and modify
the invocation to GORM itself. However, what if we wanted to allow searching of
posts after a certain date? We would need to add some more checks, first to see
if the `after` query string is present in the URL, and if so modify the query
appropriately.

没什么可争议的，我们只是检查一下我们是否有一个值并修改
对 GORM 本身的调用。但是，如果我们想允许搜索
在某个日期之后发布？我们需要添加更多的检查，首先看看
如果 URL 中存在 `after` 查询字符串，则修改查询
适当地。

```
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

So we add another check to determine if the invocation should be modified. This
is working fine so far, but things could start getting out of hand. Ideally what
we would want is some way of extending GORM with some custom callbacks that
would accept the `search`, and `after` variables regardless of their value, and
defer the logic to that custom callback. GORM does support a
[plugin system](https://gorm.io/docs/write_plugins.html), for writing custom
callbacks, however it seems this is more suited for modifying table state upon
certain operations.

所以我们添加另一个检查来确定是否应该修改调用。这个
到目前为止工作正常，但事情可能会开始失控。理想情况下是什么
我们想要的是通过一些自定义回调来扩展 GORM
将接受 `search` 和 `after` 变量，而不管它们的值如何，并且
将逻辑推迟到该自定义回调。 GORM 确实支持
[插件系统](https://gorm.io/docs/write_plugins.html)，用于编写自定义
回调，但是这似乎更适合修改表状态
某些操作。

As demonstrated above, I find GORM's biggest drawback is how cumbersone it can
be to do layered complexity. More often than not when writing SQL queries, you
will want this. Trying to determine if you want to add a `WHERE` clause to a
query based off of some user input, or how you should order the records.

如上所示，我发现 GORM 的最大缺点是它的笨拙程度
是做分层的复杂性。通常在编写 SQL 查询时，您
会想要这个。试图确定您是否想将 `WHERE` 子句添加到
基于某些用户输入的查询，或者您应该如何对记录进行排序。

I believe this comes down to one thing, and I made a
[comment](https://news.ycombinator.com/item?id=19851753) about this some time
ago on HN:

我相信这归结为一件事，我做了一个
[评论](https://news.ycombinator.com/item?id=19851753) 有时间了
以前在 HN 上：

> Personally I think an active record style ORM for Go like gorm is a poor fit
> for a language that doesn't come across as inherently OOP. Going through some
> of the documentation for gorm, it seems to rely heavily on method chaining which
> for Go seems wrong considering how errors are handled in that language. In my
> opinion, an ORM should be as idiomatic to the language as possible.

> 我个人认为像 gorm 这样的 Go 的活动记录样式 ORM 不太合适
> 对于一种天生就不是 OOP 的语言。经过一些
> 在 gorm 的文档中，它似乎严重依赖于方法链
> 考虑到该语言如何处理错误，Go 似乎是错误的。在我的
> 意见，ORM 应该尽可能符合语言习惯。

This comment was made on a submission of the blog post
[To ORM or not to ORM](https://eli.thegreenplace.net/2019/to-orm-or-not-to-orm/),
which I highly recommend you read. The author of the post approaches the same
conclusion about GORM that I did.

此评论是在博客文章的提交上发表的
[ORM 或不 ORM](https://eli.thegreenplace.net/2019/to-orm-or-not-to-orm/),
我强烈建议你阅读。该帖子的作者接近相同
我所做的关于 GORM 的结论。

## Idiomatic Query Building in Go

## Go 中的惯用查询构建

The `database/sql` package in the standard library is great for interacting with
databases. And [sqlx](https://github.com/jmoiron/sqlx) is a fine extension on
top of that for handling the return of data. However, this still doesn't fully 

标准库中的`database/sql` 包非常适合与
数据库。 [sqlx](https://github.com/jmoiron/sqlx) 是一个很好的扩展
最重要的是处理数据的返回。然而，这仍然不完全

solve the problem at hand. How can we effectively build complex queries
programmatically that is idiomatic to Go. Assume we were using
[sqlx](https://github.com/jmoiron/sqlx) for the same query above, what would
that look like?

解决手头的问题。我们如何有效地构建复杂的查询
以编程方式，这是 Go 惯用的。假设我们正在使用
[sqlx](https://github.com/jmoiron/sqlx) 对于上面相同的查询，会怎样
那个样子？

```
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

Not much better than what we did with GORM, in fact much uglier. We're checking
if `search` exists twice just so we can have the correct SQL grammar in place
for the query, storing our arguments in an `[]interface{}` slice, and
concatenating onto a string. This too, is not as extensible, or easy to maintain.
Ideally we want to be able to build up the query, and hand it off to
[sqlx](https://github.com/jmoiron/sqlx) to handle the rest. So, what would an
idiomatic query builder in Go look like? Well, it would come in one of two forms
in my opinion, the first being one that utilises option structs, and the other
that utilises first class functions.

并不比我们对 GORM 所做的好多少，实际上要丑得多。我们正在检查
如果`search`存在两次只是为了我们可以有正确的SQL语法
对于查询，将我们的参数存储在一个 `[]interface{}` 切片中，以及
连接到一个字符串上。这也不是可扩展或易于维护的。
理想情况下，我们希望能够建立查询，并将其交给
[sqlx](https://github.com/jmoiron/sqlx) 来处理剩下的。那么，什么会
Go 中的惯用查询构建器是什么样的？嗯，它会以两种形式之一出现
在我看来，第一个是利用选项结构的，另一个是
使用一流的功能。

Let's take a look at [squirrel](https://github.com/masterminds/squirrel). This
library offers the ability to build up queries, and execute them directly in a
way that I find rather idiomatic to Go. Here though, we will only be focussing
on the query building aspect.

我们来看看[松鼠](https://github.com/masterminds/squirrel)。这个
库提供了构建查询的能力，并直接在一个
我觉得 Go 相当惯用的方式。不过，在这里，我们将只关注
在查询构建方面。

With [squirrel](https://github.com/masterminds/squirrel), we can implement our
above logic like so.

使用 [squirrel](https://github.com/masterminds/squirrel)，我们可以实现我们的
上面的逻辑是这样的。

```
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

This is slightly better than what we had with GORM, and miles better than the
string concatenation we were doing before. However, it still comes across as
slightly cumbersone to write. [squirrel](https://github.com/masterminds/squirrel)
uses option structs for some of the clauses in an SQL query. Optional structs
are common pattern in Go for APIs that aim to be highly configurable.

这比我们在 GORM 上的略好，比 GORM 好几英里
我们之前做的字符串连接。然而，它仍然是这样的
写起来略麻烦。 [松鼠](https://github.com/masterminds/squirrel)
对 SQL 查询中的某些子句使用选项结构。可选结构
是 Go for API 中的常见模式，旨在高度可配置。

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

Dave Cheney has written two blog posts about first class functions, based off of
a post made by Rob Pike about the same topic. For those interested they can be
found here:

Dave Cheney 写了两篇关于一流函数的博客文章，基于
Rob Pike 发表的关于同一主题的帖子。对于那些有兴趣的人，他们可以
在这里找到：

- [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Do not fear the first class functions](https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions)

- [自引用功能及选项设计](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [友好 API 的功能选项](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [不要害怕一流的功能](https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions)

I'd highly recommend reading the above three posts, and using the pattern they
suggest when you next come to implement an API that needs to be highly
configurable.

我强烈建议阅读以上三篇文章，并使用它们的模式
建议你下次来实现一个需要高度的 API 时
可配置。

Below is an example of what this might look like for query building:

以下是查询构建的示例：

```
posts := make([]*Post, 0)

db := sqlx.Open("postgres", "...")

q := Select(
    Columns("*"),
    Table("posts"),
)

err := db.Select(&posts, q.Build(), q.Args()...)

```

A naive example, I know. But let's take a look at how we might implement an API
like this so that it can be used for query building. First, we should implement
a query struct to keep track of the query's state whilst it's being built.

一个天真的例子，我知道。但是让我们来看看我们如何实现 API
像这样，以便它可以用于查询构建。首先，我们应该实施
一个查询结构，用于在构建时跟踪查询的状态。

```
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

The above struct will keep track of the statement we're building, whether it's
`SELECT`, `UPDATE`, `INSERT`, or `DELETE`, table being operated on, the columns
we're working with, and the arguments that will be passed to the final query.
To keep this simple, let's focus on implementing the `SELECT` statement for the
query builder. 

上面的结构将跟踪我们正在构建的语句，无论是
`SELECT`、`UPDATE`、`INSERT` 或 `DELETE`，正在操作的表，列
我们正在使用，以及将传递给最终查询的参数。
为了简单起见，让我们专注于为
查询生成器。

Next, we need to define a type that can be used for modifying the query we're
building. This is the type that would be passed numerous times as a first class
function. Each time this function is called, it should return the newly modified
query, if applicable.

接下来，我们需要定义一个可用于修改我们正在查询的类型
建筑。这是将作为第一类多次通过的类型
功能。每次调用这个函数时，它应该返回新修改的
查询（如果适用）。

```
type Option func(q Query) Query

```

We can now implement the first part of the builder, the `Select` function. This
will begin building a query for the `SELECT` statement we want to build up.

我们现在可以实现构建器的第一部分，`Select` 函数。这个
将开始为我们想要构建的 `SELECT` 语句构建查询。

```
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

You should now be able to see how everything is slowly coming together, and how
the `UPDATE`, `INSERT`, and `DELETE` statements could be trivially implemented
too. Without actually implementing some options to pass to `Select`, the above
function is fairly useless, so let's do that.

您现在应该能够看到一切是如何慢慢融合在一起的，以及
`UPDATE`、`INSERT` 和 `DELETE` 语句可以简单地实现
也。没有实际实现一些选项来传递给`Select`，上面的
函数相当无用，所以让我们这样做。

```
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

As you can see, we implement these first class functions in a way so that they
return the underlying option function that will be called. It would be typically
expected for the option function to modify the query passed to it, and for a
copy to be returned.

如您所见，我们以某种方式实现了这些一流的功能，以便它们
返回将被调用的底层期权函数。这通常是
期望选项函数修改传递给它的查询，以及
要退回的副本。

For this to be useful for our use case of building complex queries, we ought to
implement the ability to add `WHERE` clauses to our query. This will require
having to keep track of the various `WHERE` clauses in the query too.

为了这对我们构建复杂查询的用例有用，我们应该
实现向我们的查询添加“WHERE”子句的能力。这将需要
还必须跟踪查询中的各种“WHERE”子句。

```
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

We define a custom type for a `WHERE` clause, and add a `wheres` property to the
original `Query` struct. Let's implement two types of `WHERE` clauses for our
needs, the first being `WHERE LIKE`, and the other being `WHERE >`.

我们为 `WHERE` 子句定义了一个自定义类型，并将 `wheres` 属性添加到
原始`查询`结构。让我们为我们的代码实现两种类型的“WHERE”子句
需要，第一个是`WHERE LIKE`，另一个是`WHERE >`。

```
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

When handling the addition of a `WHERE` clause to a query, we appropriately
handle the bindvar syntax for the underlying SQL driver, Postgres in this case,
and store the actual value itself in the `args` slice on the query.

在处理向查询添加`WHERE` 子句时，我们适当地
处理底层 SQL 驱动程序的 bindvar 语法，在本例中为 Postgres，
并将实际值本身存储在查询的 `args` 切片中。

So, with what little we have implemented we should be able to achieve what we
want in an idiomatic way.

因此，通过我们实施的很少，我们应该能够实现我们的目标
想要以惯用的方式。

```
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

Slightly better, but still not great. However, we can extend the functionality
to get what we want. So, let's implement some functions that will return options
for our specific needs.

稍微好一点，但仍然不是很好。但是，我们可以扩展功能
得到我们想要的。所以，让我们实现一些将返回选项的函数
满足我们的特定需求。

```
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

With the above two functions implemented we can now cleanly build up a somewhat
complex query for our use case. Both of these functions will only modify the
query, if the values passed to them are considered correct.

实现上述两个功能后，我们现在可以干净地构建一个
我们用例的复杂查询。这两个函数只会修改
查询，如果传递给它们的值被认为是正确的。

```
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

I find this to be a rather idiomatic way of building up complex queries in Go.
Now, of course you've made it this far in the post, and must be wondering,
"That's good and all, but you didn't imeplement the Build(), or Args() methods".
This is true, to an extent. In the interest of not wanting to prolong this post
any further than needed, I didn't bother. So, if you are interested in some of
the ideas presented here, take a look at the [code](https://github.com/andrewpillar/query),
I submitted to GitHub. It's nothing to rigorous, and doesn't cover everything a
query builder would need, it's lacking `JOIN`, for example and supports only
the Postgres bindvar.

我发现这是在 Go 中构建复杂查询的一种相当惯用的方式。
现在，当然，您已经在帖子中做到了这一点，并且一定想知道，
“这很好，但你没有实现 Build() 或 Args() 方法”。
在某种程度上，这是真的。为了不想延长这篇文章
任何超出需要的，我没有打扰。所以，如果你对一些感兴趣
这里提出的想法，看看[代码](https://github.com/andrewpillar/query)，
我提交给了 GitHub。没什么好严格的，也没有涵盖所有
查询构建器需要，例如它缺少`JOIN`，并且仅支持
Postgres 绑定变量。

If you have any disagreements with what I have said in this post, or would like
to discuss this further, then please reach out to me at
[me@andrewpillar.com](mailto:me@andrewpillar.com). 

如果你对我在这篇文章中所说的有任何异议，或者想要
进一步讨论这个，然后请与我联系
[me@andrewpillar.com](mailto:me@andrewpillar.com)。

