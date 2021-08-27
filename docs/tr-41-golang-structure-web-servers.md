# [How I Structure Web Servers in Go](https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers/)

# [我如何在 Go 中构建 Web 服务器](https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers/)

Switching to Go from a decade of C# momentum has been an interesting journey. At times I revel in Go’s [simplicity](https://www.youtube.com/watch?v=rFejpH_tAHM);at other times frustration swells when familiar OOP [patterns](https://en.wikipedia.org/wiki/Software_design_pattern) don’t harmonize in a Go codebase. Fortunately, I’ve stumbled upon some  patterns for writing HTTP services that have been working well with my  team.

从十年的 C# 势头转向 Go 是一段有趣的旅程。有时我陶醉于 Go 的 [简单](https://www.youtube.com/watch?v=rFejpH_tAHM)；有时，当熟悉的OOP [模式](https://en.wikipedia.org/wiki/Software_design_pattern) 在 Go 代码库中不协调时，挫败感会膨胀。幸运的是，我偶然发现了一些编写 HTTP 服务的模式，这些模式在我的团队中运行良好。

When working on corporate projects I tend to make  discoverability the highest priority. These applications could spend the next 20 years in production, having to be hot-patched, maintained, and  tweaked by untold legions of developers and site reliability engineers. As such, I don’t expect these patterns to be one-size fits all.

在处理公司项目时，我倾向于将可发现性放在首位。这些应用程序可能会在接下来的 20 年中投入生产，必须由无数开发人员和站点可靠性工程师进行热修补、维护和调整。因此，我不希望这些模式是一刀切的。

An example repo is available [on GitHub](https://github.com/dudleycodes/golang-microservice-structure).

[GitHub 上](https://github.com/dudleycodes/golang-microservice-structure) 提供了一个示例存储库。

> [Mat Ryer's post](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html) was one of my starting points for experimenting with HTTP services in Go and the inspiration for this post.

> [Mat Ryer 的帖子](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html) 是我尝试的起点之一Go 中的 HTTP 服务以及这篇文章的灵感。

## Code Composition

## 代码组成

### The Broker

### 经纪人

A `Broker` struct is the glue that binds distinct service packages to the HTTP  logic. No package-scoped variables are used. Interfaces of dependencies  are embedded to take advantage of [Go's composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html).

`Broker` 结构是将不同的服务包绑定到 HTTP 逻辑的粘合剂。不使用包范围的变量。嵌入了依赖的接口以利用 [Go 的组合](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)。

```go
type Broker struct {
    auth.Client             // authentication dependency imported from outside repository (interface)
    service.Service         // repository's business logic package (interface)

    cfg    Config           // the api service's configuration
    router *mux.Router      // the api service's route collection
}
```


The broker can be initialized with the [blocking](https://stackoverflow.com/questions/2407589/what-does-the-term-blocking-mean-in-programming)function `New()` which validates configurations and runs all the needed pre-flight checks.

可以使用 [blocking](https://stackoverflow.com/questions/2407589/what-does-the-term-blocking-mean-in-programming)函数 New() 来初始化代理，该函数验证配置并运行所有需要的飞行前检查。

```go
func New(cfg Config, port int) (*Broker, error) {
    r := &Broker{
        cfg: cfg,
    }

    ...

    r.auth.Client, err = auth.New(cfg.AuthConfig)
    if err != nil {
        return nil, fmt.Errorf("Unable to create new API broker: %w", err)
    }

    ...

    return r, nil
}
```


The initialized `Broker` fulfills the exposed `Server` interface which defines all functionality that can be used by the routes and middleware. The `service` package interface gets embedded, matching the embedded interface on the `Broker` struct.

初始化的 `Broker` 实现了暴露的 `Server` 接口，该接口定义了路由和中间件可以使用的所有功能。 `service` 包接口被嵌入，匹配 `Broker` 结构上的嵌入接口。

```go
type Server interface {
    PingDependencies(bool) error
    ValidateJWT(string) error

    service.Service
}
```


The web service gets launched by calling the `Start()` function. Route binding is injected via a [closure function](https://gobyexample.com/closures), ensuring circular dependencies don’t break the import cycle.

Web 服务通过调用 `Start()` 函数启动。路由绑定通过 [闭包函数](https://gobyexample.com/closures) 注入，确保循环依赖不会破坏导入循环。

```go
func (bkr *Broker) Start(binder func(s Server, r *mux.Router)) {
    ...

    bkr.router = mux.NewRouter().StrictSlash(true)
    binder(bkr, bkr.router)

    ...

    if err := http.Serve(l, bkr.router);errors.Is(err, http.ErrServerClosed) {
        log.Warn().Err(err).Msg("Web server has shut down")
    } else {
        log.Fatal().Err(err).Msg("Web server has shut down unexpectedly")
    }
}
```


Functions that could be useful in troubleshooting (eg checks used in [Kubernetes probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)) or disaster recovery scenarios hang from the `Broker`. These are only added to the `webserver.Server` interface if used by the routes/middleware.

可用于故障排除的功能（例如在 [Kubernetes 探针](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) 中使用的检查）或灾难恢复场景挂在 `Broker` 上。如果路由/中间件使用，这些只会添加到`webserver.Server` 接口。

```go
func (bkr *Broker) SetupDatabase() { ... }
func (bkr *Broker) PingDependencies(failFast bool)) { ... }
```


### Bootstrapping

### 引导

The entry point to the whole application is a `main` package. By default it’ll start the web server. We can pass in some  command line arguments to call troubleshooter functions mentioned  earlier, handy for testing proxy permissions and other network oddities  using the validated configuration that was passed into `New()`. All we have to do is exec into a running pod and use them like any other command line tool.

整个应用程序的入口点是一个 `main` 包。默认情况下，它将启动 Web 服务器。我们可以传入一些命令行参数来调用前面提到的疑难解答函数，这对于使用传递到 `New()` 的经过验证的配置来测试代理权限和其他网络异常很方便。我们所要做的就是 exec 到一个正在运行的 pod 中，并像使用任何其他命令行工具一样使用它们。

```go
func main() {
    subCommand := flag.String("start", "", "start the webserver")

    ...

    srv := webserver.New(cfg, 80)

    switch strings.ToLower(subCommand) {
    case "ping":
        srv.PingDependencies(false)
    case "start":
        srv.Start(BuildPipeline)
    default:
        fmt.Printf("Unrecognized command %q, exiting.", subCommand)
        os.Exit(1)
    }
}
```


The HTTP pipeline setup is done in the `BuildPipeline()` function, which will get injected into the server via `srv.Start()`.

HTTP 管道设置在 `BuildPipeline()` 函数中完成，它将通过 `srv.Start()` 注入服务器。

```go
func BuildPipeline(srv webserver.Server, r *mux.Router) {
    r.Use(middleware.Metrics(), middleware.Authentication(srv))
```去

r.HandleFunc("/ping", routes.Ping()).Methods(http.MethodGet)

     r.HandleFunc("/ping", routes.Ping()).Methods(http.MethodGet)

    ...

     ...

    r.HandleFunc("/makes/{makeID}/models/{modelID}", model.get(srv)).Methods(http.MethodGet)
}
```

r.HandleFunc("/makes/{makeID}/models/{modelID}", model.get(srv)).Methods(http.MethodGet)
}
``

### Middleware

### 中间件

Middleware returns a function that takes a handler which builds the needed `http.HandlerFunc`. This allows the `webserver.Server` interface to be injected and have all the safety checks only executed at start up rather than on every route invocation.

中间件返回一个函数，该函数接受一个构建所需的 `http.HandlerFunc` 的处理程序。这允许注入 `webserver.Server` 接口，并且所有安全检查只在启动时执行，而不是在每次路由调用时执行。

```go
func Authentication(srv webserver.Server) func(h http.Handler) http.Handler {
    if srv == nil ||!srv.Client.IsValid() {
        log.Fatal().Msg("a nil dependency was passed to authentication middleware")
    }

    // additional setup logic
    ...

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := strings.TrimSpace(r.Header.Get("Authorization"))
            if err := srv.ValidateJWT(token);err != nil {
                ...
                w.WriteHeader(401)
                w.Write([]byte("Access Denied"))

                return
            }

            next.ServeHTTP(w, r)
        }
    }
}
```


### Routes

### 路线

Routes have a similar footprint as middleware – a simpler setup but with the same benefits.

路由具有与中间件类似的足迹——设置更简单，但具有相同的好处。

```go
func GetLatest(srv webserver.Server) http.HandlerFunc {
    if srv == nil {
        log.Fatal().Msg("a nil dependency was passed to the `/makes/{makeID}/models/{modelID}` route")
    }

    // additional setup logic
    ...

    return func(w http.ResponseWriter, r *http.Request) {
        ...

        makeDTO, err := srv.Get
    }
}
```


## Directory Structure

## 目录结构

The directory structure is *highly* optimized for discoverability.

目录结构针对可发现性进行了*高度*优化。

```text
├── app/
|   └── service-api/**
├── cmd/
|   └── service-tool-x/
├── internal/
|   └── service/
|       └── mock/
├── pkg/
|   ├── client/
|   └── dtos/
├── (.editorconfig, .gitattributes, .gitignore)
└── go.mod
```



- **app/** is for the project applications - this is the entry point newcomers gravitate towards when exploring the codebase.
- **./service-api/** is the micro-service API for this repository; all HTTP implementation details live here.
- **cmd/** is where any command-line applications belong.
- internal/ is a special directory that cannot be imported by projects outside of this repo.
- **./service/** is where all the domain logic goes; from there it can be imported by `service-api`, `service-tool-x`, and any future applications/packages that would benefit from accessing it directly.
- **pkg/** is for any packages that are encouraged to be imported by projects outside this repo.
- **./client/** is a client library for accessing `service-api`. Other teams can import it without having to write their own and we can “[dogfood it](https://en.wikipedia.org/wiki/Eating_your_own_dog_food)” with our own CI/CD tools stored in `cmd/`.
- **./dtos/** is for the project *d*ata *t*ransfer *o*bjects, structs designed for sharing data between packages and  encoding/transmitting over the wire as json. No model-like structs are  exported from any other repo packages. `/internal/service` is responsible for mapping the DTOs to/from its internal models,  preventing implementation details from leaking (e.g. database  annotations) and allows the models to change without breaking downstream clients consuming the DTOs.
- **.editorconfig, .gitattributes, .gitignore** Because [all repos should use .editorconfig, .gitattributes, .gitignore](https://www.dudley.codes/posts/2020.02.16-git-lost-in-translation/)!
- **go.mod** even works inside [restrictive and bureaucratic corporate environments](https://www.dudley.codes/posts/2020.04.02-golang-behind-corporate-firewall/).

- **app/** 用于项目应用程序 - 这是新人在探索代码库时所趋向的切入点。
- **./service-api/** 是这个仓库的微服务API；所有 HTTP 实现细节都在这里。
- **cmd/** 是任何命令行应用程序所属的地方。
- internal/ 是一个特殊目录，不能被这个 repo 之外的项目导入。
- **./service/** 是所有域逻辑的所在；从那里它可以通过`service-api`、`service-tool-x` 以及任何将来可以从直接访问它中受益的应用程序/包导入。
- **pkg/** 用于鼓励此 repo 之外的项目导入的任何包。
- **./client/** 是访问`service-api`的客户端库。其他团队无需自己编写就可以导入它，我们可以使用存储在 `cmd/` 中的我们自己的 CI/CD 工具“[dogfood it](https://en.wikipedia.org/wiki/Eating_your_own_dog_food)”。
- **./dtos/** 用于项目 *d*ata *t*ransfer *o*bjects，结构设计用于在包之间共享数据以及作为 json 在线编码/传输。没有从任何其他 repo 包中导出类似模型的结构。 `/internal/service` 负责将 DTO 映射到/从其内部模型映射，防止实现细节泄漏（例如数据库注释），并允许模型更改而不会破坏使用 DTO 的下游客户端。
- **.editorconfig, .gitattributes, .gitignore** 因为[所有 repos 都应该使用 .editorconfig, .gitattributes, .gitignore](https://www.dudley.codes/posts/2020.02.16-git-lost-in -翻译/）！
- **go.mod** 甚至可以在 [限制性和官僚主义的企业环境] 中工作（https://www.dudley.codes/posts/2020.04.02-golang-behind-corporate-firewall/）。

> Most importantly: each package is responsible for one thing and one thing only!

> 最重要的是：每个包裹只负责一件事，只负责一件事！

### HTTP Service Structure

### HTTP 服务结构

```text
└── service-api/
    ├── cfg/
    ├── middleware/
    ├── routes/
    |   ├── makes/
    |   |   └── models/**
    |   ├── create.go
    |   ├── create_test.go
    |   ├── get.go
    |   └── get_test.go
    ├── webserver/
    ├── main.go
    └── pipeline.go
```


- **./cfg/** is for configuration  files, usually json or yaml saved in plain text files, as they should be checked into git too (except for passwords, private keys, etc).
- **./middleware** for all middleware.
- **./routes** routes get grouped and nested using directories that mirror the API application’s RESTFul-like surface.
- **./webserver** contains all shared HTTP structs and interfaces (`Broker`, configuration, `Server`, etc).
- **main.go** where the application is bootstrapped (`New()`, `Start()`).
- **pipeline.go** where the `BuildPipeline()` function lives.

- **./cfg/** 用于配置文件，通常是 json 或 yaml 保存在纯文本文件中，因为它们也应该被检入 git（密码、私钥等除外）。
- **./middleware** 用于所有中间件。
- **./routes** 使用与 API 应用程序类似 RESTFul 的表面的目录对路由进行分组和嵌套。
- **./webserver** 包含所有共享的 HTTP 结构和接口（`Broker`、配置、`Server` 等）。
- **main.go** 应用程序启动的地方（`New()`、`Start()`）。
- **pipeline.go** `BuildPipeline()` 函数所在的位置。

## See Also 
##  也可以看看
- [How I Organize Struct in Go Projects](https://www.dudley.codes/posts/2021.02.23-golang-struct-organization/) 
- [我如何在 Go 项目中组织结构](https://www.dudley.codes/posts/2021.02.23-golang-struct-organization/)
