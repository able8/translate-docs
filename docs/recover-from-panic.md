# Panics, stack traces and how to recover [best practice]

yourbasic.org/golang

- [A panic is an exception in Go](http://yourbasic.org#a-panic-is-an-exception-in-go)
- [Stack traces](http://yourbasic.org#stack-traces)
  - [Interpret a stack trace](http://yourbasic.org#interpret-a-stack-trace)
  - [Print and log a stack trace](http://yourbasic.org#print-and-log-a-stack-trace)
  - [Level of detail](http://yourbasic.org#level-of-detail)
- [Recover and catch a panic](http://yourbasic.org#recover-and-catch-a-panic)
  - [Panic handler example](http://yourbasic.org#panic-handler-example)
  - [Return a value](http://yourbasic.org#return-a-value)
- [Test a panic (utility function)](http://yourbasic.org#test-a-panic-utility-function)

## A panic is an exception in Go

Panics are similar to C++ and Java exceptions,
but are only intended for run-time errors,
such as following a nil pointer or attempting to index an array out of bounds.
To signify events such as end-of-file, Go programs use the built-in `error` type.
See [Error handling best practice](http://yourbasic.org/golang/errors-explained/) and [3 simple ways to create an error](http://yourbasic.org/golang/create-error/) for more on errors.

A panic stops the normal execution of a goroutine:

- When a program panics, it immediately starts to unwind the call stack.
- This continues until the program crashes and prints a stack trace,
- or until the built-in`recover` function is called.

A panic is caused either by a runtime error,
or an explicit call to the built-in `panic` function.

## Stack traces

A **stack trace** – a report of all active stack frames –
is typically printed to the console when a panic occurs.
Stack traces can be very useful for debugging:

- not only do you see**where** the error happened,
- but also**how** the program arrived in this place.

### Interpret a stack trace

Here’s an example of a stack trace:

```
goroutine 11 [running]:
testing.tRunner.func1(0xc420092690)
    /usr/local/go/src/testing/testing.go:711 +0x2d2
panic(0x53f820, 0x594da0)
    /usr/local/go/src/runtime/panic.go:491 +0x283
github.com/yourbasic/bit.(*Set).Max(0xc42000a940, 0x0)
    ../src/github.com/bit/set_math_bits.go:137 +0x89
github.com/yourbasic/bit.TestMax(0xc420092690)
    ../src/github.com/bit/set_test.go:165 +0x337
testing.tRunner(0xc420092690, 0x57f5e8)
    /usr/local/go/src/testing/testing.go:746 +0xd0
created by testing.(*T).Run
    /usr/local/go/src/testing/testing.go:789 +0x2de

```

It can be read from the bottom up:

- `testing.(*T).Run` has called `testing.tRunner`,
- which has called`bit.TestMax`,
- which has called`bit.(*Set).Max`,
- which has called`panic`,
- which has called`testing.tRunner.func1`.

The indented lines show the source file and line number at which the function was called.
The hexadecimal numbers refer to parameter values, including values of pointers and internal data structures.
[Stack Traces in Go](https://www.goinggo.net/2015/01/stack-traces-in-go.html) has more details.

### Print and log a stack trace

To print the stack trace for the current goroutine, use [`debug.PrintStack`](https://golang.org/pkg/runtime/debug/#PrintStack) from package [`runtime/debug`](https://golang.org/pkg/runtime/debug/).

You can also examine the current stack trace programmatically by calling [`runtime.Stack`](https://golang.org/pkg/runtime/#Stack).

### Level of detail

The [`GOTRACEBACK`](https://golang.org/pkg/runtime/#hdr-Environment_Variables) variable
controls the amount of output generated when a Go program fails.

- `GOTRACEBACK=none` omits the goroutine stack traces entirely.
- `GOTRACEBACK=single` (the default) prints a stack trace for the current goroutine,
  eliding functions internal to the run-time system.
  The failure prints stack traces for all goroutines if there is no current goroutine
  or the failure is internal to the run-time.
- `GOTRACEBACK=all` adds stack traces for all user-created goroutines.
- `GOTRACEBACK=system` is like `all` but adds stack frames for run-time functions
  and shows goroutines created internally by the run-time.

## Recover and catch a panic

The built-in `recover` function can be used to regain
control of a panicking goroutine and resume normal execution.

- A call to`recover` stops the unwinding and returns
  the argument passed to `panic`.
- If the goroutine is not panicking,
  `recover` returns `nil`.

Because the only code that runs while unwinding is inside [deferred functions](http://yourbasic.org/golang/defer/),
`recover` is only useful inside such functions.

### Panic handler example

```
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	m := 1
	panic("foo: fail")
	m = 2
	return m
}
```

```
foo: fail
main received 0
```

Since the panic occurred before `foo` returned a value,
`n` still has its initial zero value.

### Return a value

To return a value during a panic, you must use a [named return value](http://yourbasic.org/golang/named-return-values-parameters/).

```
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() (m int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			m = 2
		}
	}()
	m = 1
	panic("foo: fail")
	m = 3
	return m
}
```

```
foo: fail
main received 2
```

## Test a panic (utility function)

In this example, we use reflection to check if a list of interface variables
have types corre­sponding to the para­meters of a given function.
If so, we call the function with those para­meters to check if there is a panic.

```
// Panics tells if function f panics with parameters p.
func Panics(f interface{}, p ...interface{}) bool {
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)
	if ft.NumIn() != len(p) {
		panic("wrong argument count")
	}
	pv := make([]reflect.Value, len(p))
	for i, v := range p {
		if reflect.TypeOf(v) != ft.In(i) {
			panic("wrong argument type")
		}
		pv[i] = reflect.ValueOf(v)
	}
	return call(fv, pv)
}

func call(fv reflect.Value, pv []reflect.Value) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = true
		}
	}()
	fv.Call(pv)
	return
}
```
