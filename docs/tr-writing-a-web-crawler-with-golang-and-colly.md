# Writing a Web Crawler with Golang and Colly

# 用 Golang 和 Colly 编写一个网络爬虫

March 30, 2018



This blog features multiple posts regarding building Python web crawlers, but the subject of building a crawler in Golang has never been touched upon. There are a couple of frameworks for building web crawlers in Golang, but today we are going to look at building a web crawler using [Colly](https://github.com/gocolly/colly). When I first started playing with the framework, I was shocked how quick and easy it was to build a highly functional crawler with very few lines of Go code.

这个博客有很多关于构建 Python 网络爬虫的帖子，但是在 Golang 中构建爬虫的主题从未被触及。在 Golang 中有几个用于构建网络爬虫的框架，但今天我们将看看使用 [Colly](https://github.com/gocolly/colly) 构建网络爬虫。当我第一次开始使用这个框架时，我惊讶于用很少的 Go 代码来构建一个功能强大的爬虫是多么的快速和容易。

In this post we are going to build a crawler, which crawls this site and extracts the URL, title and code snippets from every Python post on the site. To write such a crawler we only need to write a total of 60 lines of code! Colly requires an understanding of CSS Selectors which is beyond the scope of this post, but I recommend you take a look at this [cheat sheet](https://www.w3schools.com/cssref/css_selectors.asp).

在这篇文章中，我们将构建一个爬虫，它会爬取该站点并从站点上的每个 Python 帖子中提取 URL、标题和代码片段。写这样一个爬虫我们总共只需要写60行代码！ Colly 需要了解 CSS 选择器，这超出了本文的范围，但我建议您查看此 [备忘单](https://www.w3schools.com/cssref/css_selectors.asp)。

## Setting Up A Crawler

## 设置爬虫

```
  import (
    "fmt"
    "strings"
    "time"
 
    "github.com/gocolly/colly"
)
 
type Article struct {
    ArticleTitle string
    URL          string
    CodeSnippets []string
}
 
func main() {
 
    c := colly.NewCollector(
        colly.AllowedDomains("edmundmartin.com"),
    )
 
    c.Limit(&colly.LimitRule{
        DomainGlob:  ".*edmundmartin.*",
        Parallelism: 1,
        Delay:       1 * time.Second,
    })
```


To begin with we are going to set up our crawler and create the data structure to store our results in. First, of all we need to install Colly using the go get command. Once this is done we create a new struct which will represent an article, and contains all the fields we are going to be collecting with our simple example crawler.

首先，我们将设置我们的爬虫并创建数据结构来存储我们的结果。首先，我们需要使用 go get 命令安装 Colly。完成此操作后，我们将创建一个新结构体，该结构体将代表一篇文章，并包含我们将使用我们的简单示例爬虫收集的所有字段。

With this done, we can begin writing our main function. To create a new crawler we must create a NewCollector, which itself returns a Collector instance. The NewCollector function takes a list of functions which are used to initialize our crawler. In our case we are only calling one function within our NewCollector function, which is limiting our crawler to pages found on “edmundmartin.com”.

完成后，我们可以开始编写我们的主要功能。要创建一个新的爬虫，我们必须创建一个 NewCollector，它本身返回一个 Collector 实例。 NewCollector 函数接受用于初始化我们的爬虫的函数列表。在我们的例子中，我们只调用 NewCollector 函数中的一个函数，这将我们的爬虫限制为在“edmundmartin.com”上找到的页面。

Having done this we then place some limits on our crawler. As Golang, is a very performant and many websites are running on relatively slow servers we probably want to limit the speed of our crawler. Here, we are setting up a limiter which matches everything contains “edmundmartin” in the URL. By setting the parallelism to 1 and setting a delay of a second, we are ensuring that we only crawl one URL a second.

完成此操作后，我们对爬虫设置了一些限制。由于 Golang 的性能非常好，而且许多网站都在相对较慢的服务器上运行，我们可能希望限制爬虫的速度。在这里，我们设置了一个限制器，它匹配 URL 中包含“edmundmartin”的所有内容。通过将并行度设置为 1 并设置一秒的延迟，我们可以确保我们每秒只抓取一个 URL。

## Basic Crawling Logic

## 基本爬行逻辑

```
detailCollector := c.Clone()

allArticles := []Article{}

c.OnRequest(func(r *colly.Request) {
      fmt.Println("Visiting: ", r.URL.String())
})

c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
      foundURL := e.Request.AbsoluteURL(e.Attr("href"))
      if strings.Contains(foundURL, "python") {
            detailCollector.Visit(foundURL)
      } else {
            c.Visit(foundURL)
      }
})
```


To collect data from our target site, we need to create a clone of our Colly collector. We also create a slice of our ‘Article’ struct to store the results we will be collecting. We also add a callback to our crawler which will fire every time we make a new request, this callback just prints the URL which are crawler will be visiting. 

要从目标站点收集数据，我们需要创建 Colly 收集器的克隆。我们还创建了“文章”结构的一部分来存储我们将收集的结果。我们还向我们的爬虫添加了一个回调，每次我们发出新请求时都会触发它，这个回调只打印爬虫将访问的 URL。

We then add another “OnHTML” callback which is fired every time the HTML is returned to us. This is attached to our original Colly collector instance and not the clone of the Collector. Here we pass in CSS Selector, which pulls out all of the href’s on the page. We can also use some logic contained within the Colly framework which allows us to resolve to URL in question. If URL contains ‘python’, we submit it to our cloned to Collector, while if ‘python’ is absent from the URL we simply visit the page in question. This cloning of our collector allows us to define different OnHTML parsers for each clone of original crawler.

然后我们添加另一个“OnHTML”回调，每次 HTML 返回给我们时都会触发它。这是附加到我们原始的 Colly 收集器实例，而不是收集器的克隆。这里我们传入 CSS Selector，它会拉出页面上的所有 href。我们还可以使用包含在 Colly 框架中的一些逻辑，它允许我们解析有问题的 URL。如果 URL 包含“python”，我们将其提交给我们的克隆到收集器，而如果 URL 中不包含“python”，我们只需访问相关页面。我们收集器的这种克隆允许我们为原始爬虫的每个克隆定义不同的 OnHTML 解析器。

## Extracting Details From A Post

## 从帖子中提取详细信息

```
detailCollector.OnHTML(`div.post-inner-content`, func(e *colly.HTMLElement) {
        fmt.Println("Scraping Content ", e.Request.URL.String())
        article := Article{}
        article.URL = e.Request.URL.String()
        article.ArticleTitle = e.ChildText("h1")

        e.ForEach("div.crayon-main", func(_ int, el *colly.HTMLElement) {
            codeSnip := el.ChildText("table.crayon-table")
            article.CodeSnippets = append(article.CodeSnippets, codeSnip)
        })
        fmt.Println("Found: ", article)
        allArticles = append(allArticles, article)
    })

    c.Visit("https://edmundmartin.com")
}
```


We can now add an ‘OnHTML’ callback to our ‘detailCollector’ clone. Again we use a CSS Selector to pull out the content of each post contained on the page. From this we can extract the text contained within the post’s “H1” tag. We finally, then pick out all of the ‘div’ containing the class ‘crayon-main’, we then iterate over all the elements pulling out our code snippets. We then add our collected data to our slice of Articles.

我们现在可以向我们的“detailCollector”克隆添加一个“OnHTML”回调。我们再次使用 CSS Selector 来提取页面上包含的每个帖子的内容。从中我们可以提取帖子的“H1”标签中包含的文本。最后，我们挑选出所有包含“crayon-main”类的“div”，然后遍历所有元素，提取出我们的代码片段。然后，我们将收集到的数据添加到我们的文章切片中。

All there is left to do, is start of the crawler by calling our original collector’s ‘Visit’ function with our start URL. The example crawler should finish within around 20 seconds. Colly makes it very easy to write powerful crawlers with relatively little code. It does however take a little while to get used the callback style of the programming.

剩下要做的就是通过使用我们的起始 URL 调用我们原始收集器的“访问”函数来启动爬虫。示例爬虫应在大约 20 秒内完成。 Colly 使得用相对较少的代码编写功能强大的爬虫变得非常容易。然而，使用编程的回调风格确实需要一些时间。

## Full Code

## 完整代码

```
package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/gocolly/colly"
)

type Article struct {
    ArticleTitle string
    URL          string
    CodeSnippets []string
}

func main() {

    c := colly.NewCollector(
        colly.AllowedDomains("edmundmartin.com"),
    )

    c.Limit(&colly.LimitRule{
        DomainGlob:  ".*edmundmartin.*",
        Parallelism: 1,
        Delay:       1 * time.Second,
    })

    detailCollector := c.Clone()

    allArticles := []Article{}

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting: ", r.URL.String())
    })

    c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
        foundURL := e.Request.AbsoluteURL(e.Attr("href"))
        if strings.Contains(foundURL, "python") {
            detailCollector.Visit(foundURL)
        } else {
            c.Visit(foundURL)
        }
    })

    detailCollector.OnHTML(`div.post-inner-content`, func(e *colly.HTMLElement) {
        fmt.Println("Scraping Content ", e.Request.URL.String())
        article := Article{}
        article.URL = e.Request.URL.String()
        article.ArticleTitle = e.ChildText("h1")

        e.ForEach("div.crayon-main", func(_ int, el *colly.HTMLElement) {
            codeSnip := el.ChildText("table.crayon-table")
            article.CodeSnippets = append(article.CodeSnippets, codeSnip)
        })
        fmt.Println("Found: ", article)
        allArticles = append(allArticles, article)
    })

    c.Visit("https://edmundmartin.com")
}
```



