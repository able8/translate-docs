# 12 factor configuration with Go's flag package

Mon, Sep 16, 2019

Cost-effective way to have your app conform with [12 factor](https://12factor.net) methodology with [Go](https://golang.org)’s stock [`flag`](https://golang.org/pkg/flag/) package.

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#summary)Summary

Previously, before “cloud” was a thing, it was common to have configuration part of the source code, ie Rails’ [`config/database.yaml`](https://edgeguides.rubyonrails.org/configuring.html#configuring-a-database).

These days, with immutable infrastucture, separation of configuration and code is preferred; quoting [12 factor](https://12factor.net):

```text
III. Config
Store config in the environment
An app’s config is everything that is likely to vary between deploys (staging, production, developer environments, etc). This includes:
- Resource handles to the database, Memcached, and other backing services
- Credentials to external services such as Amazon S3 or Twitter
- Per-deploy values such as the canonical hostname for the deploy
Apps sometimes store config as constants in the code. This is a violation of twelve-factor, which requires strict separation of config from code. Config varies substantially across deploys, code does not.
```

– https://12factor.net/config

This means that the app’s context sets the configuration which enables the  app to run transparently as a serverless function, in a kubernetes pod,  in a cloud run, in a docker swarm, or your laptop.

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#problem)Problem

<iframe id="twitter-widget-0" scrolling="no" allowtransparency="true" allowfullscreen="true" class="" style="position: absolute; visibility: hidden; width: 0px; height: 0px; display: block; flex-grow: 1;" title="Twitter Tweet" src="https://platform.twitter.com/embed/Tweet.html?dnt=false&amp;embedId=twitter-widget-0&amp;features=eyJ0ZndfZXhwZXJpbWVudHNfY29va2llX2V4cGlyYXRpb24iOnsiYnVja2V0IjoxMjA5NjAwLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X2hvcml6b25fdHdlZXRfZW1iZWRfOTU1NSI6eyJidWNrZXQiOiJodGUiLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3NwYWNlX2NhcmQiOnsiYnVja2V0Ijoib2ZmIiwidmVyc2lvbiI6bnVsbH19&amp;frame=false&amp;hideCard=false&amp;hideThread=false&amp;id=1166497222894055426&amp;lang=en&amp;origin=https%3A%2F%2Fwww.gmarik.info%2Fblog%2F2019%2F12-factor-golang-flag-package%2F&amp;sessionId=63e2a4145f43d9416c3a9132d4d53ccb7d2a81d7&amp;theme=light&amp;widgetsVersion=1890d59c%3A1627936082797&amp;width=550px" frameborder="0"></iframe>

> replaced `viper` with `flag` package and🤯.
>
> How do you justify adding a dependency if stdlib provides same functionality even if some plumbing required? [#golang](https://twitter.com/hashtag/golang?src=hash&ref_src=twsrc^tfw) [pic.twitter.com/4fAoXVP7vU](https://t.co/4fAoXVP7vU)
>
> — gmarik (@gmarik) 
>
> August 27, 2019

Surprisingly often, in order to fulfill [12 Factor](https://12factor.net) config requirements, people resort to packages with large API surface and as result large codebase and deep dependency graph.

Often times this is not necessary since the same functionality can be  achieved with much less code and only using Go’s standard library  packages. Here’s an example of using [`flag`](https://golang.org/pkg/flag/) package to achieve equal result.

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#12-factor-config-with-flag-package)12 factor config with `flag` package

- 2 ways to configure the app, through: 1) cli flags or 2) environment variables
- default values are configured from corresponding variables
- environment variables, if configured, set the `flag`'s defaults using `LookupOr*` helpers
- get full configuration with simple `getConfig` helper

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
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
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

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#conclusion)Conclusion

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#pros)Pros

- no dependencies other than standard library

### [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#cons)Cons

- a bit of plumbing code is required
- defaults to environment var’s value if latter is set
- env vars are manually named
- description may duplicate var’s comments

[`flag`](https://golang.org/pkg/flag/) package with combination with few helpers provides pragmatic way to configure your [12 factor](https://12factor.net)-ready apps. It’s not perfect but gets the job done.

## [¶ ](https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/#references)References

- [Simplicity Matters by Rich Hickey](https://www.youtube.com/watch?v=rI8tNMsozo0)

##### Related Posts

- [Monotonic Time, or Perfect model vs Imperfect Reality](https://www.gmarik.info/blog/2019/monotonic-time-perfect-model-vs-imperfect-reality/)
- [Testing MongoDB queries with Golang](https://www.gmarik.info/blog/2017/testing-mongodb-queries-golang/)
- [Wordfight: a multi-player word game](https://www.gmarik.info/blog/2017/wordfight-multiplayer-word-game/)
- [Understanding Go's `for` loop with closures](https://www.gmarik.info/blog/2016/understanding-golang-for-loop-with-closures/)
- [Experimenting with Go pipelines](https://www.gmarik.info/blog/2016/experimenting-with-golang-pipelines/)
