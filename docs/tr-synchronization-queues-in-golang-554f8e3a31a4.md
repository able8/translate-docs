# Synchronization queues in Golang

# Golang 中的同步队列

## how to use channels to write idiomatic code in Go

## 如何使用通道在 Go 中编写惯用代码

[Mar 12, 2018](https://medium.com/golangspec/synchronization-queues-in-golang-554f8e3a31a4?source=post_page-----554f8e3a31a4--------------------------------) · 6 min read From: https://medium.com/golangspec/synchronization-queues-in-golang-554f8e3a31a4


# Problem

#  问题

Let’s suppose we’re running an IT company employing programmers and testers. To give people a chance to get to know each other and relax a bit we’ve  bought a ping-pong table and established following rules:

假设我们经营着一家 IT 公司，雇佣程序员和测试员。为了让人们有机会相互了解并放松一下，我们买了一张乒乓球桌并制定了以下规则：

- exactly two people can play at the same time,
- next pair can start their game only when previous one finished so it’s not allowed to switch one player,
- testers can only work with programmers and vice versa (never two testers or two programmers together). If programmer or tester wants play then needs to wait for tester or programmer respectively to establish a valid pair.

- 正好两个人可以同时玩，
- 下一组只能在前一组完成后才能开始他们的比赛，因此不允许更换一名球员，
- 测试人员只能与程序员一起工作，反之亦然（永远不要两个测试人员或两个程序员一起工作）。如果程序员或测试员要玩，则需要分别等待测试员或程序员建立有效的配对。

```
func main() {
    for i := 0;i < 10;i++ {
        go programmer()
    }
    for i := 0;i < 5;i++ {
        go tester()
    }
    select {} // long day at work...
}
func programmer() {
    for {
        code()
        pingPong()
    }
}
func tester() {
    for {
        test()
        pingPong()
    }
}
```


We’ll mimic testing, coding and playing ping-pong by `time.sleep`:

我们将通过 `time.sleep` 模拟测试、编码和打乒乓球：

```
func test() {
    work()
}
func code() {
    work()
}
func work() {
    // Sleep up to 10 seconds.
    time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
}
func pingPong() {
    // Sleep up to 2 seconds.
    time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}
func programmer() {
    for {
        code()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")    }
}
func tester() {
    for {
        test()
        fmt.Println("Tester starts")
        pingPong()
        fmt.Println("Tester ends")
    }
}
```


Such program emits stream of messages like this:

这样的程序会发出这样的消息流：

```
> go run pingpong.go
Tester starts
Programmer starts
Programmer starts
Tester ends
Programmer ends
Programmer starts
Programmer ends
Programmer ends
```


But to play ping-pong in accordance with the rules our stream of messages  can consists only of such 4-lines long sequences (in any order and  repeated arbitrary number of times):

但是要按照规则打乒乓球，我们的消息流只能包含这样的 4 行长序列（以任意顺序重复任意次数）：

```
Tester starts
Programmer starts
Tester ends
Programmer endsTester starts
Programmer starts
Programmer ends
Tester endsProgrammer starts
Tester starts
Tester ends
Programmer endsProgrammer starts
Tester starts
Programmer ends
Tester ends
```


So first either tester or programmer approaches the table. Afterwards  partner joins (programmer or tester accordingly). While leaving the game they can do it in any order. This is why we’ve 4 valid sequences.

所以首先测试人员或程序员接近表格。之后合作伙伴加入（相应的程序员或测试员）。离开游戏时，他们可以按任何顺序进行。这就是为什么我们有 4 个有效序列。

Below are two solutions. 1st one is based on mutexes and 2nd one uses  separate worker which coordinates the whole process making sure  everything happens according to the policy.

下面是两种解决方案。第一个基于互斥锁，第二个使用单独的 worker，它协调整个过程，确保一切都按照策略进行。

# Solution #1

# 解决方案 #1

Both solutions are using data structure which queues programmers and testers before approaching the table. When there is at least one valid pair  available (Dev + QA) then one pair is allowed to play ping-pong.

这两种解决方案都使用数据结构，在接近表之前将程序员和测试人员排队。当至少有一对可用的有效配对 (Dev + QA) 时，则允许一对打乒乓球。

```
func tester(q *queue.Queue) {
    for {
        test()
        q.StartT()
        fmt.Println("Tester starts")
        pingPong()
        fmt.Println("Tester ends")
        q.EndT()
    }
}
func programmer(q *queue.Queue) {
    for {
        code()
        q.StartP()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")
        q.EndP()
    }
}
func main() {
    q := queue.New()
    for i := 0; i < 10; i++ {
        go programmer(q)
    }
    for i := 0; i < 5; i++ {
        go tester(q)
    }
    select {}
}
```


Package queue is defined as follows:

包队列定义如下：

```
package queue
import "sync"
type Queue struct {
    mut                   sync.Mutex
    numP, numT            int
    queueP, queueT, doneP chan int
}
func New() *Queue {
    q := Queue{
        queueP: make(chan int),
        queueT: make(chan int),
        doneP:  make(chan int),
    }
    return &q
}
func (q *Queue) StartT() {
    q.mut.Lock()
    if q.numP > 0 {
        q.numP -= 1
        q.queueP <- 1
    } else {
        q.numT += 1
        q.mut.Unlock()
        <-q.queueT
    }
}
func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}
func (q *Queue) StartP() {
    q.mut.Lock()
    if q.numT > 0 {
        q.numT -= 1
        q.queueT <- 1
    } else {
        q.numP += 1
        q.mut.Unlock()
        <-q.queueP
    }
}
func (q *Queue) EndP() {
    q.doneP <- 1
}

```


Queue has mutex which serves two purposes:

队列有互斥锁，它有两个目的：

- synchronizes access to shared counters (`numT` and `numP`) 

- 同步访问共享计数器（`numT` 和 `numP`）

- acts as a token held by playing employees which blocks others from joing the table

- 充当玩家持有的令牌，阻止其他人加入桌子

Programmers and testers are waiting for ping-pong parter using unbuffered channels:

程序员和测试人员正在使用无缓冲通道等待乒乓合作：

```
<-q.queueP
```


or

```
<-q.queueT
```


Reading from one of these channels will block goroutine if there is no opponent available.

如果没有可用的对手，从这些通道之一读取将阻塞 goroutine。

Let’s analyze `StartT` which is executed by testers:

让我们分析由测试人员执行的“StartT”：

```
func (q *Queue) StartT() {
    q.mut.Lock()
    if q.numP > 0 {
        q.numP -= 1
        q.queueP <- 1
    } else {
        q.numT += 1
        q.mut.Unlock()
        <-q.queueT
    }
}
```


If `numP` is greater than 0 (there is at least one programmer waiting for the  game) then number of waiting programmers is decreased by one and one of  waiting programmer will be allowed to join the table (`q.queueP <- 1` ). What is interesting here is during this path mutex won’t be released so it’ll serve as a token giving exclusive access to ping-pong table.

如果 `numP` 大于 0（至少有一个程序员在等待游戏），那么等待的程序员数量减一，等待的程序员将被允许加入表中（`q.queueP <- 1` ）。有趣的是，在此路径中互斥锁不会被释放，因此它将作为一个令牌，提供对乒乓表的独占访问权限。

If there is no waiting programmers then `numT` (number of waiting testers) is increased and goroutine blocks on `<-q.queueT`.

如果没有等待的程序员，则增加 `numT`（等待测试人员的数量）并且 goroutine 在 `<-q.queueT` 上阻塞。

`StartP` is basically the same but executed by programmers.

`StartP` 基本相同，但由程序员执行。

During the play, mutex will be locked so it needs to be released by either  programmer or tester. To release mutex only when both parties finished  the game a barrier `doneP` is used:

在播放过程中，互斥量会被锁定，因此需要由程序员或测试人员释放。要仅在双方完成游戏时释放互斥锁，使用屏障 `doneP`：

```
func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}func (q *Queue) EndP() {
    q.doneP <- 1
}
```


If programmer is still playing and tester finished then tester will block on `<-q.doneP`. As soon as programmer reaches `q.doneP <- 1` then barrier will open and mutex will be released to allow these employees to go back to work.

如果程序员仍在玩并且测试人员已完成，则测试人员将阻止`<-q.doneP`。一旦程序员达到`q.doneP <- 1`，屏障就会打开，互斥量将被释放以允许这些员工回去工作。

If tester is still playing then programmer will block on `q.doneP <- 1`. When tester is done then it reads from barrier `<-q.doneP` which will unblock programmer and mutex will be released to free the table.

如果测试器仍在播放，那么程序员将在 `q.doneP <- 1` 上阻塞。当测试完成后，它从屏障 `<-q.doneP` 中读取，这将解除对程序员的阻塞，并且互斥量将被释放以释放表。

What is interesting here is that mutex is always released by tester even  when either tester or programmer might locked it. It’s also part of  reason why this solution might not be so obvious at first glance…

这里有趣的是互斥锁总是由测试人员释放，即使测试人员或程序员可能锁定它。这也是为什么这个解决方案乍一看可能不那么明显的部分原因......

# Solution #2

# 解决方案#2

```
package queue
const (
    msgPStart = iota
    msgTStart
    msgPEnd
    msgTEnd
)
type Queue struct {
    waitP, waitT   int
    playP, playT   bool
    queueP, queueT chan int
    msg            chan int
}
func New() *Queue {
    q := Queue{
        msg:    make(chan int),
        queueP: make(chan int),
        queueT: make(chan int),
    }
    go func() {
        for {
            select {
            case n := <-q.msg:
                switch n {
                case msgPStart:
                    q.waitP++
                case msgPEnd:
                    q.playP = false
                case msgTStart:
                    q.waitT++
                case msgTEnd:
                    q.playT = false
                }
                if q.waitP > 0 && q.waitT > 0 && !q.playP && !q.playT {
                    q.playP = true
                    q.playT = true
                    q.waitT--
                    q.waitP--
                    q.queueP <- 1
                    q.queueT <- 1
                }
            }
        }
    }()
    return &q
}
func (q *Queue) StartT() {
    q.msg <- msgTStart
    <-q.queueT
}func (q *Queue) EndT() {
    q.msg <- msgTEnd
}func (q *Queue) StartP() {
    q.msg <- msgPStart
    <-q.queueP
}func (q *Queue) EndP() {
    q.msg <- msgPEnd
}
```


We’ve a central coordinator running inside separate goroutine which  orchestrate the whole process. Scheduler gets information about new  employee who wants to relax or if someone is done with ping-pong via `msg` channel. While receiving any message state of the scheduler is updated:

我们有一个中央协调器在单独的 goroutine 中运行，它协调整个过程。调度程序通过 `msg` 通道获取有关想要放松的新员工的信息，或者是否有人完成了乒乓球。在接收任何消息时，调度程序的状态被更新：

- number of waiting Devs or QAs is increased
- information about playing employees is updated

- 等待 Devs 或 QAs 的数量增加
- 更新了关于扮演员工的信息

After receiving any of defined messages, scheduler will check if it’s allowed to start a game by another pair:

收到任何定义的消息后，调度程序将检查是否允许另一对开始游戏：

```
if q.waitP > 0 && q.waitT > 0 && !q.playP && !q.playT {
```


If so the state is updated accordingly and one tester and one programmer are unblocked.

如果是这样，状态会相应地更新，并且一名测试员和一名程序员被解除阻塞。

Instead of using mutexes (as in solution#1) to manage access to shared data,  we’ve now a separate goroutine which talks with outside world over the  channel. This gives us more idiomatic Go program.

我们现在没有使用互斥体（如解决方案#1）来管理对共享数据的访问，我们现在有一个单独的 goroutine 通过通道与外部世界对话。这给了我们更多惯用的 Go 程序。

> *Don’t communicate by sharing memory, share memory by communicating.*

> *不要通过共享内存来通信，通过通信来共享内存。*

# Resources

#  资源

- “The Little Book of Semaphores” by Allen B. Downey 

- Allen B. Downey 的“信号量小书”

- [https://medium.com/golangspec/reusable-barriers-in-golang-156db1 f75d0b](https://medium.com/golangspec/reusable-barriers-in-golang-156db1f75d0b)
- https://blog.golang.org/share-memory-by-communicating

- [https://medium.com/golangspec/reusable-barriers-in-golang-156db1 f75d0b](https://medium.com/golangspec/reusable-barriers-in-golang-156db1f75d0b)
- https://blog.golang.org/share-memory-by-communicating

If you’ve found alternative solution you would like to share then please comment beneath.

如果您找到了想要分享的替代解决方案，请在下方发表评论。

👏👏👏 below to help others discover this story. Please follow me here or on [Twitter](https://twitter.com/mlowicki) if you want to get updates about new posts or boost work on future stories. 

👏👏👏 在下面帮助其他人发现这个故事。如果您想获取有关新帖子的更新或推动未来故事的工作，请在此处或在 [Twitter](https://twitter.com/mlowicki) 上关注我。

