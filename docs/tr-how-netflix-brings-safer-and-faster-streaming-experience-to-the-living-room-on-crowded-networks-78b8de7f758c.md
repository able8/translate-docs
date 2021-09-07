# How Netflix brings safer and faster streaming experiences to the living room on crowded networks using TLS 1.3

# Netflix 如何使用 TLS 1.3 在拥挤的网络上为客厅带来更安全、更快的流媒体体验

[Netflix Technology Blog](https://netflixtechblog.medium.com/?source=post_page-----78b8de7f758c--------------------------------)

[Apr 20, 2020](http://netflixtechblog.com/how-netflix-brings-safer-and-faster-streaming-experience-to-the-living-room-on-crowded-networks-78b8de7f758c?source=post_page-----78b8de7f758c--------------------------------)·5 min read

At Netflix, we are obsessed with the best streaming experiences. We want playback to start instantly and to never stop unexpectedly in any network environment. We are also committed to protecting users’ privacy and service security without sacrificing any part of the playback experience.

[2020 年 4 月 20 日](http://netflixtechblog.com/how-netflix-brings-safer-and-faster-streaming-experience-to-the-living-room-on-crowded-networks-78b8de7f758c?source=post_page 在 Netflix，我们痴迷于最佳的流媒体体验。我们希望播放立即开始，并且在任何网络环境中都不会意外停止。我们还致力于在不牺牲任何部分播放体验的情况下保护用户的隐私和服务安全。

To achieve that, we are efficiently using ABR (adaptive bitrate streaming) for a better playback experience, DRM (Digital Right Management) to protect our service and TLS (Transport Layer Security) to protect customer privacy and to create a safer streaming experience.

为实现这一目标，我们有效地使用 ABR（自适应比特率流媒体）来获得更好的播放体验，使用 DRM（数字版权管理）来保护我们的服务，并使用 TLS（传输层安全）来保护客户隐私并创造更安全的流媒体体验。

Netflix on consumer electronics devices such as TVs, set-top boxes and streaming sticks was until recently using TLS 1.2 for streaming traffic. Now we support TLS 1.3 for safer and faster experiences.

Netflix 在电视、机顶盒和流媒体棒等消费电子设备上直到最近才使用 TLS 1.2 来处理流媒体流量。现在我们支持 TLS 1.3，以获得更安全、更快的体验。

# What is TLS?

# 什么是 TLS？

For two parties to communicate securely, a secure channel is necessary. This needs to have the following three properties.

对于两方安全通信，安全通道是必要的。这需要具备以下三个属性。

- Authentication: Identity of the communicating party is verified.
- Confidentiality: Data sent over the channel is only visible to the endpoints.
- Integrity: Data sent over the channel cannot be modified by attackers without detection.

- 认证：验证通信方的身份。
- 机密性：通过通道发送的数据仅对端点可见。
- 完整性：通过通道发送的数据不能被攻击者修改而不被发现。

The TLS protocol is designed to provide a secure channel between two peers by providing tools and methods to achieve the above properties.

TLS 协议旨在通过提供工具和方法来实现上述属性，从而在两个对等方之间提供安全通道。

# TLS 1.3

# TLS 1.3

TLS 1.3 is the latest version of the Transport Layer Security protocol. It is simpler, more secure and more efficient than its predecessor.

TLS 1.3 是传输层安全协议的最新版本。它比其前身更简单、更安全、更高效。

## Perfect Forward Secrecy

## 完美的前向保密

One thing we believe is very important at Netflix is providing PFS (Perfect Forward Secrecy).

我们认为在 Netflix 非常重要的一件事是提供 PFS（完美前向保密）。

PFS is a feature of the key exchange algorithm that assures that session keys will not be compromised, even if the server’s private key is compromised. By generating new keys for each session, PFS protects past sessions against the future compromise of secret keys.

PFS 是密钥交换算法的一个特性，可确保会话密钥不会被泄露，即使服务器的私钥被泄露。通过为每个会话生成新密钥，PFS 可以保护过去的会话免受未来密钥泄露的影响。

TLS 1.2 supports key exchange algorithms with PFS, but it also allows key exchange algorithms that do not support PFS. Even with the previous version of TLS 1.2, Netflix has always selected a key exchange algorithm that provides PFS such as ECDHE (Elliptic Curve Diffie Hellman Ephemeral). TLS 1.3, however, enforces this concept even more by removing all the key exchange algorithms that do not provide PFS, such as static RSA.

TLS 1.2 支持与 PFS 的密钥交换算法，但它也允许不支持 PFS 的密钥交换算法。即使在之前的 TLS 1.2 版本中，Netflix 也一直选择提供 PFS 的密钥交换算法，例如 ECDHE（Elliptic Curve Diffie Hellman Ephemeral）。然而，TLS 1.3 通过删除所有不提供 PFS 的密钥交换算法（例如静态 RSA），进一步强化了这一概念。

## Authenticated Encryption

## 认证加密

For encryption, TLS 1.3 removes all weak ciphers and uses only Authenticated Encryption with Associated Data (AEAD). This assures the confidentiality, integrity, and authenticity of the data. We use AES Galois/Counter Mode, as it also provides good performance and high throughput.

对于加密，TLS 1.3 删除了所有弱密码并仅使用具有关联数据的身份验证加密 (AEAD)。这确保了数据的机密性、完整性和真实性。我们使用 AES Galois/Counter Mode，因为它也提供了良好的性能和高吞吐量。

## Secure Handshake

## 安全握手

While the above changes are important, the most important change in TLS 1.3 is perhaps its redesign of the handshake protocol.

虽然上述更改很重要，但 TLS 1.3 中最重要的更改可能是对握手协议的重新设计。

The TLS 1.2 handshake was not designed to protect the integrity of the entire handshake. It protected only the part of the handshake after the cipher suite negotiation and this opened up the possibility of downgrade attacks which may allow the attackers to force the use of insecure cipher suites.

TLS 1.2 握手并非旨在保护整个握手的完整性。它仅保护密码套件协商后的握手部分，这开启了降级攻击的可能性，这可能允许攻击者强制使用不安全的密码套件。

With TLS 1.3, the server signs the entire handshake including the cipher suite negotiation and thus prevents the attacker from downgrading the cipher suite.

使用 TLS 1.3，服务器对包括密码套件协商在内的整个握手进行签名，从而防止攻击者降级密码套件。

Also in TLS 1.2, extensions were sent in the clear in the ServerHello. Now with TLS 1.3, even extensions are encrypted and all handshake messages after ServerHello are now encrypted.

同样在 TLS 1.2 中，扩展在 ServerHello 中以明文形式发送。现在使用 TLS 1.3，甚至扩展也被加密，ServerHello 之后的所有握手消息现在都被加密。

## Reduced Handshake

## 减少握手

TLS 1.2 supports numerous key exchange algorithms, cipher suites and digital signatures, including weak and vulnerable ones. Therefore, it requires more messages to perform a handshake and two network round trips.

TLS 1.2 支持多种密钥交换算法、密码套件和数字签名，包括弱和易受攻击的签名。因此，它需要更多的消息来执行一次握手和两次网络往返。

In contrast, the handshake in TLS 1.3 now requires only one round trip, with a simplified design and with all weak and vulnerable algorithms removed. 

相比之下，TLS 1.3 中的握手现在只需要一次往返，具有简化的设计，并删除了所有弱算法和易受攻击的算法。

In addition, it has a new feature called 0-RTT, or TLS early data, for the resumed handshake. This allows an application to include application data with its initial handshake message, instead of having to wait until the handshake completes.

此外，它还具有一个称为 0-RTT 或 TLS 早期数据的新功能，用于恢复握手。这允许应用程序在其初始握手消息中包含应用程序数据，而不必等到握手完成。

At Netflix, by the efficient resumption of the TLS session and careful use of 0-RTT for the streaming data, we can reduce the play delay.

在 Netflix，通过 TLS 会话的有效恢复和对流数据的谨慎使用 0-RTT，我们可以减少播放延迟。

![](https://miro.medium.com/max/60/1*dwfkh2K4bBsQHdP_rEyH7Q.png?q=20)

# A/B Testing Result

# A/B 测试结果

We were pretty confident that TLS 1.3 would bring us better security from the analysis of its protocol composition, but we did not know how it would perform in the context of streaming.

我们非常有信心 TLS 1.3 会从其协议组成的分析中为我们带来更好的安全性，但我们不知道它在流环境中的表现如何。

Since TLS 1.3's performance-related feature is the 0-RTT mode with the resumed handshake, our hypothesis is that TLS 1.3 would reduce play delay, as we are no longer required to wait for the handshake to finish and we can instead issue the HTTP request for media data and receive the HTTP response for media data earlier.

由于 TLS 1.3 的性能相关特性是恢复握手的 0-RTT 模式，我们的假设是 TLS 1.3 将减少播放延迟，因为我们不再需要等待握手完成，我们可以改为发出媒体数据的 HTTP 请求并更早地接收媒体数据的 HTTP 响应。

To see the actual performance of TLS 1.3 in the field, we performed an experiment with

为了看到TLS 1.3在现场的实际表现，我们进行了一个实验

- User accounts: half-million user accounts per cell.
- Device type: mid-performance device with Quad ARM core @ 1.7GHz.
- Control cell: TLS 1.2
- Treatment cell: TLS 1.3

- 用户帐户：每个单元有 50 万个用户帐户。
- 设备类型：具有 Quad ARM 内核 @ 1.7GHz 的中性能设备。
- 控制单元：TLS 1.2
- 治疗室：TLS 1.3

## Play Delay

## 播放延迟

Play Delay is defined by how long it takes for playback to start. Below are the results of the play delay measured in the experiment. The results imply that on slower or congested networks, which can be represented by the quantiles of at least 0.75, TLS 1.3 achieves the largest gains, with improvements across all network conditions.

播放延迟定义为播放开始所需的时间。以下是实验中测得的播放延迟结果。结果表明，在较慢或拥塞的网络上（可以用至少 0.75 的分位数表示），TLS 1.3 实现了最大的收益，并且在所有网络条件下都有改进。

![](https://miro.medium.com/max/60/1*So7N5ohUZYZnP2ZfgLyNEg.png?q=20)

Below is the time series median play delay graph for this mid-performance device in the field. It also shows that playback starts earlier with TLS 1.3.

下面是该领域中性能设备的时间序列中值播放延迟图。它还表明，使用 TLS 1.3 更早地开始播放。

![](https://miro.medium.com/max/60/1*6Hkq1ZCussPA1s-qtcuYOg.png?q=20)

## Media Rebuffer

## 媒体重新缓冲

At Netflix, we define a media rebuffer as a non-network originated rebuffer. It typically occurs when media data is not processed quickly enough by the device due to the high load on the CPU. Comparing the control cell with TLS 1.2, the experiment cell with TLS 1.3 showed about a 7.4% improvement in media rebuffers. This result implies that using TLS 1.3 with 0-RTT is more efficient and can reduce the CPU load.

在 Netflix，我们将媒体缓存定义为非网络发起的缓存。它通常发生在由于 CPU 上的高负载而导致设备没有足够快地处理媒体数据时。将对照单元与 TLS 1.2 进行比较，具有 TLS 1.3 的实验单元在媒体重新缓冲方面表现出约 7.4% 的改进。这个结果意味着使用带有 0-RTT 的 TLS 1.3 效率更高，并且可以减少 CPU 负载。

![](https://miro.medium.com/max/60/1*sZY-mXVAxBMFMp9ZvVl5Aw.png?q=20)

# Conclusion

#  结论

From the security analysis, we are confident that TLS 1.3 improves communication security over TLS 1.2. From the field test, we are confident that TLS 1.3 provides us a better streaming experience.

从安全性分析来看，我们相信 TLS 1.3 比 TLS 1.2 提高了通信安全性。从现场测试来看，我们相信 TLS 1.3 为我们提供了更好的流媒体体验。

At the time of writing this article, the Internet is experiencing higher than usual traffic and congestion. We believe saving even small amounts of data and round trips can be meaningful and even better if it also provides a more secure and efficient streaming experience.

在撰写本文时，互联网正在经历比平常更高的流量和拥塞。我们相信，如果还能提供更安全、更高效的流媒体体验，即使保存少量数据和往返行程也很有意义，甚至会更好。

Therefore, we have started deploying TLS 1.3 on newer consumer electronics devices and we are expecting even more devices to be deployed with TLS 1.3 capability in the near future.

因此，我们已开始在较新的消费电子设备上部署 TLS 1.3，我们预计在不久的将来将部署更多具有 TLS 1.3 功能的设备。

[**Netflix TechBlog**](https://netflixtechblog.com/?source=post_sidebar--------------------------post_sidebar-----------)

Learn about Netflix’s world class engineering efforts… 

了解 Netflix 世界一流的工程工作……

