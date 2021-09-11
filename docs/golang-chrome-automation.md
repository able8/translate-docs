# Automate Chrome with Golang and ChromeDP

April 4, 2018

Until recently I never knew how simple it could be to automate a task in the browser. A client wanted me to build simple automation  script for Chrome - it would log into his Drupal website, open Bootstrap settings and change cdn’s to the one found in config file. Sounded bit  hard at the beginning, but after playing an hour with chromedp it became quite trivial. Example repo is available on [GitHub](https://www.github.com/ribice/golang-chrome-automation).

I haven’t heard of Chrome Debugging Protocol before I was given this  task. I knew of Headless Chrome only, interfering with Chrome through  APIs only, creating a CLI tool.

On the other hand, ChromeDP launches a real browser instance.  Although the chromedp project claims to work with other browsers, namely Edge, Safari and Firefox, I have tested this only on Chrome, per  requirements.

You can learn more about ChromeDP from this talk at GopherCon Singapore:

<iframe src="https://www.youtube.com/embed/_7pWCg94sKw" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" allowfullscreen="" title="YouTube Video"></iframe>

ChromeDP’s source code and examples are located at [GitHub](https://github.com/chromedp/chromedp).

In short, the client requested the following from me:

- CLI tool that changes settings on a Drupal website
- Tool should be written in Go
- The only (large) dependency should be chromedp
- Config should be read from a text file (any format)
- Config should contain Drupal URL, Credentials, and CDN URLs
- The tool should open Drupal URL, log into it using provided username/password, change and save CDN settings

Following text contains code snippets with short comments. A working example is available on [GitHub](https://www.github.com/ribice/golang-chrome-automation).

#### Config struct

In order to read the config file from JSON to Golang, we a need struct for marshaling.

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

Reading from text files and marshaling into a struct is trivial using Go. I settled for JSON since Go has support for it in the standard  library.

```go
func readConfig() (*config, error) {
    _, filePath, _, _ := runtime.Caller(0)
    pwd := filePath[:len(filePath)-7]
    txt, err := ioutil.ReadFile(pwd + "/config.json")
    if err != nil {
        return nil, err
    }
    var cfg = new(config)
    if err := json.Unmarshal(txt, cfg); err != nil {
        return nil, err
    }
    return cfg, nil
}
```

Unlike the majority of things in Go, reading files and getting file  path can be done in several ways. I prefer using runtime.Caller() as it  lets me run the code from any location - it will always point to the  same directory.

The standard library provides a method for JSON unmarshalling, in JSON package.

#### Checking error

Although the code I delivered to the client didn’t contain this  function, I think it makes the code clearer. I’ve seen plenty of  projects handling fatal errors like this.

```go
func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
```

#### Starting Chrome Debugging Protocol

Most of the code related to starting/stopping chrome is available as  an example on chromedp’s repository. The below code creates a  cancellable context and passes it to chromedp.New(), a method that  starts the browser and logs everything to stdout.

```go
// create context
ctxt, cancel := context.WithCancel(context.Background())
defer cancel()

// create chrome instance
c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
```

### Executing commands in Chrome

The main implementation is located in changeDrupalSettings function.

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

Another option that comes to my mind, that is simple to use and  understand what it does, is chromedp.Text(), which selects text into a  Go variable.

Examples of all the mentioned methods and more are available on [ChromeDP Examples repositoriry](https://github.com/chromedp/examples). When running the application, it takes a few seconds to open Chrome  settings and enable the debug options. Also, for some reason, the last click on #edit-submit did not work on  Windows (but did on Mac and Linux). I haven’t debugged the application  on Windows so far.

Rest of the code gracefully stops Chrome and logs a success message.

```go
// shutdown chrome
checkErr(c.Shutdown(ctxt))

// wait for chrome to finish
checkErr(c.Wait())

log.Println("Successfully changed Drupal settings")
```

Once I find some spare time I’ll probably play more with ChromeDP, as this simple project was very interesting to me. There are lots of cool  things that could be built using it.

### Similar articles:

- [Refactoring Gorsk - Why and how](https://www.ribice.ba/refactoring-gorsk/)
- [Marshal YAML fields into map[string\]string](https://www.ribice.ba/golang-yaml-string-map/)
- [Twisk - Golang RPC starter kit](https://www.ribice.ba/twisk/)
- [Serve SwaggerUI within your Golang application](https://www.ribice.ba/serving-swaggerui-golang/)
- [Working with Go Web Frameworks - Gin and Echo](https://www.ribice.ba/golang-web-frameworks/)

2018 © Emir Ribic - [Some rights reserved](https://creativecommons.org/licenses/by/3.0/); please attribute properly and link back. Code snippets are [MIT Licensed](https://choosealicense.com/licenses/mit/)

Powered by [Hugo](https://gohugo.io/) & [Kiss](https://github.com/ribice/kiss).
