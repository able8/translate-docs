# How to customize Go's HTTP client

# 如何自定义 Go 的 HTTP 客户端

Written on October 15, 2020 From: https://rafallorenz.com/go/customize-http-client/

In this article we are going to learn how to build extendable  HTTP client for our application while working with multiple external  APIs. We will take advantage of [http.Client](https://golang.org/pkg/net/http/#Client) and [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper). For the purpose of this article lets assume we are going to work with  only two different external APIs. Our HTTP client should let us:

在本文中，我们将学习如何在使用多个外部 API 的同时为我们的应用程序构建可扩展的 HTTP 客户端。我们将利用 [http.Client](https://golang.org/pkg/net/http/#Client) 和 [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper)。出于本文的目的，假设我们将只使用两个不同的外部 API。我们的 HTTP 客户端应该让我们：

1. allow us to handle HTTP errors using our application error types
2. allow to extend the logic with custom addons per each API

1. 允许我们使用我们的应用程序错误类型处理 HTTP 错误
2. 允许使用每个 API 的自定义插件扩展逻辑

# Custom HTTP client

# 自定义 HTTP 客户端

To take a full advantage of Go native libraries while tweaking it to our needs lets define our `Client` type as follow:

为了充分利用 Go 本地库，同时根据我们的需要对其进行调整，让我们定义我们的 `Client` 类型如下：

```
type Options struct {
    ApiURL          string
    Verbose         bool
}

type Client struct {
    httpClient *http.Client
    options    *Options
}

func New(httpClient *http.Client, options Options) *Client {
    return &Client{
        httpClient: httpClient,
        options:    &options,
    }
}
```


Doing so we can reuse our `Client` type for multiple external APIs by providing different options which holds `ApiURL` value and other configuration parameters. On top of that `Options` itself can be extended to hold much more complex configuration as  needed. Depending on your use case you might want to add more.

这样做我们可以通过提供包含 `ApiURL` 值和其他配置参数的不同选项来为多个外部 API 重用我们的 `Client` 类型。最重要的是，`Options` 本身可以根据需要进行扩展以容纳更复杂的配置。根据您的用例，您可能想要添加更多。

We want our client to conform to the following interface:

我们希望我们的客户端符合以下接口：

```
type HTTPClient interface {
    Get(ctx context.Context, path string, v interface{}) error
    Post(ctx context.Context, path string, payload interface{}, v interface{}) error
    Put(ctx context.Context, path string, payload interface{}, v interface{}) error
    Delete(ctx context.Context, path string, payload interface{}, v interface{}) error
}
```


As you can probably notice we want to be able to pass `path` with optional `payload` and `v` objects into which response body should be parsed. This will allow us  to implement quickly API calls that are specific for our service,  without unnecessary duplications. To keep this post short I am going to  implement one method:

您可能会注意到，我们希望能够将带有可选的 `payload` 和 `v` 对象的 `path` 传递给响应主体应该被解析到的对象。这将使我们能够快速实现特定于我们服务的 API 调用，而不会产生不必要的重复。为了使这篇文章简短，我将实现一种方法：

```
func (c *Client) Get(ctx context.Context, path string, v interface{}) error {
    req, err := c.newRequest(ctx, http.MethodGet, path, nil)
    if err != nil {
        return fmt.Errorf("failed to create GET request: %w", err)
    }

    if _, err := c.doRequest(req, v);err != nil {
        return err
    }

    return nil
}
```


Other methods should look very similar, with a small change of a HTTP method used. Internal implementation of our client is divided into two  parts:

其他方法看起来应该非常相似，只是对所使用的 HTTP 方法进行了小幅改动。我们客户端的内部实现分为两部分：

- new request
- do request

- 新请求
- 做请求

Notice that `GET` method has no payload therefore last value passed to `newRequest` is nil.

请注意，`GET` 方法没有有效负载，因此传递给 `newRequest` 的最后一个值为零。

Example usage of client methods by our service would look something like this:

我们的服务使用客户端方法的示例如下所示：

```
func (s *Service) CreateUser(ctx context.Context, u User) (User, error) {
    var user User
    if err := s.client.Post(ctx, "/users", u, &user);err != nil {
        return user, fmt.Errorf("failed to create user %+v: %w", u, err)
    }

    return user, nil
}

func (s *Service) GetUser(ctx context.Context, id string) (User, error) {
    var user User
    if err := s.client.Get(ctx, fmt.Sprintf("/users/%s", id), &user);err != nil {
        return user, fmt.Errorf("failed to get user %s: %w", id, err)
    }

    return user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
    if err := s.client.Delete(ctx, fmt.Sprintf("/users/%s", id), nil, nil);err != nil {
        return fmt.Errorf("failed to delete user %s: %w", id, err)
    }

    return nil
}
```


## New request

## 新请求

Implementation of this method is trivial, we need to simply create new HTTP request using [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) with body if payload is provided.

这个方法的实现很简单，如果提供了有效负载，我们只需使用 [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) 和 body 创建新的 HTTP 请求。

```
func (c *Client) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
var reqBody io.Reader
if payload != nil {
    bodyBytes, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal request body: %w", err)
    }
    reqBody = bytes.NewReader(bodyBytes)
}

req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.options.ApiURL, path), reqBody)
if err != nil {
    return nil, fmt.Errorf("failed to create HTTP request: %w", err)
}

if c.options.Verbose {
    body, _ := httputil.DumpRequest(req, true)
    log.Println(fmt.Sprintf("%s", string(body)))
}

req = req.WithContext(ctx)
return req, nil
}
```




At this stage we also set request’s context to given one, context will be responsible for timeouts etc.

在这个阶段，我们还将请求的上下文设置为给定的上下文，上下文将负责超时等。

## Do request

## 做请求

To keep our logic simple and readable in this stage lets handle  response parsing and defer request processing deeper. We could also at  this stage add here retry on error logic, before we parse response body.

为了在这个阶段保持我们的逻辑简单易读，让我们更深入地处理响应解析和延迟请求处理。在此阶段，我们还可以在解析响应正文之前在此处添加重试错误逻辑。

```
func (c *Client) doRequest(r *http.Request, v interface{}) error {
    resp, err := c.do(r)
    if err != nil {
        return err
    }

    if resp == nil {
        return nil
    }
    defer resp.Body.Close()

    if v == nil {
        return nil
    }

    var buf bytes.Buffer
    dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
    if err := dec.Decode(v);err != nil {
        return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
    }

    return nil
}
```


# Handling errors

# 处理错误

Most of the time, every application has defined its own error types, something like `ErrNotFound` or `ErrUserAccessDenied`. This allows us to leverage [errors.Is](https://golang.org/pkg/errors/#Is) method and apply custom behavior in an easy resilient way while  handling them. We already have our client type, lets implement logic to  do request and return proper application error.

大多数情况下，每个应用程序都定义了自己的错误类型，例如“ErrNotFound”或“ErrUserAccessDenied”。这使我们能够利用 [errors.Is](https://golang.org/pkg/errors/#Is) 方法并在处理它们时以一种简单的弹性方式应用自定义行为。我们已经有了我们的客户端类型，让我们实现逻辑来执行请求并返回正确的应用程序错误。

```
var (
    ErrUserAccessDenied  = errors.New("you do not have access to the requested resource")
    ErrNotFound          = errors.New("the requested resource not found")
    ErrTooManyRequests   = errors.New("you have exceeded throttle")
)

func (c *Client) do(r *http.Request) (*http.Response, error) {
    resp, err := c.httpClient.Do(r)
    if err != nil {
        return nil, fmt.Errorf("failed to make request [%s:%s]: %w", r.Method, r.URL.String(), err)
    }

    if c.options.Verbose {
        body, _ := httputil.DumpResponse(resp, true)
        log.Println(fmt.Sprintf("%s", string(body)))
    }

    switch resp.StatusCode {
    case http.StatusOK,
        http.StatusCreated,
        http.StatusNoContent:
        return resp, nil
    }

    defer resp.Body.Close()

    switch resp.StatusCode {
    case http.StatusNotFound:
        return ErrNotFound
    case http.StatusUnauthorized,
        http.StatusForbidden:
        return ErrUserAccessDenied
    case http.StatusTooManyRequests:
        return ErrTooManyRequests
    }

    return nil, fmr.Errorf("failed to do request, %d status code received", resp.StatusCode)
}
```


My method for the purpose of the example handles only couple of  request status codes, but your implementation can handle much more, or  maybe even you could introduce `RequestError` error type, which can hold information such as **request path, method, response status code** and maybe even **body**.

出于示例目的，我的方法仅处理几个请求状态代码，但您的实现可以处理更多，或者甚至您可以引入 `RequestError` 错误类型，它可以保存诸如 **请求路径、方法、响应等信息状态代码**，甚至可能是 **body**。

# Usage

#  用法

Now since we have our client implemented lets use it, at the very  beginning we have said that we are going to work with two different  external APIs. Each of them can have different/custom logic for  authorization, rate limit handling etc.

既然我们已经实现了我们的客户端，让我们使用它，一开始我们就说过我们将使用两个不同的外部 API。它们中的每一个都可以有不同的/自定义的授权逻辑、速率限制处理等。

Lets begin with service **A**, for the purpose of our example this service will provide as with an *API key* which we have to use to call their API. This key needs to be inserted in a custom header `X-AUTH-API-KEY`. To add custom (service specific logic) to our HTTP client we will use custom transport implementing ([http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) interface.

让我们从服务 **A** 开始，就我们的示例而言，该服务将提供一个 *API 密钥*，我们必须使用它来调用他们的 API。此密钥需要插入自定义标题“X-AUTH-API-KEY”中。要将自定义（特定于服务的逻辑）添加到我们的 HTTP 客户端，我们将使用自定义传输实现 ([http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) 接口。

```
const (
    apiVersion = "v1"
    baseUrl    = "https://api.serbvice-a.com"
)

type Service struct {
    client *client.Client
}

func New(key string) *Service {
    t := transport{
        apiKey: key,
    }

    return &Service{
        client: client.New(
            &http.Client{Transport: &t},
            client.Options{
                ApiURL:  fmt.Sprintf("%s/%s", baseUrl, apiVersion),
                Verbose: os.Getenv("VERBOSE") != "",
            },
        ),
    }
}

type transport struct {
    apiKey    string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
    r := req.Clone(req.Context())
    r.Header.Add("Content-Type", "application/json")
    r.Header.Add("Accept", "application/json")
    r.Header.Add("Accept-Charset", "UTF-8")

    r.Header.Add("X-AUTH-API-KEY", t.apiKey)

    return http.DefaultTransport.RoundTrip(r)
}
```


Service **B** could for example be using `oauth2` protocol for an authentication, in this case our implementation could look as follow:

例如，服务 **B** 可以使用 `oauth2` 协议进行身份验证，在这种情况下，我们的实现可能如下所示：

```
const (
    apiVersion = "v1"
    baseUrl    = "https://api.serbvice-b.com"
)

var (
    authConfig = oauth2.Config{
        ClientID:     "XXXX-XXXX-XXXX-XXXX",
        ClientSecret: "YYYY-YYYY-YYYY-YYYY",
        RedirectURL:  "https://api.our-service.com/oauth/callback",
        Scopes:       []string{"all"},
        Endpoint: oauth2.Endpoint{
            AuthStyle: oauth2.AuthStyleInParams,
            AuthURL:   "https://api.serbvice-b.com/oauth/authorize",
            TokenURL:  "https://api.serbvice-b.com/oauth/access_token",
        },
    }
)

type Service struct {
    client *client.Client
}

func New(t *Token) *Service {
    httpClient := authConfig.Client(
        context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: &transport{}}),
        t,
    )

    return &Service{
        client: client.New(
            httpClient,
            client.Options{
                ApiURL:  fmt.Sprintf("%s/%s", baseUrl, apiVersion),
                Verbose: os.Getenv("VERBOSE") != "",
            },
        ),
    }
}

type transport struct {}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
    r := req.Clone(req.Context())
    r.Header.Add("Content-Type", "application/json")
    r.Header.Add("Accept", "application/json")

    return http.DefaultTransport.RoundTrip(r)
}
```




Service **B** custom transport sets headers for content type only, because we are using *oauth2* HTTP client, they already append token and handle auto-refresh as necessary for us.

服务 **B** 自定义传输仅为内容类型设置标头，因为我们使用 *oauth2* HTTP 客户端，它们已经附加令牌并根据需要处理自动刷新。

# Conclusion

#  结论

We have learned today how to take advantage of Go's [http.Client](https://golang.org/pkg/net/http/#Client)and leverage [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) to apply custom logic to our customized HTTP client.

今天我们学习了如何利用 Go 的 [http.Client](https://golang.org/pkg/net/http/#Client)和利用 [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) 将自定义逻辑应用于我们自定义的 HTTP 客户端。

With the help of code snippets presented in this post, I hope you  have grasped the idea of how to customize HTTP client and apply custom,  service specific logic. If the API provider would return header with a  retry delay time would you be able to add rate limit logic? Try to do it yourself, if you need help leave a comment and lets play in Go’s  playground.

借助本文中提供的代码片段，我希望您已经掌握了如何自定义 HTTP 客户端并应用自定义的服务特定逻辑的想法。如果 API 提供者将返回带有重试延迟时间的标头，您是否能够添加速率限制逻辑？尝试自己做，如果您需要帮助，请发表评论并让我们在围棋的操场上玩耍。

 



