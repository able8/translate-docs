# [How to audit SSH server and client config on Linux/Unix](https://www.cyberciti.biz/tips/how-to-audit-ssh-server-and-client-config-on-linux-unix.html)

# [如何在 Linux/Unix 上审计 SSH 服务器和客户端配置](https://www.cyberciti.biz/tips/how-to-audit-ssh-server-and-client-config-on-linux-unix.html)

                        Author: Vivek Gite                     Last updated: June 14, 2021                     [3 comments](https://www.cyberciti.biz/tips/how-to-audit-ssh-server-and-client-config-on-linux-unix.html#comments)

作者：Vivek Gite 最后更新：2021 年 6 月 14 日 [3 条评论](https://www.cyberciti.biz/tips/how-to-audit-ssh-server-and-client-config-on-linux-unix.html#comments)

All developers and Unix users know how  to use an ssh client. OpenSSH is a widely used implementation of Secure  Shell (SSH) Internet communication protocol. Back in the old days, Unix  folks used Telnet which was insecure. On the other hand, SSH allows  exchanging data using a secure channel between two hosts. Therefore,  every Linux and Unix server running cloud or at home needs an OpenSSH  server for management and IT automation. Regrettably, the popularity of  SSH servers and client also brings various security issues. I wrote  about “[Top 20 OpenSSH Server Best Security Practices](https://www.cyberciti.biz/tips/linux-unix-bsd-openssh-server-best-practices.html)” a long time ago. Today, I will talk about ssh server and client  auditing tools that anyone can use to the hardened standard SSH server  and client configuration for security issues.

所有开发人员和 Unix 用户都知道如何使用 ssh 客户端。 OpenSSH 是一种广泛使用的 Secure Shell (SSH) Internet 通信协议的实现。过去，Unix 人员使用不安全的 Telnet。另一方面，SSH 允许使用两台主机之间的安全通道交换数据。因此，每个运行在云中或在家中的 Linux 和 Unix 服务器都需要一个 OpenSSH 服务器来进行管理和 IT 自动化。遗憾的是，SSH 服务器和客户端的普及也带来了各种安全问题。很久以前我写过“[Top 20 OpenSSH Server Best Security Practices](https://www.cyberciti.biz/tips/linux-unix-bsd-openssh-server-best-practices.html)”。今天，我将讨论 ssh 服务器和客户端审计工具，任何人都可以使用这些工具来解决安全问题的强化标准 SSH 服务器和客户端配置。

## What is the ssh-audit tool?

## ssh-audit 工具是什么？

In simple words, ssh-audit is a tool for ssh server and client auditing. For example, you can use this tool:

简单来说，ssh-audit 是一个用于 ssh 服务器和客户端审计的工具。例如，您可以使用此工具：

1. Scan for OpenSSH server and client config for security issues
2. Make sure the correct and recommended algorithm is used by your Linux and Unix boxes
3. Check for OpenSSH banners and recognize device or software and operating system
4. Lookup for ssh key exchange, host-keys, encryption, and message authentication code algorithms
5. Alert developers and sysadmin about config issues, weak/legacy algorithms, and features used by SSH
6. Historical information from OpenSSH, Dropbear SSH, and libssh
7. Policy scans to ensure adherence to a hardened/standard configuration

1. 扫描 OpenSSH 服务器和客户端配置是否存在安全问题
2. 确保您的 Linux 和 Unix 机器使用正确和推荐的算法
3. 检查 OpenSSH 横幅并识别设备或软件和操作系统
4. 查找 ssh 密钥交换、主机密钥、加密和消息验证代码算法
5. 就配置问题、弱/遗留算法以及 SSH 使用的功能向开发人员和系统管理员发出警报
6. 来自 OpenSSH、Dropbear SSH 和 libssh 的历史信息
7. 策略扫描以确保遵守强化/标准配置

### Requirements for auditing SSH server and client config on Linux/Unix

### 在 Linux/Unix 上审核 SSH 服务器和客户端配置的要求

You need:

你需要：

- Linux, Windows, or Unix-like systems such as macOS, FreeBSD, and so on
- Python version 3.6 – 3.9
- No other dependencies

- Linux、Windows 或类 Unix 系统，例如 macOS、FreeBSD 等
- Python 版本 3.6 – 3.9
- 没有其他依赖

## How to install ssh-audit tool

## 如何安装ssh-audit工具

There are many ways to install such tools. Let us look into some popular options to employ as per their desktop and server operating  systems.

有很多方法可以安装此类工具。让我们看看根据他们的桌面和服务器操作系统使用的一些流行选项。

### Installing ssh-audit on Ubuntu Linux

### 在 Ubuntu Linux 上安装 ssh-audit

If you have snap enabled on your Linux system, run the following snap command:
  `sudo snap install ssh-audit`

如果您在 Linux 系统上启用了 snap，请运行以下 snap 命令：
 `sudo snap install ssh-audit`

### FreeBSD install ssh-audit

### FreeBSD 安装 ssh-audit

Search it and install using the pkg command:
  `$ pkg search ssh-audit # note down output from above command and apply it # $ sudo pkg install py37-ssh-audit`
  ![img](data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjM3NyIgd2lkdGg9IjU5OSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB2ZXJzaW9uPSIxLjEiLz4=)![How to install ssh audit on FreeBSD](https://www.cyberciti.biz/tips/wp-content/uploads/2021/06/How-to-install-ssh-audit-on-FreeBSD.png)

使用 pkg 命令搜索并安装：
 `$ pkg search ssh-audit # 记下上面命令的输出并应用它# $ sudo pkg install py37-ssh-audit`
 /How-to-install-ssh-audit-on-FreeBSD.png)

### A note about macOS

### 关于 macOS 的说明

First, enable/[install Homebrew on macOS to use the brew package manager](https://www.cyberciti.biz/faq/how-to-install-homebrew-on-macos-package-manager/) and then type:
  `brew install ssh-audit`

首先，启用/[在 macOS 上安装 Homebrew 以使用 brew 包管理器](https://www.cyberciti.biz/faq/how-to-install-homebrew-on-macos-package-manager/) 然后输入：
 `brew install ssh-audit`

### Other methods

### 其他方法

Of course, we can install it from PyPI too. For instance:
  `pip3 install ssh-audit`
  To install from Dockerhub:
  `docker pull positronsecurity/ssh-audit`
  Then run it as follows:

当然，我们也可以从 PyPI 安装它。例如：
 `pip3 安装 ssh-audit`
 从 Dockerhub 安装：
 `docker pull positronsecurity/ssh-audit`
 然后按如下方式运行它：

Patreon supporters only guides 🤓

Patreon 支持者只做向导🤓

- No ads and tracking

- 没有广告和跟踪

- In-depth guides for developers and sysadmins at [Opensourceflare](https://www.opensourceflare.com)✨

- [Opensourceflare](https://www.opensourceflare.com) 上为开发人员和系统管理员提供的深入指南✨

- Join my Patreon to support independent content creators and start reading latest guides:

- 加入我的 Patreon 以支持独立内容创作者并开始阅读最新指南：

- - [How to set up Redis sentinel cluster on Ubuntu or Debian Linux](https://www.opensourceflare.com/how-to-set-up-redis-sentinel-cluster-on-ubuntu-or-debian-linux/)
   - [How To Set Up SSH Keys With YubiKey as two-factor authentication (U2F/FIDO2)](https://www.opensourceflare.com/how-to-set-up-ssh-keys-with-yubikey-as-two-factor-authentication-u2f-fido2/)
   - [How to set up Mariadb Galera cluster on Ubuntu or Debian Linux](https://www.opensourceflare.com/how-to-set-up-mariadb-galera-cluster-on-ubuntu-or-debian-linux/)
   - [A podman tutorial for beginners – part I (run Linux containers without Docker and in daemonless mode)](https://www.opensourceflare.com/a-podman-tutorial-for-beginners-part-i/)

- - [如何在 Ubuntu 或 Debian Linux 上设置 Redis sentinel 集群](https://www.opensourceflare.com/how-to-set-up-redis-sentinel-cluster-on-ubuntu-or-debian-linux/)
  - [如何使用 YubiKey 作为双因素身份验证 (U2F/FIDO2) 设置 SSH 密钥](https://www.opensourceflare.com/how-to-set-up-ssh-keys-with-yubikey-as-两因素身份验证-u2f-fido2/)
  - [如何在 Ubuntu 或 Debian Linux 上设置 Mariadb Galera 集群](https://www.opensourceflare.com/how-to-set-up-mariadb-galera-cluster-on-ubuntu-or-debian-linux/)
  - [适用于初学者的 podman 教程 - 第一部分（在没有 Docker 的情况下以无守护进程模式运行 Linux 容器）](https://www.opensourceflare.com/a-podman-tutorial-for-beginners-part-i/)

[Join **Patreon** ➔](https://www.patreon.com/nixcraft)

[加入**Patreon** ➔](https://www.patreon.com/nixcraft)

```
docker run -it -p 2222:2222 positronsecurity/ssh-audit {ssh-server-ip}
docker run -it -p 2222:2222 positronsecurity/ssh-audit 192.168.2.17
```

## Usage 

##  用法

The syntax is simple

语法很简单

```
ssh-audit [options] <ssh-server-host-ip>
ssh-audit 192.168.2.17
# set SSH port to 222
ssh-audit 192.168.2.254:222
```

For basic SSH server auditing, type:
  `ssh-audit router.home.sweet ssh-audit 192.168.2.254`

对于基本的 SSH 服务器审计，键入：
 `ssh-audit router.home.sweet ssh-audit 192.168.2.254`

```
# general
(gen) banner: SSH-2.0-OpenSSH_8.6 FreeBSD-openssh-portable-8.6.p1,1
(gen) software: OpenSSH 8.6 running on FreeBSD
(gen) compatibility: OpenSSH 7.4+, Dropbear SSH 2018.76+
(gen) compression: enabled (zlib@openssh.com)
 
# key exchange algorithms
(kex) curve25519-sha256                     -- [info] available since OpenSSH 7.4, Dropbear SSH 2018.76
(kex) curve25519-sha256@libssh.org          -- [info] available since OpenSSH 6.5, Dropbear SSH 2013.62
(kex) ecdh-sha2-nistp256                    -- [fail] using weak elliptic curves
                                            `- [info] available since OpenSSH 5.7, Dropbear SSH 2013.62
(kex) ecdh-sha2-nistp384                    -- [fail] using weak elliptic curves
                                            `- [info] available since OpenSSH 5.7, Dropbear SSH 2013.62
(kex) ecdh-sha2-nistp521                    -- [fail] using weak elliptic curves
                                            `- [info] available since OpenSSH 5.7, Dropbear SSH 2013.62
(kex) diffie-hellman-group-exchange-sha256 (2048-bit) -- [info] available since OpenSSH 4.4
(kex) diffie-hellman-group16-sha512         -- [info] available since OpenSSH 7.3, Dropbear SSH 2016.73
(kex) diffie-hellman-group18-sha512         -- [info] available since OpenSSH 7.3
(kex) diffie-hellman-group14-sha256         -- [info] available since OpenSSH 7.3, Dropbear SSH 2016.73
 
# host-key algorithms
(key) rsa-sha2-512 (3072-bit)               -- [info] available since OpenSSH 7.2
(key) rsa-sha2-256 (3072-bit)               -- [info] available since OpenSSH 7.2
(key) ssh-rsa (3072-bit)                    -- [fail] using weak hashing algorithm
                                            `- [info] available since OpenSSH 2.5.0, Dropbear SSH 0.28
                                            `- [info] a future deprecation notice has been issued in OpenSSH 8.2: https://www.openssh.com/txt/release-8.2
(key) ecdsa-sha2-nistp256                   -- [fail] using weak elliptic curves
                                            `- [warn] using weak random number generator could reveal the key
                                            `- [info] available since OpenSSH 5.7, Dropbear SSH 2013.62
(key) ssh-ed25519                           -- [info] available since OpenSSH 6.5
 
# encryption algorithms (ciphers)
(enc) chacha20-poly1305@openssh.com         -- [info] available since OpenSSH 6.5
                                            `- [info] default cipher since OpenSSH 6.9.
(enc) aes128-ctr                            -- [info] available since OpenSSH 3.7, Dropbear SSH 0.52
(enc) aes192-ctr                            -- [info] available since OpenSSH 3.7
(enc) aes256-ctr                            -- [info] available since OpenSSH 3.7, Dropbear SSH 0.52
(enc) aes128-gcm@openssh.com                -- [info] available since OpenSSH 6.2
(enc) aes256-gcm@openssh.com                -- [info] available since OpenSSH 6.2
 
# message authentication code algorithms
(mac) umac-64-etm@openssh.com               -- [warn] using small 64-bit tag size
                                            `- [info] available since OpenSSH 6.2
(mac) umac-128-etm@openssh.com              -- [info] available since OpenSSH 6.2
(mac) hmac-sha2-256-etm@openssh.com         -- [info] available since OpenSSH 6.2
(mac) hmac-sha2-512-etm@openssh.com         -- [info] available since OpenSSH 6.2
(mac) hmac-sha1-etm@openssh.com             -- [warn] using weak hashing algorithm
                                            `- [info] available since OpenSSH 6.2
(mac) umac-64@openssh.com                   -- [warn] using encrypt-and-MAC mode
                                            `- [warn] using small 64-bit tag size
                                            `- [info] available since OpenSSH 4.7
(mac) umac-128@openssh.com                  -- [warn] using encrypt-and-MAC mode
                                            `- [info] available since OpenSSH 6.2
(mac) hmac-sha2-256                         -- [warn] using encrypt-and-MAC mode
                                            `- [info] available since OpenSSH 5.9, Dropbear SSH 2013.56
(mac) hmac-sha2-512                         -- [warn] using encrypt-and-MAC mode
                                            `- [info] available since OpenSSH 5.9, Dropbear SSH 2013.56
(mac) hmac-sha1                             -- [warn] using encrypt-and-MAC mode
                                            `- [warn] using weak hashing algorithm
                                            `- [info] available since OpenSSH 2.1.0, Dropbear SSH 0.28
 
# fingerprints
(fin) ssh-ed25519: SHA256:JGOsGxcCjN5Ej+8FSYK5bo4L23W66wSgQof8xpASplc
(fin) ssh-rsa: SHA256:aM8yrCKPlLDd5kRwSS7JNj7Kho6k9JEs5aFv/VTGMRA
 
# algorithm recommendations (for OpenSSH 8.6)
(rec) -ecdh-sha2-nistp256                   -- kex algorithm to remove
(rec) -ecdh-sha2-nistp384                   -- kex algorithm to remove
(rec) -ecdh-sha2-nistp521                   -- kex algorithm to remove
(rec) -ecdsa-sha2-nistp256                  -- key algorithm to remove
(rec) -ssh-rsa                              -- key algorithm to remove
(rec) -hmac-sha1                            -- mac algorithm to remove
(rec) -hmac-sha1-etm@openssh.com            -- mac algorithm to remove
(rec) -hmac-sha2-256                        -- mac algorithm to remove
(rec) -hmac-sha2-512                        -- mac algorithm to remove
(rec) -umac-128@openssh.com                 -- mac algorithm to remove
(rec) -umac-64-etm@openssh.com              -- mac algorithm to remove
(rec) -umac-64@openssh.com                  -- mac algorithm to remove
 
# additional info
(nfo) For hardening guides on common OSes, please see: <https://www.ssh-audit.com/hardening_guides.html>
```

### Auditing many servers

### 审计多台服务器

Want to do a standard audit against many servers hosted in cloud? We need to [create a new text file](https://www.cyberciti.biz/faq/create-a-file-in-linux-using-the-bash-shell-terminal/):

想要对托管在云中的许多服务器进行标准审计吗？我们需要[创建一个新的文本文件](https://www.cyberciti.biz/faq/create-a-file-in-linux-using-the-bash-shell-terminal/)：

```
cat > ec2-server.txt
aws-server1
aws-server2
54.56.5.5
```

Then, I would run:
  `ssh-audit -T ec2-server.txt`

然后，我会运行：
 `ssh-audit -T ec2-server.txt`

### Auditing client config

### 审计客户端配置

To audit a client configuration, type:

要审核客户端配置，请键入：

```
ssh-audit -c
# client listener on port 4123
ssh-audit -c -p 4123
```

### How to run a policy audit against a server

### 如何对服务器运行策略审计

To list a policy run:
  `ssh-audit -L`
  Then I will see a list as follows:

要列出策略运行：
 `ssh-audit -L`
 然后我会看到一个列表如下：

```
Server policies:

  * "Hardened OpenSSH Server v7.7 (version 1)"
  * "Hardened OpenSSH Server v7.8 (version 1)"
  * "Hardened OpenSSH Server v7.9 (version 1)"
  * "Hardened OpenSSH Server v8.0 (version 1)"
  * "Hardened OpenSSH Server v8.1 (version 1)"
  * "Hardened OpenSSH Server v8.2 (version 1)"
  * "Hardened OpenSSH Server v8.3 (version 1)"
  * "Hardened OpenSSH Server v8.4 (version 1)"
  * "Hardened OpenSSH Server v8.5 (version 1)"
  * "Hardened Ubuntu Server 16.04 LTS (version 1)"
  * "Hardened Ubuntu Server 18.04 LTS (version 1)"
  * "Hardened Ubuntu Server 20.04 LTS (version 1)"

Client policies:

  * "Hardened Ubuntu Client 16.04 LTS (version 2)"
  * "Hardened Ubuntu Client 18.04 LTS (version 2)"
  * "Hardened Ubuntu Client 20.04 LTS (version 2)"

Hint: Use -P and provide the full name of a policy to run a policy scan with.
```

To run a policy audit against a server named ln.ncbz01, type:
  `ssh-audit -P 'Hardened Ubuntu Server 20.04 LTS (version 1)' ln.ncbz01`
  ![img](data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjM0MyIgd2lkdGg9IjU5OSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB2ZXJzaW9uPSIxLjEiLz4=)![ssh-audit policy scan output](https://www.cyberciti.biz/tips/wp-content/uploads/2021/06/ssh-audit-policy-scan-output.png)

要对名为 ln.ncbz01 的服务器运行策略审计，请键入：
 `ssh-audit -P '强化 Ubuntu Server 20.04 LTS（版本 1）' ln.ncbz01`
 ssh-audit-policy-scan-output.png)

## Fixing Ubuntu 20.04 LTS Server failed audit

## 修复 Ubuntu 20.04 LTS 服务器审核失败

We need to run the following commands as the root user:
  Re-generate the RSA and ED25519 keys:
```
# rm /etc/ssh/ssh_host_* # ssh-keygen -t rsa -b 4096 -f /etc/ssh/ssh_host_rsa_key -N "" # ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N ""
```

 
Remove small Diffie-Hellman moduli
```
# awk '$5 >= 3071' /etc/ssh/moduli > /etc/ssh/moduli.safe # mv /etc/ssh/moduli.safe /etc/ssh/moduli
```

Enable the RSA and ED25519 keys

启用 RSA 和 ED25519 密钥

```
# sed -i 's/^\#HostKey \/etc\/ssh\/ssh_host_\(rsa\|ed25519\)_key$/HostKey \/etc\/ssh\/ssh_host_\1_key/g' /etc/ssh/sshd_config
```

Restrict supported key exchange, cipher, and MAC algorithms

限制支持的密钥交换、密码和 MAC 算法

```bash
# echo -e "\n# Restrict key exchange, cipher, and MAC algorithms,  as per sshaudit.com\n# hardening guide.\nKexAlgorithms  curve25519-sha256,curve25519-sha256@libssh.org,diffie-hellman-group16-sha512,diffie-hellman-group18-sha512,diffie-hellman-group-exchange-sha256\nCiphers  chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr\nMACs  hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,umac-128-etm@openssh.com\nHostKeyAlgorithms  ssh-ed25519,ssh-ed25519-cert-v01@openssh.com,sk-ssh-ed25519@openssh.com,sk-ssh-ed25519-cert-v01@openssh.com,rsa-sha2-256,rsa-sha2-512,rsa-sha2-256-cert-v01@openssh.com,rsa-sha2-512-cert-v01@openssh.com" > /etc/ssh/sshd_config.d/ssh-audit_hardening.conf
```

 
Finally, [restart ssh service on Ubuntu Linux](https://www.cyberciti.biz/faq/howto-start-stop-ssh-server/):

最后， [在 Ubuntu Linux 上重启 ssh 服务](https://www.cyberciti.biz/faq/howto-start-stop-ssh-server/)：

```
# systemctl restart ssh
```

Now verify audit again:

现在再次验证审计：

```
$ ssh-audit -P 'Hardened Ubuntu Server 20.04 LTS (version 1)' ln.ncbz01
```

## Summing up

##  加起来

I think it is an excellent tool for sysadmin and security folks for  auditing ssh server and client on your Linux and Unix box and nice  addition to my “[Top 20 OpenSSH Server Best Security Practices](https://www.cyberciti.biz/tips/linux-unix-bsd-openssh-server-best-practices.html)” post. Make sure you [check out the project home page](https://github.com/jtesta/ssh-audit). Let me know if you found this as a valuable tool in the comment section below.

我认为它是系统管理员和安全人员在 Linux 和 Unix 机器上审计 ssh 服务器和客户端的绝佳工具，也是我的“[Top 20 OpenSSH Server Best Security Practices](https://www.cyberciti.biz/提示/linux-unix-bsd-openssh-server-best-practices.html)”的帖子。确保您[查看项目主页](https://github.com/jtesta/ssh-audit)。如果您在下面的评论部分发现这是一个有价值的工具，请告诉我。

🐧 Get the latest tutorials on Linux, Open Source & DevOps via
  **[RSS feed ➔](https://www.cyberciti.biz/atom/atom.xml) [Weekly email newsletter ➔](https://newsletter.cyberciti.biz/subscription?f=1ojtmiv8892KQzyMsTF4YPr1pPSAhX2rq7Qfe5DiHMgXwKo892di4MTWyOdd976343rcNR6LhdG1f7k9H8929kMNMdWu3g)** 

🐧 通过以下方式获取有关 Linux、开源和 DevOps 的最新教程
 ** [RSS馈送➔](https://www.cyberciti.biz/atom/atom.xml)每周电子邮件通讯➔](https://newsletter.cyberciti.biz/subscription?f=1ojtmiv8892KQzyMsTF4YPr1pPSAhX2rq7Qfe5DiHMgXwKo892di4MTWyOdd976343rcNR6LhdG1f7k9H8929kMNMdWu3g)**

