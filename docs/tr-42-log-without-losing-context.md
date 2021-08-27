# Logging without losing money or context.

# 在不损失金钱或上下文的情况下进行日志记录。

26 May 2020

2020 年 5 月 26 日

1. [Problem](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Problem)
2. [Proposed Solution](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Solution)
3. [Conclusion](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Conclusion)

1. [问题](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Problem)
2. [建议的解决方案](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Solution)
3. [结论](https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context#Conclusion)

## Problem

##  问题

In your application, you want all flows of execution to be covered with log statements. That way, if something were to go wrong, you can be able to [trace](https://www.komu.engineer/blogs/timescaledb/timescaledb-for-logs#Opinion3) at which point the flow broke at.
However, this presents a problem; if your application has lots of traffic, then the amount of logs it generates are going to be gigantic. This in itself is not a problem, were it not for the fact that you are using a logging as a service provider and [they](https://www.datadoghq.com/pricing/)all [charge](https:/ /www.loggly.com/plans-and-pricing/) an [arm](https://www.honeycomb.io/pricing/)and a [leg](https://www.sumologic.com/pricing/ ) for every log event.
Some of the pricing models are so inscrutable, that the logging service providers offer calculators ( [here](https://calculator.aws/),[here](https://cloud.google.com/products/calculator), [etc](https://azure.microsoft.com/en-us/pricing/calculator/)) to try and help you figure out what should ideally have been kindergaten arithmetic.

在您的应用程序中，您希望日志语句涵盖所有执行流程。这样，如果出现问题，您可以 [trace](https://www.komu.engineer/blogs/timescaledb/timescaledb-for-logs#Opinion3) 流在此时中断。
然而，这带来了一个问题；如果您的应用程序有大量流量，那么它生成的日志量将是巨大的。这本身不是问题，如果不是因为您将日志记录用作服务提供商并且 [他们](https://www.datadoghq.com/pricing/)[charge](https:/ /www.loggly.com/plans-and-pricing/) 一个 [arm](https://www.honeycomb.io/pricing/)和一个 [leg](https://www.sumologic.com/pricing/ ) 对于每个日志事件。
一些定价模型非常难以理解，以至于日志服务提供商提供计算器（[此处](https://calculator.aws/)、[此处](https://cloud.google.com/products/calculator)、[etc](https://azure.microsoft.com/en-us/pricing/calculator/)) 尝试帮助您找出理想情况下应该是幼儿园数学。

So it seems like for you to have your cake and eat it too, you have to part with large sums of money every month(and due to elastic pricing you can't even tell in advance how much that will be.)
There are two main ways that people try and solve this problem(usually at the [suggestion](https://docs.datadoghq.com/logs/indexes/#examples-of-exclusion-filters)of the [said](https ://docs.honeycomb.io/working-with-your-data/best-practices/sampling/) logging service providers);
1. Filtering logs by severity and only sending logs above a certain severity to the logging service provider.
2. and/or
3. Sampling logs so that you only send a certain percentage to the logging service provider.

因此，对于您来说，似乎既要吃蛋糕又要吃蛋糕，每个月都必须分出一大笔钱（由于弹性定价，您甚至无法提前知道要花多少钱。）
人们尝试解决此问题的主要方法有两种（通常在 [said](https://docs.datadoghq.com/logs/indexes/#examples-of-exclusion-filters)://docs.honeycomb.io/working-with-your-data/best-practices/sampling/) 日志服务提供商）；
1、按严重性过滤日志，只将高于一定严重性的日志发送给日志服务提供商。
2. 和/或
3. 对日志进行采样，以便您只将一定比例的日志发送给日志服务提供商。

But these two solutions pose a problem; loss of context.
Consider an application that updates multiple social media platforms with a status message.

但这两种解决方案带来了一个问题；上下文丢失。
考虑使用状态消息更新多个社交媒体平台的应用程序。

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

如果我们过滤日志并且只将 ERROR 级别的日志发送给我们的日志服务提供商，那么我们将失去关于 *facebook send failed* 错误是如何产生的上下文。我之前写过日志主要用于[帮助调试生产中的问题；](https://www.komu.engineer/blogs/timescaledb/timescaledb-for-logs) 因此，导致事件的上下文和时间顺序特定的问题很重要。你不想调查一个谋杀之谜，其中一半的线索被你以前的自己故意抹去了。
同样，如果我们对日志进行采样；导致错误的年表将缺少一些 INFO 日志语句，因为这些语句已被采样。

What we want is;
If the application has not emitted any errors, no logs gets sent to our logging service provider(INFO or otherwise.) However, if there are any errors emitted; all the logs leading to that error irrespective of their severity/level are sent to the logging service.

我们想要的是；
如果应用程序没有发出任何错误，则不会将日志发送到我们的日志记录服务提供者（INFO 或其他方式）。但是，如果发出任何错误；所有导致该错误的日志，无论其严重性/级别如何，都将发送到日志记录服务。

## Proposed Solution

## 建议的解决方案

I think we should be able to implement such a logging scheme. The basic idea is; whenever your application emits log events, all the logs go into a [circular buffer](https://en.wikipedia.org/wiki/Circular_buffer) of size X. Whenever an error log is emitted, the whole circular buffer is flushed and all its contents are sent to the logging service provider.
The circular buffer can be in memory or on disk/whatever and the size is configurable. 
我认为我们应该能够实现这样的日志记录方案。基本思想是；每当您的应用程序发出日志事件时，所有日志都会进入大小为 X 的 [循环缓冲区](https://en.wikipedia.org/wiki/Circular_buffer)。每当发出错误日志时，整个循环缓冲区都会被刷新并它的所有内容都被发送到日志服务提供者。
循环缓冲区可以在内存或磁盘/任何地方，并且大小是可配置的。
I took a stub at implementing this using [sirupsen/logrus](https://github.com/sirupsen/logrus) which is a popular logging library for the Go programming language, but the implementation should be transferable across libraries/languages.
In [sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus),you can declare a [hook](https://pkg.go.dev/github.com/sirupsen/ logrus?tab=doc#Hook) implementing the custom behaviour that you want.

我使用 [sirupsen/logrus](https://github.com/sirupsen/logrus) 实现了这一点，这是 Go 编程语言的流行日志库，但实现应该可以跨库/语言转移。
在[sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus)中，可以声明一个[hook](https://pkg.go.dev/github.com/sirupsen/logrus?tab=doc#Hook) 实现您想要的自定义行为。

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

在您的应用程序中使用它的方法是；

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

现在，如果发生任何错误；导致错误的所有日志和年表都可用并发送给日志记录服务提供商。

```go
go run .

{"level":"info","msg":"tweet send start","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"info","msg":"tweet send end.","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"info","msg":"facebook send start","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
{"level":"error","msg":"facebook send failed. error=http 500","time":"2020-05-25T21:03:36+03:00","traceID":"sa225Hqk"}
        
```


Conversely, if there are no errors been emitted by your application then no logs are sent to your logging service provider.

相反，如果您的应用程序没有发出任何错误，则不会向您的日志记录服务提供商发送日志。

```go
go run .
# no log output
```


So you do not end up spending tons of money and you also do not lose context when errors occur.

因此，您最终不会花费大量金钱，并且在发生错误时也不会丢失上下文。

## Conclusion

##  结论

You can implement a logging strategy that loses you neither money nor context.
A downside of the presented solution is that it can be hard to tell if there are no logs because the application has not emitted any errors or because the logging pipeline itself has a bug. However this can be solved by emitting a heartbeat log event every Y minutes and letting this heartbeat propagate upto the logging service provider.

您可以实施既不会损失金钱也不会损失上下文的日志记录策略。
所提出的解决方案的一个缺点是，很难判断是否没有日志，因为应用程序没有发出任何错误，或者因为日志管道本身有错误。然而，这可以通过每 Y 分钟发出一个心跳日志事件并让这个心跳传播到日志服务提供者来解决。

All the code in this blogpost, including the full source code, can be found at: https://github.com/komuw/komu.engineer/tree/master/blogs/log-without-losing-context

这篇博文中的所有代码，包括完整的源代码，可以在以下位置找到：https://github.com/komuw/komu.engineer/tree/master/blogs/log-without-losing-context

You can comment on this article [by clicking here.](https://github.com/komuw/komu.engineer/issues/17)https://github.com/komuw/komu.engineer/issues/17) 
你可以评论这篇文章 [点击这里。](https://github.com/komuw/komu.engineer/issues/17)https://github.com/komuw/komu.engineer/issues/17)
