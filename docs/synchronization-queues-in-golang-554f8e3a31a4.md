# Synchronization queues in Golang

## how to use channels to write idiomatic code in Go

[Mar 12, 2018](https://medium.com/golangspec/synchronization-queues-in-golang-554f8e3a31a4?source=post_page-----554f8e3a31a4--------------------------------) · 6 min read


# Problem

Let’s suppose we’re running an IT company employing programmers and testers.  To give people a chance to get to know each other and relax a bit we’ve  bought a ping-pong table and established following rules:

- exactly two people can play at the same time,
- next pair can start their game only when previous one finished so it’s not allowed to switch one player,
- testers can only work with programmers and vice versa (never two testers or two programmers together). If programmer or tester wants play then needs to wait for tester or programmer respectively to establish a valid pair.

```
func main() {
    for i := 0; i < 10; i++ {
        go programmer()
    }
    for i := 0; i < 5; i++ {
        go tester()
    }
    select {} // long day at work...
}func programmer() {
    for {
        code()
        pingPong()
    }
}func tester() {
    for {
        test()
        pingPong()
    }
}
```

We’ll mimic testing, coding and playing ping-pong by `time.sleep`:

```
func test() {
    work()
}func code() {
    work()
}func work() {
    // Sleep up to 10 seconds.
    time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
}func pingPong() {
    // Sleep up to 2 seconds.
    time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}func programmer() {
    for {
        code()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")    }
}func tester() {
    for {
        test()
        fmt.Println("Tester starts")
        pingPong()
        fmt.Println("Tester ends")
    }
}
```

Such program emits stream of messages like this:

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

Below are two solutions. 1st one is based on mutexes and 2nd one uses  separate worker which coordinates the whole process making sure  everything happens according to the policy.

# Solution #1

Both solutions are using data structure which queues programmers and testers before approaching the table. When there is at least one valid pair  available (Dev + QA) then one pair is allowed to play ping-pong.

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
}func programmer(q *queue.Queue) {
    for {
        code()
        q.StartP()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")
        q.EndP()
    }
}func main() {
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

```
package queueimport "sync"type Queue struct {
    mut                   sync.Mutex
    numP, numT            int
    queueP, queueT, doneP chan int
}func New() *Queue {
    q := Queue{
        queueP: make(chan int),
        queueT: make(chan int),
        doneP:  make(chan int),
    }
    return &q
}func (q *Queue) StartT() {
    q.mut.Lock()
    if q.numP > 0 {
        q.numP -= 1
        q.queueP <- 1
    } else {
        q.numT += 1
        q.mut.Unlock()
        <-q.queueT
    }
}func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}func (q *Queue) StartP() {
    q.mut.Lock()
    if q.numT > 0 {
        q.numT -= 1
        q.queueT <- 1
    } else {
        q.numP += 1
        q.mut.Unlock()
        <-q.queueP
    }
}func (q *Queue) EndP() {
    q.doneP <- 1
}
```

Queue has mutex which serves two purposes:

- synchronizes access to shared counters (`numT` and `numP`)
- acts as a token held by playing employees which blocks others from joing the table

Programmers and testers are waiting for ping-pong parter using unbuffered channels:

```
<-q.queueP
```

or

```
<-q.queueT
```

Reading from one of these channels will block goroutine if there is no opponent available.

Let’s analyze `StartT` which is executed by testers:

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

If `numP` is greater than 0 (there is at least one programmer waiting for the  game) then number of waiting programmers is decreased by one and one of  waiting programmer will be allowed to join the table (`q.queueP <- 1`). What is interesting here is during this path mutex won’t be released so it’ll serve as a token giving exclusive access to ping-pong table.

If there is no waiting programmers then `numT` (number of waiting testers) is increased and goroutine blocks on `<-q.queueT`.

`StartP` is basically the same but executed by programmers.

During the play, mutex will be locked so it needs to be released by either  programmer or tester. To release mutex only when both parties finished  the game a barrier `doneP` is used:

```
func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}func (q *Queue) EndP() {
    q.doneP <- 1
}
```

If programmer is still playing and tester finished then tester will block on `<-q.doneP`. As soon as programmer reaches `q.doneP <- 1` then barrier will open and mutex will be released to allow these employees to go back to work.

If tester is still playing then programmer will block on `q.doneP <- 1`. When tester is done then it reads from barrier `<-q.doneP` which will unblock programmer and mutex will be released to free the table.

What is interesting here is that mutex is always released by tester even  when either tester or programmer might locked it. It’s also part of  reason why this solution might not be so obvious at first glance…

# Solution #2

```
package queueconst (
    msgPStart = iota
    msgTStart
    msgPEnd
    msgTEnd
)type Queue struct {
    waitP, waitT   int
    playP, playT   bool
    queueP, queueT chan int
    msg            chan int
}func New() *Queue {
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
}func (q *Queue) StartT() {
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

- number of waiting Devs or QAs is increased
- information about playing employees is updated

After receiving any of defined messages, scheduler will check if it’s allowed to start a game by another pair:

```
if q.waitP > 0 && q.waitT > 0 && !q.playP && !q.playT {
```

If so the state is updated accordingly and one tester and one programmer are unblocked.

Instead of using mutexes (as in solution#1) to manage access to shared data,  we’ve now a separate goroutine which talks with outside world over the  channel. This gives us more idiomatic Go program.

> *Don’t communicate by sharing memory, share memory by communicating.*

# Resources

- “The Little Book of Semaphores” by Allen B. Downey
- [https://medium.com/golangspec/reusable-barriers-in-golang-156db1 f75d0b](https://medium.com/golangspec/reusable-barriers-in-golang-156db1f75d0b)
- https://blog.golang.org/share-memory-by-communicating

If you’ve found alternative solution you would like to share then please comment beneath.

👏👏👏 below to help others discover this story. Please follow me here or on [Twitter](https://twitter.com/mlowicki) if you want to get updates about new posts or boost work on future stories.
