# How to Parse a JSON Request Body in Go

# 如何在 Go 中解析 JSON 请求体

Posted on: 21st October 2019

发表于：2019 年 10 月 21 日

Let's say that you're building a JSON API with Go. And in some of the handlers — probably as part of a POST or PUT request — you want to read a JSON object from the request body and assign it to a struct in your code.

假设您正在使用 Go 构建一个 JSON API。在某些处理程序中——可能作为 POST 或 PUT 请求的一部分——您希望从请求正文中读取 JSON 对象并将其分配给代码中的结构体。

After a bit of research, there's a good chance that you'll end up with some code that looks similar to the `personCreate` handler here:

经过一些研究，很有可能你最终会得到一些类似于这里的 `personCreate` 处理程序的代码：

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

    // Try to decode the request body into the struct.If there is an error,
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

如果您正在构建一个快速原型，或者构建一个仅供个人/内部使用的 API，那么 `personCreate` 处理程序中的代码可能没问题。

But if you're building an API for public use in production then there are a few issues with this to be aware of, and things that can be improved.

但是，如果您正在构建用于生产中公共使用的 API，那么需要注意一些问题，并且可以改进一些事情。

1.  Not all errors returned by [`Decode()`](https://golang.org/pkg/encoding/json/#Decoder.Decode) are caused by a bad request from the client. Specifically, `Decode()` can return a [`json.InvalidUnmarshalError`](https://golang.org/pkg/encoding/json/#InvalidUnmarshalError) error — which is caused by an unmarshalable target destination being passed to `Decode ()`. If that happens, then it indicates a problem with our application — not the client request — so really the error should be logged and a `500 Internal Server Error` response sent to the client instead.
    
2.  The error messages returned by `Decode()` aren't ideal for sending to a client. Some are arguably too detailed and expose information about the underlying program (like `"json: cannot unmarshal number into Go struct field Person.Name of type string"`). Others aren't descriptive enough (like `"unexpected EOF"`) and some are just plain confusing (like `"invalid character 'A' looking for beginning of object key string"`). There also isn't consistency in the formatting or language used.
    
3.  A client can include extra unexpected fields in their JSON, and these fields will be silently ignored without the client receiving any error. We can fix this by using the decoder's [`DisallowUnknownFields()`](https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields) method.
    
4.  There's no upper limit on the size of the request body that will be read by the `Decode()` method. Limiting this would help prevent our server resources being wasted if a malcious client sends a very large request body, and it's something we can easily do by using the [`http.MaxBytesReader()`](https://golang.org/pkg/net/http/#MaxBytesReader) function.
    
5.  There's no check for a `Content-Type: application/json` header in the request. Of course, this header may not always be present, and mistakes and malicious clients mean that it isn't a guarantee of the _actual_ content type. But checking for an incorrect `Content-Type` header would allow us to 'fail fast' and send a helpful error message without spending unnecessary resources on parsing the body. 

1. 并非所有 [`Decode()`](https://golang.org/pkg/encoding/json/#Decoder.Decode)返回的错误都是由客户端的错误请求引起的。具体来说，`Decode()` 可以返回 [`json.InvalidUnmarshalError`](https://golang.org/pkg/encoding/json/#InvalidUnmarshalError) 错误——这是由传递给 `Decode 的不可编组目标目的地引起的()`。如果发生这种情况，那么它表明我们的应用程序有问题——而不是客户端请求——所以实际上应该记录错误并将“500 Internal Server Error”响应发送到客户端。
    
2.`Decode()` 返回的错误消息不适合发送到客户端。有些可能太详细了，并且暴露了有关底层程序的信息（例如“json：无法将数字解组到 Go struct 字段 Person.Name 类型字符串”中）。其他的描述性不够（比如“意外的 EOF”），有些只是容易混淆（比如“无效字符 'A' 寻找对象键字符串的开头”`）。使用的格式或语言也不一致。
    
3. 客户端可以在其 JSON 中包含额外的意外字段，这些字段将被静默忽略，客户端不会收到任何错误。我们可以通过使用解码器的 [`DisallowUnknownFields()`](https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields) 方法来解决这个问题。
    
4. Decode() 方法读取的请求体的大小没有上限。如果恶意客户端发送非常大的请求主体，限制这将有助于防止我们的服务器资源被浪费，这是我们可以通过使用 [`http.MaxBytesReader()`](https://golang.org/pkg/net/http/#MaxBytesReader) 函数。
    
5. 请求中没有检查 `Content-Type: application/json` 标头。当然，这个标头可能并不总是存在，错误和恶意客户端意味着它不是 _actual_ 内容类型的保证。但是检查不正确的“Content-Type”标头将使我们能够“快速失败”并发送有用的错误消息，而无需花费不必要的资源来解析正文。

6.  The decoder that we create with `json.NewDecoder()` is designed to decode streams of JSON objects and considers a request body like `'{"Name": "Bob"}{"Name": "Carol": " Age": 54}'` or `'{"Name": "Dave"}{}'` to be valid. But in the code above only the first JSON object in the request body will actually be parsed. So if the client sends multiple JSON objects in the request body, we want to alert them to the fact that only a single object is supported.
    
     There are two ways to achieve this. We can either call the decoder's `Decode()` method for a second time and make sure that it returns an `io.EOF` error (if it does, then we know there are not any additional JSON objects or other data in the request body). Or we could avoid using `Decode()` altogether and read the body into a byte slice and pass it to [](https://golang.org/pkg/encoding/json/#Unmarshal)`json.Unmarshal()` , which _would_ return an error if the body contains multiple JSON objects. The downside of using `json.Unmarshal()` is that there is no way to disallow extra unexpected fields in the JSON, so we can't address point 3 above.
    

6. 我们使用 `json.NewDecoder()` 创建的解码器旨在解码 JSON 对象流，并考虑像这样的请求主体 `'{"Name": "Bob"}{"Name": "Carol": " Age": 54}'` 或 `'{"Name": "Dave"}{}'` 才有效。但是在上面的代码中，实际上只会解析请求正文中的第一个 JSON 对象。因此，如果客户端在请求正文中发送多个 JSON 对象，我们希望提醒他们仅支持单个对象这一事实。
    
    有两种方法可以实现这一点。我们可以再次调用解码器的 `Decode()` 方法并确保它返回一个 `io.EOF` 错误（如果确实如此，那么我们知道请求中没有任何额外的 JSON 对象或其他数据）身体）。或者我们可以完全避免使用 `Decode()` 并将主体读入一个字节切片并将其传递给 [](https://golang.org/pkg/encoding/json/#Unmarshal)`json.Unmarshal()` ，如果正文包含多个 JSON 对象，则_将_返回错误。使用 `json.Unmarshal()` 的缺点是无法禁止 JSON 中额外的意外字段，因此我们无法解决上面的第 3 点。
    

An Improved Handler
-------------------

Let's implement an alternative version of the `personCreate` handler which addresses all of these issues.

改进的处理程序
让我们实现一个替代版本的 `personCreate` 处理程序来解决所有这些问题。

You'll notice here that we're using the new [`errors.Is()`](https://golang.org/pkg/errors/#Is) and [`errors.As()`](https://golang.org/pkg/errors/#As) functions, which have been introduced in Go 1.13, to help intercept the errors from `Decode()`.

您会在此处注意到我们使用了新的 [`errors.Is()`](https://golang.org/pkg/errors/#Is) 和 [`errors.As()`](https://golang.org/pkg/errors/#As) 函数，已在 Go 1.13 中引入，以帮助拦截来自 `Decode()` 的错误。

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
    // application/json.Note that we are using the gddo/httputil/header
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
    // response body.A request body larger than that will now result in
    // Decode() returning a "http: request body too large" error.
    r.Body = http.MaxBytesReader(w, r.Body, 1048576)
    
    // Setup the decoder and call the DisallowUnknownFields() method on it.
    // This will cause Decode() to return a "json: unknown field ..." error
    // if it encounters any extra unexpected fields in the JSON.Strictly
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
        // io.ErrUnexpectedEOF error for syntax errors in the JSON.There
        // is an open issue regarding this at
        // https://github.com/golang/go/issues/25956.
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            http.Error(w, msg, http.StatusBadRequest)
    
        // Catch any type errors, like trying to assign a string in the
        // JSON request body to a int field in our Person struct.We can
        // interpolate the relevant field name and position into the error
        // message to make it easier for the client to fix.
        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            http.Error(w, msg, http.StatusBadRequest)
    
        // Catch the error caused by extra unexpected fields in the request
        // body.We extract the field name from the error message and
        // interpolate it in our custom error message.There is an open
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
    
        // Catch the error caused by the request body being too large.Again
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
    // the destination.If the request body only contained a single JSON
    // object this will return an io.EOF error.So if we get anything else,
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

这里明显的缺点是这段代码更加冗长，而且 IMO 有点难看。由于`json/encoding` 有很多未解决的问题，这些问题都无济于事，这些问题正在等待 [更广泛的审查](https://github.com/golang/go/issues/29035#issuecomment-444598621) 的包。

But from a security and client perspective it's a lot better : )

但从安全和客户端的角度来看，它要好得多:)

The handler is now stricter about the content it will accept; we're reducing the amount of server resources used unnecessarily; and the client gets clear and consistent error messages that provide a decent amount of information without over-sharing.

处理程序现在对其将接受的内容更加严格；我们正在减少不必要使用的服务器资源量；并且客户端会获得清晰一致的错误消息，这些消息提供了大量的信息而不会过度共享。

As a side note, you might have noticed that the `json/encoding` package contains some other error types (like [`json.UnmarshalFieldError`](https://golang.org/pkg/encoding/json/#UnmarshalFieldError)) which aren't checked in the code above — but these have been deprecated and not used by the current version of Go.

作为旁注，您可能已经注意到 `json/encoding` 包包含一些其他错误类型（例如 [`json.UnmarshalFieldError`](https://golang.org/pkg/encoding/json/#UnmarshalFieldError))上面的代码中没有检查这些 - 但是这些已经被弃用并且不被当前版本的 Go 使用。

Making a Helper Function
------------------------

If you've got a few handlers that need to to process JSON request bodies, you probably don't want to repeat this code in all of them.

制作辅助函数
如果您有几个处理程序需要处理 JSON 请求正文，您可能不想在所有这些处理程序中重复此代码。

A solution which I've found works well is to create a `decodeJSONBody()` helper function, and have this return a custom `malformedRequest` error type which wraps the errors and relevant status codes.

我发现一个很有效的解决方案是创建一个 `decodeJSONBody()` 辅助函数，并让它返回一个自定义的 `malformedRequest` 错误类型，它包装了错误和相关的状态代码。

For example:

例如：

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

如果您喜欢这篇博文，请不要忘记查看我关于如何[使用 Go 构建专业 Web 应用程序](https://lets-go.alexedwards.net/) 的新书！

Follow me on Twitter [@ajmedwards](https://twitter.com/ajmedwards).

在 Twitter 上关注我 [@ajmedwards](https://twitter.com/ajmedwards)。

All code snippets in this post are free to use under the [MIT Licence](https://opensource.org/licenses/MIT).

本文中的所有代码片段均可在 [MIT 许可证](https://opensource.org/licenses/MIT) 下免费使用。

### Related Posts

###  相关文章

*   [Surprises and Gotchas When Working With JSON](https://www.alexedwards.net/blog/json-surprises-and-gotchas)
*   [Using PostgreSQL JSONB with Go](https://www.alexedwards.net/blog/using-postgresql-jsonb) 

* [使用 JSON 时的惊喜和陷阱](https://www.alexedwards.net/blog/json-surprises-and-gotchas)
* [在 Go 中使用 PostgreSQL JSONB](https://www.alexedwards.net/blog/using-postgresql-jsonb)

