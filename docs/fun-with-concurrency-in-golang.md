# Fun with Concurrency in Golang

October 19, 2019

> NOTE: The source for this post [is on Gitlab](https://gitlab.com/searsaw/funwithconcurrency).

Golang has strong support for concurrency. From the language itself  (goroutines, channels) to constructs in the standard library (WaitGroup, Mutex), the language tries to make it easy on the developer to write  concurrent programs. Let’s play with some of these by creating a program that spins up three different HTTP servers and allows for graceful  shutdowns when the program gets a `SIGTERM` signal.

Let’s start by creating a `main.go` file with the following contents.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	helloWorldSvr := getHelloWorldServer()
	helloNameSvr := getHelloNameServer()
	echoSvr := getEchoServer()

	helloWorldSvr.ListenAndServe()
	helloNameSvr.ListenAndServe()
	echoSvr.ListenAndServe()
	fmt.Println("all servers are started")
}

func getHelloWorldServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello, world!`))
	})

	return &http.Server{Addr: ":7000", Handler: mux}
}

func getHelloNameServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		name := params.Get("name")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	})

	return &http.Server{Addr: ":8000", Handler: mux}
}

func getEchoServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, r.Body)
	})

	return &http.Server{Addr: ":9000", Handler: mux}
}
```

Here we are creating three random HTTP servers: one who simply returns “Hello, world!”, one that takes a `name` query param and says “Hello” to the name, and one that sends back whatever is sent to it in the body of the request.

If we run this program, we can start making requests to our servers. I’ll be using [`HTTPie`](https://httpie.org/) to send requests in this post.

```bash
$ http :7000
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Oct 2019 21:53:32 GMT

Hello, world!

$ http :8000 name==Alex

http: error: ConnectionError: HTTPConnectionPool(host='localhost', port=8000): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x107e3fa10>: Failed to establish a new connection: [Errno 61] Connection refused')) while doing GET request to URL: http://localhost:8000/

$ http POST :9000 name=Alex job="Software Engineer" coffee=please

http: error: ConnectionError: HTTPConnectionPool(host='localhost', port=9000): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x104a3bb90>: Failed to establish a new connection: [Errno 61] Connection refused')) while doing GET request to URL: http://localhost:9000/
```

Well, that’s strange. The first one is running, but the other two don’t appear to be running. Well, this is because the `ListenAndServe` method on an `http.Server` struct is a blocking call. This means our program stops there until  that method returns. To get it to return, the server needs to be  shutdown. Then it would go to the next one, which will also block, and  so on and so forth.

We need a way to run all of these at the same time. Luckily, Go has  us covered with goroutines. Goroutines are lightweight threads managed  by the Go runtime which allows us to kick off another “process,” for  lack of a better word, and continue on with our program. In this case,  that means it will start each server in a separate goroutine and  continue on with our main one.

We can update our `ListenAndServe` calls like so:

```go
go helloWorldSvr.ListenAndServe()
go helloNameSvr.ListenAndServe()
go echoSvr.ListenAndServe()
```

If we run this, we see that our program exits  pretty quickly. Hmmm…well this is because we kicked off three  goroutines, which let our `main` function continue to the end since we didn’t have anything to stop it from completing. How can we  make the program wait to complete until we give it the signal to quit?  Well, once again, Go has us covered. We can use the `signal` package built into the standard library along with channels to accomplish this.

I like to think of channels as a tube that we can send messages down  from one part of our program to be picked up by another part. Channels  can be passed around in goroutines to help coordinate all kinds of  behavior. There is a little more to channels than my simplified  explanation. [Go by Example](https://gobyexample.com/channels) has a good set of posts about channels that I recommend checking out.

Let’s add the following to the end of our main function after starting our HTTP servers.

```go
signals := make(chan os.Signal, 1)
signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
<-signals
```

Here, we are creating a channel that will hold one signal at a time. Then we use the `Notify` function from the `signal` package to tell our program to tell us when SIGINT and SIGTERM are sent from the user. We will be notified by that signal being sent on the  channel. The last line is asking for a value off of the channel. This is a blocking operation that will only continue when it gets a value or  the channel is closed. Therefore, this line prevents our program from  continuing until it receives a SIGINT or SIGTERM, which would be sent by us when we stop the program .

Run the program and send the requests from earlier. They all succeed!

```bash
$ http :7000
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Oct 2019 22:34:11 GMT

Hello, world!

$ http :8000 name==Alex
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Oct 2019 22:34:16 GMT

Hello, Alex!

$ http POST :9000 name=Alex job="Software Engineer" coffee=please
HTTP/1.1 200 OK
Content-Length: 64
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Oct 2019 22:34:20 GMT

{
    "coffee": "please",
    "job": "Software Engineer",
    "name": "Alex"
}
```

Well this is great, but this isn’t as good as we  can make it. We need to try to shutdown our servers gracefully to ensure all requests are serviced before the server closes. Otherwise our end  users may send a request but never get a response! Luckily, the server  struct has a `Shutdown` method on it for us to use for just this purpose! Let’s refactor our code a bit to account for this.

```go
func getHelloWorldServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello, world!`))
	})
	server := &http.Server{Addr: ":7000", Handler: mux}

	go func() {
		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutCtx); err != nil {
			fmt.Printf("error shutting down the hello world server: %s\n", err)
		}
        fmt.Println("the hello world server is closed")
	}()

	fmt.Println("the hello world server is starting")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("error starting the hello world server: %s\n", err)
	}
	fmt.Println("the hello world server is closing")
}
```

For the sake of brevity, I’ve only shown how I updated the `getHelloWorldServer` function. I have moved the starting of the server into this function. I have also started *another* goroutine in this function that will shutdown the server. We are  creating a context that will automatically timeout and force the  shutdown after five seconds.

We need to update our main function like so:

```go
func main() {
	go getHelloWorldServer()
	go getHelloNameServer()
	go getEchoServer()

	fmt.Println("all servers are started")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
```

We are starting the goroutines when we “get” the  servers now. Try running this program. We get some errors! We can see  that we tried to start a server that was already closed. This happened  because the goroutine in each server function runs before the server is  even started. We would prefer that piece of code not run until we have  signalled for the program to shutdown. We need a way to close or cancel  something from our `main` function that will let our  goroutines know it’s time to shutdown the servers. This is a great  use-case for a context. Contexts are used all throughout the Golang  standard library as a way to signal when to stop things from running and also as a way to pass values down to functions. We will use one for the former reason.

First we need to create the context in our `main` function. It’s important we create one *and* get a cancel function for it so we can cancel it later in our program.

```go
ctx, cancel := context.WithCancel(context.Background())
```

Next, we should pass this down to each of our server functions.

```go
go getHelloWorldServer(ctx)
go getHelloNameServer(ctx)
go getEchoServer(ctx)
```

Update the function definitions to take this context.

```go
func getHelloWorldServer(ctx context.Context) {
```

Now, we need to wait for the context to be  cancelled before we try to shutdown the servers. This can be done by  waiting for a message on the `Done` channel of the context. This channel will be closed when we call cancel from our `main` function. This means any places we are waiting for a message will  continue moving forward. Add this waiting right before we create the `shutCtx` in each function.

```go
go func() {
    <-ctx.Done()
    shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
```

Now that our servers will wait for the context to  be cancelled, we need to cancel it at the right time. We want to cancel  it once we get a signal to stop everything. So let’s cancel the context  right after we get that signal!

```go
signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
<-signals
cancel()
```

Run the server again and make a few requests. It  works again! Now close the program. Notice when it stops, we don’t see  our logs saying each of our servers are closed. We need a way for the  goroutines to signal to the `main` function they are done and for the `main` function to wait until they are all gone. Go’s standard library has a `sync` package with a `WaitGroup` construct in it. We can create a group and tell it how many are in the  group. Then each goroutine will tell the group that it is done. We can  then make the `main` function wait for all the goroutines in the group to be done before moving forward with the program.

First let’s create the `WaitGroup` at the beginning of the `main` function and initialize the count in it to be three since we know we will have three servers.

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
```

Next we need to pass a reference to the group down to each of our server functions.

```go
go getHelloWorldServer(ctx, &wg)
go getHelloNameServer(ctx, &wg)
go getEchoServer(ctx, &wg)
```

Let’s also update our server function definitions to take this `WaitGroup` as a parameter.

```go
func getHelloWorldServer(ctx context.Context, wg *sync.WaitGroup) {
```

Inside our server functions, make sure to call `wg.Done()` after the server has completely shutdown to decrease the count in the `WaitGroup` by one.

```go
}
fmt.Println("the hello world server is closed")
wg.Done()
```

Lastly, we need to make our `main` function wait for the `WaitGroup` counter to be zero before continuing forward. We will put this at the bottom of the `main` function.

```go
<-signals
cancel()
wg.Wait()
```

Run the program again and then shut it down.  Notice, though they may be out of order since it all happens very  quickly, we get all our logs. Woohoo!

So we have a solid base here, but there’s one case that isn’t  handled. What if one of the servers has an error when starting up? We  should probably have the whole program fail. Right now, it will just not start up that one server. We want the program to fail fast so we can  have a quick feedback loop. The package [golang.org/x/sync/errgroup](https://godoc.org/golang.org/x/sync/errgroup) provides the `errgroup.Group` construct we can use to handle the errors returned by starting or stopping our servers.

Let’s update our code to use an `errgroup.Group` to handle the error handling for all our goroutines. First, let’s see our `main` function.

```go
import (
	// ...other stuff
	"golang.org/x/sync/errgroup"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ctx, cancel := context.WithCancel(context.Background())

	eg, egCtx := errgroup.WithContext(context.Background())
	eg.Go(getHelloWorldServer(ctx, &wg))
	eg.Go(getHelloNameServer(ctx, &wg))
	eg.Go(getEchoServer(ctx, &wg))

	go func() {
		<-egCtx.Done()
		cancel()
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	if err := eg.Wait(); err != nil {
		fmt.Printf("error in the server goroutines: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("everything closed successfully")
}
```

We now are creating an `errgroup.Group` using the `WithContext` function and are kicking off our server goroutines in the group using its `Go` method. We are using the context from the errgroup as a way to know when to cancel the rest of the goroutines. This `Done` channel is closed when either the `eg.Wait` function returns, which means all the goroutines have completed, or an  error was returned from one of the goroutines. This is the important  piece. We want to be notified when one of the goroutines errors out so  we can stop immediately. Also, we have put our signal handling logic  into a goroutine since we are not using it to halt our `main` function. We are simply using it to signal to our goroutines when to  stop when things are running correctly. Lastly, we use the `Wait` method on the errgroup to halt the `main` function until all the servers have been shutdown. It then returns the  first error that was returned from any of the goroutines or `nil`.

Now let’s see how our server functions have changed.

```go
func getHelloWorldServer(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Hello, world!`))
		})
		server := &http.Server{Addr: ":7000", Handler: mux}
		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(shutCtx); err != nil {
				errChan <- fmt.Errorf("error shutting down the hello world server: %w", err)
			}
			fmt.Println("the hello world server is closed")
			close(errChan)
			wg.Done()
		}()

		fmt.Println("the hello world server is starting")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			return fmt.Errorf("error starting the hello world server: %w", err)
		}
		fmt.Println("the hello world server is closing")
		err := <-errChan
		wg.Wait()
		return err
	}
}
```

First, our definition has been updated to return a function that itself returns an error. This is because the `Go` function of the errgroup takes a function that takes no parameters and  returns an error. On the next line, we are returning this function as an anonymous function. Inside this anonymous function is all our logic for starting and stopping our server. Below our server, we are creating  another channel that accepts errors. This channel is used to prevent the function that starts the server from returning before the goroutine  that stops the server finishes. It also has the added benefit of giving  us the error from that server if there was one.

Next, inside the goroutine for stopping the server, we are putting an error on the `errChan` if there was an error while shutting the server down. We then close the error outside of the conditional to ensure we always get something off  the channel when we wait on the channel later in the function. Lastly,  we call `Done` on the `WaitGroup` to make sure other goroutines waiting on this server to shutdown are aware it’s done.

Back outside the goroutine, we are still starting the server and  handling the error. However, handling the error, in this case, means  wrapping the error and returning it from the function. Remember, any  error returned from one of these functions will cause the errgroup  context to cancel in our `main` function, which will then cause all the other goroutines to cancel as well. If `ListenAndServe` exits without any errors, we then wait on a message on the `errChan`. Getting a message here means either there was an error while the server was shutting down or the channel was closed without any errors put on  it. Whatever is returned from the error channel we assign to a variable. We then wait for the `WaitGroup` counter to reach zero. When it does, this means all the other servers have completed and it’s time  to exit from them all. So, at the end, we simply return what came off  the error channel. If any of them are actual errors, the first one will  be returned from the `Wait` of the errgroup in our `main` function, and the program will terminate.

Run the program, send some requests, shut it down. Change the address for the hello world server to be `:8000` and see it all crash!

#   Going forward  

So we built a program that creates three HTTP servers and coordinates gracefully shutting them down while handling errors at all the  necessary points to ensure it fails fast. What’s next? Well, one easy  thing to do would be to refactor this code. The three HTTP server  functions are all the same outside of a few small things such as name  and what their handlers actually do. This could easily be abstracted out into a single function five parameters. There may be other ways to  clean up this code. I leave that as a challenge for the reader. If you  want to see how I cleaned up the code, check out [the source on Gitlab](https://gitlab.com/searsaw/funwithconcurrency).
