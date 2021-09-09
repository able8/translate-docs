# [Executable PNGs](https://djharper.dev/post/2020/12/26/executable-pngs/ "Executable PNGs")

# [可执行的 PNG](https://djharper.dev/post/2020/12/26/executable-pngs/ "可执行的 PNG")

Saturday, December 26, 2020

2020 年 12 月 26 日星期六

It's an image _and_ a program

这是一个图像_和_一个程序

A few weeks ago I was reading about [PICO-8](https://www.lexaloffle.com/pico-8.php), a fantasy games console with limited constraints. What really piqued my interest about it was the novel way games are distributed, you encode them into a PNG image. This includes the game code, assets, everything. The image can be whatever you want, screenshots from the game, cool artwork or just text. To load them you pass the image as input to the PICO-8 program and start playing.

几周前，我正在阅读 [PICO-8](https://www.lexaloffle.com/pico-8.php)，这是一个限制有限的幻想游戏机。真正引起我兴趣的是游戏分发的新颖方式，您可以将它们编码为 PNG 图像。这包括游戏代码，资产，一切。图像可以是任何你想要的，游戏截图，很酷的艺术品或只是文字。要加载它们，您将图像作为输入传递给 PICO-8 程序并开始播放。

This got me thinking, wouldn’t it be cool if you could do that for programs on Linux? No! I hear you cry, that’s a dumb idea, but whatever, herein lies an overview of possibly the dumbest things I’ve worked on this year.

这让我想到，如果你能为 Linux 上的程序做到这一点，不是很酷吗？不！我听到你哭了，这是一个愚蠢的想法，但无论如何，这里是我今年工作过的最愚蠢的事情的概述。

## Encoding

## 编码

I'm not entirely sure what PICO-8 is actually doing, but at a guess it's probably use [Steganography](https://en.wikipedia.org/wiki/Steganography) techniques to 'hide' the data within the raw bytes of the image. There are a lot of resources out there that explain how Steganography works, but the crux of it is quite simple, your image your want to hide data into is made up of bytes, an image is made up of pixels. Pixels are made up of 3 Red Green and Blue (RGB) values, represented as 3 bytes. To hide your data (the “payload”) you essentially “mix” the bytes from your payload with the bytes from the image.

我不完全确定 PICO-8 实际上在做什么，但我猜测它可能使用 [Steganography](https://en.wikipedia.org/wiki/Steganography) 技术来“隐藏”原始字节中的数据的图像。有很多资源可以解释隐写术的工作原理，但它的关键非常简单，您想要隐藏数据的图像由字节组成，图像由像素组成。像素由 3 个红绿蓝 (RGB) 值组成，表示为 3 个字节。要隐藏您的数据（“有效负载”)，您实际上是将有效负载中的字节与图像中的字节“混合”。

If you just replaced each byte in your cover image with the bytes from your payload, you would end up with sections of the image looking distorted as the colours probably wouldn’t match with what your original image was. The trick is to be as subtle as possible, or _hide in plain sight_. This can be achieved by _spreading_ your payload bytes over the bytes of the cover image by using the _least significant bits_ to hide them in. In other words, make subtle adjustments to the byte values so the colour changes are not drastic enough to be perceptive by the human eye.

如果您只是用负载中的字节替换封面图像中的每个字节，最终图像的某些部分看起来会失真，因为颜色可能与原始图像不匹配。诀窍是尽可能地微妙，或者_隐藏在视线中_。这可以通过使用_最不重要的位_将它们隐藏在封面图像的字节上_扩展_您的有效负载字节来实现。换句话说，对字节值进行细微调整，使颜色变化不会剧烈到足以被感知人眼。

For example if your payload was the letter `H`, represented as `01001000` in binary (72), and your image contained a series of black pixels

例如，如果您的负载是字母“H”，以二进制 (72) 表示为“01001000”，并且您的图像包含一系列黑色像素

[![](http://djharper.dev/img/byte-replace1.png)](http://djharper.dev/img/byte-replace1.png)

The bits from the input bytes are spread across 8 output bytes by hiding them in the least significant bit

输入字节中的位通过将它们隐藏在最低有效位中而分布在 8 个输出字节中

The output is two-and-a-bit pixels that are slightly less black than before, but can you tell the difference?

输出是两个和一个位的像素，比以前的黑色略少，但你能分辨出区别吗？

[![](http://djharper.dev/img/pixels1.png)](http://djharper.dev/img/pixels1.png)

The pixels have been adjusted in colour slightly.

像素颜色略有调整。

Well, an exceptionally trained colour connoisseur might be able to, but in reality these subtle shifts can really only be noticed by a machine. Retrieving your super secret `H` is just a matter of reading 8 bytes from the resulting image and re-assembling them back into 1 byte. Obviously hiding a single letter is lame, but this can scale to anything you want, a super secret sentence, a copy of _War and Peace_, a link to your soundcloud, the go compiler, the only limit is the amount of bytes available in your cover image as you'll require at least 8x whatever your input is.

好吧，一个训练有素的色彩鉴赏家也许能够做到，但实际上这些微妙的变化只能被机器注意到。检索您的超级秘密“H”只是从结果图像中读取 8 个字节并将它们重新组合成 1 个字节的问题。显然隐藏一个字母是蹩脚的，但这可以扩展到任何你想要的东西，一个超级秘密的句子，一个 _War and Peace_ 的副本，一个指向你的 soundcloud 的链接，go 编译器，唯一的限制是你的可用字节数封面图片，因为无论您的输入是什么，您都需要至少 8 倍。

## Hiding programs

## 隐藏程序

So, back to the whole linux-executables-in-an-image thing, that old chestnut. Well, seeing as executables are just bytes, they can be hidden in images. Just like in the PICO-8 thing.

所以，回到整个 linux-executables-in-an-image 事情，那个老栗子。好吧，由于可执行文件只是字节，它们可以隐藏在图像中。就像在 PICO-8 中一样。

Before I could achieve this I decided to write my own [Steganography library](https://github.com/djhworld/steg) and [tool](https://github.com/djhworld/stegtool) to support encoding and decoding data into PNGs. Yes, there are lots of steganography libraries and tools out there but I learn better by building.

在实现这一目标之前，我决定编写自己的 [Steganography 库](https://github.com/djhworld/steg) 和 [工具](https://github.com/djhworld/stegtool) 来支持编码和解码将数据转换为 PNG。是的，有很多隐写术库和工具，但我通过构建学习得更好。

```bash
$ stegtool encode <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--cover-image htop-logo.png <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--input-data /usr/bin/htop <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--output-image htop.png
$
$ <span style="color:#cb4b16">echo</span> <span style="color:#2aa198">"Super secret hidden message"</span> |stegtool encode <span style="color:#2aa198">\ </span>
--cover-image image.png <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--output-image image-with-hidden-message.png
$ stegtool decode --image image-with-hidden-message.png
Super secret hidden message
```




As it’s all written in [Rust](https://www.rust-lang.org/) it wasn’t that difficult to compile to WASM, so feel free to play with it here:

因为它都是用 [Rust](https://www.rust-lang.org/) 编写的，所以编译成 WASM 并不难，所以请在这里随意使用它：

Anyway, now that can embed data, including executables into an image, how do we run them?

无论如何，现在可以将数据（包括可执行文件）嵌入到图像中，我们如何运行它们？

## Get it running

## 让它运行

The simple option would be to just run the tool above, `decode` the data into a new file, `chmod +x` it and then run it. It works but that’s not fun enough. What I wanted was something similar to the PICO-8 experience, you pass something a PNG image and it takes care of the rest.

简单的选择是运行上面的工具，将数据‘解码’成一个新文件，‘chmod +x’然后运行它。它有效，但这还不够有趣。我想要的是类似于 PICO-8 的体验，你传递一个 PNG 图像，剩下的由它来处理。

However, as it turns out, you can’t just load some arbitrary set of bytes into memory and tell Linux to jump to it. Well, not in a direct way anyway, but you _can_ use some cheap tricks to fudge it.

然而，事实证明，你不能只是将一些任意的字节集加载到内存中，然后告诉 Linux 跳转到它。好吧，无论如何都不是直接的方式，但是您_可以_使用一些便宜的技巧来捏造它。

## memfd\_create

## memfd\_create

After reading [this blogpost](https://magisterquis.github.io/2018/03/31/in-memory-only-elf-execution.html) it became apparent to me you can create an in-memory file and mark it as executable

阅读 [这篇博文](https://magisterquis.github.io/2018/03/31/in-memory-only-elf-execution.html) 后，我很明显可以创建一个内存文件并标记它作为可执行文件

> Wouldn’t it be cool to just grab a chunk of memory, put our binary in there, and run it without monkey-patching the kernel, rewriting execve(2) in userland, or loading a library into another process?

> 只需获取一块内存，将我们的二进制文件放在那里，然后运行它，而不用猴子修补内核、在用户空间重写 execve(2) 或将库加载到另一个进程中，这不是很酷吗？

This method uses the syscall [memfd\_create(2)](https://man7.org/linux/man-pages/man2/memfd_create.2.html) to create a file under the `/proc/self/fd` namespace of your process and load any data you want in it using `write`. I spent quite a while messing around with the [libc](https://crates.io/crates/libc) bindings for Rust to get this to work, and had a lot of trouble understanding the data types you pass around, the documentation for these Rust bindings doesn't help much.

该方法使用系统调用[memfd\_create(2)](https://man7.org/linux/man-pages/man2/memfd_create.2.html)在`/proc/self/fd`下创建一个文件进程的命名空间，并使用 `write` 加载你想要的任何数据。我花了很长时间来处理 Rust 的 [libc](https://crates.io/crates/libc) 绑定以使其正常工作，并且在理解您传递的数据类型、文档方面遇到了很多麻烦对于这些 Rust 绑定并没有多大帮助。

I got something working eventually though

不过我最终得到了一些工作

```rust
<span style="color:#859900">unsafe</span> {
    <span style="color:#859900">let</span> <span style="color:#268bd2">write_mode</span> = <span style="color:#2aa198;font-weight:bold">119</span>;<span style="color:#93a1a1;font-style:italic">// w
</span><span style="color:#93a1a1;font-style:italic"></span>    <span style="color:#93a1a1;font-style:italic">// create executable in-memory file
</span><span style="color:#93a1a1;font-style:italic"></span>    <span style="color:#859900">let</span> <span style="color:#268bd2">fd</span> = <span style="color:#268bd2">syscall</span>(<span style="color:#268bd2">libc</span>::<span style="color:#268bd2">SYS_memfd_create</span>, &<span style="color:#268bd2">write_mode</span>, <span style="color:#2aa198;font-weight:bold">1</span>);
    <span style="color:#859900">if</span> <span style="color:#268bd2">fd</span> == -<span style="color:#2aa198;font-weight:bold">1</span> {
        <span style="color:#859900">return</span> <span style="color:#cb4b16">Err</span>(<span style="color:#cb4b16">String</span>::<span style="color:#268bd2">from</span>(<span style="color:#2aa198">"memfd_create failed"</span>));
    }

    <span style="color:#859900">let</span> <span style="color:#268bd2">file</span> = <span style="color:#268bd2">libc</span>::<span style="color:#268bd2">fdopen</span>(<span style="color:#268bd2">fd</span>, &<span style="color:#268bd2">write_mode</span>);

    <span style="color:#93a1a1;font-style:italic">// write contents of our binary
</span><span style="color:#93a1a1;font-style:italic"></span>    <span style="color:#268bd2">libc</span>::<span style="color:#268bd2">fwrite</span>(
        <span style="color:#268bd2">data</span>.<span style="color:#268bd2">as_ptr</span>() <span style="color:#859900">as</span> *<span style="color:#859900">mut</span> <span style="color:#268bd2">libc</span>::<span style="color:#268bd2">c_void</span>,
        <span style="color:#2aa198;font-weight:bold">8</span> <span style="color:#859900">as</span> <span style="color:#859900;font-weight:bold">usize</span>,
        <span style="color:#268bd2">data</span>.<span style="color:#268bd2">len</span>() <span style="color:#859900">as</span> <span style="color:#859900;font-weight:bold">usize</span>,
        <span style="color:#268bd2">file</span>,
    );
}

```


Invoking `/proc/self/fd/<fd>` as a child process from the parent that created it is enough to run your binary.

从创建它的父进程调用`/proc/self/fd/<fd>` 作为子进程就足以运行你的二进制文件。

```rust
<span style="color:#859900">let</span> <span style="color:#268bd2">output</span> = <span style="color:#268bd2">Command</span>::<span style="color:#268bd2">new</span>(<span style="color:#268bd2">format</span>!(<span style="color:#2aa198">"/proc/self/fd/{}"</span>, <span style="color:#268bd2">fd</span>))
    .<span style="color:#268bd2">args</span>(<span style="color:#268bd2">args</span>)
    .<span style="color:#268bd2">stdin</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">stdout</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">stderr</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">spawn</span>();

```




Given these building blocks, I wrote [pngrun](https://github.com/djhworld/pngrun) to run the images. It essentially…

鉴于这些构建块，我编写了 [pngrun](https://github.com/djhworld/pngrun) 来运行图像。它本质上…

1. Accepts an image that has had our binary embedded in it from the steganography tool, and any arguments
2. Decodes it (i.e. extracts and re-assembles the bytes)
3. Creates an in-memory file using`memfd_create`
4. Puts the bytes of the binary into the in-memory file
5. Invokes the file`/proc/self/fd/<fd>` as a child process, passing any arguments from the parent

1. 接受来自隐写术工具的嵌入了我们的二进制文件的图像，以及任何参数
2. 解码（即提取并重新组合字节）
3.使用`memfd_create`创建内存文件
4. 将二进制的字节放入内存文件
5. 调用文件`/proc/self/fd/<fd>` 作为子进程，从父进程传递任何参数

So you can run it like this

所以你可以像这样运行它

```bash
$ pngrun htop.png
<htop output>
$ pngrun go.png run main.go
Hello world!
```


Once `pngrun` exits the in-memory file is destroyed.

一旦 `pngrun` 退出，内存中的文件就会被销毁。

## binfmt\_misc

## binfmt\_misc

It's annoying having to type `pngrun` every time though, so my last cheap trick to this pointless gimmick was to use [binfmt\_misc](https://en.wikipedia.org/wiki/Binfmt_misc), a system that allows you to “execute” files based on its file types. I think it was mainly designed for interpreters/virtual machines, like Java. So instead of typing `java -jar my-jar.jar` you can just type `./my-jar.jar` and it will invoke the `java` process to run your JAR. The caveat is your file `my-jar.jar` needs to be marked as executable first.

但是每次都必须输入`pngrun`很烦人，所以我对这个毫无意义的噱头的最后一个廉价技巧是使用[binfmt\_misc](https://en.wikipedia.org/wiki/Binfmt_misc)，一个允许你根据文件类型“执行”文件。我认为它主要是为解释器/虚拟机设计的，比如 Java。因此，您无需键入“java -jar my-jar.jar”，而只需键入“./my-jar.jar”，它就会调用“java”进程来运行您的 JAR。需要注意的是，您的文件 `my-jar.jar` 需要首先标记为可执行文件。

So adding an entry to binfmt\_misc for `pngrun` to attempt to run any `png` files that have the `x` flag set was as simple as

因此，在 binfmt\_misc 中为 `pngrun` 添加一个条目以尝试运行任何设置了 `x` 标志的 `png` 文件就像这样简单

```bash
$ cat /etc/binfmt.d/pngrun.conf
:ExecutablePNG:E::png::/home/me/bin/pngrun:
$ sudo systemctl restart binfmt.d
$ chmod +x htop.png
$ ./htop.png
<output>
```


## What’s the point

##  重点是什么

Well, there isn’t one really. I was seduced by the idea of making PNG images run programs and got a bit carried away with it, but it was fun none the less. There’s something amusing to me about distributing programs as an image, remember the ridiculous cardboard boxes PC software used to come in with artwork on the front, why not bring that back! (lets not)

嗯，真的没有。我被让 PNG 图像运行程序的想法所吸引，并有点忘乎所以，但它仍然很有趣。将程序作为图像分发对我来说很有趣，请记住用于在前面带有艺术品的可笑的纸板箱 PC 软件，为什么不把它带回来！ （不要吧）

It’s really dumb though and comes with a lot of caveats that make it completely pointless and impractical, the main one being needing the stupid `pngrun` program on your machine. But I also noticed some weird stuff around programs like `clang`. I encoded it into this fun LLVM logo and while it runs OK, it fails when you try to compile something.

尽管如此，它真的很愚蠢，并带有很多警告，使其完全没有意义和不切实际，主要是在您的机器上需要愚蠢的“pngrun”程序。但我也注意到一些关于程序的奇怪东西，比如 `clang`。我把它编码成这个有趣的 LLVM 标志，虽然它运行正常，但当你尝试编译某些东西时它会失败。

[![](http://djharper.dev/img/DragonMedium.png)](http://djharper.dev/img/DragonMedium.png)

```bash
$ ./clang.png --version
clang version <span style="color:#2aa198;font-weight:bold">11</span>.0.0 (Fedora <span style="color:#2aa198;font-weight:bold">11</span>.0.0-2.fc33)
Target: x86_64-unknown-linux-gnu
Thread model: posix
InstalledDir: /proc/self/fd
$ ./clang.png main.c
error: unable to execute command: Executable <span style="color:#2aa198">""</span> doesn't exist!
```


This is probably a product of the anonymous file thing, which can probably be overcome if I could be bothered to investigate.

这可能是匿名文件的产物，如果我能费心去调查，这可能可以克服。

### Additional reasons why this is dumb

### 这很愚蠢的其他原因

A lot of binaries are quite large, and given the constraints of needing to fit them into an image, sometimes these need to be _big_, meaning you end up with comically large files.

许多二进制文件都非常大，并且考虑到需要将它们放入图像中的限制，有时这些文件需要 _big_，这意味着您最终会得到可笑的大文件。

Also most software isn’t just one executable so the dream of just distributing a PNG kinda falls flat for more complex software like games.

此外，大多数软件不仅仅是一个可执行文件，因此为更复杂的软件（如游戏）分发 PNG 的梦想有点落空。

## Conclusion

##  结论

This is probably the dumbest project I’ve worked on all year but it’s been fun, I’ve learned about Steganography, `memfd_create`, `binfmt_misc` and played a little more with Rust.

这可能是我一整年从事的最愚蠢的项目，但它很有趣，我了解了隐写术、`memfd_create`、`binfmt_misc` 并更多地使用了 Rust。

* * *

* * *

Load comments 

加载评论

