# 12 factor configuration with Go's flag package

# 使用 Go 的 flag 包进行 12 因子配置

Mon, Sep 16, 2019

2019 年 9 月 16 日，星期一

Cost-effective way to have your app conform with [12 factor](https://12factor.net) methodology with [Go](https://golang.org)'s stock [`flag`](https://golang.org/pkg/flag/) package.

使用 [Go](https://golang.org) 的股票 [`flag`](https://golang.org/pkg/flag/) 包。

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#summary)Summary

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#summary)总结

Previously, before “cloud” was a thing, it was common to have configuration part of the source code, ie Rails' [`config/database.yaml`](https://edgeguides.rubyonrails.org/configuring.html#configuring-a-database).

以前，在“云”成为事物之前，源代码的配置部分是很常见的，即Rails的[`config/database.yaml`](https://edgeguides.rubyonrails.org/configuring.html#configuring-a-数据库)。

These days, with immutable infrastucture, separation of configuration and code is preferred; quoting [12 factor](https://12factor.net):

如今，由于基础设施不可变，配置和代码分离是首选；引用 [12 因子](https://12factor.net)：

```text
III.Config
Store config in the environment
An app’s config is everything that is likely to vary between deploys (staging, production, developer environments, etc).This includes:
- Resource handles to the database, Memcached, and other backing services
- Credentials to external services such as Amazon S3 or Twitter
- Per-deploy values such as the canonical hostname for the deploy
Apps sometimes store config as constants in the code.This is a violation of twelve-factor, which requires strict separation of config from code.Config varies substantially across deploys, code does not.
```

– https://12factor.net/config

– https://12factor.net/config

This means that the app’s context sets the configuration which enables the  app to run transparently as a serverless function, in a kubernetes pod,  in a cloud run, in a docker swarm, or your laptop.

这意味着应用程序的上下文设置了配置，使应用程序能够作为无服务器功能在 kubernetes pod、云运行、docker swarm 或您的笔记本电脑中透明地运行。

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#problem)Problem

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#problem)问题

<iframe id="twitter-widget-0" scrolling="no" allowtransparency="true" allowfullscreen="true" class="" style="position: absolute; visibility: hidden; width: 0px; height: 0px; display : block; flex-grow: 1;" title="Twitter Tweet" src="https://platform.twitter.com/embed/Tweet.html?dnt=false&embedId=twitter-widget-0&features=eyJ0ZndfZXhwZXJpbWVudHNfY29va2llX2V4cGlyYXRpb24iOnsiYnVja2V0IjoxMjA5NjAwLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X2hvcml6b25fdHdlZXRfZW1iZWRfOTU1NSI6eyJidWNrZXQiOiJodGUiLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3NwYWNlX2NhcmQiOnsiYnVja2V0Ijoib2ZmIiwidmVyc2lvbiI6bnVsbH19&frame=false&hideCard=false&hideThread=false&id=1166497222894055426&lang=en&origin=https%3A %2F%2Fwww.gmarik.info%2Fblog%2F2019%2F12-factor-golang-flag-package%2F&sessionId=63e2a4145f43d9416c3a9132d4d53ccb7d2a81d7&theme=light&widgetsVersion=1890d59c%3A1627936082797&width=550px" frameborder="0"></iframe>

<iframe id="twitter-widget-0" scrolling="no" allowtransparency="true" allowfullscreen="true" class="" style="位置：绝对；可见性：隐藏；宽度：0px；高度：0px；显示: 块; flex-grow: 1;"标题= “微博资料Tweet” SRC =“https://platform.twitter.com/embed/Tweet.html?dnt=false&embedId=twitter-widget-0&features=eyJ0ZndfZXhwZXJpbWVudHNfY29va2llX2V4cGlyYXRpb24iOnsiYnVja2V0IjoxMjA5NjAwLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X2hvcml6b25fdHdlZXRfZW1iZWRfOTU1NSI6eyJidWNrZXQiOiJodGUiLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3NwYWNlX2NhcmQiOnsiYnVja2V0Ijoib2ZmIiwidmVyc2lvbiI6bnVsbH19&frame=false&hideCard=false&hideThread=false&id=1166497222894055426&lang=en&origin=https%3A %2F%2Fwww.gmarik.info%2Fblog%2F2019%2F12-factor-golang-flag-package%2F&sessionId=63e2a4145f43d9416c3a9132d4d53ccb7d2a81d7&theme=light&widgets%730x70x750x70x7500000帧

> replaced `viper` with `flag` package and🤯.
>
> How do you justify adding a dependency if stdlib provides same functionality even if some plumbing required? [#golang](https://twitter.com/hashtag/golang?src=hash&ref_src=twsrc^tfw)[pic.twitter.com/4fAoXVP7vU](https://t.co/4fAoXVP7vU)
>
> — gmarik (@gmarik)
>
> August 27, 2019

> 将 `viper` 替换为 `flag` 包和🤯。
>
> 如果 stdlib 提供相同的功能，即使需要一些管道，您如何证明添加依赖项的合理性？ [#golang](https://twitter.com/hashtag/golang?src=hash&ref_src=twsrc^tfw)[pic.twitter.com/4fAoXVP7vU](https://t.co/4fAoXVP7vU)
>
> — gmarik (@gmarik)
>
> 2019 年 8 月 27 日

Surprisingly often, in order to fulfill [12 Factor](https://12factor.net) config requirements, people resort to packages with large API surface and as result large codebase and deep dependency graph.

令人惊讶的是，为了满足 [12 Factor](https://12factor.net) 配置要求，人们往往求助于具有大型 API 表面的包，从而导致大型代码库和深度依赖关系图。

Often times this is not necessary since the same functionality can be  achieved with much less code and only using Go’s standard library  packages. Here’s an example of using [`flag`](https://golang.org/pkg/flag/) package to achieve equal result.

通常这不是必需的，因为可以用更少的代码并且只使用 Go 的标准库包来实现相同的功能。这是一个使用 [`flag`](https://golang.org/pkg/flag/) 包来实现相同结果的示例。

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#12-factor-config-with-flag-package)12 factor config with `flag` package

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#12-factor-config-with-flag-package)12 因子配置与`flag`包裹

- 2 ways to configure the app, through: 1) cli flags or 2) environment variables
- default values are configured from corresponding variables
- environment variables, if configured, set the `flag`'s defaults using `LookupOr*` helpers
- get full configuration with simple `getConfig` helper

- 2 种配置应用程序的方法，通过：1) cli 标志或 2) 环境变量
- 默认值由相应的变量配置
- 环境变量，如果已配置，使用 `LookupOr*` 助手设置 `flag` 的默认值
- 使用简单的 `getConfig` 助手获得完整的配置

```golang
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "log"
)

var (
    // set by build process
    Git_Revision string
    Consul_URL string = "http://consul.local:8500"
    Statsd_URL string
    HTTP_ListenAddr string = ":8080"
    HTTP_Timeout    int    = 16
)

func main() {
    flag.StringVar(&Consul_URL, "consul-url", LookupEnvOrString("CONSUL_URL", Consul_URL), "service discovery url")
    flag.StringVar(&Statsd_URL, "statsd-url", LookupEnvOrString("STATSD_URL", Statsd_URL), "statsd's host:port")
    flag.StringVar(&HTTP_ListenAddr, "http-listen-addr", LookupEnvOrString("HTTP_LISTEN_ADDR", HTTP_ListenAddr), "http service listen address")
    flag.IntVar(&HTTP_Timeout, "http-timeout", LookupEnvOrInt("HTTP_TIMEOUT", HTTP_Timeout), "http timeout requesting http services")

    flag.Parse()
    log.Printf("app.config %v\n", getConfig(flag.CommandLine))

    log.Println("app.status=starting")
    defer log.Println("app.status=shutdown")

    log.Println("hello world")
}

func LookupEnvOrString(key string, defaultVal string) string {
    if val, ok := os.LookupEnv(key);ok {
        return val
    }
    return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
    if val, ok := os.LookupEnv(key);ok {
        v, err := strconv.Atoi(val)
        if err != nil {
            log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
        }
        return v
    }
    return defaultVal
}

func getConfig(fs *flag.FlagSet) []string {
    cfg := make([]string, 0, 10)
    fs.VisitAll(func(f *flag.Flag) {
        cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
    })

    return cfg
}
```

see it in action on [Playground](https://play.golang.org/p/CPstmhyrk47)

在 [Playground](https://play.golang.org/p/CPstmhyrk47) 上看到它的实际效果

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#conclusion)Conclusion

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#conclusion)结论

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#pros)Pros

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#pros)

- no dependencies other than standard library

- 没有标准库以外的依赖项

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#cons)Cons

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#cons)

- a bit of plumbing code is required
- defaults to environment var’s value if latter is set
- env vars are manually named
- description may duplicate var’s comments

- 需要一些管道代码
- 如果设置了后者，则默认为环境变量的值
- 环境变量是手动命名的
- 描述可能会重复 var 的评论

[`flag`](https://golang.org/pkg/flag/) package with combination with few helpers provides pragmatic way to configure your [12 factor](https://12factor.net)-ready apps. It’s not perfect but gets the job done.

[`flag`](https://golang.org/pkg/flag/) 包与少量帮助程序相结合，提供了一种实用的方式来配置您的 [12 factor](https://12factor.net)-ready 应用程序。它并不完美，但可以完成工作。

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#references)References

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#references)参考

- [Simplicity Matters by Rich Hickey](https://www.youtube.com/watch?v=rI8tNMsozo0)

- [Simplicity Matters by Rich Hickey](https://www.youtube.com/watch?v=rI8tNMsozo0)

##### Related Posts

#### #  相关文章

- [Monotonic Time, or Perfect model vs Imperfect Reality](https://www.gmarik.info/blog/2019/monotonic-time-perfect-model-vs-imperfect-reality/)
- [Testing MongoDB queries with Golang](https://www.gmarik.info/blog/2017/testing-mongodb-queries-golang/)
- [Wordfight: a multi-player word game](https://www.gmarik.info/blog/2017/wordfight-multiplayer-word-game/)
- [Understanding Go's `for` loop with closures](https://www.gmarik.info/blog/2016/understanding-golang-for-loop-with-closures/)
- [Experimenting with Go pipelines](https://www.gmarik.info/blog/2016/experimenting-with-golang-pipelines/) 

- [单调时间，或完美模型 vs 不完美现实](https://www.gmarik.info/blog/2019/monotonic-time-perfect-model-vs-imperfect-reality/)
- [使用 Golang 测试 MongoDB 查询](https://www.gmarik.info/blog/2017/testing-mongodb-queries-golang/)
- [Wordfight：多人文字游戏](https://www.gmarik.info/blog/2017/wordfight-multiplayer-word-game/)
- [了解 Go 的 `for` 闭包循环](https://www.gmarik.info/blog/2016/understanding-golang-for-loop-with-closures/)
- [使用 Go 管道进行实验](https://www.gmarik.info/blog/2016/experimenting-with-golang-pipelines/)

