# Using WireGuard on macOS via the CLI

# 通过 CLI 在 macOS 上使用 WireGuard

Published on 28 Jun 2021 ·
Filed in [Tutorial](http://blog.scottlowe.org/categories/tutorial) ·
874 words (estimated 5 minutes to read)

发表于 2021 年 6 月 28 日 ·
归档于 [教程](http://blog.scottlowe.org/categories/tutorial) ·
874 字（预计阅读 5 分钟）

I've written a few different posts on [WireGuard](https://www.wireguard.com), the “simple yet fast and modern VPN” (as described by the WireGuard web site) that aims to supplant tools like IPSec and OpenVPN. My [first post on WireGuard](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/) showed how to configure WireGuard on Linux, both on the client side as well as on the server side. After that, I followed it up with posts on [using the GUI WireGuard app to configure WireGuard on macOS](http://blog.scottlowe.org/2021/04/01/using-wireguard-on-macos/) and— most recently— [making WireGuard from Homebrew work on an M1-based Mac](http://blog.scottlowe.org/2021/06/22/making-wireguard-from-homebrew-work-on-an-m1-mac/). In this post, I’m going to take a look at using WireGuard on macOS again, but this time via the CLI.

我在 [WireGuard](https://www.wireguard.com) 上写了一些不同的帖子，“简单而快速且现代的 VPN”（如 WireGuard 网站所述)旨在取代 IPSec 和开放VPN。我的 [关于 WireGuard 的第一篇文章](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/) 展示了如何在 Linux 上配置 WireGuard，无论是在客户端以及服务器端。在那之后，我在[使用 GUI WireGuard 应用程序在 macOS 上配置 WireGuard](http://blog.scottlowe.org/2021/04/01/using-wireguard-on-macos/) 和-最近——[使 WireGuard from Homebrew 在基于 M1 的 Mac 上工作](http://blog.scottlowe.org/2021/06/22/making-wireguard-from-homebrew-work-on-an-m1-mac/)。在这篇文章中，我将再次在 macOS 上使用 WireGuard，但这次是通过 CLI。

Some of this information is also found in [this WireGuard quick start](https://www.wireguard.com/quickstart/). Here I’ll focus only on using macOS as a WireGuard client, not as a server; refer to the WireGuard docs (or to [my earlier post](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/)) for information on setting up a WireGuard server. I’ll also assume that you’ve installed WireGuard via [Homebrew](https://brew.sh).

其中一些信息也可以在 [此 WireGuard 快速入门](https://www.wireguard.com/quickstart/) 中找到。在这里，我将只关注将 macOS 用作 WireGuard 客户端，而不是用作服务器；有关设置的信息，请参阅 WireGuard 文档（或 [我之前的帖子](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/))启动 WireGuard 服务器。我还假设您已经通过 [Homebrew](https://brew.sh) 安装了 WireGuard。

## Generating Keys

## 生成密钥

The first step is to generate the public/private keys you’ll need. If the `/usr/local/etc/wireguard` (or the `/opt/homebrew/etc/wireguard` for users on an M1-based Mac) directory doesn’t exist, you’ll need to first create that directory. (It didn’t exist on my system.) Then, from that directory, run these commands:

第一步是生成您需要的公钥/私钥。如果“/usr/local/etc/wireguard”（或基于 M1 的 Mac 用户的“/opt/homebrew/etc/wireguard”）目录不存在，您需要先创建该目录。 （它在我的系统上不存在。）然后，从该目录运行以下命令：

```bash
umask <span style="color:#ae81ff">077</span>
wg genkey |tee privatekey |wg pubkey > publickey

```

This generates files named `privatekey` and `publickey`, the contents of which you’ll use in the configuration of WireGuard. Some tutorials indicate the need to become root or use `sudo`, but my configuration seems to work fine without either of those.

这会生成名为“privatekey”和“publickey”的文件，您将在 WireGuard 的配置中使用它们的内容。一些教程表明需要成为 root 用户或使用 `sudo`，但我的配置在没有这些的情况下似乎也能正常工作。

## Setting up a WireGuard Interface

## 设置 WireGuard 接口

Once you have public/private keys, you’re ready to set up the interface configuration. Note that you’ll also need the public key of the server (peer) system to which you’re connecting.

拥有公钥/私钥后，您就可以设置接口配置了。请注意，您还需要要连接的服务器（对等）系统的公钥。

Create a file named `wg0.conf` in `/usr/local/etc/wireguard` (that would be in `/opt/homebrew/etc/wireguard` for M1-based Macs). The contents of the file should look like this:

在`/usr/local/etc/wireguard`（对于基于M1 的Mac 将在`/opt/homebrew/etc/wireguard` 中）创建一个名为`wg0.conf` 的文件。该文件的内容应如下所示：

```toml
[Interface]
PrivateKey = <private-key-generated-earlier>
Address = <vpn-ip-address-for-this-interface>

[Peer]
PublicKey = <public-key-for-peer-system>
Endpoint = <public-ip-address>:51280
AllowedIPs = <vpn-ip-address-for-peer-interface>, <additional-cidrs>
PersistentKeepalive = 25
```

Most of this configuration file is pretty straightforward:

这个配置文件的大部分内容都非常简单：

- You’ll need to specify the private key for the local system plus the public key of the peer system.
- Each system will need “VPN IP addresses” that are routable/reachable from each other (I use a /29 CIDR from which I assign peer VPN IP addresses). The VPN IP addresses do not have to be publicly routable (they can be RFC 1918 ranges), but they should not overlap other IP address ranges—otherwise you’ll run into routing table issues.
- For the`AllowedIPs` setting, you need to specify the VPN IP address of the peer _plus_ any additional networks that are reachable via the peer. In my case, I use WireGuard to access private EC2 instances in an AWS VPC, so I specify the VPC CIDR here.

- 您需要指定本地系统的私钥加上对等系统的公钥。
- 每个系统都需要相互可路由/可达的“VPN IP 地址”（我使用 /29 CIDR，从中分配对等 VPN IP 地址）。 VPN IP 地址不必是可公开路由的（它们可以是 RFC 1918 范围），但它们不应与其他 IP 地址范围重叠——否则您会遇到路由表问题。
- 对于`AllowedIPs` 设置，您需要指定对等方_加上_可通过对等方访问的任何其他网络的VPN IP 地址。就我而言，我使用 WireGuard 访问 AWS VPC 中的私有 EC2 实例，因此我在此处指定了 VPC CIDR。

The peer system will have/need a similar configuration. As described [here](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/), my primary use case is enabling connectivity to EC2 instances with private IP addresses inside a VPC, so the peer system for me is a Linux instance with WireGuard installed and configured.

对等系统将具有/需要类似的配置。如 [此处](http://blog.scottlowe.org/2021/02/22/setting-up-wireguard-for-aws-vpc-access/) 所述，我的主要用例是启用与私有 EC2 实例的连接VPC 内的 IP 地址，因此对我而言，对等系统是安装并配置了 WireGuard 的 Linux 实例。

## Activating the VPN 

## 激活 VPN

Once you have the WireGuard interface configuration finished, you can activate the VPN connection using `wg-quick up wg0` (if you called your configuration file from the previous section something _other_ than `wg0.conf`, then change the command accordingly). If the peer system is already configured and its interface is up, then the VPN connection should establish automatically and you _should_ be able to start routing traffic through the peer (assuming you specified some additional networks to be routed through the peer).

完成 WireGuard 接口配置后，您可以使用 `wg-quick up wg0` 激活 VPN 连接（如果您从上一节中调用了配置文件 _other_ 而不是 `wg0.conf`，则相应地更改命令）。如果对等系统已配置且其接口已启动，则 VPN 连接应自动建立，并且您_应该_能够开始通过对等路由流量（假设您指定了一些要通过对等路由的其他网络）。

Use `wg-quick down wg0` to take the VPN connection down.

使用 `wg-quick down wg0` 关闭 VPN 连接。

I did see some references to being able to use `launchd` on macOS to automatically start the WireGuard interface. It took a bit of testing and exploration, but I eventually settled on [this configuration](http://blog.scottlowe.org/2021/08/04/starting-wireguard-interfaces-automatically-launchd-macos/).

我确实看到了一些关于能够在 macOS 上使用“launchd”来自动启动 WireGuard 界面的参考资料。经过一些测试和探索，我最终决定了[这个配置](http://blog.scottlowe.org/2021/08/04/starting-wireguard-interfaces-automatically-launchd-macos/)。

## Troubleshooting the Connection

## 连接故障排除

If you run into problems getting the connection up and working, here are a few things you can check:

如果您在连接和工作时遇到问题，您可以检查以下几项内容：

- When specifying`AllowedIPs`, make sure to include the VPN interface IP address as well as any additional networks. Failing to include the VPN interface IP address can cause the connection to fail.
- Double-check that you’ve specified public and private keys correctly. It’s easy to get these mixed up. Each system’s configuration should reference its own private key and the peer’s public key.
- Make sure you’ve specified the peer’s endpoint correctly.
- Make sure that the traffic is being allowed to reach each system. Double-check firewalls, network access control lists, host-based firewalls, and security groups.
- Make sure the WireGuard interface is up and active on both systems. When troubleshooting connections, it can be easy to forget whether the interface is active or not.

- 在指定`AllowedIPs` 时，请确保包括 VPN 接口 IP 地址以及任何其他网络。未包含 VPN 接口 IP 地址可能会导致连接失败。
- 仔细检查您是否正确指定了公钥和私钥。很容易混淆这些。每个系统的配置都应该引用自己的私钥和对等方的公钥。
- 确保您已正确指定对等方的端点。
- 确保允许流量到达每个系统。仔细检查防火墙、网络访问控制列表、基于主机的防火墙和安全组。
- 确保两个系统上的 WireGuard 接口都已启动并处于活动状态。在排除连接故障时，很容易忘记接口是否处于活动状态。

Aside from the location of the configuration files, configuring WireGuard via the CLI using configuration files is very, very similar across systems. The interface configuration shown above is—as far as I’ve been able to tell so far—identical across both macOS and Linux.

除了配置文件的位置之外，使用配置文件通过 CLI 配置 WireGuard 非常非常相似。上面显示的界面配置——据我目前所知——在 macOS 和 Linux 上都是相同的。

I hope this article is helpful. If you have any questions, comments, or corrections, please feel free to contact me. You can reach [me on Twitter](https://twitter.com/scott_lowe) (DMs are open), or hit me up on one of a variety of Slack communities. I’d be happy to chat with you.

我希望这篇文章有帮助。如果您有任何问题、意见或更正，请随时与我联系。你可以在 Twitter 上联系 [我](https://twitter.com/scott_lowe)（DM是开放的)，或者在各种 Slack 社区之一上联系我。我很乐意和你聊天。

### Metadata and Navigation

### 元数据和导航

[CLI](http://blog.scottlowe.org/tags/cli)[Encryption](http://blog.scottlowe.org/tags/encryption) [macOS](http://blog.scottlowe.org/tags/macos) [Networking](http://blog.scottlowe.org/tags/networking)[VPN](http://blog.scottlowe.org/tags/vpn)

[CLI](http://blog.scottlowe.org/tags/cli)[加密](http://blog.scottlowe.org/tags/encryption) [macOS](http://blog.scottlowe.org/标签/macos)[网络](http://blog.scottlowe.org/tags/networking) [VPN](http://blog.scottlowe.org/tags/vpn)

Previous Post: [Installing Older Versions of Kumactl on an M1 Mac](https://blog.scottlowe.org/2021/06/25/installing-older-versions-of-kumactl-on-an-m1-mac/)

上一篇：[在 M1 Mac 上安装旧版本的 Kumactl](https://blog.scottlowe.org/2021/06/25/installing-older-versions-of-kumactl-on-an-m1-mac/)

Next Post: [Adding Multiple Items Using Kustomize JSON 6902 Patches](https://blog.scottlowe.org/2021/07/07/adding-multiple-items-using-kustomize-json-6902-patches/)
Be social and share this post!

下一篇： [使用 Kustomize JSON 6902 补丁添加多个项目](https://blog.scottlowe.org/2021/07/07/adding-multiple-items-using-kustomize-json-6902-patches/)
社交并分享这篇文章！

[Share on Facebook](https://www.facebook.com/sharer/sharer.php?u=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing-wireguard-on-mac- via-cli%2f "Share on Facebook")[Share on Twitter](https://twitter.com/intent/tweet?url=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing -wireguard-on-mac-via-cli%2f&text=Using%20WireGuard%20on%20macOS%20via%20the%20CLI "Share on Twitter")[Share on Google Plus](https://plus.google.com/share ?url=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing-wireguard-on-mac-via-cli%2f "Share on Google Plus")

[在 Facebook 上分享](https://www.facebook.com/sharer/sharer.php?u=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing-wireguard-on-mac- via-cli%2f "在 Facebook 上分享")[在 Twitter 上分享](https://twitter.com/intent/tweet?url=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing -wireguard-on-mac-via-cli%2f&text=Using%20WireGuard%20on%20macOS%20via%20the%20CLI“在 Twitter 上分享”)[在 Google Plus 上分享](https://plus.google.com/share ?url=https%3a%2f%2fblog.scottlowe.org%2f2021%2f06%2f28%2fusing-wireguard-on-mac-via-cli%2f“在 Google Plus 上分享”)

### Related Posts

###  相关文章

- [Making WireGuard from Homebrew Work on an M1 Mac](http://blog.scottlowe.org/2021/06/22/making-wireguard-from-homebrew-work-on-an-m1-mac/) 22 Jun 2021
- [Using WireGuard on macOS](http://blog.scottlowe.org/2021/04/01/using-wireguard-on-macos/) 1 Apr 2021
- [Technology Short Take 136](http://blog.scottlowe.org/2021/01/15/technology-short-take-136/) 15 Jan 2021 

- [使用 Homebrew 在 M1 Mac 上制作 WireGuard](http://blog.scottlowe.org/2021/06/22/making-wireguard-from-homebrew-work-on-an-m1-mac/) 6 月 22 日2021年
- [在 macOS 上使用 WireGuard](http://blog.scottlowe.org/2021/04/01/using-wireguard-on-macos/) 2021 年 4 月 1 日
- [技术短片 136](http://blog.scottlowe.org/2021/01/15/technology-short-take-136/) 2021 年 1 月 15 日

