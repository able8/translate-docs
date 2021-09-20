# How to Parse a JSON Request Body in Go

Posted on: 21st October 2019

Let's say that you're building a JSON API with Go. And in some of the handlers — probably as part of a POST or PUT request — you want to read a JSON object from the request body and assign it to a struct in your code.

After a bit of research, there's a good chance that you'll end up with some code that looks similar to the `personCreate` handler here:

```
// File: main.go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Person struct {
    Name string
    Age  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    // Declare a new Person struct.
    var p Person

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Do something with the Person struct...
    fmt.Fprintf(w, "Person: %+v", p)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
``` 

If you're putting together a quick prototype, or building an API for personal/internal use only, then the code in the `personCreate` handler is probably OK.

But if you're building an API for public use in production then there are a few issues with this to be aware of, and things that can be improved.

1.  Not all errors returned by [`Decode()`](https://golang.org/pkg/encoding/json/#Decoder.Decode) are caused by a bad request from the client. Specifically, `Decode()` can return a [`json.InvalidUnmarshalError`](https://golang.org/pkg/encoding/json/#InvalidUnmarshalError) error — which is caused by an unmarshalable target destination being passed to `Decode()`. If that happens, then it indicates a problem with our application — not the client request — so really the error should be logged and a `500 Internal Server Error` response sent to the client instead.
    
2.  The error messages returned by `Decode()` aren't ideal for sending to a client. Some are arguably too detailed and expose information about the underlying program (like `"json: cannot unmarshal number into Go struct field Person.Name of type string"`). Others aren't descriptive enough (like `"unexpected EOF"`) and some are just plain confusing (like `"invalid character 'A' looking for beginning of object key string"`). There also isn't consistency in the formatting or language used.
    
3.  A client can include extra unexpected fields in their JSON, and these fields will be silently ignored without the client receiving any error. We can fix this by using the decoder's [`DisallowUnknownFields()`](https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields) method.
    
4.  There's no upper limit on the size of the request body that will be read by the `Decode()` method. Limiting this would help prevent our server resources being wasted if a malcious client sends a very large request body, and it's something we can easily do by using the [`http.MaxBytesReader()`](https://golang.org/pkg/net/http/#MaxBytesReader) function.
    
5.  There's no check for a `Content-Type: application/json` header in the request. Of course, this header may not always be present, and mistakes and malicious clients mean that it isn't a guarantee of the _actual_ content type. But checking for an incorrect `Content-Type` header would allow us to 'fail fast' and send a helpful error message without spending unnecessary resources on parsing the body.
    
6.  The decoder that we create with `json.NewDecoder()` is designed to decode streams of JSON objects and considers a request body like `'{"Name": "Bob"}{"Name": "Carol": "Age": 54}'` or `'{"Name": "Dave"}{}'` to be valid. But in the code above only the first JSON object in the request body will actually be parsed. So if the client sends multiple JSON objects in the request body, we want to alert them to the fact that only a single object is supported.
    
    There are two ways to achieve this. We can either call the decoder's `Decode()` method for a second time and make sure that it returns an `io.EOF` error (if it does, then we know there are not any additional JSON objects or other data in the request body). Or we could avoid using `Decode()` altogether and read the body into a byte slice and pass it to [](https://golang.org/pkg/encoding/json/#Unmarshal)`json.Unmarshal()`, which _would_ return an error if the body contains multiple JSON objects. The downside of using `json.Unmarshal()` is that there is no way to disallow extra unexpected fields in the JSON, so we can't address point 3 above.
    

An Improved Handler
-------------------

Let's implement an alternative version of the `personCreate` handler which addresses all of these issues.

You'll notice here that we're using the new [`errors.Is()`](https://golang.org/pkg/errors/#Is) and [`errors.As()`](https://golang.org/pkg/errors/#As) functions, which have been introduced in Go 1.13, to help intercept the errors from `Decode()`.

```
// File: main.go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"

    "github.com/golang/gddo/httputil/header"
)

type Person struct {
    Name string
    Age  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    // If the Content-Type header is present, check that it has the value
    // application/json. Note that we are using the gddo/httputil/header
    // package to parse and extract the value here, so the check works
    // even if the client includes additional charset or boundary
    // information in the header.
    if r.Header.Get("Content-Type") != "" {
        value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
        if value != "application/json" {
            msg := "Content-Type header is not application/json"
            http.Error(w, msg, http.StatusUnsupportedMediaType)
            return
        }
    }

    // Use http.MaxBytesReader to enforce a maximum read of 1MB from the
    // response body. A request body larger than that will now result in
    // Decode() returning a "http: request body too large" error.
    r.Body = http.MaxBytesReader(w, r.Body, 1048576)
    
    // Setup the decoder and call the DisallowUnknownFields() method on it.
    // This will cause Decode() to return a "json: unknown field ..." error
    // if it encounters any extra unexpected fields in the JSON. Strictly
    // speaking, it returns an error for "keys which do not match any
    // non-ignored, exported fields in the destination".
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    
    var p Person
    err := dec.Decode(&p)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
    
        switch {
        // Catch any syntax errors in the JSON and send an error message
        // which interpolates the location of the problem to make it
        // easier for the client to fix.
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            http.Error(w, msg, http.StatusBadRequest)
    
        // In some circumstances Decode() may also return an
        // io.ErrUnexpectedEOF error for syntax errors in the JSON. There
        // is an open issue regarding this at
        // https://github.com/golang/go/issues/25956.
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            http.Error(w, msg, http.StatusBadRequest)
    
        // Catch any type errors, like trying to assign a string in the
        // JSON request body to a int field in our Person struct. We can
        // interpolate the relevant field name and position into the error
        // message to make it easier for the client to fix.
        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            http.Error(w, msg, http.StatusBadRequest)
    
        // Catch the error caused by extra unexpected fields in the request
        // body. We extract the field name from the error message and
        // interpolate it in our custom error message. There is an open
        // issue at https://github.com/golang/go/issues/29035 regarding
        // turning this into a sentinel error.
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            http.Error(w, msg, http.StatusBadRequest)
    
        // An io.EOF error is returned by Decode() if the request body is
        // empty.
        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            http.Error(w, msg, http.StatusBadRequest)
    
        // Catch the error caused by the request body being too large. Again
        // there is an open issue regarding turning this into a sentinel
        // error at https://github.com/golang/go/issues/30715.
        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            http.Error(w, msg, http.StatusRequestEntityTooLarge)
    
        // Otherwise default to logging the error and sending a 500 Internal
        // Server Error response.
        default:
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
    }
    
    // Call decode again, using a pointer to an empty anonymous struct as 
    // the destination. If the request body only contained a single JSON 
    // object this will return an io.EOF error. So if we get anything else, 
    // we know that there is additional data in the request body.
    err = dec.Decode(&struct{}{})
    if err != io.EOF {
        msg := "Request body must only contain a single JSON object"
        http.Error(w, msg, http.StatusBadRequest)
        return
    }
    
    fmt.Fprintf(w, "Person: %+v", p)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
```


The clear downside here is that this code is a lot more verbose, and IMO, a little bit ugly. Things aren't helped by the fact that there are quite a few open issues with `json/encoding` which are on hold pending [a wider review](https://github.com/golang/go/issues/29035#issuecomment-444598621) of the package.

But from a security and client perspective it's a lot better : )

The handler is now stricter about the content it will accept; we're reducing the amount of server resources used unnecessarily; and the client gets clear and consistent error messages that provide a decent amount of information without over-sharing.

As a side note, you might have noticed that the `json/encoding` package contains some other error types (like [`json.UnmarshalFieldError`](https://golang.org/pkg/encoding/json/#UnmarshalFieldError)) which aren't checked in the code above — but these have been deprecated and not used by the current version of Go.

Making a Helper Function
------------------------

If you've got a few handlers that need to to process JSON request bodies, you probably don't want to repeat this code in all of them.

A solution which I've found works well is to create a `decodeJSONBody()` helper function, and have this return a custom `malformedRequest` error type which wraps the errors and relevant status codes.

For example:

```
// File: helpers.go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "strings"

    "github.com/golang/gddo/httputil/header"
)

type malformedRequest struct {
    status int
    msg    string
}

func (mr *malformedRequest) Error() string {
    return mr.msg
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    if r.Header.Get("Content-Type") != "" {
        value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
        if value != "application/json" {
            msg := "Content-Type header is not application/json"
            return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
        }
    }

    r.Body = http.MaxBytesReader(w, r.Body, 1048576)
    
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    
    err := dec.Decode(&dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError
    
        switch {
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    
        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    
        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    
        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}
    
        default:
            return err
        }
    }
    
    err = dec.Decode(&struct{}{})
    if err != io.EOF {
        msg := "Request body must only contain a single JSON object"
        return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    }
    
    return nil
}` 

Once that's written, the code in your handlers can be kept really nice and compact:

`// File: main.go
package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
)

type Person struct {
    Name string
    Age  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    var p Person

    err := decodeJSONBody(w, r, &p)
    if err != nil {
        var mr *malformedRequest
        if errors.As(err, &mr) {
            http.Error(w, mr.msg, mr.status)
        } else {
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
    }
    
    fmt.Fprintf(w, "Person: %+v", p)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    log.Println("Starting server on :4000...")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
``` 

If you enjoyed this blog post, don't forget to check out my new book about how to [build professional web applications with Go](https://lets-go.alexedwards.net/)!

Follow me on Twitter [@ajmedwards](https://twitter.com/ajmedwards).

All code snippets in this post are free to use under the [MIT Licence](https://opensource.org/licenses/MIT).

### Related Posts

*   [Surprises and Gotchas When Working With JSON](https://www.alexedwards.net/blog/json-surprises-and-gotchas)
*   [Using PostgreSQL JSONB with Go](https://www.alexedwards.net/blog/using-postgresql-jsonb)
