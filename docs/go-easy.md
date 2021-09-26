# Go is not an easy language

Written on 22 Feb 2021

Discussions:  [Lobsters](https://lobste.rs/s/ee6nsc/go_is_not_easy_language),  [Hacker News](https://news.ycombinator.com/item?id=26220693),  [/r/golang](https://www.reddit.com/r/golang/comments/lpeafy/go_is_not_an_easy_language/),  [/r/programming](https://www.reddit.com/r/golang/comments/lpo6zh/go_is_not_an_easy_language/)

Go is not an easy programming language. It *is* simple in many ways: the syntax is simple, most of the semantics are simple. But a language is more than just syntax; it’s about doing useful *stuff*. And doing useful stuff is not always easy in Go.

Turns out that combining all those simple features in a way to do something useful can be tricky. How do you remove an item from an array in Ruby? `list.delete_at(i)`. And remove entries by value? `list.delete(value)`. Pretty easy, yeah?

In Go it’s … less easy; to remove the index `i` you need to do:

```
list = append(list[:i], list[i+1:]...)
```

And to remove the value `v` you’ll need to use a loop:

```
n := 0
for _, l := range list {
    if l != v {
        list[n] = l
        n++
    }
}
list = list[:n]
```

Is this unacceptably hard? Not really; I think most programmers can figure out what the above does even without prior Go experience. But it’s not exactly *easy* either. I’m usually lazy and copy these kind of things from the [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks) page because I want to focus on actually solving the problem at hand, rather than plumbing like this.

It’s also easy to get it (subtly) wrong or suboptimal, especially for less experienced programmers. For example compare the above to copying to a new array and copying to a new pre-allocated array (`make([]string, 0, len(list))`):

```
InPlace             116 ns/op      0 B/op   0 allocs/op
NewArrayPreAlloc    525 ns/op    896 B/op   1 allocs/op
NewArray           1529 ns/op   2040 B/op   8 allocs/op
```

While 1529ns is still plenty fast enough for many use cases and isn’t something to excessively worry about, there are plenty of cases where these things *do* matter and having the guarantee to always use the best possible algorithm with `list.delete(value)` has some value.

------

Goroutines are another good example. “Look how easy it is to start a goroutine! Just add `go` and you’re done!” Well, yes; you’re done until you have five million of those running at the same time and then you’re left wondering where all your memory went, and it’s not hard to “leak” goroutines by accident either.

There are a number of patterns to limit the number of goroutines, and none of them are exactly easy. A simple example might be something like:

```
var (
    jobs    = 20                 // Run 20 jobs in total.
    running = make(chan bool, 3) // Limit concurrent jobs to 3.
    wg      sync.WaitGroup       // Keep track of which jobs are finished.
)

wg.Add(jobs)
for i := 1; i <= jobs; i++ {
    running <- true // Fill running; this will block and wait if it's already full.

    // Start a job.
    go func(i int) {
        defer func() {
            <-running // Drain running so new jobs can be added.
            wg.Done() // Signal that this job is done.
        }()

        // "do work"
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }(i)
}

wg.Wait() // Wait until all jobs are done.
fmt.Println("done")
```

There’s a reason I annotated this with some comments: for people not intimately familiar with Go this may take some effort to understand. This also won’t ensure that the numbers are printed in order (which may or may not be a requirement).

Go’s concurrency primitives may be simple and easy to use, but combining them to solve common real-world scenarios is a lot less simple. The original version of the above example [was actually incorrect](https://lobste.rs/s/ee6nsc/go_is_not_easy_language#c_gdnw5e) 😅

------

In [Simple Made Easy](https://www.infoq.com/presentations/Simple-Made-Easy/) Rich Hickey argues that we shouldn’t confuse “simple” with “it’s easy to write”: just because you can do something useful in one or two lines doesn’t mean the underlying concepts – and therefore the entire program – are “simple” as in “simple to understand”.

I feel there is some wisdom in this; in most cases we shouldn’t sacrifice “simple” for “easy”, but that doesn’t mean we can’t think at all about how to make things easier. Just because concepts are simple doesn’t mean they’re easy to use, can’t be misused, or can’t be used in ways that lead to (subtle) bugs. Pushing Hickey’s argument to the extreme we’d end up with something like [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) and that would of course be silly.

Ideally a language should reduce the cognitive load required to reason about its behaviour; there are many ways to increase this cognitive load: complex intertwined language features is one of them, and getting “distracted” by implementing fairly basic things from those simple concepts is another: it’s another block of code I need to reason about. While I’m not overly concerned about code formatting or syntax choices, I do think it can matter to reduce this cognitive load when reading code.

The lack of generics probably plays some part here; implementing a `slices` package which does these kind of things in a generic way is hard right now. Generics make this possible and also makes things more complex (more language features are used), but they also make things easier and, arguably, less complex on other fronts.[[1\]](https://www.arp242.net/go-easy.html#fn:g)

------

Are these insurmountable problems? No. I still use (and like) Go after all. But I also don’t think that Go is a language that you “could pick up in ~5-10 minutes”, which was the comment that prompted this post; a sentiment I’ve seen expressed many times, although usually with less extreme timeframes (“1-2 days”, “1 week”).

As a corollary to all of the above; learning the language isn’t just about learning the syntax to write your `if`s and `for`s; it’s about learning a way of thinking. I’ve seen many people coming from Python or C♯ try to shoehorn concepts or patterns from those languages in Go. Common ones include using struct embedding as inheritance, panics as exceptions, “pseudo-dynamic programming” with interface{}, and so forth. It rarely ends well, if ever.

I did this as well when I was writing my first Go program; it’s only natural. And when I started as a Ruby programmer I tried to write Python code in Ruby (although this works a bit better as the languages are more similar, but there are still plenty of odd things you can do such as using `for` loops).

This is why I don’t like it when people get redirected to the Tour of Go to “learn the language”, as it just teaches basic syntax and little more. It’s nice as a little, well, *tour* to get a bit of a feel of the language and see how it roughly works and what it can roughly do, but it’s ill-suited to actually learn the language.

**Footnotes**

1. Contrary to popular belief the [Go team was never “against” generics](https://research.swtch.com/generic);  I’ve seen many comments to the effect of “the Go team doesn’t think  generics are useful”, but this was never the case. [↩](https://www.arp242.net/go-easy.html#fnref:g)

**Feedback**

Contact me at 				[martin@arp242.net](mailto:martin@arp242.net), 				[GitHub](https://github.com/arp242/arp242.net/issues/new), or 				[@arp242_martin](https://twitter.com/arp242_martin) 				for feedback, questions, etc.

**Other Go posts**

- 21 Nov 2019 [Go’s features of last resort](https://www.arp242.net/go-last-resort.html)
- 10 Dec 2020 [Bitmasks for nicer APIs](https://www.arp242.net/bitmask.html)