# Shell Field Guide

# 壳牌现场指南

## Table of Contents

##  目录

- [1. Introduction](https://raimonster.com/scripting-field-guide/#org6565fc6)
- [2. Which Shell?](https://raimonster.com/scripting-field-guide/#org151e4c5)
- [3. Level](https://raimonster.com/scripting-field-guide/#org9666cf2)
- [4. Patterns](https://raimonster.com/scripting-field-guide/#orgfbf20c8)
- [5. Interactive](https://raimonster.com/scripting-field-guide/#orga5110ad)
- [6. Debugging](https://raimonster.com/scripting-field-guide/#org0b80cda)
- [7. zsh-only]()
   - [7.1. Word spliting](https://raimonster.com/scripting-field-guide/#org6532368)
   - [7.2. globbing](https://raimonster.com/scripting-field-guide/#org37fe325)
   - [7.3. Some global aliases:](https://raimonster.com/scripting-field-guide/#org7a15bc1)
   - [7.4. Expansion of global aliases](https://raimonster.com/scripting-field-guide/#org5bac793)
   - [7.5. Autocomplete](https://raimonster.com/scripting-field-guide/#orga27a037)
   - [7.6. Create helpers and generate global aliases automagically](https://raimonster.com/scripting-field-guide/#orgc8509ba)
   - [7.7. suffix aliases don't have to match a filename](https://raimonster.com/scripting-field-guide/#org2269651)
   - [7.8. noglob](https://raimonster.com/scripting-field-guide/#org379b5d5)
   - [7.9. make noglob 'transparent' to bash](https://raimonster.com/scripting-field-guide/#org17a1f68)
   - [7.10. glob nested expansion](https://raimonster.com/scripting-field-guide/#orgfc2c1d9)
   - [7.11. Some extra shortcuts for nice things](https://raimonster.com/scripting-field-guide/#org2e56e92)
   - [7.12. =()](https://raimonster.com/scripting-field-guide/#org15bfae2)
- [8. TODO patterns](https://raimonster.com/scripting-field-guide/#org282d509)
- [9. links](https://raimonster.com/scripting-field-guide/#org9179c42)
- [10. From shell to lisp and everything in between](https://raimonster.com/scripting-field-guide/#orgf35cb8f)
- [11. Credits](https://raimonster.com/scripting-field-guide/#org6d9e75d)

- [1.介绍](https://raimonster.com/scripting-field-guide/#org6565fc6)
- [2.哪个壳？](https://raimonster.com/scripting-field-guide/#org151e4c5)
- [3.级别](https://raimonster.com/scripting-field-guide/#org9666cf2)
- [4.模式](https://raimonster.com/scripting-field-guide/#orgfbf20c8)
- [5.互动](https://raimonster.com/scripting-field-guide/#orga5110ad)
- [6.调试](https://raimonster.com/scripting-field-guide/#org0b80cda)
- [7. zsh-only]()
  - [7.1.分词](https://raimonster.com/scripting-field-guide/#org6532368)
  - [7.2.通配](https://raimonster.com/scripting-field-guide/#org37fe325)
  - [7.3.一些全局别名:](https://raimonster.com/scripting-field-guide/#org7a15bc1)
  - [7.4.全局别名的扩展](https://raimonster.com/scripting-field-guide/#org5bac793)
  - [7.5.自动完成](https://raimonster.com/scripting-field-guide/#orga27a037)
  - [7.6.创建助手并自动生成全局别名](https://raimonster.com/scripting-field-guide/#orgc8509ba)
  - [7.7.后缀别名不必与文件名匹配](https://raimonster.com/scripting-field-guide/#org2269651)
  - [7.8. noglob](https://raimonster.com/scripting-field-guide/#org379b5d5)
  - [7.9.使 noglob '透明' 到 bash](https://raimonster.com/scripting-field-guide/#org17a1f68)
  - [7.10. glob 嵌套扩展](https://raimonster.com/scripting-field-guide/#orgfc2c1d9)
  - [7.11.一些额外的好东西的捷径](https://raimonster.com/scripting-field-guide/#org2e56e92)
  - [7.12. =()](https://raimonster.com/scripting-field-guide/#org15bfae2)
- [8. TODO 模式](https://raimonster.com/scripting-field-guide/#org282d509)
- [9.链接](https://raimonster.com/scripting-field-guide/#org9179c42)
- [10.从 shell 到 lisp 以及介于两者之间的所有内容](https://raimonster.com/scripting-field-guide/#orgf35cb8f)
- [11.积分](https://raimonster.com/scripting-field-guide/#org6d9e75d)

## 1 Introduction

## 1 介绍

This booklet is intended to be a catalog of tricks and techniques you may want to use if you're doing some sort of complex scripting. Some are just useful, some are more playful, and might not have such direct impact in your day-to-day life. Some are pure entertainment. You'll have to judge by yourself which things belong to which category. I'll try to keep the rhetoric to the minimum to maximize signal/noise.

如果您正在编写某种复杂的脚本，这本小册子旨在成为您可能想要使用的技巧和技术的目录。有些只是有用，有些则更有趣，并且可能不会对您的日常生活产生如此直接的影响。有些是纯粹的娱乐。你必须自己判断哪些东西属于哪个类别。我会尽量减少言辞，以最大限度地提高信号/噪音。

The git repo is at https://github.com/kidd/scripting-field-guide/. Any feedback is greatly appreciated. Keep in mind this is not any kind of official doc. I just write MY current "state of the art" and I'll be updating the contents with useful stuff I find or discover, that are not widely explained in usual manuals/wikis.

git 存储库位于 https://github.com/kidd/scripting-field-guide/。非常感谢任何反馈。请记住，这不是任何类型的官方文档。我只是写我目前的“最先进的技术”，我会用我发现或发现的有用的东西来更新内容，这些东西在通常的手册/维基中没有得到广泛的解释。

You probably have some amount of sh/bash/zsh in your stack that probably started as one-off scripts, and probably later on started growing and being copypasted everywhere in your pipelines, or your coworkers use for their own use (with some variations) , etc. Those scripts are very difficult to kill and they have a very high mutation rate.

您的堆栈中可能有一些 sh/bash/zsh 可能是作为一次性脚本开始的，并且可能后来开始增长并在您的管道中随处复制粘贴，或者您的同事自用（有一些变化）等。这些脚本很难被杀死，而且它们的变异率非常高。

## 2 Which Shell?

##2 哪个壳？

No matter if you use Linux, Mac, or Windows, you should be living most of the time in a shell to enjoy the content shown here. Some value comes from the automated scripts, and some comes from the daily usage and refinement of your helper functions, aliases, etc. in interactive mode.

无论您使用 Linux、Mac 还是 Windows，您都应该大部分时间都在 shell 中享受这里显示的内容。一些价值来自自动化脚本，一些来自交互模式下你的辅助函数、别名等的日常使用和改进。

In general the examples here are meant to run in Bash or Zsh, which are compatible for the most part.

通常，这里的示例旨在在 Bash 或 Zsh 中运行，它们在大多数情况下是兼容的。

## 3 Level

## 3 级

These examples are based on non-trivial real world code I've written using patterns I haven't seen applied in many places over the net. A few of the snippets are stolen from public repos I find interesting. Also, important scripting stuff might be missing if I don't feel I have anything to add to the generally available info around.

这些示例基于我使用我在网上很多地方都没有见过的模式编写的重要的现实世界代码。一些片段是从我觉得有趣的公共存储库中窃取的。此外，如果我觉得我没有任何东西可以添加到普遍可用的信息中，那么重要的脚本内容可能会丢失。

## 4 Patterns

## 4 种模式

### 4.1 Use Shellcheck

### 4.1 使用 Shellcheck

First, let's get that out of the way. This is low-hanging fruit. And you will get the most of this booklet by following it.

首先，让我们把它排除在外。这是唾手可得的果实。遵循它，您将充分利用这本小册子。

A lot of the most common errors we usually make are well known ones. And in fact, we all usually fail in similar ways. Bash is known for being error prone when dealing with testing variable values, string operations, or flaky subshells and pipes.

我们通常犯的许多最常见的错误都是众所周知的错误。事实上，我们通常都会以类似的方式失败。 Bash 以在处理测试变量值、字符串操作或片状子外壳和管道时容易出错而闻名。

Installing [shellcheck](https://www.shellcheck.net/) will flag many of those ticking bombs for you. 

安装 [shellcheck](https://www.shellcheck.net/) 将为您标记许多这些滴答作响的炸弹。

No matter which editor you are using, you should be able to install a plugin to do automatic checks while you're editing.

无论您使用哪种编辑器，您都应该能够在编辑时安装插件以进行自动检查。

In emacs' case, the package is called [flymake-shellcheck](https://github.com/federicotdn/flymake-shellcheck), and a quick configuration for it is:

在 emacs 的例子中，这个包叫做 [flymake-shellcheck](https://github.com/federicotdn/flymake-shellcheck)，它的快速配置是：

```
(use-package flymake-shellcheck
  :ensure t
  :commands flymake-shellcheck-load
  :init
  (add-hook 'sh-mode-hook 'flymake-shellcheck-load))
```

Shellcheck is available on most distros, so it's just an `apt`, `brew`, or `nix-env` away.

Shellcheck 在大多数发行版上都可用，所以它只是一个 `apt`、`brew` 或 `nix-env`。

### 4.2 Overview

### 4.2 概述

In this section, we're covering the parts of the basics that are not so basic after all, or that are more unique in shellscripting languages.

在本节中，我们将介绍一些根本不那么基本的基础知识，或者在 shell 脚本语言中更独特的部分。

### 4.3 Booleans and Conditionals

### 4.3 布尔和条件

In any shell, `foo && bar` will execute `bar` only if `foo` succeeded. That means that `foo` returned 0. That means that to && (which you read like "and"), 0 is true. so yes. 0 is true, and other values are false.

在任何 shell 中，只有当 `foo` 成功时，`foo && bar` 才会执行 `bar`。这意味着`foo`返回0。这意味着对于&&（你读为“and”），0是真的。所以是的。 0 为真，其他值为假。

### 4.4 Arrays

### 4.4 数组

Ordered list of things.

有序的事物清单。

```
foo=("ls" "/tmp/")

echo ${foo[-2]}
echo ${foo[-1]}
echo ${foo[0]}
echo ${foo[1]}
echo ${foo[2]}

for i in "${foo[@]}";do
  echo $i
done

$foo
${foo[*]}
${foo[@]}
echo ${#foo[*]}
echo ${#foo[@]}
```

Are `*` and `@` equal? [nope](https://stackoverflow.com/questions/2761723/what-is-the-difference-between-and-in-shell-scripts).

`*` 和 `@` 相等吗？ [不](https://stackoverflow.com/questions/2761723/what-is-the-difference-between-and-in-shell-scripts)。

```
"${foo[@]}"
"${foo[*]}"
#!/bin/bash

main()
{
  echo 'MAIN sees ' $# ' args'
}

main $*
main $@

main "$*"
main "$@"

### end ###

and I run it like this:

my_script 'a b c' d e
```

### 4.5 Pass Arrays around

### 4.5 传递数组

```
a=('Track 01.mp3' 'Track 02.mp3')
myfun "${a[@]}" # pass array to a function
b=( "${a[@]}" ) # copy array
```

Read the great Oil Shell [blogpost](https://www.oilshell.org/blog/2016/11/06.html).

阅读伟大的石油壳牌 [博客文章](https://www.oilshell.org/blog/2016/11/06.html)。

### 4.6 Slurping arrays

### 4.6 Slurping 数组

A nice way to read a bunch of elements in one go is to use `readarray`.

一次性读取一堆元素的一种好方法是使用“readarray”。

```
parse_args() {
  [[ $# -eq 0 ]] && die "Usage: $0 <version>"
  version="$1"
  local version_split=$(echo $version | tr '.' '\n')
  readarray -t version_array <<< "$version_split"

  if [[ -z ${version_array[3]} ]];then
    die "not enough version numbers"
  fi
}
```

Even nicer would be to use IFS so we'd be able to split in one go.

更好的是使用 IFS，这样我们就可以一次性拆分。

```
IFS=.read -a ver <<<"1.23.1.0"
echo ${ver[0]}
echo "next/${ver[0]}.${ver[1]}.x.x"
```

Or, use it in a destructuring fashion:

或者，以解构方式使用它：

```
get_nix_version_parts(){
  local major minor patch
  # shellcheck disable=SC2034,SC2162
  IFS="."read major minor patch < <(get_nix_version)
  local -p
}

$ get_nix_version_parts
major=2
minor=3
patch=4
```

https://news.ycombinator.com/item?id=24408318

https://news.ycombinator.com/item?id=24408318

### 4.7 Read lines from file

### 4.7 从文件中读取行

The `read` command we used just above is part of the usual idiom to read a file line by line.

我们上面使用的“read”命令是逐行读取文件的常用习惯用法的一部分。

```
while read line;do
  echo $line
done < /tmp/file.txt
```

More related info in the [BashFAQ001](https://mywiki.wooledge.org/BashFAQ/001). But it's very rare the case where I need to iterate a file line by line.

[BashFAQ001](https://mywiki.wooledge.org/BashFAQ/001) 中有更多相关信息。但我需要逐行迭代文件的情况非常罕见。

### 4.8 Assign to $1,$2,$3…

### 4.8 分配给 $1,$2,$3...

`set --` can be used as an incantation to assign to the positional parameters. Let me show you.

`set --` 可以用作分配给位置参数的咒语。让我演示给你看。

```
set -- a b c
echo $1 $2 $3
echo $@
```

See? here's how to "unshift" a parameter to the current arg list:

看？以下是将参数“取消移位”到当前 arg 列表的方法：

```
set -- "injected" "$@"
```

### 4.9 Functions

### 4.9 函数

Functions are functions. They receive arguments, and they return a value.

函数就是函数。它们接收参数，并返回一个值。

The special thing about shell functions is that they also can use the file descriptors of the process. That means that they "inherit" STDIN, STDOUT, STDERR (and maybe more).

shell 函数的特殊之处在于它们还可以使用进程的文件描述符。这意味着他们“继承”了 STDIN、STDOUT、STDERR（或许还有更多）。

Use them.

使用它们。

Another point is that function names can be passed as parameters, because they are passed as strings, but you can call them inside as functions again.

还有一点是函数名可以作为参数传递，因为它们是作为字符串传递的，但是你可以在内部再次作为函数调用它们。

```
f() {
  $1 hi
}

f echo
f touch # will create a file 'hi'
```

### 4.10 Variables

### 4.10 变量

By default variables are global, to a file. No matter if you assign them for the first time inside a function, or at the top level.

默认情况下，变量是全局的，对于文件。无论您是第一次在函数内部分配它们，还是在顶层分配它们。

```
foo=3
bar=$foo
f() {
  echo $bar
}
f
f() {
  bar=1
}
f
echo $bar
```

You make a variable local to a function with `local`. Use it as much as you can (kinda obvious).

您可以使用“local”为函数创建一个局部变量。尽可能多地使用它（有点明显）。

```
myfun() {
  local bar
  bar=3
  echo $bar
}

bar=4
echo $bar
myfun
echo $bar
```

### 4.11 Variable Expansions

### 4.11 变量扩展

They offer some variable manipulations using shell only, not having to create another process `sed,awk,perl`.

它们仅使用 shell 提供了一些变量操作，而不必创建另一个进程 `sed,awk,perl`。

```
v=banana
# substitute one
echo ${v/na/NA}   # baNAna
# substitute many
echo ${v//na/NA}  # baNANA

# substitute from the start (think ^ in PCRE)
echo ${v/#ba/NA}  # NAnana

# substitute from the end
echo ${v/%na/NA}  # banaNA
```

Take a read on https://tldp.org/LDP/abs/html/manipulatingvars.html and https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html for more details .

阅读 https://tldp.org/LDP/abs/html/manipulatingvars.html 和 https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html 了解更多详情.

And a nice non-obvious trick from here is to prefix or suffix a variable string:

这里有一个很好的非显而易见的技巧是为变量字符串添加前缀或后缀：

```
v=banana
echo ${v/%/na}   # bananana
echo ${v/#/na}   # nabanana
```

And a less obvious trick is to prefix every element of an array with a fixed string:

还有一个不太明显的技巧是在数组的每个元素前面加上一个固定的字符串：

```
local arr=(var1=1 var2=2)
echo ${arr[*]/#/"--env "}
```

This will produce `--env var1=1 --env var2=2`. Super useful to be combined when building flags for docker.

这将产生`--env var1=1 --env var2=2`。在为 docker 构建标志时组合起来非常有用。

### 4.12 Interpolation

### 4.12 插值

We previously saw that functions can be passed around as strings, and be called later on.

我们之前看到函数可以作为字符串传递，并在以后调用。

Something that might not be obvious is that the string can be created from shorter strings, and that allows for an extra flexibility, that comes with its own dangers, but it's a very useful pattern to dispatch functions based on user input or function outputs.

可能不明显的是，字符串可以从较短的字符串创建，这允许额外的灵活性，这有其自身的危险，但它是基于用户输入或函数输出调度函数的非常有用的模式。

```
l=l
s=s
$l$s .
```

### 4.13 dispatch functions using args

### 4.13 使用 args 的调度函数

A nice usage of the previous technique is using user input as a dispatching method.

先前技术的一个很好的用法是使用用户输入作为调度方法。

You've probably seen this pattern already:

您可能已经见过这种模式：

```
while [[ $# -gt 0 ]];do

case $1 in
  foo)
    foo
    ;;
  *)
    exit 1
    ;;
esac
shift
done
```

And it is useful for its own good, and flexible.

并且它是有用的，而且它很灵活。

But for some simpler cases, we can dispatch based on the variable itself:

但是对于一些更简单的情况，我们可以根据变量本身进行调度：

```
cmd_foo() {
 do-something
}

cmd_$1
```

The problem with this is that in case we supply a `$1` that doesn't map to any `cmd_$1` we'll get something like

这样做的问题是，如果我们提供一个不映射到任何 cmd_$1 的 `$1`，我们会得到类似的东西

```
bash: cmd_notexisting: command not found
```

### 4.14 command_not_found_handle

### 4.14 command_not_found_handle

Here's a detail on a kinda obscure bash (only bash) feature.

这是一个有点晦涩的 bash（仅限 bash）功能的详细信息。

You can set a hook that will be called when bash tries to run a command and it doesn't find it.

您可以设置一个钩子，当 bash 尝试运行命令但找不到它时，将调用该钩子。

```
command_not_found_handle() {
  echo "$1 is not a correct command. Cmds allowed:"
  echo "$(typeset -F | grep cmd_ | sed -e 's/.*cmd_/cmd_/')"
}

cmd_foo() {
  echo "foo"
}

cmd_baz() {
  echo "baz"
}
cmd_bar
```

You can unset the function `command_not_found_handle` to go back to the normal behavior.

您可以取消设置函数 `command_not_found_handle` 以恢复正常行为。

### 4.15 Return Values for Conditionals

### 4.15 条件的返回值

`if` 's test condition can use the return values of commands. That's a known thing, but lots of code you see around relies on `[[]]` to test the return values of commands/functions anyway.

`if` 的测试条件可以使用命令的返回值。这是众所周知的事情，但是你看到的很多代码都依赖于 `[[]]` 来测试命令/函数的返回值。

```
if echo "foo" |grep "bar" ;then
  echo "found!"
fi
```

This is much clearer than

这比

```
if [[ !-z $( echo "foo" | grep "bar") ]];then
  echo "found!"
fi
```

As easy and trivial as it seems, this way of thinking pushes you forward to thinking about creating smaller functions that check the conditions and `return` 0 or non 0. It's syntactically smaller, and usually makes you play by the rules of the commands, more than just finding your way around the output strings.

尽管看起来简单而微不足道，但这种思维方式会促使您考虑创建更小的函数来检查条件并“返回”0 或非 0。它在语法上更小，通常会让您遵守命令的规则，不仅仅是找到解决输出字符串的方法。

```
if less_than $package "1.3.2";then
  die "can't proceed"
fi
```

### 4.16 set variable in an "if" test

### 4.16 在“if”测试中设置变量

Usual pattern to capture the output of a command and branch depending on its return value is:

根据返回值捕获命令和分支输出的常用模式是：

```
res="$(... whatever ...)"
if [ "$?"-eq 0 ];then ...
                        ...
fi
```

Well, you can test the return value AND capture the output at the same time!

好吧，您可以同时测试返回值并捕获输出！

```
if res="$(...)";then
  ...
fi
```

Unfortunately, it doesnt' work with `local`, so you can't be defining a local var in the same line. So, the variable is either global, or you spent a line to declare it local before. Still, I think I prefer to have a line to declare the variable as local rather than having explicit `$?`'s around.

不幸的是，它不适用于“local”，因此您不能在同一行中定义本地变量。所以，变量要么是全局的，要么你之前用一行来声明它是局部的。不过，我认为我更喜欢用一行来将变量声明为本地变量，而不是使用显式的“$?”。

```
local var1
if var1=$(f);then
  echo "$var1"
fi
```

Ref: https://news.ycombinator.com/item?id=27163494

参考：https://news.ycombinator.com/item?id=27163494

### 4.17 Do work on loop conditions

### 4.17 处理循环条件

I've not seen it used a lot (and there might be a reason for it, who knows), `while` conditions are just plain commands, so you can put other stuff than `[]/[[]]/test` there.

我没有看到它经常使用（谁知道可能是有原因的），`while` 条件只是简单的命令，所以你可以放置除 `[]/[[]]/test` 之外的其他东西那里。

Heres's an idiomatic way to iterate through all the arguments of a function while consuming the `$*` array.

这是在使用 `$*` 数组时遍历函数的所有参数的惯用方法。

```
while(($#)) ;do
  #...
  shift
done
```

And here's a pseudo-repl that keeps shooting one-off commands. This will keep shooting `tr` commands to whatever strings you give it, with the usual rlwrap goodies.

这是一个伪复制，它不断发出一次性命令。这将继续向您提供的任何字符串发送 `tr` 命令，以及通常的 rlwrap 好东西。

```
while rlwrap -o -S'>> ' tr a-z A-Z ;do :;done
```

Note: `:` is a nop builtin in bash.

注意：`:` 是 bash 内建的 nop。

### 4.18 One Branch Conditionals

### 4.18 单分支条件

The usual conditionals one sees everywhere look like `if`.

随处可见的常见条件看起来像“if”。

```
if [[ some-condition ]];then
  echo "yes"
fi
```

This is all good and fine, but in the same vein of using the least powerful construct for each task, it's nice to think of the one way conditionals in the form of `&&` and `||` as a way to explicitly say that we don't want to do anything else when the condition is not met. It's a hint to the reader.

这一切都很好，但与为每个任务使用最不强大的构造一样，很高兴想到以 `&&` 和 `||` 形式的单向条件作为一种明确表示的方式当条件不满足时，我们不想做任何其他事情。这是对读者的提示。

```
some-condition ||{
   echo "log: warning!"
}

other-condition && {
   echo "log: all cool"
}
```

This conveys the intention of doing something **just** in one case, and that the negation of this is not interesting at all. There's a big warning you have to be aware of. The same as with lua's `... and .. or ..`, bash `||` and `&&` are not interchangeable for `if...else...end`. [BashWiki](https://mywiki.wooledge.org/BashPitfalls#cmd1_.26.26_cmd2_.7C.7C_cmd3) has an explanation why, but, the same as in Lua's case, if the "then" part returns false, the else will run.

这传达了在一种情况下做某事**只是**的意图，并且对此的否定一点也不有趣。有一个重要的警告你必须注意。与 lua 的 `... 和 .. 或 ..` 相同，bash `||` 和 `&&` 对于 `if...else...end` 不可互换。 [BashWiki](https://mywiki.wooledge.org/BashPitfalls#cmd1_.26.26_cmd2_.7C.7C_cmd3) 有一个解释，但是，与 Lua 的情况相同，如果“then”部分返回 false，则其他会跑。

There are lots of references to this, but I like this recent post where it explains it for arrays in higher level languages like ruby: https://jesseduffield.com/array-functions-and-the-rule-of-least-power /

对此有很多参考，但我喜欢最近的这篇文章，它用 ruby 等高级语言对数组进行了解释：https://jesseduffield.com/array-functions-and-the-rule-of-least-power /

An extended article of this kind of conditionals can be found [here](https://timvisee.com/blog/elegant-bash-conditionals/).

可以在 [此处](https://timvisee.com/blog/elegant-bash-conditionals/) 找到此类条件的扩展文章。

### 4.19 pushd/popd

### 4.19 推送/弹出

pushd and popd are used to move to some directory and go back to it in a stack fashion, so nesting can happen and you never lose track. The problem is that it still is on you to have a `popd` per `pushd`.

pushd 和 popd 用于移动到某个目录并以堆栈方式返回到该目录，因此可以发生嵌套并且您永远不会迷失方向。问题是你仍然需要为每个 `pushd` 设置一个 `popd`。

```
pushd /tmp/my-dir
  echo $PWD
popd
```

Here's an alternative way, that at least makes sure that you close all pushd with a popd.

这是另一种方法，至少可以确保您使用 popd 关闭所有 pushd。

Starting a new shell and cd-ing , will make all commands in that subshell be in that directory, and will come back to the old directory after closing the new spawned shell.

启动一个新的 shell 和 cd-ing ，将使该子 shell 中的所有命令都在该目录中，并在关闭新生成的 shell 后返回到旧目录。

```
(cd /tmp/my-dir
  ls
)
```

Remember to `inherit_errexit` or `set -e` inside the subshell if you need. That's a very easy trap to fall into.

如果需要，请记住在子 shell 中使用 `inherit_errexit` 或 `set -e`。这是一个非常容易落入的陷阱。

### 4.20 wrap functions

### 4.20 包装函数

Bash can't pass blocks of code around, but the alternative is to pass functions. More on that later.

Bash 不能传递代码块，但另一种方法是传递函数。稍后再谈。

```
mute() {
  "$@" >/dev/null
}

mute pushd /tmp/foobar
```

### 4.21 use [[

### 4.21 使用 [[

Unless you want your script to be POSIX compliant, use `[[` instead of `[`. `[` is a regular command. It's like `ls`, or `true`. You can check it by searching for a file named `[` in your path.

除非您希望您的脚本符合 POSIX，否则请使用 `[[` 而不是 `[`。 `[` 是一个常规命令。就像 `ls` 或 `true`。您可以通过在路径中搜索名为 `[` 的文件来检查它。

Being a normal command it always evaluates its params, like a regular function. On the other hand though, `[[` is a special bash operator, and it evaluates the parameters lazily.

作为一个普通的命令，它总是评估它的参数，就像一个普通的函数。但另一方面，`[[` 是一个特殊的 bash 运算符，它懒惰地评估参数。

```
# [[ does lazy evaluation:
[[ a = b && $(echo foo >&2) ]]

# [ does not:
[ a = b -a "$(echo foo >&2)” ]
```

Ref: https://lists.gnu.org/archive/html/help-bash/2014-06/msg00013.html

参考：https://lists.gnu.org/archive/html/help-bash/2014-06/msg00013.html

### 4.22 eval?

### 4.22 评估？

When you have mostly small functions that are mostly pure, you compose them like you'd do in any other language.

当您拥有大部分纯小函数时，您可以像使用任何其他语言一样编写它们。

In the following snippet, we are in a release script. Some step builds a package inside a docker container, another step tests a package already built.

在以下代码段中，我们位于发布脚本中。一些步骤在 docker 容器内构建一个包，另一个步骤测试已经构建的包。

A nice way to build ubuntus, for example, is to add an ARG to the Dockerfile so we can build several ubuntu versions using the same file.

例如，构建 ubuntus 的一个好方法是向 Dockerfile 添加一个 ARG，这样我们就可以使用同一个文件构建多个 ubuntu 版本。

It'd look like this:

它看起来像这样：

```
ARG VERSION
FROM ubuntu:$VERSION

RUN apt-get ...
...
```

We build that image and do all the building inside it, mounting a volume shared with our host, so we can extract our `.deb` file easily.

我们构建该映像并在其中完成所有构建，安装与我们的主机共享的卷，因此我们可以轻松提取我们的 `.deb` 文件。

After that, to do some smoke tests on the package, the idea is to install the `.deb` file in a fresh ubuntu image.

之后，要对包进行一些冒烟测试，想法是在新的 ubuntu 映像中安装 `.deb` 文件。

Let's pick the same base image we picked to build the package.

让我们选择与构建包相同的基础镜像。

```
# evaluate the string "centos:$VERSION" (that comes from
# centos/Dockerfile) in the current scope
# DISTRO is ubuntu:18.04
local VERSION=$(get_version $DISTRO) # VERSION==18.04
run_test "file.deb" "$(eval echo $(awk '/^FROM /{print $2; exit}' $LOCAL_PATH/$(get_dockerfile_for $DISTRO)))" # ubuntu:18.04
```

The usage of eval is there to interpolate the string that we get from the `FROM` in the current environment.

eval 的用法是在当前环境中插入我们从`FROM` 获得的字符串。

WARNING: You know, anything that uses `eval` is dangerous per se. Do not use it unless you know very well what you're doing AND the input is 100% under your control. Usually, more restricted commands can achieve what you want to do. In this particular case, you could use `envsubst`, or just manually replace `$\{?VERSION\}?` in a sed.

警告：你知道，任何使用 `eval` 的东西本身都是危险的。除非您非常清楚自己在做什么并且输入 100% 在您的控制之下，否则不要使用它。通常，更受限制的命令可以实现您想要的功能。在这种特殊情况下，您可以使用 `envsubst`，或者只是手动替换 sed 中的 `$\{?VERSION\}?`。

```
test_release "$PACKAGE_PATH" $(awk '/^FROM /{print $2; exit}' $LOCAL_PATH/$(get_dockerfile_for $DISTRO) | sed -e "s/\$VERSION/$VERSION/")
```

Yet another way is using [shell parameter expansions.](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)

另一种方法是使用 [shell 参数扩展。](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)

```
var1=value
echo 'this is $var1' >/tmp/f.txt
f=$(cat /tmp/f.txt)
echo "${f}"  # this is $var1
echo "${f@P}"  # this is value
```

### 4.23 pass commands around

### 4.23 传递命令

This one uses [DRY_RUN](https://raimonster.com/scripting-field-guide/#org6614f4a). While refactoring a script that does some curls, we want to make sure that our refactored version does the exact same calls in the same order.

这个使用 [DRY_RUN](https://raimonster.com/scripting-field-guide/#org6614f4a)。在重构执行一些 curl 的脚本时，我们希望确保我们重构的版本以相同的顺序执行完全相同的调用。

```
compare_outputs() {
  export DRY_RUN=1
  git checkout b1
  $@ 2>/tmp/1.out
  git checkout b2
  $@ 2>/tmp/2.out
  echo "diffing"
  diff /tmp/1.out /tmp/2.out
}
compare_outputs ./release.sh -p rhel:6 -R 'internal-preview'
```

First we create a function `compare_outputs`, that gets a command to run as parameters. The function will run it once, redirecting the standard error to a file `/tmp/1.out`.

首先我们创建一个函数`compare_outputs`，它获取一个命令作为参数运行。该函数将运行一次，将标准错误重定向到文件`/tmp/1.out`。

Then, it checks out the branch that contains our refactored version, and will run the command again, redirecting standard error to `/tmp/2.out`, and will diff the two outputs.

然后，它会检查包含我们重构版本的分支，并将再次运行命令，将标准错误重定向到 `/tmp/2.out`，并对两个输出进行比较。

In case there's a difference between the two, `diff` will output them, and the function will return the non-zero exit value of diff. If everything went fine, `compare_outputs` will succeed.

如果两者之间存在差异，`diff` 将输出它们，并且该函数将返回 diff 的非零退出值。如果一切顺利，`compare_outputs` 就会成功。

Now that we know that for these inputs the command runs fine, we want to find out if it works for other types of releases, not only internal-preview.

现在我们知道对于这些输入命令运行良好，我们想知道它是否适用于其他类型的版本，而不仅仅是内部预览。

Here I'm using zsh's global aliases to give a much more fluid interface to the commands, but you can use the regular while/for loops:

在这里，我使用 zsh 的全局别名为命令提供更流畅的界面，但您可以使用常规的 while/for 循环：

```
alias -g SPLIT='|tr " " "\n" '
alias -g FORI='|while read i ;do '
alias -g IROF=';done '

set -e
echo "ga internal-preview rc1 rc2" SPLIT FORI
   noglob compare_outputs ./release.sh -p rhel:8 -R "$i"
IROF
```

So, combining the two, we can have a really smooth way of iterating over the possibilities, without really messing into the details of loops.

因此，将两者结合起来，我们可以有一种非常流畅的方式来迭代各种可能性，而不会真正弄乱循环的细节。

WARNING: This approach is not robust enough to put it anywhere in production, but to write quick one off scripts is a killer. Experimenting in a shell and creating tools and 2nd order tools to make interaction faster builds a language that grows on you, and keeps improving your toolbelt.

警告：这种方法不够健壮，无法将其放在生产中的任何地方，但快速编写一个脚本是一个杀手。在 shell 中进行试验并创建工具和二阶工具以加快交互速度，构建了一种在您身上成长的语言，并不断改进您的工具带。

### 4.24 The Toplevel Is Hopeless[1](https://raimonster.com/scripting-field-guide/#fn.1)

### 4.24 顶层无望[1](https://raimonster.com/scripting-field-guide/#fn.1)

Shellscripts are thought as quick one-off programs, but when they are useful, they are sticky, so you better write them from the start as if it would be permanent. The upfront cost is very low anyway. Structure the script like a regular app.

Shellscripts 被认为是一次性的快速程序，但是当它们有用时，它们是粘性的，所以你最好从一开始就编写它们，就好像它是永久的一样。无论如何，前期成本非常低。像普通应用程序一样构建脚本。

Bash is extremely permissive in what it allows to be coded and ran. By default, failures do not make the program exit or throw an exception (no exceptions). And for some reason, the common usage of shellscripts is to put everything in the top level. Don't do that. Do the least possible things in the toplevel.

Bash 在允许编码和运行方面极为宽松。默认情况下，失败不会使程序退出或抛出异常（无异常）。出于某种原因，shellscripts 的常见用法是将所有内容都放在顶层。不要那样做。在顶层做尽可能少的事情。

A way to improve the defaults, is setting a bunch of flags that make the script stricter, so it fails on many situations you'd want to stop anyway because something went wrong.

改进默认值的一种方法是设置一堆使脚本更严格的标志，因此它在许多情况下失败，您无论如何都想停止因为出现问题。

```
#!/usr/bin/env bash
set -eEuo pipefail
shopt -s inherit_errexit

main() {
  parse_args
  validate_args
  do_things
  cleanup
}

main "$@"
```

Ref: https://dougrichardson.us/2018/08/03/fail-fast-bash-scripting.html

参考：https://dougrichardson.us/2018/08/03/fail-fast-bash-scripting.html

### 4.25 Check your deps

### 4.25 检查你的 deps

Giving useful information to the users will help them using the script, and you debugging it. Script dependencies is a common use case that we'll do it in a nice way.

向用户提供有用的信息将帮助他们使用脚本，并且您可以对其进行调试。脚本依赖是一个常见的用例，我们会以一种很好的方式做到这一点。

```
deps() {
  for dep in "$@";do
    mute which "$dep" ||die "$dep dependency missing"
  done
}

main() {
  deps jq curl
  # ...
}
```

### 4.26 source files

### 4.26 源文件

`source` is like `require` or `import` in some programming languages. It evaluates the sourced file in the context of the current script, so you get all definitions in your environment.

`source` 类似于某些编程语言中的 `require` 或 `import`。它在当前脚本的上下文中评估源文件，因此您可以获得环境中的所有定义。

It's simple, but it helps you get used to modularize your code into libraries.

这很简单，但它可以帮助您习惯将代码模块化为库。

Be careful, it's very rudimentary, and it will be overwriting old vars or functions if names clash. There's no namespacing happening there.

小心，它非常初级，如果名称冲突，它将覆盖旧的 vars 或函数。那里没有命名空间。

```
source file.sh

# the same
.file.sh
```

### 4.27 Use Scripts as a Libs

### 4.27 使用脚本作为 Libs

A python-inspired way of using scripts as loadable libraries is to check whether the current file was the one that was called originally or it's being just sourced.

将脚本用作可加载库的一种受 Python 启发的方法是检查当前文件是最初调用的文件还是刚刚被调用的文件。

Again, no side effects in load time makes this functionality possible. otherwise, you're on your own.

同样，加载时间没有副作用使此功能成为可能。否则，你就靠自己了。

```
# Allow sourcing of this script
if [[ $(basename "$(realpath "$0")") == "${BASH_SOURCE}" ]];then
  setup
  parse_args "$@"
  main
fi
```

### 4.28 Tmpfiles Everywhere 

### 4.28 Tmpfiles 无处不在

Your script is not going to run alone. Don't assume paths are fixed or known.

您的脚本不会单独运行。不要假设路径是固定的或已知的。

CI/CD Pipelines run many jobs in the same node and files can start clashing.

CI/CD 管道在同一个节点中运行许多作业，文件可能会开始发生冲突。

Make use of `$(mktemp -d /tmp/foo-bar.XXXXX)`. If you have to patch a file, do it in a clean fresh copy. Don't modify files in old paths

使用`$(mktemp -d /tmp/foo-bar.XXXXX)`。如果您必须修补文件，请在干净的新副本中进行。不要修改旧路径中的文件

If you HAVE TO modify paths, do it idempotently. But really, don't do it.

如果您必须修改路径，请以幂等方式进行。但真的，不要这样做。

```
git_clone_tmp() {
  local repo=${1:?repo is required}
  local ref=${2:?ref is required}
  tmpath=$(mktemp -d "/tmp/cloned-$repo-XXXXX")
  on_exit "rm -rf $tmpath"
  git clone -b ${ref} $repo $tmpath
}
```

CAVEAT: You have to manually delete the directory if you want it cleaned.

警告：如果你想清理目录，你必须手动删除它。

Here's an article with very good advice on [tempfiles](https://www.netmeister.org/blog/mktemp.html).

这是一篇关于 [tempfiles](https://www.netmeister.org/blog/mktemp.html) 的非常好的建议的文章。

### 4.29 Cleanup tasks with trap

### 4.29 使用陷阱清理任务

`trap` is used to 'subscribe' a callback when something happens. Many times it's used on exit. It's a good thing to cleanup tmpdirs after your script exits, so you can use the output of `mktemp -d` and subscribe a cleanup function for it.

`trap` 用于在发生某些事情时“订阅”回调。很多时候它在退出时使用。在脚本退出后清理 tmpdirs 是一件好事，因此您可以使用 `mktemp -d` 的输出并为其订阅清理函数。

```
on_exit() {
  rm -rf $1
}
local tmpath=$(mktemp -d /tmp/foo-bar.XXXXX)
trap "on_exit $tmpath" EXIT SIGINT
```

### 4.30 array of callbacks on_exit

### 4.30 on_exit 回调数组

Level up that pattern, we can have a helper to add callbacks to run on exit. Get used to these kind of patterns, they are super powerful and save you lots of manual bookkeeping.

升级该模式，我们可以有一个助手来添加回调以在退出时运行。习惯这种模式，它们非常强大，可以为您节省大量手动簿记。

```
ON_EXIT=()
EXIT_RES=

function on_exit_fn {
  EXIT_RES=$?
  for cb in "${ON_EXIT[@]}";do $cb ||true;done
  return $EXIT_RES
}

trap on_exit_fn EXIT SIGINT

function on_exit {
  ON_EXIT+=("$@")
}

local v_id=$(docker volume create)
on_exit "docker volume rm $v_id"
# Use your v_id knowing that it'll be available during your script but
# will be cleaned up before exiting.
```

### 4.31 stacktrace on error

### 4.31 错误堆栈跟踪

Here's a nice helper for debugging errors in bash. In case of non-0 exit, it prints a stacktrace.

这是在 bash 中调试错误的好帮手。在非 0 退出的情况下，它会打印堆栈跟踪。

```
set -Eeuo pipefail
trap stacktrace EXIT
stacktrace() {
    rc="$?"
    if [ $rc != 0 ];then
        printf '\nThe command "%s" triggerd a stacktrace:\n' "$BASH_COMMAND"
        for i in $(seq 1 $((${#FUNCNAME[@]} - 2)));do
          j=$((i+1));
          printf '\t%s: %s() called in %s:%s\n' "${BASH_SOURCE[$i]}" "${FUNCNAME[$i]}" "${BASH_SOURCE[$j]}" "${BASH_LINENO[$i]}";
        done
    fi
}
```

ref: https://news.ycombinator.com/item?id=26644110

参考：https://news.ycombinator.com/item?id=26644110

### 4.32 Dots and colons allowed in function names!

### 4.32 函数名中允许使用点和冒号！

A way to split the namespace is to have libs define functions with their own namespace.

拆分命名空间的一种方法是让库定义具有自己命名空间的函数。

I've gotten used to use dots or colons as namespace separator.

我已经习惯使用点或冒号作为命名空间分隔符。

```
semver.greater() {
 # ...
}
```

or

或者

```
semver:greater() {
 # ...
}
```

### 4.33 make steps of the process as composable as possible by using "$@"

### 4.33 通过使用“$@”使流程的步骤尽可能组合

By using `$@` to pass commands as parameters around you can get to a degree of composability that allows for a nice chaining of commands.

通过使用`$@` 将命令作为参数传递，您可以获得一定程度的可组合性，从而可以很好地链接命令。

here's a very simple version of `watch`. See how you can `every 2   ls -la`. I think that style is called Bernstein Chaining. But I'm not sure if it's exactly the same. It also looks like currying or partial evaluation to me if you squint a little bit.

这是一个非常简单的 `watch` 版本。看看如何“每 2 ls -la”。我认为这种风格被称为伯恩斯坦链。但我不确定它是否完全相同。如果你眯着眼睛看，它对我来说也像是咖喱或部分评估。

```
every() {
   secs=$1
   shift
   while true;do
     "$@"
     sleep $secs;
   done
 }
```

As you know by now, bash doesn't pass blocks of code around, but the alternative is to pass function names.

正如您现在所知，bash 不会传递代码块，但另一种方法是传递函数名称。

```
mute() {
  $@ >/dev/null 2>/dev/null
}
mute ls
```

So now we can create the most stupid command composition ever:

所以现在我们可以创建有史以来最愚蠢的命令组合：

```
every 1 mute echo hi
#or
mute every 1 echo hi
```

For the particular redirection problem, another option is to use aliases. Redirects can be written anywhere on your CLI (not just at the end), so the following will work using a plain alias:

对于特定的重定向问题，另一种选择是使用别名。重定向可以写在你的 CLI 的任何地方（不仅仅是在最后），所以以下内容将使用普通别名：

```
alias mute='>/dev/null 2>/dev/null'
mute ls
```

- https://www.oilshell.org/blog/2017/01/13.html

- https://www.oilshell.org/blog/2017/01/13.html

### 4.34 do_times/foreach_*

### 4.34 do_times/foreach_*

shellscripts are highly side-efffecty, and even though the scoping of variables is not very empowering, you can get a limited amount of decomposition of loops by passing function names.

shellscripts 具有很高的副作用，即使变量的范围不是很强大，您也可以通过传递函数名称来获得有限数量的循环分解。

This is a lame example, but I hope it shows the use case, it allows you to group already existing functions while taking advantage of a fixed looping iterator, and leaving traces of the current loop vars in the global "variable" environment.

这是一个蹩脚的例子，但我希望它展示了用例，它允许您在利用固定循环迭代器的同时对现有函数进行分组，并在全局“变量”环境中留下当前循环变量的痕迹。

```
create_user() {
  uname="u$1" # leave uname in the global env so later functions see it
  http :8080/users name="$uname"
}

create_pet() {
  pname="p$1"
  http :8080/users/$uname/pets name="$pname"
}

create_bundle() {
  create_user $1
  create_pet $1
}

do_times() {
  local n=$1;shift
  for i in $(seq $n);do
    "$@" $i
  done
}

do_times 15 create_bundle
```

A bit more complex is runnning a command to every repo in an org:

更复杂的是对组织中的每个 repo 运行命令：

```
run_tests() {
  ./ci/test.sh
}

foreach_repo_with_index() {
  local counter=0
  local repos=$(http https://api.github.com/users/$1/repos)
  shift
  for entry in $(echo $repos | jq -r '.[].git_url');do
    (git_clone_tmp $entry master
     cd $tmpath
     "$@" $counter $entry
    )
    ((counter=counter+1))
  done
}

foreach_repo_with_index kidd run_tests
```

### 4.35 <(foo) and >(foo)

### 4.35 <(foo) 和 >(foo)

Some commands ask for files as inputs. And sometimes you have that file, but sometimes you're only creating that file to pass it to the command. In those cases, creating temporary files is not necessary if you use `<(cmd)`. Here's a way to diff the output of 2 commands without putting them in a temporary file.

一些命令要求文件作为输入。有时您拥有该文件，但有时您只是创建该文件以将其传递给命令。在这些情况下，如果您使用 `<(cmd)`，则不需要创建临时文件。这是一种比较 2 个命令的输出而不将它们放在临时文件中的方法。

```
diff <(date) <(date)
diff <(date) <(sleep 1; date)
```

The same happens with outputs. Commands that ask you for a destination file. You can trick them by using `>(command)` as a file. A nice trick is to use `>(cat)` to know what's going on there. Also useful to send stuff to the clipboard `>(xclip)` before running something on the output.

输出也是如此。要求您提供目标文件的命令。你可以通过使用 `>(command)` 作为文件来欺骗它们。一个很好的技巧是使用 `>(cat)` 来了解那里发生了什么。在对输出运行某些内容之前将内容发送到剪贴板 `>(xclip)` 也很有用。

What the shell does in those cases is to bind a file descriptor of the process created inside `< or >` to the first process.

在这些情况下，shell 所做的是将在 `< 或 >` 中创建的进程的文件描述符绑定到第一个进程。

You can experiment with those using commands like `echo <(pwd)`.

您可以使用像 `echo <(pwd)` 这样的命令来试验那些。

In Zsh you can use `m-x expand-word` to see the file descriptors being expanded.

在 Zsh 中，您可以使用 `m-x expand-word` 来查看被扩展的文件描述符。

A way to peek into a huge pipe is to `tee >(cat)`

窥视巨大管道的一种方法是“tee >(cat)”

### 4.36 Use xargs

### 4.36 使用 xargs

Continuing with other ways of plumbing commands into other commands, there's `xargs`. Some commands work seamlessly with pipes, by taking inputs from stdin and printing to stdout. But some others like to work with files, and they ask for their parameters in their args list. For example, `evince`. It wouldn't be even expected to cat a pdf and pass it to evince through stdin.

继续使用将命令连接到其他命令的其他方法，还有 `xargs`。通过从 stdin 获取输入并打印到 stdout，一些命令可以与管道无缝协作。但是其他一些人喜欢处理文件，他们在他们的 args 列表中要求他们的参数。例如，`evince`。甚至不希望将 pdf 转换为 pdf 并通过 stdin 将其传递给 evince。

In general, to convert from this pattern: `cmd param` to `echo   param| cmd`, xargs can be helpful. Look at its man page to know how to split or batch args in multiple `cmd` calls.

一般来说，要从这种模式转换：`cmd param` 到 `echo param| cmd`, xargs 可能会有所帮助。查看其手册页以了解如何在多个 `cmd` 调用中拆分或批处理 args。

Xargs is helpful for parallelizing work. You should look at its man page, but just know it can help in running parallel processes (check `-P` in its man).

Xargs 有助于并行化工作。您应该查看它的手册页，但只知道它可以帮助运行并行进程（检查其手册中的“-P”）。

Other tips on this great [Guide to xargs](https://www.oilshell.org/blog/2021/08/xargs.html).

关于这个很棒的 [xargs 指南](https://www.oilshell.org/blog/2021/08/xargs.html) 的其他提示。

### 4.37 change loops for "mapping/reducing" functions

### 4.37 更改“映射/减少”函数的循环

#### 4.37.1 find

#### 4.37.1 查找

Many times we want to run the same operation or test to lots of files. Instead of looping for each file, think if `find -exec` would solve it. Also, find supports multiple directories.

很多时候我们想要对大量文件运行相同的操作或测试。与其为每个文件循环，不如想想 `find -exec` 是否能解决它。此外，find 支持多个目录。

```
dirs=("/usr/local/bin" "/usr/bin")
for d in "${dirs[@]}";do
  for f in $(find "$d");do
    echo "check if owner of $f is johndoe and group is johndoe"
    [ `stat -c %U:%G $f` == "johndoe:johndoe" ] ||die "error"
  done
done
```

Compare it to:

比较一下：

```
[ $(find "$dirs[@]" -exec stat -c '%U:%G' {} \; | grep -vc "johndoe:johndoe") == "0" ] ||die "error"
```

Other examples might be:

其他示例可能是：

```
# count all lines of all docx in this dir
find .-type f -name "*docx" -exec pandoc "{}" -t plain \;|wc -l

#All your files have the same owner and group permissions
[ $(find "$files[@]" -exec stat -c '%a' {} \; | grep -Evc "^(.)\1") == "0" ]
```

#### 4.37.2 grep -Fvf

#### 4.37.2 grep -Fvf

Three magical flags that go well together.

三个神奇的旗帜搭配得很好。

`-f` list of patterns to match as a file. `-F` interpret the "pattern" as a Fixed string, not a pattern/regex `-v` negate the output. Print non matching lines.

`-f` 模式列表作为文件匹配。 `-F` 将“模式”解释为固定字符串，而不是模式/正则表达式 `-v` 否定输出。打印不匹配的行。

The cool thing about combining `-f` and `-v` is that the negative matches mean "lines that are not ANY of the ones in the pattern list". So you can do list diffing. like `sort + diff` but more flexible.

结合`-f` 和`-v` 的一个很酷的事情是负匹配意味着“不是模式列表中的任何行”。所以你可以做列表差异。像`sort + diff`但更灵活。

Here's a practical case of finding version numbers that we have a tag for, that do not have a title in the readme

这是查找我们有标签的版本号的实际案例，这些版本号在自述文件中没有标题

```
f="readme.md"
!grep -Fvf <(grep -P "^# \d\.\d\.\d\.\d$" "$f" | sed -e 's/^# //') \
            <(git tag | grep -P "^\d\.\d\.\d\.\d$")
```

### 4.38 pass flags as a splatted array 

### 4.38 将标志作为一个 splatted 数组传递

There's quite a bit to chew on this example. First of all, the core pattern is to build up your commandline options with an array, and splat it in the final command line. For complex commands like `docker` where you easily have 10+ flags it's a visual aid, and also opens up the opportunities for reusing or abstracting sets of options to logical blocks.

这个例子有很多值得咀嚼的地方。首先，核心模式是使用数组构建命令行选项，并将其放入最终命令行。对于像 `docker` 这样的复杂命令，你可以轻松拥有 10 多个标志，它是一种视觉辅助，也为重用或抽象逻辑块的选项集提供了机会。

Once it's an array, we can add elements conditionally to that array depending on the current run, and build the line that we'll be running in the end.

一旦它是一个数组，我们就可以根据当前的运行有条件地向该数组添加元素，并构建我们最终将运行的行。

```
# Allows Ctrl-C'ing on interactive shells
INTERACTIVE=
if [[ -t 1 ]];then INTERACTIVE="-it";fi

local flags=(
  # We mount it as read-only, so we make sure we are not writing anything
  # in there, and that everything is explicitly defined
  "-v $LOCAL_PATH/build-dir:/build-dir:ro,delegated"
  "-v $OUTPUT_DIR:/output:rw,consistent"
  "-v $tmp_dir:/tmp/work:rw,delegated"
)
if [[ -n $LOCAL_PATH ]];then
  flags+=("-v $(realpath $LOCAL_PATH)/overrides/my-other-file:/build-dir/build.json:ro")
  flags+=("-e LOCAL_PATH=/tmp/local")
fi

local v_id=$(docker volume create)
flags+=("-v $v_id:/tmp/build")
on_exit "docker volume rm $v_id"

docker run --rm $INTERACTIVE ${flags[*]} $image touch /tmp/build/foo.txt

docker run --rm $INTERACTIVE ${flags[*]} fpm:latest fpm-build /tmp/build/foo.txt
on_exit "chown_cache $tmp_dir"
```

In this example we see another cool trick. Mounting a volume in 2 differrent containers, so not for the purpose of sharing a local file/dir with the host but to share it between themselves. In that case, the 2 containers don't even coexist temporarily, but use the volume as a conveyor belt, passing it from container to container, and each one applies "its thing".

在这个例子中，我们看到了另一个很酷的技巧。在 2 个不同的容器中安装卷，因此不是为了与主机共享本地文件/目录，而是为了在它们之间共享它。在这种情况下，两个容器甚至暂时不共存，而是将体积用作传送带，将其从一个容器传递到另一个容器，并且每个容器都应用“它的东西”。

After all the mess, someone has to cleanup everything, but we know how to do it with `on_exit` trick.

在所有的混乱之后，必须有人清理所有内容，但我们知道如何使用“on_exit”技巧来做到这一点。

### 4.39 inherit_errexit

### 4.39 inherit_errexit

bash 4.4+ , you can `shopt -s inherit_errexit`, and subshells will inherit the errexit flag value. meaning that if you `set -Ee`, anything that runs inside a subshell will throw an error at the moment any command exits with `!=0`.

bash 4.4+ ，你可以`shopt -s inherit_errexit`，子shell将继承errexit标志值。这意味着如果你`set -Ee`，在子shell中运行的任何东西都会在任何命令以`!=0`退出时抛出错误。

### 4.40 GNU Parallel

### 4.40 GNU 并行

I can't recommend [parallel](https://www.gnu.org/software/parallel/) enough. The same as xargs, but in a much more flexible way, parallel lets you run various jobs at a time. If you have this tool into account, it doesn't just speed up your runtimes, but it will force you write cleaner code. Parallel execution will test your scripts so if they are not using randomized tmp working directories, things will clash, etc…

我不能完全推荐 [parallel](https://www.gnu.org/software/parallel/)。与 xargs 相同，但以更灵活的方式，parallel 允许您一次运行各种作业。如果您将这个工具考虑在内，它不仅会加速您的运行时，还会迫使您编写更简洁的代码。并行执行将测试您的脚本，因此如果它们不使用随机 tmp 工作目录，事情就会发生冲突，等等......

Parallel in itself is such a hackerfriendly tool it deserves to be deeply learned. You can use it just locally to run a process per core, you can send jobs to several machines connected via a simple ssh, you can bind tmux or sqlite to it, or you can write a trivial job queuing system.

Parallel 本身就是一个对黑客友好的工具，值得深入学习。您可以仅在本地使用它来为每个核心运行一个进程，您可以将作业发送到通过简单的 ssh 连接的多台机器，您可以将 tmux 或 sqlite 绑定到它，或者您可以编写一个简单的作业排队系统。

Man pages and official examples are a goldmine.

手册页和官方示例是一座金矿。

### 4.41 HEREDOCS

### 4.41 HEREDOCS

- Basic usage of heredocs:

- heredocs的基本用法：

```
      echo <<EOF
$interpolated
\$non_interpolated
EOF
```

- A dash after `<<` replaces trailing spaces in here docs

- `<<` 后的破折号替换此处文档中的尾随空格

```
echo <<-EOF
$var
there
EOF
```

- quoting the identifier disables interpolation of variables

- 引用标识符禁用变量插值

```
echo <<'EOF'
$non_interpolated
there
EOF
```

The [bash manual](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Here-Documents) is super concise and to the point there.

[bash 手册](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Here-Documents) 非常简明扼要。

The most complex case is having a heredoc which contains some bash code that is ment to run remotely, and wanting to interpolate some variables in the local env, and some in the remote one. [Escaping is enough for regular](https://stackoverflow.com/questions/4937792/using-variables-inside-a-bash-heredoc) vars, but if you want to use special (`$1, $!,   $? `…) vars, check this out:

最复杂的情况是有一个 heredoc，其中包含一些可以远程运行的 bash 代码，并且想要在本地环境中插入一些变量，在远程环境中插入一些变量。 [普通的转义就足够了](https://stackoverflow.com/questions/4937792/using-variables-inside-a-bash-heredoc) vars，但如果你想使用特殊的 (`$1, $!, $? `...) 变量，检查一下：

```
deploy () {
  local deploy_dir="$1"
  ssh server1 'bash -s' <<EOSSH
  echo "deploying => ${deploy_dir}"
  ls -tdr ~/quote-clojure/artifacts/deploy-be-* |head -n -5 |xargs rm -rf
  ln -nfs "\$HOME/quote-clojure/artifacts/${deploy_dir}/quote-clojure-0.1.0-SNAPSHOT-standalone.jar" "\$HOME/quote-clojure/artifacts/current.jar"
  cat ~/quote-clojure/artifacts/current.pid |xargs kill
  java -jar ~/quote-clojure/artifacts/current.jar &>/dev/null &
  lastpid=\$(echo \$!)
  echo \$lastpid > ~/quote-clojure/artifacts/current.pid
EOSSH
}
```

### 4.42 HERESTRINGS

### 4.42 杂种

It's the stripped down version of HEREDOCS. Inline a single string (or output of a single command) as an input string. 

这是HEREDOCS 的精简版。内联单个字符串（或单个命令的输出）作为输入字符串。

It's [kinda similar](https://stackoverflow.com/questions/18116343/whats-the-difference-between-here-string-and-echo-pipe) to what you could do with a pipe.

它[有点相似](https://stackoverflow.com/questions/18116343/whats-the-difference-between-here-string-and-echo-pipe) 与您可以用管道做的事情类似。

```
cat <<<"HELLO"
cat <<<$(echo "HELLO")
echo "ECHOPIPE" |/bin/cat <(seq 5) <<<"HERESTRING"
echo "ECHOPIPE" |/bin/cat <(seq 5) - <<<"HERESTRING"
```

### 4.43 \__DATA__

### 4.43 \__数据__

I loved the Perl [\__END__ and \__DATA__](https://perldoc.perl.org/perldata#Special-Literals), features and realized it's possible to do it in shellscripts.

我喜欢 Perl [\__END__ 和 \__DATA__](https://perldoc.perl.org/perldata#Special-Literals) 的特性，并意识到可以在 shellscripts 中做到这一点。

You can append to the current file. Here's an example of a super simple bookmark "manager":

您可以附加到当前文件。下面是一个超级简单的书签“管理器”的例子：

```
#!/usr/bin/env sh

cmd_add() {
    shift
    echo "$@" >> "$0"
}

cmd_go() {
    sed '0,/^__DATA__$/d' "$0" |
    dmenu -i -l 20  |
    rev |cut -f1 -d' ' |rev |
    xargs xdg-open
}

main() {
    cmd_${1:-go} $@
}

main $@
exit

__DATA__
r/emacs             https://www.reddit.com/r/emacs
```

### 4.44 Secrets

### 4.44 秘密

As, [Smallstep article](https://smallstep.com/blog/command-line-secrets/) explains, you should be careful when passing around secret data like tokens between processes.

正如 [Smallstep 文章](https://smallstep.com/blog/command-line-secrets/) 解释的那样，在进程之间传递令牌等秘密数据时应该小心。

Env vars are not really safe, and there are a few tricks you can use to cover your assets.

环境变量并不是真正安全的，您可以使用一些技巧来覆盖您的资产。

## 5 Interactive

## 5 互动

### 5.1 Save your small scripts

### 5.1 保存你的小脚本

Rome wasn't built in a day, and like having a journal log, most of the little scripts you create, once you have enough discipline will be useful for some other cases, and your functions will be reusable.

罗马不是一天建成的，就像拥有日志日志一样，您创建的大多数小脚本，一旦您有足够的纪律，将在其他一些情况下有用，并且您的功能将是可重用的。

Save your scripts into files early on, instead of crunching everything in the repl. learn how to use a decent editor that shortens the feedback cycle as much as possible.

尽早将脚本保存到文件中，而不是在 repl 中处理所有内容。学习如何使用一个体面的编辑器来尽可能缩短反馈周期。

### 5.2 Increased Interactivity

### 5.2 增加交互性

Knowing your shell's shortcuts for interactive use is a must. The same way you learned to touchtype and you learned your editor, you should learn all the shortcuts for your shell. Here's some of them.

了解用于交互使用的 shell 快捷方式是必须的。就像您学习触摸输入和学习编辑器一样，您应该学习 shell 的所有快捷方式。这是其中一些。

| key    | action                              |
| ------ |----------------------------------- |
| Ctrl-r | reverse-history-search              |
| C-a    | beginning-of-line                   |
| C-e    | end-of-line                         |
| C-w    | delete-word-backwards               |
| C-k    | kill-line (from point to eol)       |
| C-y    | paste last killed thing             |
| A-y    | previous killed thing (after a c-y) |
| C-p    | previous-line                       |
| C-n    | next-line                           |
| A-. | insert last agument                 |
| A-/    | dabbrev-expand                      |

|关键|行动 |
| | Ctrl-r |反向历史搜索|
| C-a |行首 |
| C-e |线尾|
| C-w |向后删除字|
| C-k | kill-line（从点到eol）|
| C-y |粘贴最后杀死的东西|
| A-y |以前被杀死的东西（在 c-y 之后）|
| C-p |上一条 |
| C-n |下一行 |
|一种-。 |插入最后一个参数 |
| A-/ | dabbrev-expand |

A written form of `A-.` is `$_`. It retains the last argment and puts it in $_. `test -f "FILE" && source "$_"`.

`A-.` 的书面形式是 `$_`。它保留最后一个参数并将其放入 $_。 `test -f "FILE" && source "$_"`。

### 5.3 Aliases

### 5.3 别名

Aliases are very simple substitutions of commands for a sequence of other commands. Usual example is

别名是非常简单的命令替换一系列其他命令。通常的例子是

```
alias ls='ls --auto-color'
```

Now let's move on to the interesting stuff.

现在让我们继续讨论有趣的事情。

### 5.4 functions can generate aliases

### 5.4 函数可以生成别名

Aliases live in a global namespace for the shell, so no matter where you define them, they take effect globally, possibly overwriting older aliases with the same name.

别名存在于 shell 的全局命名空间中，因此无论您在何处定义它们，它们都会全局生效，可能会覆盖旧的同名别名。

Well, it's not lexical scope (far from it), but using aliases you can create a string that snapshots the value you want, and capture it to run it later.

嗯，它不是词法作用域（远非如此），但是使用别名您可以创建一个字符串来对您想要的值进行快照，并捕获它以便稍后运行。

Some fun stuff:

一些有趣的东西：

- aliasgen. Create an alias for each directory in ~/workspace/. This is superceeded by `CDPATH`, but the trick is still cool.

- 别名。为 ~/workspace/ 中的每个目录创建一个别名。这被`CDPATH`取代，但这个技巧仍然很酷。

```
aliasgen() {
  for i in ~/workspace/*(/) ;do
      DIR=$(basename $i) ;
       alias $DIR="cd ~/workspace/$i";
  done
}
aliasgen
```

- a make a shortcut to the current directory.

- 制作当前目录的快捷方式。

```
function a() { alias $1=cd\ $PWD;}

mkdir -p /tmp/fing-longer
cd /tmp/fing-longer
a fl
cd /
fl
echo $PWD   # /tmp/fing-longer
```

A man can dream…

一个人可以梦想……

- unhist. functions can create aliases, and functions can receive functions as parameters (as a string (function name)), so we can combine them to advise existing functions.

   - 非历史性的。函数可以创建别名，函数可以接收函数作为参数（作为字符串（函数名称）），因此我们可以将它们组合起来为现有函数提供建议。

  ```
   unhist () {
    alias $1=" $1"
  }
  unhist unhist
  unhist grep
  unhist rg
  
  noglobber() {
      alias $1="noglob $1"
  }
  noglobber http
  noglobber curl
  noglobber git
  ```

- Problem: These commands do not compose. Combination of 2 of those doesn't work, because the second acts just on the textual representation that it received, not the current value of the alias.

- 问题：这些命令不构成。其中 2 个的组合不起作用，因为第二个仅作用于它收到的文本表示，而不是别名的当前值。

### 5.5 Override (advise?) common functions 

### 5.5 覆盖（建议？）常用函数

Overriding commands is generally a bad practice as it violates the principle of least surprise, but there might be occasions (mostly in your local machine) where you can integrate awesome finetunnings to your toolbelt.

覆盖命令通常是一种不好的做法，因为它违反了最小惊喜原则，但在某些情况下（主要是在您的本地机器中），您可以将很棒的微调集成到您的工具带中。

Here we're going to get the original docker binary file location. After that we declare a function called `docker` that will proxy the parameters to the original `docker` program UNLESS you're calling `docker run`. In that case, we're injecting a mouted volume that mounts `/root/.bash_history` of the container to a file hosted in the host (duh). That's a pretty cool way of keeping a history of your recent commands in your containers, no matter how many times you start and kill them.

在这里，我们将获取原始 docker 二进制文件位置。之后，我们声明了一个名为 `docker` 的函数，它将把参数代理给原始的 `docker` 程序，除非你正在调用 `docker run`。在这种情况下，我们将注入一个将容器的 `/root/.bash_history` 挂载到主机中托管的文件的 mouted 卷（废话）。这是一种非常酷的方式，可以在容器中保存最近命令的历史记录，无论您启动和终止它们多少次。

```
DOCKER_ORIG=$(which docker)
docker () {
    if [[ $1 == "run" ]];then
        shift
        $DOCKER_ORIG run -v $HOME/.shared_bash_history:/root/.bash_history "$@"
    else
        $DOCKER_ORIG "$@"
    fi
}
```

I'm particularly fond of this trick, as it saved me tons of typing. But at a personal level, it was mindblowing that sharing this around the internet caused the most disparity of opinions. Also, I recently read the greate book "Docker in Practice" by [Ian Miell](https://github.com/ianmiell) and there's a snippet that is 99.9% like the one I created myself. That was a very cool moment.

我特别喜欢这个技巧，因为它为我节省了大量的打字时间。但在个人层面上，令人震惊的是，在互联网上分享这一点引起了最大的意见分歧。此外，我最近阅读了 [Ian Miell](https://github.com/ianmiell) 的伟大著作“实践中的 Docker”，其中有一个片段与我自己创建的片段 99.9% 相似。那是一个非常酷的时刻。

### 5.6 Faster iteration on pipes

### 5.6 更快的管道迭代

When testing complex pipelines:

测试复杂管道时：

- Make them pure (no side effects).
- One command per line.
- End lines with the pipe character.
- During development, end the pipeline with `cat`.

- 使它们纯净（无副作用）。
- 每行一个命令。
- 以竖线字符结束行。
- 在开发过程中，以 `cat` 结束管道。

I usually use `watch -n1 'code.sh'` in a split window so I see the results of what I'm doing. The advantage of

我通常在拆分窗口中使用`watch -n1 'code.sh'`，这样我就能看到我正在做的事情的结果。的优势

```
curl https://www.example.com/videos/              |
  pup 'figure.mg > a attr{href}'                  |
  head -1                                         |
  xargs -I{} curl -L "https://www.example.com/{}" |
  pup 'script'                                    |
  grep file:                                      |
  sed -e "s/.*\(http[^ \"']*\).*/\1/"             |
# xargs vlc                                       |
  cat
```

Over

超过

```
curl https://www.example.com/videos/                \
  |pup 'figure.mg > a attr{href}'                  \
  |head -1                                         \
  |xargs -I{} curl -L "https://www.example.com/{}" \
  |pup 'script'                                    \
  |grep file:                                      \
  |sed -e "s/.*\(http[^ \"']*\).*/\1/"             \
# xargs vlc   # doesn't work
```

Is that you can comment out lines on the former one, but you can't do that on the latter. The `cat` trick makes it so that you have an 'exit' point, and you don't have to comment that one. Also, some editors will indent the first one correctly, while you'll have to manually indent the second one.

是您可以注释掉前一行的行，但不能对后者进行注释。 `cat` 技巧使得你有一个“退出”点，你不必评论那个点。此外，一些编辑器会正确缩进第一个，而您必须手动缩进第二个。

Small wins that compose just fine :)

组成很好的小胜利:)

### 5.7 Use ${var?error msg} on templates

### 5.7 在模板上使用 ${var?error msg}

If you write something to be copypasted by your user and filled in, instead of `<var>`, try `${var?You need to set var}`. it allows for the user to set the variable in the environment without having to replace inline, and if the user forgets any, the shell will barf.

如果你写的东西要被你的用户复制粘贴并填写，而不是`<var>`，试试`${var?You need to set var}`。它允许用户在环境中设置变量而不必替换内联，如果用户忘记了任何，shell 将 barf。

### 5.8 Idempotent functions

### 5.8 幂等函数

"My favourite shell scripting function definition technique: idempotent functions by redefining the function as a noop inside its body:

“我最喜欢的 shell 脚本函数定义技术：通过将函数重新定义为函数体内的 noop 来实现幂等函数：

(The `true;` body is needed by some shells, e.g. bash, and not others, e.g. zsh.)  " – [chrismorgan](https://news.ycombinator.com/item?id=27729397)

（`true;` 主体是某些 shell 所需要的，例如 bash，而不是其他 shell，例如 zsh。）” – [chrismorgan](https://news.ycombinator.com/item?id=27729397)

```
foo() {
    foo() { true;}
    echo "This bit will only be executed on the first foo call"
}
```

## 6 Debugging

## 6 调试

### 6.1 adding `bash` to a script to debug

### 6.1 将 `bash` 添加到脚本中进行调试

You can add `bash` inside any script, and it'll add a sort of a breakpoint, allowing you to check the state of the env and manually call functions around.

您可以在任何脚本中添加 `bash`，它会添加一种断点，允许您检查 env 的状态并手动调用周围的函数。

If you orgainse your code in small functions, it's easy to add breakpoints by just spawning bash processes inside your script.

如果您在小函数中组织代码，只需在脚本中生成 bash 进程即可轻松添加断点。

This works also inside docker containers (if you provide `-ti` flag on run).

这也适用于 docker 容器（如果您在运行时提供 `-ti` 标志）。

Let's see some usual uses of docker and how we can debug our scripts there:

让我们看看 docker 的一些常用用法以及我们如何在那里调试我们的脚本：

```
# leaves you at a shell to fiddle if all is in place after build
docker run -it mycomplex-image bash

# Runs /tmp/file.sh from the host inside.That's cool to make the
# container less hermetic.Even if the image is not originally ment
# to, you can even override it and 'monkeypatch' the file with the one
# from the host anyway.
docker run -it -v $PWD:/tmp/ mycomplex-image /tmp/file.sh

# So now you can really add wtv you want there.
echo 'bash' >>$PWD/file.sh

# run+open shell at runtime to inspect the state of the script
docker run -it -v $PWD:/tmp/ mycomplex-image /tmp/file.sh
```

### 6.2 DRY_RUN

### 6.2 DRY_RUN

```
if [[-n "$DRY_RUN" ]];then
  curl () {
    echo curl "$@"
  }
fi
```

use `command curl` to force the command, not the alias or anything

使用 `command curl` 强制命令，而不是别名或任何东西

### 6.3 Cheap debugging flag

### 6.3 便宜的调试标志

```
optargs "V" option;do
case $option in
  V)
    set -xa
  ;;
```

### 6.4 explore a pipe with tee >(some_command) |

### 6.4 使用 tee 探索管道 >(some_command) |

the `>()` is not very easy to use. Very few places where it fits. Here's a nice pipe inspector though, using `tee >(cat 1>&2)` trick.

`>()` 不是很容易使用。适合的地方很少。不过，这是一个不错的管道检查器，使用了`tee >(cat 1>&2)` 技巧。

```
plog() {
  # tee >(cat 1>&2)
  local msg=${1:-plog}
  tee >(sed -e "s/^/[$msg] /" 1>&2)
}
alias -g 'PL'=' |plog ' #zsh only

echo "a\nb" PL foo |tr 'a-z' 'A-Z' PL bar
# output:
# [foo] a     # stderr
# [foo] b     # stderr
# A           # stdout
# B           # stdout
# [bar] A     # stderr
# [bar] b     # stderr
```

ref: https://stackoverflow.com/questions/17983777/shell-pipe-to-multiple-commands-in-a-file

参考：https://stackoverflow.com/questions/17983777/shell-pipe-to-multiple-commands-in-a-file

### 6.5 tee+sudo

### 6.5 三通+须藤

If you wat to store a file in a root-owned dir, in the middle of your pipeline, instead of running the whole thing as root, you can use `sudo tee file`:

如果您想在管道中间的 root 拥有的目录中存储文件，而不是以 root 身份运行整个程序，您可以使用 `sudo tee file`：

```
ls |grep m >/usr/local/garbage  # fail
ls |grep m |sudo tee /usr/local/garbage # success!
```

### 6.6 quoting

### 6.6 引用

Bash: To get a quoted version of a given string, here's what you can do:

Bash：要获取给定字符串的引用版本，您可以执行以下操作：

```
# this is my "string" I want to 'comment "on"'
!:q
```

That gives us `'#this is my "string" I want to '\''comment   "on"'\'''`. Neat!

这给了我们`'#this is my "string" I want to '\''comment "on"'\'''`。整洁的！

I just found this trick [here](https://til.simonwillison.net/til/til/bash_escaping-a-string.md). From the associated HN thread:

我刚刚发现这个技巧 [here](https://til.simonwillison.net/til/til/bash_escaping-a-string.md)。从关联的 HN 线程：

```
function bashquote () {
  printf '%q' "$(cat)"
  echo
}
```

Zsh: If you're on zsh, `a-'` quotes the current line.

Zsh：如果你在 zsh 上，`a-'` 引用当前行。

## 7 zsh-only

## 7 zsh-only

### 7.1 Word spliting

### 7.1 分词

Word splitting works differently by default in zsh than in bash.

默认情况下，zsh 中分词的工作方式与 bash 中不同。

```
foo="ls -las"
$foo
```

This works in bash, but zsh will not split by words. To make zsh expand by words, there are 2 ways: `setopt SH_WORD_SPLIT` and `${=foo}`. zsh has `unsetop` command, which allows to scope where you want the expansions to happen. `unsetop SH_WORD_SPLIT`.

这适用于 bash，但 zsh 不会按单词拆分。要让zsh按词展开，有两种方式：`setopt SH_WORD_SPLIT`和`${=foo}`。 zsh 有 `unsetop` 命令，它允许在你想要扩展发生的地方确定范围。 `取消设置 SH_WORD_SPLIT`。

The problem with both solutions is that none of them are compatible with bash, so you'll be cornering yourself to "this only works in zsh". A way to overcome this is to use arrays, which are expanded in the same way in both shells.

两种解决方案的问题在于它们都不与 bash 兼容，因此您将陷入困境，“这仅适用于 zsh”。解决这个问题的一种方法是使用数组，它们在两个 shell 中以相同的方式扩展。

Or, use the same hack as you'll see later with noglob.

或者，使用您稍后将在 noglob 中看到的相同 hack。

Refs:

参考：

- https://stackoverflow.com/questions/6715388/variable-expansion-is-different-in-zsh-from-that-in-bash
- http://zsh.sourceforge.net/FAQ/zshfaq03.html#l18

- https://stackoverflow.com/questions/6715388/variable-expansion-is-different-in-zsh-from-that-in-bash
- http://zsh.sourceforge.net/FAQ/zshfaq03.html#l18

### 7.2 globbing

### 7.2 通配

In zsh, getting a list of files that match some characteristics is doable using globbing. Bash has globbing also, but in a less sophisticated way.

在 zsh 中，使用通配符可以获得匹配某些特征的文件列表。 Bash 也有 globbing，但方式不太复杂。

The basic structure of a `glob` is `pattern(qualifiers)`. Patterns can contain:

`glob` 的基本结构是 `pattern(qualifiers)`。模式可以包含：

- strings: they do exact match
- wildcards:  `*`, `?`, `**/`
- character classes: `[0-9]`
- choices:  `(.pdf|.djvu)`

- 字符串：它们完全匹配
- 通配符：`*`、`?`、`**/`
- 字符类：`[0-9]`
- 选择：`(.pdf|.djvu)`

The qualifiers are extra constraints you put on the matches. There are lots of different qualifiers. Look at `zshexpn` for the complete list. The ones I use more are:

预选赛是您对比赛施加的额外限制。有很多不同的限定词。查看“zshexpn”以获取完整列表。我用的比较多的是：

- `.` Files
- `/` Directories
- `om[numberhere]`. Nth latest modified

- `.` 文件
- `/` 目录
- `om[numberhere]`。第N次最新修改

### 7.3 Some global aliases:

### 7.3 一些全局别名：

These are some aliases I have in my ~/.zshrc that somehow help me use a shell in a more fluid way.

这些是我在 ~/.zshrc 中的一些别名，它们以某种方式帮助我以更流畅的方式使用 shell。

```
alias -g P1='|awk "{print \$1}"'
alias -g P2='|awk "{print \$2}"'
alias -g P3='|awk "{print \$3}"'
alias -g P4='|awk "{print \$4}"'
alias -g P5='|awk "{print \$5}"'
alias -g P6='|awk "{print \$6}"'
alias -g PL='|awk "{print \$NF}"'
alias -g PN='|awk "{print \$NF}"'
alias -g HL='|head -20'
alias -g H='|head '
alias -g H1='|head -1'
alias -g TL='|tail -20'
alias -g T='|tail '
alias -g T1='T -1'
#alias -g tr='-ltr'
alias -g X='|xclip  '
alias -g TB='|nc termbin.com 9999 '
alias -g L='|less -R '
alias -g LR='|less -r '
alias -g G='|grep '
alias -g GI='|grep -i '
alias -g GG=' 2>&1 |grep '
alias -g GGI=' 2>&1 |grep -i '
alias -g GV='|grep -v '
alias -g V='|grep -v '
alias -g TAC='|tac '
alias -g DU='du -B1'

alias -g E2O=' 2>&1 '
alias -g NE=' 2>/dev/null '
alias -g NO=' >/dev/null '

alias -g WC='|wc -l '

alias -g J='|noglob jq'
alias -g JQ='|noglob jq'
alias -g jq='noglob jq'
alias -g JL='|noglob jq -C .|less -R '
alias -g JQL='|noglob jq -C .|less -R '
alias -g XMEL='|xmlstarlet el'
alias -g XML='|xmlstarlet sel -t -v '

alias -g LYNX="| lynx -dump -stdin "
alias -g H2T="| html2text "
alias -g TRIM="| xargs "
alias -g XA='|xargs -d"\n" '
alias -g XE="| xargs e"
alias -g P="| pick "
alias -g PP="| percol | xargs "
alias -g W5="watch -n5 "
alias -g W1="watch -n1 "

alias -g CB="| col -b "
alias -g NC="| col -b "
alias -g U='|uniq '
alias -g XT='urxvt -e '
alias -g DM='|dmenu '
alias -g DMV='|dmenu -i -l 20 '

alias -g ...='../..'
alias -g ....='../../..'
alias -g .....='../../../..'

alias -g l10='*(om[1,10])'
alias -g l20='*(om[1,20])'
alias -g l5='*(om[1,5])'
alias -g l='*(om[1])'
alias -g '**.'='**/*(.)'
alias -g lpdf='*.pdf(om[1])'
alias -g lpng='*.png(om[1])'

alias -g u='*(om[1])'

alias lsmov='ls *.(mp4|mpg|mpeg|avi|mkv)'
alias lspdf='ls *.(pdf|djvu)'
alias lsmp3='ls *.mp3'
alias lspng='ls *.png'
```

Now, some sequences of words can start making sense:

现在，一些单词序列可以开始变得有意义了：

- `lspdf -tr TL DM XA evince`
- `docker exec -u root -ti $(docker ps -q H1) bash`
- `docker ps DM P1 XA docker stop`

-`lspdf -tr TL DM XA evince`
-`docker exec -u root -ti $(docker ps -q H1) bash`
- `docker ps DM P1 XA docker stop`

### 7.4 Expansion of global aliases

### 7.4 全局别名的扩展

Let's dig deeper.

让我们深入挖掘。

```
alias -g DOCK='docker ps 2..N P P1'
DOCK  # works fine.Echoes the id of the chosen container
docker stop DOCK   # does not work , because it expands to docker stop docker ps 2..N P P1
docker stop $(DOCK) # works fine again
alias -g DOCK='$(docker ps 2..N P P1)'
docker stop DOCK   # yay!
```

So, you can bind words to expansion-time results of the aliases. It feels like a very powerful thing, to have this "compile time" expansions. Reminds me of CL's symbol-macrolet, or IMMEDIATE Forth words.

因此，您可以将单词绑定到别名的扩展时间结果。拥有这种“编译时间”扩展，感觉是一件非常强大的事情。让我想起了 CL 的符号宏，或者 IMMEDIATE Forth 词。

### 7.5 Autocomplete

### 7.5 自动完成

Writting smart autocompletion scripts is not easy.

编写智能自动完成脚本并不容易。

zsh supports `compdef _gnu_generic` type of completion, which gets you very far with 0 effort.

zsh 支持 `compdef _gnu_generic` 类型的补全，这让你以 0 努力走得更远。

When autocompleting after a `-` in the commandline, if your command is configured like `compdef _gnu_generic mycommand`, zsh will call the script with `--help` and parse the output, trying to find flags, and will use them as suggestions . It's really great.

在命令行中的 `-` 后自动完成时，如果您的命令配置为类似 `compdef _gnu_generic mycommand`，zsh 将使用 `--help` 调用脚本并解析输出，尝试查找标志，并将它们用作建议.这真的很棒。

The compromise is to write a decent "–help" for your script. Which is cool because your user will love it too, and you just have to write it once.

妥协是为您的脚本编写一个像样的“-help”。这很酷，因为您的用户也会喜欢它，而您只需编写一次。

The completion is not context aware though, so you can't autocomplete flags after the first non-flag argument. It seems this could be improved in zsh-land, by asking for the –help like `mycommand args-so-far --help`. But it doesn't work like that.

但是，完成并不了解上下文，因此您无法在第一个非标志参数之后自动完成标志。似乎这可以在 zsh-land 中得到改进，通过请求像`mycommand args-so-far --help`这样的-help。但它不是那样工作的。

```
#!/usr/bin/env bash
# The script can be bash-only, while the completion work in zsh-only
set -Eeuo pipefail

help() {
  echo "  -h,--help    Show help"
  echo "  -c,--command Another thing"
}

if [ "$1" == "--help" ];
  help
fi
```

Now you can play with `mycommand -<TAB>`. Amazing, wow.

现在你可以玩`mycommand -<TAB>`。了不起，哇。

### 7.6 Create helpers and generate global aliases automagically

### 7.6 创建助手并自动生成全局别名

Borrowing a bit from Perl, a bit from Forth, and a bit from PicoLisp, I've come to create a few helpers that abstract words into a bit higher level concepts. Unifying the option selectors is one, and then, other line oriented operations like `chomp, from,   till`.

我借鉴了 Perl、Forth 和 PicoLisp 的一些知识，创建了一些帮助程序，将单词抽象为更高级的概念。统一选项选择器是一种，然后是其他面向行的操作，如“chomp、from、till”。

```
pick() {
  if [ -z "$DISPLAY" ];then
    percol ||fzf ||slmenu -i -l 20
  else
    dmenu -i -l 20
  fi
}
alias -g P='|pick'

globalias() {
  alias -g `echo -n $1 |tr '[a-z]' '[A-Z]'`=" | $1 "
}

globalias fzf

# uniquify column
function uc () {
  awk -F" " "!_[\$$1]++"
}
globalias uc

function from() { perl -pe "s|.*?$1|\1|"}
globalias from

function till() { sed -e "s|$1.*|$1|"}
globalias till

function chomp () { sed -e "s|.$||"}
globalias chomp
```

Again, it's a pity those do not compose well. Just be well organized, or build a more elaborate hack so you can compose aliases with some sort of confidence. It'll always be a hack though.

再一次，遗憾的是那些写得不好。只是组织得很好，或者构建一个更复杂的 hack，这样你就可以有信心地编写别名。不过，这将永远是一个黑客。

### 7.7 suffix aliases don't have to match a filename

### 7.7 后缀别名不必匹配文件名

zsh has another type of aliases called "suffix alias". Those alias allow you to define programs to open/run file types.

zsh 有另一种别名，称为“后缀别名”。这些别名允许您定义程序来打开/运行文件类型。

```
alias -s docx="libreoffice"
```

With this said, if you write a name of a file ending with `docx` as the first token in a command line, it will use libreoffice to open it.

话虽如此，如果您在命令行中写入以“docx”结尾的文件名作为第一个标记，它将使用 libreoffice 打开它。

```
invoice1.docx
# will effectively call libreoffice invoice1.docx
```

The trick here is that the parser doesn't check that the file is indeed an existing file. It can be any string.

这里的技巧是解析器不检查文件是否确实是现有文件。它可以是任何字符串。

Let's look at an example of it.

让我们看一个例子。

```
alias -s git="git clone"
```

In this case, we can easily copy a `git@github.com:.....git` from a browser, and paste it into a zsh console. Then, zsh will run that "file" with the command `git clone`, effectively cloning that repository.

在这种情况下，我们可以轻松地从浏览器中复制一个 `git@github.com:.....git`，并将其粘贴到 zsh 控制台中。然后，zsh 将使用命令 `git clone` 运行该“文件”，有效地克隆该存储库。

Cool, ain't it?

很酷，不是吗？

### 7.8 noglob

### 7.8 noglob

zsh has more aggressive parameter expansion, to the level that `[,],...` have special meanings, and will be interpreted and expanded before calling the final commands in your shell.

zsh 具有更激进的参数扩展，以至于`[,],...` 具有特殊含义，并且会在调用 shell 中的最终命令之前进行解释和扩展。

There are commands that you don't want ever expanded , for example, when using `curl`, it's much more likely that an open bracket will be ment to be there verbatim rather than expanded.

有些命令是你不想展开的，例如，当使用 `curl` 时，很可能会逐字显示一个开括号而不是展开。

Zsh provides a command to quote the following expansions. And it's called noglob.

Zsh 提供了一个命令来引用以下扩展。它被称为 noglob。

```
noglob curl http://example.com\&a[]=1
```

### 7.9 make noglob 'transparent' to bash

### 7.9 使 noglob“透明”到 bash

zsh and bash are mostly compatible, but there's a few things not supported in bash. `noglob` is one of them. To build a cushion inbetween, an easy way is to just create a `~/bin/noglob` file

zsh 和 bash 大多兼容，但 bash 不支持一些东西。 `noglob` 就是其中之一。要在两者之间建立一个缓冲，一个简单的方法是创建一个 `~/bin/noglob` 文件

```
$*
```

### 7.10 glob nested expansion

### 7.10 glob 嵌套扩展

In https://news.ycombinator.com/item?id=26175894 there's a nice advanced example:

在 https://news.ycombinator.com/item?id=26175894 中有一个很好的高级示例：

```
Variable expansion syntax, glob qualifiers, and history modifiers can
be combined/nested quite nicely.For example, this outputs all the
commands available from $PATH: `echo
${~${path/%/\/*(*N:t)}}`.`${~foo}` is to enable glob expansion on the
result of foo.`${foo/%/bar}` substitutes the end of the result of foo
to "bar" (i.e. it appends it);when foo is an array, it does it for
each element.In `/*(*N:t)`, we're adding the slash and star to the
paths from `$path`, then the parentheses are glob qualifiers.`*`
inside means only match the executables, `N` is to activate NULL_GLOB
for the match so that we don't get errors for globs that didn't match
anything, `:t` is a history mod used for globs that returns just the
"tail" of the result, i.e. the basename.IIRC, bash can't even nest
multiple parameter expansions;you need to save each step separately.
```

### 7.11 Some extra shortcuts for nice things

### 7.11 一些额外的好东西的捷径

- `alt-'` quotes the current line. It's like `quotemeta`. great to help you fight double and triple quoting when writing scripts.
- `alt-#` comment and accept. Nice way to store the current line for later.
- `ctrl-o` kill-current-line, wait for a command, and paste.

- `alt-'` 引用当前行。这就像`quotemeta`。非常有助于您在编写脚本时对抗双引号和三引号。
- `alt-#` 评论并接受。存储当前行以备后用的好方法。
- `ctrl-o` kill-current-line，等待命令，然后粘贴。

### 7.12 =()

### 7.12 =()

Zsh has `<()` and `>()` like Bash, but it also has `=()`. This varant is similar to `<()` but instead of creating a temporary pipe, it creates a temporary file. That is useful if we want to run commands that require a file instead of a pipe (most times, because it uses lseek to go through it).

Zsh 和 Bash 一样有 `<()` 和 `>()`，但它也有 `=()`。这个变量类似于`<()`，但它不是创建一个临时管道，而是创建一个临时文件。如果我们想运行需要文件而不是管道的命令（大多数情况下，因为它使用 lseek 来通过它），这很有用。

Node is an example of this.

Node 就是一个例子。

```
node <(echo 'setTimeout(() => console.log("foo"), 400)')   # fails
node =(echo 'setTimeout(() => console.log("foo"), 400)')   # works!
```

Or,

或者，

```
docker run --rm -ti -v =(echo "OHAI"):/tmp/foo ubuntu cat /tmp/foo
```

## 8 TODO patterns

## 8 个 TODO 模式

### 8.1 just use cat/netcat/pipes with <()

### 8.1 只使用带有 <() 的 cat/netcat/pipes

- input

   - 输入

  `python logger.py executable` will run the executable and monitor it for error messages. Depending on the error messages it will be doing.

   `python logger.py executable` 将运行可执行文件并监视它的错误消息。根据它将执行的错误消息。

  In order to test it, I want to run it with my own output. So what I do is `python logger.py cat`. That way I can type my stuff there, and even better, I can use a stream from the shell. `myexecutable | python logger.py cat` still works.

为了测试它，我想用我自己的输出运行它。所以我做的是`python logger.py cat`。这样我就可以在那里输入我的东西，甚至更好的是，我可以使用来自 shell 的流。 `我的可执行文件| python logger.py cat` 仍然有效。

#### 8.1.1 what's the unifying theory behind all that?

#### 8.1.1 这一切背后的统一理论是什么？

It's still not clear to me how they relate, but the feeling is that there's a common thread ruling all those commands. as if they generalize over the same things, or just a couple of very interrelated things.

我仍然不清楚它们之间的关系，但感觉是有一个共同的主线支配着所有这些命令。好像他们概括了相同的事情，或者只是几个非常相关的事情。

`echo` is to `cat` what `|` is to `xargs`. and `<()` and `>()` are able to make static files be dynamic streams. putting `cat` and `echo` inside `<()` seem like either a noop, or a leap in what can be done there. Still have to figure it out.

`echo` 之于 `cat` 就像 `|` 之于 `xargs`。和 `<()` 和 `>()` 能够使静态文件成为动态流。将 `cat` 和 `echo` 放在 `<()` 中似乎要么是一个空洞，要么是在那里可以做的事情的飞跃。还是得想办法。

<(grep a file.txt) , | xargs , cat, echo

<(grep a file.txt) , | xargs ，猫，回声

| you-have\it-wants | executable  | file                 | stream      |
| ----------------- |----------- |-------------------- |----------- |
| executable        | X           | <(exe)               | exe \| |
| file              | <(cat file) | X                    | cat file \| |
| stream            | cat         | <(grep foo file.txt) | X           |

|你有\它想要|可执行文件|档案 |流 |
| |可执行文件| X | <(exe) | exe \| |
|档案 | <（猫文件）| X |猫文件\| |
|流 |猫 | <(grep foo 文件.txt) | X |

- output

   - 输出

  Most of those can be tested with and `tee`. Sometimes you would like the output to be an output to a file to be extramassaged.

   大部分都可以用和`tee`来测试。有时您希望输出是要额外按摩的文件的输出。

  | you-have\it-wants | executable | file   | stream |
   | ----------------- |---------- |------ |------ |
  | executable        | X          | >()    | |
   | file              | | X      | |
   | stream            | | >(cat) | X      |

   |你有\它想要|可执行文件|档案 |流 |
  | |可执行文件| X | >() | |
  |档案 | | X | |
  |流 | | >(猫) | X |

  lnav <(tail -F /my/logfile-that-gets-rotated-or-truncated.log) cat <(date)

lnav <(tail -F /my/logfile-that-gets-rotated-or-truncated.log) cat <(日期)

### 8.2 redirects & streams

### 8.2 重定向和流

- https://catonmat.net/ftp/bash-redirections-cheat-sheet.pdf
- https://catonmat.net/bash-one-liners-explained-part-three
- https://github.com/miguelmota/bash-streams-handbook
- https://www2.dmst.aueb.gr/dds/sw/dgsh/
- https://wiki.bash-hackers.org/howto/redirection_tutorial

- https://catonmat.net/ftp/bash-redirections-cheat-sheet.pdf
- https://catonmat.net/bash-one-liners-explained-part-three
- https://github.com/miguelmota/bash-streams-handbook
- https://www2.dmst.aueb.gr/dds/sw/dgsh/
- https://wiki.bash-hackers.org/howto/redirection_tutorial

### 8.3 The $0 pattern

### 8.3 $0 模式

https://www.reddit.com/r/oilshell/comments/f6of85/four_more_posts_in_shell_the_good_parts/

https://www.reddit.com/r/oilshell/comments/f6of85/four_more_posts_in_shell_the_good_parts/

### 8.4 use git staging area to diff outputs of commands

### 8.4 使用 git staging area 来区分命令的输出

https://chrismorgan.info/blog/make-and-git-diff-test-harness/

https://chrismorgan.info/blog/make-and-git-diff-test-harness/

### 8.5 coprocs

### 8.5 coprocs

- https://stackoverflow.com/questions/7942632/how-to-extrace-pg-backend-pid-from-postgresql-in-shell-script-and-pass-it-to-ano/8305578#8305578
- https://unix.stackexchange.com/questions/86270/how-do-you-use-the-command-coproc-in-various-shells
- https://mbuki-mvuki.org/posts/2021-05-30-memoize-commands-or-bash-functions-with-coprocs/

- https://stackoverflow.com/questions/7942632/how-to-extrace-pg-backend-pid-from-postgresql-in-shell-script-and-pass-it-to-ano/8305578#8305578
- https://unix.stackexchange.com/questions/86270/how-do-you-use-the-command-coproc-in-various-shells
- https://mbuki-mvuki.org/posts/2021-05-30-memoize-commands-or-bash-functions-with-coprocs/

## 9 links

## 9 个链接

- https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html
- https://www.gnu.org/software/bash/manual/html_node/
- https://tldp.org/LDP/abs/html/
- https://mywiki.wooledge.org/BashPitfalls
- [Gary Bernhardt. The Unix Chainsaw](https://www.youtube.com/watch?v=sCZJblyT_XM)
- https://github.com/spencertipping/. This guy has some bash sick snippets
- https://news.ycombinator.com/item?id=23765123
- https://medium.com/@joydeepubuntu/functional-programming-in-bash-145b6db336b7
- https://www.youtube.com/watch?v=yD2ekOEP9sU
- http://catern.com/posts/pipes.html
- https://ebzzry.io/en/zsh-tips-1/
- https://github.com/ssledz/bash-fun
- https://news.ycombinator.com/item?id=24556022
- https://www.datafix.com.au/BASHing/index.html
- https://susam.github.io/tucl/the-unix-command-language.html 

- https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html
- https://www.gnu.org/software/bash/manual/html_node/
- https://tldp.org/LDP/abs/html/
- https://mywiki.wooledge.org/BashPitfalls
- [加里伯恩哈特。 Unix 链锯](https://www.youtube.com/watch?v=sCZJblyT_XM)
- https://github.com/spencertipping/。这家伙有一些 bash 生病的片段
- https://news.ycombinator.com/item?id=23765123
- https://medium.com/@joydeepubuntu/functional-programming-in-bash-145b6db336b7
- https://www.youtube.com/watch?v=yD2ekOEP9sU
- http://catern.com/posts/pipes.html
- https://ebzzry.io/en/zsh-tips-1/
- https://github.com/ssledz/bash-fun
- https://news.ycombinator.com/item?id=24556022
- https://www.datafix.com.au/BASHing/index.html
- https://susam.github.io/tucl/the-unix-command-language.html

- https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html
- https://www.grymoire.com/Unix/Sh.html
- https://github.com/dylanaraps/pure-sh-bible
- https://shatterealm.netlify.app/programming/2021_01_02_shiv_lets_build_a_vcs
- https://news.ycombinator.com/item?id=24401085
- https://git.sr.ht/~sircmpwn/shit
- [bocker](https://github.com/p8952/bocker). Docker implemented in around 100 lines of bash.
- https://github.com/simplenetes-io/simplenetes wow
- https://bakkenbaeck.github.io/a-random-walk-through-git/
- https://github.com/WeilerWebServices/Bash
- https://www.netmeister.org/blog/consistent-tools.html.

- https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html
- https://www.grymoire.com/Unix/Sh.html
- https://github.com/dylanaraps/pure-sh-bible
- https://shatterealm.netlify.app/programming/2021_01_02_shiv_lets_build_a_vcs
- https://news.ycombinator.com/item?id=24401085
- https://git.sr.ht/~sircmpwn/shit
- [博克](https://github.com/p8952/bocker)。 Docker 在大约 100 行 bash 中实现。
- https://github.com/simplenetes-io/simplenetes 哇
- https://bakkenbaeck.github.io/a-random-walk-through-git/
- https://github.com/WeilerWebServices/Bash
- https://www.netmeister.org/blog/consistent-tools.html。

## 10 From shell to lisp and everything in between

## 10 从 shell 到 lisp 以及介于两者之间的所有内容

- [Oil Shell](https://github.com/oilshell/oil).
- [Rash](https://rash-lang.org/) (Racket shell)
- [PaSh](https://arxiv.org/pdf/2007.09436.pdf): Light-touch Data-Parallel Shell Processing.
- [Painless emacs remote shells](https://www.eigenbahn.com/2020/07/08/painless-emacs-remote-shells). Because emacs has you covered
- https://news.ycombinator.com/item?id=24249646 rust
- https://github.com/liljencrantz/crush
- https://github.com/artyom-poptsov/metabash
- https://www.nushell.sh/
- [Babashka](https://github.com/borkdude/babashka)
- Bash to Perl/Python/Ruby using ```` and growing from there.

- [油壳](https://github.com/oilshell/oil)。
- [皮疹](https://rash-lang.org/)（球拍外壳)
- [PaSh](https://arxiv.org/pdf/2007.09436.pdf)：轻触数据并行外壳处理。
- [Painless emacs remote shells](https://www.eigenbahn.com/2020/07/08/painless-emacs-remote-shells)。因为 emacs 已经涵盖了你
- https://news.ycombinator.com/item?id=24249646 锈
- https://github.com/liljencrantz/crush
- https://github.com/artyom-poptsov/metabash
- https://www.nushell.sh/
- [Babashka](https://github.com/borkdude/babashka)
- Bash 到 Perl/Python/Ruby 使用 ```` 并从那里发展。

## 11 Credits

## 11 学分

- Raimon Grau <[raimonster@gmail.com](mailto:raimonster@gmail.com)>.
- Some examples are result of Raimon's and Lluís Esquerda's conversations or real world examples.
- people in https://news.ycombinator.com/item?id=24402571 which I'll be pulling in as time allows.

- 雷蒙·格劳 <[raimonster@gmail.com](mailto:raimonster@gmail.com)>。
- 一些示例是 Raimon 和 Lluís Esquerda 对话或现实世界示例的结果。
- https://news.ycombinator.com/item?id=24402571 中的人，我会在时间允许的情况下加入。

[1](https://raimonster.com/scripting-field-guide/#fnr.1)https://gist.github.com/samth/3083053

[1](https://raimonster.com/scripting-field-guide/#fnr.1)https://gist.github.com/samth/3083053

Author: Raimon Grau 

添加一名作者

