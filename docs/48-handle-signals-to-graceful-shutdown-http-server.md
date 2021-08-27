# How to handle signals with Go to graceful shutdown HTTP server

Written on June 18, 2020  ​From: https://rafallorenz.com/go/handle-signals-to-graceful-shutdown-http-server/

In this article we are going to learn how to handle os incoming [signals](https://golang.org/pkg/os/signal) for performing graceful shutdown of http server. To do so we are going to take advantage of [os/signal](https://golang.org/pkg/os/signal) package.

> Signals are primarily used on Unix-like systems.

# Types of signals

We are going to focus on asynchronous signals. They are not triggered by program errors, but are instead sent from the kernel or from some  other program.

Of the asynchronous signals:
- the `SIGHUP` signal is sent when a program loses its controlling terminal
- the `SIGINT` signal is sent when the user at the controlling terminal presses the interrupt character, which by default is **^C (Control-C)**
- The `SIGQUIT` signal is sent when the user at the controlling terminal presses the quit character, which by default is **^\ (Control-Backslash)**

In general you can cause a program to simply exit by pressing ^C, and you can cause it to exit with a stack dump by pressing ^.

# Default behavior of signals in Go programs

> By default, a synchronous signal is converted into a run-time panic. A `SIGHUP`, `SIGINT`, or `SIGTERM` signal causes the program to exit. If the Go program is started with either SIGHUP or SIGINT ignored  (signal handler set to SIG_IGN), they will remain ignored. If the Go program is started with a non-empty signal mask, that will  generally be honored. However, some signals are explicitly unblocked:  the synchronous signals.

You can read more about it on the [package documentation](https://golang.org/pkg/os/signal/#hdr-Default_behavior_of_signals_in_Go_programs)

# Handling signals

The idea is to catch incoming signal and perform graceful stop of our http application. We can create signal channel using, and use it to  notify on incoming signal. `signal.Notify` disables the default behavior for a given set of asynchronous signals  and instead delivers them over one or more registered channels.

```
  signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	// terminate after second signal before callback is done
	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	// PERFORM GRACEFUL SHUTDOWN HERE

	os.Exit(0)
```

When signal is received we will call a callback followed after by `os.Exit(0)`, if second signal is received will terminate process by a call to `os.Exit(1)`.

# Server

## Graceful shutdown

To gracefully shutdown [http.Server](https://golang.org/pkg/net/http/#Server) we can use [Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown) method.

> Shutdown gracefully shuts down the server without interrupting any  active connections. Shutdown works by first closing all open listeners,  then closing all idle connections, and then waiting indefinitely for  connections to return to idle and then shut down. If the provided  context expires before the shutdown is complete, Shutdown returns the  context’s error, otherwise it returns any error returned from closing  the Server’s underlying Listener(s).

```
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := httpServer.Shutdown(ctx); err != nil {
    log.Fatalf("shutdown error: %v\n", err)
  } else {
    log.Printf("gracefully stopped\n")
  }
```

> Shutdown does not attempt to close nor wait for hijacked  connections such as WebSockets. The caller of Shutdown should separately notify such long-lived connections of shutdown and wait for them to  close, if desired. See [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown) for a way to register shutdown notification functions.

## Handling hijacked connections

> [RegisterOnShutdown](https://golang.org/pkg/net/http/#Server.RegisterOnShutdown) registers a function to call on Shutdown. This can be used to  gracefully shutdown connections that have undergone ALPN protocol  upgrade or that have been hijacked. This function should start  protocol-specific graceful shutdown, but should not wait for shutdown to complete.

```
	ctx, cancel := context.WithCancel(context.Background())

	httpServer.RegisterOnShutdown(cancel)
```

We want to ask context passed down to the socket handlers/goroutines to stop. To do so we can set `BaseContext` property on [http.Server](https://golang.org/pkg/net/http/#Server).

> BaseContext optionally specifies a function that returns the base  context for incoming requests on this server. The provided Listener is  the specific Listener that’s about to start accepting requests. If  BaseContext is nil, the default is context.Background(). If non-nil, it  must return a non-nil context.

```
	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	httpServer.RegisterOnShutdown(cancel)
```

Keep in mind that doing so will make [Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown) cancel context via `RegisterOnShutdown` which will terminate all handlers using `BaseContext` immediately.

Full correct solution would require you to separate base context for  WeSocket connections and other HTTP handlers. Also seems like  introducing simple timeout to cancel `BaseContext` will not be enough, connections idleness has to be checked as well.

If you do not care about notifying long-lived connections during  shutdown and don’t want to wait for them to close gracefully. Quick  solution would be to manually cancel `BaseContext` instead of using `RegisterOnShutdown` which makes context get canceled as the first procedure during shutdown.

```
	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	
	// GRACEFULLY SHUTDOWN

	cancel()
```

# Gluing up all pieces together

```
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	// if your BaseContext is more complex you might want to use this instead of doing it manually
	// httpServer.RegisterOnShutdown(cancel)

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
```

When Shutdown is called, `Serve`, `ListenAndServe`, and `ListenAndServeTLS` immediately return `ErrServerClosed`. Make sure the program doesn’t exit and waits instead for Shutdown to return.

**Note**: that we are deferring `os.Exit()` followed by `return`. Defers don’t run on `Fatal()`.

> Calling Goexit from the main goroutine terminates that goroutine  without func main returning. Since func main has not returned, the  program continues execution of other goroutines. If all other goroutines exit, the program crashes.

You can read conversation about it [here](https://www.reddit.com/r/golang/comments/hbalgf/how_to_handle_signals_with_go_to_graceful/fv800rj?utm_source=share&utm_medium=web2x)

Run this example on [The Go Playground](https://play.golang.org/p/6w3yFxU1N24)

# Conclusion

By using tools provided by go environment, we can easily handle  graceful shutdown of our application. We could simply create a package  to handle that for us, in fact I already have created one if you are  interested. [shutdown](https://github.com/vardius/shutdown) - is a simple go signals handler for performing graceful shutdown by executing callback function.
