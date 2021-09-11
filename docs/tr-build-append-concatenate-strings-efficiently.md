# Efficient string concatenation [full guide]

# 高效的字符串连接 [完整指南]

yourbasic.org/golang

## Clean and simple string building

## 干净简单的字符串构建

For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. 
It’s clean, simple and fairly efficient.

对于性能不是问题的简单情况，[`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) 是你的朋友。
它干净、简单且相当有效。

```
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](http://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/)
lists the most common formatting verbs and flags.

[fmt 备忘单](http://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/)
列出了最常见的格式化动词和标志。

## High-performance string concatenation[Go 1.10](https://golang.org/doc/go1.10 "Go 1.10 Release Notes")

## 高性能字符串连接[Go 1.10](https://golang.org/doc/go1.10 "Go 1.10 Release Notes")

A [`strings.Builder`](https://golang.org/pkg/strings/#Builder) is used to efficiently append strings using write methods.

[`strings.Builder](https://golang.org/pkg/strings/#Builder) 用于使用写入方法有效地附加字符串。

- It offers a subset of the[`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) methods that allows it to safely avoid extra copying when converting a builder to a string.
- You can use the[`fmt`](https://golang.org/pkg/fmt/) package for formatting since the builder implements the [`io.Writer`](http://yourbasic.org/golang/io-writer-interface-explained/) interface.
- The[`Grow`](https://golang.org/pkg/strings/#Builder.Grow) method can be used to preallocate memory when the maximum size of the string is known.
   
- 它提供了 [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) 方法的一个子集这允许它在将构建器转换为字符串时安全地避免额外的复制。
- 您可以使用 [`fmt`](https://golang.org/pkg/fmt/) 包进行格式化因为构建器实现了 [`io.Writer`](http://yourbasic.org/golang/io-writer-interface-explained/) 接口。
- [`Grow`](https://golang.org/pkg/strings/#Builder.Grow) 方法当字符串的最大大小已知时，可用于预分配内存。

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

## 在 Go 1.10 之前

Use [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) to print into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer).

使用 [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) 打印到 [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer)。

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


This solution is pretty efficient but may generate some excess garbage. For higher performance, you can try to use the append functions in package [`strconv`](https://golang.org/pkg/strconv/).

此解决方案非常有效，但可能会产生一些多余的垃圾。 为了获得更高的性能，您可以尝试使用 append 函数
在包 [`strconv`](https://golang.org/pkg/strconv/) 中。

```
buf := []byte("Size: ")
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```


If the expected maximum length of the string is known, you may want to preallocate the slice.

如果字符串的预期最大长度已知，您可能想要预先分配切片。

```
buf := make([]byte, 0, 16)
buf = append(buf, "Size: "...)
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```



