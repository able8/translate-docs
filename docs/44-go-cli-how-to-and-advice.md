# [Writing Go CLIs With Just Enough Architecture](https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/)

Thursday, June 4, 2020

As someone who has written [quite a few](https://blog.carlmjohnson.net/post/2018/go-cli-tools/) command line applications in Go, I was interested the other day when I saw [Diving into Go by building a CLI application](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) by [Eryb](https://eryb.space) on social media. In the post, the author describes writing a simple application to fetch comics from the [XKCD API](https://xkcd.com/json.html) and display the results. It looks like this:

```
$ go-grab-xkcd --help

Usage of go-grab-xkcd:
  -n int
        Comic number to fetch (default latest)
  -o string
        Print output in format: text/json (default "text")
  -s    Save image to current directory
  -t int
        Client timeout in seconds (default 30)
```

I came away a little disappointed though because I felt  like the final result was both a little undercooked and a little  overcooked: undercooked in that it didn’t handle errors robustly enough  and overcooked in that it created some abstractions speculatively in a  way that I felt were [unlikely to pay off in the long run](https://en.wikipedia.org/wiki/You_aren't_gonna_need_it).

Naturally, one thing lead to another, and I [forked and rewrote](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) the demo app myself to demonstrate what I consider to be **just enough** architecture for a Go command line app.

------

Before I go too much further with this post, let me say that I don’t mean to  be overly negative. For such a small app, none of what follows  particularly matters. This app is simple enough that you could write it  in Bash if you wanted to. The point was to write an command line  interface for fun and get some experience doing it. So, with that in  mind, go read [Eryb’s post](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) first and then come back and let me share how I would approach writing  the same CLI in Go with the experience I’ve gained by writing so many  CLIs and having to extend them on and off over a period of years.

My core principle is that while there are many possible architectures for an app—[MVC](https://en.wikipedia.org/wiki/Model–view–controller), [hexagonal](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)), etc—there are three layers that you almost always need and never regret writing. You want one layer to **handle user input** and get it into a normalized form. You want one layer to **do your actual task**. And you want one layer to **handle formatting and output** to the user. Those are the three layers you always need. Other layers  you may or may not need and can add it when it becomes clear that they  might help. For a Go command line interface, this means the most  important thing is to separate flag parsing stuff from execution, and  everything else is not a big deal to let evolve over time. The main  challenge is avoiding create abstractions that don’t actually pay for  themselves in setup time versus time saved in writing future extensions  to your program.

The first thing I noticed looking at [the source to the demo CLI](https://github.com/carlmjohnson/go-grab-xkcd/tree/v0.0.0) is that it was divided into three packages: main (necessarily), model, and client. This is one package **too few** and two packages **too many**! A Go CLI should start with two packages: main (in which func main has  only one line) and everything else (in this case, call it package  grabxkcd).

Before I explain about the single line `main` function, let’s look at a common (anti-)pattern for Go CLIs (which the  demo app fortunately doesn’t apply), in which you check for errors with a helper function like this:

```go
func check(err error) { // or try() or must() or die()
    if err != nil {
        log.Fatal(err) // or panic or os.Exit
    }
}

func main() {
    v, err := doSomething()
    check(err)
    // …
}
```

I myself have written this more than a few times,  but for any serious CLI, this pattern breaks down quickly because it  allows for only one response to errors, immediately aborting. Often that is a good beginning for a CLI, but as it grows you find that you need  to handle some errors by retrying or ignoring or user notification, but  because your functions all return a single value instead of value plus  error, it becomes a ton of work to add proper error handling in after  the fact.

Okay, so if we shouldn’t use a `check` function, why should we use a single line `main` function? The original demo app continues even after it finds an error  in fetching the comic. This was because it couldn’t return err in main  and it also didn’t use the `check` func pattern. It also  didn’t exit with a non-zero code on error, which matters if you want to  use your CLI in a scripting pipeline. The single line main func solves  these problems.

I took the idea of a single line main function from a [blog post by Nate Finch](https://npf.io/2016/10/reusable-commands/). He wrote:

I think there’s only approximately one right answer to “what’s in your package main?” and that’s this:

```go
// command main documentation here.
package main

import (
    "os"

    "github.com/you/proj/cli"
)
func main{
    os.Exit(cli.Run())
}
```

You can read [his post](https://npf.io/2016/10/reusable-commands/) for more explanation of why, but for me the argument comes down to that it makes your program more testable and more extensible without being  much more work. Like a `check` function, the single line main lets you easily handle errors in execution by creating a clear pathway  between errors and program termination, but unlike a `check` function, the single line main does so without closing off the path for future enhancements.

The demo CLI uses the standard Go flag package. The flag package has many  detractors because it is (like Go itself) quite idiosyncratic, but I  find that it is mostly great. The one thing I don’t like about the flag  package is that [the documentation](https://pkg.go.dev/flag?tab=doc#pkg-overview) suggests using it by setting global variables and using `init` functions, like

```go
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

In his post [A Theory of Modern Go](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html), Peter Bourgon proposes the following guidelines:

> - No package level variables
> - No func init

He does so because

> Package-global objects can encode state and/or behavior that is hidden from external  callers. Code that calls on those globals can have surprising side  effects, which subverts the reader’s ability to understand and mentally  model the program.

I agree with this whole heartedly, so I think that when you use the flag package, it is important not to  use it with global variables in your package. Using global variables for flags inside of package main leads to spooky action at a distance very  quickly. Where did this variable come from? Why is it set this way? When you use global flags, these confusions crop up all the time. One thing I like about the original demo is that it didn’t use globals, but I think it could be refined even a little bit more.

One more bit of background before I get into the nitty-gritty of my rewrite. In [How I write HTTP services after eight years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html), Mat Ryer explains,

> All of my components have a single server structure that usually ends up looking something like this:
>
> ```go
> type server struct {
>     db     *someDatabase
>     router *someRouter
>     email  EmailSender
> }
> ```

He then creates his routes and  handlers as methods hanging off the server type. I use the same  principle when writing a CLI. As the flag package processes my  arguments, I use them to create an `appEnv` struct that holds the options and controls the execution environment.

Putting these pieces together, my recently created CLIs all follow this  pattern: a single line main function calls a function (usually named `app.CLI` and in a separate package) that first uses the flag package to parse  user input and then kicks off actual execution as a method on an `appEnv` struct. Here is the code for [my version](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) of the go-grab-xkcd app:

```go
// In main.go / package main
func main() {
    os.Exit(grabxkcd.CLI(os.Args[1:]))
}

// In grabxkcd/grabxkcd.go
func CLI(args []string) int {
    var app appEnv
    err := app.fromArgs(args)
    if err != nil {
        return 2
    }
    if err = app.run(); err != nil {
        fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
        return 1
    }
    return 0
}
```

Next, as we wrap up handling user input, let’s take a look at the `appEnv` type and `fromArgs` method:

```go
// In grabxkcd/grabxkcd.go

type appEnv struct {
    hc         http.Client
    comicNo    int
    saveImage  bool
    outputJSON bool
}

func (app *appEnv) fromArgs(args []string) error {
    // Shallow copy of default client
    app.hc = *http.DefaultClient
    fl := flag.NewFlagSet("xkcd-grab", flag.ContinueOnError)
    fl.IntVar(
        &app.comicNo, "n", LatestComic, "Comic number to fetch (default latest)",
    )
    fl.DurationVar(&app.hc.Timeout, "t", 30*time.Second, "Client timeout")
    fl.BoolVar(
        &app.saveImage, "s", false, "Save image to current directory",
    )
    outputType := fl.String(
        "o", "text", "Print output in format: text/json",
    )
    if err := fl.Parse(args); err != nil {
        return err
    }
    if *outputType != "text" && *outputType != "json" {
        fmt.Fprintf(os.Stderr, "got bad output type: %q\n", *outputType)
        fl.Usage()
        return flag.ErrHelp
    }
    app.outputJSON = *outputType == "json"
    return nil
}
```

This is pretty similar to the original demo CLI, but I’ve made two tweaks beyond just assigning to `appEnv` members. Here is the original code for comparison:

```go
    comicNo := flag.Int(
        "n", int(client.LatestComic), "Comic number to fetch (default latest)",
    )
    clientTimeout := flag.Int64(
        "t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
    )
    saveImage := flag.Bool(
        "s", false, "Save image to current directory",
    )
    outputType := flag.String(
        "o", "text", "Print output in format: text/json",
    )
```

The original code takes a timeout of `int64` seconds (`int32` is 136 years, which ought to be enough IMO), but Go has a built in `time.Duration` type. In other languages like JavaScript and Python it’s common to use  integers for durations and just have a convention about whether the  duration is in seconds or milliseconds or whatever, but in Go, you  should always use `time.Duration`. A nice trick is to set the timeout directly on your `http.Client` by using `flag.DurationVar` and passing in a pointer to `myclient.Timeout`.

Second, outputType is used to make a binary choice between text and json  output. In my code, I check for that and return an initialization error  if someone sets `-o unsupported`.

------

That wraps up the user input layer! Whew. Now let’s look at **actual task execution**.

The original demo had separate client and model packages, but if you look at the client, it imports returns a `model.Comic` object, so the client package cannot be used without the model package:

```go
func NewXKCDClient() *XKCDClient
func (hc *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error)
func (hc *XKCDClient) SaveToDisk(url, savePath string) error
func (hc *XKCDClient) SetTimeout(d time.Duration)
```

If you look at the private members of the `client.XKCDClient` (the name repeats “client” which is often [a bad code smell](https://blog.golang.org/package-names) in Go!), you can see that it’s actually just a `*http.Client` plus a URL builder for the XKCD API:

```go
type XKCDClient struct {
    client  *http.Client
    baseURL string
}

func (hc *XKCDClient) buildURL(n ComicNumber) string {
    var finalURL string
    if n == LatestComic {
        finalURL = fmt.Sprintf("%s/info.0.json", hc.baseURL)
    } else {
        finalURL = fmt.Sprintf("%s/%d/info.0.json", hc.baseURL, n)
    }
    return finalURL
}
```

Coming from other languages, there is a temptation to build clients like this that “encapsulate” the details of the API  by, for example, hiding the base URL. The idea is that you may want to  point the base URL at localhost or a staging server for testing. I think that in Go this is basically a mistake. The `http.Client` type has a member called `Transport` of type `http.RoundTripper` interface. By changing the client transport, you can make an `http.Client` do whatever you want in testing—or even production. For example, you might create a `RoundTripper` that adds auth headers (Google does this for [their APIs](https://pkg.go.dev/cloud.google.com/go?tab=doc)) or does caching or reads test responses from a test file. Controlling the `http.Client` is extremely powerful, so any Go package designed to help users with an API should let them supply their own clients. Once you do that, there’s no longer any point in monkeying with base URLs.

Since the client and models packages cannot be used apart from one another, they should  be combined. Because the client is just a wrapper around `http.Client` and a private URL builder, the URL builder should be made public and anyone can use their own `http.Client`. Once we’ve broken things down to that point, there’s not really enough  for a separate package, hence my joke that the original app had one  package too few and two too many. As Dave Cheney argued, [consider fewer, larger packages](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_consider_fewer_larger_packages). If we do want to separate the API stuff from the command line stuff  just for organization purposes, we can put it into a separate file in  the same package. My version looks like this:

```go
// In grabxkcd/api.go

func BuildURL(comicNumber int) string {
    if comicNumber == LatestComic {
        return "https://xkcd.com/info.0.json"
    }
    return fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
}

// APIResponse returned by the XKCD API
type APIResponse struct {
    Month       string `json:"month"`
    Number      int    `json:"num"`
    Link        string `json:"link"`
    Year        string `json:"year"`
    News        string `json:"news"`
    SafeTitle   string `json:"safe_title"`
    Transcript  string `json:"transcript"`
    Description string `json:"alt"`
    Image       string `json:"img"`
    Title       string `json:"title"`
    Day         string `json:"day"`
}
```

Back in `grabxkcd/grabxkcd.go`, I add `appEnv` methods to replace what we lost from the client:

```go
func (app *appEnv) fetchJSON(url string, data interface{}) error
func (app *appEnv) fetchAndSave(url, destPath string) error
```

The original demo app’s `Fetch` and `SaveToDisk` methods were tied into the XKCD API and couldn’t really be reused in other situations. By writing generic `fetchJSON` and `fetchAndSave` helpers that work with any JSON response or GET-able URL, it’s more  likely that we’ll be able to reuse our code, if not in this project then by copy-pasting it into a future project. When working with JSON in  particular, it’s often better to write functions/methods that take a  dynamic type (`interface{}`) instead of a static type (e.g. `APIResponse`). Accepting a dynamic type allows you to reuse your code in more  situations, and using a static type doesn’t actually add any practical  type safety, since the underlying calls to `json.Marshal` or whatever will still be dynamic themselves.

There are two execution details that the demo app hid away inside of other  functions that I want to bring to top level in mine. The [XKCD API](https://xkcd.com/json.html) is okay, but it doesn’t quite return the information we want.  Specifically, instead of having an ISO-encoded date field, it has three  fields for day, month, and year, and instead of telling us the name of  its images, it gives us the whole URL of an image and the filename is  the last component of the path. Let’s fix those issues by adding methods onto the `APIResponse` type:

```go
// Date returns a *time.Time based on the API strings (or nil if the response is malformed).
func (ar APIResponse) Date() *time.Time {
    t, err := time.Parse(
        "2006-1-2",
        fmt.Sprintf("%s-%s-%s", ar.Year, ar.Month, ar.Day),
    )
    if err != nil {
        return nil
    }
    return &t
}

func (ar APIResponse) Filename() string {
    return path.Base(ar.Image)
}
```

Notice that by explicitly parsing a date out of  the XKCD API, we were faced with a choice about error handling. Should  we return an error here? Should we return a blank object? Returning a  nil is a compromise choice. It means that we think it is unlikely that  the XKCD API will return a bad date, but if it does, we’re willing to  accept a runtime panic if we forget to check that the date is not nil  first. It’s not the most robust error handling strategy, but it’s a  middle ground for a validation error that we expect to never happen.

------

We’ve looked at looked handling user input and actual execution. Now it’s time for **output**.

The original demo app has a `Comic` model type for output, which is a good practice, but I find the name `models.Comic` less than clear about its intent (doesn’t the API return a comic?). Let’s call it `grabxkcd.Output` instead.

The original demo app also uses the same type for both JSON output and  plain text output, but I think that’s a little bit premature. Right now  we want these types to use the same underlying values, but there’s no  reason in principle they might not diverge in the future. Let’s put off  the decision for now and just have separate functions for JSON output  and pretty printing. If we want to unify them in the future (for  example, by replacing the pretty printer with a `template.Template` output that consumes `grabxkcd.Output`), it shouldn’t be too hard.

Here’s what our `appEnv.run` method looks like with these changes:

```go
func (app *appEnv) run() error {
    u := BuildURL(app.comicNo)

    var resp APIResponse
    if err := app.fetchJSON(u, &resp); err != nil {
        return err
    }
    if resp.Date() == nil {
        return fmt.Errorf("could not parse date of comic: %q/%q/%q",
            resp.Year, resp.Month, resp.Day)
    }
    if app.saveImage {
        if err := app.fetchAndSave(resp.Image, resp.Filename()); err != nil {
            return err
        }
        fmt.Fprintf(os.Stderr, "Saved: %q\n", resp.Filename())
    }
    if app.outputJSON {
        return printJSON(resp)
    }

    return prettyPrint(resp)
}
```

------

With these changes in place, our [public API for package grabxkcd](https://godoc.org/github.com/carlmjohnson/go-grab-xkcd/grabxkcd) looks like this:

```go
func CLI(args []string) int
func BuildURL(comicNumber int) string
type APIResponse
func (ar APIResponse) Date() *time.Time
func (ar APIResponse) Filename() string
type Output
```

What I like about this is I feel it balances our  convenience as a developer when writing the code and the convenience of  any potential users of our package. This is a simple CLI, so probably no one will ever use this code but us. On the one hand, we found it more  convenient for our own development purposes to break out a second  package instead of just putting everything into package main. But on the other hand, we haven’t invested much in creating a big public API full  of features and interfaces that we’re just speculating might be useful  for someone, someday. This is the code we needed, and if it ever helps  someone else, great, and if not, no big deal.

[My final commit](https://github.com/carlmjohnson/go-grab-xkcd/commit/d5d51722ef95a5ebdf49c453cd7f00ed4a9e3910) has **197 additions** and **197 deletions**, meaning that for all my writing about different abstractions we may or may not want to use, in the end, **the code is the exact same length** as written by me as it was when written by Eryb. (Okay, yes, I may have noticed I was three lines away and squeezed a little to make it the exact same.) My code actually handles several error conditions that were overlooked  by the original demo, so we’re actually doing a little bit more work  than the original app did. We’ve added just enough architecture to make  our CLI more robust and extensible without actually doing any more work  than normal.

My code is designed to be easier to test… but I  haven’t actually written any tests yet. Why? Because for a simple CLI if you’re not following [TDD](https://en.wikipedia.org/wiki/Test-driven_development), you may not want to take the time to write tests. The idea is to leave  room open so that if you want or need to add tests later, there are  clear and easy processes for doing so.

The point of all of this,  again, is not that my method is the only way to write a CLI in Go. For  simple CLIs, anything will work, including having a single main.go file  with a bunch of global variables and a `check` function. I  argue however that unlike creating a hexagonal architecture with many  layers of interfaces just in case you might need them in the future, my  method is **no more work** than the simple method while being **much better for extending in the future**.

------

If you like the way I describe writing Go CLIs but want a template to  start off of instead of writing the boilerplate yourself, check out [go-cli](https://github.com/carlmjohnson/go-cli), a [Github Template Repo](https://github.com/carlmjohnson/go-cli/generate) I have created that acts as a `cat` clone. It follows the patterns described in this post and is pretty easy to adapt into new tools.

------

## Bash Double Bonus

The [original post](https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html) ends with a trick for downloading multiple comics serially in Bash, so  here’s a bonus trick for downloading the comics in parallel and waiting  for them all to finish:

```bash
$ for i in {1..10}; do ./go-grab-xkcd -n $i -s & done; wait
```

I hope this helped you think about how to structure your CLIs, and you have fun experimenting with finding your own patterns*!