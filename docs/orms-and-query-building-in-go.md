# ORMs and Query Building in Go

_Posted on Sat 13 2019 to [Programming](http://andrewpillar.com/programming)_ _Updated at 22:30 on Mon 23 Mar 2020_

Recently, I have been looking into various solutions for interacting with
databases with ease in Go. My go to library for database work in Go is
[sqlx](https://github.com/jmoiron/sqlx), this makes unmarshalling the data from
the database into structs a cinch. You write out your SQL query, tag your
structs using the `db` tag, and let [sqlx](https://github.com/jmoiron/sqlx)
handle the rest. However, the main problem I have encountered, was with
idiomatic query building. This led me to investigate this problem, and jot down
some of my thoughts in this post.

**TL;DR** First class functions are an idiomatic way of doing SQL query building
in Go. Check out the repository containing some example code I wrote testing
this out: [https://github.com/andrewpillar/query](https://github.com/andrewpillar/query).

## GORM, Layered Complexity, and the Active Record Pattern

Most people who dip their toe into database work in Go, will most likely be
pointed towords [gorm](https://gorm.io) for working with databases. It is a
fairly fully featured ORM, that supports migrations, relations, transactions,
and much more. For those who have worked with ActiveRecord, or Eloquent, GORM's
usage would be some what familiar to you.

I have used GORM briefly before, and for simple CRUD based applications this is
fine. However, when it comes to more layered complexity, I find that it falls
short. Assume we are building a blogging application, and we allow users to
search for posts via the `search` query string in a URL. If this is present, we
want to constrain the query with a `WHERE title LIKE`, otherwise we do not.

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

As demonstrated above, I find GORM's biggest drawback is how cumbersone it can
be to do layered complexity. More often than not when writing SQL queries, you
will want this. Trying to determine if you want to add a `WHERE` clause to a
query based off of some user input, or how you should order the records.

I believe this comes down to one thing, and I made a
[comment](https://news.ycombinator.com/item?id=19851753) about this some time
ago on HN:

> Personally I think an active record style ORM for Go like gorm is a poor fit
> for a language that doesn't come across as inherently OOP. Going through some
> of the documentation for gorm, it seems to rely heavily on method chaining which
> for Go seems wrong considering how errors are handled in that language. In my
> opinion, an ORM should be as idiomatic to the language as possible.

This comment was made on a submission of the blog post
[To ORM or not to ORM](https://eli.thegreenplace.net/2019/to-orm-or-not-to-orm/),
which I highly recommend you read. The author of the post approaches the same
conclusion about GORM that I did.

## Idiomatic Query Building in Go

The `database/sql` package in the standard library is great for interacting with
databases. And [sqlx](https://github.com/jmoiron/sqlx) is a fine extension on
top of that for handling the return of data. However, this still doesn't fully
solve the problem at hand. How can we effectively build complex queries
programmatically that is idiomatic to Go. Assume we were using
[sqlx](https://github.com/jmoiron/sqlx) for the same query above, what would
that look like?

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

Let's take a look at [squirrel](https://github.com/masterminds/squirrel). This
library offers the ability to build up queries, and execute them directly in a
way that I find rather idiomatic to Go. Here though, we will only be focussing
on the query building aspect.

With [squirrel](https://github.com/masterminds/squirrel), we can implement our
above logic like so.

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

An API for query building in Go should fulfill both of these needs:

- Idiomacy
- Extensibility

How can this be achieved with Go?

## First Class Functions for Query Building

Dave Cheney has written two blog posts about first class functions, based off of
a post made by Rob Pike about the same topic. For those interested they can be
found here:

- [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Do not fear the first class functions](https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions)

I'd highly recommend reading the above three posts, and using the pattern they
suggest when you next come to implement an API that needs to be highly
configurable.

Below is an example of what this might look like for query building:

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

Next, we need to define a type that can be used for modifying the query we're
building. This is the type that would be passed numerous times as a first class
function. Each time this function is called, it should return the newly modified
query, if applicable.

```
type Option func(q Query) Query

```

We can now implement the first part of the builder, the `Select` function. This
will begin building a query for the `SELECT` statement we want to build up.

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

For this to be useful for our use case of building complex queries, we ought to
implement the ability to add `WHERE` clauses to our query. This will require
having to keep track of the various `WHERE` clauses in the query too.

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

So, with what little we have implemented we should be able to achieve what we
want in an idiomatic way.

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

If you have any disagreements with what I have said in this post, or would like
to discuss this further, then please reach out to me at
[me@andrewpillar.com](mailto:me@andrewpillar.com).
