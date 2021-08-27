# How to handle errors in Go?

# 如何处理 Go 中的错误？

## Practically and efficiently!

## 切实有效！

While handlings errors in Go is exceptionally annoying, I like the explicit error checks much more than throwing an exception 5 levels up the stack and hoping someone will catch it. I am looking at you Java!

虽然在 Go 中处理错误非常烦人，但我更喜欢显式错误检查，而不是在堆栈上抛出 5 级异常并希望有人能抓住它。我在看着你Java！

Here are my 5 rules on handling errors in Go.

这是我在 Go 中处理错误的 5 条规则。

## Rule 1 - Don't ignore the error

## 规则 1 - 不要忽略错误

Sooner or later your function will fail and you will waste hours figuring why and restoring your program.
Just handle it. If you are in a rush or too tired - take a break.

您的函数迟早会失败，您将浪费数小时弄清楚原因并恢复您的程序。
处理一下就好了如果您赶时间或太累了 - 休息一下。

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

## 规则 2 - 提前返回

It may feel natural to focus on the "happy path" of the code execution first, but I prefer to start with validation and return the value at the end when everything went 100% fine.
I don't scale:

首先关注代码执行的“快乐路径”可能感觉很自然，但我更喜欢从验证开始，并在一切顺利 100% 时返回值。
我不缩放：

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

♛ 专业：

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

**好处**
- 更容易阅读
- 更容易添加更多验证
- 更少的嵌套代码（尤其是在循环中）
- 明确关注安全和错误处理
- 每个“if”条件可能的特定错误消息

## Rule 3 - Return value or Error (but not both)

## 规则 3 - 返回值或错误（但不能同时返回）

I have seen developers using the return values in combination with an error at the same time. This is a bad practice. Avoid doing this.
Confusing:

我见过开发人员同时使用返回值和错误。这是一个不好的做法。避免这样做。
令人困惑：

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

**缺点？**
- 方法签名不明确
- 必须对方法进行逆向工程以了解返回什么值以及何时返回

You are right, and sometimes you need to return some additional information about the error, in which case, create a new dedicated Error object.

你是对的，有时你需要返回一些关于错误的附加信息，在这种情况下，创建一个新的专用 Error 对象。

♛ PRO:

♛ 专业：

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

## 规则 4 - 记录或返回（但不能同时进行）

When you log an error, you are handling it. Do NOT return the error back to the caller - forcing him to handle it as well!

当您记录错误时，您正在处理它。不要将错误返回给调用者 - 强迫他也处理它！

[![Alt Text](https://res.cloudinary.com/practicaldev/image/fetch/s--S-nVw0UM--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)](https://res.cloudinary.com/practicaldev/image/fetch/s--S-nVw0UM--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)

uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)](https://res.cloudinary.com/practicaldev/image/fetch/s--S-nVw0UM--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/ph3vyxeh9l0mtukitcdj.png)

Why? Because you don't want to log the same message twice or more:

为什么？因为您不想将同一条消息记录两次或更多次：

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

记录和返回时的混乱输出：

```
token validation failed.token expired at 115431493validating token failed.token expired at block 115431493
```


♛ PRO either logs OR returns:

♛ PRO 要么记录要么返回：

```
validating token failed.token expired at block 599480733
```


## Rule 5 - Configure an `if err != nil` macro in your IDE

## 规则 5 - 在您的 IDE 中配置一个 `if err != nil` 宏

I couldn't keep typing the error check, so I just created a quick video guide I created on how to set it up in GoLand from Intellij. I bound the macro on my Mouse 4 button that I usually use for healing my Necromancer in Guild Wars 2 :)

我无法继续输入错误检查，因此我创建了一个快速视频指南，介绍如何从 Intellij 在 GoLand 中进行设置。我在我的鼠标 4 按钮上绑定了宏，我通常用它来治疗激战 2 中的死灵法师 :)

## Do you like Go?

## 你喜欢 Go 吗？

I am writing an eBook on `how to build a peer-to-peer system in Go` from scratch!

Check it out: https://web3.coach/#book

I tweet about it at: https://twitter.com/Web3Coach 

我正在写一本关于“如何从头开始在 Go 中构建点对点系统”的电子书！
看看：https://web3.coach/#book
我在推特上介绍它：https://twitter.com/Web3Coach

