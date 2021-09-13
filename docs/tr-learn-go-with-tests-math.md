# Mathematics

# 数学

**[You can find all the code for this chapter here](https://github.com/quii/learn-go-with-tests/tree/main/math)**

**[你可以在这里找到本章的所有代码](https://github.com/quii/learn-go-with-tests/tree/main/math)**

For all the power of modern computers to perform huge sums at lightning speed, the average developer rarely uses any mathematics to do their job. But not today! Today we'll use mathematics to solve a *real* problem. And not boring mathematics - we're going to use trigonometry and vectors and all sorts of stuff that you always said you'd never have to use after highschool.

尽管现代计算机具有以闪电般的速度执行巨额计算的所有能力，但普通开发人员很少使用任何数学来完成他们的工作。但不是今天！今天我们将使用数学来解决一个*真实*的问题。而不是枯燥的数学——我们将使用三角学和向量以及你一直说高中毕业后永远不会使用的各种东西。

## The Problem

##  问题

You want to make an SVG of a clock. Not a digital clock - no, that would be easy - an *analogue* clock, with hands. You're not looking for anything fancy, just a nice function that takes a `Time` from the `time` package and spits out an SVG of a clock with all the hands - hour, minute and second - pointing in the right direction. How hard can that be?

您想制作时钟的 SVG。不是数字时钟 - 不，这很容易 - 一个*模拟*时钟，带指针。你不是在寻找任何花哨的东西，只是一个很好的函数，它从 `time` 包中获取一个 `Time` 并吐出一个时钟的 SVG，所有的指针 - 小时、分钟和秒 - 指向正确的方向。这有多难？

First we're going to need an SVG of a clock for us to play with. SVGs are a fantastic image format to manipulate programmatically because they're written as a series of shapes, described in XML. So this clock:

首先，我们需要一个时钟的 SVG 来玩。 SVG 是一种奇妙的图像格式，可以以编程方式进行操作，因为它们被编写为一系列形状，并以 XML 进行描述。所以这个时钟：

[![an svg of a clock](https://github.com/quii/learn-go-with-tests/raw/main/math/example_clock.svg)](https://github.com/quii/learn-go-with-tests/blob/main/math/example_clock.svg)

is described like this:

是这样描述的：

```
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

  <!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

  <!-- hour hand -->
  <line x1="150" y1="150" x2="114.150000" y2="132.260000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- minute hand -->
  <line x1="150" y1="150" x2="101.290000" y2="99.730000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- second hand -->
  <line x1="150" y1="150" x2="77.190000" y2="202.900000"
        style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
```

It's a circle with three lines, each of the lines starting in the middle of the circle (x=150, y=150), and ending some distance away.

它是一个包含三条线的圆，每条线都从圆的中间 (x=150, y=150) 开始，到一定距离结束。

So what we're going to do is reconstruct the above somehow, but change the lines so they point in the appropriate directions for a given time.

所以我们要做的是以某种方式重建上面的内容，但是改变线条，使它们在给定的时间内指向适当的方向。

## An Acceptance Test

## 验收测试

Before we get too stuck in, lets think about an acceptance test.

在我们陷入困境之前，让我们考虑一下验收测试。

Wait, you don't know what an acceptance test is yet. Look, let me try to explain.

等等，你还不知道验收测试是什么。看，让我试着解释一下。

Let me ask you: what does winning look like? How do we know we've finished work? TDD provides a good way of knowing when you've finished: when the test passes. Sometimes it's nice - actually, almost all of the time it's nice - to write a test that tells you when you've finished writing the whole usable feature. Not just a test that tells you that a particular function is working in the way you expect, but a test that tells you that the whole thing you're trying to achieve - the 'feature' - is complete.

让我问你：获胜是什么样子的？我们怎么知道我们已经完成了工作？ TDD 提供了一种了解何时完成的好方法：测试何时通过。有时很好 - 实际上，几乎所有时间都很好 - 编写一个测试，告诉您何时完成了整个可用功能的编写。不仅是测试告诉您特定功能以您期望的方式工作，而且还告诉您您试图实现的整个事情 - “功能” - 已经完成。

These tests are sometimes called 'acceptance tests', sometimes called 'feature test'. The idea is that you write a really high level test to describe what you're trying to achieve - a user clicks a button on a website, and they see a complete list of the Pokémon they've caught, for instance. When we've written that test, we can then write test more tests - unit tests - that build towards a working system that will pass the acceptance test. So for our example these tests might be about rendering a webpage with a button, testing route handlers on a web server, performing database look ups, etc. All of these things will be TDD'd, and all of them will go towards making the original acceptance test pass.

这些测试有时称为“验收测试”，有时称为“功能测试”。这个想法是你编写一个非常高级的测试来描述你想要实现的目标 - 例如，用户单击网站上的按钮，他们会看到他们捕获的 Pokémon 的完整列表。当我们编写了那个测试之后，我们就可以编写更多的测试——单元测试——构建一个将通过验收测试的工作系统。因此，对于我们的示例，这些测试可能是关于渲染带有按钮的网页、测试 Web 服务器上的路由处理程序、执行数据库查找等。所有这些都将是 TDD，并且所有这些都将用于使原始验收测试通过。

Something like this *classic* picture by Nat Pryce and Steve Freeman

类似于 Nat Pryce 和 Steve Freeman 的这张*经典*图片

[![img.png](https://github.com/quii/learn-go-with-tests/raw/main/TDD-outside-in.jpg)](https://github.com/quii/learn-go-with-tests/blob/main/TDD-outside-in.jpg)

学习测试/blob/main/TDD-outside-in.jpg）

Anyway, let's try and write that acceptance test - the one that will let us know when we're done.

无论如何，让我们尝试编写验收测试 - 完成后会通知我们的测试。

We've got an example clock, so let's think about what the important parameters are going to be.

我们有一个示例时钟，所以让我们考虑一下重要的参数是什么。

```
<line x1="150" y1="150" x2="114.150000" y2="132.260000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>
```

The centre of the clock (the attributes `x1` and `y1` for this line) is the same for each hand of the clock. The numbers that need to change for each hand of the clock - the parameters to whatever builds the SVG - are the `x2` and `y2` attributes. We'll need an X and a Y for each of the hands of the clock.

时钟的中心（这条线的属性“x1”和“y1”）对于时钟的每一根指针都是相同的。需要为时钟的每根指针更改的数字 - 构建 SVG 的任何参数 - 是 `x2` 和 `y2` 属性。我们需要一个 X 和一个 Y 来代表时钟的每个指针。

I *could* think about more parameters - the radius of the clockface circle, the size of the SVG, the colours of the hands, their shape, etc... but it's better to start off by solving a simple, concrete problem with a simple, concrete solution, and then to start adding parameters to make it generalised.

我*可以*考虑更多参数 - 表盘圆的半径、SVG 的大小、指针的颜色、它们的形状等......但最好从解决一个简单、具体的问题开始简单、具体的解决方案，然后开始添加参数使其泛化。

So we'll say that

所以我们会说

- every clock has a centre of (150, 150)
- the hour hand is 50 long
- the minute hand is 80 long
- the second hand is 90 long.

- 每个时钟的中心是 (150, 150)
- 时针长 50
- 分针长 80
- 秒针长 90。

A thing to note about SVGs: the origin - point (0,0) - is at the *top left* hand corner, not the *bottom left* as we might expect. It'll be important to remember this when we're working out where what numbers to plug in to our lines.

关于 SVG 需要注意的一点：原点 - 点 (0,0) - 位于*左上*手角，而不是我们预期的*左下*。当我们确定将哪些数字插入到我们的线路中时，记住这一点很重要。

Finally, I'm not deciding *how* to construct the SVG - we could use a template from the [`text/template`](https://golang.org/pkg/text/template/) package, or we could just send bytes into a `bytes.Buffer` or a writer. But we know we'll need those numbers, so let's focus on testing something that creates them.

最后，我不决定*如何*构建 SVG - 我们可以使用 [`text/template`](https://golang.org/pkg/text/template/) 包中的模板，或者我们也可以只需将字节发送到 `bytes.Buffer` 或写入器。但是我们知道我们需要这些数字，所以让我们专注于测试创建它们的东西。

### Write the test first

### 先写测试

So my first test looks like this:

所以我的第一个测试是这样的：

```
package clockface_test

import (
    "testing"
    "time"

    "github.com/gypsydave5/learn-go-with-tests/math/v1/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

    want := clockface.Point{X: 150, Y: 150 - 90}
    got := clockface.SecondHand(tm)

    if got != want {
        t.Errorf("Got %v, wanted %v", got, want)
    }
}
```

Remember how SVGs plot their coordinates from the top left hand corner? To place the second hand at midnight we expect that it hasn't moved from the centre of the clockface on the X axis - still 150 - and the Y axis is the length of the hand 'up' from the centre; 150 minus 90.

还记得 SVG 如何从左上角绘制它们的坐标吗？为了在午夜放置秒针，我们希望它没有从 X 轴上的钟面中心移动 - 仍然是 150 - 而 Y 轴是指针从中心“向上”的长度； 150 减 90。

### Try to run the test

### 尝试运行测试

This drives out the expected failures around the missing functions and types:

这消除了围绕缺少的功能和类型的预期故障：

```
--- FAIL: TestSecondHandAtMidnight (0.00s)
./clockface_test.go:13:10: undefined: clockface.Point
./clockface_test.go:14:9: undefined: clockface.SecondHand
```

So a `Point` where the tip of the second hand should go, and a function to get it.

所以一个“点”应该放在秒针的尖端，以及一个获取它的函数。

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

Let's implement those types to get the code to compile

让我们实现这些类型来编译代码

```
package clockface

import "time"

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
    X float64
    Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
    return Point{}
}
```

and now we get:

现在我们得到：

```
--- FAIL: TestSecondHandAtMidnight (0.00s)
    clockface_test.go:17: Got {0 0}, wanted {150 60}
FAIL
exit status 1
FAIL    github.com/gypsydave5/learn-go-with-tests/math/v1/clockface    0.006s
```

### Write enough code to make it pass

### 编写足够的代码使其通过

When we get the expected failure, we can fill in the return value of `SecondHand`:

当我们得到预期的失败时，我们可以填写`SecondHand`的返回值：

```
// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
    return Point{150, 60}
}
```

Behold, a passing test.

看，一个通过的测试。

```
PASS
ok          clockface    0.006s
```

### Refactor

### 重构

No need to refactor yet - there's barely enough code!

还不需要重构——几乎没有足够的代码！

### Repeat for new requirements

### 重复新要求

We probably need to do some work here that doesn't just involve returning a clock that shows midnight for every time...

我们可能需要在这里做一些工作，而不仅仅是返回一个每次都显示午夜的时钟......

### Write the test first

### 先写测试

```
func TestSecondHandAt30Seconds(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

    want := clockface.Point{X: 150, Y: 150 + 90}
    got := clockface.SecondHand(tm)

    if got != want {
        t.Errorf("Got %v, wanted %v", got, want)
    }
}
```

Same idea, but now the second hand is pointing *downwards* so we *add* the length to the Y axis.

同样的想法，但现在秒针指向*向下*，所以我们*添加* Y 轴的长度。

This will compile... but how do we make it pass?

这将编译...但我们如何使它通过？

## Thinking time

## 思考时间

How are we going to solve this problem?

我们将如何解决这个问题？

Every minute the second hand goes through the same 60 states, pointing in 60 different directions. When it's 0 seconds it points to the top of the clockface, when it's 30 seconds it points to the bottom of the clockface. Easy enough. 

每分钟秒针都会经过相同的 60 个状态，指向 60 个不同的方向。 0 秒时指向表盘顶部，30 秒时指向表盘底部。很容易。

So if I wanted to think about in what direction the second hand was pointing at, say, 37 seconds, I'd want the angle between 12 o'clock and 37/60ths around the circle. In degrees this is `(360 / 60 ) * 37 = 222`, but it's easier just to remember that it's `37/60` of a complete rotation.

因此，如果我想考虑秒针指向的方向，例如 37 秒，我会想要 12 点钟和 37/60 点之间围绕圆圈的角度。以度为单位，这是 `(360 / 60 ) * 37 = 222`，但记住它是完整旋转的 `37/60` 会更容易。

But the angle is only half the story; we need to know the X and Y coordinate that the tip of the second hand is pointing at. How can we work that out?

但角度只是故事的一半；我们需要知道秒针指尖所指的 X 和 Y 坐标。我们怎样才能解决这个问题？

## Math

##  数学

Imagine a circle with a radius of 1 drawn around the origin - the coordinate `0, 0`.

想象一个围绕原点绘制的半径为 1 的圆 - 坐标为“0, 0”。

[![picture of the unit circle](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle.png)

quii/learn-go-with-tests/blob/main/math/images/unit_circle.png）

This is called the 'unit circle' because... well, the radius is 1 unit!

这被称为“单位圆”，因为......好吧，半径是 1 个单位！

The circumference of the circle is made of points on the grid - more coordinates. The x and y components of each of these coordinates form a triangle, the hypotenuse of which is always 1 - the radius of the circle

圆的圆周由网格上的点组成——更多的坐标。每个坐标的 x 和 y 分量形成一个三角形，其斜边始终为 1 - 圆的半径

[![picture of the unit circle with a point defined on the circumference](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_coords.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_coords.png)

://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_coords.png）

Now, trigonometry will let us work out the lengths of X and Y for each triangle if we know the angle they make with the origin. The X coordinate will be cos(a), and the Y coordinate will be sin(a), where a is the angle made between the line and the (positive) x axis.

现在，如果我们知道它们与原点所成的角度，三角学将让我们计算出每个三角形的 X 和 Y 的长度。 X 坐标为 cos(a)，Y 坐标为 sin(a)，其中 a 是直线与（正）x 轴之间的夹角。

[![picture of the unit circle with the x and y elements of a ray defined as cos(a) and sin(a) respectively, where a is the angle made by the ray with the x axis](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_params.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_params.png)

.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_params.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_params.png)

(If you don't believe this, [go and look at Wikipedia...](https://en.wikipedia.org/wiki/Sine#Unit_circle_definition))

（如果你不相信这个，[去看看维基百科...](https://en.wikipedia.org/wiki/Sine#Unit_circle_definition))

One final twist - because we want to measure the angle from 12 o'clock rather than from the X axis (3 o'clock), we need to swap the axis around; now x = sin(a) and y = cos(a).

最后一点——因为我们想从 12 点钟而不是从 X 轴（3 点钟）测量角度，我们需要交换轴；现在 x = sin(a) 和 y = cos(a)。

[![unit circle ray defined from by angle from y axis](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_12_oclock.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_12_oclock.png)

/github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_12_oclock.png)

So now we know how to get the angle of the second hand (1/60th of a circle for each second) and the X and Y coordinates. We'll need functions for both `sin` and `cos`.

所以现在我们知道如何获得秒针的角度（每秒一个圆的 1/60）以及 X 和 Y 坐标。我们需要`sin` 和`cos` 的函数。

## `math`

## `数学`

Happily the Go `math` package has both, with one small snag we'll need to get our heads around; if we look at the description of [`math.Cos`](https://golang.org/pkg/math/#Cos):

令人高兴的是，Go`math` 包两者兼具，我们需要解决一个小问题；如果我们看一下 [`math.Cos`](https://golang.org/pkg/math/#Cos) 的描述：

> Cos returns the cosine of the radian argument x.

> Cos 返回弧度参数 x 的余弦值。

It wants the angle to be in radians. So what's a radian? Instead of  defining the full turn of a circle to be made up of 360 degrees, we  define a full turn as being 2π radians. There are good reasons to do  this that we won't go in to.[^2]

它希望角度以弧度为单位。那么什么是弧度？我们没有将圆的整圈定义为 360 度，而是将整圈定义为 2π 弧度。这样做有充分的理由，我们不会讨论。[^2]

Now that we've done some reading, some learning and some thinking, we can write our next test.

现在我们已经完成了一些阅读、一些学习和一些思考，我们可以编写下一个测试了。

### Write the test first

### 先写测试

All this maths is hard and confusing. I'm not confident I understand what's going on - so let's write a test! We don't need to solve the whole problem in one go - let's start off with working out the correct angle, in radians, for the second hand at a particular time.

所有这些数学都很困难且令人困惑。我不确定我是否理解发生了什么 - 所以让我们编写一个测试！我们不需要一次性解决整个问题——让我们从在特定时间为秒针计算出正确的角度（以弧度为单位）开始。

I'm going to write these tests *within* the `clockface` package; they may never get exported, and they may get deleted (or moved) once I have a better grip on what's going on.

我将在 `clockface` 包中*内*编写这些测试；它们可能永远不会被导出，一旦我更好地掌握了正在发生的事情，它们可能会被删除（或移动）。

I'm also going to *comment out* the acceptance test that I was working on while I'm working on these tests - I don't want to get distracted by that test while I'm getting this one to pass.

我还将*注释*我在进行这些测试时正在进行的验收测试 - 我不想在我通过该测试时被该测试分心。

```
package clockface

import (
    "math"
    "testing"
    "time"
)

func TestSecondsInRadians(t *testing.T) {
    thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
    want := math.Pi
    got := secondsInRadians(thirtySeconds)

    if want != got {
        t.Fatalf("Wanted %v radians, but got %v", want, got)
    }
}
```

Here we're testing that 30 seconds past the minute should put the second hand at halfway around the clock. And it's our first use of the `math` package! If a full turn of a circle is 2π radians, we know that halfway round should just be π radians. `math.Pi` provides us with a value for π.

在这里，我们正在测试一分钟后 30 秒应将秒针置于半小时。这是我们第一次使用 `math` 包！如果圆的一整圈是 2π 弧度，我们知道半圆应该是 π 弧度。 `math.Pi` 为我们提供了一个 π 值。

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:12:9: undefined: secondsInRadians
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func secondsInRadians(t time.Time) float64 {
    return 0
}
clockface_test.go:15: Wanted 3.141592653589793 radians, but got 0
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func secondsInRadians(t time.Time) float64 {
    return math.Pi
}
PASS
ok      clockface    0.011s
```

### Refactor

### 重构

Nothing needs refactoring yet

什么都不需要重构

### Repeat for new requirements

### 重复新要求

Now we can extend the test to cover a few more scenarios. I'm going to skip forward a bit and show some already refactored test code - it should be clear enough how I got where I want to.

现在我们可以扩展测试以涵盖更多场景。我将向前跳过一点并展示一些已经重构的测试代码 - 应该很清楚我是如何到达我想要的地方的。

```
func TestSecondsInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(0, 0, 30), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(0, 0, 45), (math.Pi / 2) * 3},
        {simpleTime(0, 0, 7), (math.Pi / 30) * 7},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := secondsInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

I added a couple of helper functions to make writing this table based test a little less tedious. `testName` converts a time into a digital watch format (HH:MM:SS), and `simpleTime` constructs a `time.Time` using only the parts we actually care about (again, hours, minutes and seconds).[^ 1] Here they are:

我添加了几个辅助函数，使编写这个基于表的测试变得不那么乏味。 `testName` 将时间转换为数字手表格式（HH:MM:SS），而 `simpleTime` 仅使用我们真正关心的部分（同样是小时、分钟和秒）构造一个 `time.Time`。[^ 1] 他们是：

```
func simpleTime(hours, minutes, seconds int) time.Time {
    return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
    return t.Format("15:04:05")
}
```

These two functions should help make these tests (and future tests) a little easier to write and maintain.

这两个函数应该有助于使这些测试（和未来的测试）更容易编写和维护。

This gives us some nice test output:

这给了我们一些不错的测试输出：

```
clockface_test.go:24: Wanted 0 radians, but got 3.141592653589793

clockface_test.go:24: Wanted 4.71238898038469 radians, but got 3.141592653589793
```

Time to implement all of that maths stuff we were talking about above:

是时候实现我们上面讨论的所有数学内容了：

```
func secondsInRadians(t time.Time) float64 {
    return float64(t.Second()) * (math.Pi / 30)
}
```

One second is (2π / 60) radians... cancel out the 2 and we get π/30 radians. Multiply that by the number of seconds (as a `float64`) and we should now have all the tests passing...

一秒是 (2π / 60) 弧度...抵消 2，我们得到 π/30 弧度。将其乘以秒数（作为“float64”），我们现在应该让所有测试都通过了……

```
clockface_test.go:24: Wanted 3.141592653589793 radians, but got 3.1415926535897936
```

Wait, what?

等等，什么？

### Floats are horrible

### 浮动是可怕的

Floating point arithmetic is [notoriously inaccurate](https://0.30000000000000004.com/). Computers can only really handle integers, and rational numbers to some extent. Decimal numbers start to become inaccurate, especially when we factor them up and down as we are in the `secondsInRadians` function. By dividing `math.Pi` by 30 and then by multiplying it by 30 we've ended up with *a number that's no longer the same as `math.Pi`*.

浮点运算是 [众所周知的不准确](https://0.30000000000000004.com/)。计算机只能在一定程度上真正处理整数和有理数。十进制数开始变得不准确，尤其是当我们在“secondsInRadians”函数中将它们上下分解时。通过将`math.Pi`除以30，然后再乘以30，我们最终得到了*一个不再与`math.Pi`相同的数字*。

There are two ways around this:

有两种方法可以解决这个问题：

1. Live with it
2. Refactor our function by refactoring our equation

1. 与它共存
2. 通过重构我们的方程来重构我们的函数

Now (1) may not seem all that appealing, but it's often the only way to make floating point equality work. Being inaccurate by some infinitesimal fraction is frankly not going to matter for the purposes of drawing a clockface, so we could write a function that defines a 'close enough' equality for our angles. But there's a simple way we can get the accuracy back: we rearrange the equation so that we're no longer dividing down and then multiplying up. We can do it all by just dividing.

现在 (1) 可能看起来并不那么吸引人，但它通常是使浮点相等工作的唯一方法。坦率地说，对于绘制表盘的目的而言，由于某个无穷小分数不准确并不重要，因此我们可以编写一个函数，为我们的角度定义“足够接近”的相等性。但是有一个简单的方法可以让我们恢复准确度：我们重新排列方程，这样我们就不再先除以再乘以。我们只需分开就可以做到这一切。

So instead of

所以代替

```
numberOfSeconds * π / 30
```

we can write

我们可以写

```
π / (30 / numberOfSeconds)
```

which is equivalent.

这是等效的。

In Go:

在围棋中：

```
func secondsInRadians(t time.Time) float64 {
    return (math.Pi / (30 / (float64(t.Second()))))
}
```

And we get a pass.

我们得到了通行证。

```
PASS
ok      clockface     0.005s
```

It should all look [something like this](https://github.com/quii/learn-go-with-tests/tree/main/math/v3/clockface).

它应该看起来 [像这样](https://github.com/quii/learn-go-with-tests/tree/main/math/v3/clockface)。

### A note on dividing by zero

### 关于除以零的注释

Computers often don't like dividing by zero because infinity is a bit strange.

计算机通常不喜欢除以零，因为无穷大有点奇怪。

In Go if you try to explicitly divide by zero you will get a compilation error.

在 Go 中，如果您尝试显式地除以零，则会出现编译错误。

```
package main

import (
    "fmt"
)

func main() {
    fmt.Println(10.0 / 0.0) // fails to compile
}
```

Obviously the compiler can't always predict that you'll divide by zero, such as our `t.Second()`

显然编译器不能总是预测你会被零除，比如我们的 `t.Second()`

Try this

尝试这个

```
func main() {
  fmt.Println(10.0/zero())
}

func zero() float64 {
  return 0.0
}
```

It will print `+Inf` (infinity). Dividing by +Inf seems to result in zero and we can see this with the following:

它将打印`+Inf`（无穷大）。除以 +Inf 似乎结果为零，我们可以通过以下方式看到这一点：

```
package main

import (
    "fmt"
    "math"
)

func main() {
  fmt.Println(secondsinradians())
}

func zero() float64 {
  return 0.0
}

func secondsinradians() float64 {
    return (math.Pi / (30 / (float64(zero()))))
}
```

### Repeat for new requirements

### 重复新要求

So we've got the first part covered here - we know what angle the second hand will be pointing at in radians. Now we need to work out the coordinates.

所以我们已经在这里介绍了第一部分 - 我们知道秒针将以弧度表示的角度。现在我们需要计算出坐标。

Again, let's keep this as simple as possible and only work with the *unit circle*; the circle with a radius of 1. This means that our hands will all have a length of one but, on the bright side, it means that the maths will be easy for us to swallow.

再一次，让我们尽可能简单，只使用*单位圆*；半径为 1 的圆。这意味着我们的手的长度都为 1，但从好的方面来说，这意味着我们很容易接受数学运算。

### Write the test first

### 先写测试

```
func TestSecondHandVector(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := secondHandPoint(c.time)
            if got != c.point {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:40:11: undefined: secondHandPoint
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func secondHandPoint(t time.Time) Point {
    return Point{}
}
clockface_test.go:42: Wanted {0 -1} Point, but got {0 0}
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func secondHandPoint(t time.Time) Point {
    return Point{0, -1}
}
PASS
ok      clockface    0.007s
```

### Repeat for new requirements

### 重复新要求

```
func TestSecondHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
        {simpleTime(0, 0, 45), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := secondHandPoint(c.time)
            if got != c.point {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_test.go:43: Wanted {-1 0} Point, but got {0 -1}
```

### Write enough code to make it pass

### 编写足够的代码使其通过

Remember our unit circle picture?

还记得我们的单位圆图吗？

[![picture of the unit circle with the x and y elements of a ray defined as cos(a) and sin(a) respectively, where a is the angle made by the ray with the x axis](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_params.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_params.png)



Also recall that we want to measure the angle from 12 o'clock which  is the Y axis instead of from the X axis which we would like measuring  the angle between the second hand and 3 o'clock.

还记得我们要从 12 点钟方向测量角度，这是 Y 轴，而不是从 X 轴测量角度，我们想测量秒针和 3 点钟方向之间的角度。

[![unit circle ray defined from by angle from y axis](https://github.com/quii/learn-go-with-tests/raw/main/math/images/unit_circle_12_oclock.png)](https://github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_12_oclock.png)

/github.com/quii/learn-go-with-tests/blob/main/math/images/unit_circle_12_oclock.png)

We now want the equation that produces X and Y. Let's write it into seconds:

我们现在想要产生 X 和 Y 的方程。让我们把它写成秒：

```
func secondHandPoint(t time.Time) Point {
    angle := secondsInRadians(t)
    x := math.Sin(angle)
    y := math.Cos(angle)

    return Point{x, y}
}
```

Now we get

现在我们得到

```
clockface_test.go:43: Wanted {0 -1} Point, but got {1.2246467991473515e-16 -1}

clockface_test.go:43: Wanted {-1 0} Point, but got {-1 -1.8369701987210272e-16}
```

Wait, what (again)? Looks like we've been cursed by the floats once more - both of those unexpected numbers are *infinitesimal* - way down at the 16th decimal place. So again we can either choose to try to increase precision, or to just say that they're roughly equal and get on with our lives.

等等，什么（又）？看起来我们又一次被浮点数诅咒了——这两个意想不到的数字都是*无穷小的*——小数点后第16位。因此，我们再次可以选择尝试提高精确度，或者只是说它们大致相等并继续我们的生活。

One option to increase the accuracy of these angles would be to use the rational type `Rat` from the `math/big` package. But given the objective is to draw an SVG and not land on the moon landings I think we can live with a bit of fuzziness.

提高这些角度的准确性的一种选择是使用“math/big”包中的有理类型“Rat”。但鉴于目标是绘制 SVG 而不是登陆月球，我认为我们可以忍受一些模糊。

```
func TestSecondHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
        {simpleTime(0, 0, 45), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := secondHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point) {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}

func roughlyEqualFloat64(a, b float64) bool {
    const equalityThreshold = 1e-7
    return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
    return roughlyEqualFloat64(a.X, b.X) &&
        roughlyEqualFloat64(a.Y, b.Y)
}
```

We've defined two functions to define approximate equality between two `Points` - they'll work if the X and Y elements are within 0.0000001 of each other. That's still pretty accurate.

我们定义了两个函数来定义两个“点”之间的近似相等性——如果 X 和 Y 元素彼此在 0.0000001 以内，它们就会起作用。这还是很准确的。

And now we get:

现在我们得到：

```
PASS
ok      clockface    0.007s
```

### Refactor

### 重构

I'm still pretty happy with this.

我对此还是很满意的。

Here's [what it looks like now](https://github.com/quii/learn-go-with-tests/tree/main/math/v4/clockface)

这是[现在的样子](https://github.com/quii/learn-go-with-tests/tree/main/math/v4/clockface)

### Repeat for new requirements

### 重复新要求

Well, saying *new* isn't entirely accurate - really what we can do now is get that acceptance test passing! Let's remind ourselves of what it looks like:

好吧，说 *new* 并不完全准确——我们现在能做的就是让验收测试通过！让我们回忆一下它的样子：

```
func TestSecondHandAt30Seconds(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

    want := clockface.Point{X: 150, Y: 150 + 90}
    got := clockface.SecondHand(tm)

    if got != want {
        t.Errorf("Got %v, wanted %v", got, want)
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_acceptance_test.go:28: Got {150 60}, wanted {150 240}
```

### Write enough code to make it pass

### 编写足够的代码使其通过

We need to do three things to convert our unit vector into a point on the SVG:

我们需要做三件事来将我们的单位向量转换成 SVG 上的一个点：

1. Scale it to the length of the hand
2. Flip it over the X axis because to account for the SVG having an origin in the top left hand corner
3. Translate it to the right position (so that it's coming from an origin of (150,150))

1.将其缩放到手的长度
2. 将它翻转到 X 轴上，因为 SVG 的原点在左上角
3. 将其翻译到正确的位置（使其来自 (150,150) 的原点）

Fun times!

娱乐时间！

```
// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
    p := secondHandPoint(t)
    p = Point{p.X * 90, p.Y * 90}   // scale
    p = Point{p.X, -p.Y}            // flip
    p = Point{p.X + 150, p.Y + 150} // translate
    return p
}
```

Scale, flip, and translate in exactly that order. Hooray maths!

完全按照该顺序缩放、翻转和平移。数学万岁！

```
PASS
ok      clockface    0.007s
```

### Refactor

### 重构

There's a few magic numbers here that should get pulled out as constants, so let's do that

这里有一些神奇的数字应该作为常量提取出来，所以让我们这样做

```
const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
    p := secondHandPoint(t)
    p = Point{p.X * secondHandLength, p.Y * secondHandLength}
    p = Point{p.X, -p.Y}
    p = Point{p.X + clockCentreX, p.Y + clockCentreY} //translate
    return p
}
```

## Draw the clock

## 绘制时钟

Well... the second hand anyway...

嗯...反正二手...

Let's do this thing - because there's nothing worse than not delivering some value when it's just sitting there waiting to get out into the world to dazzle people. Let's draw a second hand!

让我们做这件事——因为当它只是坐在那里等待进入世界让人们眼花缭乱时，没有什么比不提供一些价值更糟糕的了。让我们画第二只手！

We're going to stick a new directory under our main `clockface` package directory, called (confusingly), `clockface`. In there we'll put the `main` package that will create the binary that will build an SVG:

我们将在我们的主 `clockface` 包目录下粘贴一个新目录，称为（令人困惑的）`clockface`。在那里，我们将放置 `main` 包，该包将创建用于构建 SVG 的二进制文件：

```
├── clockface
│   └── main.go
├── clockface.go
├── clockface_acceptance_test.go
└── clockface_test.go
```

Inside `main.go`, you'll start with this code but change the import for the clockface package to point at your own version:

在 `main.go` 中，您将从以下代码开始，但更改 clockface 包的导入以指向您自己的版本：

```
package main

import (
    "fmt"
    "io"
    "os"
    "time"

    "github.com/quii/learn-go-with-tests/math/clockface" // REPLACE THIS!
)

func main() {
    t := time.Now()
    sh := clockface.SecondHand(t)
    io.WriteString(os.Stdout, svgStart)
    io.WriteString(os.Stdout, bezel)
    io.WriteString(os.Stdout, secondHandTag(sh))
    io.WriteString(os.Stdout, svgEnd)
}

func secondHandTag(p clockface.Point) string {
    return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pX, pY)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
```

Oh boy am I not trying to win any prizes for beautiful code with *this* mess - but it does the job. It's writing an SVG out to `os.Stdout` - one string at a time.

哦，天哪，我是不是不想因为*这个* 混乱的漂亮代码而赢得任何奖品——但它确实做到了。它正在将一个 SVG 输出到 `os.Stdout` - 一次一个字符串。

If we build this

如果我们建立这个

```
go build
```

and run it, sending the output into a file

并运行它，将输出发送到文件中

```
./clockface > clock.svg
```

We should see something like 

我们应该看到类似的东西

[![a clock with only a second hand](https://github.com/quii/learn-go-with-tests/raw/main/math/v6/clockface/clockface/clock.svg)](https://github.com/quii/learn-go-with-tests/blob/main/math/v6/clockface/clockface/clock.svg)



And this is [how the code looks](https://github.com/quii/learn-go-with-tests/tree/main/math/v6/clockface).

这就是[代码的样子](https://github.com/quii/learn-go-with-tests/tree/main/math/v6/clockface)。

### Refactor

### 重构

This stinks. Well, it doesn't quite *stink* stink, but I'm not happy about it.

这很臭。好吧，它并不很*臭*臭，但我对此并不满意。

1. That whole `SecondHand` function is *super* tied to being an SVG... without mentioning SVGs or actually producing an SVG...
2. ... while at the same time I'm not testing any of my SVG code.

1. 整个`SecondHand` 函数是*super* 绑定到成为一个SVG...没有提到SVGs 或实际生成一个SVG...
2. ... 同时我没有测试我的任何 SVG 代码。

Yeah, I guess I screwed up. This feels wrong. Let's try to recover with a more SVG-centric test.

是的，我想我搞砸了。这感觉不对。让我们尝试通过更以 SVG 为中心的测试来恢复。

What are our options? Well, we could try testing that the characters spewing out of the `SVGWriter` contain things that look like the sort of SVG tag we're expecting for a particular time. For instance:

我们有哪些选择？好吧，我们可以尝试测试从 `SVGWriter` 中喷出的字符是否包含看起来像我们在特定时间期望的那种 SVG 标签的东西。例如：

```
func TestSVGWriterAtMidnight(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

    var b strings.Builder
    clockface.SVGWriter(&b, tm)
    got := b.String()

    want := `<line x1="150" y1="150" x2="150" y2="60"`

    if !strings.Contains(got, want) {
        t.Errorf("Expected to find the second hand %v, in the SVG output %v", want, got)
    }
}
```

But is this really an improvement?

但这真的是一种进步吗？

Not only will it still pass if I don't produce a valid SVG (as it's only testing that a string appears in the output), but it will also fail if I make the smallest, unimportant change to that string - if I add an extra space between the attributes, for instance.

如果我不生成有效的 SVG，它不仅会通过（因为它只是测试输出中是否出现一个字符串），而且如果我对该字符串进行最小的、不重要的更改，它也会失败 - 如果我添加一个例如，属性之间的额外空间。

The *biggest* smell is that I'm testing a data structure - XML - by looking at its representation as a series of characters - as a string. This is *never*, *ever* a good idea as it produces problems just like the ones I outline above: a test that's both too fragile and not sensitive enough. A test that's testing the wrong thing!

*最大的*气味是我正在测试数据结构 - XML - 通过将其表示形式视为一系列字符 - 作为字符串。这是*永远*，*永远*一个好主意，因为它会产生与我上面概述的问题一样的问题：既太脆弱又不够灵敏的测试。一个测试错误的东西！

So the only solution is to test the output *as XML*. And to do that we'll need to parse it.

所以唯一的解决方案是测试输出 *as XML*。为此，我们需要解析它。

## Parsing XML

## 解析 XML

[`encoding/xml`](https://godoc.org/encoding/xml) is the Go package that can handle all things to do with simple XML parsing.

[`encoding/xml`](https://godoc.org/encoding/xml) 是 Go 包，可以处理所有与简单 XML 解析有关的事情。

The function [`xml.Unmarshall`](https://godoc.org/encoding/xml#Unmarshal) takes a `[]byte` of XML data, and a pointer to a struct for it to get unmarshalled in to.

函数 [`xml.Unmarshall`](https://godoc.org/encoding/xml#Unmarshal) 接受一个 `[]byte` 的 XML 数据，以及一个指向结构的指针，以便对其进行解组。

So we'll need a struct to unmarshall our XML into. We could spend some time working out what the correct names for all of the nodes and attributes, and how to write the correct structure but, happily, someone has written [`zek`](https://github.com/miku/zek) a program that will automate all of that hard work for us. Even better, there's an online version at https://www.onlinetool.io/xmltogo/. Just paste the SVG from the top of the file into one box and - bam - out pops:

所以我们需要一个结构来将我们的 XML 解组。我们可以花一些时间来确定所有节点和属性的正确名称，以及如何编写正确的结构，但令人高兴的是，有人写了 [`zek`](https://github.com/miku/zek) 一个可以为我们自动完成所有艰苦工作的程序。更好的是，https://www.onlinetool.io/xmltogo/ 上有一个在线版本。只需将文件顶部的 SVG 粘贴到一个框中，然后 - bam - 弹出：

```
type Svg struct {
    XMLName xml.Name `xml:"svg"`
    Text    string   `xml:",chardata"`
    Xmlns   string   `xml:"xmlns,attr"`
    Width   string   `xml:"width,attr"`
    Height  string   `xml:"height,attr"`
    ViewBox string   `xml:"viewBox,attr"`
    Version string   `xml:"version,attr"`
    Circle  struct {
        Text  string `xml:",chardata"`
        Cx    string `xml:"cx,attr"`
        Cy    string `xml:"cy,attr"`
        R     string `xml:"r,attr"`
        Style string `xml:"style,attr"`
    } `xml:"circle"`
    Line []struct {
        Text  string `xml:",chardata"`
        X1    string `xml:"x1,attr"`
        Y1    string `xml:"y1,attr"`
        X2    string `xml:"x2,attr"`
        Y2    string `xml:"y2,attr"`
        Style string `xml:"style,attr"`
    } `xml:"line"`
}
```

We could make adjustments to this if we needed to (like changing the name of the struct to `SVG`) but it's definitely good enough to start us off. Paste the struct into the `clockface_test` file and let's write a test with it:

如果需要，我们可以对此进行调整（例如将结构的名称更改为 `SVG`），但这绝对足以让我们开始。将结构粘贴到 `clockface_test` 文件中，让我们用它编写一个测试：

```
func TestSVGWriterAtMidnight(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

    b := bytes.Buffer{}
    clockface.SVGWriter(&b, tm)

    svg := Svg{}
    xml.Unmarshal(b.Bytes(), &svg)

    x2 := "150"
    y2 := "60"

    for _, line := range svg.Line {
        if line.X2 == x2 && line.Y2 == y2 {
            return
        }
    }

    t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", x2, y2, b.String())
}
```

We write the output of `clockface.SVGWriter` to a `bytes.Buffer` and then `Unmarshall` it into an `Svg`. We then look at each `Line` in the `Svg` to see if any of them have the expected `X2` and `Y2` values. If we get a match we return early (passing the test); if not we fail with a (hopefully) informative message.

我们将 `clockface.SVGWriter` 的输出写入一个 `bytes.Buffer`，然后将它`Unmarshall` 写入一个 `Svg`。然后我们查看 `Svg` 中的每个 `Line`，看看它们中是否有任何一个具有预期的 `X2` 和 `Y2` 值。如果我们得到匹配，我们会提前返回（通过测试）；如果不是，我们会以（希望）信息丰富的消息失败。

```
./clockface_acceptance_test.go:41:2: undefined: clockface.SVGWriter
```

Looks like we'd better write that `SVGWriter`...

看起来我们最好写那个 `SVGWriter`...

```
package clockface

import (
    "fmt"
    "io"
    "time"
)

const (
    secondHandLength = 90
    clockCentreX     = 150
    clockCentreY     = 150
)

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
    io.WriteString(w, svgStart)
    io.WriteString(w, bezel)
    secondHand(w, t)
    io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
    p := secondHandPoint(t)
    p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
    p = Point{p.X, -p.Y}                                      // flip
    p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pX, pY)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
```

The most beautiful SVG writer? No. But hopefully it'll do the job...

最漂亮的 SVG 作家？不，但希望它能完成这项工作......

```
clockface_acceptance_test.go:56: Expected to find the second hand with x2 of 150 and y2 of 60, in the SVG output <?xml version="1.0" encoding="UTF-8" standalone="no"?>
    <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
    <svg xmlns="http://www.w3.org/2000/svg"
         width="100%"
         height="100%"
         viewBox="0 0 300 300"
         version="2.0"><circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/><line x1="150" y1="150" x2="150.000000" y2="60.000000" style="fill:none;stroke:#f00;stroke-width:3px;"/></svg>
```

Oooops! The `%f` format directive is printing our coordinates to the default level of precision - six decimal places. We should be explicit as to what level of precision we're expecting for the coordinates. Let's say three decimal places.

哎呀！ `%f` 格式指令将我们的坐标打印到默认精度级别 - 小数点后六位。我们应该明确我们期望坐标的精度水平。让我们说三位小数。

```
     fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pX, pY)
```

And after we update our expectations in the test

在我们在测试中更新我们的期望之后

```
     x2 := "150.000"
    y2 := "60.000"
```

We get:

我们得到：

```
PASS
ok      clockface    0.006s
```

We can now shorten our `main` function:

我们现在可以缩短我们的 main 函数：

```
package main

import (
    "os"
    "time"

    "github.com/gypsydave5/learn-go-with-tests/math/v7b/clockface"
)

func main() {
    t := time.Now()
    clockface.SVGWriter(os.Stdout, t)
}
```

This is what [things should look like now](https://github.com/quii/learn-go-with-tests/blob/main/math/v7b/clockface).

这就是[事情现在应该是什么样子](https://github.com/quii/learn-go-with-tests/blob/main/math/v7b/clockface)。

And we can write a test for another time following the same pattern, but not before...

我们可以按照相同的模式再次编写测试，但之前不能......

### Refactor

### 重构

Three things stick out:

突出三点：

1. We're not really testing for all of the information we need to ensure is present - what about the `x1` values, for instance?
2. Also, those attributes for `x1` etc. aren't really `strings` are they? They're numbers!
3. Do I really care about the `style` of the hand? Or, for that matter, the empty `Text` node that's been generated by `zak`?

1. 我们并没有真正测试我们需要确保存在的所有信息 - 例如，`x1` 值怎么样？
2. 另外，`x1` 等的那些属性并不是真正的 `strings`，是吗？他们是数字！
3. 我真的很在意手的‘风格’吗？或者，就此而言，由`zak` 生成的空`Text` 节点？

We can do better. Let's make a few adjustments to the `Svg` struct, and the tests, to sharpen everything up.

我们可以做得更好。让我们对 `Svg` 结构和测试进行一些调整，以加强一切。

```
type SVG struct {
    XMLName xml.Name `xml:"svg"`
    Xmlns   string   `xml:"xmlns,attr"`
    Width   string   `xml:"width,attr"`
    Height  string   `xml:"height,attr"`
    ViewBox string   `xml:"viewBox,attr"`
    Version string   `xml:"version,attr"`
    Circle  Circle   `xml:"circle"`
    Line    []Line   `xml:"line"`
}

type Circle struct {
    Cx float64 `xml:"cx,attr"`
    Cy float64 `xml:"cy,attr"`
    R  float64 `xml:"r,attr"`
}

type Line struct {
    X1 float64 `xml:"x1,attr"`
    Y1 float64 `xml:"y1,attr"`
    X2 float64 `xml:"x2,attr"`
    Y2 float64 `xml:"y2,attr"`
}
```

Here I've

我在这里

- Made the important parts of the struct named types -- the `Line` and the `Circle`
- Turned the numeric attributes into `float64`s instead of `string`s.
- Deleted unused attributes like `Style` and `Text` 

- 制作了结构命名类型的重要部分——`Line` 和 `Circle`
- 将数字属性转换为 `float64`s 而不是 `string`s。
- 删除了未使用的属性，如`Style` 和`Text`

- Renamed `Svg` to `SVG` because *it's the right thing to do*.

- 将`Svg`重命名为`SVG`，因为*这是正确的做法*。

This will let us assert more precisely on the line we're looking for:

这将使我们更准确地断言我们正在寻找的行：

```
func TestSVGWriterAtMidnight(t *testing.T) {
    tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
    b := bytes.Buffer{}

    clockface.SVGWriter(&b, tm)

    svg := SVG{}

    xml.Unmarshal(b.Bytes(), &svg)

    want := Line{150, 150, 150, 60}

    for _, line := range svg.Line {
        if line == want {
            return
        }
    }

    t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}
```

Finally we can take a leaf out of the unit tests' tables, and we can write a helper function `containsLine(line Line, lines []Line) bool` to really make these tests shine:

最后，我们可以从单元测试的表中取出一片叶子，我们可以编写一个辅助函数 `containsLine(line Line,lines []Line) bool` 来真正让这些测试发光：

```
func TestSVGWriterSecondHand(t *testing.T) {
    cases := []struct {
        time time.Time
        line Line
    }{
        {
            simpleTime(0, 0, 0),
            Line{150, 150, 150, 60},
        },
        {
            simpleTime(0, 0, 30),
            Line{150, 150, 150, 240},
        },
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            b := bytes.Buffer{}
            clockface.SVGWriter(&b, c.time)

            svg := SVG{}
            xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Line) {
                t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
            }
        })
    }
}

func containsLine(l Line, ls []Line) bool {
    for _, line := range ls {
        if line == l {
            return true
        }
    }
    return false
}
```

Here's what [it looks like](https://github.com/quii/learn-go-with-tests/blob/main/math/v7c/clockface)

这是[看起来像](https://github.com/quii/learn-go-with-tests/blob/main/math/v7c/clockface)

Now *that's* what I call an acceptance test!

现在*这就是*我所说的验收测试！

### Write the test first

### 先写测试

So that's the second hand done. Now let's get started on the minute hand.

所以这是二手的。现在让我们开始分针。

```
func TestSVGWriterMinuteHand(t *testing.T) {
    cases := []struct {
        time time.Time
        line Line
    }{
        {
            simpleTime(0, 0, 0),
            Line{150, 150, 150, 70},
        },
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            b := bytes.Buffer{}
            clockface.SVGWriter(&b, c.time)

            svg := SVG{}
            xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Line) {
                t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_acceptance_test.go:87: Expected to find the minute hand line {X1:150 Y1:150 X2:150 Y2:70}, in the SVG lines [{X1:150 Y1:150 X2:150 Y2:60}]
```

We'd better start building some other clock hands, Much in the same way as we produced the tests for the second hand, we can iterate to produce the following set of tests. Again we'll comment out our acceptance test while we get this working:

我们最好开始构建一些其他时钟指针，与我们为秒针进行测试的方式非常相似，我们可以迭代以生成以下测试集。再次，当我们开始工作时，我们将注释掉我们的验收测试：

```
func TestMinutesInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(0, 30, 0), math.Pi},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := minutesInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:59:11: undefined: minutesInRadians
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func minutesInRadians(t time.Time) float64 {
    return math.Pi
}
```

### Repeat for new requirements

### 重复新要求

Well, OK - now let's make ourselves do some *real* work. We could model the minute hand as only moving every full minute - so that it 'jumps' from 30 to 31 minutes past without moving in between. But that would look a bit rubbish. What we want it to do is move a *tiny little bit* every second.

好吧，现在让我们自己做一些*真正的*工作。我们可以将分针建模为每整整一分钟才移动一次——这样它就会从 30 分钟到 31 分钟“跳跃”而不会在中间移动。但这看起来有点垃圾。我们想要它做的是每秒移动*一点点*。

```
func TestMinutesInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(0, 30, 0), math.Pi},
        {simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := minutesInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

How much is that tiny little bit? Well...

那一点点是多少？好...

- Sixty seconds in a minute
- thirty minutes in a half turn of the circle (`math.Pi` radians)
- so `30 * 60` seconds in a half turn.
- So if the time is 7 seconds past the hour ...
- ... we're expecting to see the minute hand at `7 * (math.Pi / (30 * 60))` radians past the 12.

- 一分钟六十秒
- 半圈三十分钟（`math.Pi`弧度）
- 所以`30 * 60` 秒半圈。
- 所以如果时间是一小时后的 7 秒......
- ...我们期望看到分针在 7 * (math.Pi / (30 * 60)) 弧度超过 12。

### Try to run the test

### 尝试运行测试

```
clockface_test.go:62: Wanted 0.012217304763960306 radians, but got 3.141592653589793
```

### Write enough code to make it pass 

### 编写足够的代码使其通过

In the immortal words of Jennifer Aniston: [Here comes the science bit](https://www.youtube.com/watch?v=29Im23SPNok)

用詹妮弗安妮斯顿不朽的话来说：[这里是科学位](https://www.youtube.com/watch?v=29Im23SPNok)

```
func minutesInRadians(t time.Time) float64 {
    return (secondsInRadians(t) / 60) +
        (math.Pi / (30 / float64(t.Minute())))
}
```

Rather than working out how far to push the minute hand around the clockface for every second from scratch, here we can just leverage the `secondsInRadians` function. For every second the minute hand will move 1/60th of the angle the second hand moves.

与其从头开始计算分针每秒钟在表盘上移动多远，我们可以利用“secondsInRadians”函数。每秒钟，分针将移动秒针移动角度的 1/60。

```
secondsInRadians(t) / 60
```

Then we just add on the movement for the minutes - similar to the movement of the second hand.

然后我们只需添加分钟的运动 - 类似于秒针的运动。

```
math.Pi / (30 / float64(t.Minute()))
```

And...

和...

```
PASS
ok      clockface    0.007s
```

Nice and easy. This is what things [look like now](https://github.com/quii/learn-go-with-tests/blob/main/math/v8/clockface/clockface_acceptance_test.go)

好，易于。这就是[现在的样子](https://github.com/quii/learn-go-with-tests/blob/main/math/v8/clockface/clockface_acceptance_test.go)

### Repeat for new requirements

### 重复新要求

Should I add more cases to the `minutesInRadians` test? At the moment there are only two. How many cases do I need before I move on to the testing the `minuteHandPoint` function?

我应该在“minutesInRadians”测试中添加更多案例吗？目前只有两个。在继续测试 `minuteHandPoint` 函数之前，我需要多少个案例？

One of my favourite TDD quotes, often attributed to Kent Beck,[^3] is

我最喜欢的 TDD 引语之一，通常归功于 Kent Beck，[^3] 是

> Write tests until fear is transformed into boredom.

> 编写测试，直到恐惧变成无聊为止。

And, frankly, I'm bored of testing that function. I'm confident I know how it works. So it's on to the next one.

而且，坦率地说，我对测试该功能感到厌烦。我相信我知道它是如何工作的。所以它在下一个。

### Write the test first

### 先写测试

```
func TestMinuteHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(0, 30, 0), Point{0, -1}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := minuteHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point) {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:79:11: undefined: minuteHandPoint
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func minuteHandPoint(t time.Time) Point {
    return Point{}
}
clockface_test.go:80: Wanted {0 -1} Point, but got {0 0}
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func minuteHandPoint(t time.Time) Point {
    return Point{0, -1}
}
PASS
ok      clockface    0.007s
```

### Repeat for new requirements

### 重复新要求

And now for some actual work

现在进行一些实际工作

```
func TestMinuteHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(0, 30, 0), Point{0, -1}},
        {simpleTime(0, 45, 0), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := minuteHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point) {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
clockface_test.go:81: Wanted {-1 0} Point, but got {0 -1}
```

### Write enough code to make it pass

### 编写足够的代码使其通过

A quick copy and paste of the `secondHandPoint` function with some minor changes ought to do it...

快速复制和粘贴 `secondHandPoint` 函数并进行一些小的更改应该可以做到...

```
func minuteHandPoint(t time.Time) Point {
    angle := minutesInRadians(t)
    x := math.Sin(angle)
    y := math.Cos(angle)

    return Point{x, y}
}
PASS
ok      clockface    0.009s
```

### Refactor

### 重构

We've definitely got a bit of repetition in the `minuteHandPoint` and `secondHandPoint` - I know because we just copied and pasted one to make the other. Let's DRY it out with a function.

我们肯定在 `minuteHandPoint` 和 `secondHandPoint` 中有一些重复 - 我知道因为我们只是复制并粘贴了一个来制作另一个。让我们用一个函数来干燥它。

```
func angleToPoint(angle float64) Point {
    x := math.Sin(angle)
    y := math.Cos(angle)

    return Point{x, y}
}
```

and we can rewrite `minuteHandPoint` and `secondHandPoint` as one liners:

我们可以将 `minuteHandPoint` 和 `secondHandPoint` 重写为一行：

```
func minuteHandPoint(t time.Time) Point {
    return angleToPoint(minutesInRadians(t))
}
func secondHandPoint(t time.Time) Point {
    return angleToPoint(secondsInRadians(t))
}
PASS
ok      clockface    0.007s
```

Now we can uncomment the acceptance test and get to work drawing the minute hand.

现在我们可以取消验收测试的注释并开始绘制分针。

### Write enough code to make it pass

### 编写足够的代码使其通过

The `minuteHand` function is a copy-and-paste of `secondHand` with some minor adjustments, such as declaring a `minuteHandLength`:

`minuteHand` 函数是 `secondHand` 的复制和粘贴，并进行了一些细微的调整，例如声明一个 `minuteHandLength`：

```
const minuteHandLength = 80

//...

func minuteHand(w io.Writer, t time.Time) {
    p := minuteHandPoint(t)
    p = Point{p.X * minuteHandLength, p.Y * minuteHandLength}
    p = Point{p.X, -p.Y}
    p = Point{p.X + clockCentreX, p.Y + clockCentreY}
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pX, pY)
}
```

And a call to it in our `SVGWriter` function:

并在我们的 `SVGWriter` 函数中调用它：

```
func SVGWriter(w io.Writer, t time.Time) {
    io.WriteString(w, svgStart)
    io.WriteString(w, bezel)
    secondHand(w, t)
    minuteHand(w, t)
    io.WriteString(w, svgEnd)
}
```

Now we should see that `TestSVGWriterMinuteHand` passes:

现在我们应该看到`TestSVGWriterMinuteHand`通过了：

```
PASS
ok      clockface    0.006s
```

But the proof of the pudding is in the eating - if we now compile and run our `clockface` program, we should see something like

但是布丁的证据在于吃——如果我们现在编译并运行我们的 `clockface` 程序，我们应该看到类似

[![a clock with second and minute hands](https://github.com/quii/learn-go-with-tests/raw/main/math/v9/clockface/clockface/clock.svg)](https://github.com/quii/learn-go-with-tests/blob/main/math/v9/clockface/clockface/clock.svg)

//github.com/quii/learn-go-with-tests/blob/main/math/v9/clockface/clockface/clock.svg)

### Refactor

### 重构

Let's remove the duplication from the `secondHand` and `minuteHand` functions, putting all of that scale, flip and translate logic all in one place.

让我们从 `secondHand` 和 `minuteHand` 函数中删除重复，将所有的缩放、翻转和翻译逻辑都放在一个地方。

```
func secondHand(w io.Writer, t time.Time) {
    p := makeHand(secondHandPoint(t), secondHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pX, pY)
}

func minuteHand(w io.Writer, t time.Time) {
    p := makeHand(minuteHandPoint(t), minuteHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pX, pY)
}

func makeHand(p Point, length float64) Point {
    p = Point{p.X * length, p.Y * length}
    p = Point{p.X, -p.Y}
    return Point{p.X + clockCentreX, p.Y + clockCentreY}
}
PASS
ok      clockface    0.007s
```

This is [where we're up to now](https://github.com/quii/learn-go-with-tests/tree/main/math/v9/clockface).

这是[我们现在的位置](https://github.com/quii/learn-go-with-tests/tree/main/math/v9/clockface)。

There... now it's just the hour hand to do!

那里......现在只是时针要做！

### Write the test first

### 先写测试

```
func TestSVGWriterHourHand(t *testing.T) {
    cases := []struct {
        time time.Time
        line Line
    }{
        {
            simpleTime(6, 0, 0),
            Line{150, 150, 150, 200},
        },
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            b := bytes.Buffer{}
            clockface.SVGWriter(&b, c.time)

            svg := SVG{}
            xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Line) {
                t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_acceptance_test.go:113: Expected to find the hour hand line {X1:150 Y1:150 X2:150 Y2:200}, in the SVG lines [{X1:150 Y1:150 X2:150 Y2:60} {X1:150 Y1:150 X2:150 Y2:70}]
```

Again, let's comment this one out until we've got the some coverage with the lower level tests:

再次，让我们注释掉这个，直到我们得到了一些较低级别测试的覆盖：

### Write the test first

### 先写测试

```
func TestHoursInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hoursInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:97:11: undefined: hoursInRadians
```

### Write the minimal amount of code for the test to run and check the failing test output

### 为测试编写最少的代码以运行并检查失败的测试输出

```
func hoursInRadians(t time.Time) float64 {
    return math.Pi
}
PASS
ok      clockface    0.007s
```

### Repeat for new requirements

### 重复新要求

```
func TestHoursInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
        {simpleTime(0, 0, 0), 0},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hoursInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_test.go:100: Wanted 0 radians, but got 3.141592653589793
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func hoursInRadians(t time.Time) float64 {
    return (math.Pi / (6 / float64(t.Hour())))
}
```

### Repeat for new requirements

### 重复新要求

```
func TestHoursInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(21, 0, 0), math.Pi * 1.5},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hoursInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_test.go:101: Wanted 4.71238898038469 radians, but got 10.995574287564276
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func hoursInRadians(t time.Time) float64 {
    return (math.Pi / (6 / (float64(t.Hour() % 12))))
}
```

Remember, this is not a 24-hour clock; we have to use the remainder operator to get the remainder of the current hour divided by 12.

请记住，这不是 24 小时制；我们必须使用余数运算符来得到当前小时的余数除以 12。

```
PASS
ok      github.com/gypsydave5/learn-go-with-tests/math/v10/clockface    0.008s
```

### Write the test first

### 先写测试

Now let's try to move the hour hand around the clockface based on the minutes and the seconds that have passed.

现在让我们尝试根据经过的分钟和秒数在表盘上移动时针。

```
func TestHoursInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(21, 0, 0), math.Pi * 1.5},
        {simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hoursInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_test.go:102: Wanted 0.013089969389957472 radians, but got 0
```

### Write enough code to make it pass

### 编写足够的代码使其通过

Again, a bit of thinking is now required. We need to move the hour hand along a little bit for both the minutes and the seconds. Luckily we have an angle already to hand for the minutes and the seconds - the one returned by `minutesInRadians`. We can reuse it!

同样，现在需要一些思考。我们需要为分钟和秒移动时针。幸运的是，我们已经有了一个角度来处理分钟和秒 - 由`minutesInRadians`返回的角度。我们可以重复使用它！

So the only question is by what factor to reduce the size of that angle. One full turn is one hour for the minute hand, but for the hour hand it's twelve hours. So we just divide the angle returned by `minutesInRadians` by twelve:

所以唯一的问题是通过什么因素来减小该角度的大小。分针转一整圈是一小时，而时针则是十二小时。所以我们只需将 `minutesInRadians` 返回的角度除以 12：

```
func hoursInRadians(t time.Time) float64 {
    return (minutesInRadians(t) / 12) +
        (math.Pi / (6 / float64(t.Hour()%12)))
}
```

and behold:

看看：

```
clockface_test.go:104: Wanted 0.013089969389957472 radians, but got 0.01308996938995747
```

Floating point arithmetic strikes again.

浮点运算再次来袭。

Let's update our test to use `roughlyEqualFloat64` for the comparison of the angles.

让我们更新我们的测试以使用 `roughlyEqualFloat64` 来比较角度。

```
func TestHoursInRadians(t *testing.T) {
    cases := []struct {
        time  time.Time
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(21, 0, 0), math.Pi * 1.5},
        {simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hoursInRadians(c.time)
            if !roughlyEqualFloat64(got, c.angle) {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}
PASS
ok      clockface    0.007s
```

### Refactor

### 重构

If we're going to use `roughlyEqualFloat64` in *one* of our radians tests, we should probably use it for *all* of them. That's a nice and simple refactor, which will leave things [looking like this](https://github.com/quii/learn-go-with-tests/tree/main/math/v10/clockface).

如果我们要在*一个*弧度测试中使用`roughlyEqualFloat64`，我们可能应该将它用于*所有*。这是一个很好且简单的重构，它将使事情[看起来像这样](https://github.com/quii/learn-go-with-tests/tree/main/math/v10/clockface)。

## Hour Hand Point

## 时针点

Right, it's time to calculate where the hour hand point is going to go by working out the unit vector.

是的，是时候通过计算单位向量来计算时针点的位置了。

### Write the test first

### 先写测试

```
func TestHourHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
        {simpleTime(6, 0, 0), Point{0, -1}},
        {simpleTime(21, 0, 0), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := hourHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point) {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}
```

Wait, am I going to write *two* test cases *at once*? Isn't this *bad TDD*?

等等，我要*一次*编写*两个*测试用例吗？这不是*糟糕的TDD*吗？

### On TDD Zealotry

### 关于 TDD 狂热

Test driven development is not a religion. Some people might act like it is - usually people who don't do TDD but are happy to moan on Twitter or Dev.to that it's only done by zealots and that they're 'being pragmatic' when they don't write tests. But it's not a religion. It's a tool.

测试驱动的开发不是一种信仰。有些人可能会表现得如此——通常是那些不做 TDD 但很高兴在 Twitter 或 Dev. 上抱怨它只是由狂热者完成的人，并且当他们不编写测试时他们是“务实的”。但它不是一种宗教。它是一个工具。

I *know* what the two tests are going to be - I've tested two other clock hands in exactly the same way - and I already know what my implementation is going to be - I wrote a function for the general case of changing an angle into a point in the minute hand iteration.

我*知道*这两个测试将是什么 - 我已经以完全相同的方式测试了另外两个时钟指针 - 我已经知道我的实现将是什么 - 我为更改一个的一般情况编写了一个函数角度到分针迭代中的一个点。

I'm not going to plough through TDD ceremony for the sake of it. TDD is a technique that helps me understand the code I'm writing - and the code that I'm going to write - better. TDD gives me feedback, knowledge and insight. But if I've already got that knowledge, then I'm not going to plough through the ceremony for no reason. Neither tests nor TDD are an end in themselves.

我不会为了它而努力完成 TDD 仪式。 TDD 是一种技术，可以帮助我更好地理解我正在编写的代码——以及我将要编写的代码。 TDD 给了我反馈、知识和洞察力。但如果我已经掌握了这些知识，那么我不会无缘无故地参加仪式。测试和 TDD 本身都不是目的。

My confidence has increased, so I feel I can make larger strides forward. I'm going to 'skip' a few steps, because I know where I am, I know where I'm going and I've been down this road before.

我的信心增加了，所以我觉得我可以向前迈出更大的步伐。我将“跳过”几步，因为我知道我在哪里，我知道我要去哪里，而且我以前也走过这条路。

But also note: I'm not skipping writing the tests entirely - I'm still writing them first. They're just appearing in less granular chunks.

但还要注意：我并没有完全跳过编写测试 - 我仍然首先编写它们。它们只是出现在粒度较小的块中。

### Try to run the test

### 尝试运行测试

```
./clockface_test.go:119:11: undefined: hourHandPoint
```

### Write enough code to make it pass

### 编写足够的代码使其通过

```
func hourHandPoint(t time.Time) Point {
    return angleToPoint(hoursInRadians(t))
}
```

As I said, I know where I am, and I know where I'm going. Why pretend otherwise? The tests will soon tell me if I'm wrong.

正如我所说，我知道我在哪里，我知道我要去哪里。为什么要假装不一样？如果我错了，测试很快就会告诉我。

```
PASS
ok      github.com/gypsydave5/learn-go-with-tests/math/v11/clockface    0.009s
```

## Draw the hour hand

## 绘制时针

And finally we get to draw in the hour hand. We can bring in that acceptance test by uncommenting it:

最后我们可以在时针上画画。我们可以通过取消注释来引入验收测试：

```
func TestSVGWriterHourHand(t *testing.T) {
    cases := []struct {
        time time.Time
        line Line
    }{
        {
            simpleTime(6, 0, 0),
            Line{150, 150, 150, 200},
        },
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            b := bytes.Buffer{}
            clockface.SVGWriter(&b, c.time)

            svg := SVG{}
            xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Line) {
                t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
            }
        })
    }
}
```

### Try to run the test

### 尝试运行测试

```
clockface_acceptance_test.go:113: Expected to find the hour hand line {X1:150 Y1:150 X2:150 Y2:200},
    in the SVG lines [{X1:150 Y1:150 X2:150 Y2:60} {X1:150 Y1:150 X2:150 Y2:70}]
```

### Write enough code to make it pass

### 编写足够的代码使其通过

And we can now make our final adjustments to the SVG writing constants and functions:

我们现在可以对 SVG 写入常量和函数进行最终调整：

```
const (
    secondHandLength = 90
    minuteHandLength = 80
    hourHandLength   = 50
    clockCentreX     = 150
    clockCentreY     = 150
)

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
    io.WriteString(w, svgStart)
    io.WriteString(w, bezel)
    secondHand(w, t)
    minuteHand(w, t)
    hourHand(w, t)
    io.WriteString(w, svgEnd)
}

// ...

func hourHand(w io.Writer, t time.Time) {
    p := makeHand(hourHandPoint(t), hourHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pX, pY)
}
```

And so...

所以...

```
ok      clockface    0.007s
```

Let's just check by compiling and running our `clockface` program.

让我们通过编译和运行我们的 `clockface` 程序来检查一下。

[![a clock](https://github.com/quii/learn-go-with-tests/raw/main/math/v12/clockface/clockface/clock.svg)](https://github.com/quii/learn-go-with-tests/blob/main/math/v12/clockface/clockface/clock.svg)

/quii/learn-go-with-tests/blob/main/math/v12/clockface/clockface/clock.svg)

### Refactor

### 重构

Looking at `clockface.go`, there are a few 'magic numbers' floating about. They are all based around how many hours/minutes/seconds there are in a half-turn around a clockface. Let's refactor so that we make explicit their meaning.

看看“clockface.go”，有一些“神奇的数字”在浮动。它们都基于围绕表盘半圈有多少小时/分钟/秒。让我们重构，以便明确它们的含义。

```
const (
    secondsInHalfClock = 30
    secondsInClock     = 2 * secondsInHalfClock
    minutesInHalfClock = 30
    minutesInClock     = 2 * minutesInHalfClock
    hoursInHalfClock   = 6
    hoursInClock       = 2 * hoursInHalfClock
)
```

Why do this? Well, it makes explicit what each number *means* in the equation. If - *when* - we come back to this code, these names will help us to understand what's going on.

为什么要这样做？好吧，它明确了方程中每个数字的*含义*。如果 - *when* - 我们回到这段代码，这些名称将帮助我们了解发生了什么。

Moreover, should we ever want to make some really, really WEIRD clocks - ones with 4 hours for the hour hand, and 20 seconds for the second hand say - these constants could easily become parameters. We're helping to leave that door open (even if we never go through it).

此外，如果我们想要制作一些非常非常奇怪的时钟——时针为 4 小时，秒针为 20 秒——这些常数很容易成为参数。我们正在帮助打开那扇门（即使我们从未穿过它）。

## Wrapping up

##  包起来

Do we need to do anything else?

我们还需要做些什么吗？

First, let's pat ourselves on the back - we've written a program that makes an SVG clockface. It works and it's great. It will only ever make one sort of clockface - but that's fine! Maybe you only *want* one sort of clockface. There's nothing wrong with a program that solves a specific problem and nothing else.

首先，让我们拍拍自己的后背——我们已经编写了一个制作 SVG 表盘的程序。它有效而且很棒。它只会制作一种表盘 - 但这很好！也许你只*想要*一种表盘。解决特定问题的程序没有错，没有其他问题。

### A Program... and a Library

### 一个程序...和一个库

But the code we've written *does* solve a more general set of problems to do with drawing a clockface. Because we used tests to think about each small part of the problem in isolation, and because we codified that isolation with functions, we've built a very reasonable little API for clockface calculations.

但是我们编写的代码*确实*解决了一系列与绘制表盘有关的更一般的问题。因为我们使用测试来孤立地考虑问题的每一小部分，并且因为我们用函数将这种隔离进行了编码，所以我们为表盘计算构建了一个非常合理的小 API。

We can work on this project and turn it into something more general - a library for calculating clockface angles and/or vectors.

我们可以在这个项目上工作，把它变成更通用的东西——一个用于计算表盘角度和/或矢量的库。

In fact, providing the library along with the program is *a really good idea*. It costs us nothing, while increasing the utility of our program and helping to document how it works.

事实上，与程序一起提供库是*一个非常好的主意*。它不会花费我们任何费用，同时增加我们程序的效用并帮助记录它是如何工作的。

> APIs should come with programs, and vice versa. An API that you must write C code to use, which cannot be invoked easily from the command line, is harder to learn and use. And contrariwise, it's a royal pain to have interfaces whose only open, documented form is a program, so you cannot invoke them easily from a C program. -- Henry Spencer, in *The Art of Unix Programming*

> API 应该随程序一起提供，反之亦然。必须编写 C 代码才能使用的 API，不能从命令行轻松调用，更难学习和使用。相反，如果接口的唯一公开的、文档化的形式是一个程序，那么你就不能从 C 程序中轻松调用它们，这是一种极大的痛苦。 ——亨利·斯宾塞，*Unix 编程艺术*

In [my final take on this program](https://github.com/quii/learn-go-with-tests/tree/main/math/vFinal/clockface), I've made the unexported functions within `clockface` into a public API for the library, with functions to calculate the angle and unit vector for each of the clock hands. I've also split the SVG generation part into its own package, `svg`, which is then used by the `clockface` program directly. Naturally I've documented each of the functions and packages.

在 [我对这个程序的最终看法](https://github.com/quii/learn-go-with-tests/tree/main/math/vFinal/clockface) 中，我在 `clockface` 中创建了未导出的函数进入库的公共 API，具有计算每个时钟指针的角度和单位向量的函数。我还将 SVG 生成部分拆分为自己的包 `svg`，然后由 `clockface` 程序直接使用。当然，我已经记录了每个功能和包。

Talking about SVGs...

谈论 SVG...

### The Most Valuable Test

### 最有价值的测试

I'm sure you've noticed that the most sophisticated piece of code for handling SVGs isn't in our application code at all; it's in the test code. Should this make us feel uncomfortable? Shouldn't we do something like

我相信您已经注意到，用于处理 SVG 的最复杂的一段代码根本不在我们的应用程序代码中；它在测试代码中。这应该让我们感到不舒服吗？我们不应该做类似的事情吗

- use a template from `text/template`?
- use an XML library (much as we're doing in our test)?
- use an SVG library? 

- 使用来自`text/template`的模板吗？
- 使用 XML 库（就像我们在测试中所做的那样）？
- 使用 SVG 库？

We could refactor our code to do any of these things, and we can do so because it doesn't matter *how* we produce our SVG, what is important is *what* we produce -  *an SVG*. As such, the part of our system that needs to know the most about SVGs - that needs to be the strictest about what constitutes an SVG - is the test for the SVG output: it needs to have enough context and knowledge about what an SVG is for us to be confident that we're outputting an SVG. The *what* of an SVG lives in our tests; the *how* in the code.

我们可以重构我们的代码来做任何这些事情，我们可以这样做，因为我们*如何*生产我们的SVG并不重要，重要的是我们*什么*我们生产 - *一个SVG*。因此，我们系统中最需要了解 SVG 的部分——需要对什么构成 SVG 最严格——是对 SVG 输出的测试：它需要有足够的上下文和关于什么是 SVG 的知识让我们确信我们正在输出 SVG。 SVG 的*什么*存在于我们的测试中；代码中的 *how*。

We may have felt odd that we were pouring a lot of time and effort into those SVG tests - importing an XML library, parsing XML, refactoring the structs - but that test code is a valuable part of our codebase - possibly more valuable than the current production code. It will help guarantee that the output is always a valid SVG, no matter what we choose to use to produce it.

我们可能会觉得奇怪，我们在这些 SVG 测试中投入了大量时间和精力——导入 XML 库、解析 XML、重构结构——但该测试代码是我们代码库的一个有价值的部分——可能比当前更有价值生产代码。这将有助于保证输出始终是有效的 SVG，无论我们选择使用什么来生成它。

Tests are not second class citizens - they are not 'throwaway' code. Good tests will last a lot longer than the version of the code they are testing. You should never feel like you're spending 'too much time' writing your tests. It is an investment. 

测试不是二等公民——它们不是“一次性”代码。好的测试将比他们正在测试的代码版本持续更长的时间。您永远不应该觉得自己在编写测试上花费了“太多时间”。这是一项投资。

