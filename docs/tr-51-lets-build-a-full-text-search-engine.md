## Let's build a Full-Text Search engine

## 让我们建立一个全文搜索引擎

- July 28, 2020 From: https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/

Full-Text Search is one of those tools people use every day without realizing it. If you ever googled "golang coverage report" or tried to find "indoor wireless camera" on an e-commerce website, you used some kind of full-text search.

全文搜索是人们每天都在不知不觉中使用的工具之一。如果您曾经在电子商务网站上搜索过“golang 覆盖率报告”或尝试查找“室内无线摄像头”，那么您使用了某种全文搜索。

Full-Text Search (FTS) is a technique for searching text in a collection of documents. A document can refer to a web page, a newspaper article, an email message, or any structured text.

全文搜索 (FTS) 是一种在文档集合中搜索文本的技术。文档可以指网页、报纸文章、电子邮件或任何结构化文本。

Today we are going to build our own FTS engine. By the end of this  post, we'll be able to search across millions of documents in less than a millisecond. We'll start with simple search queries like "give me all  documents that contain the word *cat*" and we'll extend the engine to support more sophisticated boolean queries.

今天我们将构建我们自己的 FTS 引擎。到本文结束时，我们将能够在不到一毫秒的时间内搜索数百万个文档。我们将从简单的搜索查询开始，例如“给我包含单词 *cat* 的所有文档”，我们将扩展引擎以支持更复杂的布尔查询。

Note: Most well-known FTS engine is [Lucene](https://lucene.apache.org/)(as well as [Elasticsearch](https://github.com/elastic/elasticsearch) and Solr built on top of it).

注意：最著名的 FTS 引擎是 [Lucene](https://lucene.apache.org/)（以及[Elasticsearch](https://github.com/elastic/elasticsearch) 和 Solr 建立在其之上）它）。

### Why FTS

### 为什么是 FTS

Before we start writing code, you may ask "can't we just use *grep* or have a loop that checks if every document contains the word I'm  looking for?". Yes, we can. However, it's not always the best idea.

在我们开始编写代码之前，您可能会问“我们不能只使用 *grep* 或有一个循环来检查每个文档是否包含我正在寻找的单词？”。我们可以。然而，这并不总是最好的主意。

### Corpus

### 语料库

We are going to search a part of the abstract of English Wikipedia. The latest dump is available at [dumps.wikimedia.org](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz). As of today, the file size after decompression is 913 MB. The XML file contains over 600K documents.

我们将搜索英文维基百科摘要的一部分。最新转储可在 [dumps.wikimedia.org](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz) 获得。截至今天，解压后的文件大小为 913 MB。 XML 文件包含超过 60 万个文档。

Document example:

文档示例：

```
<title>Wikipedia: Kit-Cat Klock</title>
<url>https://en.wikipedia.org/wiki/Kit-Cat_Klock</url>
<abstract>The Kit-Cat Klock is an art deco novelty wall clock shaped like a grinning cat with cartoon eyes that swivel in time with its pendulum tail.</abstract>
```


### Loading documents

### 加载文件

First, we need to load all the documents from the dump. The built-in encoding/xml package comes very handy:

首先，我们需要从转储中加载所有文档。内置的 encoding/xml 包非常方便：

```
import (
    "encoding/xml"
    "os"
)

type document struct {
    Title string `xml:"title"`
    URL   string `xml:"url"`
    Text  string `xml:"abstract"`
    ID    int
}

func loadDocuments(path string) ([]document, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    dec := xml.NewDecoder(f)
    dump := struct {
        Documents []document `xml:"doc"`
    }{}
    if err := dec.Decode(&dump);err != nil {
        return nil, err
    }

    docs := dump.Documents
    for i := range docs {
        docs[i].ID = i
    }
    return docs, nil
}
```


Every loaded document gets assigned a unique identifier. To keep things simple, the first loaded document gets assigned ID=0, the second ID=1 and so on.

每个加载的文档都会被分配一个唯一的标识符。为简单起见，第一个加载的文档被分配 ID=0，第二个 ID=1，依此类推。

### First attempt

###  第一次尝试

#### Searching the content

#### 搜索内容

Now that we have all documents loaded into memory, we can try to find the ones about cats. At first, let's loop through all documents and  check if they contain the substring *cat*:

现在我们已将所有文档加载到内存中，我们可以尝试查找有关猫的文档。首先，让我们遍历所有文档并检查它们是否包含子字符串 *cat*：

```
func search(docs []document, term string) []document {
    var r []document
    for _, doc := range docs {
        if strings.Contains(doc.Text, term) {
            r = append(r, doc)
        }
    }
    return r
}
```


On my laptop, the search phase takes 103ms - not too bad. If you spot check a few documents from the output, you may notice that the function matches *caterpillar* and *category*, but doesn't match *Cat* with the capital *C*. That's not quite what I was looking for.

在我的笔记本电脑上，搜索阶段需要 103 毫秒 - 还不错。如果您从输出中抽查一些文档，您可能会注意到该函数匹配 *caterpillar* 和 *category*，但不匹配带有大写字母 *C* 的 *Cat*。那不是我想要的。

We need to fix two things before moving forward:

在继续之前，我们需要解决两件事：

- Make the search case-insensitive (so *Cat* matches as well).
- Match on a word boundary rather than on a substring (so *caterpillar* and *communication* don't match).

- 使搜索不区分大小写（因此 *Cat* 也匹配）。
- 匹配单词边界而不是子字符串（因此 *caterpillar* 和 *communication* 不匹配）。

#### Searching with regular expressions

#### 使用正则表达式搜索

One solution that quickly comes to mind and allows implementing both requirements is *regular expressions*.

很快就会想到并允许实现这两个要求的一种解决方案是*正则表达式*。

Here it is - (?i)\bcat\b:

这是 - (?i)\bcat\b：

- (?i) makes the regex case-insensitive
- \b matches a word boundary (position where one side is a word character and another side is not a word character)

- (?i) 使正则表达式不区分大小写
- \b 匹配单词边界（一侧是单词字符而另一侧不是单词字符的位置）

```
func search(docs []document, term string) []document {
    re := regexp.MustCompile(`(?i)\b` + term + `\b`) // Don't do this in production, it's a security risk.term needs to be sanitized.
    var r []document
    for _, doc := range docs {
        if re.MatchString(doc.Text) {
            r = append(r, doc)
        }
    }
    return r
}
```




Ugh, the search took more than 2 seconds. As you can see, things  started getting slow even with 600K documents. While the approach is  easy to implement, it doesn't scale well. As the dataset grows larger,  we need to scan more and more documents. The time complexity of this  algorithm is linear - the number of documents required to scan is equal  to the total number of documents. If we had 6M documents instead of  600K, the search would take 20 seconds. We need to do better than that.

呃，搜索用了2秒多。如您所见，即使有 60 万个文档，事情也开始变慢。虽然该方法易于实施，但它的扩展性不佳。随着数据集越来越大，我们需要扫描越来越多的文档。该算法的时间复杂度是线性的——需要扫描的文档数量等于文档总数。如果我们有 6M 文档而不是 600K，搜索将需要 20 秒。我们需要做得比这更好。

### Inverted Index

### 倒排索引

To make search queries faster, we'll preprocess the text and build an index in advance.

为了使搜索查询更快，我们将预处理文本并提前构建索引。

The core of FTS is a data structure called *Inverted Index*. The Inverted Index associates every word in documents with documents that contain the word.

FTS 的核心是一个叫做*Inverted Index* 的数据结构。倒排索引将文档中的每个单词与包含该单词的文档相关联。

Example:

例子：

```
documents = {
    1: "a donut on a glass plate",
    2: "only the donut",
    3: "listen to the drum machine",
}

index = {
    "a": [1],
    "donut": [1, 2],
    "on": [1],
    "glass": [1],
    "plate": [1],
    "only": [2],
    "the": [2, 3],
    "listen": [3],
    "to": [3],
    "drum": [3],
    "machine": [3],
}
```


Below is a real-world example of the Inverted Index. An index in a book where a term references a page number:

以下是倒排索引的真实示例。书籍中的索引，其中术语引用了页码：

![img](https://artem.krylysov.com/images/2020-fts/book-index.png)

### Text analysis

### 文本分析

Before we start building the index, we need to break the raw text  down into a list of words (tokens) suitable for indexing and searching.

在开始构建索引之前，我们需要将原始文本分解为适合索引和搜索的单词（标记）列表。

The text analyzer consists of a tokenizer and multiple filters.

文本分析器由一个分词器和多个过滤器组成。

![img](https://artem.krylysov.com/images/2020-fts/text-analysis.png)

### Tokenizer

### 分词器

The tokenizer is the first step of text analysis. Its job is to  convert text into a list of tokens. Our implementation splits the text  on a word boundary and removes punctuation marks:

分词器是文本分析的第一步。它的工作是将文本转换为标记列表。我们的实现在单词边界上拆分文本并删除标点符号：

```
func tokenize(text string) []string {
    return strings.FieldsFunc(text, func(r rune) bool {
        // Split on any character that is not a letter or a number.
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    })
}
> tokenize("A donut on a glass plate. Only the donuts.")

["A", "donut", "on", "a", "glass", "plate", "Only", "the", "donuts"]
```


### Filters

### 过滤器

In most cases, just converting text into a list of tokens is not  enough. To make the text easier to index and search, we'll need to do  additional normalization.

在大多数情况下，仅将文本转换为标记列表是不够的。为了使文本更易于索引和搜索，我们需要进行额外的规范化。

#### Lowercase

#### 小写

In order to make the search case-insensitive, the lowercase filter converts tokens to lower case. *cAt*, *Cat* and *caT* are normalized to *cat*. Later, when we query the index, we'll lower case the search terms as well. This will make the search term *cAt* match the text *Cat*.

为了使搜索不区分大小写，小写过滤器将标记转换为小写。 *cAt*、*Cat* 和 *caT* 标准化为 *cat*。稍后，当我们查询索引时，我们也会小写搜索词。这将使搜索词 *cAt* 与文本 *Cat* 匹配。

```
func lowercaseFilter(tokens []string) []string {
    r := make([]string, len(tokens))
    for i, token := range tokens {
        r[i] = strings.ToLower(token)
    }
    return r
}
> lowercaseFilter([]string{"A", "donut", "on", "a", "glass", "plate", "Only", "the", "donuts"})

["a", "donut", "on", "a", "glass", "plate", "only", "the", "donuts"]
```


#### Dropping common words

#### 删除常用词

Almost any English text contains commonly used words like *a*, *I*, *the* or *be*. Such words are called *stop words*. We are going to remove them since almost any document would match the stop words.

几乎所有英文文本都包含常用词，例如 *a*、*I*、*the* 或 *be*。这样的词被称为*停止词*。我们将删除它们，因为几乎所有文档都会匹配停用词。

There is no "official" list of stop words. Let's exclude the top 10 by the [OEC rank](https://en.wikipedia.org/wiki/Most_common_words_in_English). Feel free to add more:

没有“官方”停用词列表。让我们按 [OEC 排名](https://en.wikipedia.org/wiki/Most_common_words_in_English) 排除前 10 名。随意添加更多：

```
var stopwords = map[string]struct{}{ // I wish Go had built-in sets.
    "a": {}, "and": {}, "be": {}, "have": {}, "i": {},
    "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
    r := make([]string, 0, len(tokens))
    for _, token := range tokens {
        if _, ok := stopwords[token];!ok {
            r = append(r, token)
        }
    }
    return r
}
> stopwordFilter([]string{"a", "donut", "on", "a", "glass", "plate", "only", "the", "donuts"})

["donut", "on", "glass", "plate", "only", "donuts"]
```


#### Stemming

#### 词干

Because of the grammar rules, documents may include different forms of the same word. Stemming reduces words into their base form. For example, *fishing*, *fished* and *fisher* may be reduced to the base form (stem) *fish*.

由于语法规则，文档可能包含同一个词的不同形式。词干将单词简化为基本形式。例如，*fishing*、*fished* 和 *fisher* 可以简化为基本形式（词干）*fish*。

Implementing a stemmer is a non-trivial task, it's not covered in this post. We'll take one of the [existing](https://github.com/kljensen/snowball) modules:

实现词干分析器是一项非常重要的任务，本文未涉及。我们将采用 [现有](https://github.com/kljensen/snowball) 模块之一：

```
import snowballeng "github.com/kljensen/snowball/english"

func stemmerFilter(tokens []string) []string {
    r := make([]string, len(tokens))
    for i, token := range tokens {
        r[i] = snowballeng.Stem(token, false)
    }
    return r
}
> stemmerFilter([]string{"donut", "on", "glass", "plate", "only", "donuts"})

["donut", "on", "glass", "plate", "only", "donut"]
```


Note

笔记

A stem is not always a valid word. For example, some stemmers may reduce *airline* to *airlin*.

词干并不总是一个有效的词。例如，一些词干分析器可能会将 *airline* 减少为 *airlin*。

### Putting the analyzer together

### 将分析器放在一起

```
func analyze(text string) []string {
    tokens := tokenize(text)
    tokens = lowercaseFilter(tokens)
    tokens = stopwordFilter(tokens)
    tokens = stemmerFilter(tokens)
    return tokens
}
```


The tokenizer and filters convert sentences into a list of tokens:

标记器和过滤器将句子转换为标记列表：

```
> analyze("A donut on a glass plate. Only the donuts.")

["donut", "on", "glass", "plate", "only", "donut"]
```


The tokens are ready for indexing.

令牌已准备好进行索引。

### Building the index

### 建立索引

Back to the inverted index. It maps every word in documents to document IDs. The built-in map is a good candidate for storing the mapping. The key in the map is a token (string) and the value is a list of document IDs:

回到倒排索引。它将文档中的每个单词映射到文档 ID。内置地图是存储映射的一个很好的候选者。映射中的键是一个标记（字符串），值是一个文档 ID 列表：

```
type index map[string][]int
```


Building the index consists of analyzing the documents and adding their IDs to the map:

构建索引包括分析文档并将其 ID 添加到 Map 中：

```
func (idx index) add(docs []document) {
    for _, doc := range docs {
        for _, token := range analyze(doc.Text) {
            ids := idx[token]
            if ids != nil && ids[len(ids)-1] == doc.ID {
                // Don't add same ID twice.
                continue
            }
            idx[token] = append(ids, doc.ID)
        }
    }
}

func main() {
    idx := make(index)
    idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
    idx.add([]document{{ID: 2, Text: "donut is a donut"}})
    fmt.Println(idx)
}
```


It works! Each token in the map refers to IDs of the documents that contain the token:

有用！map 中的每个令牌指的是包含该令牌的文档的 ID：

```
map[donut:[1 2] glass:[1] is:[2] on:[1] only:[1] plate:[1]]
```


### Querying

### 查询

To query the index, we are going to apply the same tokenizer and filters we used for indexing:

要查询索引，我们将应用我们用于索引的相同标记器和过滤器：

```
func (idx index) search(text string) [][]int {
    var r [][]int
    for _, token := range analyze(text) {
        if ids, ok := idx[token];ok {
            r = append(r, ids)
        }
    }
    return r
}
> idx.search("Small wild cat")

[[24, 173, 303, ...], [98, 173, 765, ...], [[24, 51, 173, ...]]
```


And finally, we can find all documents that mention cats. Searching 600K documents took less than a millisecond (18µs)!

最后，我们可以找到所有提到猫的文件。搜索 60 万个文档只用了不到一毫秒（18 微秒）！

With the inverted index, the time complexity of the search query is  linear to the number of search tokens. In the example query above, other than analyzing the input text, search had to perform only three map lookups.

使用倒排索引，搜索查询的时间复杂度与搜索标记的数量成线性关系。在上面的示例查询中，除了分析输入文本外，搜索只需要执行三个map查找。

### Boolean queries

### 布尔查询

The query from the previous section returned a disjoined list of documents for each token. What we normally expect to find when we type *small wild cat* in a search box is a list of results that contain *small*, *wild* and *cat* at the same time. The next step is to compute the set intersection  between the lists. This way we'll get a list of documents matching all  tokens.

上一节中的查询为每个标记返回了一个不连贯的文档列表。当我们在搜索框中输入 *small wild cat* 时，我们通常期望找到的是同时包含 *small*、*wild* 和 *cat* 的结果列表。下一步是计算列表之间的集合交集。通过这种方式，我们将获得与所有标记匹配的文档列表。

![img](https://artem.krylysov.com/images/2020-fts/venn.png)

Luckily, IDs in our inverted index are inserted in ascending order. Since the IDs are sorted, it's possible to compute the intersection between two lists in linear time. The intersection function iterates two lists simultaneously and collect IDs that exist in both:

幸运的是，倒排索引中的 ID 是按升序插入的。由于 ID 已排序，因此可以在线性时间内计算两个列表之间的交集。交集函数同时迭代两个列表并收集两者中存在的 ID：

```
func intersection(a []int, b []int) []int {
    maxLen := len(a)
    if len(b) > maxLen {
        maxLen = len(b)
    }
    r := make([]int, 0, maxLen)
    var i, j int
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            i++
        } else if a[i] > b[j] {
            j++
        } else {
            r = append(r, a[i])
            i++
            j++
        }
    }
    return r
}
```


Updated search analyzes the given query text, lookups tokens and computes the set intersection between lists of IDs:

更新的搜索分析给定的查询文本，查找标记并计算 ID 列表之间的集合交集：

```
func (idx index) search(text string) []int {
    var r []int
    for _, token := range analyze(text) {
        if ids, ok := idx[token];ok {
            if r == nil {
                r = ids
            } else {
                r = intersection(r, ids)
            }
        } else {
            // Token doesn't exist.
            return nil
        }
    }
    return r
}
```


The Wikipedia dump contains only two documents that match *small*, *wild* and *cat* at the same time:

维基百科转储仅包含两个同时匹配 *small*、*wild* 和 *cat* 的文档：

```
> idx.search("Small wild cat")

130764  The wildcat is a species complex comprising two small wild cat species, the European wildcat (Felis silvestris) and the African wildcat (F. lybica).
131692  Catopuma is a genus containing two Asian small wild cat species, the Asian golden cat (C. temminckii) and the bay cat.
```


The search is working as expected!

搜索按预期工作！

### Conclusions 

### 结论

We just built a Full-Text Search engine. Despite its simplicity, it can be a solid foundation for more advanced projects.

我们刚刚构建了一个全文搜索引擎。尽管它很简单，但它可以成为更高级项目的坚实基础。

I didn't touch on a lot of things that can significantly improve the performance and make the engine more user friendly. Here are some ideas for further improvements:

我没有涉及很多可以显着提高性能并使引擎更加用户友好的事情。以下是进一步改进的一些想法：

- Extend boolean queries to support *OR* and *NOT*.
- Store the index on disk:
   - Rebuilding the index on every application restart may take a while.
   - Large indexes may not fit in memory.
- Experiment with memory and CPU-efficient data formats for storing sets of document IDs. Take a look at [Roaring Bitmaps](https://roaringbitmap.org/).
- Support indexing multiple document fields.
- Sort results by relevance.



- 扩展布尔查询以支持 *OR* 和 *NOT*。
- 将索引存储在磁盘上：
  - 每次重新启动应用程序时重建索引可能需要一段时间。
  - 大索引可能不适合内存。
- 尝试使用内存和 CPU 高效的数据格式来存储文档 ID 集。看看 [Roaring Bitmaps](https://roaringbitmap.org/)。
- 支持索引多个文档字段。
- 按相关性对结果进行排序。

The full source code is available on [GitHub](https://github.com/akrylysov/simplefts).

完整的源代码可在 [GitHub](https://github.com/akrylysov/simplefts) 上找到。

I'm not a native English speaker and I'm trying to improve my language skills. Feel free to correct me if you spot any spelling or  grammatical error! 

我的母语不是英语，我正在努力提高我的语言技能。如果您发现任何拼写或语法错误，请随时纠正我！


