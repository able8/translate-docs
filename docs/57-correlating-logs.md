# Correlating Logs

**01.10.2020 19:39** From: https://filipnikolovski.com/posts/correlating-logs/

When something goes wrong in your system, logs are crucial to finding out  exactly what’s happened. Usually, this involves following the logs as a trail of breadcrumbs that lead to the root cause of the failure. If your application is generating a lot of logs, it can become strenuous to tie everything together that reveals the failing scenario.

This  can become especially challenging in a distributed system, where one  HTTP request to your API can pass through dozens of different services, each outputting logs that have no context of the flow of the request.

In this post firstly we’ll define what is a structured log and then explore a solution on how to tie in the mutual logs in a **Go application**, using context, structured logging, and correlation ids. Although the following example  will be entirely in Go, the principle should be the same for apps  written in different programming languages.

## Structured logs

Before we talk about correlating logs, we need a way to make sense of the data that we are writing. Logs are typically just an unstructured text which makes it hard to extract useful information from them by another  machine. To be able to query or filter them by something, the logs need  to be written in a format that can be easily parsed and indexed.

Here is an example of an unstructured log record:

```go
INFO: 2009/11/10 23:00:00 192.168.0.2 Hello World! 
```

Versus a structured one:

```json
{"time": 1562212768, "host": "192.168.0.2", "level": "INFO", "message": "Hello World!"}
```

The structured log essentially contains the same  information as the unstructured one, but the key difference is that the  message is in a format that another machine can understand (it can be JSON, XML, whatever) and the  fields can be identified, indexed, and other programs such as logging  systems can analyze them so that later on we can use these fields to search and  filter results.

Let’s see this in action in a **Go app**. There are plenty of libraries for structured logging in Go, some of the more popular ones are [Zerolog](https://github.com/rs/zerolog), [Zap](https://github.com/uber-go/zap), and [Logrus](https://github.com/sirupsen/logrus).

We’ll use Zerolog as an example. From the package description, it says that  it provides “a fast and simple logger dedicated to JSON output”. It has features such as contextual fields, log levels, and passing the logger  by context, which is exactly what we’re looking for.

Here’s how we can create the log from the example above:

```go
package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Info().
        Str("host", "192.168.0.2").
        Msg("Hello World!") 
}
// Output: {"time": 1562212768, "host": "192.168.0.2", "level": "info", "message": "Hello World!"}
```

## Correlation IDs

In a typical web app you would have your controllers, middleware, services and domain logic code that will be invoked on each user request. As we go down the stack, we can have multiple logs that are being written for  the same action, each log telling only a part of the story. Then, when  something bad happens and we want to investigate, we would need to bring the pieces together  like a puzzle, in order to make sense of the situation.

The  problem gets more convoluted in a distributed systems scenario, where  separate services are part of this one big user action, each outputting  their own logs that are difficult to follow and make sense of, when it comes to understanding  the bigger picture. Finding the service that is the culprit will be like playing detective.

To avoid having to become Sherlock as you  intuitively search through the big pile of messages, you could simply  connect them by having an additional field in the structure of the log - a **correlation id**. This id can be generated on each request, somewhere on top of the call chain and propagated down to be used as part of each logged message.

We can illustrate this concept in a Go app by creating and using a **HTTP middleware**, the **context object**, and our trusty old logger.

The basic principle of a middleware is that it’s a way of organizing a  shared functionality that we want to run on each HTTP request. This code usually sits between the router and the application  controllers.

**Note:** I won’t cover the whole story on how to create and use middlewares since there are [great posts](https://www.alexedwards.net/blog/making-and-using-middleware) about it that go into [detail](https://drstearns.github.io/tutorials/gomiddleware/), as well as some great libraries (such as [Negroni](https://github.com/urfave/negroni)), check them out for more info.

Anyway, the gist of it is that a middleware is essentially a function that implements the [http.Handler](https://pkg.go.dev/net/http?tab=doc#Handler) interface and it takes a handler function as a parameter which is the `next` handler that should be invoked, after the middleware code is done.

```go
func someMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Middleware logic.. 
    next.ServeHTTP(w, r)
  })
}
```

This allows us to chain multiple handlers  together, so in our scenario we’ll end up having a handler that  generates IDs and passes them down that wraps each application handler.

The flow of control will look something like this:

```
Router --> Generate correlation ID Middleware --> App handler
```

## An example

We can illustrate this in code with a simple example, where we will:

- Fetch the correlation id from a header if it’s present (something like **X-Correlation-Id**)
- In case the header is not present, we’ll generate a random id and add it to the response header. This will make our debugging easier since we’ll receive the id in our response.
- Pass it down to our logger instance
- Call the next handler

Okay, so first thing that we’ll need to do is initialize the logger and add it to the request context. This will allow us to use the logger object in our handlers. We’ll create a middleware that will inject the log into the request context, so that we can fetch the logger later on in our handlers.

```go
func logMiddleware(log zerolog.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            l := log.With().Logger()
 
            // l.WithContext returns a copy of the context with the log object associated
            r = r.WithContext(l.WithContext(r.Context()))
            
            next.ServeHTTP(w, r)
        })
    }
}
```

Next up is our middleware that will handle the correlation id. This handler will set a unique id to the request, which can be fetched from a header that we’ll define (such as X-Correlation-Id), or if the header is not present in the request, the handler will generate a new unique id.

This id will be added as a field to the log object and will be present in  the log structure for each message created down the line. The id is also added as a response header, which is great for debugging  purposes, since we’ll immediately have the id to query the log system.

In case we’ll need to send the id to another service, we’ll add it to the  request context so that we can easily fetch it when the need arises.

```go
func correlationIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        id := r.Header.Get("X-Correlation-Id")
        if id == "" {
            // generate new version 4 uuid
            newid := uuid.New()
            id = newid.String()
        } 
        // set the id to the request context
        ctx = context.WithValue(ctx, "correlation_id", id)
		r = r.WithContext(ctx) 
        // fetch the logger from context and update the context
        // with the correlation id value
        log := zerolog.Ctx(ctx)
        log.UpdateContext(func(c zerolog.Context) zerolog.Context {
            return c.Str("correlation_id", id)
        }) 
        // set the response header
        w.Header().Set("X-Correlation-Id", id)
        next.ServeHTTP(w, r) 
    })
}
```

**Note:** Zerolog provides a handy integration with the `net/http` package, and it has a bunch of useful middleware that you can use if you choose to use this as your logging library. These are just  simplified examples, you can choose a more robust and flexible solution  provided by the library. You can read more about that [here](https://github.com/rs/zerolog#integration-with-nethttp).

Okay so now let’s see our middleware in action and test it out in our full example:

```go
package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logMiddleware(log zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := log.With().Logger()

			// l.WithContext returns a copy of the context with the log object associated
			r = r.WithContext(l.WithContext(r.Context()))

			next.ServeHTTP(w, r)
		})
	}
}

func correlationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get("X-Correlation-Id")
		if id == "" {
			// generate new version 4 uuid
			newid := uuid.New()
			id = newid.String()
		}
		// set the id to the request context
		ctx = context.WithValue(ctx, "correlation_id", id)
		r = r.WithContext(ctx)
		// fetch the logger from context and update the context
		// with the correlation id value
		log := zerolog.Ctx(ctx)
		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("correlation_id", id)
		})
		// set the response header
		w.Header().Set("X-Correlation-Id", id)
		next.ServeHTTP(w, r)
	})
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logger := log.Ctx(r.Context())

	logger.Info().Msg("Processing request.")

	w.Write([]byte("OK"))

	logger.Info().Dur("elapsed", time.Since(start)).Msg("Done.")
}

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	mux := http.NewServeMux()

	mux.Handle("/",
		logMiddleware(log)(
			correlationIDMiddleware(
				http.HandlerFunc(testHandler),
			),
		),
	)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal().Err(err)
}
```

Here we have a `testHandler` function that we’ve wrapped in our log and correlation ID middleware (remember you can chain multiple middlewares together) and we have added two log messages so we can show that the correlation id is correctly added to our logs.

When we start the example program and execute a request, the logs will appear in the terminal like so:

```json
{"level":"info","correlation_id":"2fda2736-7d2d-441e-82a2-24df9a1199a0","time":1601569068,"message":"Processing request."}

{"level":"info","correlation_id":"2fda2736-7d2d-441e-82a2-24df9a1199a0","elapsed":0.030087,"time":1601569068,"message":"Done."}
```

When running `curl -I localhost:8080` in the command line, we could also see that our correlation id is added in the response headers:

```
HTTP/1.1 200 OK
X-Correlation-Id: 2fda2736-7d2d-441e-82a2-24df9a1199a0 
Date: Thu, 01 Oct 2020 16:23:37 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8
```

## Conclusion

Troubleshooting a  problem in a system is difficult, even more so when the components that  the system is made of are distributed on multiple machines. The logs are not guaranteed to arrive sequentially, and it’s a challenge trying to make sense of  what’s happened.

Instead of wasting time on piecing together the  scattered information, a correlation ID will provide cohesion to the  logs, unifying the related messages and providing you a way of tracking  down the problem sooner, rather than later.

### LATEST POSTS
- [What is the future of databases?](https://filipnikolovski.com/posts/what-is-the-future-of-databases/)
- [TIL: eBPF is awesome](https://filipnikolovski.com/posts/ebpf/)
- [Correlating Logs](https://filipnikolovski.com/posts/correlating-logs/)
- [Bazel Performance in a CI Environment](https://filipnikolovski.com/posts/bazel-performance-in-a-ci-environment/)
- [Avoiding Pitfalls During Service Deployments](https://filipnikolovski.com/posts/avoiding-pitfalls-during-service-deployments/)
