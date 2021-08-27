# Logging without losing money or context. 

26 May 2020

1. [Problem](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Problem)
2. [Proposed Solution](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Solution)
3. [Conclusion](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Conclusion)

## Problem

In your application, you want all flows of execution to be covered with log statements. That way, if something were to go wrong, you can be able to [trace](https://www.komu.engineer/blogs/timescaledb/timescaledb-for-logs#Opinion3) at which point the flow broke at.
However, this presents a problem; if your application has lots of traffic, then the amount of logs it generates are going to be gigantic. This in itself is not a problem, were it not for the fact that you are using a logging as a service provider and [they](https://www.datadoghq.com/pricing/) all [charge](https://www.loggly.com/plans-and-pricing/) an [arm](https://www.honeycomb.io/pricing/) and a [leg](https://www.sumologic.com/pricing/) for every log event.
Some of the pricing models are so inscrutable, that the logging service providers offer calculators ( [here](https://calculator.aws/), [here](https://cloud.google.com/products/calculator), [etc](https://azure.microsoft.com/en-us/pricing/calculator/)) to try and help you figure out what should ideally have been kindergaten arithmetic.

So it seems like for you to have your cake and eat it too, you have to part with large sums of money every month(and due to elastic pricing you can't even tell in advance how much that will be.)
There are two main ways that people try and solve this problem(usually at the [suggestion](https://docs.datadoghq.com/logs/indexes/#examples-of-exclusion-filters) of the [said](https://docs.honeycomb.io/working-with-your-data/best-practices/sampling/) logging service providers);
1. Filtering logs by severity and only sending logs above a certain severity to the logging service provider.
2. and/or
3. Sampling logs so that you only send a certain percentage to the logging service provider.

But these two solutions pose a problem; loss of context.
Consider an application that updates multiple social media platforms with a status message.

```go
func main() {
    updateSocialMedia("Sending out my first social media status message")
}

func updateSocialMedia(msg string) {
	traceID := "sa225Hqk" //should be randomly generated per call
	logger := logrus.WithFields(logrus.Fields{"traceID": traceID})

	tweet(msg, logger)
	facebook(msg, logger)
	linkedin(msg, logger)
}

func tweet(msg string, logger *logrus.Entry) {
	logger.Info("tweet send start")
	// code to call twitter API goes here
	logger.Info("tweet send end.")
}

func facebook(msg string, logger *logrus.Entry) {
	logger.Info("facebook send start")
	err := facebookAPI(msg)
	if err != nil {
		logger.Errorf("facebook send failed. error=%s", err)
	}
	logger.Info("facebook send end.")
}

func linkedin(msg string, logger *logrus.Entry) {
	logger.Info("linkedin send start")
	// code to call linkedin API goes here
	logger.Info("linkedin send end.")
}
                
```

If we were filtering logs and only sending logs of ERROR level to our logging service provider, then we would lose context on how the *facebook send failed* error came to be. I had previosuly written that logs are primarily used to [help debug issues in production;](https://www.komu.engineer/blogs/timescaledb/timescaledb-for-logs) thus, context and chronology of events that led to a particular issue are of importance. You do not want to investigate a murder mystery where half the clues have been deliberately wiped out by your earlier self.
In the same way, if we were sampling logs; the chronology leading upto the error would be missing a few INFO log statements since those would have been sampled out.

What we want is;
If the application has not emitted any errors, no logs gets sent to our logging service provider(INFO or otherwise.) However, if there are any errors emitted; all the logs leading to that error irrespective of their severity/level are sent to the logging service. 

## Proposed Solution

I think we should be able to implement such a logging scheme. The basic idea is; whenever your application emits log events, all the logs go into a [circular buffer](https://en.wikipedia.org/wiki/Circular_buffer) of size X. Whenever an error log is emitted, the whole circular buffer is flushed and all its contents are sent to the logging service provider.
The circular buffer can be in memory or on disk/whatever and the size is configurable.

I took a stub at implementing this using [sirupsen/logrus](https://github.com/sirupsen/logrus) which is a popular logging library for the Go programming language, but the implementation should be transferable across libraries/languages.
In [sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus), you can declare a [hook](https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#Hook) implementing the custom behaviour that you want.

```go
package main

import (
    "io"

    "github.com/sirupsen/logrus"
)

// hook to buffer logs and only send at right severity.
type hook struct {
    writer io.Writer

    // Note: in production, lineBuffer should use a circular buffer instead of a slice.
    // otherwise you may have unbounded memory growth.
    // we are just using a slice of []bytes here for brevity and blogging purposes.
    lineBuffer [][]byte
}

// Fire will append all logs to a circular buffer and only 'flush'
// them when a log of sufficient severity(ERROR) is emitted.
func (h *hook) Fire(entry *logrus.Entry) error {
    line, err := entry.Bytes()
    if err != nil {
        return err
    }
    h.lineBuffer = append(h.lineBuffer, line)

    if entry.Level <= logrus.ErrorLevel {
        var writeError error
        for _, line := range h.lineBuffer {
            _, writeError = h.writer.Write(line)
        }
        h.lineBuffer = nil // clear the buffer
        return writeError
    }

    return nil
}

// Levels define on which log levels this hook would trigger
func (h *hook) Levels() []logrus.Level {
    return logrus.AllLevels
}
                    
```

And the way to use it in your application is;

```go
package main

import (
    "errors"
    "io/ioutil"
    "math/rand"
    "os"
    "time"

    "github.com/sirupsen/logrus"
)

func main() {
    // send logs to nowhere by default
    logrus.SetOutput(ioutil.Discard)
    logrus.SetFormatter(&logrus.JSONFormatter{})
    
    // Use our custom hook that will append logs to a circular buffer
    // and ONLY flush them to stderr when errors occur.
    logrus.AddHook(&hook{writer: os.Stderr})

    updateSocialMedia("Sending out my first social media status message")
}
        
```

Now, if any error occurs; all the logs and chronology leading upto the error are available and are sent to the logging service provider.

```go
go run .

{"level":"info","msg":"tweet send start","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"info","msg":"tweet send end.","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"info","msg":"facebook send start","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"error","msg":"facebook send failed. error=http 500","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
        
```

Conversely, if there are no errors been emitted by your application then no logs are sent to your logging service provider.

```go
go run .
# no log output
```

So you do not end up spending tons of money and you also do not lose context when errors occur. 

## Conclusion

You can implement a logging strategy that loses you neither money nor context.
A downside of the presented solution is that it can be hard to tell if there are no logs because the application has not emitted any errors or because the logging pipeline itself has a bug. However this can be solved by emitting a heartbeat log event every Y minutes and letting this heartbeat propagate upto the logging service provider.

All the code in this blogpost, including the full source code, can be found at: https://github.com/komuw/komu.engineer/tree/master/blogs/log-without-losing-context

You can comment on this article [by clicking here.](https://github.com/komuw/komu.engineer/issues/17) https://github.com/komuw/komu.engineer/issues/17)