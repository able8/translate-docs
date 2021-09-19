# Top 6 security best practices for Go

- December 19, 2019
- 6 minute read

Golang’s adoption has been increasing over the years. Successful projects like [Docker](https://blog.sqreen.com/docker-security/), [Kubernetes](https://blog.sqreen.com/kubernetes-security-best-practices/), and Terraform have bet heavily on this programming language. More recently, Go has been the de facto standard for building command-line tools. And for security matters, Go happens to be doing pretty well in their reports for vulnerabilities, with [only one CVE registry since 2002](https://www.cvedetails.com/vendor/1485/Golang.html).

However, not having vulnerabilities doesn’t mean that the programming language is super secure. We humans can create insecure apps if we don’t follow certain practices. For example, by following [the secure coding practices from OWASP](https://www.owasp.org/index.php/OWASP_Secure_Coding_Practices_-_Quick_Reference_Guide), we can determine how to apply these practices when using Go. And that’s exactly what I’ll do this time. In this post, I’ll show you six practices that you need to consider when developing with Go.

## 1\. Validate input entries

Validating entries from the user is not only for functionality purposes, but also helps avoid attackers who send us intrusive data that could damage the system. Moreover, you help users to use the tool more confidently by preventing them from making silly and common mistakes. For instance, you could prevent a user from trying to delete several records at the same time.

To validate user input, you can use native Go packages like `strconv` to handle string conversions to other data types. Go also has support for regular expressions with `regexp` for complex validations. Even though Go’s preference is to use native libraries, there are third-party packages like [validator](https://github.com/go-playground/validator). With validator, you can include validations for structs or individual fields more easily. For instance, the following code validates that the `User` struct contains a valid email address:

## 2\. Use HTML templates

One critical and common vulnerability is cross-site scripting or XSS. This exploit consists basically of the attacker being able to inject malicious code into the app to modify the output. For example, someone could send a JavaScript code as part of the query string in a URL. When the application returns the user’s value, the JavaScript code could be executed. Therefore, as a developer, you need to be aware of this and sanitize the user’s input.

Go has the package [html/template](https://golang.org/pkg/html/template/) to encode what the app will return to the user. So, instead of the browser executing an input like `<script>alert(‘You’ve Been Hacked!’);</script>`, popping up an alert message; you could encode the input, and the app will treat the input as a typical HTML code printed in the browser. An HTTP server that returns an HTML template will look like this:

But there are also third-party libraries you can use when developing web apps in Go. For example, there’s [Gorilla web toolkit](https://www.gorillatoolkit.org/), which includes libraries to help developers to do things like encoding authentication cookie values. There’s also [nosurf](https://github.com/justinas/nosurf), which is an HTTP package that helps with the prevention of cross-site request forgery ( [CSRF](https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF))).

## 3\. Protect yourself from SQL injections

If you’ve been a developer for a while, you might be aware of SQL injections, which is still number one on [OWASP’s Top 10](https://www.owasp.org/images/7/72/OWASP_Top_10-2017_%28en%29.pdf.pdf) list. However, there are some specific things that you need to consider when using Go. The first thing you need you to do is make sure the user that connects to the database has limited permissions. A good practice is to also sanitize the user’s input, as I described in a previous section, or to escape special characters and use [HTMLEscapeString](https://golang.org/pkg/html/template/#HTMLEscapeString) function from the HTML template package.

But, the most critical piece of code you’d need to include is the use of parameterized queries. In Go, you don’t prepare a statement in a connection; you prepare it on the DB. Here’s an example of how to use parameterized queries:

However, what if the database engine doesn’t support the use of prepared statements? Or what if it affects the performance of queries? Well, you can use the `db.Query()` function, but make sure you sanitize the user’s input first, as seen in previous sections. There are also third-party libraries like [sqlmap](https://github.com/sqlmapproject/sqlmap) to prevent SQL injections.

Despite our best efforts, sometimes vulnerabilities slip through, or enter our apps via third parties. To ensure that you protect your web apps from critical attacks like SQL injections, consider an application security management platform like [Sqreen](https://www.sqreen.com/).

## 4\. Encrypt sensitive information

Just because a string is hard to read, like a base-64 format, doesn’t mean that the hidden value is secret. You need a way to encrypt information that attackers can’t decode easily. Typical information that you’d like to encrypt are database passwords, user passwords, or even social security numbers.

OWASP has a few [recommendations](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#leverage-an-adaptive-one-way-function) of which encryption algorithms to use, such as `bcrypt`, `PDKDF2`, `Argon2` **,** or `scrypt`. Fortunately, there’s a Go package that includes robust implementations to encrypt information like [crypto](https://godoc.org/golang.org/x/crypto). For instance, the following code is a sample of how to use `bcrypt`:

Notice that you still need to be careful about how you transmit the information between services. You wouldn’t like to send the user’s data in plain text. It doesn’t matter if the app encrypts users’ inputs before persisting the data. Assume that someone on the internet could be sniffing your traffic and keeping request logs of your system. An attacker might use this information to correlate it with other data from other systems.

## 5\. Enforce HTTPS communication

Nowadays, most of the browsers require that HTTPS works on every site. Chrome, for example, will show you an alert if in the address bar if the site isn’t using HTTPS. An Infosec team could have as a policy to enforce in-transit encryption for communication between services. So, to secure in-transit connection in the system isn’t only about the app listening in port 443. You also need to use proper certificates and enforce HTTPS to avoid attackers downgrading the protocol to HTTP.

Here’s a code snippet for a web app that uses and enforces the HTTPS protocol:

Notice that the app will be listening in port 443. The following line is the one enforcing the HTTPS configuration:

You might also want to specify the server name in the TLS configuration, like this:

It’s always a good practice to implement in-transit encryption even if your web app is only for internal communication. Imagine if, for some reason, an attacker could sniff your internal traffic. Whenever you can, it’s  always best to raise the difficulty bar for possible future attackers.

## 6\. Be mindful with errors and logs

Last, but definitely not least, are error handling and logging in your Go apps.

To successfully troubleshoot in production, you need to instrument your apps properly. But you need to be mindful about the errors you show to users. You wouldn’t like users to know what exactly went wrong. Attackers might use this information to infer which services and technologies you’re using. Moreover, you have to remember that even though logs are great, they’re stored somewhere. And if logs end up in the wrong hands, they can be used to build an upcoming attack into the system.

So, the first thing you need to know or remember is that Go doesn’t have exceptions. This means that you’d need to handle errors differently than with other languages. The standard looks like this:

Also, Go offers a native library to work with logs. The most simple code is like this:

But there are third-party libraries for logging as well. A few of them are `logrus`, `glog`, and `loggo`. Here’s a small code snippet using `logrus`:

Finally, make sure you apply all the previous rules like encryption and sanitization of the data you put into the logs.

## There’s always room for improvement

These recommendations are the minimum things your Go apps should have. However, if your application is a command-line tool, you don’t need the in-transit encryption practices. But the rest of the above security tips apply to almost all app types. If you want to learn more, there’s a [book from OWASP specific to Go](https://github.com/OWASP/Go-SCP) that goes deeper into some topics. And there’s also a [repository](https://github.com/guardrailsio/awesome-golang-security) that includes links to frameworks, libraries, and other tools for security in Go.

No matter what you end up doing, remember that you can always improve your [app’s security](https://blog.sqreen.com/asm-series-a/). Attackers will always find new ways to exploit vulnerabilities, so do your best to continuously improve your security.

—-

_This post was written by Christian Meléndez._ [_Christian_](https://cmelendeztech.com/) _is a technologist that started as a software developer and has more recently become a cloud architect focused on implementing continuous delivery pipelines with applications in several flavors, including .NET, Node.js, and Java, often using Docker containers._
