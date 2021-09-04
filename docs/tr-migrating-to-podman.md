# [Migrating from Docker to Podman](https://qulogic.gitlab.io/posts/2019-10-20-migrating-to-podman/)

# [从Docker迁移到Podman](https://qulogic.gitlab.io/posts/2019-10-20-migrating-to-podman/)

2019-10-20

If you use [Docker](https://www.docker.com/), you may or may not have already heard of [Podman](https://podman.io/). It is an alternative container engine, and while I don’t have much knowledge of the details, there are a few reasons why I’m switching:

1. Podman runs in rootless mode, i.e., it does not need a daemon running as root;
2. Podman supports new things like cgroupsv2 (coming in Fedora 31);
3. Docker (actually moby-engine) is difficult to keep up-to-date in Fedora (which may correlate with point 2), and people seem to complain about this (though I’ve not cared too much.)

如果您使用 [Docker](https://www.docker.com/)，您可能听说过也可能没有听说过[Podman](https://podman.io/)。它是一个替代容器引擎，虽然我对细节知之甚少，但有几个原因让我切换：

1. Podman 以无根模式运行，即不需要以 root 身份运行的守护进程；
2. Podman 支持 cgroupsv2 之类的新东西（Fedora 31 中即将推出）；
3. Docker（实际上是 moby-engine）在 Fedora 中很难保持最新（这可能与第 2 点有关），人们似乎对此有所抱怨（尽管我并没有太在意。）

You can probably find many other posts declaring why you should or should not switch to Podman. I will not attempt to convince you here; this is instead a record of what I did to achieve the migration.

您可能会找到许多其他帖子，说明您为什么应该或不应该切换到 Podman。我不会在这里试图说服你；这是我为实现迁移所做的工作的记录。

# Requirements

#  要求

I don’t have very complicated requirements, and fortunately, Podman supports essentially the same CLI as the `docker` command, with some additions for pod orchestration.

我没有很复杂的需求，幸运的是，Podman 支持与 `docker` 命令基本相同的 CLI，但增加了一些 pod 编排。

The main complication is that storage moves from `/var` to `/home`. My `/var` is on `/` on an SSD, whereas `/home` is on spinning disk RAID. Thus with Docker, the images would be stored on a faster device than with Podman. In testing out Podman, I have pulled or built some images in its storage that I would like to keep as well. Fortunately, someone has had [this discussion in an issue already](https://github.com/containers/libpod/issues/1916), and all I had to do was implement it.

主要的复杂之处在于存储从`/var` 移动到`/home`。我的`/var` 在SSD 上的`/` 上，而`/home` 在旋转磁盘RAID 上。因此，使用 Docker，镜像将存储在比使用 Podman 更快的设备上。在测试 Podman 时，我在其存储中提取或构建了一些我也想保留的镜像。幸运的是，有人已经[在一个问题中讨论过](https://github.com/containers/libpod/issues/1916)，而我所要做的就是实现它。

# Migration

# 迁移

## Purge excess images

## 清除多余的镜像

First, I removed all intermediate or old images that were unnamed or untagged in Docker:

```bash
$ docker images
REPOSITORY       TAG               IMAGE ID      CREATED       SIZE
pidgin/builders  mingw-w64-x86_64  4b489a486f31  2 days ago    3.03GB
<none>           <none>            138cf2b0006c  2 days ago    2.3 GB
alpine           latest            961769676411  2 months ago  5.58MB
...

# Delete images that are <none> or that I don't need any more.
# Also, delete images that are already in Podman.
$ docker rmi 138cf2b0006c ...
```


And the same thing for `podman` (using the same commands), just to simplify moving storage. Note, you can achieve something similar using `prune`, but I opted to be a bit more surgical in what I removed, at least for now.

`podman` 也是如此（使用相同的命令），只是为了简化移动存储。请注意，您可以使用 `prune` 来实现类似的效果，但我选择对我删除的内容进行更多手术，至少现在是这样。

## Move storage

## 移动存储

I have a partition `/var/container`, which I use for container-like things without it gobbling up space on `/`, i.e., Conda, Docker, Mock, etc. So I moved the existing images I have there:

```bash
# Create a podman container location on the SSD.
$ sudo mkdir /var/container/podman

# Create a user-specific container path.
$ sudo mkdir /var/container/podman/1000
$ sudo chown 1000:1000 /var/container/podman/1000
# (Permissions are like /run/user/1000)
$ sudo chmod 0700 /var/container/podman/1000

# Move my images to the new location.
# (Even though this is a user directory, you need sudo because the container's
# filesystem may use other UIDs.)
$ sudo mv ~/.local/share/containers/* /var/container/podman/1000/

# Add SELinux equivalence;I'm not sure how necessary this is, but it is noted
# in the 'SELINUX LABELING' section of `man containers-storage.conf`.
$ sudo semanage fcontext -a -e $HOME/.local/share/containers \
      /var/container/podman/1000
$ sudo restorecon -R -v /var/container/podman/1000
```



Then edit `~/.config/containers/libpod.conf` and set the paths to the new location:

```
volume_path = "/var/container/podman/1000/storage/volumes"
static_dir = "/var/container/podman/1000/storage/libpod"
```


and in `~/.config/containers/storage.conf`, set its paths as well:

```
   graphroot = "/var/container/podman/1000/storage"
```


Finally, delete the libpod state database. This stores the old path and will whine about it being different. **Be careful with this!** I was only able to delete the file because I had no running containers, no pods, etc. and I checked its contents were safe to delete first.

最后，删除 libpod 状态数据库。这会存储旧路径，并且会抱怨它不同。 **小心这个！**我只能删除文件，因为我没有运行容器，没有 pods 等，我检查了它的内容是否可以安全删除。

```bash
# Be careful deleting this file!Check the caveat above.
$ rm /var/container/podman/1000/storage/libpod/bolt_state.db
```



Now, I can verify that everything was moved as it should be, and still works:

```bash
$ podman info |rg 1000
  GraphRoot: /var/container/podman/1000/storage
  RunRoot: /run/user/1000
  VolumePath: /var/container/podman/1000/storage/volumes

$ podman images
REPOSITORY                  TAG                IMAGE ID       CREATED        SIZE
docker.io/pidgin/builders   mingw-w64-x86_64   4b489a486f31   2 days ago     3.06 GB
docker.io/library/fedora    30                 9eff96f4c827   3 weeks ago    256 MB
...
```





## Move Docker images to Podman

## 将 Docker 镜像移动到 Podman

Since Podman has a `docker-daemon:` transport that fetches from Docker, this is a straightforward loop over the images in Docker:

```bash
$ for img in $(docker images --format '{{.Repository}}:{{.Tag}}');do
> echo $img
> podman pull docker-daemon:$img
> done
```



This covered all but two images. The first missing image is 0 bytes, and Docker doesn’t know how to export it (I also tried `docker save` which failed.) So instead, I just pulled this one from Docker Hub again.

这涵盖了除两个镜像之外的所有镜像。第一个丢失的镜像是 0 字节，Docker 不知道如何导出它（我也试过 `docker save` 失败了。）所以相反，我再次从 Docker Hub 中提取了这个。

The second missing image was one without a TAG. Podman isn’t able to parse this reference for whatever reason. A trick to transfer that one is to do:

```bash
$ docker save -o temp.tar image
$ podman pull docker-archive:temp.tar
```



*However*, this also failed with:

```
Error processing tar file(exit status 1):
  there might not be enough IDs available in the namespace (requested
  197608:197121 for /windows/etc/DIR_COLORS):
  lchown /windows/etc/DIR_COLORS: invalid argument
```


This is a [known limitation](https://www.redhat.com/sysadmin/rootless-podman) of how Podman works with user namespaces. I could attempt to work around it, but since I have access to the original builders of this image, I’ll just get it fixed instead.

这是 Podman 如何处理用户命名空间的[已知限制](https://www.redhat.com/sysadmin/rootless-podman)。我可以尝试解决它，但由于我可以访问此镜像的原始构建器，因此我将对其进行修复。

## Remove Docker

## 删除 Docker

Now, to clean up Docker, all its storage, and the package:

```bash
$ docker system prune -a --volumes
$ sudo rm -rf /var/container/docker  # There were 50G here!
$ sudo rm -rf /var/lib/docker /var/run/docker*

$ sudo systemctl stop docker.service
$ sudo systemctl disable docker.service

$ sudo dnf remove moby-engine
$ sudo rm -rf /etc/docker ~/.docker/
$ sudo groupdel docker  # This still exists because my user is in the group.
```



Finally, I also installed `podman-docker`, which provides a `docker` alias, to help out other projects/scripts that happen to run the `docker` command.

最后，我还安装了 `podman-docker`，它提供了一个 `docker` 别名，以帮助其他碰巧运行 `docker` 命令的项目/脚本。

# Gotchas

# 陷阱

Not everything worked perfectly out-of-the-box. I ran into a few bugs or misdocumented behaviours:

1. As noted earlier, there can be [limitations](https://www.redhat.com/sysadmin/rootless-podman) of UID/GID ranges in containers.
2. The [tutorial](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md) is confused about whether it wishes to be rootless or not. Almost [everything works with rootless](https://github.com/containers/libpod/issues/4250) there, but some bits do not, and since users don't share images or containers, you can't really mix the two.
3. Podman and Buildah images are shared, but their *containers* are not. The `podman build` command uses Buildah under the hood, and when builds are only partially completed, the latter’s containers are not deleted. This produces ‘sticky’ images in Podman: they are in `podman images`, but cannot be deleted because they’re in use, just not by Podman (i.e., invisible in `podman ps`). The fix is to delete the containers using the `buildah` command instead. The images can then be deleted with either command.

并非所有东西都能完美地开箱即用。我遇到了一些错误或错误记录的行为：

1. 如前所述，容器中的 UID/GID 范围可能存在 [限制](https://www.redhat.com/sysadmin/rootless-podman)。
2. [教程](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md) 对是否希望无根感到困惑。几乎 [一切都适用于 rootless](https://github.com/containers/libpod/issues/4250) 那里，但有些位没有，而且由于用户不共享镜像或容器，你不能真正混合二。
3. Podman 和 Buildah 镜像是共享的，但它们的*容器*不是。 `podman build` 命令在幕后使用 Buildah，当构建只是部分完成时，后者的容器不会被删除。这会在 Podman 中产生“粘性”镜像：它们在 `podman images` 中，但无法删除，因为它们正在使用中，只是 Podman 没有删除（即，在 `podman ps` 中不可见）。解决方法是使用 `buildah` 命令来删除容器。然后可以使用任一命令删除镜像。

Read other posts 

阅读其他帖子
