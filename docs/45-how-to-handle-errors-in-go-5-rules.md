# How to handle errors in Go?

## Practically and efficiently!

While handlings errors in Go is exceptionally annoying, I like the explicit error checks much more than throwing an exception 5 levels up the stack and hoping someone will catch it. I am looking at you Java!

Here are my 5 rules on handling errors in Go.

## Rule 1 - Don't ignore the error

Sooner or later your function will fail and you will waste hours figuring why and restoring your program.
Just handle it. If you are in a rush or too tired - take a break.

```
package main

import (
    "fmt"
    "time"
)

func main() {
    // DO NOT IGNORE THE ERROR
    lucky, _ := ifItCanFailItWill()
}

func ifItCanFailItWill() (string, error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 == 0 {
        return "shinny desired value", nil
    }

    return "", fmt.Errorf("I will fail one day, handle me")
}
```



## Rule 2 - Return early

It may feel natural to focus on the "happy path" of the code execution first, but I prefer to start with validation and return the value at the end when everything went 100% fine.
I don't scale:

```
func nah() (string, error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 == 0 && isValid() {
        return "shinny desired value", nil
    }

    return "", fmt.Errorf("I will fail one day, handle me")
}
```
♛ PRO:

```
func earlyReturnRocks() (string, error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 > 0 {
        return "", fmt.Errorf("time dividability must be OCD compliant")
    }

    if !isValid() {
        return "", fmt.Errorf("a different custom, specific, helpful error message")
    }

    return "shinny desired value", nil
}
```



**Advantages**
- Easier to read
- Easier to add more validation
- Less nested code (especially in loops)
- A clear focus on safety and error handling
- Specific error message per `if` condition possible

## Rule 3 - Return value or Error (but not both)

I have seen developers using the return values in combination with an error at the same time. This is a bad practice. Avoid doing this.
Confusing:

```
func validateToken() (desiredValue string, expiredAt int, err error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 > 0 {
        // THE expiredAt (nowNs) SHOULD NOT BE RETURNED TOGETHER WITH THE ERR
        return "", nowNs, fmt.Errorf("token expired")
    }

    return "shinny desired value", 0, nil
}
```



**Disadvantages?**
- Unclear method signature
- One must reverse-engineer the method to know what values are returned and when

You are right, and sometimes you need to return some additional information about the error, in which case, create a new dedicated Error object.

♛ PRO:

```
package main

import (
    "fmt"
    "github.com/davecgh/go-spew/spew"
    "time"
)

func main() {
    value, err := validateToken()
    if err != nil {
        spew.Dump(err.Error())
    }

    spew.Dump(value)
}

// Compatible with error built-in interface.
//
// type error interface {
//  Error() string
// }
type TokenExpiredErr struct {
    expiredAt int
}

func (e TokenExpiredErr) Error() string {
    return fmt.Sprintf("token expired at block %d", e.expiredAt)
}

func validateToken() (desiredValue string, err error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 > 0 {
        return "", TokenExpiredErr{expiredAt: nowNs}
    }

    return "shinny desired value", nil
}
```



## Rule 4 - Log or Return (but not both)

When you log an error, you are handling it. Do NOT return the error back to the caller - forcing him to handle it as well!

[![Alt Text](https://res.cloudinary.com/practicaldev/image/fetch/s--S-nVw0UM--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)](https://res.cloudinary.com/practicaldev/image/fetch/s--S-nVw0UM--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)

Why? Because you don't want to log the same message twice or more:

```
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // validateToken() is already doing the logging,
    // but I didn't reverse engineer the method so I don't know about that
    // and now I will unfortunately end up with the same message being logged twice
    _, err := validateToken()
    if err != nil {
        // I have nowhere to return it, SO I RIGHTFULLY LOG IT
        // And I will not ignore a possible error writing err
        _, err = fmt.Fprint(os.Stderr, fmt.Errorf("validating token failed. %s", err.Error()))
        if err != nil {
            // Extremely rare, no other choice
            panic(err)
        }

        os.Exit(1)
    }
}

type TokenExpiredErr struct {
    expiredAt int
}

func (e TokenExpiredErr) Error() string {
    return fmt.Sprintf("token expired at block %d", e.expiredAt)
}

func validateToken() (desiredValue string, err error) {
    nowNs := time.Now().Nanosecond()
    if nowNs % 2 > 0 {
        // DO NOT LOG AND RETURN
        // DO NOT LOG AND RETURN
        // DO NOT LOG AND RETURN
        fmt.Printf("token validation failed. token expired at %d", nowNs)
        return "", TokenExpiredErr{expiredAt: nowNs}
    }

    return "shinny desired value", nil
}
```

Messy output when logging AND returning:

```
token validation failed. token expired at 115431493validating token failed. token expired at block 115431493
```

♛ PRO either logs OR returns:

```
validating token failed. token expired at block 599480733
```

## Rule 5 - Configure an `if err != nil` macro in your IDE

I couldn't keep typing the error check, so I just created a quick video guide I created on how to set it up in GoLand from Intellij. I bound the macro on my Mouse 4 button that I usually use for healing my Necromancer in Guild Wars 2 :)

## Do you like Go?

I am writing an eBook on `how to build a peer-to-peer system in Go` from scratch!
Check it out: https://web3.coach/#book
I tweet about it at: https://twitter.com/Web3Coach