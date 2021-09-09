# [Executable PNGs](https://djharper.dev/post/2020/12/26/executable-pngs/ "Executable PNGs")

Saturday, December 26, 2020

It's an image _and_ a program

A few weeks ago I was reading about [PICO-8](https://www.lexaloffle.com/pico-8.php), a fantasy games console with limited constraints. What really piqued my interest about it was the novel way games are distributed, you encode them into a PNG image. This includes the game code, assets, everything. The image can be whatever you want, screenshots from the game, cool artwork or just text. To load them you pass the image as input to the PICO-8 program and start playing.

This got me thinking, wouldn’t it be cool if you could do that for programs on Linux? No! I hear you cry, that’s a dumb idea, but whatever, herein lies an overview of possibly the dumbest things I’ve worked on this year.

## Encoding

I’m not entirely sure what PICO-8 is actually doing, but at a guess it’s probably use [Steganography](https://en.wikipedia.org/wiki/Steganography) techniques to ‘hide’ the data within the raw bytes of the image. There are a lot of resources out there that explain how Steganography works, but the crux of it is quite simple, your image your want to hide data into is made up of bytes, an image is made up of pixels. Pixels are made up of 3 Red Green and Blue (RGB) values, represented as 3 bytes. To hide your data (the “payload”) you essentially “mix” the bytes from your payload with the bytes from the image.

If you just replaced each byte in your cover image with the bytes from your payload, you would end up with sections of the image looking distorted as the colours probably wouldn’t match with what your original image was. The trick is to be as subtle as possible, or _hide in plain sight_. This can be achieved by _spreading_ your payload bytes over the bytes of the cover image by using the _least significant bits_ to hide them in. In other words, make subtle adjustments to the byte values so the colour changes are not drastic enough to be perceptive by the human eye.

For example if your payload was the letter `H`, represented as `01001000` in binary (72), and your image contained a series of black pixels

[![](http://djharper.dev/img/byte-replace1.png)](http://djharper.dev/img/byte-replace1.png)

The bits from the input bytes are spread across 8 output bytes by hiding them in the least significant bit

The output is two-and-a-bit pixels that are slightly less black than before, but can you tell the difference?

[![](http://djharper.dev/img/pixels1.png)](http://djharper.dev/img/pixels1.png)

The pixels have been adjusted in colour slightly.

Well, an exceptionally trained colour connoisseur might be able to, but in reality these subtle shifts can really only be noticed by a machine. Retrieving your super secret `H` is just a matter of reading 8 bytes from the resulting image and re-assembling them back into 1 byte. Obviously hiding a single letter is lame, but this can scale to anything you want, a super secret sentence, a copy of _War and Peace_, a link to your soundcloud, the go compiler, the only limit is the amount of bytes available in your cover image as you’ll require at least 8x whatever your input is.

## Hiding programs

So, back to the whole linux-executables-in-an-image thing, that old chestnut. Well, seeing as executables are just bytes, they can be hidden in images. Just like in the PICO-8 thing.

Before I could achieve this I decided to write my own [Steganography library](https://github.com/djhworld/steg) and [tool](https://github.com/djhworld/stegtool) to support encoding and decoding data into PNGs. Yes, there are lots of steganography libraries and tools out there but I learn better by building.

```bash
$ stegtool encode <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--cover-image htop-logo.png <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--input-data /usr/bin/htop <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--output-image htop.png
$
$ <span style="color:#cb4b16">echo</span> <span style="color:#2aa198">"Super secret hidden message"</span> | stegtool encode <span style="color:#2aa198">\ </span>
--cover-image image.png <span style="color:#2aa198">\
</span><span style="color:#2aa198"></span>--output-image image-with-hidden-message.png
$ stegtool decode --image image-with-hidden-message.png
Super secret hidden message
```

As it’s all written in [Rust](https://www.rust-lang.org/) it wasn’t that difficult to compile to WASM, so feel free to play with it here:

Anyway, now that can embed data, including executables into an image, how do we run them?

## Get it running

The simple option would be to just run the tool above, `decode` the data into a new file, `chmod +x` it and then run it. It works but that’s not fun enough. What I wanted was something similar to the PICO-8 experience, you pass something a PNG image and it takes care of the rest.

However, as it turns out, you can’t just load some arbitrary set of bytes into memory and tell Linux to jump to it. Well, not in a direct way anyway, but you _can_ use some cheap tricks to fudge it.

## memfd\_create

After reading [this blogpost](https://magisterquis.github.io/2018/03/31/in-memory-only-elf-execution.html) it became apparent to me you can create an in-memory file and mark it as executable

> Wouldn’t it be cool to just grab a chunk of memory, put our binary in there, and run it without monkey-patching the kernel, rewriting execve(2) in userland, or loading a library into another process?

This method uses the syscall [memfd\_create(2)](https://man7.org/linux/man-pages/man2/memfd_create.2.html) to create a file under the `/proc/self/fd` namespace of your process and load any data you want in it using `write`. I spent quite a while messing around with the [libc](https://crates.io/crates/libc) bindings for Rust to get this to work, and had a lot of trouble understanding the data types you pass around, the documentation for these Rust bindings doesn’t help much.

I got something working eventually though

```rust
<span style="color:#859900">unsafe</span> {
    <span style="color:#859900">let</span> <span style="color:#268bd2">write_mode</span> = <span style="color:#2aa198;font-weight:bold">119</span>; <span style="color:#93a1a1;font-style:italic">// w
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

```rust
<span style="color:#859900">let</span> <span style="color:#268bd2">output</span> = <span style="color:#268bd2">Command</span>::<span style="color:#268bd2">new</span>(<span style="color:#268bd2">format</span>!(<span style="color:#2aa198">"/proc/self/fd/{}"</span>, <span style="color:#268bd2">fd</span>))
    .<span style="color:#268bd2">args</span>(<span style="color:#268bd2">args</span>)
    .<span style="color:#268bd2">stdin</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">stdout</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">stderr</span>(<span style="color:#268bd2">std</span>::<span style="color:#268bd2">process</span>::<span style="color:#268bd2">Stdio</span>::<span style="color:#268bd2">inherit</span>())
    .<span style="color:#268bd2">spawn</span>();

```

Given these building blocks, I wrote [pngrun](https://github.com/djhworld/pngrun) to run the images. It essentially…

1. Accepts an image that has had our binary embedded in it from the steganography tool, and any arguments
2. Decodes it (i.e. extracts and re-assembles the bytes)
3. Creates an in-memory file using`memfd_create`
4. Puts the bytes of the binary into the in-memory file
5. Invokes the file`/proc/self/fd/<fd>` as a child process, passing any arguments from the parent

So you can run it like this

```bash
$ pngrun htop.png
<htop output>
$ pngrun go.png run main.go
Hello world!
```

Once `pngrun` exits the in-memory file is destroyed.

## binfmt\_misc

It’s annoying having to type `pngrun` every time though, so my last cheap trick to this pointless gimmick was to use [binfmt\_misc](https://en.wikipedia.org/wiki/Binfmt_misc), a system that allows you to “execute” files based on its file types. I think it was mainly designed for interpreters/virtual machines, like Java. So instead of typing `java -jar my-jar.jar` you can just type `./my-jar.jar` and it will invoke the `java` process to run your JAR. The caveat is your file `my-jar.jar` needs to be marked as executable first.

So adding an entry to binfmt\_misc for `pngrun` to attempt to run any `png` files that have the `x` flag set was as simple as

```bash
$ cat /etc/binfmt.d/pngrun.conf
:ExecutablePNG:E::png::/home/me/bin/pngrun:
$ sudo systemctl restart binfmt.d
$ chmod +x htop.png
$ ./htop.png
<output>
```

## What’s the point

Well, there isn’t one really. I was seduced by the idea of making PNG images run programs and got a bit carried away with it, but it was fun none the less. There’s something amusing to me about distributing programs as an image, remember the ridiculous cardboard boxes PC software used to come in with artwork on the front, why not bring that back! (lets not)

It’s really dumb though and comes with a lot of caveats that make it completely pointless and impractical, the main one being needing the stupid `pngrun` program on your machine. But I also noticed some weird stuff around programs like `clang`. I encoded it into this fun LLVM logo and while it runs OK, it fails when you try to compile something.

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

### Additional reasons why this is dumb

A lot of binaries are quite large, and given the constraints of needing to fit them into an image, sometimes these need to be _big_, meaning you end up with comically large files.

Also most software isn’t just one executable so the dream of just distributing a PNG kinda falls flat for more complex software like games.

## Conclusion

This is probably the dumbest project I’ve worked on all year but it’s been fun, I’ve learned about Steganography, `memfd_create`, `binfmt_misc` and played a little more with Rust.

* * *

Load comments

