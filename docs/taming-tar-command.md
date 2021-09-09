# Taming the tar command: Tips for managing backups in Linux

Put tar to work creating and managing your backups smartly. Learn how tar can create, extract, append, split, verify integrity, and much more.

Posted:
September 18, 2020

![newly constructed parking lot](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201000%20600'%2F%3E)

Photo by K HOWARD from Pexels

## More Linux resources

- [Advanced Linux Commands Cheat Sheet for Developers](https://developers.redhat.com/cheat-sheets/advanced-linux-commands/?intcmp=701f20000012ngPAAQ)
- [Get Started with Red Hat Insights](https://access.redhat.com/products/red-hat-insights/?intcmp=701f20000012ngPAAQ)
- [Download Now: Basic Linux Commands Cheat Sheet](https://developers.redhat.com/cheat-sheets/linux-commands-cheat-sheet/?intcmp=701f20000012ngPAAQ)
- [Linux System Administration Skills Assessment](https://rhtapps.redhat.com/assessment/?intcmp=701f20000012ngPAAQ)

Ever try something, it didn’t work, and you didn’t make a backup first?

One of the key rules for working as a system administrator is _always_ to make a backup. You never know when you might need it. In my personal experience, it has saved me more times than I can count. It’s a common practice to complete and sometimes makes a difference in your finished work.

The `tar` utility has a ton of options and available usage. _Tar_ stands for _tape archive_ and allows you to create backups using: `tar`, `gzip`, and `bzip`. It compresses files and directories into an archive file, known as a _tarball_. This command is one of the most widely-used commands for this purpose. Also, the tarball is easily movable from one server to the next.

## How to create a tar backup

In this example, we create a backup called _backup.tar_ of the directory `/home/user`.

```shell
# tar -cvf backup.tar /home/user
```

Let’s break down these options:

`-c` \- Create the archive

`-v` \- Show the process verbosely

`-f` \- Name the archive

## How to create a `tar.gz` backup

In this example, we create a `gzip` archive backup called _backup.tar.gz_ of the directory `/home/user`.

```shell
# tar -cvfz backup.tar.gz /home/user
```

Let’s break down these options:

`-c` \- Create the archive

`-v` \- Show the process verbosely

`-f` \- Name the archive

`-z` \- Compressed gzip archive file

## How to exclude files when creating a tar backup

In this example, we create a `gzip` backup called _backup.tar.gz_, but exclude the files named _file.txt_ and _file.sh_ by using the `--exclude [filename]` option.

```shell
# tar --exclude file.txt --exclude file.sh -cvfz backup.tar.gz
```

## How to extract content from a tar (.gz) backup

In this example, we extract content from a `gzip` backup _backup.tar.gz_, specifically a file called _file.txt_ from the directory `/backup/directory` in the `gzip` file.

```shell
# tar -xvfz backup.tar.gz /backup/directory/file.txt
```

Let’s break down these options:

`-x` \- Extract the content

`-v` \- Show the process verbosely

`-f` \- Name the archive

`-z` \- compressed gzip archive file

## How to list contents of a tar(.gz) backup

In this example, we list the contents from a `gzip` backup _backup.tar.gz_ without extracting it.

```shell
# tar -ztvf backup.tar.gz
```

Let’s break down these options:

`-t` \- List the contents

`-v` \- Show the process verbosely

`-f` \- Name the archive

`-z` \- compressed gzip archive file

## How to use the wildcard option

In this example, we use the wildcard option on a backup _backup.tar_. Wildcards allow you to select files without having a specific search for keywords. This is helpful in situations where you are trying to locate something but do not want to specify the name or want to add in all options matching that particular search.

```shell
# tar -cf backup.tar “*.xml”
```

Let’s break down these options:

`-c` \- Create the backup

`-f` \- Name the archive

## How to append or add files to a backup

In this example, we add onto a backup _backup.tar_. This allows you to add additional files to the pre-existing backup _backup.tar_.

```shell
# tar -rvf backup.tar /path/to/file.xml
```

Let’s break down these options:

`-r` \- Append to archive

`-v` \- Verbose output

`-f` \- Name the file

## How to split a backup into smaller backups

In this example, we split the existing backup into smaller archived files. You can `pipe` the `tar` command into the `split` command.

```shell
# tar cvf - /dir | split --bytes=200MB - backup.tar
```

Let’s break down these options:

`-c` \- Create the archive

`-v` \- Verbose output

`-f` \- Name the file

In this example, the `dir/` is the directory that you want to split the backup content from. We are making 200MB backups from the `/dir` folder.

## How to check the integrity of a tar.gz backup

In this example, we check the integrity of an existing `tar` archive.

To test the `gzip` file is not corrupt:

```shell
#gunzip -t backup.tar.gz
```

To test the `tar` file content's integrity:

```shell
#gunzip -c backup.tar.gz | tar t > /dev/null
```

OR

```shell
#tar -tvWF backup.tar
```

Let’s break down these options:

`-W` \- Verify an archive file

`-t` \- List files of archived file

`-v` \- Verbose output

## Use pipes and greps to locate content

In this example, we use `pipes` and `greps` to locate content. The best option is already made for you. `Zgrep` can be utilized for `gzip` archives.

```shell
#zgrep <keyword> backup.tar.gz
```

You can also use the `zcat` command. This shows the content of the archive, then `pipes` that output to a `grep`.

```shell
#zcat backup.tar.gz | grep <keyword>
```

`Egrep` is a great one to use just for regular file types.

## Wrap up

`Tar` has a lot of things you can do with it. It allows you to create the archive and manage it easily with the available tools in your terminal. If `tar` is not installed, you can do so depending on your operating system. `Tar` is useful in several different cases. As a system administrator, I created plenty of backups and recovered from some of them, too. It’s always safer to make a backup of a file or directory before making changes, in case you need to revert to the original setup. Having that security is something we all need.

_**[ Good backups are an important part of any security and disaster recovery plan. Want to learn more? Check out the [IT security and compliance checklist](https://www.redhat.com/en/resources/it-optimization-security-compliance-checklist?intcmp=701f20000012ngPAAQ). ]**_

### Check out these related articles on Enable Sysadmin

Image

![Backup tips from the trenches](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

[5 Linux backup and restore tips from the trenches](http://www.redhat.com/sysadmin/5-backup-tips)

Find out five backup and restore tips from someone who's been there, failed, and then succeeded.

Posted:
March 26, 2020


Author: [Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

Image

![Something went wrong and someone is to blame](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

[When backups fail: A cautionary sysadmin tale](http://www.redhat.com/sysadmin/backups-cautionary-tale)

When anything fails, fingers start to point. Here's one story of failed backups and inherited responsibility.

Posted:
August 11, 2020


Author: [Ken Hess (Red Hat)](http://www.redhat.com/sysadmin/users/khess)

Image

![When backups saved the day](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%20700%20420'%2F%3E)

[Linux stories: When backups saved the day](http://www.redhat.com/sysadmin/backups-saved-day)

Are your backups running too long? Would the time required for a full recovery put your business at risk? Here's a better solution.

Posted:
May 7, 2020


Author: [Jörg Kastning (Red Hat Accelerator, Sudoer)](http://www.redhat.com/sysadmin/users/joerg-kastning)

**Topics:** [Backups](http://www.redhat.com/sysadmin/topics/backups) [Linux](http://www.redhat.com/sysadmin/topics/linux)

![Author’s photo](http://www.redhat.com/sysadmin/sites/default/files/styles/user_picture_square/public/pictures/2019-10/gabby_taylor.jpg?itok=wxnUhoxL)

## Gabby Taylor

I currently work as a Manager of Content Support for Linux Academy. I have been working with Linux and OpenSource tools for a decade, constantly wanting to make new resolutions for obstacles and always training others on improving systems as a systems administrator.
[More about me](http://www.redhat.com/sysadmin/users/gabby-taylor)

#### On Demand: Red Hat Summit 2021 Virtual Experience

Relive our April event with demos, keynotes, and technical sessions from

experts, all available on demand.

[Watch Now](https://www.redhat.com/en/summit?intcmp=7013a0000026RhqAAE)

## Related Content


[Configure DNS with a Linux command, build a lab in five minutes, and more tips for sysadmins](http://www.redhat.com/sysadmin/top-sysadmin-articles-august-2021)

Check out Enable Sysadmin's top 10 articles from August 2021.
