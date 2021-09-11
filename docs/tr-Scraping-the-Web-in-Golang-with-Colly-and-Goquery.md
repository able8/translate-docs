# Scraping the Web in Golang with Colly and Goquery

使用 Colly 和 Goquery 在 Golang 中抓取 Web

March 1, 2018 From: https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/

If told to write a web crawler, the tools at the top of  my mind would be Python based: BeautifulSoup or Scrapy. However, the  ecosystem for writing web scrapers and crawlers in Go is quite robust. In particular, Colly and Goquery are extremely powerful tools that  afford a similar amount of expressiveness and flexibility to their  Python-based counterparts.

如果被告知要编写一个网络爬虫，我脑海中最重要的工具将是基于 Python 的：BeautifulSoup 或 Scrapy。然而，用 Go 编写网络爬虫和爬虫的生态系统非常强大。特别是，Colly 和 Goquery 是非常强大的工具，它们提供了与基于 Python 的对应工具相似的表达能力和灵活性。

## A Brief Introduction to Web Crawling

## 网络爬虫简介

What is a web crawler? Essentially, a web crawler works by inspecting the  HTML content of web pages and performing some type of action based on that content. Usually, pages are scraped for outbound links, which  the crawler places in a queue to visit. We can also save data extracted  from the current page. For example, if our crawler lands on a Wikipedia  page, we may save that page’s text and title.

什么是网络爬虫？本质上，网络爬虫通过检查网页的 HTML 内容并根据该内容执行某种类型的操作来工作。通常，页面会被抓取以获取出站链接，爬虫将其放入队列中以供访问。我们还可以保存从当前页面提取的数据。例如，如果我们的爬虫登陆维基百科页面，我们可能会保存该页面的文本和标题。

The simplest web crawlers perform the following algorithm:

最简单的网络爬虫执行以下算法：

```
initialize Queue
enqueue SeedURL

while Queue is not empty:
    URL = Pop element from Queue
    Page = Visit(URL)
    Links = ExtractLinks(Page)
    Enqueue Links on Queue
```


Our `Visit`  and `ExtractLinks`  functions are what changes; both are application specific. We might have a crawler that tries to interpret the entire graph of the web, like  Google does, or something simple that just scrapes Wikipedia.

我们的 `Visit` 和 `ExtractLinks` 函数发生了变化；两者都是特定于应用程序的。我们可能有一个爬虫试图解释整个网络图，就像谷歌那样，或者一些简单的东西只是抓取维基百科。

Things quickly become more complicated as your use case grows. Want  many, many more pages to be scraped? You might have to start looking  into a more sophisticated crawler that runs in parallel. Want to scrape  more complicated pages? You may need to find a more powerful HTML  parser.

随着您的用例的增长，事情很快变得更加复杂。想要抓取更多页面吗？您可能必须开始研究并行运行的更复杂的爬虫。想要抓取更复杂的页面？您可能需要找到更强大的 HTML 解析器。

## Colly

## 科利

Colly is a flexible framework for writing web crawlers in Go. It’s  very much batteries-included. Out of the box, you get support for: *Rate limiting* Parallel crawling*  Respecting `robots.txt` * HTML/Link parsing

Colly 是一个灵活的框架，用于在 Go 中编写网络爬虫。它包含很多电池。开箱即用，您将获得以下支持： * 速率限制 * 并行抓取 * 尊重 `robots.txt` * HTML/链接解析

The fundamental component of a Colly crawler is a “Collector”. Collectors keep track of pages that are queued to visit, and maintain callbacks for when a page is being scraped.

Colly 爬虫的基本组件是“收集器”。收集器跟踪排队访问的页面，并在页面被抓取时维护回调。

### Setup

###  设置

Creating a Colly collector is simple, but we have lots of options that we may elect to use:

创建一个 Colly 收集器很简单，但我们可以选择使用很多选项：

```go
c := colly.NewCollector(
    // Restrict crawling to specific domains
    colly.AllowedDomains("godoc.org"),
    // Allow visiting the same page multiple times
    colly.AllowURLRevisit(),
    // Allow crawling to be done in parallel / async
    colly.Async(true),
)
```


Of course, you can also just stick with a bare `colly.NewCollector()` and handle these addons yourself.

当然，你也可以坚持使用一个空的 `colly.NewCollector()` 并自己处理这些插件。

We might also want to place specific limits on our crawler’s behavior to be good web citizens. Colly makes it easy to introduce rate  limiting:

为了成为优秀的网络公民，我们可能还想对我们的爬虫行为设置特定的限制。 Colly 可以轻松引入速率限制：

```go
c.Limit(&colly.LimitRule{
    // Filter domains affected by this rule
    DomainGlob:  "godoc.org/*",
    // Set a delay between requests to these domains
    Delay: 1 * time.Second
    // Add an additional random delay
    RandomDelay: 1 * time.Second,
})
```


Some websites are more picky than others when it  comes to the amount of traffic they allow before cutting you off. Generally, setting a delay of a couple seconds should keep you off the  “naughty list”.

有些网站在切断您之前允许的流量方面比其他网站更挑剔。一般来说，设置几秒钟的延迟应该会让你远离“顽皮的名单”。

From here, we can start our collector by seeding it with a URL:

从这里，我们可以通过用 URL 播种来启动我们的收集器：

```go
c.Visit("https://godoc.org")
```


### OnHTML



We have a collector that plays nice which can start at an arbitrary website. Now, we want our collector to *do* something — it needs to inspect pages so it can extract links and other data.

我们有一个很好的收集器，它可以从任意网站开始。现在，我们希望我们的收集器*做*一些事情——它需要检查页面，以便它可以提取链接和其他数据。

The `colly.Collector.OnHTML` method allows you to register a callback for when the collector reaches a portion of a page that  matches a specific HTML tag specifier. For starters, we can get a  callback whenever our crawler sees an `<a>` tag that contains an `href` link.

`colly.Collector.OnHTML` 方法允许您在收集器到达与特定 HTML 标记说明符匹配的页面部分时注册回调。对于初学者来说，只要我们的爬虫看到一个包含 `href` 链接的 `<a>` 标签，我们就可以得到一个回调。

```
c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    // Extract the link from the anchor HTML element
    link := e.Attr("href")
    // Tell the collector to visit the link
    c.Visit(e.Request.AbsoluteURL(link))
})
```


As seen above, in the callback you’re given a `colly.HTMLElement` that contains the matching HTML data.

如上所示，在回调中，您将获得一个包含匹配 HTML 数据的 `colly.HTMLElement`。

Now, we have the beginnings of an actual web crawler: we find links  on the pages we visit, and tell our collector to visit those links in  subsequent requests. 

现在，我们开始了一个真正的网络爬虫：我们在我们访问的页面上找到链接，并告诉我们的收集器在后续请求中访问这些链接。

`OnHTML` is a powerful tool. It can search for CSS selectors (i.e. `div.my_fancy_class` or `#someElementId`), and you can attach multiple `OnHTML` callbacks to your collector to handle different page types.

`OnHTML` 是一个强大的工具。它可以搜索 CSS 选择器（即 `div.my_fancy_class` 或 `#someElementId`），并且您可以将多个 `OnHTML` 回调附加到您的收集器以处理不同的页面类型。

Colly’s `HTMLElement`  struct is quite useful. In addition to getting attributes with the `Attr` function, you can also extract text. For example, we may want to print a page’s title:

Colly 的 `HTMLElement` 结构非常有用。除了使用 Attr 函数获取属性外，您还可以提取文本。例如，我们可能想打印一个页面的标题：

```
c.OnHTML("title", func(e *colly.HTMLElement) {
    fmt.Println(e.Text)
})
```


### OnRequest / OnResponse

### 

There may be times when you don’t need a specific HTML element from a page, but instead want to know when your crawler is about to retrieve  or has just retrieved a page. For this, Colly exposes the  `OnRequest` and `OnResponse` callbacks.

有时您可能不需要页面中的特定 HTML 元素，而是想知道您的爬虫何时将要检索或刚刚检索到一个页面。为此，Colly 公开了`OnRequest` 和`OnResponse` 回调。

All of these callbacks will be called for each visited page. As for how this fits in with `OnHTML`. Here is the order in which callbacks are called per page:

将为每个访问的页面调用所有这些回调。至于这如何与 `OnHTML` 配合。以下是每页调用回调的顺序：

1. `OnRequest`
2. `OnResponse`
3. `OnHTML`
4. `OnScraped` (not referenced in this post, but may be useful to you)

1.`请求`
2.`OnResponse`
3.`OnHTML`
4.`OnScraped`（本文未提及，但可能对你有用）

Of particular use is the ability to abort a request within the `OnRequest` callback. This may be useful for when you want your collector to stop.

特别有用的是在“OnRequest”回调中中止请求的能力。当您希望收集器停止时，这可能很有用。

```go
numVisited := 0
c.OnRequest(func(r *colly.Request) {
    if numVisited > 100 {
        r.Abort()
    }
    numVisited++
})
```


In `OnResponse`, you have access to the entire HTML document, which could be useful in certain contexts:

在 OnResponse 中，您可以访问整个 HTML 文档，这在某些情况下可能很有用：

```go
c.OnResponse(func(r *colly.Response) {
    fmt.Println(r.Body)
})
```


### HTMLElement

### HTML元素

In addition to the `Attr()` method and `Text` property that `colly.HTMLElement` has, we can also use it to traverse child elements. The [`ChildAttr()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ChildAttr),[`ChildText()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ChildText), and [`ForEach()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ForEach) methods in particular are quite useful.

除了`colly.HTMLElement`具有的`Attr()`方法和`Text`属性，我们还可以使用它来遍历子元素。 [`ChildAttr()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ChildAttr),[`ChildText()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ChildText) 和 [`ForEach()`](https://godoc.org/github.com/gocolly/colly#HTMLElement.ForEach) 方法特别有用。

For example, we can use `ChildText()` to get the text of all the paragraphs in a section:

例如，我们可以使用 `ChildText()` 来获取一个部分中所有段落的文本：

```go
c.OnHTML("#myCoolSection", func(e *colly.HTMLElement) {
    fmt.Println(e.ChildText("p"))
})
```


And we can use `ForEach()` to iterate over an elements children that match a specific selector:

我们可以使用 `ForEach()` 来迭代匹配特定选择器的元素子元素：

```go
c.OnHTML("#myCoolSection", func(e *colly.HTMLElement) {
    e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
        if strings.Contains(elem.Text, "golang") {
            fmt.Println(elem.Text)
        }
    })
})
```


## Bringing in Goquery

## 引入 Goquery

Colly’s built-in `HTMLElement` is useful for most scraping tasks, but if we want to do particularly advanced traversals of a DOM,  we’ll have to look elsewhere. For example, there’s no way (currently) to traverse up the DOM to parent elements or traverse laterally through  sibling elements.

Colly 的内置 `HTMLElement` 对大多数抓取任务都很有用，但是如果我们想要对 DOM 进行特别高级的遍历，我们将不得不寻找其他地方。例如，（目前）没有办法向上遍历 DOM 到父元素或横向遍历兄弟元素。

Enter *Goquery*, “like that j-thing, only in Go”. It’s  basically jQuery. In Go. (Which is awesome) For anything you’d like to  scrape from an HTML document, it can probably be done using Goquery.

输入 *Goquery*，“就像那个 j-thing，只在 Go 中”。它基本上是 jQuery。在去。 （这太棒了）对于您想从 HTML 文档中抓取的任何内容，可能可以使用 Goquery 来完成。

While Goquery is modeled off of jQuery, I found it to be pretty  similar in many respects to the BeautifulSoup API. So, if you’re coming  from the Python scraping world, then you’ll probably find yourself  comfortable with Goquery.

虽然 Goquery 是基于 jQuery 建模的，但我发现它在许多方面与 BeautifulSoup API 非常相似。因此，如果您来自 Python 抓取世界，那么您可能会发现自己对 Goquery 很满意。

Goquery allows us to do more complicated HTML selections and DOM traversals than Colly’s `HTMLElement` affords. For example, we may want to find the sibling elements of our  anchor element, to get some context around the link we’ve scraped:

Goquery 允许我们进行比 Colly 的 HTMLElement 更复杂的 HTML 选择和 DOM 遍历。例如，我们可能想要找到锚元素的兄弟元素，以获取我们抓取的链接的上下文：

```go
dom, _ := qoquery.NewDocument(htmlData)
dom.Find("a").Siblings().Each(func(i int, s *goquery.Selection) {
    fmt.Printf("%d, Sibling text: %s\n", i, s.Text())
})
```


Also, we can easily find the parent of a selected  element. This might be useful if we’re given an anchor tag from Colly,  and we want to find the content of the pages `<title>` tag:

此外，我们可以轻松找到所选元素的父级。如果我们从 Colly 获得了一个锚标签，并且我们想要找到页面 `<title>` 标签的内容，这可能会很有用：

```go
anchor.ParentsUntil("~").Find("title").Text()
```


`ParentsUntil` traverses up the DOM until it finds something that matches the passed selector. We can use `~` to traverse all the way up to the top of the DOM, which then allows us to easily grab the title tag. 

`ParentsUntil` 向上遍历 DOM，直到找到与传递的选择器匹配的内容。我们可以使用 `~` 一直遍历到 DOM 的顶部，这样我们就可以轻松地获取标题标签。

This is really just scratching the surface of what Goquery can do. So far, we’ve seen examples of DOM traversal, but Goquery also has robust  support for DOM manipulation — editing text, adding/removing classes or  properties, inserting/removing HTML elements, etc.

这实际上只是 Goquery 可以做的事情的皮毛。到目前为止，我们已经看到了 DOM 遍历的例子，但 Goquery 也对 DOM 操作提供了强大的支持——编辑文本、添加/删除类或属性、插入/删除 HTML 元素等。

Bringing it back to web scraping, how do we use Goquery with Colly? It’s straightforward: each Colly HTMLElement contains a Goquery  selection, which you can access through the `DOM` property.

回到网络抓取，我们如何将 Goquery 与 Colly 一起使用？很简单：每个 Colly HTMLElement 都包含一个 Goquery 选择，您可以通过 `DOM` 属性访问它。

```go
c.OnHTML("div", func(e *colly.HTMLElement) {
    // Goquery selection of the HTMLElement is in e.DOM
    goquerySelection := e.DOM

    // Example Goquery usage
    fmt.Println(qoquerySelection.Find(" span").Children().Text())
})
```


It’s worth noting that most scraping tasks can be framed in such a way that you don’t *need* to use Goquery! Simply add an OnHTML callback for `html`, and you can get access to the entire page that way. However, I still  found that Goquery was a nice addition to my DOM traversal toolbelt.

值得注意的是，大多数抓取任务都可以以这样的方式构建，您不需要*使用 Goquery！只需为 `html` 添加一个 OnHTML 回调，您就可以通过这种方式访问整个页面。然而，我仍然发现 Goquery 是我的 DOM 遍历工具带的一个很好的补充。

## Writing a full web crawler

## 编写一个完整的网络爬虫

Using Colly and Goquery[1](https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/#fn:1), we can pretty easily piece together a simple web crawler.

使用 Colly 和 Goquery[1](https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/#fn:1)，我们可以很容易拼凑一个简单的网络爬虫。

With all the pieces explored above, we can write a simple web crawler that scrapes [Emojipedia](https://emojipedia.org/) for emoji descriptions.

有了上面探索的所有部分，我们可以编写一个简单的网络爬虫来抓取 [Emojipedia](https://emojipedia.org/) 以获取表情符号描述。

```go
package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("emojipedia.org"),
    )

    // Callback for when a scraped page contains an article element
    c.OnHTML("article", func(e *colly.HTMLElement) {
        isEmojiPage := false

        // Extract meta tags from the document
        metaTags := e.DOM.ParentsUntil("~").Find("meta")
        metaTags.Each(func(_ int, s *goquery.Selection) {
            // Search for og:type meta tags
            property, _ := s.Attr("property")
            if strings.EqualFold(property, "og:type") {
                content, _ := s.Attr("content")

                // Emoji pages have "article" as their og:type
                isEmojiPage = strings.EqualFold(content, "article")
            }
        })

        if isEmojiPage {
            // Find the emoji page title
            fmt.Println("Emoji: ", e.DOM.Find("h1").Text())
            // Grab all the text from the emoji's description
            fmt.Println(
                "Description: ",
                e.DOM.Find(".description").Find("p").Text())
        }
    })

    // Callback for links on scraped pages
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        // Extract the linked URL from the anchor tag
        link := e.Attr("href")
        // Have our crawler visit the linked URL
        c.Visit(e.Request.AbsoluteURL(link))
    })

    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        RandomDelay: 1 * time.Second,
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.Visit("https://emojipedia.org")
}
```


Full code can be found [here](https://github.com/bcongdon/colly-example) on Github.

完整代码可以在 Github 上找到 [这里](https://github.com/bcongdon/colly-example)。

And that’s it! After compiling and running, you’ll see the crawler  visiting a number of pages, and print out emoji names / descriptions  when it stumbles onto an emoji page.

就是这样！编译和运行后，你会看到爬虫访问了许多页面，并在它偶然发现一个表情符号页面时打印出表情符号名称/描述。

Clearly, this is just the beginning. One could easily save this data  in a graph structure, or expose a web parser/scraper as a distinct  package for a site that doesn’t have a public API.

显然，这只是开始。人们可以轻松地将这些数据保存在图形结构中，或者将 Web 解析器/抓取器公开为没有公共 API 的站点的独特包。

The nice thing about Colly is that it scales with your use case. On the more advanced end of the spectrum, it [supports](http://go-colly.org/docs/examples/redis_backend/) using Redis as a backend for holding queued pages, [parallel scraping](http://go-colly.org/docs/examples/parallel/), and the use of [multiple collectors](http://go-colly.org/docs/best_practices/multi_collector/) running simultaneously.

Colly 的好处在于它可以根据您的用例进行扩展。在更高级的一端，它 [支持](http://go-colly.org/docs/examples/redis_backend/) 使用 Redis 作为后端来保存排队页面，[并行抓取](http://go-colly.org/docs/examples/parallel/)，并使用 [多个收集器](http://go-colly.org/docs/best_practices/multi_collector/) 同时运行。

## Where to Go from here?

##  从这往哪儿走？

The Colly [documentation](http://go-colly.org/) website is a great resource, and has tons of practicable examples. Colly’s [Godoc](https://godoc.org/github.com/gocolly/colly) and Goquery’s [Godoc](https://godoc.org/github.com/PuerkitoBio/goquery) are also good places to look.

Colly [文档](http://go-colly.org/) 网站是一个很好的资源，有大量实用的例子。 Colly 的 [Godoc](https://godoc.org/github.com/gocolly/colly) 和 Goquery 的 [Godoc](https://godoc.org/github.com/PuerkitoBio/goquery) 也是不错的地方。

You may also find [gocrawl](https://github.com/PuerkitoBio/gocrawl) to be worth looking into. It’s written by the same developer that created Goquery.

您可能还会发现 [gocrawl](https://github.com/PuerkitoBio/gocrawl) 值得研究。它是由创建 Goquery 的同一位开发人员编写的。

## Footnotes

## 脚注

------



1. I got a very nice followup from the creator of colly, [asciimoo](https://github.com/asciimoo), who pointed out that this example could be done entirely with colly’s `HTMLElement`. Below is his implimentation of the emoji page scraper that works without calling out to Goquery:

    1. 我从 colly 的创建者 [asciimoo](https://github.com/asciimoo) 那里得到了一个非常好的跟进，他指出这个例子完全可以用 colly 的 `HTMLElement` 来完成。下面是他对无需调用 Goquery 即可工作的表情符号页面抓取器的实现：

   ```go
    c.OnHTML("html", func(e *colly.HTMLElement) {
       if strings.EqualFold(e.ChildAttr(`meta[property="og:type"]`, "content"), "article") {
           // Find the emoji page title
           fmt.Println("Emoji: ", e.ChildText("article h1"))
           // Grab all the text from the emoji's description
           fmt.Println("Description: ", e.ChildText("article .description p"))
       }
   })
   ```

   



------



标记为：[golang](https://benjamincongdon.me/tags/#golang)、[web](https://benjamincongdon.me/tags/#web)、[教程](https://benjamincongdon.我/标签/#tutorial)

Related Posts:

相关文章：

- [How to Deploy a Secure Static Site to AWS with S3 and CloudFront](https://benjamincongdon.me/blog/2017/06/13/How-to-Deploy-a-Secure-Static-Site-to-AWS-with-S3-and-CloudFront/)
- [25 Days of using Golang](https://benjamincongdon.me/blog/2018/01/17/25-Days-of-using-Golang/)
- [Learning to Like Java](https://benjamincongdon.me/blog/2017/11/21/Learning-to-Like-Java/)

- [如何使用 S3 和 CloudFront 将安全静态站点部署到 AWS](https://benjamincongdon.me/blog/2017/06/13/How-to-Deploy-a-Secure-Static-Site-to-AWS-with-S3-and-CloudFront/)
- [使用 Golang 的 25 天](https://benjamincongdon.me/blog/2018/01/17/25-Days-of-using-Golang/)
- [学习喜欢Java](https://benjamincongdon.me/blog/2017/11/21/Learning-to-Like-Java/)

© Ben Congdon 2015-2021 

