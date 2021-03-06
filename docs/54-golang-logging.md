# Logging in Go: Choosing a System and Using it

Apr 1, 2020 From: https://www.honeybadger.io/blog/golang-logging/

Go has built-in features to make it easier  for programmers to implement logging. Third parties have also built  additional tools to make logging easier. What's the difference between  them? Which should you choose? In this article Ayooluwa Isaiah describes both of these and discusses when you'd prefer one over the other.

You're relatively new to the Go language. You're probably using it to write a web app or a server, and you need to create a log  file. So, you do a quick web search and find that there are a ton of  options for logging in go. How do you know which one to pick? This  article will equip you to answer that question. 

We will take a look at the built-in `log` package and  determine what projects it is suited for before exploring other logging  solutions that are prevalent in the Go ecosystem.

## What to log

I don't need to tell you how important logging is. Logs are used by  every production web application to help developers and operations:

- Spot bugs in the application's code
- Discover performance problems
- Do post-mortem analysis of outages and security incidents

The data you actually log will depend on the type of application  you're building. Most of the time, you will have some variation in the  following: 

- The timestamp for when an event occurred or a log was generated
- Log levels such as debug, error, or info
- Contextual data to help understand what happened and make it possible to easily reproduce the situation

## What not to log

In general, you shouldn't log any form of sensitive business data or  personally identifiable information. This includes, but is not limited  to:

- Names
- IP addresses
- Credit card numbers

These restrictions can make logs less useful from an engineering  perspective, but they make your application more secure. In many cases,  regulations such as GDPR and HIPAA may forbid the logging of personal  data. 

## Introducing the log package

The Go standard library has a built-in `log` package that  provides most basic logging features. While it does not have log levels  (such as debug, warning, or error), it still provides everything you  need to get a basic logging strategy set up.

Here's the most basic logging example:

```
package main

import "log"

func main() {
    log.Println("Hello world!")
}
```

The code above prints the text "Hello world!" to the standard error,  but it also includes the date and time, which is handy for filtering log messages by date.

```
2019/12/09 17:21:53 Hello world!
```

> By default, the `log` package prints to the standard error (`stderr`) output stream, but you can make it write to local files or any destination that supports the `io.Writer` interface. It also adds a timestamp to the log message without any additional configuration.

## Logging to a file

If you need to store log messages in a file, you can do so by  creating a new file or opening an existing file and setting it as the  output of the log. Here's an example:

```
package main

import (
    "log"
    "os"
)

func main() {
    // If the file doesn't exist, create it or append to the file
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)

    log.Println("Hello world!")
}
```

When we run the code, the following is written to `logs.txt.` 

```
2019/12/09 17:22:47 Hello world!
```

As mentioned earlier, you can basically output your logs to any destination that implements the `io.Writer` interface, so you have a lot of flexibility when deciding where to log messages in your application.

## Creating custom loggers

Although the `log` package implements a predefined `logger` that writes to the standard error, we can create custom logger types using the `log.New()` method.

When creating a new logger, you need to pass in three arguments to `log.New()`:

- `out`: Any type that implements the `io.Writer` interface, which is where the log data will be written to
- `prefix`: A string that is appended to the beginning of each log line
- `flag`: A set of constants that allow us to define which  logging properties to include in each log entry generated by the logger  (more on this in the next section)

We can take advantage of this feature to create custom loggers. Here's an example that implements `Info`, `Warning` and `Error` loggers:

```
package main

import (
    "log"
    "os"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
    InfoLogger.Println("Starting the application...")
    InfoLogger.Println("Something noteworthy happened")
    WarningLogger.Println("There is something you should know about")
    ErrorLogger.Println("Something went wrong")
}
```

After creating or opening the `logs.txt` file at the top of the `init` function, we then initialize the three defined loggers by providing the output destination, prefix string, and log flags.

In the `main` function, the loggers are utilized by calling the `Println` function, which writes a new log entry to the log file. When you run this program, the following will be written to `logs.txt`.

```
INFO: 2019/12/09 12:01:06 main.go:26: Starting the application...
INFO: 2019/12/09 12:01:06 main.go:27: Something noteworthy happened
WARNING: 2019/12/09 12:01:06 main.go:28: There is something you should know about
ERROR: 2019/12/09 12:01:06 main.go:29: Something went wrong
```

Note that in this example, we are logging to a single file, but you  can use a separate file for each logger by passing a different file when creating the logger.

## Log flags

You can use [log flag constants](https://golang.org/pkg/log/#pkg-constants) to enrich a log message by providing additional context information,  such as the file, line number, date, and time. For example, passing the  message "Something went wrong" through a logger with a flag combination  shown below:

```
log.Ldate|log.Ltime|log.Lshortfile
```

will print

```
2019/12/09 12:01:06 main.go:29: Something went wrong
```

Unfortunately, there is no control over the order in which they appear or the format in which they are presented.

## Introducing logging frameworks

Using the `log` package is great for local development  when getting fast feedback is more important than generating rich,  structured logs. Beyond that, you will mostly likely be better off using a logging framework. 

A major advantage of using a logging framework is that it helps to standardize the log data. This means that:

- It's easier to read and understand the log data.
- It's easier to gather logs from several sources and feed them to a central platform to be analyzed.

In addition, logging is pretty much a solved problem. Why reinvent the wheel?

## Choosing a logging framework

Deciding which framework to use can be a challenge, as there are [several options](https://github.com/avelino/awesome-go#logging) to choose from.

The two most popular logging frameworks for Go appear to be [glog](https://github.com/golang/glog) and [logrus](https://github.com/Sirupsen/logrus). The popularity of glog is surprising, since it hasn't been updated in  several years. logrus is better maintained and used in popular projects  like Docker, so we'll be focusing on it. 

## Getting started with logrus

Installing logrus is as simple as running the command below in your terminal:

```
go get "github.com/Sirupsen/logrus"
```

One great thing about logrus is that it's completely compatible with the `log` package of the standard library, so you can replace your log imports everywhere with `log "github.com/sirupsen/logrus"` and it will just work!

Let's modify our earlier "hello world" example that used the log package and use logrus instead:

```
package main

import (
  log "github.com/sirupsen/logrus"
)

func main() {
    log.Println("Hello world!")
}
```

Running this code produces the output:

```
INFO[0000] Hello world!
```

It couldn't be any easier!

### Logging in JSON

`logrus` is well suited for structured logging in JSON  which ??? as JSON is a well-defined standard ??? makes it easy for external  services to parse your logs and also makes the addition of context to a  log message relatively straightforward through the use of fields, as  shown below:

```
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.JSONFormatter{})
    log.WithFields(
        log.Fields{
            "foo": "foo",
            "bar": "bar",
        },
    ).Info("Something happened")
}
```

The log output generated will be a JSON object that includes the message, log level, timestamp, and included fields.

```
{"bar":"bar","foo":"foo","level":"info","msg":"Something happened","time":"2019-12-09T15:55:24+01:00"}
```

If you're not interested in outputting your logs as JSON, be aware  that several third-party formatters exist for logrus, which you can view on its [Github page](https://github.com/Sirupsen/logrus#formatters). You can even write your own formatter if you prefer.

### Log levels

Unlike the standard log package, logrus supports log levels. 

logrus has seven log levels: Trace, Debug, Info, Warn, Error, Fatal,  and Panic. The severity of each level increases as you go down the list.

```
log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
```

By setting a logging level on a logger, you can log only the entries  you need depending on your environment. By default, logrus will log  anything that is Info or above (Warn, Error, Fatal, or Panic).

```
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.JSONFormatter{})

    log.Debug("Useful debugging information.")
    log.Info("Something noteworthy happened!")
    log.Warn("You should probably take a look at this.")
    log.Error("Something failed but I'm not quitting.")
}
```

Running the above code will produce the following output:

```
{"level":"info","msg":"Something noteworthy happened!","time":"2019-12-09T16:18:21+01:00"}
{"level":"warning","msg":"You should probably take a look at this.","time":"2019-12-09T16:18:21+01:00"}
{"level":"error","msg":"Something failed but I'm not quitting.","time":"2019-12-09T16:18:21+01:00"}
```

Notice that the Debug level message was not printed. To include it in the logs, set `log.Level` to equal `log.DebugLevel`:

```
log.SetLevel(log.DebugLevel)
```

## Wrap up

In this post, we explored the use of the built-in log package and  established that it should only be used for trivial applications or when building a quick prototype. For everything else, the use of a  mainstream logging framework is a must. 

We also looked at ways to ensure that the information contained in your logs is consistent and easy to analyze, especially when aggregating it on a centralized platform.

Thanks for reading!