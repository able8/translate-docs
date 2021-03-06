# Handling Panics in Go

October  3, 2019 40.2k views

### Introduction

Errors that a program encounters fall into two broad categories:  those the programmer has anticipated and those the programmer has not.  The `error` interface that we have covered in our previous two articles on [error handling](https://www.digitalocean.com/community/tutorials/handling-errors-in-go) largely deal with errors that we expect as we are writing Go programs. The `error` interface even allows us to acknowledge the rare possibility of an  error occurring from function calls, so we can respond appropriately in  those situations.

Panics fall into the second category of errors, which are  unanticipated by the programmer. These unforeseen errors lead a program  to spontaneously terminate and exit the running Go program. Common  mistakes are often responsible for creating panics. In this tutorial,  we’ll examine a few ways that common operations can produce panics in  Go, and we’ll also see ways to avoid those panics. We’ll also use [`defer`](https://www.digitalocean.com/community/tutorials/understanding-defer-in-go) statements along with the `recover` function to capture panics before they have a chance to unexpectedly terminate our running Go programs.

## Understanding Panics

There are certain operations in Go that automatically return panics and stop the program. Common operations include indexing an [array](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go#arrays) beyond its capacity, performing type assertions, calling methods on nil pointers, incorrectly using mutexes, and attempting to work with closed channels. Most of these situations result from mistakes made while  programming that the compiler has no ability to detect while compiling  your program.

Since panics include detail that is useful for resolving an issue,  developers commonly use panics as an indication that they have made a  mistake during a program’s development.

### Out of Bounds Panics

When you attempt to access an index beyond the length of a slice or  the capacity of an array, the Go runtime will generate a panic.

The following example makes the common mistake of attempting to  access the last element of a slice using the length of the slice  returned by the `len` builtin. Try running this code to see why this might produce a panic:

```go
package main

import (
    "fmt"
)

func main() {
    names := []string{
        "lobster",
        "sea urchin",
        "sea cucumber",
    }
    fmt.Println("My favorite sea creature is:", names[len(names)])
}
```

 

This will have the following output:

```
Outputpanic: runtime error: index out of range [3] with length 3

goroutine 1 [running]:
main.main()
    /tmp/sandbox879828148/prog.go:13 +0x20
```

The name of the panic’s output provides a hint: `panic: runtime error: index out of range`. We created a slice with three sea creatures. We then tried to get the  last element of the slice by indexing that slice with the length of the  slice using the `len` builtin function. Remember that slices  and arrays are zero-based; so the first element is zero and the last  element in this slice is at index `2`. Since we try to access the slice at the third index, `3`, there is no element in the slice to return because it is beyond the  bounds of the slice. The runtime has no option but to terminate and exit since we have asked it to do something impossible. Go also can’t prove  during compilation that this code will try to do this, so the compiler  cannot catch this.

Notice also that the subsequent code did not run. This is because a  panic is an event that completely halts the execution of your Go program. The message produced contains multiple pieces of information  helpful for diagnosing the cause of the panic.

## Anatomy of a Panic

Panics are composed of a message indicating the cause of the panic and a [stack trace](https://en.wikipedia.org/wiki/Stack_trace) that helps you locate where in your code the panic was produced.

The first part of any panic is the message. It will always begin with the string `panic:`, which will be followed with a string that varies depending on the cause of the panic. The panic from the previous exercise has the message:

```
panic: runtime error: index out of range [3] with length 3
```

The string `runtime error:` following the `panic:` prefix tells us that the panic was generated by the language runtime.  This panic is telling us that we attempted to use an index `[3]` that was out of range of the slice’s length `3`.

Following this message is the stack trace. Stack traces form a map  that we can follow to locate exactly what line of code was executing  when the panic was generated, and how that code was invoked by earlier  code.

```
goroutine 1 [running]:
main.main()
    /tmp/sandbox879828148/prog.go:13 +0x20
```

This stack trace, from the previous example, shows that our program generated the panic from the file `/tmp/sandbox879828148/prog.go` at line number 13. It also tells us that this panic was generated in the `main()` function from the `main` package.

The stack trace is broken into separate blocks—one for each [goroutine](https://tour.golang.org/concurrency/1) in your program. Every Go program’s execution is accomplished by one or more goroutines that can each independently and simultaneously execute  parts of your Go code. Each block begins with the header `goroutine X [state]:`. The header gives the ID number of the goroutine along with the state  that it was in when the panic occurred. After the header, the stack  trace shows the function that the program was executing when the panic  happened, along with the filename and line number where the function  executed.

The panic in the previous example was generated by an out-of-bounds  access to a slice. Panics can also be generated when methods are called  on pointers that are unset.

## Nil Receivers

The Go programming language has pointers to refer to a specific  instance of some type existing in the computer’s memory at runtime.  Pointers can assume the value `nil` indicating that they are not pointing at anything. When we attempt to call methods on a pointer that is `nil`, the Go runtime will generate a panic. Similarly, variables that are  interface types will also produce panics when methods are called on  them. To see the panics generated in these cases, try the following  example:

```go
package main

import (
    "fmt"
)

type Shark struct {
    Name string
}

func (s *Shark) SayHello() {
    fmt.Println("Hi! My name is", s.Name)
}

func main() {
    s := &Shark{"Sammy"}
    s = nil
    s.SayHello()
}
```

 

The panics produced will look like this:

```
Outputpanic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xdfeba]

goroutine 1 [running]:
main.(*Shark).SayHello(...)
    /tmp/sandbox160713813/prog.go:12
main.main()
    /tmp/sandbox160713813/prog.go:18 +0x1a
```

In this example, we defined a struct called `Shark`. `Shark` has one method defined on its pointer receiver called `SayHello` that will print a greeting to standard out when called. Within the body of our `main` function, we create a new instance of this `Shark` struct and request a pointer to it using the `&` operator. This pointer is assigned to the `s` variable. We then reassign the `s` variable to the value `nil` with the statement `s = nil`. Finally we attempt to call the `SayHello` method on the variable `s`. Instead of receiving a friendly message from Sammy, we receive a panic  that we have attempted to access an invalid memory address. Because the `s` variable is `nil`, when the `SayHello` function is called, it tries to access the field `Name` on the `*Shark` type.  Because this is a pointer receiver, and the receiver in this case is `nil`, it panics because it can’t dereference a `nil` pointer.

While we have set `s` to `nil` explicitly in this example, in practice this happens less obviously. When you see panics involving `nil pointer dereference`, be sure that you have properly assigned any pointer variables that you may have created.

Panics generated from nil pointers and out-of-bounds accesses are two commonly occurring panics generated by the runtime. It is also possible to manually generate a panic using a builtin function.

## Using the `panic` Builtin Function

We can also generate panics of our own using the `panic`  built-in function. It takes a single string as an argument, which is the message the panic will produce. Typically this message is less verbose  than rewriting our code to return an error. Furthermore, we can use this within our own packages to indicate to developers that they may have  made a mistake when using our package’s code. Whenever possible, best  practice is to try to return `error` values to consumers of our package.

Run this code to see a panic generated from a function called from another function:

```go
package main

func main() {
    foo()
}

func foo() {
    panic("oh no!")
}
```

 

The panic output produced looks like:

```
Outputpanic: oh no!

goroutine 1 [running]:
main.foo(...)
    /tmp/sandbox494710869/prog.go:8
main.main()
    /tmp/sandbox494710869/prog.go:4 +0x40
```

Here we define a function `foo` that calls the `panic` builtin with the string `"oh no!"`. This function is called by our `main` function. Notice how the output has the message `panic: oh no!` and the stack trace shows a single goroutine with two lines in the stack trace: one for the `main()` function and one for our `foo()` function.

We’ve seen that panics appear to terminate our program where they are generated. This can create problems when there are open resources that  need to be properly closed. Go provides a mechanism to execute some code always, even in the presence of a panic.

## Deferred Functions

Your program may have resources that it must clean up properly, even  while a panic is being processed by the runtime. Go allows you to defer  the execution of a function call until its calling function has  completed execution. Deferred functions run even in the presence of a  panic, and are used as a safety mechanism to guard against the chaotic  nature of panics. Functions are deferred by calling them as usual, then  prefixing the entire statement with the `defer` keyword, as in `defer sayHello()`. Run this example to see how a message can be printed even though a panic was produced:

```go
package main

import "fmt"

func main() {
    defer func() {
        fmt.Println("hello from the deferred function!")
    }()

    panic("oh no!")
}
```

 

The output produced from this example will look like:

```
Outputhello from the deferred function!
panic: oh no!

goroutine 1 [running]:
main.main()
    /Users/gopherguides/learn/src/github.com/gopherguides/learn//handle-panics/src/main.go:10 +0x55
```

Within the `main` function of this example, we first `defer` a call to an anonymous function that prints the message `"hello from the deferred function!"`. The `main` function then immediately produces a panic using the `panic` function. In the output from this program, we first see that the  deferred function is executed and prints its message. Following this is  the panic we generated in `main`.

Deferred functions provide protection against the surprising nature  of panics. Within deferred functions, Go also provides us the  opportunity to stop a panic from terminating our Go program using  another built-in function.

## Handling Panics

Panics have a single recovery mechanism—the `recover`  builtin function. This function allows you to intercept a panic on its  way up through the call stack and prevent it from unexpectedly  terminating your program. It has strict rules for its use, but can be  invaluable in a production application.

Since it is part of the `builtin` package, `recover` can be called without importing any additional packages:

```
package main

import (
    "fmt"
    "log"
)

func main() {
    divideByZero()
    fmt.Println("we survived dividing by zero!")

}

func divideByZero() {
    defer func() {
        if err := recover(); err != nil {
            log.Println("panic occurred:", err)
        }
    }()
    fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
    return a / b
}
```

This example will output:

```
Output2009/11/10 23:00:00 panic occurred: runtime error: integer divide by zero
we survived dividing by zero!
```

Our `main` function in this example calls a function we define, `divideByZero`. Within this function, we `defer` a call to an anonymous function responsible for dealing with any panics that may arise while executing `divideByZero`. Within this deferred anonymous function, we call the `recover` builtin function and assign the error it returns to a variable. If `divideByZero` is panicking, this `error` value will be set, otherwise it will be `nil`. By comparing the `err` variable against `nil`, we can detect if a panic occurred, and in this case we log the panic using the `log.Println` function, as though it were any other `error`.

Following this deferred anonymous function, we call another function that we defined, `divide`, and attempt to print its results using `fmt.Println`. The arguments provided will cause `divide` to perform a division by zero, which will produce a panic.

In the output to this example, we first see the log message from the  anonymous function that recovers the panic, followed by the message `we survived dividing by zero!`. We have indeed done this, thanks to the `recover` builtin function stopping an otherwise catastrophic panic that would terminate our Go program.

The `err` value returned from `recover()` is exactly the value that was provided to the call to `panic()`. It’s therefore critical to ensure that the `err` value is only nil when a panic has not occurred.

## Detecting Panics with `recover`

The `recover` function relies on the value of the error to make determinations as to whether a panic occurred or not. Since the  argument to the `panic` function is an empty interface, it can be any type. The zero value for any interface type, including the empty interface, is `nil`. Care must be taken to avoid `nil` as an argument to `panic` as demonstrated by this example:

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    divideByZero()
    fmt.Println("we survived dividing by zero!")

}

func divideByZero() {
    defer func() {
        if err := recover(); err != nil {
            log.Println("panic occurred:", err)
        }
    }()
    fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
    if b == 0 {
        panic(nil)
    }
    return a / b
}
```

 

This will output:

```
Outputwe survived dividing by zero!
```

This example is the same as the previous example involving `recover` with some slight modifications. The `divide` function has been altered to check if its divisor, `b`, is equal to `0`. If it is, it will generate a panic using the `panic` builtin with an argument of `nil`. The output, this time, does not include the log message showing that a panic occurred even though one was created by `divide`. This silent behavior is why it is very important to ensure that the argument to the `panic` builtin function is not `nil`.

## Conclusion

We have seen a number of ways that `panic`s can be created in Go and how they can be recovered from using the `recover` builtin. While you may not necessarily use `panic` yourself, proper recovery from panics is an important step of making Go applications production-ready.

You can also explore [our entire How To Code in Go series](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go).
