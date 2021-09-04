# [Migrating from Docker to Podman](https://qulogic.gitlab.io/posts/2019-10-20-migrating-to-podman/)

2019-10-20

If you use [Docker](https://www.docker.com/), you may or may not have already heard of [Podman](https://podman.io/). It is an alternative container engine, and while I don’t have much knowledge of the details, there are a few reasons why I’m switching:

1. Podman runs in rootless mode, i.e., it does not need a daemon running as root;
2. Podman supports new things like cgroupsv2 (coming in Fedora 31);
3. Docker (actually moby-engine) is difficult to keep up-to-date in Fedora (which may correlate with point 2), and people seem to complain about this (though I’ve not cared too much.)

You can probably find many other posts declaring why you should or should not switch to Podman. I will not attempt to convince you here; this is instead a record of what I did to achieve the migration.

# Requirements

I don’t have very complicated requirements, and fortunately, Podman supports essentially the same CLI as the `docker` command, with some additions for pod orchestration.

The main complication is that storage moves from `/var` to `/home`. My `/var` is on `/` on an SSD, whereas `/home` is on spinning disk RAID. Thus with Docker, the images would be stored on a faster device than with Podman. In testing out Podman, I have pulled or built some images in its storage that I would like to keep as well. Fortunately, someone has had [this discussion in an issue already](https://github.com/containers/libpod/issues/1916), and all I had to do was implement it.

# Migration

## Purge excess images

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

Copy

And the same thing for `podman` (using the same commands), just to simplify moving storage. Note, you can achieve something similar using `prune`, but I opted to be a bit more surgical in what I removed, at least for now.

## Move storage

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

# Add SELinux equivalence; I'm not sure how necessary this is, but it is noted
# in the 'SELINUX LABELING' section of `man containers-storage.conf`.
$ sudo semanage fcontext -a -e $HOME/.local/share/containers \
      /var/container/podman/1000
$ sudo restorecon -R -v /var/container/podman/1000
```

Copy

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

```bash
# Be careful deleting this file! Check the caveat above.
$ rm /var/container/podman/1000/storage/libpod/bolt_state.db
```

Copy

Now, I can verify that everything was moved as it should be, and still works:

```bash
$ podman info | rg 1000
  GraphRoot: /var/container/podman/1000/storage
  RunRoot: /run/user/1000
  VolumePath: /var/container/podman/1000/storage/volumes

$ podman images
REPOSITORY                  TAG                IMAGE ID       CREATED        SIZE
docker.io/pidgin/builders   mingw-w64-x86_64   4b489a486f31   2 days ago     3.06 GB
docker.io/library/fedora    30                 9eff96f4c827   3 weeks ago    256 MB
...
```

Copy

## Move Docker images to Podman

Since Podman has a `docker-daemon:` transport that fetches from Docker, this is a straightforward loop over the images in Docker:

```bash
$ for img in $(docker images --format '{{.Repository}}:{{.Tag}}'); do
> echo $img
> podman pull docker-daemon:$img
> done
```

Copy

This covered all but two images. The first missing image is 0 bytes, and Docker doesn’t know how to export it (I also tried `docker save` which failed.) So instead, I just pulled this one from Docker Hub again.

The second missing image was one without a TAG. Podman isn’t able to parse this reference for whatever reason. A trick to transfer that one is to do:

```bash
$ docker save -o temp.tar image
$ podman pull docker-archive:temp.tar
```

Copy

*However*, this also failed with:

```
Error processing tar file(exit status 1):
  there might not be enough IDs available in the namespace (requested
  197608:197121 for /windows/etc/DIR_COLORS):
  lchown /windows/etc/DIR_COLORS: invalid argument
```

This is a [known limitation](https://www.redhat.com/sysadmin/rootless-podman) of how Podman works with user namespaces. I could attempt to work around it, but since I have access to the original builders of this image, I’ll just get it fixed instead.

## Remove Docker

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

Copy

Finally, I also installed `podman-docker`, which provides a `docker` alias, to help out other projects/scripts that happen to run the `docker` command.

# Gotchas

Not everything worked perfectly out-of-the-box. I ran into a few bugs or misdocumented behaviours:

1. As noted earlier, there can be [limitations](https://www.redhat.com/sysadmin/rootless-podman) of UID/GID ranges in containers.
2. The [tutorial](https://github.com/containers/libpod/blob/master/docs/tutorials/podman_tutorial.md) is confused about whether it wishes to be rootless or not. Almost [everything works with rootless](https://github.com/containers/libpod/issues/4250) there, but some bits do not, and since users don’t share images or containers, you can’t really mix the two.
3. Podman and Buildah images are shared, but their *containers* are not. The `podman build` command uses Buildah under the hood, and when builds are only partially completed, the latter’s containers are not deleted. This produces ‘sticky’ images in Podman: they are in `podman images`, but cannot be deleted because they’re in use, just not by Podman (i.e., invisible in `podman ps`). The fix is to delete the containers using the `buildah` command instead. The images can then be deleted with either command.

Read other posts