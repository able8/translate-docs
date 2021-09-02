# The underutilized usefulness of sync.Once

# sync.Once 未被充分利用的用处

2021-08-14 From: https://blog.chuie.io/posts/synconce/

If you've ever used goroutines in Go, you've probably come across a couple of concurrency primitives. Probably `sync.Mutex`, `sync.WaitGroup` and maybe `sync.Map`, but have you heard of `sync.Once`?

如果您曾经在 Go 中使用过 goroutine，那么您可能会遇到几个并发原语。可能是`sync.Mutex`、`sync.WaitGroup` 和`sync.Map`，但你听说过`sync.Once` 吗？

Maybe you have, but what does the [godoc say about it](https://golang.org/pkg/sync/#Once)?

也许你有，但 [godoc 对此有何评论](https://golang.org/pkg/sync/#Once)？

> Once is an object that will perform exactly one action.

> 一次是一个只执行一个动作的对象。

Sounds simple enough, what's so useful about it then?

听起来很简单，那么它有什么用呢？

Well for some reason this isn't particularly well documented, but a `sync.Once` will *wait* until the execution inside the first `.Do` completes. This makes it incredibly useful when performing relatively expensive operations that you would typically cache in a map.

好吧，出于某种原因，这并不是特别好的文档，但是 `sync.Once` 将*等待*直到第一个 `.Do` 中的执行完成。这在执行通常会缓存在地图中的相对昂贵的操作时非常有用。

## Naive caching

## 原始缓存

Say for example you have a popular website that hits a backend API that isn't particularly fast, so you decide to cache API results in-memory with a map. A naive solution might look like this:

```go
package main

type QueryClient struct {
    cache map[string][]byte
    mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) []byte {
    // Check if the result is already cached.
    c.mutex.Lock()
    if cached, found := c.cache[name];found {
        c.mutex.Unlock()
        return cached, nil
    }
    c.mutex.Unlock()

    // Make the request if it's uncached.
    resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
    // Error handling and resp.Body.Close omitted for brevity.
    result, err := ioutil.ReadAll(resp)

    // Store the result in the cache.
    c.mutex.Lock()
    c.cache[name] = result
    c.mutex.Unlock()

    return result
}
```


Looks good, right?

看起来不错，对吧？

Well what happens if there are two calls to `DoQuery` that happen simultaneously? The calls would race, neither would see the cache is populated, and both would perform the HTTP request to `upstream.api` unnecessarily, when only one would need to complete it.

如果同时发生两次对 DoQuery 的调用会发生什么？调用会竞争，两者都不会看到缓存被填充，并且当只有一个人需要完成它时，两者都会不必要地执行对 `upstream.api` 的 HTTP 请求。

## Ugly but better caching

## 丑陋但更好的缓存

I don't have statistics on this, but one way I would imagine people solving this is by using channels, contexts or mutexes. For example you could turn this into:

```go
package main

type CacheEntry struct {
    data []byte
    wait <-chan struct{}
}

type QueryClient struct {
    cache map[string]*CacheEntry
    mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) []byte {
    // Check if the operation has already been started.
    c.mutex.Lock()
    if cached, found := c.cache[name];found {
        c.mutex.Unlock()
        // Wait for it to complete.
        <-cached.wait
        return cached.data, nil
    }

    entry := &CacheEntry{
        data: result,
        wait: make(chan struct{}),
    }
    c.cache[name] = entry
    c.mutex.Unlock()

    // Make the request if it's uncached.
    resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
    // Error handling and resp.Body.Close omitted for brevity
    entry.data, err = ioutil.ReadAll(resp)

    // Signal that the operation is complete, receiving on closed channels
    // returns immediately.
    close(entry.wait)

    return entry.data
}
```


That's good and all but the code's readability has taken a hit. It's not immediately clear what's going on with `cached.wait` and the flow of operations under different situations is not very intuitive.

这很好，但代码的可读性受到了打击。目前还不清楚`cached.wait` 发生了什么，不同情况下的操作流程也不是很直观。

## Applying `sync.Once`

## 应用`sync.Once`

Let's try to apply `sync.Once` to this instead:

```go
package main

type CacheEntry struct {
    data []byte
    once *sync.Once
}

type QueryClient struct {
    cache map[string]*CacheEntry
    mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) []byte {
    c.mutex.Lock()
    entry, found := c.cache[name]
    if !found {
        // Create a new entry if one does not exist already.
        entry = &CacheEntry{
            once: new(sync.Once),
        }
        c.cache[name] = entry
    }
    c.mutex.Unlock()

    // Now when we invoke `.Do`, if there is an on-going simultaneous operation,
    // it will block until it has completed (and `entry.data` is populated).
    // Or if the operation has already completed once before,
    // this call is a no-op and doesn't block.
    entry.once.Do(func() {
        resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
        // Error handling and resp.Body.Close omitted for brevity
        entry.data, err = ioutil.ReadAll(resp)
    })

    return entry.data
}
```


That's it. This achieves the same as the previous example, but is now much easier to understand (at least in my opinion). There is only a single return, and the code flows intuitively from top to bottom without having to read and understand what's going on with the `entry.wait` channel as before.

就是这样。这与前面的示例实现相同，但现在更容易理解（至少在我看来）。只有一个返回，代码从上到下直观地流动，无需像以前一样阅读和理解 `entry.wait` 通道发生了什么。

## Further reading/additional considerations

## 进一步阅读/其他注意事项

Another mechanism similar to `sync.Once` is [golang.org/x/sync/singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight). However `singleflight` only deduplicates requests that are in-flight (i.e. doesn't cache persistently). `singleflight` however may be cleaner to implement with contexts compared to `sync.Once` (through the use of a `select` and `ctx.Done()`), in production environments this may be important as to be able to cancel out with a context. The pattern with `singleflight` is quite similar to `sync.Once` but you would early return if a value is present inside the map.

另一种类似于`sync.Once`的机制是 [golang.org/x/sync/singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)。然而，`singleflight` 只对进行中的请求进行重复数据删除（即不会持久缓存）。然而，与“sync.Once”（通过使用“select”和“ctx.Done()”）相比，“singleflight”可能更清晰地实现上下文，在生产环境中这可能很重要，因为能够取消有上下文。 `singleflight` 的模式与 `sync.Once` 非常相似，但如果地图中存在值，你会提前返回。

[ianlancetaylor](https://github.com/golang/go/issues/25312#issuecomment-387800105) suggested the following pattern to use `sync.Once` with contexts:

```go
c := make(chan bool, 1)
go func() {
    once.Do(f)
    c <- true
}()
select {
case <-c:
case <-ctxt.Done():
    return
}
```
