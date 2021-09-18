# Running Go CLI programs in the browser

# 在浏览器中运行 Go CLI 程序

Written on 5 Feb 2020

写于 2020 年 2 月 5 日

Turns out it’s almost shockingly easy to run Go CLI programs in the browser with WebAssembly (WASM); as an example I’ll use my [uni](https://github.com/arp242/uni) program. Building is as easy as:

事实证明，使用 WebAssembly (WASM) 在浏览器中运行 Go CLI 程序几乎非常容易；作为示例，我将使用我的 [uni](https://github.com/arp242/uni) 程序。构建非常简单：

```
GOOS=js GOARCH=wasm go build -o wasm/main.wasm
```

The resulting binary is rather large (5.1M); TinyGo can be used to create smaller builds, but it [doesn't support `os.Args` yet](https://github.com/tinygo-org/tinygo/issues/541), so it won't work here . After gzip compression it’s only 1.3M, so that’s manageable (and still smaller than many “text-only” websites).

生成的二进制文件相当大（5.1M）； TinyGo 可用于创建较小的构建，但它[尚不支持`os.Args`](https://github.com/tinygo-org/tinygo/issues/541)，所以它在这里不起作用. gzip 压缩后只有 1.3M，所以这是可以管理的（并且仍然小于许多“纯文本”网站)。

We then need to load the `main.wasm` binary:

然后我们需要加载 `main.wasm` 二进制文件：

```
<html>
<head>
    <meta charset="utf-8">
</head>
<body>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</body>
</html>
```

Copy the `wasm_exec.js` file from the Go source repo (`$(go env GOROOT)/misc/wasm/wasm_exec.js`), or [GitHub](https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js).

从 Go 源代码库 (`$(go env GOROOT)/misc/wasm/wasm_exec.js`) 或 [GitHub](https://github.com/golang/go/blob) 复制 `wasm_exec.js` 文件/master/misc/wasm/wasm_exec.js)。

You can’t load the HTML file the local filesystem as the browser will refuse to load the wasm file; you’ll have to use a webserver which serves wasm files with the correct MIME type, for example with Python:

您无法将 HTML 文件加载到本地文件系统，因为浏览器将拒绝加载 wasm 文件；你必须使用一个 web 服务器来提供具有正确 MIME 类型的 wasm 文件，例如使用 Python：

```
#!/usr/bin/env python3
import http.server
h = http.server.SimpleHTTPRequestHandler
h.extensions_map = {'': 'text/html', '.wasm': 'application/wasm', '.js': 'application/javascript'}
http.server.HTTPServer(('127.0.0.1', 2000), h).serve_forever()
```

Going to http://localhost:2000 will fetch the file and run `uni`; the JS console should display:

转到 http://localhost:2000 将获取文件并运行 `uni`； JS 控制台应显示：

```
uni: no command given
exit code 1
```

As if we typed `uni` on the CLI. To give it some arguments set `go.argv`:

就像我们在 CLI 上输入了 `uni` 一样。给它一些参数设置`go.argv`：

```
<script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            // Remember that argv[0] is the program name.
            go.argv = ['uni', '-q', 'identify', 'wasm'];
            go.run(result.instance);
        });
</script>
```

Which will give the expected output in the console:

这将在控制台中给出预期的输出：

```
'w'  U+0077  119    77          w     LATIN SMALL LETTER W (Lowercase_Letter)
'a'  U+0061  97     61          a     LATIN SMALL LETTER A (Lowercase_Letter)
's'  U+0073  115    73          s     LATIN SMALL LETTER S (Lowercase_Letter)
'm'  U+006D  109    6d          m     LATIN SMALL LETTER M (Lowercase_Letter)
```

Now it’s a simple matter of connecting an `input` element to `go.argv`; this also fetches the `main.wasm` just once and re-runs it, instead of re-fetching it every time:

现在是将 `input` 元素连接到 `go.argv` 的简单问题；这也只获取一次 `main.wasm` 并重新运行它，而不是每次都重新获取它：

```
<input id="input" style="font: 16px monospace">
<script src="wasm_exec.js"></script>
<script>
    fetch('main.wasm').then(response => response.arrayBuffer()).then(function(bin) {
            input.addEventListener('keydown', function(e) {
                if (e.keyCode !== 13)  // Enter
                    return;

                e.preventDefault();

                const go = new Go();
                go.argv = ['uni'].concat(this.value.split(' '));
                this.value = '';
                WebAssembly.instantiate(bin, go.importObject).then((result) => {
                    go.run(result.instance);
                });
            });
        });
</script>
```

Overwrite the `global.fs.writeSync` from `wasm_exec.js` to display the output in the HTML page instead of the console:

覆盖 `wasm_exec.js` 中的 `global.fs.writeSync` 以在 HTML 页面而不是控制台中显示输出：

```
<script>
    fetch('main.wasm').then(response => response.arrayBuffer()).then(function(bin) {
            input.addEventListener('keydown', function(e) {
                if (e.keyCode !== 13)  // Enter
                    return;

                e.preventDefault();

                const go = new Go();
                go.argv = ['uni'].concat(this.value.split(' '));
                this.value = '';

                // Write stdout to terminal.
                let outputBuf = '';
                const decoder = new TextDecoder("utf-8");
                global.fs.writeSync = function(fd, buf) {
                    outputBuf += decoder.decode(buf);
                    const nl = outputBuf.lastIndexOf("\n");
                    if (nl != -1) {
                        window.output.innerText += outputBuf.substr(0, nl + 1);
                        window.scrollTo(0, document.body.scrollHeight);
                        outputBuf = outputBuf.substr(nl + 1);
                    }
                    return buf.length;
                };

                WebAssembly.instantiate(bin, go.importObject).then((result) => {
                    go.run(result.instance);
                });
            });
        });
</script>
```

[📋 Copy](https://www.arp242.net/wasm-cli.html#)

[📋 复制](https://www.arp242.net/wasm-cli.html#)

And that’s pretty much it; 30 lines of JavaScript to run CLI applications in the browser :-) The only change I had to make to `uni` Go code was [adding a build tag](https://github.com/arp242/uni/commit/bfd9a565343bce6469c67ea2ae3accad597afcb4#diff-c5818bddd7e55bf1374be45465e95062).

差不多就是这样；在浏览器中运行 CLI 应用程序的 30 行 JavaScript :-) 我必须对 `uni` Go 代码所做的唯一更改是[添加构建标记](https://github.com/arp242/uni/commit/bfd9a565343bce6469c67ea2ae3accad597afcb4#diff-c5818bddd7e55bf1374be45465e95062)。

------

There are plenty of other things that can be improved: some better styling, reading from stdin, keybinds, loading indicator, etc. The [full version](https://arp242.github.io/uni-wasm/) does some of that. Take a look at [index.html](https://github.com/arp242/uni/blob/master/wasm/index.html) and [term.js](https://github.com/arp242/uni/blob/master/wasm/term.js) in case you're interested. It could still be improved further, but I thought this was “good enough” for a basic demo :-)

还有很多其他方面可以改进：一些更好的样式、从标准输入读取、键绑定、加载指示器等。 [完整版](https://arp242.github.io/uni-wasm/)做了一些那。看看[index.html](https://github.com/arp242/uni/blob/master/wasm/index.html)和[term.js](https://github.com/arp242/uni/blob/master/wasm/term.js) 以防您感兴趣。它仍然可以进一步改进，但我认为这对于基本演示来说“足够好”:-)

**Related articles**

**相关文章**

- [WebAssembly on the Go wiki](https://github.com/golang/go/wiki/WebAssembly)

- [Go wiki 上的 WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

**Feedback**

**回馈**

Contact me at                 [martin@arp242.net](mailto:martin@arp242.net),                 [GitHub](https://github.com/arp242/arp242.net/issues/new), or                 [@arp242_martin](https://twitter.com/arp242_martin)                 for feedback, questions, etc. 

通过 [martin@arp242.net](mailto:martin@arp242.net)、[GitHub](https://github.com/arp242/arp242.net/issues/new) 或 [@arp242_martin](https)与我联系://twitter.com/arp242_martin) 以获取反馈、问题等。

