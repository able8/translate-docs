# Writing a Web Crawler with Golang and Colly

March 30, 2018

This blog features multiple posts regarding building Python web crawlers, but the subject of building a crawler in Golang has never been touched upon. There are a couple of frameworks for building web crawlers in Golang, but today we are going to look at building a web crawler using [Colly](https://github.com/gocolly/colly). When I first started playing with the framework, I was shocked how quick and easy it was to build a highly functional crawler with very few lines of Go code.

In this post we are going to build a crawler, which crawls this site and extracts the URL, title and code snippets from every Python post on the site. To write such a crawler we only need to write a total of 60 lines of code! Colly requires an understanding of CSS Selectors which is beyond the scope of this post, but I recommend you take a look at this [cheat sheet](https://www.w3schools.com/cssref/css_selectors.asp).

## Setting Up A Crawler

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

With this done, we can begin writing our main function. To create a new crawler we must create a NewCollector, which itself returns a Collector instance. The NewCollector function takes a list of functions which are used to initialize our crawler. In our case we are only calling one function within our NewCollector function, which is limiting our crawler to pages found on “edmundmartin.com”.

Having done this we then place some limits on our crawler. As Golang, is a very performant and many websites are running on relatively slow servers we probably want to limit the speed of our crawler. Here, we are setting up a limiter which matches everything contains “edmundmartin” in the URL. By setting the parallelism to 1 and setting a delay of a second, we are ensuring that we only crawl one URL a second.

## Basic Crawling Logic

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

We then add another “OnHTML” callback which is fired every time the HTML is returned to us. This is attached to our original Colly collector instance and not the clone of the Collector. Here we pass in CSS Selector, which pulls out all of the href’s on the page. We can also use some logic contained within the Colly framework which allows us to resolve to URL in question. If URL contains ‘python’, we submit it to our cloned to Collector, while if ‘python’ is absent from the URL we simply visit the page in question. This cloning of our collector allows us to define different OnHTML parsers for each clone of original crawler.

## Extracting Details From A Post

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

All there is left to do, is start of the crawler by calling our original collector’s ‘Visit’ function with our start URL. The example crawler should finish within around 20 seconds. Colly makes it very easy to write powerful crawlers with relatively little code. It does however take a little while to get used the callback style of the programming.

## Full Code

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

