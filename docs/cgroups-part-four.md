# Managing cgroups with systemd

Find out how much easier cgroup management is with systemd in this four-part series finale covering cgroups and resource management.

Posted:
October 9, 2020

![Managing cgroups with systemd](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201000%20600'%2F%3E)

Image by [R\_Winkelmann](https://pixabay.com/users/R_Winkelmann-6830448/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749) from [Pixabay](https://pixabay.com/?utm_source=link-attribution&utm_medium=referral&utm_campaign=image&utm_content=4567749)

In this final installment of my four-part cgroups article series, I cover cgroup integration with systemd. Be sure you also check out parts [one](https://redhat.com/sysadmin/cgroups-part-one), [two](https://redhat.com/sysadmin/cgroups-part-two), and [three](https://redhat.com/sysadmin/cgroups-part-three) in the series.

## Cgroups with systemd

## More Linux resources

- [Advanced Linux Commands Cheat Sheet for Developers](https://developers.redhat.com/cheat-sheets/advanced-linux-commands/?intcmp=701f20000012ngPAAQ)
- [Get Started with Red Hat Insights](https://access.redhat.com/products/red-hat-insights/?intcmp=701f20000012ngPAAQ)
- [Download Now: Basic Linux Commands Cheat Sheet](https://developers.redhat.com/cheat-sheets/linux-commands-cheat-sheet/?intcmp=701f20000012ngPAAQ)
- [Linux System Administration Skills Assessment](https://rhtapps.redhat.com/assessment/?intcmp=701f20000012ngPAAQ)

By default, systemd creates a new cgroup under the `system.slice` for each service it monitors. Going back to our OpenShift Control Plane host, running `systemd-cgls` shows the following services under the `system.slice` (output is truncated for brevity):

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

- Editing the service file itself.
- Using drop-in files.
- Using`systemctl set-property` commands, which are the same as manually editing the files, but `systemctl` creates the required entries for you.

I cover these in more detail below.

#### Editing service files

Let's edit the unit file itself. To do this, I created a very simple unit file which runs a script:

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

```plaintext
#!/bin/bash
/usr/bin/cat /dev/urandom > /dev/null &
```

When you examine the output of `systemd-cgls`, you see that our new service is nested under the `system.slice` (output truncated):

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

```plaintext
Slice=my-beautiful-slice.slice
```

The output of `systemd-cgls` shows something curious. The `cat.service` file is now deeply nested:

```plaintext
Control group /:
├─my.slice
│ └─my-beautiful.slice
│   └─my-beautiful-slice.slice
│     └─cat.service
│       └─4010 /usr/bin/cat /dev/urandom
```

Why is this? The answer has to do with the way that systemd interprets nested cgroups. Children are declared in the following fashion: `<parent>-<child>.slice</child></parent>`. Since systemd attempts to be helpful, if a parent does not exist, systemd creates it for you. If I had used underscores `_` instead of dashes `-` the result would have been what you would have expected:

```plaintext
Control group /:
├─my_beautiful_slice.slice
│ └─cat.service
│   └─4123 /usr/bin/cat /dev/urandom
```

#### Using drop-in files

Drop-in files for systemd are fairly trivial to set up. Start by making an appropriate directory based on your service's name in `/etc/systemd/system`. In the `cat` example, run the following command:

```shell
# mkdir -p /etc/systemd/system/cat.service.d/
```

These files can be organized any way you like them. They are actioned based on numerical order, so you should name your configuration files something like `10-CPUSettings.conf`. All files in this directory should have the file extension `.conf` and require you to run `systemctl daemon-reload` every time you adjust one of these files.

I have created two drop-in files to show how you can split out different configurations. The first is `00-slice.conf`. As seen below, it sets up the default options for a separate slice for the `cat` service:

```plaintext
[Service]
Slice=AWESOME.slice
MemoryAccounting=yes
CPUAccounting=yes
```

The other file sets the number of CPUShares, and it's called `10-CPUSettings.conf`.

```plaintext
[Service]
CPUShares=256
```

To show that this method works, I create a second service in the same slice. To make it easier to tell the processes apart, the second script is slightly different:

```shell
#!/bin/bash
/usr/bin/sha256sum /dev/urandom > /dev/null &
```

I then simply created copies of the `cat` files, replacing the script and changing the CPUShares value:

```shell
# sed 's/load\.sh/load2\.sh/g' cat.service > sha256sum.service
# cp -r cat.service.d sha256sum.service.d
# sed -i 's/256/2048/g' sha256sum.service.d/10-CPUSettings.conf
```

Finally, reload the daemon and start the services:

```shell
# systemctl daemon-reload
# systemctl start cat.service
# systemctl start sha256sum.service
```

_**[ Readers also liked: [What happens behind the scenes of a rootless Podman container?](https://www.redhat.com/sysadmin/behind-scenes-podman) ]**_

Instead of showing you the output from `top`, now is a good time to introduce you to `systemd-cgtop`. It works in a similar fashion to regular `top` except it gives you a breakdown per slice, and then again by services in each slice. This is very helpful in determining whether you are making good use of cgroups in general on your system. As seen below, `systemd-cgtop` shows both the aggregation for all services in a particular slice as part of the overall system and the resource utilization of each service in a slice:

Image

![ctop showing aggregation of services in a slice and resource utilization](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201232%20297'%2F%3E)

#### Using systemctl set-property

The last method that can be used to configure cgroups is the `systemctl set-property` command. I'll start with a basic service file `md5sum.service`:

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

After I have set up the unit file and reloaded the daemon, I use the `systemctl` command similar to the following:

```shell
# systemctl set-property md5sum.service CPUShares=1024
```

This creates a drop-in file for you located at `/etc/systemd/system.control/md5sum.service.d/50-CPUShares.conf`. Feel free to look at the files if you are curious as to their contents. As these files are not meant to be edited by hand, I won't spend any time on them.

You can test to see if the changes have taken effect by running:

```shell
systemctl start md5sum.service cat.service sha256sum.service
```

As you see in the screenshot below, the changes appear to be successful. `sha256sum.service` is configured for 2048 CPUShares, while `md5sum.service` has 1024. Finally, `cat.service` has 256.

Image

![ctop displaying different CPUShare configurations for different processes](data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D'http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg'%20viewBox%3D'0%200%201232%20165'%2F%3E)

_**[ Thinking about security? [Check out this free guide to boosting hybrid cloud security and protecting your business.](https://www.redhat.com/en/resources/hybrid-cloud-security-ebook?intcmp=701f20000012ngPAAQ) ]**_

## Wrap up

Hopefully, you learned something new throughout our journey together. There was a lot to tackle, and we barely even scratched the surface on what is possible with cgroups. Aside from the role that cgroups play in keeping your system healthy, they also play a part in a "defense-in-depth" strategy. Additionally, cgroups are a critical component for modern Kubernetes workloads, where they aid in the proper running of containerized processes. Cgroups are responsible for so many things, including:

- Limiting resources of processes.
- Deciding priorities when contentions do arise.
- Controlling access to read/write and mknod devices.
- Providing a high level of accounting for processes that are running on a system.

One could argue that containerization, Kubernetes, and a host of other business-critical implementations would not be possible without leveraging cgroups.

If you have any questions or comments or perhaps other article ideas, feel free to reach out to me on Twitter. I look forward to hearing all your feedback.
