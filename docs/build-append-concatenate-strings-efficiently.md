# Efficient string concatenation [full guide]

yourbasic.org/golang

- [Clean and simple string building (fmt)](http://yourbasic.org#clean-and-simple-string-building)
- [High-performance string concatenation (stringbuilder)](http://yourbasic.org#high-performance-string-concatenation)
- [Before Go 1.10 (bytebuffer)](http://yourbasic.org#before-go-1-10)

## Clean and simple string building

For simple cases where performance is a non-issue,
[`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend.
It’s clean, simple and fairly efficient.

```
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](http://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/)
lists the most common formatting verbs and flags.

## High-performance string concatenation[Go 1.10](https://golang.org/doc/go1.10 "Go 1.10 Release Notes")

A [`strings.Builder`](https://golang.org/pkg/strings/#Builder)
is used to efficiently append strings using write methods.

- It offers a subset of the[`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) methods
  that allows it to safely avoid extra copying when converting a builder to a string.
- You can use the[`fmt`](https://golang.org/pkg/fmt/) package for formatting
  since the builder implements the [`io.Writer`](http://yourbasic.org/golang/io-writer-interface-explained/) interface.
- The[`Grow`](https://golang.org/pkg/strings/#Builder.Grow) method
  can be used to preallocate memory when the maximum size of the string is known.

```
var b strings.Builder
b.Grow(32)
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&b, "%d:%d, ", i+1, p)
}
s := b.String()   // no copying
s = s[:b.Len()-2] // no copying (removes trailing ", ")
fmt.Println(s)
```

```
1:2, 2:3, 3:5, 4:7, 5:11, 6:13

```

## Before Go 1.10

Use [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf)
to print into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer).

```
var buf bytes.Buffer
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&buf, "%d:%d, ", i+1, p)
}
buf.Truncate(buf.Len() - 2) // Remove trailing ", "
s := buf.String()           // Copy into a new string
fmt.Println(s)
```

```
1:2, 2:3, 3:5, 4:7, 5:11, 6:13

```

This solution is pretty efficient but may generate some excess garbage.
For higher performance, you can try to use the append functions
in package [`strconv`](https://golang.org/pkg/strconv/).

```
buf := []byte("Size: ")
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```

If the expected maximum length of the string is known,
you may want to preallocate the slice.

```
buf := make([]byte, 0, 16)
buf = append(buf, "Size: "...)
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```
