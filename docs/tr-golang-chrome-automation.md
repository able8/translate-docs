# Automate Chrome with Golang and ChromeDP

# 使用 Golang 和 ChromeDP 自动化 Chrome

April 4, 2018

Until recently I never knew how simple it could be to automate a task in the browser. A client wanted me to build simple automation  script for Chrome - it would log into his Drupal website, open Bootstrap settings and change cdn’s to the one found in config file. Sounded bit  hard at the beginning, but after playing an hour with chromedp it became quite trivial. Example repo is available on [GitHub](https://www.github.com/ribice/golang-chrome-automation).

直到最近，我才知道在浏览器中自动执行任务是多么简单。一位客户希望我为 Chrome 构建简单的自动化脚本——它会登录他的 Drupal 网站，打开 Bootstrap 设置并将 cdn 更改为在配置文件中找到的那个。一开始听起来有点难，但是在用 chromedp 玩了一个小时后，它变得非常简单。 [GitHub](https://www.github.com/ribice/golang-chrome-automation) 上提供了示例存储库。

I haven’t heard of Chrome Debugging Protocol before I was given this  task. I knew of Headless Chrome only, interfering with Chrome through  APIs only, creating a CLI tool.

在我接到这个任务之前，我还没有听说过 Chrome 调试协议。我只知道 Headless Chrome，仅通过 API 干扰 Chrome，创建 CLI 工具。

On the other hand, ChromeDP launches a real browser instance. Although the chromedp project claims to work with other browsers, namely Edge, Safari and Firefox, I have tested this only on Chrome, per  requirements.

另一方面，ChromeDP 启动了一个真正的浏览器实例。尽管 chromedp 项目声称可以与其他浏览器（即 Edge、Safari 和 Firefox）一起使用，但我仅根据要求在 Chrome 上对此进行了测试。

You can learn more about ChromeDP from this talk at GopherCon Singapore:

你可以从 GopherCon Singapore 的这个演讲中了解更多关于 ChromeDP 的信息：

<iframe src="https://www.youtube.com/embed/_7pWCg94sKw" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" allowfullscreen="" title="YouTube Video"></iframe>

<iframe src="https://www.youtube.com/embed/_7pWCg94sKw" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" allowfullscreen="" title="YouTube 视频"></iframe>

ChromeDP’s source code and examples are located at [GitHub](https://github.com/chromedp/chromedp).

ChromeDP 的源代码和示例位于 [GitHub](https://github.com/chromedp/chromedp)。

In short, the client requested the following from me:

简而言之，客户向我提出以下要求：

- CLI tool that changes settings on a Drupal website
- Tool should be written in Go
- The only (large) dependency should be chromedp
- Config should be read from a text file (any format)
- Config should contain Drupal URL, Credentials, and CDN URLs
- The tool should open Drupal URL, log into it using provided username/password, change and save CDN settings

- 更改 Drupal 网站上的设置的 CLI 工具
- 工具应该用 Go 编写
- 唯一的（大）依赖应该是 chromedp
- 应该从文本文件（任何格式）中读取配置
- 配置应包含 Drupal URL、凭据和 CDN URL
- 该工具应打开 Drupal URL，使用提供的用户名/密码登录，更改并保存 CDN 设置

Following text contains code snippets with short comments. A working example is available on [GitHub](https://www.github.com/ribice/golang-chrome-automation).

以下文本包含带有简短注释的代码片段。 [GitHub](https://www.github.com/ribice/golang-chrome-automation) 上提供了一个工作示例。

#### Config struct

#### 配置结构

In order to read the config file from JSON to Golang, we need a struct for marshaling.

为了将配置文件从 JSON 读取到 Golang，我们需要一个用于编组的结构。

```go
type config struct {
    URL             string `json:"url"`
    Username        string `json:"username"`
    Password        string `json:"password"`
    BootstrapCSS    string `json:"bootstrap_css"`
    BootstrapCSSMin string `json:"bootstrap_css_min"`
    BootstrapJS     string `json:"bootstrap_js"`
    BootstrapJSMin  string `json:"bootstrap_js_min"`
}
```


#### Read config file

#### 读取配置文件

Reading from text files and marshaling into a struct is trivial using Go. I settled for JSON since Go has support for it in the standard  library.

使用 Go 从文本文件中读取并将其编组到结构中是微不足道的。我选择了 JSON，因为 Go 在标准库中支持它。

```go
func readConfig() (*config, error) {
    _, filePath, _, _ := runtime.Caller(0)
    pwd := filePath[:len(filePath)-7]
    txt, err := ioutil.ReadFile(pwd + "/config.json")
    if err != nil {
        return nil, err
    }
    var cfg = new(config)
    if err := json.Unmarshal(txt, cfg);err != nil {
        return nil, err
    }
    return cfg, nil
}
```


Unlike the majority of things in Go, reading files and getting file  path can be done in several ways. I prefer using runtime.Caller() as it  lets me run the code from any location - it will always point to the  same directory.

与 Go 中的大多数事情不同，读取文件和获取文件路径可以通过多种方式完成。我更喜欢使用 runtime.Caller() 因为它让我可以从任何位置运行代码 - 它总是指向同一个目录。

The standard library provides a method for JSON unmarshalling, in JSON package.

标准库在 JSON 包中提供了 JSON 解组的方法。

#### Checking error

#### 检查错误

Although the code I delivered to the client didn’t contain this  function, I think it makes the code clearer. I’ve seen plenty of  projects handling fatal errors like this.

虽然我提交给客户端的代码没有包含这个函数，但我认为它使代码更清晰。我见过很多项目处理这样的致命错误。

```go
func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
```


#### Starting Chrome Debugging Protocol

#### 启动 Chrome 调试协议

Most of the code related to starting/stopping chrome is available as  an example on chromedp’s repository. The below code creates a  cancellable context and passes it to chromedp.New(), a method that  starts the browser and logs everything to stdout.

大多数与启动/停止 chrome 相关的代码都可以在 chromedp 的存储库中作为示例使用。下面的代码创建了一个可取消的上下文并将其传递给 chromedp.New()，这是一个启动浏览器并将所有内容记录到标准输出的方法。

```go
// create context
ctxt, cancel := context.WithCancel(context.Background())
defer cancel()

// create chrome instance
c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
```


### Executing commands in Chrome

### 在 Chrome 中执行命令

The main implementation is located in changeDrupalSettings function.

主要实现位于 changeDrupalSettings 函数中。

```go
// Run executes the changeDrupalSettings on current chromeDP instance using the supplied context.
c.Run(ctxt, changeDrupalSettings(cfg))
func changeDrupalSettings(cfg *config) chromedp.Tasks {
return chromedp.Tasks{
    // Open Provided URL
    chromedp.Navigate(cfg.URL + "user/login"),

    // Wait until a CSS element is visible
    chromedp.WaitVisible(`#edit-name`, chromedp.ByID),

    // Pass value to a css element, in this case username and password from config file
    chromedp.SendKeys(`#edit-name`, cfg.Username, chromedp.ByID),
    chromedp.SendKeys(`#edit-pass`, cfg.Password, chromedp.ByID),

    // Click on log-in button
    chromedp.Click("#edit-submit"),

    // Wait until the user is logged in
    chromedp.Sleep(1 * time.Second),

    // Open bootstrap settings page
    chromedp.Navigate(cfg.URL + "admin/appearance/settings/bootstrap#edit-advanced"),
    chromedp.WaitVisible(`#edit-cdn`, chromedp.ByID),
    chromedp.Click(`#edit-cdn`),

    // Clicks on a dropdown select
    chromedp.Click(`#edit-cdn-provider`),

    // Chooses option that starts with c, "Custom", and selects it.
    // Alternatively down arrow could be pressed until Custom option was reached
    chromedp.SendKeys(`#edit-cdn-provider`, "c"+kb.Select, chromedp.ByID),
    chromedp.WaitVisible(`#edit-cdn-custom-css`, chromedp.ByID),

    // Clears the text box
    chromedp.Clear(`#edit-cdn-custom-css`),
    chromedp.Clear(`#edit-cdn-custom-css-min`),
    chromedp.Clear(`#edit-cdn-custom-js`),
    chromedp.Clear(`#edit-cdn-custom-js-min`),

    // Sends URLs from config file to css elements
    chromedp.SendKeys(`#edit-cdn-custom-css`, cfg.BootstrapCSS, chromedp.ByID),
    chromedp.SendKeys(`#edit-cdn-custom-css-min`, cfg.BootstrapCSSMin, chromedp.ByID),
    chromedp.SendKeys(`#edit-cdn-custom-js`, cfg.BootstrapJS, chromedp.ByID),
    chromedp.SendKeys(`#edit-cdn-custom-js-min`, cfg.BootstrapJSMin, chromedp.ByID),

    // Clicks on save button
    chromedp.Click("#edit-submit"),

    // Wait before closing Chrome
    chromedp.Sleep(1 * time.Second),
    }
}
```




Most of the options used below are easy to understand and work with - inspecting CSS and applying a command to it.

下面使用的大多数选项都很容易理解和使用 - 检查 CSS 并对其应用命令。

Another option that comes to my mind, that is simple to use and  understand what it does, is chromedp.Text(), which selects text into a  Go variable.

我想到的另一个选项是 chromedp.Text()，它易于使用且易于理解，它可以将文本选择到 Go 变量中。

Examples of all the mentioned methods and more are available on [ChromeDP Examples repositoriry](https://github.com/chromedp/examples). When running the application, it takes a few seconds to open Chrome  settings and enable the debug options. Also, for some reason, the last click on #edit-submit did not work on  Windows (but did on Mac and Linux). I haven’t debugged the application  on Windows so far.

[ChromeDP 示例存储库](https://github.com/chromedp/examples) 上提供了所有提到的方法的示例以及更多。运行应用程序时，打开 Chrome 设置并启用调试选项需要几秒钟的时间。此外，出于某种原因，最后一次点击 #edit-submit 在 Windows 上不起作用（但在 Mac 和 Linux 上起作用)。到目前为止，我还没有在 Windows 上调试该应用程序。

Rest of the code gracefully stops Chrome and logs a success message.

其余代码优雅地停止 Chrome 并记录成功消息。

```go
// shutdown chrome
checkErr(c.Shutdown(ctxt))

// wait for chrome to finish
checkErr(c.Wait())

log.Println("Successfully changed Drupal settings")
```


Once I find some spare time I’ll probably play more with ChromeDP, as this simple project was very interesting to me. There are lots of cool  things that could be built using it.

一旦我找到一些空闲时间，我可能会更多地使用 ChromeDP，因为这个简单的项目对我来说非常有趣。有很多很酷的东西可以使用它来构建。

### Similar articles:

### 类似文章：

- [Refactoring Gorsk - Why and how](https://www.ribice.ba/refactoring-gorsk/)
- [Twisk - Golang RPC starter kit](https://www.ribice.ba/twisk/)
- [Serve SwaggerUI within your Golang application](https://www.ribice.ba/serving-swaggerui-golang/)
- [Working with Go Web Frameworks - Gin and Echo](https://www.ribice.ba/golang-web-frameworks/)


2018 © Emir Ribic - [Some rights reserved](https://creativecommons.org/licenses/by/3.0/); please attribute properly and link back. Code snippets are [MIT Licensed](https://choosealicense.com/licenses/mit/)

2018 © Emir Ribic - [保留部分权利](https://creativecommons.org/licenses/by/3.0/)；请正确属性并链接回来。代码片段是 [MIT 许可](https://choosealicense.com/licenses/mit/)



