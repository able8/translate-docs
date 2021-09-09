# Managing cgroups with systemd

# 使用 systemd 管理 cgroup

Find out how much easier cgroup management is with systemd in this four-part series finale covering cgroups and resource management.

在这个涵盖 cgroup 和资源管理的四部分系列结局中，了解使用 systemd 使 cgroup 管理变得多么容易。

Posted:
October 9, 2020

发表：
2020 年 10 月 9 日

![Managing cgroups with systemd](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201000%20600'%2F%3E)

0%200%201000%20600'%2F%3E)

Image by [R\_Winkelmann](https://pixabay.com/users/R_Winkelmann-6830448/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749)

图片由 [R\_Winkelmann](https://pixabay.com/users/R_Winkelmann-6830448/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749) 来自 [Pixabay](https://pixabay.com/?utm_source)=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749)

In this final installment of my four-part cgroups article series, I cover cgroup integration with systemd. Be sure you also check out parts [one](https://redhat.com/sysadmin/cgroups-part-one),[two](https://redhat.com/sysadmin/cgroups-part-two), and [three](https://redhat.com/sysadmin/cgroups-part-three) in the series.

在我的四部分 cgroups 文章系列的最后一部分中，我介绍了 cgroup 与 systemd 的集成。确保您还检查了部分 [one](https://redhat.com/sysadmin/cgroups-part-one)、[two](https://redhat.com/sysadmin/cgroups-part-two)和[三](https://redhat.com/sysadmin/cgroups-part-three) 系列。

## Cgroups with systemd

## Cgroups 与 systemd

## More Linux resources

## 更多 Linux 资源

- [Advanced Linux Commands Cheat Sheet for Developers](https://developers.redhat.com/cheat-sheets/advanced-linux-commands/?intcmp=701f20000012ngPAAQ)
- [Get Started with Red Hat Insights](https://access.redhat.com/products/red-hat-insights/?intcmp=701f20000012ngPAAQ)
- [Download Now: Basic Linux Commands Cheat Sheet](https://developers.redhat.com/cheat-sheets/linux-commands-cheat-sheet/?intcmp=701f20000012ngPAAQ)
- [Linux System Administration Skills Assessment](https://rhtapps.redhat.com/assessment/?intcmp=701f20000012ngPAAQ)

- [开发人员高级 Linux 命令备忘单](https://developers.redhat.com/cheat-sheets/advanced-linux-commands/?intcmp=701f20000012ngPAAQ)
- [红帽洞察入门](https://access.redhat.com/products/red-hat-insights/?intcmp=701f20000012ngPAAQ)
- [立即下载：基本 Linux 命令备忘单](https://developers.redhat.com/cheat-sheets/linux-commands-cheat-sheet/?intcmp=701f20000012ngPAAQ)
- [Linux 系统管理技能测评](https://rhtapps.redhat.com/assessment/?intcmp=701f20000012ngPAAQ)

By default, systemd creates a new cgroup under the `system.slice` for each service it monitors. Going back to our OpenShift Control Plane host, running `systemd-cgls` shows the following services under the `system.slice` (output is truncated for brevity):

默认情况下，systemd 会在`system.slice` 下为其监控的每个服务创建一个新的 cgroup。回到我们的 OpenShift 控制平面主机，运行 `systemd-cgls` 会在 `system.slice` 下显示以下服务（为简洁起见，输出被截断）：

```plaintext
└─system.slice
├─sssd.service
├─lvm2-lvmetad.service
├─rsyslog.service
├─systemd-udevd.service
├─systemd-logind.service
├─systemd-journald.service
├─crond.service
├─origin-node.service
├─docker.service
├─dnsmasq.service
├─tuned.service
├─sshd.service
├─NetworkManager.service
├─dbus.service
├─polkit.service
├─chronyd.service
├─auditd.service
└─getty@tty1.service
```


You can change this behavior by editing the systemd service file. There are three options with regard to cgroup management with systemd:

您可以通过编辑 systemd 服务文件来更改此行为。关于使用 systemd 进行 cgroup 管理的三个选项：

- Editing the service file itself.
- Using drop-in files.
- Using`systemctl set-property` commands, which are the same as manually editing the files, but `systemctl` creates the required entries for you.

- 编辑服务文件本身。
- 使用插入文件。
- 使用`systemctl set-property` 命令，与手动编辑文件相同，但`systemctl` 会为您创建所需的条目。

I cover these in more detail below.

我在下面更详细地介绍了这些。

#### Editing service files

#### 编辑服务文件

Let's edit the unit file itself. To do this, I created a very simple unit file which runs a script:

让我们编辑单元文件本身。为此，我创建了一个运行脚本的非常简单的单元文件：

```plaintext
[Service]
Type=oneshot
ExecStart=/root/generate_load.sh
TimeoutSec=0
StandardOutput=tty
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target
```


The bash script has only two lines:

bash 脚本只有两行：

```plaintext
#!/bin/bash
/usr/bin/cat /dev/urandom > /dev/null &
```


When you examine the output of `systemd-cgls`, you see that our new service is nested under the `system.slice` (output truncated):

当您检查 `systemd-cgls` 的输出时，您会看到我们的新服务嵌套在 `system.slice` 下（输出被截断）：

```plaintext
└─system.slice
├─cat.service
├─tuned.service
├─sshd.service
├─NetworkManager.service
├─sssd.service
├─dbus.service
│ └─getty@tty1.service
└─systemd-logind.service
```


What happens if I add the following line to the systemd service file?

如果我将以下行添加到 systemd 服务文件中会发生什么？

```plaintext
Slice=my-beautiful-slice.slice
```


The output of `systemd-cgls` shows something curious. The `cat.service` file is now deeply nested:

`systemd-cgls` 的输出显示了一些奇怪的东西。 `cat.service` 文件现在是深度嵌套的：

```plaintext
Control group /:
├─my.slice
│ └─my-beautiful.slice
│   └─my-beautiful-slice.slice
│     └─cat.service
│       └─4010 /usr/bin/cat /dev/urandom
```




Why is this? The answer has to do with the way that systemd interprets nested cgroups. Children are declared in the following fashion: `<parent>-<child>.slice</child></parent>`. Since systemd attempts to be helpful, if a parent does not exist, systemd creates it for you. If I had used underscores `_` instead of dashes `-` the result would have been what you would have expected:

为什么是这样？答案与 systemd 解释嵌套 cgroup 的方式有关。子项以以下方式声明：`<parent>-<child>.slice</child></parent>`。由于 systemd 试图提供帮助，如果父级不存在，systemd 会为您创建它。如果我使用下划线 `_` 而不是破折号 `-`，结果将是您所期望的：

```plaintext
Control group /:
├─my_beautiful_slice.slice
│ └─cat.service
│   └─4123 /usr/bin/cat /dev/urandom
```


#### Using drop-in files

#### 使用插入文件

Drop-in files for systemd are fairly trivial to set up. Start by making an appropriate directory based on your service's name in `/etc/systemd/system`. In the `cat` example, run the following command:

systemd 的插入文件设置起来相当简单。首先根据您的服务名称在 `/etc/systemd/system` 中创建一个适当的目录。在 `cat` 示例中，运行以下命令：

```shell
# mkdir -p /etc/systemd/system/cat.service.d/
```


These files can be organized any way you like them. They are actioned based on numerical order, so you should name your configuration files something like `10-CPUSettings.conf`. All files in this directory should have the file extension `.conf` and require you to run `systemctl daemon-reload` every time you adjust one of these files.

这些文件可以按您喜欢的任何方式组织。它们是根据数字顺序操作的，因此您应该将配置文件命名为“10-CPUSettings.conf”。此目录中的所有文件都应具有文件扩展名“.conf”，并且每次调整这些文件之一时都要求您运行“systemctl daemon-reload”。

I have created two drop-in files to show how you can split out different configurations. The first is `00-slice.conf`. As seen below, it sets up the default options for a separate slice for the `cat` service:

我创建了两个插入文件来展示如何拆分不同的配置。第一个是`00-slice.conf`。如下所示，它为 `cat` 服务的单独切片设置了默认选项：

```plaintext
[Service]
Slice=AWESOME.slice
MemoryAccounting=yes
CPUAccounting=yes
```


The other file sets the number of CPUShares, and it's called `10-CPUSettings.conf`.

另一个文件设置 CPUShares 的数量，它被称为 `10-CPUSettings.conf`。

```plaintext
[Service]
CPUShares=256
```


To show that this method works, I create a second service in the same slice. To make it easier to tell the processes apart, the second script is slightly different:

为了证明此方法有效，我在同一个切片中创建了第二个服务。为了更容易区分进程，第二个脚本略有不同：

```shell
#!/bin/bash
/usr/bin/sha256sum /dev/urandom > /dev/null &
```


I then simply created copies of the `cat` files, replacing the script and changing the CPUShares value:

然后我简单地创建了 `cat` 文件的副本，替换了脚本并更改了 CPUShares 值：

```shell
# sed 's/load\.sh/load2\.sh/g' cat.service > sha256sum.service
# cp -r cat.service.d sha256sum.service.d
# sed -i 's/256/2048/g' sha256sum.service.d/10-CPUSettings.conf
```


Finally, reload the daemon and start the services:

最后，重新加载守护进程并启动服务：

```shell
# systemctl daemon-reload
# systemctl start cat.service
# systemctl start sha256sum.service
```


_**[ Readers also liked: [What happens behind the scenes of a rootless Podman container?](https://www.redhat.com/sysadmin/behind-scenes-podman) ]**_

_**[ 读者还喜欢：[无根 Podman 容器的幕后发生了什么？](https://www.redhat.com/sysadmin/behind-scenes-podman) ]**_

Instead of showing you the output from `top`, now is a good time to introduce you to `systemd-cgtop`. It works in a similar fashion to regular `top` except it gives you a breakdown per slice, and then again by services in each slice. This is very helpful in determining whether you are making good use of cgroups in general on your system. As seen below, `systemd-cgtop` shows both the aggregation for all services in a particular slice as part of the overall system and the resource utilization of each service in a slice:

现在不是向您展示 `top` 的输出，而是向您介绍 `systemd-cgtop` 的好时机。它的工作方式与常规的“top”类似，除了它为您提供每个切片的细分，然后再按每个切片中的服务细分。这对于确定您是否在系统上充分利用了 cgroup 非常有帮助。如下所示，`systemd-cgtop` 显示了作为整个系统一部分的特定切片中所有服务的聚合以及切片中每个服务的资源利用率：

Image

图片

![ctop showing aggregation of services in a slice and resource utilization](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201232%20297'%2F%3E)

2Fsvg'%20viewBox%3D'0%200%201232%20297'%2F%3E)

#### Using systemctl set-property

#### 使用 systemctl set-property

The last method that can be used to configure cgroups is the `systemctl set-property` command. I'll start with a basic service file `md5sum.service`:

可用于配置 cgroup 的最后一种方法是 `systemctl set-property` 命令。我将从一个基本的服务文件 `md5sum.service` 开始：

```plaintext
[Service]
Type=oneshot
ExecStart=/root/generate_load3.sh
TimeoutSec=0
StandardOutput=tty
RemainAfterExit=yes
Slice=AWESOME.slice

[Install]
WantedBy=multi-user.target
```


Using the `systemctl set-property` command places the files in `/etc/systemd/system.control`. These files are not to be edited by hand. Not every property is recognized by the `set-property` command, so the `Slice` definition was put in the service file itself.

使用`systemctl set-property` 命令将文件放在`/etc/systemd/system.control` 中。这些文件不能手动编辑。 `set-property` 命令并不能识别每个属性，因此 `Slice` 定义被放在服务文件本身中。

After I have set up the unit file and reloaded the daemon, I use the `systemctl` command similar to the following:

设置单元文件并重新加载守护程序后，我使用类似于以下内容的 `systemctl` 命令：

```shell
# systemctl set-property md5sum.service CPUShares=1024
```


This creates a drop-in file for you located at `/etc/systemd/system.control/md5sum.service.d/50-CPUShares.conf`. Feel free to look at the files if you are curious as to their contents. As these files are not meant to be edited by hand, I won't spend any time on them.

这会为您创建一个位于 `/etc/systemd/system.control/md5sum.service.d/50-CPUShares.conf` 的插入文件。如果您对它们的内容感到好奇，请随意查看这些文件。由于这些文件不打算手动编辑，因此我不会在它们上花费任何时间。

You can test to see if the changes have taken effect by running:

您可以通过运行以下命令来测试更改是否生效：

```shell
systemctl start md5sum.service cat.service sha256sum.service
```




As you see in the screenshot below, the changes appear to be successful. `sha256sum.service` is configured for 2048 CPUShares, while `md5sum.service` has 1024. Finally, `cat.service` has 256.

正如您在下面的屏幕截图中看到的，更改似乎是成功的。 `sha256sum.service` 配置为 2048 个 CPUShares，而 `md5sum.service` 有 1024 个。最后，`cat.service` 有 256 个。

Image

图片

![ctop displaying different CPUShare configurations for different processes](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201232%20165'%2F%3E)

20viewBox%3D'0%200%201232%20165'%2F%3E)

_**[ Thinking about security? [Check out this free guide to boosting hybrid cloud security and protecting your business.](https://www.redhat.com/en/resources/hybrid-cloud-security-ebook?intcmp=701f20000012ngPAAQ) ]**_

_**[ 考虑安全？ [查看此免费指南，以提高混合云安全性和保护您的业务。](https://www.redhat.com/en/resources/hybrid-cloud-security-ebook?intcmp=701f20000012ngPAAQ) ]**_

## Wrap up

##  包起来

Hopefully, you learned something new throughout our journey together. There was a lot to tackle, and we barely even scratched the surface on what is possible with cgroups. Aside from the role that cgroups play in keeping your system healthy, they also play a part in a "defense-in-depth" strategy. Additionally, cgroups are a critical component for modern Kubernetes workloads, where they aid in the proper running of containerized processes. Cgroups are responsible for so many things, including:

希望您在我们一起的旅程中学到了一些新东西。有很多事情需要解决，我们几乎没有触及 cgroups 可能的表面。除了 cgroup 在保持系统健康方面发挥的作用之外，它们还在“深度防御”策略中发挥作用。此外，cgroup 是现代 Kubernetes 工作负载的关键组件，它们有助于容器化进程的正确运行。 Cgroup 负责很多事情，包括：

- Limiting resources of processes.
- Deciding priorities when contentions do arise.
- Controlling access to read/write and mknod devices.
- Providing a high level of accounting for processes that are running on a system.

- 限制进程的资源。
- 在确实出现争用时确定优先级。
- 控制对读/写和 mknod 设备的访问。
- 为系统上运行的进程提供高级别的会计。

One could argue that containerization, Kubernetes, and a host of other business-critical implementations would not be possible without leveraging cgroups.

有人可能会争辩说，如果不利用 cgroup，就不可能实现容器化、Kubernetes 和许多其他关键业务实现。

If you have any questions or comments or perhaps other article ideas, feel free to reach out to me on Twitter. I look forward to hearing all your feedback. 

如果您有任何问题或评论或其他文章想法，请随时在 Twitter 上与我联系。我期待听到您的所有反馈。

