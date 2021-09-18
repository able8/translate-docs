# Better Error Handling, in Go

Tuesday March 26th 2019

It’s no secret that we’re big golang fans at bet365. It has allowed us to create exciting new products, like Bet Builder and our custom sports search engine. Within the very near future, the majority of our sports website will be powered by _Go_.

We’ve used this move in technology as an opportunity to review what simplicity means within our vast codebase, and embrace the idiomatic ‘Go way’.

Our internal standards document bares signage to “leave your OO baggage at the door”.

And we did; Our code organization has become more literal, with less obfuscating OO patterns, clearer logic paths and a focus on reducing ‘code golf’ for maintainability. Go promotes this kind of behavior and generally makes it obvious when it’s punishing you for not abiding by its rules.

## I’ll get to the point

One great example of this simplicity is the way it handles errors at runtime; Rather than adding a bespoke language construct for dealing with the raising and propagation of errors throughout your call stacks, go chose to simply make use of its ability to return more than one value from a function. This is great because there is no confusion when you’re invoking a function or method as to whether or not it is your responsibility to handle what happens when something goes wrong. If you see an `error` is part of that functions return signature, then it’s on you to deal with it.

Now, whilst this does make your code declaratively more obvious, one thing we noticed was if you’re calling a few functions which return errors it’s quite easy for things to get a little _too obvious_ that you’re doing a bit of error handling by falling into the pattern I’m about to demonstrate.

You may have seen this pattern occurring yourself, you may have solved this problem yourself in a different way, you may be totally fine with this code and not see it as a problem at all, and that’s okay.

## Our problem, an example

Let’s take a look at a somewhat contrived example of what I’m talking about. We’ll use a typical HTTP handler for some context

``` js

func myHandler(w http.Response, r *http.Request) {

    err := validateRequest(r)
    if err != nil {
        log.Printf("error validating request to myHandler - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    user, err := getUserFromRequest(r)
    if err != nil {
        log.Printf("error getting user from request in myHandler - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    dataset, err := db.GetUserData(user)
    if err != nil {
        log.Printf("error retrieving user data in myHandler - err: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    buffer := newBuffer()
    err := serialize.UserData(dataset, &buffer)
    if err != nil {
        log.Printf("error serializing user data in myHandler - err %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    err := buffer.WriteTo(w);
    if err != nil {
        log.Printf("error writing buffer to response in myHandler - err %v", err)
        return
    }
}

```

Okay, so a nice straightforward HTTP handler. Let’s quickly break down what tasks it was performing for us.

- Validate the request
- Grab a user from the request
- Fetch a user dataset from the database
- Serialize the dataset into a buffer
- Write the buffer to the HTTP response

So we’ve done those five things, fine, but take a look at that code again – it doesn’t feel like we just did five basic things there, it feels like a lot more was happening due to the noise created by the error handling after each action we took.

The first thing to consider here is the amount of additional context we’re wrapping into each one of these errors for logging purposes. The messages generally start off with `"there was an error doing thing"` – which seems like an obvious thing to pad out your error message with, but it could actually be redundant when you consider the function it originated from could’ve (should’ve) added enough context about where the error originated from itself. We found we were able to remove a lot of `return fmt.Errorf("there was a db error %v", err)` statements that were simply adding zero additional information.

Of course, that doesn’t necessarily mean the error was raised in the function we called. The point is wherever we have an error ‘origin site’ within our own codebase, we should take the opportunity to decorate the error with enough context that it doesn’t need wrapping multiple times as we pass it back up the stack. This could be parameter information, generated URLs or other context useful to someone diagnosing an issue later on. We can then simply bare-return errors up the stack up until the point we can handle them – in our case, exit safely and log the error.

So how do we remove the need to wrap these errors multiple times during the execution path of our handler but still maintain the same level of information? After all, the information we wrapped it with wasn’t useless, it also stated `in myHandler` as the error ‘destination site’, which again could be vital in a diagnosis investigation by an engineer looking at the issue later on.

After trying a few different approaches, and there are many ways to do this, we took inspiration from how go 2.0 had proposed solving the problem ( [check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)) and came up with a more concise solution we can use today using what we call the error defer pattern. A halfway house, if you like.

## The solution, an example

This is the same code, but we’re going to switch out the error handling to use error defer. Let’s take a look

``` js

func myHandler(w http.Response, r *http.Request) {

    var err error
    defer func() {
        if err != nil {
            log.Printf("error in myHandler - error: %v", err)
            w.WriteHeader(http.StatusInternalServerErrror)
        }
    }()

    err = validateRequest(r)
    if err != nil { return }

    user, err := getUserFromRequest(r)
    if err != nil { return  }

    dataset, err := db.GetUserData(user)
    if err != nil { return }

    buffer := newBuffer()
    err = serialize.UserData(dataset, &buffer)
    if err != nil { return }

    err2 := buffer.WriteTo(w)
    if err2 != nil {
        log.Printf("error writing buffer to response in myHandler - error %v", err2)
        return
    }
}

```

Straight away, that feels much easier and quicker to digest. We can see the application logic hasn’t changed but has become more visible due to the error handling blocks being reduced to a more basic form. We’ve also made more use of the context added at the origin sites within the functions we’ve called.

The error we respond with is now tracked by the top level `var err error`. Each function which returns an error writes to that single declaration, then simply checks the value is not `nil` each time it can be set, before it continues to the next statement. If the value isn’t `nil` we just return. When we return, no matter what stage we do, we always drop into the `defer func` we declared at the top of the handler. From there, we check the error, wrap with the destination site information and continue to handle as previously by logging and setting an http status code.

You’ll notice the last block shows an exception to the pattern and how we can deal with scenarios where sharing an error handler is not sufficient – We failed the write, so there’s little point setting an http status code at that point, we just need to log the error, so we don’t write to `err`.

Without wanting to bang on about errors too much more, there are also other options within this approach, for example we can tag more information into custom `error` structs, then check their types in the `defer func` to give ourselves additional context or configuration whilst still keeping this handling code out of the main execution path, like so

``` js

if custom, ok := err.(CustomErrorType); ok {
     /// handle the custom type
}

```

## Conclusion

Go is a fantastic language to work with, but like any tool it can easily become unwieldy without proper care for how it is used. Error handling has been identified by the wider community as being somewhat of a bugbear, so this is how we solved that problem in a way that worked for us. We look forward to using [check/handle](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md) when it finally arrives in golang 2.0!

#### Pete G

Senior Software Architect
