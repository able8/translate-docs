# Top 6 security best practices for Go

# Go 的 6 大安全最佳实践

- December 19, 2019
- 6 minute read

- 2019 年 12 月 19 日
- 6 分钟阅读

Golang’s adoption has been increasing over the years. Successful projects like [Docker](https://blog.sqreen.com/docker-security/),[Kubernetes](https://blog.sqreen.com/kubernetes-security-best-practices/), and Terraform have bet heavily on this programming language. More recently, Go has been the de facto standard for building command-line tools. And for security matters, Go happens to be doing pretty well in their reports for vulnerabilities, with [only one CVE registry since 2002](https://www.cvedetails.com/vendor/1485/Golang.html).

多年来，Golang 的采用率一直在增加。 [Docker](https://blog.sqreen.com/docker-security/)、[Kubernetes](https://blog.sqreen.com/kubernetes-security-best-practices/)、Terraform等成功项目都有押注这种编程语言。最近，Go 已成为构建命令行工具的事实上的标准。在安全问题上，Go 恰好在他们的漏洞报告中做得很好，[自 2002 年以来只有一个 CVE 注册表](https://www.cvedetails.com/vendor/1485/Golang.html)。

However, not having vulnerabilities doesn’t mean that the programming language is super secure. We humans can create insecure apps if we don’t follow certain practices. For example, by following [the secure coding practices from OWASP](https://www.owasp.org/index.php/OWASP_Secure_Coding_Practices_-_Quick_Reference_Guide), we can determine how to apply these practices when using Go. And that’s exactly what I’ll do this time. In this post, I’ll show you six practices that you need to consider when developing with Go.

然而，没有漏洞并不意味着编程语言是超级安全的。如果我们不遵循某些做法，我们人类可能会创建不安全的应用程序。例如，通过遵循[来自 OWASP 的安全编码实践](https://www.owasp.org/index.php/OWASP_Secure_Coding_Practices_-_Quick_Reference_Guide)，我们可以确定在使用 Go 时如何应用这些实践。而这正是我这次要做的。在这篇文章中，我将向您展示使用 Go 进行开发时需要考虑的六个实践。

## 1\. Validate input entries

## 1\.验证输入条目

Validating entries from the user is not only for functionality purposes, but also helps avoid attackers who send us intrusive data that could damage the system. Moreover, you help users to use the tool more confidently by preventing them from making silly and common mistakes. For instance, you could prevent a user from trying to delete several records at the same time.

验证来自用户的条目不仅出于功能目的，还有助于避免攻击者向我们发送可能损坏系统的侵入性数据。此外，您可以防止用户犯愚蠢和常见的错误，从而帮助他们更自信地使用该工具。例如，您可以阻止用户尝试同时删除多条记录。

To validate user input, you can use native Go packages like `strconv` to handle string conversions to other data types. Go also has support for regular expressions with `regexp` for complex validations. Even though Go’s preference is to use native libraries, there are third-party packages like [validator](https://github.com/go-playground/validator). With validator, you can include validations for structs or individual fields more easily. For instance, the following code validates that the `User` struct contains a valid email address:

要验证用户输入，您可以使用像 `strconv` 这样的原生 Go 包来处理字符串到其他数据类型的转换。 Go 还支持带有 `regexp` 的正则表达式，用于复杂的验证。尽管 Go 偏好使用原生库，但也有第三方包，如 [validator](https://github.com/go-playground/validator)。使用验证器，您可以更轻松地包含对结构或单个字段的验证。例如，以下代码验证 `User` 结构是否包含有效的电子邮件地址：

## 2\. Use HTML templates

## 2\.使用 HTML 模板

One critical and common vulnerability is cross-site scripting or XSS. This exploit consists basically of the attacker being able to inject malicious code into the app to modify the output. For example, someone could send a JavaScript code as part of the query string in a URL. When the application returns the user’s value, the JavaScript code could be executed. Therefore, as a developer, you need to be aware of this and sanitize the user’s input.

跨站点脚本或 XSS 是一种关键且常见的漏洞。此漏洞利用基本上包括攻击者能够将恶意代码注入应用程序以修改输出。例如，某人可以将 JavaScript 代码作为 URL 中查询字符串的一部分发送。当应用程序返回用户的值时，可以执行 JavaScript 代码。因此，作为开发人员，您需要意识到这一点并清理用户的输入。

Go has the package [html/template](https://golang.org/pkg/html/template/) to encode what the app will return to the user. So, instead of the browser executing an input like `<script>alert(‘You’ve Been Hacked!’);</script>`, popping up an alert message; you could encode the input, and the app will treat the input as a typical HTML code printed in the browser. An HTTP server that returns an HTML template will look like this:

Go 有一个包 [html/template](https://golang.org/pkg/html/template/) 来编码应用程序将返回给用户的内容。因此，浏览器不会执行像 `<script>alert('Youve Being Hacked!');</script>` 这样的输入，而是弹出一条警告消息；您可以对输入进行编码，应用程序会将输入视为打印在浏览器中的典型 HTML 代码。返回 HTML 模板的 HTTP 服务器将如下所示：

But there are also third-party libraries you can use when developing web apps in Go. For example, there’s [Gorilla web toolkit](https://www.gorillatoolkit.org/), which includes libraries to help developers to do things like encoding authentication cookie values. There's also [nosurf](https://github.com/justinas/nosurf), which is an HTTP package that helps with the prevention of cross-site request forgery ( [CSRF](https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF))).

但是，在 Go 中开发 Web 应用程序时，您也可以使用第三方库。例如，有 [Gorilla web toolkit](https://www.gorillatoolkit.org/)，其中包含帮助开发人员执行诸如编码身份验证 cookie 值之类的库。还有 [nosurf](https://github.com/justinas/nosurf)，这是一个有助于防止跨站点请求伪造的 HTTP 包（[CSRF](https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)))。

## 3\. Protect yourself from SQL injections 

## 3. 保护自己免受 SQL 注入

If you've been a developer for a while, you might be aware of SQL injections, which is still number one on [OWASP's Top 10](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf) list. However, there are some specific things that you need to consider when using Go. The first thing you need you to do is make sure the user that connects to the database has limited permissions. A good practice is to also sanitize the user's input, as I described in a previous section, or to escape special characters and use [HTMLEscapeString](https://golang.org/pkg/html/template/#HTMLEscapeString) function from the HTML template package.

如果你做过一段时间的开发者，你可能知道 SQL 注入，它仍然是 [OWASP 的 Top 10](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf) 列表。但是，在使用 Go 时需要考虑一些特定的事情。您需要做的第一件事是确保连接到数据库的用户具有有限的权限。一个好的做法是也清理用户的输入，正如我在上一节中描述的那样，或者转义特殊字符并使用 [HTMLEscapeString](https://golang.org/pkg/html/template/#HTMLEscapeString) 函数从HTML 模板包。

But, the most critical piece of code you’d need to include is the use of parameterized queries. In Go, you don’t prepare a statement in a connection; you prepare it on the DB. Here’s an example of how to use parameterized queries:

但是，您需要包含的最关键的一段代码是使用参数化查询。在 Go 中，你不需要在连接中准备语句；你在数据库上准备它。以下是如何使用参数化查询的示例：

However, what if the database engine doesn’t support the use of prepared statements? Or what if it affects the performance of queries? Well, you can use the `db.Query()` function, but make sure you sanitize the user’s input first, as seen in previous sections. There are also third-party libraries like [sqlmap](https://github.com/sqlmapproject/sqlmap) to prevent SQL injections.

但是，如果数据库引擎不支持使用准备好的语句怎么办？或者如果它影响查询的性能怎么办？好吧，您可以使用 `db.Query()` 函数，但请确保首先清理用户的输入，如前几节所述。还有像[sqlmap](https://github.com/sqlmapproject/sqlmap)这样的第三方库来防止SQL注入。

Despite our best efforts, sometimes vulnerabilities slip through, or enter our apps via third parties. To ensure that you protect your web apps from critical attacks like SQL injections, consider an application security management platform like [Sqreen](https://www.sqreen.com/).

尽管我们尽了最大努力，但有时漏洞还是会漏掉，或通过第三方进入我们的应用程序。为确保您的 Web 应用程序免受 SQL 注入等关键攻击，请考虑使用应用程序安全管理平台，如 [Sqreen](https://www.sqreen.com/)。

## 4\. Encrypt sensitive information

## 4. 加密敏感信息

Just because a string is hard to read, like a base-64 format, doesn’t mean that the hidden value is secret. You need a way to encrypt information that attackers can’t decode easily. Typical information that you’d like to encrypt are database passwords, user passwords, or even social security numbers.

仅仅因为字符串难以阅读，如 base-64 格式，并不意味着隐藏的值是秘密的。您需要一种方法来加密攻击者无法轻松解码的信息。您想要加密的典型信息是数据库密码、用户密码，甚至社会安全号码。

OWASP has a few [recommendations](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#leverage-an-adaptive-one-way-function) of which encryption algorithms to use, such as `bcrypt`, ` PDKDF2`, `Argon2` **,** or `scrypt`. Fortunately, there’s a Go package that includes robust implementations to encrypt information like [crypto](https://godoc.org/golang.org/x/crypto). For instance, the following code is a sample of how to use `bcrypt`:

OWASP 有一些[推荐](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#leverage-an-adaptive-one-way-function) 使用哪些加密算法，例如`bcrypt`、` PDKDF2`、`Argon2` **、** 或 `scrypt`。幸运的是，有一个 Go 包，其中包含用于加密信息的强大实现，例如 [crypto](https://godoc.org/golang.org/x/crypto)。例如，以下代码是如何使用 `bcrypt` 的示例：

Notice that you still need to be careful about how you transmit the information between services. You wouldn’t like to send the user’s data in plain text. It doesn’t matter if the app encrypts users’ inputs before persisting the data. Assume that someone on the internet could be sniffing your traffic and keeping request logs of your system. An attacker might use this information to correlate it with other data from other systems.

请注意，您仍然需要注意如何在服务之间传输信息。您不希望以纯文本形式发送用户数据。应用程序是否在保存数据之前加密用户的输入并不重要。假设互联网上的某个人可能正在嗅探您的流量并保留您系统的请求日志。攻击者可能会使用此信息将其与来自其他系统的其他数据相关联。

## 5\. Enforce HTTPS communication

## 5. 强制 HTTPS 通信

Nowadays, most of the browsers require that HTTPS works on every site. Chrome, for example, will show you an alert if in the address bar if the site isn’t using HTTPS. An Infosec team could have as a policy to enforce in-transit encryption for communication between services. So, to secure in-transit connection in the system isn’t only about the app listening in port 443. You also need to use proper certificates and enforce HTTPS to avoid attackers downgrading the protocol to HTTP.

如今，大多数浏览器都要求 HTTPS 在每个站点上都能正常工作。例如，如果站点未使用 HTTPS，Chrome 会在地址栏中向您显示警报。信息安全团队可以将其作为一项策略来强制执行服务之间通信的传输中加密。因此，要保护系统中的传输中连接，不仅仅是应用程序在端口 443 中侦听。您还需要使用适当的证书并强制执行 HTTPS 以避免攻击者将协议降级为 HTTP。

Here’s a code snippet for a web app that uses and enforces the HTTPS protocol:

以下是使用和强制执行 HTTPS 协议的 Web 应用程序的代码片段：

Notice that the app will be listening in port 443. The following line is the one enforcing the HTTPS configuration:

请注意，该应用程序将侦听端口 443。以下行是强制执行 HTTPS 配置的行：

You might also want to specify the server name in the TLS configuration, like this:

您可能还想在 TLS 配置中指定服务器名称，如下所示：

It’s always a good practice to implement in-transit encryption even if your web app is only for internal communication. Imagine if, for some reason, an attacker could sniff your internal traffic. Whenever you can, it’s  always best to raise the difficulty bar for possible future attackers.

即使您的 Web 应用程序仅用于内部通信，实现传输中加密始终是一个好习惯。想象一下，如果出于某种原因，攻击者可以嗅探您的内部流量。只要有可能，最好为未来可能的攻击者提高难度。

## 6\. Be mindful with errors and logs

## 6\.注意错误和日志

Last, but definitely not least, are error handling and logging in your Go apps. 

最后但并非最不重要的是，错误处理和登录 Go 应用程序。

To successfully troubleshoot in production, you need to instrument your apps properly. But you need to be mindful about the errors you show to users. You wouldn’t like users to know what exactly went wrong. Attackers might use this information to infer which services and technologies you’re using. Moreover, you have to remember that even though logs are great, they’re stored somewhere. And if logs end up in the wrong hands, they can be used to build an upcoming attack into the system.

要在生产中成功排除故障，您需要正确检测您的应用程序。但是您需要注意向用户显示的错误。你不希望用户知道到底出了什么问题。攻击者可能会使用此信息来推断您正在使用哪些服务和技术。此外，您必须记住，即使日志很棒，它们也存储在某个地方。如果日志最终落入坏人之手，它们可用于在系统中构建即将到来的攻击。

So, the first thing you need to know or remember is that Go doesn’t have exceptions. This means that you’d need to handle errors differently than with other languages. The standard looks like this:

所以，你需要知道或记住的第一件事是 Go 没有例外。这意味着您需要以不同于其他语言的方式处理错误。该标准如下所示：

Also, Go offers a native library to work with logs. The most simple code is like this:

此外，Go 提供了一个本地库来处理日志。最简单的代码是这样的：

But there are third-party libraries for logging as well. A few of them are `logrus`, `glog`, and `loggo`. Here’s a small code snippet using `logrus`:

但是也有用于日志记录的第三方库。其中一些是`logrus`、`glog`和`loggo`。这是一个使用 `logrus` 的小代码片段：

Finally, make sure you apply all the previous rules like encryption and sanitization of the data you put into the logs.

最后，确保您应用了所有先前的规则，例如对您放入日志的数据进行加密和清理。

## There’s always room for improvement

## 总有改进的余地

These recommendations are the minimum things your Go apps should have. However, if your application is a command-line tool, you don’t need the in-transit encryption practices. But the rest of the above security tips apply to almost all app types. If you want to learn more, there’s a [book from OWASP specific to Go](https://github.com/OWASP/Go-SCP) that goes deeper into some topics. And there’s also a [repository](https://github.com/guardrailsio/awesome-golang-security) that includes links to frameworks, libraries, and other tools for security in Go.

这些建议是您的 Go 应用程序应该具备的最低限度的东西。但是，如果您的应用程序是命令行工具，则不需要传输中加密实践。但上述安全提示的其余部分适用于几乎所有应用程序类型。如果你想了解更多，有一本 [OWASP 专门针对 Go 的书](https://github.com/OWASP/Go-SCP) 可以更深入地探讨一些主题。还有一个 [repository](https://github.com/guardrailsio/awesome-golang-security)，其中包含框架、库和其他 Go 安全工具的链接。

No matter what you end up doing, remember that you can always improve your [app’s security](https://blog.sqreen.com/asm-series-a/). Attackers will always find new ways to exploit vulnerabilities, so do your best to continuously improve your security.

无论您最终做什么，请记住，您始终可以提高 [应用程序的安全性](https://blog.sqreen.com/asm-series-a/)。攻击者总会找到利用漏洞的新方法，因此请尽最大努力不断提高您的安全性。

—-

----

_This post was written by Christian Meléndez._ [_Christian_](https://cmelendeztech.com/) _is a technologist that started as a software developer and has more recently become a cloud architect focused on implementing continuous delivery pipelines with applications in several flavors , including .NET, Node.js, and Java, often using Docker containers._ 

_这篇文章由 Christian Meléndez 撰写。_ [_Christian_](https://cmelendeztech.com/) _是一名技术专家，最初是一名软件开发人员，最近成为一名云架构师，专注于使用多种风格的应用程序实现持续交付管道，包括 .NET、Node.js 和 Java，经常使用 Docker 容器。_

