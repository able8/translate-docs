# Structuring and testing HTTP handlers in Go

Written 15th of October 2020 From: https://www.maragu.dk/blog/structuring-and-testing-http-handlers-in-go/

There are many ways to structure your HTTP handlers in your web application code in Go. It would be nice to have a default way to do this that makes it easy to:
- Inject your dependencies, to make the handlers and the rest of your code loosely coupled
- See which route paths go to which handlers, and have them close together in code, for readability
- Unit test the handlers in isolation

After quite a few different designs, I've found a way I like, and in this post, I'll show you.

If you want to check out a simple project implementing this, see[github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing).

## The handler

I'll start by showing you the design, and then breaking it down. A handler generally looks like this:

```go
package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

type partyStarterRepo interface {
	StartParty(id string) error
}

func PartyStarter(mux chi.Router, s partyStarterRepo) {
	mux.Post("/start/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := s.StartParty(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}
```

### Request multiplexer as parameter

The `PartyStarter` function takes the request multiplexer `mux` (in this case [chi](https://github.com/go-chi/chi), but use any you like) as the first parameter. This means that the handler registers itself, including defining the route and its parameters. It's nice to have this very close to the handler code, both for increased readability, but also that it's very clear that they belong together and should be changed together. For example, if the `id` parameter changes in name or content, the code right below should reflect that.

### Dependency as private interface parameter

The business logic dependency is passed as an interface, `partyStarterRepo`, that is defined specifically for this handler. We can do this in Go because interfaces are implicit, meaning that anything that has a method with signature `StartParty(id string) error` can be passed to this function. We will use this in testing.

This enables us to define exactly what this handler needs from its dependencies, nothing more, nothing less. So if your dependency has a lot of extra functionality (for example, a`StopParty` function), this handler doesn't know about it.

### Handlers in a separate package

To isolate the handlers, they are in a separate package called `handlers`. Note that because of the use of private interfaces for dependencies, we don't import our business logic packages in the handlers. This reduces coupling, and makes it easier to swap the underlying dependencies, for example.

## Testing the handler

To test the handler, we can use the `httptest` package from the standard library, along with a very small mock for the dependency.

```go
package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

type partyStarterRepoMock struct {
	err error
}

func (m *partyStarterRepoMock) StartParty(id string) error {
	return m.err
}

func TestPartyStarter(t *testing.T) {
	t.Run("sends bad gateway on start party error", func(t *testing.T) {
		mux := chi.NewMux()
		PartyStarter(mux, &partyStarterRepoMock{err: errors.New("no snacks")})

		req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		res := rec.Result()
		if res.StatusCode != http.StatusBadGateway {
			t.FailNow()
		}
	})

	t.Run("sends accepted on start party success", func(t *testing.T) {
		mux := chi.NewMux()
		PartyStarter(mux, &partyStarterRepoMock{})

		req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		res := rec.Result()
		if res.StatusCode != http.StatusAccepted {
			t.FailNow()
		}
	})
}
```

See how the mock is tiny, because we're testing only exactly what this handler needs? No more autogenerating huge mocks with all your business logic functions on it.

Also note that we don't have to start our server to check that our routes work as expected, because the routes are right there in the handler.

## Conclusion

In this post, I've shown you how to structure your HTTP handlers in Go so they are loosely coupled with their dependencies, using private interfaces, and easy to test, using routes that are defined inside the handler. To see a simple project showing you all of this, check out [github.com/maragudk/http-handler-testing](https://github.com/maragudk/http-handler-testing).

If you have any questions or comments about this, feel free to reach out to me on Twitter. I'm at [@markusrgw](https://twitter.com/markusrgw).

## About me

I'm Markus, a professional software consultant and developer. 🤓✨ You can reach me at [markus@maragu.dk](mailto:Markus from maragu) or [@markusrgw](https://twitter.com/markusrgw).

I'm currently [building Go courses over at golang.dk](https://www.golang.dk/).