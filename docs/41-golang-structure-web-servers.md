# [How I Structure Web Servers in Go](https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers/)

Switching to Go from a decade of C# momentum has been an interesting journey. At times I revel in Go’s [simplicity](https://www.youtube.com/watch?v=rFejpH_tAHM); at other times frustration swells when familiar OOP [patterns](https://en.wikipedia.org/wiki/Software_design_pattern) don’t harmonize in a Go codebase. Fortunately, I’ve stumbled upon some  patterns for writing HTTP services that have been working well with my  team.

When working on corporate projects I tend to make  discoverability the highest priority. These applications could spend the next 20 years in production, having to be hot-patched, maintained, and  tweaked by untold legions of developers and site reliability engineers.  As such, I don’t expect these patterns to be one-size fits all.

An example repo is available [on GitHub](https://github.com/dudleycodes/golang-microservice-structure).

> [Mat Ryer's post](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html) was one of my starting points for experimenting with HTTP services in Go and the inspiration for this post.

## Code Composition

### The Broker

A `Broker` struct is the glue that binds distinct service packages to the HTTP  logic. No package-scoped variables are used. Interfaces of dependencies  are embedded to take advantage of [Go's composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html).

```go
type Broker struct {
    auth.Client             // authentication dependency imported from outside repository (interface)
    service.Service         // repository's business logic package (interface)

    cfg    Config           // the api service's configuration
    router *mux.Router      // the api service's route collection
}
```

The broker can be initialized with the [blocking](https://stackoverflow.com/questions/2407589/what-does-the-term-blocking-mean-in-programming) function `New()` which validates configurations and runs all the needed pre-flight checks.

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

```go
type Server interface {
    PingDependencies(bool) error
    ValidateJWT(string) error

    service.Service
}
```

The web service gets launched by calling the `Start()` function. Route binding is injected via a [closure function](https://gobyexample.com/closures), ensuring circular dependencies don’t break the import cycle.

```go
func (bkr *Broker) Start(binder func(s Server, r *mux.Router)) {
    ...

    bkr.router = mux.NewRouter().StrictSlash(true)
    binder(bkr, bkr.router)

    ...

    if err := http.Serve(l, bkr.router); errors.Is(err, http.ErrServerClosed) {
        log.Warn().Err(err).Msg("Web server has shut down")
    } else {
        log.Fatal().Err(err).Msg("Web server has shut down unexpectedly")
    }
}
```

Functions that could be useful in troubleshooting (e.g. checks used in [Kubernetes probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)) or disaster recovery scenarios hang from the `Broker`. These are only added to the `webserver.Server` interface if used by the routes/middleware.

```go
func (bkr *Broker) SetupDatabase() { ... }
func (bkr *Broker) PingDependencies(failFast bool)) { ... }
```

### Bootstrapping

The entry point to the whole application is a `main` package. By default it’ll start the web server. We can pass in some  command line arguments to call troubleshooter functions mentioned  earlier, handy for testing proxy permissions and other network oddities  using the validated configuration that was passed into `New()`. All we have to do is exec into a running pod and use them like any other command line tool.

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

```go
func BuildPipeline(srv webserver.Server, r *mux.Router) {
    r.Use(middleware.Metrics(), middleware.Authentication(srv))
    r.HandleFunc("/ping", routes.Ping()).Methods(http.MethodGet)

    ...

    r.HandleFunc("/makes/{makeID}/models/{modelID}", model.get(srv)).Methods(http.MethodGet)
}
```

### Middleware

Middleware returns a function that takes a handler which builds the needed `http.HandlerFunc`. This allows the `webserver.Server` interface to be injected and have all the safety checks only executed at start up rather than on every route invocation.

```go
func Authentication(srv webserver.Server) func(h http.Handler) http.Handler {
    if srv == nil || !srv.Client.IsValid() {
        log.Fatal().Msg("a nil dependency was passed to authentication middleware")
    }

    // additional setup logic
    ...

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := strings.TrimSpace(r.Header.Get("Authorization"))
            if err := srv.ValidateJWT(token); err != nil {
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

Routes have a similar footprint as middleware – a simpler setup but with the same benefits.

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

The directory structure is *highly* optimized for discoverability.

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

> Most importantly: each package is responsible for one thing and one thing only!

### HTTP Service Structure

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

## See Also

- [How I Organize Struct in Go Projects](https://www.dudley.codes/posts/2021.02.23-golang-struct-organization/)