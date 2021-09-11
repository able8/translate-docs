# Taming the tar command: Tips for managing backups in Linux

# 驯服 tar 命令：在 Linux 中管理备份的技巧

Put tar to work creating and managing your backups smartly. Learn how tar can create, extract, append, split, verify integrity, and much more.

使用 tar 来巧妙地创建和管理您的备份。了解 tar 如何创建、提取、附加、拆分、验证完整性等等。

Posted:
September 18, 2020



Ever try something, it didn’t work, and you didn’t make a backup first?

曾经尝试过什么，它没有用，而且你没有先备份？

One of the key rules for working as a system administrator is _always_ to make a backup. You never know when you might need it. In my personal experience, it has saved me more times than I can count. It’s a common practice to complete and sometimes makes a difference in your finished work.

作为系统管理员工作的关键规则之一是_始终_进行备份。你永远不知道什么时候可能需要它。根据我的个人经验，它为我节省的时间多得数不过来。完成这是一种常见的做法，有时会对您完成的工作产生影响。

The `tar` utility has a ton of options and available usage. _Tar_ stands for _tape archive_ and allows you to create backups using: `tar`, `gzip`, and `bzip`. It compresses files and directories into an archive file, known as a _tarball_. This command is one of the most widely-used commands for this purpose. Also, the tarball is easily movable from one server to the next.

`tar` 实用程序有大量的选项和可用的用法。 _Tar_ 代表 _tape archive_ 并允许您使用：`tar`、`gzip` 和 `bzip` 创建备份。它将文件和目录压缩成一个存档文件，称为 _tarball_。此命令是为此目的使用最广泛的命令之一。此外，tarball 很容易从一台服务器移动到下一台服务器。

## How to create a tar backup

## 如何创建tar备份

In this example, we create a backup called _backup.tar_ of the directory `/home/user`.

在这个例子中，我们创建了一个名为 _backup.tar_ 的备份目录`/home/user`。

```shell
# tar -cvf backup.tar /home/user
```


Let’s break down these options:

让我们分解这些选项：

`-c` \- Create the archive

`-c` \- 创建存档

`-v` \- Show the process verbosely

`-v` \- 详细显示进程

`-f` \- Name the archive

`-f` \- 命名存档

## How to create a `tar.gz` backup

## 如何创建`tar.gz` 备份

In this example, we create a `gzip` archive backup called _backup.tar.gz_ of the directory `/home/user`.

在这个例子中，我们创建了一个名为 _backup.tar.gz_ 的 `gzip` 存档备份，位于目录 `/home/user` 中。

```shell
# tar -cvfz backup.tar.gz /home/user
```


Let’s break down these options:

让我们分解这些选项：

`-c` \- Create the archive

`-c` \- 创建存档

`-v` \- Show the process verbosely

`-v` \- 详细显示进程

`-f` \- Name the archive

`-f` \- 命名存档

`-z` \- Compressed gzip archive file

`-z` \- 压缩的 gzip 存档文件

## How to exclude files when creating a tar backup

## 如何在创建 tar 备份时排除文件

In this example, we create a `gzip` backup called _backup.tar.gz_, but exclude the files named _file.txt_ and _file.sh_ by using the `--exclude [filename]` option.

在本例中，我们创建了一个名为 _backup.tar.gz_ 的 `gzip` 备份，但使用 `--exclude [filename]` 选项排除名为 _file.txt_ 和 _file.sh_ 的文件。

```shell
# tar --exclude file.txt --exclude file.sh -cvfz backup.tar.gz
```


## How to extract content from a tar (.gz) backup

## 如何从 tar (.gz) 备份中提取内容

In this example, we extract content from a `gzip` backup _backup.tar.gz_, specifically a file called _file.txt_ from the directory `/backup/directory` in the `gzip` file.

在此示例中，我们从 `gzip` 备份 _backup.tar.gz_ 中提取内容，特别是从 `gzip` 文件中的目录 `/backup/directory` 中名为 _file.txt_ 的文件。

```shell
# tar -xvfz backup.tar.gz /backup/directory/file.txt
```


Let’s break down these options:

让我们分解这些选项：

`-x` \- Extract the content

`-x` \- 提取内容

`-v` \- Show the process verbosely

`-v` \- 详细显示进程

`-f` \- Name the archive

`-f` \- 命名存档

`-z` \- compressed gzip archive file

`-z` \- 压缩的 gzip 存档文件

## How to list contents of a tar(.gz) backup

## 如何列出 tar(.gz) 备份的内容

In this example, we list the contents from a `gzip` backup _backup.tar.gz_ without extracting it.

在这个例子中，我们列出了一个 `gzip` 备份 _backup.tar.gz_ 中的内容，而不提取它。

```shell
# tar -ztvf backup.tar.gz
```


Let’s break down these options:

让我们分解这些选项：

`-t` \- List the contents

`-t` \- 列出内容

`-v` \- Show the process verbosely

`-v` \- 详细显示进程

`-f` \- Name the archive

`-f` \- 命名存档

`-z` \- compressed gzip archive file

`-z` \- 压缩的 gzip 存档文件

## How to use the wildcard option

## 如何使用通配符选项

In this example, we use the wildcard option on a backup _backup.tar_. Wildcards allow you to select files without having a specific search for keywords. This is helpful in situations where you are trying to locate something but do not want to specify the name or want to add in all options matching that particular search.

在此示例中，我们在备份 _backup.tar_ 上使用通配符选项。通配符允许您选择文件而无需对关键字进行特定搜索。这在您尝试查找某些内容但不想指定名称或想要添加与该特定搜索匹配的所有选项的情况下很有帮助。

```shell
# tar -cf backup.tar “*.xml”
```


Let’s break down these options:

让我们分解这些选项：

`-c` \- Create the backup

`-c` \- 创建备份

`-f` \- Name the archive

`-f` \- 命名存档

## How to append or add files to a backup

## 如何将文件附加或添加到备份

In this example, we add onto a backup _backup.tar_. This allows you to add additional files to the pre-existing backup _backup.tar_.

在这个例子中，我们添加了一个备份 _backup.tar_。这允许您将其他文件添加到预先存在的备份 _backup.tar_。

```shell
# tar -rvf backup.tar /path/to/file.xml
```


Let’s break down these options:

让我们分解这些选项：

`-r` \- Append to archive

`-r` \- 附加到存档

`-v` \- Verbose output

`-v` \- 详细输出

`-f` \- Name the file

`-f` \- 命名文件

## How to split a backup into smaller backups

## 如何将备份拆分为较小的备份

In this example, we split the existing backup into smaller archived files. You can `pipe` the `tar` command into the `split` command.

在此示例中，我们将现有备份拆分为较小的存档文件。您可以通过管道将 `tar` 命令导入到 `split` 命令中。

```shell
# tar cvf - /dir |split --bytes=200MB - backup.tar
```


Let’s break down these options: 

让我们分解这些选项：

`-c` \- Create the archive

`-c` \- 创建存档

`-v` \- Verbose output

`-v` \- 详细输出

`-f` \- Name the file

`-f` \- 命名文件

In this example, the `dir/` is the directory that you want to split the backup content from. We are making 200MB backups from the `/dir` folder.

在此示例中，`dir/` 是您要从中拆分备份内容的目录。我们正在从 `/dir` 文件夹制作 200MB 的备份。

## How to check the integrity of a tar.gz backup

## 如何检查 tar.gz 备份的完整性

In this example, we check the integrity of an existing `tar` archive.

在这个例子中，我们检查现有 `tar` 存档的完整性。

To test the `gzip` file is not corrupt:

要测试 `gzip` 文件是否损坏：

```shell
#gunzip -t backup.tar.gz
```


To test the `tar` file content's integrity:

测试 `tar` 文件内容的完整性：

```shell
#gunzip -c backup.tar.gz |tar t > /dev/null
```


OR

```shell
#tar -tvWF backup.tar
```


Let’s break down these options:

让我们分解这些选项：

`-W` \- Verify an archive file

`-W` \- 验证存档文件

`-t` \- List files of archived file

`-t` \- 列出归档文件的文件

`-v` \- Verbose output

`-v` \- 详细输出

## Use pipes and greps to locate content

## 使用管道和 grep 来定位内容

In this example, we use `pipes` and `greps` to locate content. The best option is already made for you. `Zgrep` can be utilized for `gzip` archives.

在这个例子中，我们使用 `pipes` 和 `greps` 来定位内容。最好的选择已经为您准备好了。 `Zgrep` 可用于 `gzip` 档案。

```shell
#zgrep <keyword> backup.tar.gz
```


You can also use the `zcat` command. This shows the content of the archive, then `pipes` that output to a `grep`.

你也可以使用 `zcat` 命令。这显示了存档的内容，然后是输出到 `grep` 的 `pipes`。

```shell
#zcat backup.tar.gz |grep <keyword>
```


`Egrep` is a great one to use just for regular file types.

`Egrep` 是一个很好的工具，仅用于常规文件类型。

## Wrap up

##  总结

`Tar` has a lot of things you can do with it. It allows you to create the archive and manage it easily with the available tools in your terminal. If `tar` is not installed, you can do so depending on your operating system. `Tar` is useful in several different cases. As a system administrator, I created plenty of backups and recovered from some of them, too. It’s always safer to make a backup of a file or directory before making changes, in case you need to revert to the original setup. Having that security is something we all need.

`Tar` 有很多你可以用它做的事情。它允许您使用终端中的可用工具轻松创建档案并对其进行管理。如果没有安装`tar`，你可以根据你的操作系统安装。 `Tar` 在几种不同的情况下都很有用。作为系统管理员，我创建了大量备份并从中恢复了一些。在进行更改之前备份文件或目录总是更安全，以防您需要恢复到原始设置。拥有这种安全感是我们所有人都需要的。

_**[ Good backups are an important part of any security and disaster recovery plan. Want to learn more? Check out the [IT security and compliance checklist](https://www.redhat.com/en/resources/it-optimization-security-compliance-checklist?intcmp=701f20000012ngPAAQ). ]**_

_**[ 好的备份是任何安全和灾难恢复计划的重要组成部分。想了解更多？查看 [IT 安全和合规检查表](https://www.redhat.com/en/resources/it-optimization-security-compliance-checklist?intcmp=701f20000012ngPAAQ)。 ]**_

### Check out these related articles on Enable Sysadmin

### 查看有关启用系统管理员的这些相关文章



![Backup tips from the trenches](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

'0%200%20700%20420'%2F%3E)

[5 Linux backup and restore tips from the trenches](http://www.redhat.com/sysadmin/5-backup-tips)

 [来自战壕的5个Linux备份和恢复技巧](http://www.redhat.com/sysadmin/5-backup-tips)

Find out five backup and restore tips from someone who's been there, failed, and then succeeded.

从曾经去过那里、失败然后又成功的人那里找出五个备份和恢复技巧。

Posted:
March 26, 2020


发表：
2020 年 3 月 26 日


Author: [Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

作者：[Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

Image

图片

![Something went wrong and someone is to blame](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

20viewBox%3D'0%200%20700%20420'%2F%3E)

[When backups fail: A cautionary sysadmin tale](http://www.redhat.com/sysadmin/backups-cautionary-tale)

[当备份失败时：一个警告系统管理员的故事](http://www.redhat.com/sysadmin/backups-cautionary-tale)

When anything fails, fingers start to point. Here's one story of failed backups and inherited responsibility.

当任何事情失败时，手指开始指向。这是一个关于失败备份和继承责任的故事。

Posted:
August 11, 2020


发表：
2020 年 8 月 11 日


Author: [Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

作者：[Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

Image

图片

![When backups saved the day](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

'0%200%20700%20420'%2F%3E)

[Linux stories: When backups saved the day](http://www.redhat.com/sysadmin/backups-saved-day)

[Linux 故事：当备份挽救了一天](http://www.redhat.com/sysadmin/backups-saved-day)

Are your backups running too long? Would the time required for a full recovery put your business at risk? Here's a better solution.

您的备份运行时间是否过长？完全恢复所需的时间是否会使您的企业面临风险？这是一个更好的解决方案。

Posted:
May 7, 2020


发表：
2020 年 5 月 7 日


Author: [Jörg Kastning (Red Hat Accelerator, Sudoer)](http://www.redhat.com/sysadmin/users/joerg-kastning)

作者：[Jörg Kastning（红帽加速器，Sudoer）](http://www.redhat.com/sysadmin/users/joerg-kastning)

**Topics:** [Backups](http://www.redhat.com/sysadmin/topics/backups)[Linux](http://www.redhat.com/sysadmin/topics/linux)

**主题：** [备份](http://www.redhat.com/sysadmin/topics/backups)[Linux](http://www.redhat.com/sysadmin/topics/linux)

![Author’s photo](http://www.redhat.com/sysadmin/sites/default/files/styles/user_picture_square/public/pictures/2019-10/gabby_taylor.jpg?itok=wxnUhoxL)

## Gabby Taylor

## 加比泰勒

I currently work as a Manager of Content Support for Linux Academy. I have been working with Linux and OpenSource tools for a decade, constantly wanting to make new resolutions for obstacles and always training others on improving systems as a systems administrator.
[More about me](http://www.redhat.com/sysadmin/users/gabby-taylor)

我目前担任 Linux Academy 的内容支持经理。十年来，我一直在使用 Linux 和 OpenSource 工具，一直希望为障碍制定新的解决方案，并始终以系统管理员的身份培训其他人改进系统。
[更多关于我](http://www.redhat.com/sysadmin/users/gabby-taylor)

#### On Demand: Red Hat Summit 2021 Virtual Experience

#### 点播：2021 年红帽峰会虚拟体验

Relive our April event with demos, keynotes, and technical sessions from

通过演示、主题演讲和技术会议重温我们 4 月的活动

experts, all available on demand.

专家，均可按需提供。

[Watch Now](https://www.redhat.com/en/summit?intcmp=7013a0000026RhqAAE)

[立即观看](https://www.redhat.com/en/summit?intcmp=7013a0000026RhqAAE)

## Related Content 

##  相关内容

[Configure DNS with a Linux command, build a lab in five minutes, and more tips for sysadmins](http://www.redhat.com/sysadmin/top-sysadmin-articles-august-2021)

 [Linux 命令配置DNS，五分钟搭建实验室，更多系统管理员小贴士](http://www.redhat.com/sysadmin/top-sysadmin-articles-august-2021)

Check out Enable Sysadmin's top 10 articles from August 2021. 

查看 Enable Sysadmin 的 2021 年 8 月前 10 篇文章。

