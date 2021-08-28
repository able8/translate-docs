

# How to customize Go's HTTP client

Written on October 15, 2020 From: https://rafallorenz.com/go/customize-http-client/

In this article we are going to learn how to build extendable  HTTP client for our application while working with multiple external  APIs. We will take advantage of [http.Client](https://golang.org/pkg/net/http/#Client) and [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper). For the purpose of this article lets assume we are going to work with  only two different external APIs. Our HTTP client should let us:

1. allow us to handle HTTP errors using our application error types
2. allow to extend the logic with custom addons per each API

# Custom HTTP client

To take a full advantage of Go native libraries while tweaking it to our needs lets define our `Client` type as follow:

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

We want our client to conform to the following interface:

```
type HTTPClient interface {
	Get(ctx context.Context, path string, v interface{}) error
	Post(ctx context.Context, path string, payload interface{}, v interface{}) error
	Put(ctx context.Context, path string, payload interface{}, v interface{}) error
	Delete(ctx context.Context, path string, payload interface{}, v interface{}) error
}
```

As you can probably notice we want to be able to pass `path` with optional `payload` and `v` objects into which response body should be parsed. This will allow us  to implement quickly API calls that are specific for our service,  without unnecessary duplications. To keep this post short I am going to  implement one method:

```
func (c *Client) Get(ctx context.Context, path string, v interface{}) error {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	if _, err := c.doRequest(req, v); err != nil {
		return err
	}

	return nil
}
```

Other methods should look very similar, with a small change of a HTTP method used. Internal implementation of our client is divided into two  parts:

- new request
- do request

Notice that `GET` method has no payload therefore last value passed to `newRequest` is nil.

Example usage of client methods by our service would look something like this:

```
func (s *Service) CreateUser(ctx context.Context, u User) (User, error) {
	var user User
	if err := s.client.Post(ctx, "/users", u, &user); err != nil {
		return user, fmt.Errorf("failed to create user %+v: %w", u, err)
	}

	return user, nil
}

func (s *Service) GetUser(ctx context.Context, id string) (User, error) {
	var user User
	if err := s.client.Get(ctx, fmt.Sprintf("/users/%s", id), &user); err != nil {
		return user, fmt.Errorf("failed to get user %s: %w", id, err)
	}

	return user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/users/%s", id), nil, nil); err != nil {
		return fmt.Errorf("failed to delete user %s: %w", id, err)
	}

	return nil
}
```

## New request

Implementation of this method is trivial, we need to simply create new HTTP request using [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) with body if payload is provided.

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

## Do request

To keep our logic simple and readable in this stage lets handle  response parsing and defer request processing deeper. We could also at  this stage add here retry on error logic, before we parse response body.

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
	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
	}

	return nil
}
```

# Handling errors

Most of the time, every application has defined its own error types, something like `ErrNotFound` or `ErrUserAccessDenied`. This allows us to leverage [errors.Is](https://golang.org/pkg/errors/#Is) method and apply custom behavior in an easy resilient way while  handling them. We already have our client type, lets implement logic to  do request and return proper application error.

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

# Usage

Now since we have our client implemented lets use it, at the very  beginning we have said that we are going to work with two different  external APIs. Each of them can have different/custom logic for  authorization, rate limit handling etc.

Lets begin with service **A**, for the purpose of our example this service will provide as with an *API key* which we have to use to call their API. This key needs to be inserted in a custom header `X-AUTH-API-KEY`. To add custom (service specific logic) to our HTTP client we will use custom transport implementing ([http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) interface.

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

# Conclusion

We have learned today how to take advantage of Go’s [http.Client](https://golang.org/pkg/net/http/#Client) and leverage [http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) to apply custom logic to our customized HTTP client.

With the help of code snippets presented in this post, I hope you  have grasped the idea of how to customize HTTP client and apply custom,  service specific logic. If the API provider would return header with a  retry delay time would you be able to add rate limit logic? Try to do it yourself, if you need help leave a comment and lets play in Go’s  playground.

​     