# Web API Integration Testing with Go

 Tuesday. June 11, 2019 -  31 mins  

I’m learning Go by building a small API-backed web application, and wanted  to share the process in case it helps someone else. In this post, we’ll  continue where we left off last time with the [Go web API](https://rshipp.com/go-web-api) for managing [GitHub stars](https://help.github.com/en/articles/about-stars), adding automated tests to ensure our code functions as expected. If  you’d like to follow along with this post without going through the  previous one, you can grab a copy of the API (`main.go`) from [this GitHub repo](https://github.com/rshipp/StarManager/tree/0d7cdd291711c7dd6a706bc38844f250343c7b1f).

We’ll start by writing [integration tests](https://en.wikipedia.org/wiki/Integration_testing), which will rely on a real database backend rather than a [mock](https://en.wikipedia.org/wiki/Mock_object) or [stub](https://en.wikipedia.org/wiki/Unit_testing). Compared to [unit tests](https://en.wikipedia.org/wiki/Unit_testing), integration tests have a few drawbacks and benefits for our purposes.

Drawbacks:

1. Using a real database makes integration tests slower than unit tests that rely on stubbed methods or mocked interfaces.
2. Testing the full stack at once instead of each small piece at a time can make it harder to tell where bugs are coming from.

Benefits:

1. Integration tests allow us to ensure our SQL queries and schema are correct  (especially important here because that’s where most of our API’s  functionality comes from).
2. Since we don’t have to use stubs or mocks, we can write less code for the tests and get them up and running faster.

We will eventually want unit tests, for the reasons mentioned above, but  starting with integration gives us the biggest return for the time  being.

## What You’ll Need

Before we get started, you’ll need a few things:

- [Go](https://en.wikipedia.org/wiki/Code_smell) installed on your computer.
- The [Go web API](https://github.com/rshipp/StarManager/tree/0d7cdd291711c7dd6a706bc38844f250343c7b1f) (`main.go`) from the [previous post](https://rshipp.com/go-web-api).
- A text editor.

You may also want to skim through the [tour of Go](https://tour.golang.org/welcome/1) or another introduction to the Go language if you haven’t already, though you shouldn’t need more than a basic understanding.

If you run into anything unclear in this post, feel free to [open an issue](https://github.com/rshipp/rshipp.github.io/issues) on GitHub and let me know!

## Create Handler

Let’s go through and write an integration test function for each of our five HTTP request handlers from `main.go`, starting with the Create handler.

Go tests are expected to be in a file ending with `_test.go`. Since we’re writing tests for `main.go`, the conventional test file name is `main_test.go`. (Make sure `main_test.go` and `main.go` are in the same folder.)

Set up a basic outline for our first test in `main_test.go`:

```gogo
package main

import (
	"testing"
)

func setup() *App {
	// Initialize an in-memory database for full integration testing.
	app := &App{}
	app.Initialize("sqlite3", ":memory:")
	return app
}

func teardown(app *App) {
	// Closing the connection discards the in-memory database.
	app.DB.Close()
}

func TestCreateHandler(t *testing.T) {
	app := setup()

	// Test body will be here!

	teardown(app)
}
```

The documentation for the Go [testing](https://golang.org/pkg/testing/#pkg-overview) package goes over the requirements for test functions: they must each be named like `TestXxx` (where `X` is a capital letter), and have an argument `t *testing.T`.

There are no special “setup” or “teardown” functions as we might be used to  from other languages, nor is there a testing class with instance  variables we can use to access our database. To get around this, we  define a `setup` function to initialize an [in-memory SQLite database](https://sqlite.org/inmemorydb.html) and return an `App` pointer (we defined `App` in `main.go`), and `teardown` to accept that same pointer and close our database connection. We then call `app := setup()` at the start of every integration test function, and `teardown(app)` at the end of each function. This ensures that our database is always  clean and in a consistent state at the start of each test.

If we run the tests with `go test`, they should pass:

```go
PASS
ok      _/<...>/StarManager     0.004s
```

Let’s start filling out the test body for `TestCreateHandler`:

```go
	testStar := &Star{
		ID:          1,
		Name:        "test/name",
		Description: "test desc",
		URL:         "test url",
	}

	// Transform Star record into *strings.Reader suitable for use in HTTP POST forms.
	data := url.Values{
		"name":        {testStar.Name},
		"description": {testStar.Description},
		"url":         {testStar.URL},
	}

	form := strings.NewReader(data.Encode())

	// Set up a new request.
	req, err := http.NewRequest("POST", "/stars", form)
	if err != nil {
		t.Fatal(err)
	}
	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	http.HandlerFunc(app.CreateHandler).ServeHTTP(rr, req)
```

We’re using a few new Go packages here, so we’ll have to add them to the import list at the top of `main_test.go`:

```go
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)
```

The [url.Values.Encode()](https://golang.org/pkg/net/url/#Values) and [strings.NewReader](https://golang.org/pkg/strings/#NewReader) code takes our `testStar` record and converts it into the string format that [http.NewRequest](https://golang.org/pkg/net/http/#NewRequest) expects. We then use that `form` string with `http.NewRequest` to set up a request to the Create endpoint defined in `main.go`, `POST /stars`. Note that we are not actually making an HTTP request here, just preparing one.

This is where Go’s built-in `httptest` package comes in handy: we set up a [httptest.ResponseRecorder](https://golang.org/pkg/net/http/httptest/#ResponseRecorder), then pass it in to `http.HandlerFunc().ServeHTTP()`. With both this response record and the request we prepared earlier, we can test our `app.CreateHandler()` directly, without needing to set up a local HTTP server or client. In  essence, we’re passing variables around in Go’s internal functions  without using network requests or responses at all.

If there was an unexpected error forming the request, we can call [t.Fatal](https://golang.org/pkg/testing/#T.Fatal) to stop executing the test function immediately and print the error message.

With our POST request “sent” to our handler, and the response recorded in `rr`, we can continue filling out `TestCreateHandler`, checking that our API works as expected:

```go
	// Test that the status code is correct.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusCreated, status)
	}

	// Test that the Location header is correct.
	expectedURL := fmt.Sprintf("/stars/%s", testStar.Name)
	if location := rr.Header().Get("Location"); location != expectedURL {
		t.Errorf("Location header is invalid. Expected %s. Got %s instead", expectedURL, location)
	}

	// Test that the created star is correct.
	// Note: There is only one star in the database.
	createdStar := Star{}
	app.DB.First(&createdStar)
	if createdStar != *testStar {
		t.Errorf("Created star is invalid. Expected %+v. Got %+v instead", testStar, createdStar)
	}
```

First up, we expect to see `201 Created` as the status code. Go provides some nicely named [aliases](https://golang.org/pkg/net/http/#pkg-constants) for HTTP status codes, so we can reference this as `http.StatusCreated`, and compare it to the actual response code we got in `rr.Code`. If they’re different, we use [t.Errorf](https://golang.org/pkg/testing/#T.Errorf) to print out a useful message and then fail the test.

Next, the `Location` header: we expect this to be set to the URL of the star that was just created, which we defined in `main.go` as `/stars/{star.Name}`. The actual header is in `rr.Header()`, so we can compare that to the expected URL to verify correctness.

Finally, we want to see if a star was actually created. Since we’re working with a real database, we can use our [GORM](http://gorm.io/) database pointer `app.DB` directly to fetch the first (and only) star in the database, and compare it to our original `testStar`. (Using GORM directly like this is a bit of a [code smell](https://en.wikipedia.org/wiki/Code_smell), but we’ll worry about refactoring later.)

When we run `go test` again, it should report a `PASS`. Looks like our Create handler passed the test!

For a more complete project, we’d want to have additional tests for edge  cases: what happens if we try to create a star that already exists? what if the database is down? are there invalid characters that break the  SQL query? Our API in `main.go` is  pretty naive right now, so a lot of these will probably fail in  unexpected ways. As we build more functionality into the API, we’ll  continue to add integration and unit tests that make sure everything  works correctly.

## Update Handler

The  Update handler, like the Create handler, expects a form-encoded request  body with star attributes. The code we’re using to do that is a little  fragile (we may have to manually update it when we add new fields to the Star struct), so let’s start by refactoring it out into a function  inside `main_test.go` so we only have to maintain it in once place:

```go
func StarFormValues(star Star) *strings.Reader {
	// Transforms Star record into *strings.Reader suitable for use in HTTP POST forms.
	data := url.Values{
		"name":        {star.Name},
		"description": {star.Description},
		"url":         {star.URL},
	}

	return strings.NewReader(data.Encode())
}
```

In the Create handler, remove the code we pulled out into `StarFormValues`, and update the `NewRequest` call:

```go
	// Set up a new request.
	req, err := http.NewRequest("POST", "/stars", StarFormValues(*testStar))
```

Now we can reuse that function in the Update handler too.

Since we want to test updating a star, let’s start by putting one in the database:

```go
func TestUpdateHandler(t *testing.T) {
	app := setup()

	// Create a star for us to update.
	testStar := &Star{
		ID:          1,
		Name:        "test/name",
		Description: "test desc",
		URL:         "test url",
	}
	app.DB.Create(testStar)
```

There are two main things we want to test here: updating a star’s name (which changes the URL used to  reference it), and updating other fields in a star. We could do this  manually, but that sounds like a lot of duplicated code. Luckily, Go has a pattern called [table-driven tests](https://github.com/golang/go/wiki/TableDrivenTests) that will save us a lot of effort.

The basic pattern for a table-driven test looks something like this:

```go
myTests = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 4},
	{4, 16},
}

for _, tt := range myTests {
	if actual := MySquareFunction(tt.input); actual != tt.expected {
		t.Errorf("MySquareFunction(%d): expected: %d, actual: %d", tt.input, tt.expected, actual)
	}
}
```

The “table” is a [slice](https://gobyexample.com/slices) of anonymous (unnamed) [structs](https://gobyexample.com/structs). We can define as many fields as we need - in this case `input` and `expected` - and fill the slice with as many records as we want with different  values for those fields. Here, we have 3 records, each representing `{input, expected}`. We then loop over the table with `range` (`tt` is the conventional name for table-driven test elements, but you could  call it something else if you wanted), and run the same test on each  record in the table.

Applying this to our use case, back in `TestUpdateHandler` in `main_test.go`, we can set up a table of stars:

```go
	// Set up a test table.
	starTests := []struct {
		original Star
		update   Star
	}{
		{original: *testStar,
			update: Star{ID: 1, Name: "test/name", Description: "updated desc", URL: "test URL"},
		},
		{original: Star{ID: 1, Name: "test/name", Description: "updated desc", URL: "test URL"},
			update: Star{ID: 1, Name: "updated name", Description: "updated desc", URL: "test URL"},
		},
	}

	for _, tt := range starTests {
```

The “original” star is what we know will be in the database when the test runs (note the second `original` is the same as the first `update`), and the “update” star is what we want to update it to.

Inside that loop, we send a PUT request to the update endpoint `/stars/{star.Name}`, with the contents of the updated fields:

```go
		// Set up a new request.
		req, err := http.NewRequest("PUT", fmt.Sprintf("/stars/%s", tt.original.Name), StarFormValues(tt.update))
		if err != nil {
			t.Fatal(err)
		}
		// Our API expects a form body, so set the content-type header appropriately.
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		// We need a mux router in order to pass in the `name` variable.
		r := mux.NewRouter()

		r.HandleFunc("/stars/{name:.*}", app.UpdateHandler).Methods("PUT")
		r.ServeHTTP(rr, req)
```

One difference here from the Create test: we need a custom router, since the Update handler expects a `name` variable with the name of the star we want to update. We use the same `{name:.*}` pattern here as we do in the routes at the bottom of `main.go`.

Be sure to add mux (`"github.com/gorilla/mux"`) to the import list at the top of `main_test.go`, since we’re calling `mux.NewRouter()`.

The rest of the test function is about the same as it was for create; we check the return code (`204 No Content`) and make sure the database was updated successfully:

```go
		// Test that the status code is correct.
		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusNoContent, status)
		}

		// Test that the updated star is correct.
		// Note: There is only one star in the database.
		updatedStar := Star{}
		app.DB.First(&updatedStar)
		if updatedStar != tt.update {
			t.Errorf("Updated star is invalid. Expected %+v. Got %+v instead", tt.update, updatedStar)
		}
	}

	teardown(app)
}
```

## View Handler

In the View handler test, we’ll need a couple new techniques: reading the HTTP response body, and [unmarshalling](https://en.wikipedia.org/wiki/Unmarshalling) JSON. Let’s go ahead and add the imports we’ll use to the list at the top of `main_test.go`:

```go
	"encoding/json"
	"io/ioutil"
```

For the test function, we’ll use a loop  that’s similar to our table-driven tests. We don’t really need the  anonymous struct though, so we can simplify it a bit to just a slice of  Star records:

```go
func TestViewHandler(t *testing.T) {
	app := setup()

	// Set up a test table.
	starTests := []Star{
		Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"},
		Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"},
	}

	for _, star := range starTests {
		// Create a star for us to view.
		app.DB.Create(star)

		// Set up a new request.
		req, err := http.NewRequest("GET", fmt.Sprintf("/stars/%s", star.Name), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		// We need a mux router in order to pass in the `name` variable.
		r := mux.NewRouter()

		r.HandleFunc("/stars/{name:.*}", app.ViewHandler).Methods("GET")
		r.ServeHTTP(rr, req)

		// Test that the status code is correct.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
		}
```

We haven’t needed it before now, but `rr` does hold the complete response body returned from our API. We can access it through [rr.Result().Body](https://golang.org/pkg/net/http/httptest/#ResponseRecorder.Result), by [reading](https://stackoverflow.com/questions/39945968/most-efficient-way-to-convert-io-readcloser-to-byte-array) with [ioutil.ReadAll()](https://golang.org/pkg/io/ioutil/#ReadAll):

```go
		// Read the response body.
		data, err := ioutil.ReadAll(rr.Result().Body)
		if err != nil {
			t.Fatal(err)
		}
```

Now that we have the response content in `data`, we can [unmarshal it](https://golang.org/pkg/encoding/json/#Unmarshal) into a Star struct and compare with the star we created in the database:

```go
		// Test that the updated star is correct.
		returnedStar := Star{}
		if err := json.Unmarshal(data, &returnedStar); err != nil {
			t.Errorf("Returned star is invalid JSON. Got: %s", data)
		}
		if returnedStar != star {
			t.Errorf("Returned star is invalid. Expected %+v. Got %+v instead", star, returnedStar)
		}
	}

	teardown(app)
}
```

## List Handler

The List handler test is pretty similar to the View test, in that we make a GET request and check the JSON from the response body. Here, we don’t  need a custom `mux` router, since there aren’t any variables to pass in to the List handler.

```go
func TestListHandler(t *testing.T) {
	app := setup()

	// Create a couple stars to list.
	stars := []Star{
		Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"},
		Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"},
	}

	for _, star := range stars {
		app.DB.Create(star)
	}

	// Set up a new request.
	req, err := http.NewRequest("GET", "/stars", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(app.ListHandler).ServeHTTP(rr, req)

	// Test that the status code is correct.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	data, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	// Test that our stars list is the same as what was returned.
	returnedStars := []Star{}
	if err := json.Unmarshal(data, &returnedStars); err != nil {
		t.Errorf("Returned star list is invalid JSON. Got: %s", data)
	}
	if len(returnedStars) != len(stars) {
		t.Errorf("Returned star list is an invalid length. Expected %d. Got %d instead", len(stars), len(returnedStars))
	}
	for index, returnedStar := range returnedStars {
		if returnedStar != stars[index] {
			t.Errorf("Returned star is invalid. Expected %+v. Got %+v instead", stars[index], returnedStar)
		}
	}

	teardown(app)
}
```

Note the loop at the bottom of this  test, where we make sure each item in the JSON matches each of our test  stars, in order. If the order returned from the List handler ever  changes, we’ll have to revisit this test.

## Delete Handler

The test for the Delete handler doesn’t really have anything new either,  just recycling the same concepts used above in a slightly different way:

```go
func TestDeleteHandler(t *testing.T) {
	app := setup()

	// Set up a test table.
	starTests := []struct {
		star Star
	}{
		{star: Star{ID: 1, Name: "test/name", Description: "test desc", URL: "test URL"}},
		{star: Star{ID: 2, Name: "test/another_name", Description: "test desc 2", URL: "http://example.com/"}},
	}

	for _, tt := range starTests {
		// Create a star for us to delete.
		app.DB.Create(tt.star)

		// Set up a new request.
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/stars/%s", tt.star.Name), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		// We need a mux router in order to pass in the `name` variable.
		r := mux.NewRouter()

		r.HandleFunc("/stars/{name:.*}", app.DeleteHandler).Methods("DELETE")
		r.ServeHTTP(rr, req)

		// Test that the status code is correct.
		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusNoContent, status)
		}

		// Test that the star is no longer in the db.
		deletedStar := Star{}
		app.DB.Where("name = ?", tt.star.Name).First(&deletedStar)
		if deletedStar != (Star{}) {
			t.Errorf("Star still exists in db: %+v", tt.star)
		}
	}

	teardown(app)
}
```

Run the tests with `go test`, and watch them all pass!

```go
PASS
ok      _/<...>/StarManager     0.008s
```

If you want to see the results of each individual test, use the “verbose” flag - `go test -v`:

```go
=== RUN   TestCreateHandler
--- PASS: TestCreateHandler (0.00s)
=== RUN   TestUpdateHandler
--- PASS: TestUpdateHandler (0.00s)
=== RUN   TestViewHandler
--- PASS: TestViewHandler (0.00s)
=== RUN   TestListHandler
--- PASS: TestListHandler (0.00s)
=== RUN   TestDeleteHandler
--- PASS: TestDeleteHandler (0.00s)
PASS
ok      _/home/ryan/dev/StarManager     0.008s
```

You can check out the complete code for this post [on GitHub](https://github.com/rshipp/StarManager/tree/38bb56732e89dc5c70d486cf71ce0dfa9ee88d2c).

## Conclusion

A quick recap:

1. We wrote setup/teardown functions for our integration tests.
2. Tested HTTP response codes for all five HTTP request handlers.
3. Tested HTTP response headers, and database contents for the Create handler.
4. Tested database contents for the Update and Delete handlers.
5. Tested HTTP response body JSON for the View and List handlers.

In the process, we covered several Go features and concepts:

- Writing Go test files and functions.
- Integration tests vs unit tests.
- Basic use of the `testing` and `httptest` packages.
- Table-driven tests.
- Anonymous structs.
- JSON unmarshalling.
- Reading a byte stream with `ioutil.ReadAll()`.
- Using custom routes in HTTP handler tests to pass in variables.
- And more!

In future posts, I’ll revisit this API and walk through adding some new functionality and creating a frontend for the star app.

## Additional References

I found these three resources especially helpful! If you’re new to  testing in Go and want to learn more, I highly recommend them as a  starting place.

- [Integration Tests in Go - Philosophical Hacker](https://www.philosophicalhacker.com/post/integration-tests-in-go/)
- [A Quick Guide to Testing in Golang - CaitieM](https://caitiem.com/2016/08/18/a-quick-guide-to-testing-in-golang/)
- [Testing HTTP handlers in Go - Lanre Adelowo](https://lanre.wtf/blog/2017/04/08/testing-http-handlers-go/)

There’s also the free ebook “Learn Go With Tests” on GitHub that looks really  nice, though I only used it a little for this post:

- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

Other references I used while researching, but didn’t mention in the post:
