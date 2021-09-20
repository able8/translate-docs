# Bad Go: Adventures with JSON marshalling

Adventures for the indoors

Sat, Oct 5, 2019

This is a story about encoding/json in the Go standard library. I’m not going to say this is bad Go. That would be heresy. But there is an  aspect of marshalling that could be improved. Because it is in the  standard library it isn’t bad Go, but if you followed the pattern in  your own code then that would be a mistake. Outside of the standard  library it would lose its magical aura, and it would be bad Go.

My frustration is with the Marshaler interface and the MarshalJSON  method. This method makes it pretty much impossible for custom JSON  marshalling to be efficient. The inimitable Mr. Cheney has recently  warned us about this very issue [here](https://dave.cheney.net/2019/09/05/dont-force-allocations-on-the-callers-of-your-api).

(To be clear, although I did sit next to Mr Cheney at a meetup, and  once he did like one of my tweets, that does not mean he in any way  endorses this blog or its content)

Let’s try to demonstrate the problem.  We’ll start by marshalling a very simple struct in a simple benchmark.

```go
type mystruct struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ReportAllocs()
	var data = mystruct{
		A: 42,
		B: "42",
	}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&data)
		if err != nil {
			b.Errorf("failed to marshal json. %s", err)
		}
	}
}
BenchmarkJSONMarshal-8   3627376    316 ns/op   32 B/op   1 allocs/op
```

If we run this we find there’s just 1 allocation per  marshalling attempt, which is the byte slice containing the marshalled  data. It would be nice if we could re-use a slice for this, but one  allocation is not too upsetting. And if we really want to we can use an [encoder](https://golang.org/pkg/encoding/json/#Encoder) to avoid this.

So what am I complaining about? Well, let’s modify our struct a little to add a time.

```go
type mystruct struct {
	A int       `json:"a,omitempty"`
	B string    `json:"b,omitempty"`
	C time.Time `json:"c"`
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ReportAllocs()
	var data = mystruct{
		A: 42,
		B: "42",
		C: time.Now(),
	}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&data)
		if err != nil {
			b.Errorf("failed to marshal json. %s", err)
		}
	}
}
BenchmarkJSONMarshal-8    981222   1345 ns/op  208 B/op   4 allocs/op
```

Suddenly we’re making 4 allocations per marshalling  attempt! 3 additional allocations because we’ve added a time. Why would  that be? Well, one issue is that the json package does not natively  understand time.Time, and marshals it via the Marshaler interface.  time.Time implements [MarshalJSON](https://golang.org/pkg/time/#Time.MarshalJSON). This forces an additional allocation because the method is defined to return a `[]byte` with the marshalled time. There’s no mechanism in the API to allow this custom marshaler to append it’s data to the data marshalled so far. It  needs to allocate a separate slice that it returns (thus forcing a heap  allocation), and which the json library then appends to its output.

That explains 1 additional allocation. Why are there 3? Well, we can benchmark Time.MarshalJSON to see what it is doing.

```go
func BenchmarkTimeMarshal(b *testing.B) {
	b.ReportAllocs()
	var t time.Time

	for i := 0; i < b.N; i++ {
		_, err := t.MarshalJSON()
		if err != nil {
			b.Errorf("failed to marshal. %s", err)
		}
	}
}
BenchmarkTimeMarshal-8   3400222    378 ns/op   48 B/op   1 allocs/op
```

This only creates 1 allocation. So the other 2 must  somehow come about within the json package itself, presumably as  additional overhead joining up the results.

If we run the benchmark under the profiler we discover the causes of the 4 allocations.

1. The byte slice that holds the final marshalled JSON.
2. The byte slice Time.MarshalJSON is forced to generate.
3. Some additional overhead copying the marshalled JSON from  Time.MarshalJSON into the result byte slice. This uses json.Compact,  which allocates a scanner while it does the copying because it also  checks the JSON is valid and ensures insignificant space is removed from the JSON.
4. To access the Marshaler interface, json uses the reflect package, and in fact creates a new `interface{}` value pointing to the time value. This somehow causes an allocation.

As far as I can tell all 3 of these allocations are currently unavoidable if you use a custom JSON marshaler for a type.

Why do I find this so frustrating? To me the existence of the  json.Marshaler interface looks like an escape hatch: a mechanism to do  things that are out of the ordinary; to put effort in and improve  performance. But it isn’t that. It’s a garbage chute - use it and you’ll end up stuck in a bin covered in garbage.

- Have lots of timestamps in your data => covered in garbage
- Want to use json.RawMessage to avoid encoding parts of your data => covered in garbage
- Need to express null fields, but want to avoid using pointers so you don’t get covered in garbage? Well, you’ll do a lot of work and end up  covered in garbage.

Now, none of this is a problem if you’re not marshalling a lot of  JSON. But if you are it starts to make Go look like a poor choice. Or  you have to look at third-party JSON encoders, which isn’t an  unreasonable option but is somehow unsatisfying.

How could we improve on this? What if we added a second marshaler interface?

```go
type MarshalAppender interface {
    MarshalAppendJSON(buf []byte) ([]byte, error)
}
```

Implementers of this interface append their json directly to the `buf` parameter passed in. We define things so that MarshalAppendJSON must  append valid JSON without any redundant white space. Finally we work out why accessing the interface method causes an allocation and fix it.  Then we’ll have the possibility of allocation-free custom JSON  marshalling.

## Is it Bad Go?

MarshalAppender is perhaps a little more complicated than Marshaler.  And simple is often best. But if your code is a fundamental building  block, either within your own project or for projects throughout the  world, I’d argue it’s worth going the extra mile to provide both  efficient implementations and APIs that can be used efficiently.

Providing just the simple interface may seem simpler and clearer. But what happens when someone needs that greater efficiency? Either they’re stuck, or they create a whole new implementation, or they go to extreme lengths to deal with the garbage collector. You’ve not reduced the  complexity in the world - you’ve deferred it. And increased it.

## Next steps

I’m actually going to [propose this](https://github.com/golang/go/issues/34701) to the Go team and try to contribute the change. I intend to write  about the experience in a future blog. Hopefully it won’t be terribly  interesting!
